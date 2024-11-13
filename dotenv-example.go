package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type DBConfig struct {
	User     string `env:"username, required"`
	Password string `env:"password, required"`
	Host     string `env:"host, required"`
	Database string `env:"db"`
}

func (db DBConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@%s/%s", db.User, db.Password, db.Host, db.Database)
}

func main() {
	db := GetConfig[DBConfig]()

	fmt.Println(db.ConnectionString())
}

func GetConfig[T any]() T {
	godotenv.Load(".env")

	var cfg T
	if err := envconfig.Process(context.TODO(), &cfg); err != nil {
		panic(fmt.Errorf("failed to load config: %s", err.Error()))
	}

	return cfg
}
