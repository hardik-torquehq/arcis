#!/bin/bash

KEY="mykey"
CHAINID="arcis_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t arcis-datadir.XXXXX)

echo "create and add new keys"
./arcisd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Arcis with moniker=$MONIKER and chain-id=$CHAINID"
./arcisd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./arcisd add-genesis-account \
"$(./arcisd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aarcis,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./arcisd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./arcisd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./arcisd validate-genesis --home $DATA_DIR

echo "starting arcis node $i in background ..."
./arcisd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started arcis node"
tail -f /dev/null