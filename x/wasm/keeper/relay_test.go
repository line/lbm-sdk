package keeper

import (
	"encoding/json"
	"errors"
	"math"
	"testing"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/wasm/keeper/wasmtesting"
	"github.com/line/lbm-sdk/x/wasm/types"
	wasmvm "github.com/line/wasmvm"
	wasmvmtypes "github.com/line/wasmvm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOnOpenChannel(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr sdk.AccAddress
		contractGas  sdk.Gas
		contractErr  error
		expGas       uint64
		expErr       bool
	}{
		"consume contract gas": {
			contractAddr: example.Contract,
			contractGas:  myContractGas,
			expGas:       myContractGas,
		},
		"consume max gas": {
			contractAddr: example.Contract,
			contractGas:  math.MaxUint64 / types.DefaultGasMultiplier,
			expGas:       math.MaxUint64 / types.DefaultGasMultiplier,
		},
		"consume gas on error": {
			contractAddr: example.Contract,
			contractGas:  myContractGas,
			contractErr:  errors.New("test, ignore"),
			expErr:       true,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			myChannel := wasmvmtypes.IBCChannel{Version: "my test channel"}
			m.IBCChannelOpenFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, channel wasmvmtypes.IBCChannel, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (uint64, error) {
				assert.Equal(t, myChannel, channel)
				return spec.contractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			before := ctx.GasMeter().GasConsumed()

			// when
			err := keepers.WasmKeeper.OnOpenChannel(ctx, spec.contractAddr, myChannel)

			// then
			if spec.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expGas, ctx.GasMeter().GasConsumed()-before-storageCosts)
		})
	}
}

func TestOnConnectChannel(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr          sdk.AccAddress
		contractResp          *wasmvmtypes.IBCBasicResponse
		contractErr           error
		overwriteMessenger    *wasmtesting.MockMessageHandler
		expContractGas        sdk.Gas
		expErr                bool
		expContractEventAttrs int
		expNoEvents           bool
	}{
		"consume contract gas": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{},
		},
		"consume gas on error, ignore events + messages": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			contractErr: errors.New("test, ignore"),
			expErr:      true,
			expNoEvents: true,
		},
		"dispatch contract messages on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
			},
		},
		"emit contract events on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			expContractEventAttrs: 1,
		},
		"messenger errors returned, events stored": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			overwriteMessenger:    wasmtesting.NewErroringMessageHandler(),
			expErr:                true,
			expContractEventAttrs: 1,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
			expNoEvents:  true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			myChannel := wasmvmtypes.IBCChannel{Version: "my test channel"}
			m.IBCChannelConnectFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, channel wasmvmtypes.IBCChannel, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
				assert.Equal(t, channel, myChannel)
				return spec.contractResp, myContractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			before := ctx.GasMeter().GasConsumed()
			msger, capturedMsgs := wasmtesting.NewCapturingMessageHandler()
			*messenger = *msger
			if spec.overwriteMessenger != nil {
				*messenger = *spec.overwriteMessenger
			}

			// when
			err := keepers.WasmKeeper.OnConnectChannel(ctx, spec.contractAddr, myChannel)

			// then
			events := ctx.EventManager().Events()
			if spec.expErr {
				require.Error(t, err)
				assert.Empty(t, capturedMsgs) // no messages captured on error
				if spec.expNoEvents {
					require.Len(t, events, 0)
				} else {
					require.Len(t, events, 1)
					assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
				}
				return
			}
			require.NoError(t, err)

			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expContractGas, ctx.GasMeter().GasConsumed()-before-storageCosts)

			// verify msgs dispatched
			assert.Equal(t, spec.contractResp.Messages, *capturedMsgs)
			assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
		})
	}
}

func TestOnCloseChannel(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr          sdk.AccAddress
		contractResp          *wasmvmtypes.IBCBasicResponse
		contractErr           error
		overwriteMessenger    *wasmtesting.MockMessageHandler
		expContractGas        sdk.Gas
		expErr                bool
		expContractEventAttrs int
		expNoEvents           bool
	}{
		"consume contract gas": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{},
		},
		"consume gas on error, ignore events + messages": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			contractErr: errors.New("test, ignore"),
			expErr:      true,
			expNoEvents: true,
		},
		"dispatch contract messages on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
			},
		},
		"emit contract events on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			expContractEventAttrs: 1,
		},
		"messenger errors returned, events stored": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			overwriteMessenger:    wasmtesting.NewErroringMessageHandler(),
			expErr:                true,
			expContractEventAttrs: 1,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
			expNoEvents:  true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			myChannel := wasmvmtypes.IBCChannel{Version: "my test channel"}
			m.IBCChannelCloseFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, channel wasmvmtypes.IBCChannel, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
				assert.Equal(t, channel, myChannel)
				return spec.contractResp, myContractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			before := ctx.GasMeter().GasConsumed()
			msger, capturedMsgs := wasmtesting.NewCapturingMessageHandler()
			*messenger = *msger

			if spec.overwriteMessenger != nil {
				*messenger = *spec.overwriteMessenger
			}

			// when
			err := keepers.WasmKeeper.OnCloseChannel(ctx, spec.contractAddr, myChannel)

			// then
			events := ctx.EventManager().Events()
			if spec.expErr {
				require.Error(t, err)
				assert.Empty(t, capturedMsgs) // no messages captured on error
				if spec.expNoEvents {
					require.Len(t, events, 0)
				} else {
					require.Len(t, events, 1)
					assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
				}
				return
			}
			require.NoError(t, err)
			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expContractGas, ctx.GasMeter().GasConsumed()-before-storageCosts)

			// verify msgs dispatched
			assert.Equal(t, spec.contractResp.Messages, *capturedMsgs)
			require.Len(t, events, 1)
			assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
		})
	}
}

