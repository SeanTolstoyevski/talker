package main

import (
	"fmt"
	"github.com/SeanTolstoyevski/talker"
)

func main() {
	talker.Load()
	talker.Speak("This is test...", true)
	fmt.Println("Currently screen reader support print braille", talker.HasBraille(), "and speech", talker.HasSpeech())
}