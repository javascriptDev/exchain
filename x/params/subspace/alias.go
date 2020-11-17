package subspace

import (
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

type (
	ParamSetPairs = subspace.ParamSetPairs
	KeyTable      = subspace.KeyTable
)

var (
	NewKeyTable     = subspace.NewKeyTable
	NewParamSetPair = subspace.NewParamSetPair

	StoreKey  = subspace.StoreKey
	TStoreKey = subspace.TStoreKey
)
