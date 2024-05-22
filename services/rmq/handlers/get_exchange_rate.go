package handlers

import (
	"context"
	gitProto "github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"quote-manager/config"
	"quote-manager/dto/proto"
	"quote-manager/services/redis"
	"quote-manager/services/rmq/senders"
	"quote-manager/util"
)

type GetExchangeRateHandler struct {
	sender senders.Sender
	rCl    *redis.Client
}

func NewGetExchangeRateHandler(sender senders.Sender, client *redis.Client) GetExchangeRateHandler {
	return GetExchangeRateHandler{
		sender: sender,
		rCl:    client,
	}
}

func (h GetExchangeRateHandler) HandleMessage(body []byte) {
	req := &proto.GetExchangeRateRequest{}
	err := gitProto.Unmarshal(body, req)
	util.IsError(err, "Failed to unmarshal message")

	log.Info("Received GetExchangeRateRequest: %s\n", req)

	allRates := h.rCl.GetRatesFromRedis(context.TODO())

	resp := &proto.GetExchangeRateResponse{
		Id:    req.Id,
		Rates: allRates,
		Error: nil,
	}

	h.sender.SendMessage(
		config.RabbitExchangeRateExchange,
		config.GetExchangeRateResponseRoutingKey,
		"topic",
		resp,
	)
}
