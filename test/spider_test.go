package test

import (
	"fmt"
	"net/http"
	"testing"
	"wallpaper/common"
	"wallpaper/spider/msbing"
	"wallpaper/spider/reddit"
)

// TestReddit 首页
func TestRedditIndex(t *testing.T) {
	reddit.SpiderIndex("./pic")
}

// TestRedditPage 分页
func TestRedditPage(t *testing.T) {
	reddit.SpiderPage("./pic", "t3_xerjfq")
}
func TestDownloadImage(t *testing.T) {
	img := "https://p4.ssl.qhimg.com/t01ee06478bb11783ee.jpg"
	// "https://i.redd.it/tfbtm8c2wep91.png"
	fmt.Println(common.DownloadImg(img, "t01ee06478bb11783ee.jpg", "./", nil))
}

func TestSendRequest(t *testing.T) {
	common.SendRequest(http.MethodGet, "https://www.360.cn", nil, 500, nil)
}

func TestRequestGetByCert(t *testing.T) {
	resp, err := common.RequestGetByCert("https://www.reddit.com/r/wallpapers.json", []byte(reddit.CaCert), 5000, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp))
}

func TestBing(t *testing.T) {
	msbing.Spider("./pic")
}
