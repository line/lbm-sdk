package keeper

import (
	sdk "github.com/line/lbm-sdk/types"

	"github.com/line/lbm-sdk/x/collection"
)

var (
	balanceKeyPrefix     = []byte{0x00}
	contractKeyPrefix    = []byte{0x01}
	classKeyPrefix       = []byte{0x02}
	nextClassIDKeyPrefix = []byte{0x03}

	authorizationKeyPrefix = []byte{0x04}
	grantKeyPrefix         = []byte{0x05}
)

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

// func splitBalanceKey(key []byte) (contractID string, address sdk.AccAddress, tokenID string) {
// 	begin := len(balanceKeyPrefix) + 1
// 	end := begin + int(key[begin-1])
// 	contractID = string(key[begin:end])

// 	begin = end + 1
// 	end = begin + int(key[begin-1])
// 	address = sdk.AccAddress(key[begin:end])

// 	begin = end
// 	tokenID = string(key[begin:])

// 	return
// }

func classKey(contractID string, classID string) []byte {
	prefix := classKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+len(classID))

	copy(key, prefix)
	copy(key[len(prefix):], classID)

	return key
}

func classKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(classKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, classKeyPrefix)

	begin += len(classKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

func contractKey(contractID string) []byte {
	key := make([]byte, len(contractKeyPrefix)+len(contractID))

	copy(key, contractKeyPrefix)
	copy(key[len(contractKeyPrefix):], contractID)

	return key
}

func nextClassIDKey(contractID string) []byte {
	key := make([]byte, len(nextClassIDKeyPrefix)+len(contractID))

	copy(key, nextClassIDKeyPrefix)
	copy(key[len(nextClassIDKeyPrefix):], contractID)

	return key
}

func authorizationKey(contractID string, operator, holder sdk.AccAddress) []byte {
	prefix := authorizationKeyPrefixByOperator(contractID, operator)
	key := make([]byte, len(prefix)+len(holder))

	copy(key, prefix)
	copy(key[len(prefix):], holder)

	return key
}

func authorizationKeyPrefixByOperator(contractID string, operator sdk.AccAddress) []byte {
	prefix := authorizationKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+1+len(operator))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(operator))

	begin++
	copy(key[begin:], operator)

	return key
}

func authorizationKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(authorizationKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, authorizationKeyPrefix)

	begin += len(authorizationKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

// func splitAuthorizationKey(key []byte) (contractID string, operator, holder sdk.AccAddress) {
// 	begin := len(authorizationKeyPrefix) + 1
// 	end := begin + int(key[begin-1])
// 	contractID = string(key[begin:end])

// 	begin = end + 1
// 	end = begin + int(key[begin-1])
// 	operator = sdk.AccAddress(key[begin:end])

// 	begin = end
// 	holder = sdk.AccAddress(key[begin:])

// 	return
// }

func grantKey(contractID string, grantee sdk.AccAddress, permission collection.Permission) []byte {
	prefix := grantKeyPrefixByGrantee(contractID, grantee)
	key := make([]byte, len(prefix)+1)

	copy(key, prefix)
	key[len(prefix)] = byte(permission)

	return key
}

func grantKeyPrefixByGrantee(contractID string, grantee sdk.AccAddress) []byte {
	prefix := grantKeyPrefixByContractID(contractID)
	key := make([]byte, len(prefix)+1+len(grantee))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(grantee))

	begin++
	copy(key[begin:], grantee)

	return key
}

func grantKeyPrefixByContractID(contractID string) []byte {
	key := make([]byte, len(grantKeyPrefix)+1+len(contractID))

	begin := 0
	copy(key, grantKeyPrefix)

	begin += len(grantKeyPrefix)
	key[begin] = byte(len(contractID))

	begin++
	copy(key[begin:], contractID)

	return key
}

// func splitGrantKey(key []byte) (contractID string, grantee sdk.AccAddress, permission collection.Permission) {
// 	begin := len(grantKeyPrefix) + 1
// 	end := begin + int(key[begin-1])
// 	contractID = string(key[begin:end])

// 	begin = end + 1
// 	end = begin + int(key[begin-1])
// 	grantee = sdk.AccAddress(key[begin:end])

// 	begin = end
// 	permission = collection.Permission(key[begin])

// 	return
// }