func TestOnRecvPacket(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr          sdk.AccAddress
		contractResp          *wasmvmtypes.IBCReceiveResponse
		contractErr           error
		overwriteMessenger    *wasmtesting.MockMessageHandler
		expContractGas        sdk.Gas
		expErr                bool
		expContractEventAttrs int
		expNoEvents           bool
	}{
		"consume contract gas": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCReceiveResponse{
				Acknowledgement: []byte("myAck"),
			},
		},
		"can return empty ack": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCReceiveResponse{},
		},
		"consume gas on error, ignore events + messages": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCReceiveResponse{
				Acknowledgement: []byte("myAck"),
				Messages:        []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}},
				Attributes:      []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			contractErr: errors.New("test, ignore"),
			expErr:      true,
			expNoEvents: true,
		},
		"dispatch contract messages on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCReceiveResponse{
				Acknowledgement: []byte("myAck"),
				Messages:        []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
			},
		},
		"emit contract events on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp: &wasmvmtypes.IBCReceiveResponse{
				Acknowledgement: []byte("myAck"),
				Attributes:      []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			expContractEventAttrs: 1,
		},
		"messenger errors returned, events stored": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp: &wasmvmtypes.IBCReceiveResponse{
				Acknowledgement: []byte("myAck"),
				Messages:        []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
				Attributes:      []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			overwriteMessenger:    wasmtesting.NewErroringMessageHandler(),
			expErr:                true,
			expContractEventAttrs: 1,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
			expNoEvents:  true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			myPacket := wasmvmtypes.IBCPacket{Data: []byte("my data")}

			m.IBCPacketReceiveFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, packet wasmvmtypes.IBCPacket, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (*wasmvmtypes.IBCReceiveResponse, uint64, error) {
				assert.Equal(t, myPacket, packet)
				return spec.contractResp, myContractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			before := ctx.GasMeter().GasConsumed()

			msger, capturedMsgs := wasmtesting.NewCapturingMessageHandler()
			*messenger = *msger

			if spec.overwriteMessenger != nil {
				*messenger = *spec.overwriteMessenger
			}

			// when
			gotAck, err := keepers.WasmKeeper.OnRecvPacket(ctx, spec.contractAddr, myPacket)

			// then
			events := ctx.EventManager().Events()
			if spec.expErr {
				require.Error(t, err)
				assert.Empty(t, capturedMsgs) // no messages captured on error
				if spec.expNoEvents {
					require.Len(t, events, 0)
				} else {
					require.Len(t, events, 1)
					assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
				}
				return
			}
			require.NoError(t, err)
			require.Equal(t, spec.contractResp.Acknowledgement, gotAck)

			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expContractGas, ctx.GasMeter().GasConsumed()-before-storageCosts)

			// verify msgs dispatched
			assert.Equal(t, spec.contractResp.Messages, *capturedMsgs)
			require.Len(t, events, 1)
			assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
		})
	}
}

