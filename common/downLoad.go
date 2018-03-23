package common

import (
	"net/http"
	"io/ioutil"
	"os"
	"io"
	"bytes"
	"strings"
)

func DownLoadImage(imageUrl string) {
	resp, _ := http.Get(imageUrl)
	body, _ := ioutil.ReadAll(resp.Body)

	out, _ := os.Create(getImageName(imageUrl))

	io.Copy(out, bytes.NewReader(body))
}

func getImageName(imageUrl string) string {
	i := strings.LastIndex(imageUrl, "/")

	j := strings.LastIndex(imageUrl, ".")

	return string(imageUrl[i:j])
}
