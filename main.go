package main

import ("os"
"github.com/Dorivika/Go-Blockchain/blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()

}
