package main

import (
	"fmt"
	"myPackage/config"
	"myPackage/utils"
)

func main() {
	utils.Add(12, 13)
	fmt.Println("hello world")
	config.ConfNew()
}
