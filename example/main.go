package main

import (
	"fmt"
	"github.com/eurozulu/commandline"
	"log"
	"strings"
)

type cmdFunc func(args ...string) error

var commands = map[string]cmdFunc{
	"this": doThis,
	"that": doThat,
}

func main() {
	cli := commandline.NewCommandLine()
	if err := cli.LoadHistory("$PWD/.history"); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := cli.SaveHistory("$PWD/.history"); err != nil {
			log.Println(err)
		}
	}()

	for {
		fmt.Print(">")
		cmdLine, err := cli.ReadCommand()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
		if cmdLine == "" {
			continue
		}
		cmds := strings.SplitN(cmdLine, " ", 2)
		if strings.EqualFold(cmds[0], "exit") {
			return
		}
		fn, ok := commands[strings.ToLower(cmds[0])]
		if !ok {
			fmt.Printf("%q is not a known command\n", cmds[0])
			continue
		}
		if err := fn(cmds[1:]...); err != nil {
			log.Fatalln(err)
		}
	}

}

func doThis(args ...string) error {
	fmt.Printf("doing this: %s\n", strings.Join(args, " "))
	return nil
}

func doThat(args ...string) error {
	fmt.Printf("doing that: %s\n", strings.Join(args, " "))
	return nil
}
