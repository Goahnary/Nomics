#!/bin/bash

echo "curl -XGET https://api.nomics.com/v1/currencies/sparkline"
curl -s -o result.json -XGET "https://api.nomics.com/v1/currencies/sparkline?key=${NOMIC_API_KEY}&ids=BTC&start=2020-01-01T00%3A00%3A00Z&end=2020-10-26T00%3A00%3A00Z"

#timestamps=(`cat result.json | jq '.[].timestamps[]'`)
#prices=(`cat result.json | jq '.[].prices[]'`)

#echo ${!timestamps[*]}
#: > export.json
#echo "[" >> export.json

#switched=false
#for i in ${!timestamps[*]}; do
#	if [ !$switched ]; then
#		switched=true
#		echo "switching!"
#	else
#		echo "," >> export.json
#	fi
#	echo "{\"timestamp\": ${timestamps[$i]},\"price\": ${prices[$i]}}" >> export.json
#done

#echo "]" >> export.json
