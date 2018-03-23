package common

import (
	"io/ioutil"
	"os"
	"io"
	"bytes"
	"strings"
	"fmt"
)

func DownLoadImage(imageUrl string, dir string) {
	fmt.Println(imageUrl)
	resp := ProxyHttpGet(imageUrl)
	body, _ := ioutil.ReadAll(resp.Body)

	out, _ := os.Create(dir + getImageName(imageUrl))

	io.Copy(out, bytes.NewReader(body))
}

func getImageName(imageUrl string) string {
	i := strings.LastIndex(imageUrl, "/")
	return string(imageUrl[i:])
}
