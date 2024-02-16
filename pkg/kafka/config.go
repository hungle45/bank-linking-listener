package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/rcrowley/go-metrics"
)

const (
	// time sarama-cluster assumes the processing of an event may take
	defaultMaxProcessingTime = 1 * time.Second

	// producer flush configuration
	defaultFlushFrequency = 100 * time.Millisecond
	defaultFlushBytes     = 64 * 1024
)

// DefaultConfig creates a new config used per default
func DefaultConfig() *sarama.Config {
	metrics.UseNilMetrics = true

	config := sarama.NewConfig()

	// config.Version = sarama.V2_0_0_0

	// consumer configuration
	config.Consumer.Return.Errors = true
	config.Consumer.MaxProcessingTime = defaultMaxProcessingTime
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // this configures the initial offset for streams. Tables are always
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}

	// producer configuration
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = defaultFlushFrequency
	config.Producer.Flush.Bytes = defaultFlushBytes
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	return config
}
