package config

import "time"

type Config struct {
	BrokerHost       string        `yaml:"broker_host"`
	BrokerPort       int           `yaml:"broker_port"`
	WALDirectory     string        `yaml:"wal_directory"`
	SyncInterval     time.Duration `yaml:"sync_interval"`
	MaxSegmentSizeMB int           `yaml:"max_segment_size_mb"`
}

func DefaultConfig() *Config {
	return &Config{
		BrokerHost:       "0.0.0.0",
		BrokerPort:       9092,
		WALDirectory:     "/var/data/event-mesh/wal",
		SyncInterval:     500 * time.Millisecond,
		MaxSegmentSizeMB: 1024,
	}
}
