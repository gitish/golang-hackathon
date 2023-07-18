package web

import (
	"os"
)

var DIR_NAME = "/Users/shashail/data/go"

func GetDataDir() string {
	val, ok := os.LookupEnv("DATA_DIR")
	if ok {
		DIR_NAME = val
	}
	return DIR_NAME
}
