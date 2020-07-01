package main

import (
	"context"
	"fmt"

	"github.com/luno/luno-go"
)

func main() {
	ctx := context.Background()
	//START OMIT
	lunoClient := luno.NewClient()
	lunoClient.SetAuth("api key", "api secret") // <-- Paste your Luno key and secret here // HL

	res, err := lunoClient.GetBalances(ctx, &luno.GetBalancesRequest{}) // HL
	if err != nil {
		panic(err)
	}
	for _, b := range res.Balance {
		fmt.Println(b.Asset, b.Balance) // HL
	}
	//END OMIT
}
