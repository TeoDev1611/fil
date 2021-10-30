package cmds

import (
	"fmt"

	"github.com/TeoDev1611/fil/utils"
	"github.com/fatih/color"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/printer"
	"github.com/mattn/go-colorable"
)

/*
	Code adapted from https://github.com/goccy/go-yaml/blob/master/cmd/ycat/ycat.go thanks for your amazing work
*/

const escape = "\x1b"

func format(attr color.Attribute) string {
	return fmt.Sprintf("%s[%dm", escape, attr)
}

func YcatPrintter(file string) error {
	tokens := lexer.Tokenize(string(file))
	var p printer.Printer
	p.LineNumber = true
	p.LineNumberFormat = func(num int) string {
		fn := color.New(color.Bold, color.FgHiWhite).SprintFunc()
		return fn(fmt.Sprintf("%2d | ", num))
	}
	p.Bool = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiMagenta),
			Suffix: format(color.Reset),
		}
	}
	p.Number = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiMagenta),
			Suffix: format(color.Reset),
		}
	}
	p.MapKey = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiCyan),
			Suffix: format(color.Reset),
		}
	}
	p.Anchor = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiYellow),
			Suffix: format(color.Reset),
		}
	}
	p.Alias = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiYellow),
			Suffix: format(color.Reset),
		}
	}
	p.String = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiGreen),
			Suffix: format(color.Reset),
		}
	}
	writer := colorable.NewColorableStdout()
	writer.Write([]byte(p.PrintTokens(tokens) + "\n"))
	return nil
}

func PrettyPrintYamlFile(file string) {
	fileData := utils.OpenFiles(file)
	err := YcatPrintter(fileData)
	utils.CheckErrors(err)
}
