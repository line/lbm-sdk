package keeper

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

const (
	defaultName       = "name"
	defaultSymbol     = "token001"
	defaultSymbol2    = "token002"
	defaultTokenURI   = "token-uri"
	defaultDecimals   = 6
	defaultAmount     = 1000
	defaultTokenType  = "1001"
	defaultTokenType2 = "1002"
	defaultTokenType3 = "1003"
	defaultTokenType4 = "1004"
	defaultTokenID1   = defaultTokenType + "0001"
	defaultTokenID2   = defaultTokenType + "0002"
	defaultTokenID3   = defaultTokenType + "0003"
	defaultTokenID4   = defaultTokenType + "0004"
	defaultTokenID5   = defaultTokenType + "0005"
	defaultTokenID6   = defaultTokenType + "0006"
	defaultTokenID8   = defaultTokenType + "0008"
	defaultTokenIDFT  = "00010000"
	defaultTokenIDFT2 = "00020000"
	defaultTokenIDFT3 = "00030000"
	defaultTokenIDFT4 = "00040000"
	defaultTokenIDFT5 = "00050000"
)

var (
	addr1 = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	addr2 = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	addr3 = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
)

var (
	ms     store.CommitMultiStore
	ctx    sdk.Context
	keeper Keeper
)

func setup() {
	println("setup")
	ctx, ms, keeper = TestKeeper()
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	os.Exit(ret)
}

func cacheKeeper() sdk.Context {
	msCache := ms.CacheMultiStore()
	return ctx.WithMultiStore(msCache)
}

func TestKeeper_MarshalJSONLogger(t *testing.T) {
	ctx := cacheKeeper()
	dummy := struct {
		Key   string
		Value string
	}{
		Key:   "key",
		Value: "value",
	}
	bz, err := keeper.MarshalJSON(dummy)
	require.NoError(t, err)

	dummy2 := struct {
		Key   string
		Value string
	}{}

	err = keeper.UnmarshalJSON(bz, &dummy2)
	require.NoError(t, err)
	require.Equal(t, dummy.Key, dummy2.Key)
	require.Equal(t, dummy.Value, dummy2.Value)
	logger := keeper.Logger(ctx)
	logger.Info("test", dummy, dummy2)
}

func prepareCollectionTokens(ctx sdk.Context, t *testing.T) {
	// prepare collection
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(defaultSymbol, "name"), addr1))

	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(defaultSymbol2, "name"), addr1))

	// issue 6 tokens
	// token1 = symbol1id1 by addr1
	// token2 = symbol1id2 by addr1
	// token3 = symbol1id3 by addr1
	// token4 = symbol1id4 by addr1
	// token5 = symbol2id5 by addr1
	// token6 = symbol1id6 by addr2
	// token7 = symbol1 by addr1
	require.NoError(t, keeper.IssueNFT(ctx, defaultSymbol, types.NewBaseTokenType(defaultSymbol, defaultTokenType, defaultName), addr1))
	require.NoError(t, keeper.IssueNFT(ctx, defaultSymbol2, types.NewBaseTokenType(defaultSymbol2, defaultTokenType, defaultName), addr1))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol, addr1, types.NewNFT(defaultSymbol, defaultTokenID1, defaultName, defaultTokenURI, addr1)))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol, addr1, types.NewNFT(defaultSymbol, defaultTokenID2, defaultName, defaultTokenURI, addr1)))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol, addr1, types.NewNFT(defaultSymbol, defaultTokenID3, defaultName, defaultTokenURI, addr1)))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol, addr1, types.NewNFT(defaultSymbol, defaultTokenID4, defaultName, defaultTokenURI, addr1)))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol2, addr1, types.NewNFT(defaultSymbol2, defaultTokenID1, defaultName, defaultTokenURI, addr1)))
	require.NoError(t, keeper.GrantPermission(ctx, addr1, addr2, types.NewMintPermission(defaultSymbol, defaultTokenType)))
	require.NoError(t, keeper.MintNFT(ctx, defaultSymbol, addr2, types.NewNFT(defaultSymbol, defaultTokenID5, defaultName, defaultTokenURI, addr2)))
	require.NoError(t, keeper.IssueFT(ctx, defaultSymbol, addr1, types.NewFT(defaultSymbol, defaultTokenIDFT, defaultName, defaultTokenURI, sdk.NewInt(1), true), sdk.NewInt(defaultAmount)))
}

func verifyTokenFunc(t *testing.T, expected types.Token, actual types.Token) {
	switch e := expected.(type) {
	case types.FT:
		a, ok := actual.(types.FT)
		require.True(t, ok)
		require.Equal(t, e.GetSymbol(), a.GetSymbol())
		require.Equal(t, e.GetName(), a.GetName())
		require.Equal(t, e.GetTokenID(), a.GetTokenID())
		require.Equal(t, e.GetTokenType(), a.GetTokenType())
		require.Equal(t, e.GetTokenIndex(), a.GetTokenIndex())
		require.Equal(t, e.GetTokenURI(), a.GetTokenURI())
		require.Equal(t, e.GetDecimals(), a.GetDecimals())
		require.Equal(t, e.GetMintable(), a.GetMintable())
	case types.NFT:
		a, ok := actual.(types.NFT)
		require.True(t, ok)
		require.Equal(t, e.GetSymbol(), a.GetSymbol())
		require.Equal(t, e.GetName(), a.GetName())
		require.Equal(t, e.GetTokenID(), a.GetTokenID())
		require.Equal(t, e.GetTokenType(), a.GetTokenType())
		require.Equal(t, e.GetTokenIndex(), a.GetTokenIndex())
		require.Equal(t, e.GetTokenURI(), a.GetTokenURI())
		require.Equal(t, e.GetOwner(), a.GetOwner())
	default:
		panic("never happen")
	}
}

func verifyTokenTypeFunc(t *testing.T, expected types.TokenType, actual types.TokenType) {
	require.Equal(t, expected.GetSymbol(), actual.GetSymbol())
	require.Equal(t, expected.GetName(), actual.GetName())
	require.Equal(t, expected.GetTokenType(), actual.GetTokenType())
}
