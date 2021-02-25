package main

import (
	"wikitermia/wikitermia"

	"flag"
	"fmt"
	"os"
)

var VERSION string = "v0.1.0"

func main(){
	fmt.Println("          _ _    _ _                      _")
	fmt.Println("__      _(_) | _(_) |_ ___ _ __ _ __ ___ (_) __ _")
	fmt.Println("\\ \\ /\\ / / | |/ / | __/ _ \\ '__| '_ ` _ \\| |/ _` |")
	fmt.Println(" \\ V  V /| |   <| | ||  __/ |  | | | | | | | (_| |")
	fmt.Println("  \\_/\\_/ |_|_|\\_\\_|\\__\\___|_|  |_| |_| |_|_|\\__,_|", VERSION, "\n")

	options := ParseArguments(os.Args[1:])

	if options.Query == "" {
		fmt.Println("[\033[31mERROR\033[0m] Enter query")
		os.Exit(1)
	}

	fmt.Println("[\033[34mquery\033[0m]  ", options.Query)

	wikitermia.SetLang(options.Lang)

	search := wikitermia.Search(options.Query, options.Limit)

	if len(search) == 0 {
		fmt.Println("[\033[31msearch\033[0m]", len(search))
		fmt.Println("[\033[31mERROR\033[0m] Not found")
	} else {
		fmt.Println("[\033[34msearch\033[0m] ", "Find", len(search), "page")
		page := wikitermia.Page(search[0].Pageid)

		fmt.Println("[\033[34mextract\033[0m]", page.Extract)
	}
}

type Options struct {
	Query   string
	Lang    string
	Version bool
	Limit   string
}

func ParseArguments(arguments []string) (Options) {
    var defaultOptions Options

    options := defaultOptions

    fs := flag.NewFlagSet("main", flag.ExitOnError)

    fs.StringVar(&options.Lang, "l", "en", "Select language")
    fs.BoolVar(&options.Version, "v", false, "Show version")
    fs.StringVar(&options.Limit, "limit", "1", "Limit")
    fs.StringVar(&options.Query, "q", "", "Query")

    fs.Parse(arguments)

    return options
}
