package authz

import (
	"github.com/Finschia/finschia-rdk/codec"
	"github.com/Finschia/finschia-rdk/codec/legacy"
	"github.com/Finschia/finschia-rdk/codec/types"
	sdk "github.com/Finschia/finschia-rdk/types"
	"github.com/Finschia/finschia-rdk/types/msgservice"
	authzcodec "github.com/Finschia/finschia-rdk/x/authz/codec"
	fdncodec "github.com/Finschia/finschia-rdk/x/foundation/codec"
	govcodec "github.com/Finschia/finschia-rdk/x/gov/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/authz interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgGrant{}, "cosmos-sdk/MsgGrant")
	legacy.RegisterAminoMsg(cdc, &MsgRevoke{}, "cosmos-sdk/MsgRevoke")
	legacy.RegisterAminoMsg(cdc, &MsgExec{}, "cosmos-sdk/MsgExec")

	cdc.RegisterInterface((*Authorization)(nil), nil)
	cdc.RegisterConcrete(&GenericAuthorization{}, "cosmos-sdk/GenericAuthorization", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrant{},
		&MsgRevoke{},
		&MsgExec{},
	)

	registry.RegisterInterface(
		"cosmos.v1beta1.Authorization",
		(*Authorization)(nil),
		&GenericAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, MsgServiceDesc())
}
func init() {
	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(fdncodec.Amino)
}
