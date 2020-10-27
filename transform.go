package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	var jsonCoinPrices, err = ioutil.ReadFile("result.json")

	if err != nil {
		fmt.Println("error:", err)
	}

	type Price struct {
		Currency string
		timestamps  []string
		prices []string
	}

	var prices []Price

	erro := json.Unmarshal(jsonCoinPrices, &prices)

	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("%+v", prices)
}

