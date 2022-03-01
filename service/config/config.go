package config

import "github.com/joho/godotenv"

type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return err
}

func (c *Config) CatchError(err error) {
	if err != nil {
		panic(err)
	}
}
