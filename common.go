package main

import "log"

func ExitIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
