// 获取bing 4k 高清壁纸

package main

import (
	"flag"
	"fmt"
	// "os"
)

var savePicDir string
var groupByMonth string
var platform string

func main() {
	flag.StringVar(&savePicDir, "download", "./pic/", "文件夹路径")
	flag.StringVar(&groupByMonth, "group", "N", "按月划分文件夹")
	flag.StringVar(&platform, "platform", "", "platform:bing,reddit")

	flag.Usage = func() {
		fmt.Println("Usage: [-download] [-group] -platform")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !isFlagPassed("platform") {
		flag.Usage()
		return
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
