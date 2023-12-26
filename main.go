package main

import ("os"
"github.com/Dorivika/Go-Blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()

}
