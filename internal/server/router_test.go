package server

import (
	"testing"

	"github.com/singer-stack-lab/emqx-to-kafka/config"
)

func TestTopicRouter_Map_RegexDeviceSegment(t *testing.T) {
	router := NewTopicRouter([]config.MappingRule{
		{
			EmqxTopicPrefix: "/device/[^/]+/message/up/ivs_result",
			KafkaTopic:      "device_message_up_ivs_result",
		},
	})

	tests := []struct {
		name      string
		emqxTopic string
		wantOK    bool
	}{
		{name: "matches-asdf", emqxTopic: "/device/asdf/message/up/ivs_result", wantOK: true},
		{name: "matches-123", emqxTopic: "/device/123/message/up/ivs_result", wantOK: true},
		{name: "anchored", emqxTopic: "x/device/123/message/up/ivs_result", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, kafkaTopic := router.Map(tt.emqxTopic)
			if ok != tt.wantOK {
				t.Fatalf("Map(%q) ok=%v, want %v", tt.emqxTopic, ok, tt.wantOK)
			}
			if ok && kafkaTopic != "device_message_up_ivs_result" {
				t.Fatalf("Map(%q) kafkaTopic=%q, want %q", tt.emqxTopic, kafkaTopic, "device_message_up_ivs_result")
			}
		})
	}
}

func TestTopicRouter_Map_LiteralStillActsLikePrefix(t *testing.T) {
	router := NewTopicRouter([]config.MappingRule{
		{
			EmqxTopicPrefix: "/device/123/message/up/ivs_result",
			KafkaTopic:      "device_123_message_up_ivs_result",
		},
	})

	ok, kafkaTopic := router.Map("/device/123/message/up/ivs_result/extra")
	if !ok {
		t.Fatalf("expected match")
	}
	if kafkaTopic != "device_123_message_up_ivs_result" {
		t.Fatalf("kafkaTopic=%q, want %q", kafkaTopic, "device_123_message_up_ivs_result")
	}
}
