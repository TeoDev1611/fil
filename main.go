package main

import (
	"github.com/TeoDev1611/fil/cmds"
	"github.com/TeoDev1611/fil/utils"
	"github.com/integrii/flaggy"
)

var (
	VERSION            = "unknow"
	JsonSubCommand     *flaggy.Subcommand
	OpenTheRepoBrowser bool
	JsonFileValue      = "defaultVal"
)

// init function  
func init() {
	flaggy.SetName("Fil < Files is live >")
	flaggy.SetDescription(`A simple file parser formatter and pretty printer :D`)
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/TeoDev1611/fil"

	flaggy.Bool(&OpenTheRepoBrowser, "o", "open", "Open the browser for show the repo and more help")

	JsonSubCommand = flaggy.NewSubcommand("json")
	JsonSubCommand.Description = "The json parser utils"
	JsonSubCommand.String(&JsonFileValue, "f", "file", "Open a file with json pretty print")

	flaggy.SetVersion(VERSION)
	flaggy.AttachSubcommand(JsonSubCommand, 1)

	flaggy.Parse()
}

// main function  
func main() {
	// Main flag open the browser
	if OpenTheRepoBrowser {
		utils.OpenBrowser("https://github.com/TeoDev1611/fil")
	}

	/*
		JsonSubCommand Flags functions
	*/
	// Open the json file and pretty print
	if JsonFileValue != "defaultVal" {
		cmds.PrettyPrintFileJson(JsonFileValue)
	}
}
