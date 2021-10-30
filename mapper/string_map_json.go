package mapper

import (
	"encoding/json"

	"github.com/TeoDev1611/fil/utils"
)

var dataMapper map[string]interface{}

func FileToMapperJson(file string) map[string]interface{} {
	fileData := utils.OpenFiles(file)
	err := json.Unmarshal([]byte(fileData), &dataMapper)
	utils.CheckErrors(err)
	return dataMapper
}
