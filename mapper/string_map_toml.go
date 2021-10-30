package mapper

import (
	"github.com/TeoDev1611/fil/utils"
	"github.com/pelletier/go-toml"
)

var dataMapperToml map[string]interface{}

func FileToMapperToml(file string) map[string]interface{} {
	fileData := utils.OpenFiles(file)
	err := toml.Unmarshal([]byte(fileData), &dataMapperToml)
	utils.CheckErrors(err)
	return dataMapperToml
}
