package main

import (
	"fmt"
	"mini/dev"
	"mini/generator"
	"mini/helpers"
	"os"
)

func main() {
	const usageHelp = `
commands: 
	mini gen  -- generates web page
	mini init -- initializes folder structure
	mini serve -- serve content for development 
	
usage:
	To create a webpage run 'mini init' inside 
	the directory you want the page to be created.

	'mini init' can also be used to update information
	
	Create posts with the format: 
	'postnumber-name-year.md' 
	
	Run mini gen to generate page 
	Serve folder
	`

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usageHelp)
		os.Exit(0)
	}

	switch args[0] {
	case "gen":
		mini := generator.New()
		mini.GeneratePage()
		fmt.Println("Page generated into ./webpage")
	case "init":
		helpers.InitConf()
	case "serve":
		dev.Serve()
	default:
		fmt.Println(usageHelp)
	}
}
