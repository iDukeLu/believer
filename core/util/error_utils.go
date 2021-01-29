package util

import "log"

func Panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
