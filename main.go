package main

import (
)

func main() {
	go startServer()
	select {}
}