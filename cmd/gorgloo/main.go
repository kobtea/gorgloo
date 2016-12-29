package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"flag"
	"io"
)

const Version string = "1.0"
const (
	ExitCodeOK             = iota
	ExitCodeParseFlagError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("gorgloo", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	path := *flags.String("path", "./*.org", "filepath")
	version := *flags.Bool("version", false, "show version info")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	fmt.Println(path)
	fp, err := os.OpenFile("testdata/fixture.org", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	if version {
		fmt.Fprintf(c.errStream, "gorgloo version: %s", Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
