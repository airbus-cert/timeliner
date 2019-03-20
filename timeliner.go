package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.cert.corp/nbareil/bodyfile"
)

var filter = flag.String("filter", "", "Event filter, like \"hour > 14\"")
var strict = flag.Bool("strict", false, "Only show the entries maching the date restrictions")

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		os.Exit(1)
	}
	filename := flag.Arg(0)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	body := bodyfile.NewReader(f)
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
