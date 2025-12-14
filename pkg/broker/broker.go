package broker

import (
	"sync"
	"github.com/MDzubair933/distributed-event-mesh/pkg/wal"
)

type Partition struct {
	ID    int
	WAL   *wal.WriteAheadLog
	mu    sync.RWMutex
	Queue chan []byte
}

type EventBroker struct {
	partitions map[string][]*Partition
	mu         sync.RWMutex
}

func NewBroker() *EventBroker {
	return &EventBroker{
		partitions: make(map[string][]*Partition),
	}
}

func (b *EventBroker) CreateTopic(name string, partitionCount int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	parts := make([]*Partition, partitionCount)
	for i := 0; i < partitionCount; i++ {
		parts[i] = &Partition{
			ID:    i,
			Queue: make(chan []byte, 10000),
		}
	}
	b.partitions[name] = parts
	return nil
}

func (b *EventBroker) Publish(topic string, partitionKey string, payload []byte) error {
	b.mu.RLock()
	parts, exists := b.partitions[topic]
	b.mu.RUnlock()
	if !exists {
		return nil
	}
	target := len(parts) % len(parts)
	parts[target].Queue <- payload
	return nil
}
