package config

import (
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
var (
	version    string = "dev"
	commitHash string = "-"
)

// Project contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// GRPC contains all parameters of gRPC.
type GRPC struct {
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	MaxConnectionIdle time.Duration `yaml:"maxConnectionIdle"`
	Timeout           time.Duration `yaml:"timeout"`
	MaxConnectionAge  time.Duration `yaml:"maxConnectionAge"`
}

// REST contains all parameters of REST.
type REST struct {
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	AllowedOrigins []string `yaml:"allowedOrigins"`
	AllowedMethods []string `yaml:"allowedMethods"`
}

// Jaeger contains all parameters tracer information.
type Jaeger struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Service string `yaml:"service"`
}

// Metrics contains all parameters metrics information.
type Metrics struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Path string `yaml:"path"`
}

// Status contains all parameters status information.
type Status struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
	VersionPath   string `yaml:"versionPath"`
	LoggerPath    string `yaml:"loggerPath"`
	SwaggerDir    string `yaml:"swaggerDir"`
	SwaggerPath   string `yaml:"swaggerPath"`
}

// Config contains all configuration parameters in the config package.
type Config struct {
	Project Project    `yaml:"project"`
	Logger  zap.Config `yaml:"logger"`
	GRPC    GRPC       `yaml:"grpc"`
	REST    REST       `yaml:"rest"`
	Jaeger  Jaeger     `yaml:"jaeger"`
	Metrics Metrics    `yaml:"metrics"`
	Status  Status     `yaml:"status"`
}

// NewConfigWithYAML loads config from yaml.
func NewConfigWithYAML(configYML string) (*Config, error) {
	file, err := os.Open(configYML)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg *Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return cfg, nil
}
