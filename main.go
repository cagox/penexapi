package main

import (
	"fmt"

	"penexapi/app"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(app.Config.SiteName)

}
