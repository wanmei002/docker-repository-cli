package requests

import (
	"fmt"
	"os"
)

func printLog(params ...any) {
	val := os.Getenv("DEBUG")
	if val != "true" {
		return
	}

	fmt.Println(params...)
}
