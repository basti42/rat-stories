package system

import (
	"log"
	"os"
)

var (
	PORT         = mustExist("PORT")
	SERVICE_NAME = mustExist("SERVICE_NAME")
	DB_HOST      = mustExist("DB_HOST")
	DB_PORT      = mustExist("DB_PORT")
	DB_NAME      = mustExist("DB_NAME")
	DB_USER      = mustExist("DB_USER")
	DB_PASSWORD  = mustExist("DB_PASSWORD")
)

func mustExist(envVar string) string {
	val, exist := os.LookupEnv(envVar)
	if !exist {
		log.Panicf("missing env var ['%v'], make sure it is set", envVar)
	}
	return val
}
