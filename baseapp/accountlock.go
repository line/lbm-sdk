package baseapp

import (
	"encoding/binary"
	"sort"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NOTE should 1 <= sampleBytes <= 4. If modify it, you should revise `getAddressKey()` as well
const sampleBytes = 2

type AccountLock struct {
	accMtx [1 << (sampleBytes * 8)]sync.Mutex
}

func (al *AccountLock) Lock(ctx sdk.Context, tx sdk.Tx) []uint32 {
	if !ctx.IsCheckTx() || ctx.IsReCheckTx() {
		return nil
	}

	signers := getSigners(tx)
	accKeys := getUniqSortedAddressKey(signers)

	for _, key := range accKeys {
		al.accMtx[key].Lock()
	}

	return accKeys
}

func (al *AccountLock) Unlock(accKeys []uint32) {
	// NOTE reverse order
	for i, length := 0, len(accKeys); i < length; i++ {
		key := accKeys[length-1-i]
		al.accMtx[key].Unlock()
	}
}

func getSigners(tx sdk.Tx) []sdk.AccAddress {
	seen := map[string]bool{}
	var signers []sdk.AccAddress
	for _, msg := range tx.GetMsgs() {
		for _, addr := range msg.GetSigners() {
			if !seen[addr.String()] {
				signers = append(signers, addr)
				seen[addr.String()] = true
			}
		}
	}
	return signers
}

func getUniqSortedAddressKey(addrs []sdk.AccAddress) []uint32 {
	accKeys := make([]uint32, 0, len(addrs))
	for _, addr := range addrs {
		accKeys = append(accKeys, getAddressKey(addr))
	}

	accKeys = uniq(accKeys)
	sort.Slice(accKeys, func(i, j int) bool {
		return i < j
	})

	return accKeys
}

func getAddressKey(addr sdk.AccAddress) uint32 {
	return uint32(binary.BigEndian.Uint16(addr))
}

func uniq(u []uint32) []uint32 {
	seen := map[uint32]bool{}
	var ret []uint32
	for _, v := range u {
		if !seen[v] {
			ret = append(ret, v)
			seen[v] = true
		}
	}
	return ret
}
