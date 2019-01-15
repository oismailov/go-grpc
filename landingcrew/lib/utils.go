package lib

import (
	"encoding/json"
)

//ConvertStructToJson converts golang struct to json string.
func ConvertStructToJson(customStruct interface{}) string {
	b, err := json.Marshal(customStruct)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
