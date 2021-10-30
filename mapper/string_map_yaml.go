package mapper

import (
	"github.com/TeoDev1611/fil/utils"
	"github.com/goccy/go-yaml"
)

var dataMapperYaml map[string]interface{}

func FileToMapperYaml(file string) map[string]interface{} {
	fileDataYml := utils.OpenFiles(file)
	err := yaml.Unmarshal([]byte(fileDataYml), &dataMapperYaml)
	utils.CheckErrors(err)
	return dataMapper
}
