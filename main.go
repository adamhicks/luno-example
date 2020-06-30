package main

import (
	"context"
	"fmt"

	"github.com/luno/luno-go"
	"github.com/luno/luno-go/decimal"
)

// printMyBalances will print out the balance of each of your wallets
func printMyBalances(lunoClient *luno.Client) error {
	var req luno.GetBalancesRequest
	ctx := context.Background()

	res, err := lunoClient.GetBalances(ctx, &req)
	if err != nil {
		return err
	}

	for _, b := range res.Balance {
		fmt.Println(b.Asset, b.Balance)
	}
	return nil
}

// placeAnOrder will place a Ask for 0.0005 BTC at 20,000 EUR
func placeAnOrder(lunoClient *luno.Client) error {
	req := luno.PostLimitOrderRequest{
		Pair:     "XBTEUR",
		Type:     luno.OrderTypeAsk,
		Price:    decimal.NewFromInt64(20000),
		Volume:   decimal.NewFromFloat64(0.0005, 8),
		PostOnly: true,
	}
	ctx := context.Background()

	res, err := lunoClient.PostLimitOrder(ctx, &req)
	if err != nil {
		return err
	}

	fmt.Println("Placed", res.OrderId)
	return nil
}

func main() {
	lunoClient := luno.NewClient()
	lunoClient.SetAuth("api key", "api secret")

	err := printMyBalances(lunoClient)
	if err != nil {
		panic(err)
	}

	// Uncomment the next part if you want to place an order
	//
	// err = placeAnOrder(lunoClient)
	// if err != nil {
	// 	panic(err)
	// }
}
