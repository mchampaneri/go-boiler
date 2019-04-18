package main

import (
	"os"
	"log"
)

// Log Structure
// Used by the logging system
type Log struct {
}

// DefaultLogger is sington which 
// is used for event & data logging purpose
var DefaultLogger Log

// Logger Configurations and functions
// DefaultLogger is the System Logger
func (l *Log) Info(info string) {
	f, err := os.OpenFile(Config.StoragePath+"logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Println("Info :", info)
}

func (l *Log) Warning(info string) {
	f, err := os.OpenFile(Config.StoragePath+"logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Println("Warning :", info)
}

func (l *Log) Error(info string) {
	f, err := os.OpenFile(Config.StoragePath+"logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Println("Error :", info)
}

func (l *Log) Track(info string) {
	f, err := os.OpenFile(Config.StoragePath+"logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Println("Track :", info)
}