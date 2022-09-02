# wallpaper
## _目录说明_
- bing 获取bing每日图片


## 介绍
项目通过调用官方相关接口获取4k图片

执行程序下载今日bing的背景图片，默认图片会下载到当前目录 pic目录下。
下载完成后会将图片名称和Copyright信息写入到introduction.txt文件中。


## 使用
linux:
```sh
bing -download <your download dir> -group n
```
| Param | Type |README |
| ------ | ------ | ------ |
| -download |string| 参数不指定图片将下载到当前目录下的 pic目录中 |
| -group |y/n| 默认值为n，图片保存在当前目录下。y：图片会在当前目录下按月建立文件夹分组 |

