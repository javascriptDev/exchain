killall exchaind_biao
rm -rf multi_run.log_biao
make mainnet
cp /Users/oker/scf/bin/exchaind /Users/oker/scf/bin/exchaind_biao
export EXCHAIND_PATH=~/.exchaind_biao
#rm -rf ${EXCHAIND_PATH}

#exchaind_biao init multi_run --chain-id exchain-66 --home ${EXCHAIND_PATH}
#cp /Users/oker/scf/genesis.json ${EXCHAIND_PATH}/config/genesis.json

#cp -rf  ~/.exchaind_back  ~/.exchaind

export EXCHAIN_SEEDS="e926c8154a2af4390de02303f0977802f15eafe2@3.16.103.80:26656,7fa5b1d1f1e48659fa750b6aec702418a0e75f13@175.41.191.69:26656,c8f32b793871b56a11d94336d9ce6472f893524b@35.74.8.189:26656"
#exchaind replay -d /Users/oker/scf/src/data  --home ${EXCHAIND_PATH} S > multi_run.log 2>&1 &
exchaind_biao start --chain-id exchain-66 --mempool.sort_tx_by_gp --home ${EXCHAIND_PATH} --p2p.seeds $EXCHAIN_SEEDS > multi_run.log_biao 2>&1 &