func TestOnAckPacket(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr          sdk.AccAddress
		contractResp          *wasmvmtypes.IBCBasicResponse
		contractErr           error
		overwriteMessenger    *wasmtesting.MockMessageHandler
		expContractGas        sdk.Gas
		expErr                bool
		expContractEventAttrs int
		expNoEvents           bool
	}{
		"consume contract gas": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{},
		},
		"consume gas on error, ignore events + messages": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			contractErr: errors.New("test, ignore"),
			expErr:      true,
			expNoEvents: true,
		},
		"dispatch contract messages on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
			},
		},
		"emit contract events on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp:   &wasmvmtypes.IBCBasicResponse{
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			expContractEventAttrs: 1,
		},
		"messenger errors returned, events stored": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			overwriteMessenger:    wasmtesting.NewErroringMessageHandler(),
			expErr:                true,
			expContractEventAttrs: 1,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
			expNoEvents:  true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {

			myAck := wasmvmtypes.IBCAcknowledgement{Acknowledgement: []byte("myAck")}
			m.IBCPacketAckFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, ack wasmvmtypes.IBCAcknowledgement, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
				assert.Equal(t, myAck, ack)
				return spec.contractResp, myContractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			before := ctx.GasMeter().GasConsumed()
			msger, capturedMsgs := wasmtesting.NewCapturingMessageHandler()
			*messenger = *msger

			if spec.overwriteMessenger != nil {
				*messenger = *spec.overwriteMessenger
			}

			// when
			err := keepers.WasmKeeper.OnAckPacket(ctx, spec.contractAddr, myAck)

			// then
			events := ctx.EventManager().Events()
			if spec.expErr {
				require.Error(t, err)
				assert.Empty(t, capturedMsgs) // no messages captured on error
				if spec.expNoEvents {
					require.Len(t, events, 0)
				} else {
					require.Len(t, events, 1)
					assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
				}
				return
			}
			require.NoError(t, err)

			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expContractGas, ctx.GasMeter().GasConsumed()-before-storageCosts)

			// verify msgs dispatched
			assert.Equal(t, spec.contractResp.Messages, *capturedMsgs)
			require.Len(t, events, 1)
			assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
		})
	}
}

func TestOnTimeoutPacket(t *testing.T) {
	var m wasmtesting.MockWasmer
	wasmtesting.MakeIBCInstantiable(&m)
	var messenger = &wasmtesting.MockMessageHandler{}
	parentCtx, keepers := CreateTestInput(t, false, SupportedFeatures, nil, nil, WithMessageHandler(messenger))
	example := SeedNewContractInstance(t, parentCtx, keepers, &m)
	const myContractGas = 40

	specs := map[string]struct {
		contractAddr          sdk.AccAddress
		contractResp          *wasmvmtypes.IBCBasicResponse
		contractErr           error
		overwriteMessenger    *wasmtesting.MockMessageHandler
		expContractGas        sdk.Gas
		expErr                bool
		expContractEventAttrs int
		expNoEvents           bool
	}{
		"consume contract gas": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp:   &wasmvmtypes.IBCBasicResponse{},
		},
		"consume gas on error, ignore events + messages": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			contractErr: errors.New("test, ignore"),
			expErr:      true,
			expNoEvents: true,
		},
		"dispatch contract messages on success": {
			contractAddr: example.Contract,
			expContractGas: myContractGas,

			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages: []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
			},
		},
		"emit contract events on success": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			expContractEventAttrs: 1,
		},
		"messenger errors returned, events stored": {
			contractAddr:   example.Contract,
			expContractGas: myContractGas + 10,
			contractResp: &wasmvmtypes.IBCBasicResponse{
				Messages:   []wasmvmtypes.CosmosMsg{{Bank: &wasmvmtypes.BankMsg{}}, {Custom: json.RawMessage(`{"foo":"bar"}`)}},
				Attributes: []wasmvmtypes.EventAttribute{{Key: "Foo", Value: "Bar"}},
			},
			overwriteMessenger:    wasmtesting.NewErroringMessageHandler(),
			expErr:                true,
			expContractEventAttrs: 1,
		},
		"unknown contract address": {
			contractAddr: RandomAccountAddress(t),
			expErr:       true,
			expNoEvents:  true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			myPacket := wasmvmtypes.IBCPacket{Data: []byte("my test packet")}
			m.IBCPacketTimeoutFn = func(codeID wasmvm.Checksum, env wasmvmtypes.Env, packet wasmvmtypes.IBCPacket, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
				assert.Equal(t, myPacket, packet)
				return spec.contractResp, myContractGas * types.DefaultGasMultiplier, spec.contractErr
			}

			ctx, _ := parentCtx.CacheContext()
			before := ctx.GasMeter().GasConsumed()
			msger, capturedMsgs := wasmtesting.NewCapturingMessageHandler()
			*messenger = *msger

			if spec.overwriteMessenger != nil {
				*messenger = *spec.overwriteMessenger
			}

			// when
			err := keepers.WasmKeeper.OnTimeoutPacket(ctx, spec.contractAddr, myPacket)

			// then
			events := ctx.EventManager().Events()
			if spec.expErr {
				require.Error(t, err)
				assert.Empty(t, capturedMsgs) // no messages captured on error
				if spec.expNoEvents {
					require.Len(t, events, 0)
				} else {
					require.Len(t, events, 1)
					assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
				}
				return
			}
			require.NoError(t, err)
			// verify gas consumed
			const storageCosts = sdk.Gas(0xa8b)
			assert.Equal(t, spec.expContractGas, ctx.GasMeter().GasConsumed()-before-storageCosts)

			// verify msgs dispatched
			assert.Equal(t, spec.contractResp.Messages, *capturedMsgs)
			require.Len(t, events, 1)
			assert.Len(t, events[0].Attributes, 1+spec.expContractEventAttrs)
		})
	}
}
