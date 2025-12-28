package consumer

import "sync"

type ConsumerGroup struct {
	GroupID     string
	Members     map[string]bool
	Offsets     map[string]uint64
	mu          sync.RWMutex
}

func NewConsumerGroup(id string) *ConsumerGroup {
	return &ConsumerGroup{
		GroupID: id,
		Members: make(map[string]bool),
		Offsets: make(map[string]uint64),
	}
}

func (cg *ConsumerGroup) CommitOffset(topic string, offset uint64) {
	cg.mu.Lock()
	defer cg.mu.Unlock()
	cg.Offsets[topic] = offset
}
