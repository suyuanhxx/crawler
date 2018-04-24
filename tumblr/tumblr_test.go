package tumblr

import (
	"testing"
	"os"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"time"
)

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

func TestTumblrCrawler_DownloadVideo(t *testing.T) {
	tu := New()
	tu.DownloadVideo("travelingcolors")
	time.Sleep(100000)

}
