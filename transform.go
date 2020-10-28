package main

import (
	"encoding/json"
	"encoding/csv"
	"io/ioutil"
	"fmt"
	"log"
	"os"
)

func main() {

	//read json data for transformation
	var jsonCoinPrices, err = ioutil.ReadFile("price-data.json")

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

	allPrices := make([]map[string]string, 0)

	i := 0
	for i < len(prices) {
		//Iterate over the timestamps to transform into something BI Tools can ingest
		n := 0
		priceArrayLength := len(prices[i].Timestamps)
		result := make([]map[string]string, priceArrayLength)
		for n < priceArrayLength {
			//add row to array
			coinPrice := map[string]string{
				"currency":prices[i].Currency,
				"timestamp":prices[i].Timestamps[n],
				"price":prices[i].Prices[n],
			}

			result[n] = coinPrice
			n += 1
		}

		//collect for mass export of data
		allPrices = append(allPrices, result...)

		//encode result map to JSON
		data, _ := json.Marshal(result)

		fname := prices[i].Currency + "-prices.json"
		ioutil.WriteFile(fname, data, 0644)

		//Also write to csv
		csvName := prices[i].Currency + "-prices.csv"
		f, err := os.Create(csvName)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		w := csv.NewWriter(f)

		//write header
		header := []string{"currency", "timestamp", "price"}
		w.Write(header)
		for _, obj := range result {
		//	fmt.Println(obj)
			var record []string
			record = append(record, obj["currency"], obj["timestamp"], obj["price"])
			w.Write(record)
			record = nil
		}
		w.Flush()

		i += 1 //Always increment last
	}
	allData, _ := json.Marshal(allPrices)
	ioutil.WriteFile("coin-prices.json", allData, 0644)
}

