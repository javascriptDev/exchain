package keeper

import (
	"github.com/okex/exchain/x/evm/types"
	"math/big"
	"sync"
)

type LogsManager struct {
	mu      sync.RWMutex
	Results map[string]TxResult
}

func NewLogManager() *LogsManager {
	return &LogsManager{
		mu:      sync.RWMutex{},
		Results: make(map[string]TxResult),
	}
}

func (l *LogsManager) Set(txBytes string, value TxResult) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Results[txBytes] = value
}

func (l *LogsManager) Get(txBytes string) (TxResult, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	data, ok := l.Results[txBytes]
	return data, ok
}

func (l *LogsManager) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.Results)
}

type TxResult struct {
	ResultData *types.ResultData
	Err        error
}

func (k *Keeper) FixLog(isAnteFailed [][]string) map[int][]byte {
	res := make(map[int][]byte, 0)
	logSize := uint(0)
	txInBlock := int(-1)
	k.Bloom = new(big.Int)

	for index := 0; index < len(isAnteFailed); index++ {
		rs, ok := k.LogsManages.Get(isAnteFailed[index][0])
		if !ok || isAnteFailed[index][1] != "" {
			continue
		}
		txInBlock++

		if rs.ResultData == nil {
			continue
		}

		for _, v := range rs.ResultData.Logs {
			v.Index = logSize
			v.TxIndex = uint(txInBlock)
			logSize++
		}
		k.Bloom = k.Bloom.Or(k.Bloom, rs.ResultData.Bloom.Big())
		data, err := types.EncodeResultData(*rs.ResultData)
		if err != nil {
			panic(err)
		}
		res[index] = data
	}
	return res
}
