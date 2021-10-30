package mapper

import (
	"encoding/xml"

	"github.com/TeoDev1611/fil/utils"
)

var dataMapperXml map[string]interface{}

func FileToMapperXml(file string) map[string]interface{} {
	fileDataXml := utils.OpenFiles(file)
	err := xml.Unmarshal([]byte(fileDataXml), &dataMapperXml)
	utils.CheckErrors(err)
	return dataMapperXml
}
