// 获取bing 4k 高清壁纸

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var baseDomain = "https://cn.bing.com/"

type bingResponse struct {
	Images []BingImage `json:"images"`
}

// BingImage 图片信息
type BingImage struct {
	Startdate     string        `json:"startdate"`
	Fullstartdate string        `json:"fullstartdate"`
	Enddate       string        `json:"enddate"`
	URL           string        `json:"url"`
	Urlbase       string        `json:"urlbase"`
	Copyright     string        `json:"copyright"`
	Copyrightlink string        `json:"copyrightlink"`
	Title         string        `json:"title"`
	Quiz          string        `json:"quiz"`
	Wp            bool          `json:"wp"`
	Hsh           string        `json:"hsh"`
	Drk           int           `json:"drk"`
	Top           int           `json:"top"`
	Bot           int           `json:"bot"`
	Hs            []interface{} `json:"hs"`
}

var savePicDir string
var groupByMonth string

// https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=10&nc=1612409408851&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160
func main() {

	flag.StringVar(&savePicDir, "download", "./pic/", "文件夹路径")
	flag.StringVar(&groupByMonth, "group", "N", "文件夹路径")
	flag.Parse()

	i := len(savePicDir)
	if !os.IsPathSeparator(savePicDir[i-1]) {
		savePicDir += string(os.PathSeparator)
	}

	image, err := getImageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.ToUpper(groupByMonth) == "Y" {
		savePicDir += image.Startdate[:6]
	}
	_, err = os.Stat(savePicDir)
	if os.IsNotExist(err) {
		os.MkdirAll(savePicDir, 0666)
	}
	saveImage(image, savePicDir)
}

func getImageInfo() (*BingImage, error) {
	picNum := 1 //获取图片数
	nc := time.Now().UnixNano()
	baseURI := fmt.Sprintf("/HPImageArchive.aspx?format=js&idx=0&n=%d&nc=%d&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160", picNum, nc)
	reqURL := fmt.Sprintf("%s%s", baseDomain, baseURI)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	//
	var bingResp bingResponse
	if err := json.Unmarshal(body, &bingResp); err != nil {
		return nil, err
	}
	if len(bingResp.Images) > 1 {
		return nil, fmt.Errorf("获取图片数量错误")
	}
	return &bingResp.Images[0], nil
}

func saveImage(img *BingImage, savePicDir string) {
	imgURL := fmt.Sprintf("%s%s", baseDomain, img.URL)
	u, err := url.Parse(img.URL)
	if err != nil {
		fmt.Println(err)
	}
	q := u.Query()

	respImg, err := http.Get(imgURL)
	imgFilePath := fmt.Sprintf("%s/%s", savePicDir, q.Get("id"))
	_, err = os.Stat(imgFilePath)
	if !os.IsNotExist(err) {
		fmt.Println("文件已存在")
		return
	}

	imageFd, err := os.Create(imgFilePath)
	if err != nil {
		panic(err)
	}
	defer imageFd.Close()
	_, err = io.Copy(imageFd, respImg.Body)
	if err != nil {
		panic(err)
	}

	//
	copyrightInfoPath := fmt.Sprintf("%s/introduction.txt", savePicDir)
	content := fmt.Sprintf("%s\t%s\t%s\n", img.Startdate, q.Get("id"), img.Copyright)
	writeFile(copyrightInfoPath, []byte(content))
}

func writeFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
