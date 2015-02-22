package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

type AtticConfig struct {
	Layout    string
	Pages     []string
	InputDir  string
	OutputDir string
}

type AtticContext struct {
	ActivePage string
}

func (ctx AtticContext) IsActive(page string) bool {
	return ctx.ActivePage == page
}

var configFile = flag.String("config", "", "Configuration file.")

func main() {
	flag.Parse()

	if *configFile == "" {
		fmt.Println("Config file must be defined.")
		flag.Usage()
		os.Exit(1)
	}

	configBlob, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Could not read configuration file \"%s\"\n", *configFile)
		os.Exit(1)
	}

	config := AtticConfig{}
	err = json.Unmarshal(configBlob, &config)
	if err != nil {
		fmt.Printf("Configuration file is not valid JSON \"%s\"\n", *configFile)
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := AtticContext{}
	files := []string{path.Join(config.InputDir, config.Layout), ""}

	for _, page := range config.Pages {
		files[1] = path.Join(config.InputDir, page)

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			fmt.Printf("Could not parse template \"%s\"\n", files[1])
			fmt.Println(err)
			os.Exit(1)
		}

		outfilePath := path.Join(config.OutputDir, page)
		outfile, err := os.OpenFile(outfilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer outfile.Close()
		if err != nil {
			fmt.Printf("Could not open file \"%s\" for writing\n", outfilePath)
			fmt.Println(err)
			os.Exit(1)
		}

		ctx.ActivePage = page
		err = tmpl.Execute(outfile, ctx)
		if err != nil {
			fmt.Printf("Could not execute template \"%s\"\n", files[1])
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
