package handlers

import (
	"context"
	gitProto "github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"quote-manager/dto/proto"
	"quote-manager/services/redis"
	"quote-manager/services/rmq/senders"
	"quote-manager/util"
	"strconv"
)

type UpdateOrderEventHandler struct {
	sender   senders.Sender
	rCl      *redis.Client
	respChan chan []byte
}

func NewUpdateOrderEventHandler(sender senders.Sender, client *redis.Client, respChan chan []byte) UpdateOrderEventHandler {
	return UpdateOrderEventHandler{
		sender:   sender,
		rCl:      client,
		respChan: respChan,
	}
}

func (h UpdateOrderEventHandler) HandleMessage(body []byte) {
	event := &proto.OrderUpdateEvent{}
	err := gitProto.Unmarshal(body, event)
	util.IsError(err, "Failed to unmarshal message")

	if event.Error != nil || event.Order.FillVolume != 0 {
		return
	}
	h.rCl.IncrOrderCount(context.Background(), event.Order.Pair)
	log.Info("tregregregregregregegregre %s\n", h.rCl.GetRatesFromRedis(context.Background()))
	log.Info("Received OrderUpdateEvent: %s\n", event)

	rate := &proto.ExchangeRate{
		Pair: event.Order.Pair,
	}
	var price float64
	rates := h.rCl.GetRatesFromRedis(context.Background())

	for _, exchangeRate := range rates {
		if exchangeRate.Pair == event.Order.Pair {
			rate = exchangeRate
			break
		}
	}

	if event.Order.Direction == proto.Direction_ORDER_DIRECTION_SELL {
		price = 1 / event.Order.InitPrice
	} else {
		price = event.Order.InitPrice
	}

	if rate.Min > price || rate.Min == 0 {
		rate.Min = price
	}
	if rate.Max < price || rate.Max == 0 {
		rate.Max = price
	}

	rate.Average = h.movingAverage(price, rate.Average, event.Order.Pair)

	h.rCl.SetExchangeRateData(context.Background(), rate)
	f := h.rCl.GetRatesFromRedis(context.TODO())

	log.Info(f)
}

func (h UpdateOrderEventHandler) movingAverage(initPrice, previousAvg float64, pair string) float64 {
	orderCount := h.rCl.GetOrderCount(context.Background(), pair)
	count, err := strconv.ParseFloat(strconv.Itoa(orderCount), 64)
	util.IsError(err, "err parse to float(movingAverage)")

	return ((count-1)*previousAvg + initPrice) / count
}
