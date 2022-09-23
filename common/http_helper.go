package common

import (
	"fmt"
	"io/ioutil"

	"crypto/tls"
	"crypto/x509"
	// "golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var requestStartTime time.Time
var defaultHeaders map[string]string

func init() {
	defaultHeaders = map[string]string{
		"Content-Type":              "application/x-www-form-urlencoded",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"accept-encoding":           "gzip, deflate, br",
		"accept-language":           "zh-CN,zh;q=0.9",
		"cache-control":             "no-cache",
		"pragma":                    "no-cache",
		"sec-ch-ua":                 `"Google Chrome";v="105":"Not)A;Brand";v="8":"Chromium";v="105"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Windows"`,
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "none",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	}
}

// RequestGetByCert 自定义ssl证书
func RequestGetByCert(apiURL string, caCert []byte, timeout int, headers map[string]string) ([]byte, error) {

	requestStartTime = time.Now()
	// Create a pool with the server certificate since it is not signed
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	resp, err := sendRequestDo(http.MethodGet, apiURL, nil, 5000, nil, transport)
	return respProcess(apiURL, nil, resp, err)
}

// SendRequest 带超时的请求
func SendRequest(method, apiURL string, bodyData url.Values, timeout int, headers map[string]string) ([]byte, error) {
	requestStartTime = time.Now()
	resp, err := sendRequestDo(method, apiURL, bodyData, timeout, headers, nil)
	return respProcess(apiURL, bodyData, resp, err)
}

// DownloadImg 下载图片
func DownloadImg(imgURL, imageName, downDir string, headers map[string]string) error {
	_, err := os.Stat(downDir)
	if os.IsNotExist(err) {
		os.MkdirAll(downDir, 0666)
	}

	imgFilePath := fmt.Sprintf("%s/%s", downDir, imageName)
	_, err = os.Stat(imgFilePath)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file exists: %s", imgFilePath)
	}

	imageFd, err := os.Create(imgFilePath)
	if err != nil {
		return err
	}
	defer imageFd.Close()

	start := time.Now()
	resp, err := sendRequestDo(http.MethodGet, imgURL, nil, 5000, headers, nil)
	usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	log.Printf("download usage time:\t%s\t%.3fms\n", imgURL, usageTime)

	_, err = io.Copy(imageFd, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func respProcess(apiURL string, bodyData url.Values, resp *http.Response, err error) ([]byte, error) {
	logMap := map[string]interface{}{}
	logMap["url"] = apiURL
	logMap["body"] = bodyData.Encode()

	if err != nil || resp == nil {
		usageTime := float64(time.Now().Sub(requestStartTime).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		log.Println(logMap, msg, resp, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		usageTime := float64(time.Now().Sub(requestStartTime).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		log.Println(logMap, msg, resp.StatusCode, resp)
		return nil, fmt.Errorf("statusCode:%d", resp.StatusCode)
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		usageTime := float64(time.Now().Sub(requestStartTime).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		log.Println(logMap, msg, err.Error())
		return nil, err
	}

	usageTime := float64(time.Now().Sub(requestStartTime).Nanoseconds()) / 1e6
	msg := fmt.Sprintf("success\t%.3fms", usageTime)
	log.Println(logMap, msg)
	return respBody, nil
}

// SendRequestDo 请求装
func sendRequestDo(method, apiURL string, bodyData url.Values, timeout int, headers map[string]string, trans *http.Transport) (resp *http.Response, err error) {
	request, _ := http.NewRequest(method, apiURL, strings.NewReader(bodyData.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("pragma", "no-cache")
	request.Header.Add("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-ch-ua-platform", `"Windows"`)
	request.Header.Add("sec-fetch-dest", "document")
	request.Header.Add("sec-fetch-mode", "navigate")
	request.Header.Add("sec-fetch-site", "none")
	request.Header.Add("sec-fetch-user", "?1")
	request.Header.Add("upgrade-insecure-requests", "1")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	for k, v := range headers {
		request.Header.Set(k, v)
	}
	if trans == nil {
		trans = http.DefaultTransport.(*http.Transport).Clone()
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Millisecond,
		Transport: trans,
	}
	return client.Do(request)
}
