package functions

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func CreateOrders(ctx context.Context, client *binance.Client, crypto string, qtd string, price string) {
	order, err := client.NewCreateOrderService().Symbol(crypto).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(qtd).
		Price(price).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

func CreateOcoOrders(ctx context.Context, client *binance.Client, crypto string, qtd string, price string) {
}

func OpenOrders(ctx context.Context, client *binance.Client, crypto string) ([]*binance.Order, error) {
	openOrders, err := client.NewListOpenOrdersService().Symbol(crypto).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return openOrders, nil
}

func ListKlines(ctx context.Context, client *binance.Client, crypto string) ([]*binance.Kline, error) {
	klines, err := client.NewKlinesService().Symbol(crypto).
		Interval("15m").Do(ctx)
	if err != nil {
		return nil, err
	}
	return klines, nil
}

func ShowDeep(ctx context.Context, client *binance.Client, crypto string) (*binance.DepthResponse, error) {
	res, err := client.NewDepthService().Symbol(crypto).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
