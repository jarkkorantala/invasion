package main

import "log"

const debug = false

// Panic on error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Debug logging
func debugLog(logEntry string) {
	if debug {
		log.Printf(logEntry)
	}
}
