package config

// Config - Configuration struct.
type Config struct {
	// App - App config values
	App AppConfig `yaml:"app"`
	// Repo - App repo config values.
	Repo RepoConfig `yaml:"repo"`
	// Broker - App repo config values.
	Broker BrokerConfig `yaml:"broker"`
	// Cache - App cache config values.
	Cache CacheConfig `yaml:"cache"`
}

// AppConfig - App configuration struct.
type AppConfig struct {
	LogLevel logLevel `yaml:"logLevel"`
}

// RepoConfig - Repo configuration struct.
type RepoConfig struct {
	Type    string        `yaml:"repo"`
	MongoDB MongoDBConfig `yaml:"mongoDB"`
}

// MongoDBConfig - MongoDBConfig configuration struct.
type MongoDBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// CacheConfig - Cache configuration struct.
type CacheConfig struct {
	Type  string      `yaml:"cache"`
	Redis RedisConfig `yaml:"redis"`
}

// RedisConfig - REdis configuration struct.
type RedisConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Exchange string `yaml:"exchange"`
	Queue    string `yaml:"queue"`
}

// BrokerConfig - Repo configuration struct.
type BrokerConfig struct {
	Type     string         `yaml:"broker"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
}

// RabbitMQConfig - RabbitMQ configuration struct.
type RabbitMQConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Exchange string `yaml:"exchange"`
	Queue    string `yaml:"queue"`
}

// LogLevel - App log level.
type logLevel string

// LogLevels - App log levels.
type LogLevels struct {
	// Debug - Debug log level.
	Debug logLevel
	// Info - Info log level.
	Info logLevel
	// Warn - Warn log level.
	Warn logLevel
	// Error - Error log level.
	Error logLevel
	// Fatal - Fatal log level.
	Fatal logLevel
}
