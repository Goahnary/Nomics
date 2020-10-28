package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	//read json data for transformation
	var jsonCoinPrices, err = ioutil.ReadFile("result.json")

	if err != nil {
		fmt.Println("error:", err)
	}

	//create structure to unpack json data into
	type Price struct {
		Currency string
		Timestamps []string
		Prices []string
	}

	var prices []Price
	//Extract json into struct
	erro := json.Unmarshal(jsonCoinPrices, &prices)

	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Println("%v", prices)
	fmt.Println()

	//Iterate over the timestamps to transform into something BI Tools can ingest
	n := 0
	result := make(map[int][]map[string]string)
	for n < len(prices[0].Timestamps) {
		//add row to array
		coinPrice := map[string]string{
			"currency":prices[0].Currency,
			"timestamp":prices[0].Timestamps[n],
			"price":prices[0].Prices[n],
		}

		result[n] = append(result[n], coinPrice)
		n += 1
	}

	fmt.Println()

	//encode result map to JSON
	data, _ := json.Marshal(result)

	fmt.Println(string(data))
}

