package config

import (
	"os"
	"time"
)

const (
	OtpSecret = "some_secret"
)

type webAddress struct {
	Host string
	Port string
}

type Config struct {
	APP         string
	Environment string
	LogLevel    string

	Server struct {
		Host         string
		Port         string
		ReadTimeout  string
		WriteTimeout string
		IdleTimeout  string
	}

	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSLMode  string
	}

	Context struct {
		Timeout string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Name     string
	}

	Token struct {
		Secret     string
		AccessTTL  time.Duration
		RefreshTTL time.Duration
	}

	Minio struct {
		Endpoint              string
		AccessKey             string
		SecretKey             string
		Location              string
		MovieUploadBucketName string
	}

	ClientService webAddress
	JobService    webAddress
	OTLPCollector webAddress
}

func NewConfig() (*Config, error) {
	var config Config

	// general configuration
	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// server configuration
	config.Server.Host = getEnv("SERVER_HOST", "api-gateway")
	config.Server.Port = getEnv("SERVER_PORT", ":5555")
	config.Server.ReadTimeout = getEnv("SERVER_READ_TIMEOUT", "10s")
	config.Server.WriteTimeout = getEnv("SERVER_WRITE_TIMEOUT", "10s")
	config.Server.IdleTimeout = getEnv("SERVER_IDLE_TIMEOUT", "120s")

	// db configuration
	config.DB.Host = getEnv("POSTGRES_HOST", "postgres")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "client_job_services_db")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "root")
	config.DB.SSLMode = getEnv("POSTGRES_SSLMODE", "disable")

	// redis configuration
	config.Redis.Host = getEnv("REDIS_HOST", "localhost")
	config.Redis.Port = getEnv("REDIS_PORT", "6379")
	config.Redis.Password = getEnv("REDIS_PASSWORD", "")
	config.Redis.Name = getEnv("REDIS_DATABASE", "0")

	// client-job services configuration
	config.ClientService.Host = getEnv("CLIENT_SERVICE_GRPC_HOST", "client-service")
	config.ClientService.Port = getEnv("CLIENT_SERVICE_GRPC_PORT", ":1111")
	config.JobService.Host = getEnv("JOB_SERVICE_GRPC_HOST", "job-service")
	config.JobService.Port = getEnv("JOB_SERVICE_GRPC_PORT", ":2222")

	// token configuration
	config.Token.Secret = getEnv("TOKEN_SECRET", "token_secret")

	// access ttl parse
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil, err
	}
	// refresh ttl parse
	refreshTTL, err := time.ParseDuration(getEnv("TOKEN_REFRESH_TTL", "24h"))
	if err != nil {
		return nil, err
	}
	config.Token.AccessTTL = accessTTl
	config.Token.RefreshTTL = refreshTTL

	// otlp collector configuration
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "localhost")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	// kafka configuration
	// config.Kafka.Address = strings.Split(getEnv("KAFKA_ADDRESS", "localhost:9092"), ",")
	// config.Kafka.Topic.UserCreateTopic = getEnv("KAFKA_USER_CREATE_TOPIC", "api.create.user")

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
