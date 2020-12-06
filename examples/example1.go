package main

import (
	"github.com/SeanTolstoyevski/talker"
)

func main() {
	talker.Load()
	talker.Speak("This is test...", true)
}