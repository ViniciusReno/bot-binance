package stream

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
)

type TradeEvent struct {
	Event                 string
	Time                  int64
	Symbol                string
	AggTradeID            int64
	Price                 float64
	Quantity              int64
	FirstBreakdownTradeID int64
	LastBreakdownTradeID  int64
	TradeTime             int64
	IsBuyerMaker          bool
	Placeholder           bool
}

func StartService(ctx context.Context, client *binance.Client) {
	res, err := client.NewStartUserStreamService().Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Deph(crypto string) {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsDepthServe(crypto, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}

// Kline bring the candle graphic
func Kline(crypto string) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println("event", event.Kline)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe(crypto, "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Aggregate(crypto string, priceChan chan TradeEvent) {
	go func() {
		wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
			priceChan <- parseaggregateEvent(event)
		}
		errHandler := func(err error) {
			fmt.Println(err)
		}
		doneC, _, err := binance.WsAggTradeServe(crypto, wsAggTradeHandler, errHandler)
		if err != nil {
			fmt.Println(err)
			return
		}
		<-doneC
	}()
}

func parseaggregateEvent(event *binance.WsAggTradeEvent) TradeEvent {
	price, _ := strconv.ParseFloat(event.Price, 64)
	qtd, _ := strconv.ParseInt(event.Quantity, 6, 12)
	return TradeEvent{
		Event:                 event.Event,
		Time:                  event.Time,
		Symbol:                event.Symbol,
		AggTradeID:            event.AggTradeID,
		Price:                 price,
		Quantity:              qtd,
		FirstBreakdownTradeID: event.FirstBreakdownTradeID,
		LastBreakdownTradeID:  event.LastBreakdownTradeID,
		TradeTime:             event.TradeTime,
		IsBuyerMaker:          event.IsBuyerMaker,
		Placeholder:           event.Placeholder,
	}
}
