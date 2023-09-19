#!/bin/bash
L2BINARYNAME=rollupd
ROLLUP_NAME=test-rollup
L1_CHAIN_ID=sim
L2_CHAIN_ID=sim2
L1_KEYRING_DIR=~/.simapp
L2_KEYRING_DIR=~/.l2simapp
RPC_URI=http://localhost:26659
TEST_SEQUENCER_ADDRESS=link1twsfmuj28ndph54k4nw8crwu8h9c8mh3rtx705
SEQUENCER_DIR=simapp0
TEST_MNEMONIC="mind flame tobacco sense move hammer drift crime ring globe art gaze cinnamon helmet cruise special produce notable negative wait path scrap recall have"

cd ..

# Build & rename
make build && cp -r build/simd $GOPATH/bin/$L2BINARYNAME
${L2BINARYNAME} version

# Init sequencer

${L2BINARYNAME} init rollupdemo --home $L2_KEYRING_DIR/$SEQUENCER_DIR --chain-id $L2_CHAIN_ID > /dev/null 2>&1
${L2BINARYNAME} keys add validator --keyring-backend=test --home $L2_KEYRING_DIR/$SEQUENCER_DIR --recover --account=0 <<< ${TEST_MNEMONIC} > /dev/null 2>&1
${L2BINARYNAME} add-genesis-account $(${L2BINARYNAME} --home $L2_KEYRING_DIR/$SEQUENCER_DIR keys show validator -a --keyring-backend=test) 100000000000stake,100000000000tcony --home $L2_KEYRING_DIR/$SEQUENCER_DIR > /dev/null 2>&1
${L2BINARYNAME} gentx validator 10000000000stake --keyring-backend=test --home $L2_KEYRING_DIR/$SEQUENCER_DIR --chain-id=$L2_CHAIN_ID > /dev/null 2>&1
${L2BINARYNAME} collect-gentxs --home $L2_KEYRING_DIR/$SEQUENCER_DIR > /dev/null 2>&1

# Run L2 sequencer
${L2BINARYNAME} start --home $L2_KEYRING_DIR/$SEQUENCER_DIR --p2p.laddr "tcp://0.0.0.0:26556" --grpc.address "0.0.0.0:9190" --grpc-web.address "0.0.0.0:9191" --rollkit.sequencer "true" --rollkit.da_layer finschia --rollkit.da_config='{"rpc_uri":"'$RPC_URI'","chain_id":"'$L1_CHAIN_ID'","keyring_dir":"'$L1_KEYRING_DIR'","from":"'$TEST_SEQUENCER_ADDRESS'", "rollup_name":"'$ROLLUP_NAME'"}' > $L2_KEYRING_DIR/$L2_CHAIN_ID.log 2>&1 &
