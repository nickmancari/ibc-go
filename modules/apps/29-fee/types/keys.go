package types

import fmt "fmt"

const (
	// ModuleName defines the 29-fee name
	ModuleName = "ibcfee"

	// StoreKey is the store key string for IBC transfer
	StoreKey = ModuleName

	// PortKey is the port id that is wrapped by fee middleware
	PortKey = "feetransfer"

	// RouterKey is the message route for IBC transfer
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = ModuleName
)

// Key for relayer source address -> counteryparty address mapping
func KeySourceAddress(sourceAddress string) []byte {
	return []byte(fmt.Sprintf("relayerSourceAddress/%s", sourceAddress))
}
