package main

import (
	"context"
	"fmt"

	"github.com/luno/luno-go"
)

func main() {
	ctx := context.Background()
	// START OMIT
	lunoClient := luno.NewClient()
	lunoClient.SetAuth("api key", "api secret") // <-- Paste your Luno key and secret here

	req := luno.StopOrderRequest{OrderId: "BX97B352A6D789C"}
	res, err := lunoClient.StopOrder(ctx, &req)
	if err != nil {
		panic(err)
	}

	if res.Success {
		fmt.Println("Stopped successfully")
	} else {
		fmt.Println("Failed to stop")
	}
	// END OMIT
}
