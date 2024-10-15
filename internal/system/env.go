package system

import (
	"log"
	"os"
)

var (
	PORT         = mustExist("PORT")
	SERVICE_NAME = mustExist("SERVICE_NAME")
	DB_PATH      = mustExist("DB_PATH")
)

func mustExist(envVar string) string {
	val, exist := os.LookupEnv(envVar)
	if !exist {
		log.Panicf("missing env var ['%v'], make sure it is set", envVar)
	}
	return val
}
