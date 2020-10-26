#!/bin/bash

echo "curl -XGET https://api.nomics.com/v1/currencies/ticker"
curl -s -o result.json -XGET "https://api.nomics.com/v1/currencies/ticker?key=${NOMIC_API_KEY}&ids=BTC,ETH,XRP&interval=1d,30d&per-page=100&page=1"

