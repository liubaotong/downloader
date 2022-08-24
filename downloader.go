package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Downloader struct {
	url  string
	file string
}

func NewDownloader(urlString string) *Downloader {
	r := &Downloader{
		url: urlString,
	}

	u, _ := url.Parse(r.url)
	r.file = strings.TrimLeft(u.Path, "/")
	dir := filepath.Dir(r.file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	return r
}

func (d *Downloader) Download() {
	fmt.Printf("下载[%s]", d.url)

	data, err := d.getData()
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Create(d.file)
	file.Write(data)

	fmt.Println("完成")
}

func (d *Downloader) getData() ([]byte, error) {
	resp, err := http.Get(d.url)
	if err != nil {
		return []byte(""), nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
