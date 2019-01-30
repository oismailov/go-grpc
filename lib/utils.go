package lib

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
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

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}

func FindReplaceText(files []string, replace string) {
	for _, file := range files {
		funcMap := template.FuncMap{
			"Name": strings.Title,
		}

		fi, err := os.Stat(file)

		if err != nil {
			log.Fatal(err)
		}

		if fi.Mode().IsDir() {
			continue
		}

		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Could not read file %s", err)
		}

		fileContent := string(bytes)

		tmpl, err := template.New("").Funcs(funcMap).Parse(fileContent)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}

		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("Could not create file: %s", err)
		}

		err = tmpl.Execute(f, replace)
		if err != nil {
			log.Fatalf("Could not replace text: %s", err)
		}
	}
}
