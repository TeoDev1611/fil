package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/TeoDev1611/fil/utils"
	"github.com/TylerBrock/colorjson"
)

var objMap map[string]interface{}

func PrettyPrintFileJson(file string) {
	fileData := utils.OpenFiles(file)
	json.Unmarshal([]byte(fileData), &objMap)
	s, _ := colorjson.Marshal(objMap)
	fmt.Println(string(s))
}
