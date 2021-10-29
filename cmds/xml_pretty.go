package cmds

import (
	"fmt"
	"log"

	"github.com/TeoDev1611/fil/utils"
	"github.com/go-xmlfmt/xmlfmt"
)

func PrettyFormatXmlFile(file string) {
	fileData := utils.OpenFiles(file)
	x := xmlfmt.FormatXML(fileData, "  ", "  ")
	fmt.Println(x)
	log.Print("Currently not support the XML color format :c")
}
