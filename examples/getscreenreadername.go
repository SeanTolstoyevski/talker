package main

import (
	"fmt"
	"github.com/SeanTolstoyevski/talker"
)

func main() {
	talker.Load()
	fmt.Println("currently working screen reader:", talker.DetectScreenReader())
}
