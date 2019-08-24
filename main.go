package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
  "log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
  Data  [][]string
  Json  bytes.Buffer
  Path  string
  Head  []string
}

func main() {
	flag.Parse()
  file := &File{Path: flag.Args()[0]}

	err := file.Read()
  if err != nil {
		log.Printf("read: Error: %s", err)
  }

  err = file.JSONify()
  if err != nil {
		log.Printf("jsonify: Error: %s", err)
  }

	err = file.Write()
  if err != nil {
		log.Printf("write: Error: %s", err)
  }

  log.Print(`Done`)
}

func (f *File) Read() error {
	csvFile, err := os.Open(f.Path)
	if err != nil {
    return err
  }
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	f.Data, _ = reader.ReadAll()

	if len(f.Data) < 1 {
    return err
  }

	for _, col := range f.Data[0] {
		f.Head = append(f.Head, col)
	}

	// Remove the header row
	f.Data = f.Data[1:]

  return nil
}

func (f *File) JSONify() error {
	var buf bytes.Buffer

	buf.WriteString("[")

	for i, d := range f.Data {
		buf.WriteString("{")

    for j, y := range d {
			buf.WriteString(`"` + f.Head[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 64)
			_, bErr := strconv.ParseBool(y)

			if fErr == nil {
				buf.WriteString(y)
			} else if bErr == nil {
				buf.WriteString(strings.ToLower(y))
			} else {
				buf.WriteString((`"` + y + `"`))
			}

			// End of property
			if j < len(d)-1 {
				buf.WriteString(",")
			}

		}
		// End of object of the array
		buf.WriteString("}")
		if i < len(f.Data)-1 {
			buf.WriteString(",")
		}
	}

	buf.WriteString(`]`)

  // Add the buffer to the struct
  f.Json = buf

  return nil
}

func (f *File) Write() error {
	rawMessage := json.RawMessage(f.Json.String())
	jsonStr, err := json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		return err
	}

	// Figuring out the new file name
	dir, file := filepath.Split(f.Path)
	file = strings.Replace(file, filepath.Ext(file), ".json", 1)
	path := filepath.Join(dir, file)
	log.Println("write: New file: ", path)

	// Write the new file
	err = ioutil.WriteFile(path, jsonStr, 0644)
  if err != nil {
		return err
	}

	return nil
}
