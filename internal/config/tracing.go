package config

// import (
// 	"os"

// 	"github.com/elastic/go-elasticsearch/v8"
// 	"github.com/sirupsen/logrus"
// )

// type ELKConfig struct {
// 	ElasticURL string
// }

// func NewELKConfig() *ELKConfig {
// 	return &ELKConfig{
// 		ElasticURL: getEnv("ELASTIC_URL", "http://localhost:9200"),
// 	}
// }

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

// func NewElasticClient(cfg *ELKConfig) (*elasticsearch.Client, error) {
// 	es, err := elasticsearch.NewClient(elasticsearch.Config{
// 		Addresses: []string{cfg.ElasticURL},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return es, nil
// }

// func SetupLogger() *logrus.Logger {
// 	logger := logrus.New()
// 	logger.SetFormatter(&logrus.JSONFormatter{})
// 	return logger
// }