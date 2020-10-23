package main

import (
	"flag"
	"fmt"
	"strings"
)

var name = flag.String("name", "Casey", "The name to greet")
var times = flag.Int("times", 1, "The number to times to greet")

func MakeMsg(n string, t int) string {
	var messages []string

	for i := 0; i < t; i++ {
		messages = append(messages, fmt.Sprintf("Hello %s", n))
	}

	return strings.Join(messages, "\n")
}

func main() {
	flag.Parse()

	println(MakeMsg(*name, *times))
}
