package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/okex/exchain/x/evm"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

// feeCollectorHandler set or get the value of feeCollectorAcc
func feeCollectorHandler(ak auth.AccountKeeper, sk supply.Keeper) sdk.FeeCollectorAccHandler {
	return func(ctx sdk.Context, updateValue bool, balance sdk.Coins) sdk.Coins {
		acc := ak.GetAccount(ctx, sk.GetModuleAddress(auth.FeeCollectorName))
		if updateValue {
			acc.SetCoins(balance)
			ak.SetAccount(ctx, acc)
		}
		return acc.GetCoins()
	}
}

// evmTxFeeHandler get tx fee for evm tx
func evmTxFeeHandler() sdk.GetTxFeeHandler {
	return func(tx sdk.Tx) (fee sdk.Coins, isEvm bool) {
		if _, ok := tx.(evmtypes.MsgEthereumTx); ok {
			isEvm = true
		}
		if feeTx, ok := tx.(authante.FeeTx); ok {
			fee = feeTx.GetFee()
		}
		return
	}
}

// fixLogForParallelTxHandler fix log for parallel tx
func fixLogForParallelTxHandler(ek *evm.Keeper) sdk.LogFix {
	return func(execResults [][]string) (logs map[int][]byte) {
		return ek.FixLog(execResults)
	}
}
