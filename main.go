package main

import (
	"github.com/TeoDev1611/fil/utils"
	"github.com/integrii/flaggy"
)

var (
	VERSION            = "unknow"
	JsonSubCommand     *flaggy.Subcommand
	OpenTheRepoBrowser bool
)

func init() {
	flaggy.SetName("Fil < Files is live >")
	flaggy.SetDescription(`A simple file parser formatter and pretty printer :D`)
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/TeoDev1611/fil"

	flaggy.Bool(&OpenTheRepoBrowser, "o", "open", "Open the browser for show the repo and more help")

	JsonSubCommand = flaggy.NewSubcommand("json")
	JsonSubCommand.Description = "The json parser utils"

	flaggy.SetVersion(VERSION)

	flaggy.Parse()
}

func main() {
	if JsonSubCommand.Used {
		println("asdasd")
	}

	if OpenTheRepoBrowser {
		utils.OpenBrowser("https://github.com/TeoDev1611/fil")
	}
}
