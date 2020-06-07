package main

import (
	"fmt"

	_ "github.com/uesrname/project/common"
	_ "github.com/uesrname/project/plugins"
	"github.com/uesrname/project/utils"
)

func main() {
	utils.Add(12, 13)
	fmt.Println("hello")
}
