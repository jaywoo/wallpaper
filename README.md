# wallpaper
## _目录说明_
- bing 获取bing壁纸代码


## 介绍
项目通过调用相关接口获取4k壁纸地址进行下载

执行程序下载今日bing壁纸（4k分辨率），默认图片会当前目录创建pic目录，将图片下载到pic目录。
修改下载目录参见[参数说明](https://github.com/jaywoo/wallpaper/#%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E)


下载完成后会将图片名称和Copyright信息写入到introduction.txt文件中。


## 使用
linux:
```sh
bing -download <your download dir> -group n
```
源码:
```sh
go run bing/main.go -group y -download ./pic_test
```
#### 参数说明
| Param | Type |README |
| ------ | ------ | ------ |
| -download |string| 参数不指定图片将下载到当前目录下的 pic目录中 |
| -group |y/n| 默认值为n，图片保存在当前目录下。y：图片会在当前目录下按月建立文件夹分组 |

#### 添加任务：
##### Linux使用Crond,例：
1. 打开命令行
2. 输入： crontab -e
3. 在打开的文档中输入：10 9 * * * /home/bing -downland $HOME/Pictures  #(每天9点10分执行命令)
4. 按esc-->按 Shift+: --> 输入 wq #(保存并退出)

> PS: mac os 应该也可以这样添加

##### Windows10 添加启动项:
1. 右键可执行程序，创建快捷方式
2. 右键快捷方式-->属性。 在目标框中可以添加下载目录（可选），例 ：
  
  ![image](https://user-images.githubusercontent.com/7802535/188383830-34f24711-f9ab-4e52-82f2-0fddd447e6b5.png)
  
3. 按 Windows 徽标键  + R，键入“shell:startup”，然后选择“确定”。这将打开“启动”文件夹。
4. 将该应用的快捷方式从文件位置复制并粘贴到“启动”文件夹中。

##### Windows10 添加定时任务:
1. 复制“任务计划程序”
2. 按 Windows 徽标键, --> Ctrl+v ，打开“任务计划程序”
3. 点击“创建基本任务...”

![image](https://user-images.githubusercontent.com/7802535/188389802-edc2d298-3658-4552-99a9-2c9dc096dc4b.png)

4. 按照提示完成操作即可






