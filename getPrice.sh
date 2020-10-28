#!/bin/bash

startDate="$(date +'%Y')-01-01T00%3A00%3A00Z"
stopDate="$(date +'%Y-%m-%d')T00%3A00%3A00Z"
echo "curl -XGET https://api.nomics.com/v1/currencies/sparkline"
curl -s -o result.json -XGET "https://api.nomics.com/v1/currencies/sparkline?key=${NOMIC_API_KEY}&ids=BTC&start=${startDate}&end=${stopDate}"

