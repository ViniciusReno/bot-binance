package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ViniciusReno/bot-binance/config"
	"github.com/ViniciusReno/bot-binance/stream"
	"github.com/adshao/go-binance/v2"
)

func main() {
	config.Start()
	event := make(chan stream.TradeEvent)
	ctx := context.Background()
	newClientStream(ctx)

	stream.Aggregate(config.Coin, event)

moneyLoop:
	for {
		select {
		case <-ctx.Done():
			break moneyLoop
		case <-event:
			e := <-event
			fmt.Printf("\n%s: %f - %s", config.Coin, e.Price, time.Now().Format("15:04:05"))
		}
	}
}

func newClientStream(ctx context.Context) {
	client := binance.NewClient(config.ApiKey, config.SecretKey)
	client.NewSetServerTimeService().Do(ctx)
	stream.StartService(ctx, client)
}
