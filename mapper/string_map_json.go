package mapper

import (
	"encoding/json"

	"github.com/TeoDev1611/fil/utils"
)

var dataMapper map[string]interface{}

func FileToMapper(file string) map[string]interface{} {
	fileData := utils.OpenFiles(file)
	byteData := []byte(fileData)
	err := json.Unmarshal(byteData, &dataMapper)
	utils.CheckErrors(err)
	return dataMapper
}
