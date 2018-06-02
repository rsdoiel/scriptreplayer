package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	// Caltech Library
	"github.com/caltechlibrary/tmplfn"
)

const (
	Version = `v0.0.0-hacking`
)

var (
	showHelp                 bool
	showVersion              bool
	timingFName              string
	typescriptFName          string
	templateFName            string
	jsonFName                string
	htmlFName                string
	srcTiming, srcTypescript []byte
	err                      error
	pageTitle                string
	cssPath                  string

	usage = `
USAGE

	%s [OPTIONS] TEMPLATE_FILENAME JSON_FILE

%s generates an HTML page suitable to show the output
from "script" and "scripreplay" commands in a browser.

EXAMPLE

	%s -t demo.timing -s demo.log demo.json > demo.html

`
)

type Timing struct {
	T float64 `json:"t"`
	C int     `json:"c"`
}

type Performance struct {
	Typescript string    `json:"typescript"`
	Timing     []*Timing `json:"timing"`
}

func (p *Performance) Parse(srcTiming []byte, srcTypescript []byte) error {
	var timing []*Timing

	lines := strings.Split(fmt.Sprintf("%s", srcTiming), "\n")
	for i, line := range lines {
		if strings.Contains(line, " ") {
			parts := strings.SplitN(line, " ", 2)
			f, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return fmt.Errorf("ERROR: can't parse float line %d, %s", i, err)
			}
			c, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("ERROR: can't parse int line %d, %s", i, err)
			}
			t := new(Timing)
			t.T = f
			t.C = c
			timing = append(timing, t)
		}
	}
	p.Timing = timing
	p.Typescript = fmt.Sprintf("%s", srcTypescript)
	return nil
}

func help(w io.Writer, appName, msg string, exitCode int) {
	fmt.Fprintf(w, usage, appName, appName, appName)
	flag.PrintDefaults()
	if msg != "" {
		fmt.Fprintf(w, "\n%s\n", msg)
	} else {
		fmt.Fprintf(w, "\n%s %s\n", appName, Version)
	}
	os.Exit(exitCode)
}

func main() {
	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&timingFName, "t", "", "(required) set script timing filename")
	flag.StringVar(&timingFName, "timing", "", "set script timing filename")
	flag.StringVar(&typescriptFName, "s", "", "(required) set script typescript filename")
	flag.StringVar(&typescriptFName, "typescript", "", "set script typescript filename")
	flag.StringVar(&pageTitle, "title", "", "(optional) set the HTML page title")
	flag.StringVar(&cssPath, "csspath", "", "(optional) set custom CSS path")
	flag.Parse()

	if showHelp {
		help(os.Stdout, appName, "", 0)
	}
	if showVersion {
		fmt.Printf("%s %s\n", appName, Version)
		os.Exit(0)
	}

	if timingFName == "" {
		log.Fatal("ERROR: Missing timing filename")
	}
	if typescriptFName == "" {
		log.Fatal("ERROR: Missing typescript filename")
	}

	srcTiming, err = ioutil.ReadFile(timingFName)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	srcTypescript, err = ioutil.ReadFile(typescriptFName)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	performance := new(Performance)
	err = performance.Parse(srcTiming, srcTypescript)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	srcPerformance, err := json.Marshal(performance)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	srcTemplate, err := ioutil.ReadFile("scriptreplayer.tmpl")
	if err != nil {
		log.Fatalf("ERROR: reading scriptreplayer.tmpl, %s", err)
	}

	extName := path.Ext(typescriptFName)
	jsonName := fmt.Sprintf("%s.json", strings.TrimSuffix(typescriptFName, extName))
	err = ioutil.WriteFile(jsonName, srcPerformance, 0664)
	if err != nil {
		log.Fatalf("ERROR: writing %s, %s", jsonName, err)
	}

	// Create our Tmpl with its function map
	tmpl := tmplfn.New(tmplfn.AllFuncs())
	tmpl.Add("scriptreplayer.tmpl", srcTemplate)

	t, err := tmpl.Assemble()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	data := map[string]interface{}{
		"performance": jsonName,
	}
	if pageTitle != "" {
		data["title"] = pageTitle
	}
	if cssPath != "" {
		data["csspath"] = cssPath
	}
	err = t.ExecuteTemplate(os.Stdout, "scriptreplayer.tmpl", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
