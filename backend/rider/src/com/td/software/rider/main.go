package main

import (
	"os"
	"rider/src/com/td/software/rider/application"
)

func main() {

	if err := application.Init(); err != nil {
		os.Exit(-1)
	}

}
