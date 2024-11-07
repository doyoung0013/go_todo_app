package config

import (
	"github.com/caarlos0/env/v6"
)

// 환경 변수로부터 값을 읽어옴
type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

// New는 환경 변수로부터 설정값을 읽어와 Config 구조체를 생성하고 초기화하여 반환
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
