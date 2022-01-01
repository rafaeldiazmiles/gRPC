package main

import "log"

func Run() error {
	// responsible for intitialiazing and starting gRPC server
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
