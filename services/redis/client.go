package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"quote-manager/dto/proto"
	"quote-manager/util"
	"strconv"
)

type Client struct {
	client *redis.Client
}

func NewRedisClient(addr, password string, db int) *Client {
	return &Client{
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:7000",
			DB:   0,
		}),
	}
}

func (rc *Client) IncrOrderCount(ctx context.Context, pair string) {
	err := rc.client.Incr(ctx, fmt.Sprintf("order-count-%s", pair)).Err()
	util.IsError(err, "error incr")
}

func (rc *Client) GetOrderCount(ctx context.Context, pair string) int {
	result, err := rc.client.Get(ctx, fmt.Sprintf("order-count-%s", pair)).Result()
	util.IsError(err, "error incr")

	count, err := strconv.Atoi(result)
	util.IsError(err, "error atoi count (str to int)")

	return count
}

func (rc *Client) SetExchangeRateData(ctx context.Context, rate *proto.ExchangeRate) {

	err := rc.client.SAdd(context.TODO(), "rates", rate.Pair).Err()
	util.IsError(err, "error adding to set")

	err = rc.client.HSet(context.TODO(), fmt.Sprintf("rate:%s", rate.Pair), map[string]interface{}{
		"pair":    rate.Pair,
		"min":     rate.Min,
		"average": rate.Average,
		"max":     rate.Max,
	}).Err()

	util.IsError(err, "err set data to redis")
}

func (rc *Client) GetRatesFromRedis(ctx context.Context) []*proto.ExchangeRate {
	ratePairs, err := rc.client.SMembers(ctx, "rates").Result()
	util.IsError(err, "failed get all rates")

	log.Println(ratePairs)
	var rates []*proto.ExchangeRate
	for _, ratePair := range ratePairs {
		rateData, err := rc.client.HGetAll(ctx, fmt.Sprintf("rate:%s", ratePair)).Result()
		util.IsError(err, fmt.Sprintf("failed get rate [id: %s]", ratePair))

		rate := &proto.ExchangeRate{
			Pair:    rateData["pair"],
			Min:     convertToFloat(rateData["min"]),
			Average: convertToFloat(rateData["average"]),
			Max:     convertToFloat(rateData["max"]),
		}
		rates = append(rates, rate)
	}

	return rates
}
func convertToFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	util.IsError(err, "failed pars to float")
	return result
}
