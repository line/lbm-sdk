package v3

import sdk "github.com/cosmos/cosmos-sdk/types"

const ClassStoreKey = "class"

var (
	classStorePrefix = []byte{0x50}

	paramsKey        = []byte{0x00}
	nonceKey         = []byte{0x01}
	idKeyPrefix      = []byte{0x02}
	balanceKeyPrefix = []byte{0x20}
	ownerKeyPrefix   = []byte{0x21}
)

// Deprecated
var (
	parentKeyPrefix = []byte{0x23}
	childKeyPrefix  = []byte{0x24}
)

func idKey(id string) []byte {
	key := make([]byte, len(idKeyPrefix)+len(id))
	copy(key, idKeyPrefix)
	copy(key[len(idKeyPrefix):], id)
	return key
}

func splitIDKey(key []byte) (id string) {
	return string(key[len(idKeyPrefix):])
}

func balanceKey(contractID string, address sdk.AccAddress, tokenID string) []byte {
	prefix := balanceKeyPrefixByAddress(contractID, address)
	key := make([]byte, len(prefix)+len(tokenID))

	copy(key, prefix)
	copy(key[len(prefix):], tokenID)

	return key
}

func balanceKeyPrefixByAddress(contractID string, address sdk.AccAddress) []byte {
	prefix := balanceKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+1+len(address))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(address))

	begin++
	copy(key[begin:], address)

	return key
}

func balanceKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(balanceKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, balanceKeyPrefix)

	begin += len(balanceKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

// ----------------------------------------------------------------------------
// parent
func parentKey(contractID, tokenID string) []byte {
	prefix := parentKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+len(tokenID))

	copy(key, prefix)
	copy(key[len(prefix):], tokenID)

	return key
}

func parentKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(parentKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, parentKeyPrefix)

	begin += len(parentKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

func splitParentKey(key []byte) (contractID, tokenID string) {
	begin := len(parentKeyPrefix) + 1
	end := begin + int(key[begin-1])
	contractID = string(key[begin:end])

	begin = end
	tokenID = string(key[begin:])

	return
}

// ----------------------------------------------------------------------------
// child
func childKey(contractID, tokenID, childID string) []byte {
	prefix := childKeyPrefixByTokenID(contractID, tokenID)
	key := make([]byte, len(prefix)+len(childID))

	copy(key, prefix)
	copy(key[len(prefix):], childID)

	return key
}

func childKeyPrefixByTokenID(contractID, tokenID string) []byte {
	prefix := childKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+1+len(tokenID))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(tokenID))

	begin++
	copy(key[begin:], tokenID)

	return key
}

func childKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(childKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, childKeyPrefix)

	begin += len(childKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

func ownerKey(contractID, tokenID string) []byte {
	prefix := ownerKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+len(tokenID))

	copy(key, prefix)
	copy(key[len(prefix):], tokenID)

	return key
}

func ownerKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(ownerKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, ownerKeyPrefix)

	begin += len(ownerKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}
