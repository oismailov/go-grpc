package lib

import (
	"encoding/json"
	"log"
	"strconv"
)

//ConvertStructToJson converts golang struct to json string.
func ConvertStructToJson(customStruct interface{}) string {
	b, err := json.Marshal(customStruct)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

//ConvertStringToInt64 converts string to integer 64.
func ConvertStringToInt64(s string) int64 {
	int, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Fatalf("Cannot convert string to integer: %s", err)
	}

	return int
}
