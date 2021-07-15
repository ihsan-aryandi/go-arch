package env

import (
	"github.com/joho/godotenv"
)

type Env interface {
	ReadEnv(fileNames ...string) error
}

type env struct {}

func NewEnv() Env {
	return &env{}
}

func (*env) ReadEnv(fileNames ...string) error {
	return godotenv.Load(fileNames...)
}
