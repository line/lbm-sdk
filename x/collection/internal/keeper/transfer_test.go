package keeper

import (
	"testing"

	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestTransferFTScenario(t *testing.T) {
	ctx := cacheKeeper()

	// issue idf token
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(defaultContractID, defaultName, defaultImgURI), addr1))
	require.NoError(t, keeper.IssueFT(ctx, defaultContractID, addr1, types.NewFT(defaultContractID, defaultTokenIDFT, defaultName, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount)))

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferFT(ctx, defaultContractID, addr1, addr2, types.NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount/2))))

	//
	// transfer failure cases
	//
	// Insufficient coins
	require.EqualError(t, keeper.TransferFT(ctx, defaultContractID, addr1, addr2, types.NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount))), types.ErrInsufficientToken(types.DefaultCodespace, "insufficient account funds[abcdef01]; 500:0000000100000000 < 1000:0000000100000000").Error())
}

func TestTransferNFTScenario(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID3))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID4))

	//
	// transfer failure cases
	//

	// transfer non-exist token : failure
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID8), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID8).Error())

	// transfer a child : failure
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID2), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultTokenID2).Error())
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID3), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultTokenID3).Error())
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID4), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultTokenID4).Error())

	// transfer non-mine : failure
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID5), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID5, addr1).Error())

	// transfer-cnft cft : failure
	require.EqualError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenIDFT), types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID1))
	require.NoError(t, keeper.TransferNFT(ctx, defaultContractID, addr2, addr1, defaultTokenID1))
	require.NoError(t, keeper.TransferNFT(ctx, defaultContractID, addr1, addr2, defaultTokenID1))

	// verify the owner of transferred tokens
	// owner of token1 is addr2
	token1, err1 := keeper.GetToken(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err1)
	require.Equal(t, token1.(types.NFT).GetOwner(), addr2)

	// owner of token2 is addr2
	token2, err2 := keeper.GetToken(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err2)
	require.Equal(t, token2.(types.NFT).GetOwner(), addr2)

	// owner of token3 is addr2
	token3, err3 := keeper.GetToken(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err3)
	require.Equal(t, token3.(types.NFT).GetOwner(), addr2)

	// owner of token4 is addr2
	token4, err4 := keeper.GetToken(ctx, defaultContractID, defaultTokenID4)
	require.NoError(t, err4)
	require.Equal(t, token4.(types.NFT).GetOwner(), addr2)
}
