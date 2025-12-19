package tests

import (
	"testing"
	"github.com/MDzubair933/distributed-event-mesh/pkg/broker"
)

func TestTopicPartitioning(t *testing.T) {
	b := broker.NewBroker()
	if err := b.CreateTopic("telemetry.events", 4); err != nil {
		t.Fatalf("Failed to create topic: %v", err)
	}
	if err := b.Publish("telemetry.events", "sensor_01", []byte("payload_data")); err != nil {
		t.Fatalf("Publish failed: %v", err)
	}
}
