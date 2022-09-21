package msbing

// 获取微软件bing 4k 高清壁纸

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"wallpaper/common"
)

var baseDomain = "https://cn.bing.com/"

type bingResponse struct {
	Images []BingImgInfo `json:"images"`
}

// BingImgInfo 图片信息
type BingImgInfo struct {
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

//Spider 抓取
func Spider(savePicDir string) {
	images, err := getImageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, img := range images {
		imgURL := baseDomain + img.URL
		u, _ := url.Parse(imgURL)
		q := u.Query()
		err := common.DownloadImg(imgURL, q.Get("id"), savePicDir)
		fmt.Println(err)
	}
}

func getImageInfo() ([]BingImgInfo, error) {
	picNum := 10 //获取图片数 ，观察接口一页仅显示8个（猜测不会返回更多）
	nc := time.Now().UnixNano()
	baseURI := fmt.Sprintf("/HPImageArchive.aspx?format=js&idx=0&n=%d&nc=%d&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160", picNum, nc)
	reqURL := fmt.Sprintf("%s%s", baseDomain, baseURI)
	// fmt.Println(reqURL)
	resp, _ := common.SendRequest(http.MethodGet, reqURL, nil, 500)
	// resp := []byte(`{"images":[{"startdate":"20220920","fullstartdate":"202209201600","enddate":"20220921","url":"/th?id=OHR.SyltNordseeHoernum_ZH-CN6316415332_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.SyltNordseeHoernum_ZH-CN6316415332","copyright":"赫努姆的茅草屋顶, 德国叙尔特岛 (© Francesco Carovillano/eStock Photo)","copyrightlink":"/search?q=H%c3%b6rnum+(Sylt)&form=hpcapt&mkt=zh-cn","title":"在西尔特的最南端","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220920_SyltNordseeHoernum%22&FORM=HPQUIZ","wp":true,"hsh":"43dcec68d2759026a1cbedbfc9779220","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220919","fullstartdate":"202209191600","enddate":"20220920","url":"/th?id=OHR.SitkaOtters_ZH-CN4715326633_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.SitkaOtters_ZH-CN4715326633","copyright":"阿拉斯加州锡特卡海峡的海獭，美国 (© Robert Harding/Offset/Shutterstock)","copyrightlink":"/search?q=%e6%b5%b7%e7%8d%ad&form=hpcapt&mkt=zh-cn","title":"海上泰迪熊","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220919_SitkaOtters%22&FORM=HPQUIZ","wp":true,"hsh":"0ddb27529ac03cb45b42ebf032b78b20","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220918","fullstartdate":"202209181600","enddate":"20220919","url":"/th?id=OHR.SanMartinoVillage_ZH-CN4623104087_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.SanMartinoVillage_ZH-CN4623104087","copyright":"巴斯利卡塔的卡斯泰尔梅扎诺村，意大利 (© Roberto Moiola/Getty Images)","copyrightlink":"/search?q=%e5%b7%b4%e6%96%af%e5%88%a9%e5%8d%a1%e5%a1%94&form=hpcapt&mkt=zh-cn","title":"山峦间的光辉之城","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220918_SanMartinoVillage%22&FORM=HPQUIZ","wp":true,"hsh":"0f60af7fb556a3ce206726abb5963525","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220917","fullstartdate":"202209171600","enddate":"20220918","url":"/th?id=OHR.EmeraldYoho_ZH-CN4524610330_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.EmeraldYoho_ZH-CN4524610330","copyright":"幽鹤国家公园的翡翠湖, 加拿大不列颠哥伦比亚省 (© Cavan Images/Offset)","copyrightlink":"/search?q=%e5%b9%bd%e9%b9%a4%e5%9b%bd%e5%ae%b6%e5%85%ac%e5%9b%ad&form=hpcapt&mkt=zh-cn","title":"沉浸在大自然中","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220917_EmeraldYoho%22&FORM=HPQUIZ","wp":true,"hsh":"83df53244825d026164daf2fc3bf3986","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220916","fullstartdate":"202209161600","enddate":"20220917","url":"/th?id=OHR.BlackpoolBeach_ZH-CN2646268897_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.BlackpoolBeach_ZH-CN2646268897","copyright":"布莱克浦塔和中央码头，英国兰开夏郡 (© Bailey-Cooper Photography/Alamy)","copyrightlink":"/search?q=%e8%8b%b1%e5%9b%bd%e5%b8%83%e8%8e%b1%e5%85%8b%e6%b5%a6&form=hpcapt&mkt=zh-cn","title":"布莱克浦的灯光太棒了","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220916_BlackpoolBeach%22&FORM=HPQUIZ","wp":true,"hsh":"e31b52a06076cbee5e7ff89c77494c38","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220915","fullstartdate":"202209151600","enddate":"20220916","url":"/th?id=OHR.PianePuma_ZH-CN1482049046_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.PianePuma_ZH-CN1482049046","copyright":"百内国家公园中的一头美洲狮，智利巴塔哥尼亚 (© Ingo Arndt/Minden Pictures)","copyrightlink":"/search?q=%e7%be%8e%e6%b4%b2%e7%8b%ae&FORM=hpcapt&mkt=zh-cn","title":"敏捷而隐秘","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220915_PianePuma%22&FORM=HPQUIZ","wp":true,"hsh":"e26cdbfc792dcdcfc13e653a8a2ee2df","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220914","fullstartdate":"202209141600","enddate":"20220915","url":"/th?id=OHR.PyreneesPark_ZH-CN1341030921_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.PyreneesPark_ZH-CN1341030921","copyright":"罗兰豁口上空的银河，法国上比利牛斯省 (© SPANI Arnaud/Alamy)","copyrightlink":"/search?q=%e6%af%94%e5%88%a9%e7%89%9b%e6%96%af%e5%b1%b1&form=hpcapt&mkt=zh-cn","title":"从天而降的魔法","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220914_PyreneesPark%22&FORM=HPQUIZ","wp":true,"hsh":"41da8034eb97b40cbfcf4241ea76625f","drk":1,"top":1,"bot":1,"hs":[]},{"startdate":"20220913","fullstartdate":"202209131600","enddate":"20220914","url":"/th?id=OHR.MarbleCanyon_ZH-CN1066862981_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4","urlbase":"/th?id=OHR.MarbleCanyon_ZH-CN1066862981","copyright":"大理石峡谷中横跨科罗拉多河的纳瓦霍桥，美国亚利桑那州北部  (© trekandshoot/Alamy)","copyrightlink":"/search?q=%e5%a4%a7%e7%90%86%e7%9f%b3%e5%b3%a1%e8%b0%b7&form=hpcapt&mkt=zh-cn","title":"横跨峡谷裂缝的两座桥","quiz":"/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20220913_MarbleCanyon%22&FORM=HPQUIZ","wp":true,"hsh":"a078796a85005346f130ced55989a4d5","drk":1,"top":1,"bot":1,"hs":[]}],"tooltips":{"loading":"正在加载...","previous":"上一个图像","next":"下一个图像","walle":"此图片不能下载用作壁纸。","walls":"下载今日美图。仅限用作桌面壁纸。"}}`)
	var bingResp bingResponse
	if err := json.Unmarshal(resp, &bingResp); err != nil {
		return nil, err
	}
	if len(bingResp.Images) < 1 {
		return nil, fmt.Errorf("获取图片数量错误")
	}
	return bingResp.Images, nil
}

// func saveImageInfo(img *BingImgInfo, savePicDir string) {
// copyrightInfoPath := fmt.Sprintf("%s/introduction.txt", savePicDir)
// content := fmt.Sprintf("%s\t%s\t%s\n", img.Startdate, img.Get("id"), img.Copyright)
// writeFile(copyrightInfoPath, []byte(content))
// }
