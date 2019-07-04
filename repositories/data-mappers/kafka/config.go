package kafka

type Config struct {
	Addrs    []string `conf:"env"`
	Net      NetConfig
	Producer ProducerConfig
	Consumer ConsumerConfig
}
type NetConfig struct {
	MaxOpenRequests int
	DialTimeout     int64
	ReadTimeout     int64
	WriteTimeout    int64
	KeepAlive       int64
}
type ProducerConfig struct {
	MaxMessageBytes int
	RequiredAcks    int
	Timeout         int64
	Compression     int8
	Return          ReturnConfig
	Flush           FlushConfig
	Retry           RetryConfig
}
type ReturnConfig struct {
	Successes bool
	Errors    bool
}
type FlushConfig struct {
	MaxMessages int
}
type RetryConfig struct {
	Max     int
	Backoff int64
}

type ConsumerConfig struct {
	Fetch             FetchConfig
	MaxWaitTime       int64
	MaxProcessingTime int64
}
type FetchConfig struct {
	Min     int32
	Default int32
	Max     int32
}
