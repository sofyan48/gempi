package main

import (
	"github.com/joho/godotenv"
	"github.com/sofyan48/gempi/libs"
)

func main() {
	godotenv.Load()
	libs.Publis()
	libs.ConsumeMessages()
}
