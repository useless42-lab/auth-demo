package main

import (
	"UserCenter/routes"
	"fmt"
)

func main() {
	routes.InitApiRoute()
	var str string
	fmt.Scan(&str)
}
