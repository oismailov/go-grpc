package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

//DownloadDFile will download file from remote server.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
