package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/okex/exchain/x/evm"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

func NewIsEvmTxHandler(tx sdk.Tx) bool {
	if tx != nil {
		switch tx.(type) {
		case evmtypes.MsgEthereumTx:
			return true
		}
	}
	return false
}

func NewFeeCollectorAccHandler(ak auth.AccountKeeper, sk supply.Keeper) sdk.FeeCollectorAccHandler {
	return func(ctx sdk.Context, updateValue bool, balance sdk.Coins) sdk.Coins {
		acc := ak.GetAccount(ctx, sk.GetModuleAddress(auth.FeeCollectorName))
		if updateValue {
			acc.SetCoins(balance)
			ak.SetAccount(ctx, acc)
		}
		return acc.GetCoins()
	}
}

func NewGetTxFeeHandler() sdk.GetTxFeeHandler {
	return func(tx sdk.Tx) sdk.Coins {
		feeTx, ok := tx.(authante.FeeTx)
		if ok {
			return feeTx.GetFee()
		}
		return sdk.Coins{}
	}
}

func NewFixLog(ek *evm.Keeper) sdk.LogFix {
	return func(isAnteFailed [][]string) (logs map[int][]byte) {
		return ek.FixLog(isAnteFailed)
	}
}
