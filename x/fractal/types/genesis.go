package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:       PortID,
		ExchangeList: []Exchange{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in exchange
	exchangeIdMap := make(map[uint64]bool)
	exchangeCount := gs.GetExchangeCount()
	for _, elem := range gs.ExchangeList {
		if _, ok := exchangeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for exchange")
		}
		if elem.Id >= exchangeCount {
			return fmt.Errorf("exchange id should be lower or equal than the last id")
		}
		exchangeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
