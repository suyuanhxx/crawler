package common

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"golang.org/x/net/proxy"
	"os"
)

func ProxyHttpGet(mediaUrl string) *http.Response {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1086", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	if resp, err := httpClient.Get(mediaUrl); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		data := string(body)
		fmt.Print(data)
		return resp
	}
	return nil
}
