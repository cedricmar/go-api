package utils

import "log"

// LogFatal is shorthand for log.Fatal
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Panic is shorthand for panic
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
