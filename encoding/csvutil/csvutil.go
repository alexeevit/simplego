package csvutil

import (
	"encoding/csv"
	"os"
)

/*

For UTF-8 BOM, csv.Reader.Read() will return error = "line 1, column 1: bare \" in non-quoted-field"

If you encounter this close the file and call again with stripBom = true

*/

func NewReader(path string, comma rune, stripBom bool) (*csv.Reader, *os.File, error) {
	var myCsv *csv.Reader
	var file *os.File
	file, err := os.Open(path)
	if err != nil {
		return myCsv, file, err
	}
	if stripBom {
		b3 := make([]byte, 3)
		_, err := file.Read(b3)
		if err != nil {
			return myCsv, file, err
		}
	}
	reader := csv.NewReader(file)
	reader.Comma = comma
	return reader, file, nil
}
