package tumblr

import (
	"testing"
	"os"
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

//type Post struct {
//	PhotoUrl  []string `xml:"photo-url"`
//	Tumblelog string   `xml:"tumblelog"`
//}
//
type Result struct {
	XMLName xml.Name `xml:"tumblr"`
	//Posts   []Post   `xml:"posts>post"`
}

func TestReadFile(t *testing.T) {
	xmlFile, _ := os.Open("E:\\Develop\\crawler\\read.xml")
	result := new(Result)
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, result)
	fmt.Print(result)
}
