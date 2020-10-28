package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {

	//read json data for transformation
	var jsonCoinPrices, err = ioutil.ReadFile("result.json")

	if err != nil {
		log.Fatal(err)
	}

	//create structure to unpack json data into
	type Price struct {
		Currency string
		Timestamps []string
		Prices []string
	}

	var prices []Price
	//Extract json into struct
	if err := json.Unmarshal(jsonCoinPrices, &prices); err != nil {
		log.Fatal(err)
	}

	//Iterate over the timestamps to transform into something BI Tools can ingest
	n := 0
	priceArrayLength := len(prices[0].Timestamps)
	result := make([]map[string]string, priceArrayLength)
	for n < priceArrayLength {
		//add row to array
		coinPrice := map[string]string{
			"currency":prices[0].Currency,
			"timestamp":prices[0].Timestamps[n],
			"price":prices[0].Prices[n],
		}

		result[n] = coinPrice
		n += 1
	}

	//encode result map to JSON
	data, _ := json.Marshal(result)

	ioutil.WriteFile("coin-prices.json", data, 0644)
}

