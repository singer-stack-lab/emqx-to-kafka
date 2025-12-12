package server

import (
	"context"
	"log"

	pb "github.com/singer-stack-lab/emqx-to-kafka/gen/go/proto"
)

type ExHookServer struct {
	producer *KafkaProducer
	router   *TopicRouter
	pb.UnimplementedHookProviderServer
}

func NewExHookServer(p *KafkaProducer, r *TopicRouter) *ExHookServer {
	return &ExHookServer{
		producer: p,
		router:   r,
	}
}

func (s *ExHookServer) OnMessagePublish(ctx context.Context, req *pb.MessagePublishRequest) (*pb.ValuedResponse, error) {
	log.Println("OnMessagePublish", req)
	topic := req.Message.Topic
	value := req.Message.Payload

	clientId := req.Message.GetFrom()
	ok, kafkaTopic := s.router.Map(topic)
	if ok {
		// 不转发
		_ = s.producer.Send(kafkaTopic, clientId, value)
	}

	return &pb.ValuedResponse{
		Type: pb.ValuedResponse_CONTINUE,
		Value: &pb.ValuedResponse_Message{
			Message: req.Message,
		},
	}, nil
}

func (s *ExHookServer) OnProviderLoaded(ctx context.Context, _ *pb.ProviderLoadedRequest) (*pb.LoadedResponse, error) {
	log.Println("OnProviderLoaded")
	return &pb.LoadedResponse{
		Hooks: []*pb.HookSpec{
			{
				Name: "message.publish",
			},
		},
	}, nil
}

func (s *ExHookServer) OnProviderUnloaded(ctx context.Context, _ *pb.ProviderUnloadedRequest) (*pb.EmptySuccess, error) {
	log.Println("OnProviderUnloaded")
	return &pb.EmptySuccess{}, nil
}
