package main

import (
	"fmt"
	"mangacli/internal/argsHandler"
	"os"
)
func main() {
	action, args, err := argsHandler.Consume(os.Args)
	handleError(err)
	res, err := action.Execute(args)
	handleError(err)
	fmt.Print(res)
}

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
