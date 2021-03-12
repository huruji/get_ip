package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := &cli.App{
		Name: "getIP",
		Usage: "get the ip of your net",
		Action: func(context *cli.Context) error {
			var ip []byte
			var err error
			cmd := exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
			if ip, err = cmd.Output(); err != nil {
				fmt.Println(err)
				os.Exit(1)
				return nil
			}
			fmt.Println(string(ip))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}