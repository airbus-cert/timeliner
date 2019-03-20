package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"

	"github.cert.corp/nbareil/bodyfile"
)

var filter = flag.String("filter", "", "Event filter, like \"hour > 14\"")
var strict = flag.Bool("strict", false, "Only show the entries maching the date restrictions")

func getInput() io.Reader {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		return os.Stdin
	}

	if flag.NArg() == 0 {
		flag.Usage()
	}

	filename := flag.Arg(0)

	if filename == "-" {
		return os.Stdin
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %s", filename, err)
		os.Exit(1)
	}

	return f
}

func main() {
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\t%s [options] MFT.txt\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	body := bodyfile.NewReader(getInput())
	if *filter != "" {
		err := body.AddFilter(*filter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not add filter: %s", err)
			os.Exit(2)
		}
	}

	body.Strict = *strict

	if _, err := body.Slurp(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not read all the content: %s", err)
		os.Exit(3)
	}

	prev := ""
	for {
		tsEntry, err := body.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while iterating: %s", err)
			os.Exit(4)
		}

		currentDate := tsEntry.Time.Format("2006-01-02")
		if currentDate == prev {
			fmt.Printf("           ")
		} else {
			fmt.Printf("%s ", currentDate)
		}
		fmt.Printf("%s: %s\n", tsEntry.Time.Format("15:04:05"), tsEntry.Entry.Name)
		prev = currentDate
	}
}
