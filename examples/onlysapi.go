package main

import (
	"fmt"
	"github.com/SeanTolstoyevski/talker"
)

func main() {
	talker.PreferSAPI(false)
	talker.TrySAPI(true)
	talker.Load()
	talker.Output("This is test...", false)
	fmt.Println("currently working screen reader:", talker.DetectScreenReader())
}
