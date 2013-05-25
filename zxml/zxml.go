// zxml.go
package zxml

import (
	"encoding/xml"
	"fmt"
	"os"
	//"errors"
)

const (
	Extensions = iota
)

type Query struct {
	XMLName xml.Name `xml:config`
	ExtensionList []string `xml:"extensions>t"`
}
type Extension struct {
	t string
}

func XmlOps(t int, filePath string) (ext []string, e error) {
	var q Query
	switch t {
	case Extensions:
		xmlFile, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer xmlFile.Close()

		//err := xml.Unmarshal([]byte(data), &q)
		errU := xml.NewDecoder(xmlFile).Decode(&q)
		if errU != nil {
			return nil, errU
		}

	default:
		fmt.Printf("Not a valid option: %v", t)
	}
	return q.ExtensionList, nil
}
