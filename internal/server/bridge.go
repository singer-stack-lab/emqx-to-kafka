package server

import (
	"log"
	"net"

	"github.com/singer-stack-lab/emqx-to-kafka/config"
	pb "github.com/singer-stack-lab/emqx-to-kafka/gen/go/proto"

	"github.com/IBM/sarama"
	"google.golang.org/grpc"
)

type Bridge struct {
	cfg    *config.Config
	server *grpc.Server
}

func NewBridge(cfg *config.Config, kafkaCli sarama.Client) *Bridge {
	router := NewTopicRouter(cfg.Rules)
	producer, err := NewKafkaProducer(kafkaCli)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	exServer := NewExHookServer(producer, router)
	pb.RegisterHookProviderServer(s, exServer)

	return &Bridge{
		cfg:    cfg,
		server: s,
	}
}

func (b *Bridge) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("listening on %s", addr)
	return b.server.Serve(lis)
}
