package main

import (
	"context"
	"fmt"

	"github.com/luno/luno-go"
	"github.com/luno/luno-go/decimal"
)

func main() {
	ctx := context.Background()
	// START OMIT
	lunoClient := luno.NewClient()
	lunoClient.SetAuth("api key", "api secret") // <-- Paste your Luno key and secret here

	// will place an Ask for 0.0005 BTC at 20,000 EUR
	req := luno.PostLimitOrderRequest{
		Pair:     "XBTEUR",                          // Which market (XBT = BTC) // HL
		Type:     luno.OrderTypeAsk,                 // Buying or selling? // HL
		Price:    decimal.NewFromInt64(20000),       // Price, 20,000 EUR // HL
		Volume:   decimal.NewFromFloat64(0.0005, 8), // How much BTC? 0.0005 // HL
		PostOnly: true,                              // Don't trade immediately // HL
	}
	res, err := lunoClient.PostLimitOrder(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Placed", res.OrderId)
	// END OMIT
}
