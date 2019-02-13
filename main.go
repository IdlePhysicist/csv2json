package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	flag "github.com/spf13/pflag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	to_skip := flag.IntP("skip", "s", 1, "Skip the first how many rows")
	header  := flag.IntP("header", "h", 0, "Specify the header row for the input file")
	flag.Parse()

	args := flag.Args()
	path := args[0]

	fmt.Println("Path to file:",path, " Skip:", *to_skip, "Header:", *header)

	fileBytes, fileNPath := ReadCSV(path, *to_skip, *header)
	SaveFile(fileBytes, fileNPath)
	fmt.Println(strings.Repeat("=", 10), "Done", strings.Repeat("=", 10))
}

func ReadCSV(path string, to_skip int, header int) ([]byte, string) {
	// ReadCSV to read the content of CSV File
	csvFile, err := os.Open(path)
	if err != nil { log.Fatal("The file is not found || wrong root") }
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var content [][]string
	i := 0
	for {
		line, err := reader.Read()
		if err == io.EOF { break }
		if i < to_skip == true { // Skip rows
			i++
			continue
		} else {
			content = append(content, line)
			//fmt.Println(line) // For testing
			i++
		}
		//if err != nil { log.Fatal("Error: ",err) }
	}

	if len(content) < 1 { log.Fatal("Error: ", err ) }

	headersArr := make([]string, 0)
	for _, headE := range content[header-1] {
		headersArr = append(headersArr, headE)
	}

	// Remove the header row
	content = content[header-1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	x, _ := json.MarshalIndent(rawMessage, "", "  ")
	newFileName := filepath.Base(path)
	newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
	r := filepath.Dir(path)
	return x, filepath.Join(r, newFileName)
}

func SaveFile(myFile []byte, path string) {
	// SaveFile Will Save the file, magic right?
	if err := ioutil.WriteFile(path, myFile, os.FileMode(0644)); err != nil {
		panic(err)
	}
}
