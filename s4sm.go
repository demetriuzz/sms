package main

import (
	"flag"
	"fmt"
)

var quietFlag = flag.Bool("q", false, FLAG_Q_DESC)
var confFlag = flag.String("conf", "", FLAG_CONF_DESC)

func main() {
	flag.Parse()
	if !*quietFlag {
		fmt.Println("configation file:", *confFlag)
	}
}
