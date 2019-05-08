/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package config

import (
	"io/ioutil"
	"strconv"

	b64 "encoding/base64"

	"github.com/go-kit/kit/log"
	"gopkg.in/yaml.v2"
)

const (
	defConfigPath = "/config/config.yml"
)

var (
	logger log.Logger
)

// Load - Generic config loader.
func Load() (*Config, error) {
	// return loadFromSecretsPath()
	return loadFromEnvvar()

}

// loadDefault - Load default values
func loadDefault() (*Config, error) {
	cfg := Config{}
	// Repo
	cfg.Repo.Type = "mongodb"
	cfg.Repo.MongoDB.Host = "localhost"
	cfg.Repo.MongoDB.Port = 27017
	cfg.Repo.MongoDB.Db = "granica"
	cfg.Repo.MongoDB.User = "granica"
	cfg.Repo.MongoDB.Password = "granica"
	// Broker
	cfg.Broker.RabbitMQ.Host = "localhost"
	cfg.Broker.RabbitMQ.Port = 5672
	cfg.Broker.Type = "rabbitmq"
	cfg.Broker.RabbitMQ.User = "granica"
	cfg.Broker.RabbitMQ.Password = "ganica"
	cfg.Broker.RabbitMQ.Exchange = "user"
	cfg.Broker.RabbitMQ.Queue = "user"
	// Cache
	cfg.Cache.Type = "redis"
	cfg.Cache.Redis.Host = "localhost"
	cfg.Cache.Redis.Port = 8086
	cfg.Cache.Redis.User = "granica"
	cfg.Cache.Redis.Password = "granica"
	return &cfg, nil
}

// loadFromEnvvar - Load from envvars.
// TODO: Default values must be corrected after establishing the appropriate ones.
func loadFromEnvvar() (*Config, error) {
	// Repo
	repoType := "mongodb"
	mongodbHost := GetEnvOrDef("MONGODB_HOST", "localhost")
	mongodbPort, _ := strconv.Atoi(GetEnvOrDef("MONGODB_PORT", "27017"))
	mongodbDb := GetEnvOrDef("MONGODB_DB", "0")
	mongodbUser := GetEnvOrDef("MONGODB_USER", "granica")
	mongodbPassword := GetEnvOrDef("MONGODB_PASSWORD", "granica")
	// Broker
	brokerType := "rabbitmq"
	rabbitmqHost := GetEnvOrDef("RABBITMQ_HOST", "localhost")
	rabbitmqPort, _ := strconv.Atoi(GetEnvOrDef("RABBITMQ_PORT", "5672"))
	rabbitmqUser := GetEnvOrDef("RABBITMQ_USER", "granica")
	rabbitmqPassword := GetEnvOrDef("RABBITMQ_PASSWORD", "granica")
	rabbitmqExchange := GetEnvOrDef("RABBITMQ_EXCHANGE", "default")
	rabbitmqQueue := GetEnvOrDef("RABBITMQ_QUEUE", "main")
	// Cache
	cacheType := "redis"
	redisHost := GetEnvOrDef("REDIS_HOST", "localhost")
	redisPort, _ := strconv.Atoi(GetEnvOrDef("REDIS_PORT", "6379"))
	redisUser := GetEnvOrDef("REDIS_USER", "granica")
	redisPassword := GetEnvOrDef("REDIS_PASSWORD", "granica")
	redisExchange := GetEnvOrDef("REDIS_EXCHANGE", "default")
	redisQueue := GetEnvOrDef("REDIS_QUEUE", "main")

	mongodb := MongoDBConfig{
		Host:     mongodbHost,
		Port:     mongodbPort,
		Db:       mongodbDb,
		User:     mongodbUser,
		Password: mongodbPassword,
	}

	repo := RepoConfig{
		Type:    repoType,
		MongoDB: mongodb,
	}

	redis := RedisConfig{
		Host:     redisHost,
		Port:     redisPort,
		User:     redisUser,
		Password: redisPassword,
		Exchange: redisExchange,
		Queue:    redisQueue,
	}

	cache := CacheConfig{
		Type:  cacheType,
		Redis: redis,
	}

	rabbitmq := RabbitMQConfig{
		Host:     rabbitmqHost,
		Port:     rabbitmqPort,
		User:     rabbitmqUser,
		Password: rabbitmqPassword,
		Exchange: rabbitmqExchange,
		Queue:    rabbitmqQueue,
	}

	broker := BrokerConfig{
		Type:     brokerType,
		RabbitMQ: rabbitmq,
	}

	cfg := &Config{
		Repo:   repo,
		Broker: broker,
		Cache:  cache,
	}

	// fmt.Printf("[DEBUG] - Config: %+v", cfg)

	return cfg, nil
}

// loadFromSecretsPath - Load from k8s secrets mount path.
func loadFromSecretsPath() (*Config, error) {
	var cfg *Config
	fileBytes, err := ioutil.ReadFile(defConfigPath)
	if err != nil {
		// logger.Log("level", LogLevel.Error, "message", "File open error", "file", defConfigPath)
		return nil, err
	}

	configYAMLBytes, err := b64.StdEncoding.DecodeString(string(fileBytes))
	if err != nil {
		// logger.Log("level", LogLevel.Error, "message", "Error decoding config file", "file", defConfigPath)
		return cfg, err
	}

	err = yaml.Unmarshal(configYAMLBytes, &cfg)
	if err != nil {
		return nil, err
	}
	// logger.Log("level", LogLevel.Debug, "message", "Config", "file", cfg)
	return cfg, nil
}

// loadFromFile - Load from standard configuration yaml file.
func loadFromFile(filePath string) (*Config, error) {
	var cfg *Config
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		// logger.Log("level", LogLevel.Error, "message", "File open error", "file", filePath)
		return nil, err
	}
	err = yaml.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, err
	}
	// logger.Log("level", LogLevel.Debug, "message", "Config", "file", cfg)
	return cfg, nil
}
