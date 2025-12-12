package server

import (
	"context"

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
	topic := req.Message.Topic
	value := req.Message.Payload
	headers := req.Message.Headers

	clientId := ""
	if v, ok := headers["clientid"]; ok {
		clientId = v
	}
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
	return &pb.LoadedResponse{
		Hooks: []*pb.HookSpec{
			{
				Name: "message.publish",
			},
		},
	}, nil
}

func (s *ExHookServer) OnProviderUnloaded(ctx context.Context, _ *pb.ProviderUnloadedRequest) (*pb.EmptySuccess, error) {
	return &pb.EmptySuccess{}, nil
}
