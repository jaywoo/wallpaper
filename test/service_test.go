package test

import (
	"fmt"
	"testing"
	"wallpaper/common"
	"wallpaper/reddit"
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
	fmt.Println(common.DownloadImg(img, "./"))
}
