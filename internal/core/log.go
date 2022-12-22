package core

import log2 "log"

var log *log2.Logger

func Init(logger *log2.Logger) {
	log = logger
}
