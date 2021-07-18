package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
)

// 这里填你 OneIndex 图床服务的地址
const url = "https://www.oneindex.com/images"

func upload(url string, postImgPath string) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("file", postImgPath)

	file, _ := os.Open(postImgPath)
	defer file.Close()

	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, _ := http.Post(url, contentType, bodyBuffer)
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	root, err := htmlquery.Parse(strings.NewReader(string(respBody)))
	if err != nil {
		fmt.Errorf("parse error: %s", err)
	}

	//find := htmlquery.FindOne(root, "/html/body/div/div/div[1]/input")
	//list := htmlquery.FindOne(root, "//input[@type='text']")
	item := htmlquery.FindOne(root, "/html/body/div[1]/div/div[1]/input/@value")

	//log.Println(htmlquery.OutputHTML(find, true))
	//log.Println(htmlquery.OutputHTML(list, true))
	fmt.Println(htmlquery.InnerText(item))

}

func main() {
	//postImgPath := "/Applications/Typora.app/Contents/Resources/TypeMark/assets/icon/icon_256x256.png"
	args := os.Args[1:]
	for _, v := range args {
		upload(url, v)
	}

}
