package services

import (
	"fmt"
	"github.com/streadway/amqp"
	"quote-manager/config"
	"quote-manager/services/redis"
	"quote-manager/services/rmq/consumers"
	"quote-manager/services/rmq/handlers"
	"quote-manager/services/rmq/senders"
	"quote-manager/util"
)

type Server struct {
	clRedis   *redis.Client
	rmqConn   *amqp.Connection
	senders   senders.Sender
	handlers  map[string]handlers.MessageHandler
	consumers map[string]*consumers.Consumer
	ch        chan []byte
}

func NewServer(cfg *config.Config) *Server {
	conn, err := amqp.Dial(fmt.Sprintf(config.RmqUrlConnectionPattern,
		cfg.RabbitConfig.Username,
		cfg.RabbitConfig.Password,
		cfg.RabbitConfig.Host,
		cfg.RabbitConfig.Port,
	))
	util.IsError(err, "err")

	rCl := redis.NewRedisClient(
		fmt.Sprintf(config.RedisConnectionPattern, cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		cfg.RedisConfig.Password,
		8,
	)
	/*f := rCl.GetRatesFromRedis(context.TODO())
	log.Info(f)*/
	return &Server{
		clRedis:   rCl,
		rmqConn:   conn,
		senders:   senders.NewSender(conn),
		handlers:  map[string]handlers.MessageHandler{},
		consumers: map[string]*consumers.Consumer{},
		ch:        make(chan []byte),
	}
}

func (s *Server) InitHandlers() {
	s.handlers["UpdateOrderEventHandler"] = handlers.NewUpdateOrderEventHandler(s.senders, s.clRedis, s.ch)
	s.handlers["GetExchangeRateHandler"] = handlers.NewGetExchangeRateHandler(s.senders, s.clRedis)
}

func (s *Server) InitConsumers() {
	s.consumers["UpdateOrderEventConsumer"] = consumers.NewConsumer(
		s.rmqConn,
		config.RabbitEventsExchange,
		config.UpdatedOrderEventRoutingKey,
		config.UpdatedOrderEventQueueName,
		s.handlers["UpdateOrderEventHandler"],
		nil)

	s.consumers["GetExchangeRateConsumer"] = consumers.NewConsumer(
		s.rmqConn,
		config.RabbitExchangeRateExchange,
		config.GetExchangeRateRequestRoutingKey,
		config.CreateOrderRequestQueueName,
		s.handlers["GetExchangeRateHandler"],
		s.ch)
}

func (s *Server) Run() {
	forever := make(chan bool)
	s.runAllConsumers()
	<-forever
}

func (s *Server) runAllConsumers() {
	for _, client := range s.consumers {
		go client.ConsumeMessages()
	}
}
