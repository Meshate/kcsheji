package main

import (
	"fmt"
	. "kcsheji"
	"kcsheji/router"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	router.Init()
	_ = G.Run(":8080")
}
