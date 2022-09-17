package test

import (
	"testing"
	"wallpaper/common"
	"wallpaper/reddit"
)

//TestReddit 测试 字典表数树
func TestReddit(t *testing.T) {
	reddit.SpiderIndexPage()
}

func TestDownloadImage(t *testing.T) {
	img := "https://p4.ssl.qhimg.com/t01ee06478bb11783ee.jpg"
	common.DownloadImg(img, "./")
}
