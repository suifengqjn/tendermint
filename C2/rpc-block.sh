#!/bin/bash
decoded=`curl http://localhost:26657/block?height=2 -s | jq '.result.block.data.txs[0]' | sed 's/\"//g' | base64 -d`
echo "tx => $decoded"