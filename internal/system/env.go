package system

import (
	"log"
	"os"
)

var (
	PORT              = mustExist("PORT")
	SERVICE_NAME      = mustExist("SERVICE_NAME")
	DB_SERVER         = mustExist("DB_SERVER")
	DB_PORT           = mustExist("DB_PORT")
	DATABASE          = mustExist("DATABASE")
	DATABASE_USER     = mustExist("DATABASE_USER")
	DATABASE_PASSWORD = mustExist("DATABASE_PASSWORD")
)

func mustExist(envVar string) string {
	val, exist := os.LookupEnv(envVar)
	if !exist {
		log.Panicf("missing env var ['%v'], make sure it is set", envVar)
	}
	return val
}
