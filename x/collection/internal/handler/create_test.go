package handler

import (
	"testing"

	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	contractID = "9be17165"
)

func TestHandleMsgCreateCollection(t *testing.T) {
	ctx, h := cacheKeeper()
	{
		msg := types.NewMsgCreateCollection(addr1, defaultName, defaultMeta, defaultImgURI)
		res, err := h(ctx, msg)
		require.NoError(t, err)

		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("create_collection", sdk.NewAttribute("contract_id", contractID)),
			sdk.NewEvent("create_collection", sdk.NewAttribute("name", defaultName)),
			sdk.NewEvent("create_collection", sdk.NewAttribute("owner", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", contractID)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "issue")),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "mint")),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "burn")),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "modify")),
		}
		verifyEventFunc(t, e, res.Events)
	}
}
