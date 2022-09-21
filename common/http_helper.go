package common

import (
	"fmt"
	"io/ioutil"

	// "log"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// SendRequest 带超时的请求
func SendRequest(method, apiURL string, bodyData url.Values, timeout int) ([]byte, error) {
	start := time.Now()

	resp, err := sendRequestDo(method, apiURL, bodyData, timeout)

	logMap := map[string]interface{}{}
	logMap["url"] = apiURL
	logMap["body"] = bodyData.Encode()

	if err != nil || resp == nil {
		usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		fmt.Println(logMap, msg, resp, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		fmt.Println(logMap, msg, resp.StatusCode)
		return nil, fmt.Errorf("statusCode:%d", resp.StatusCode)
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
		msg := fmt.Sprintf("failed\t%.3fms", usageTime)
		fmt.Println(logMap, msg, err.Error())
		return nil, err
	}

	usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	msg := fmt.Sprintf("success\t%.3fms", usageTime)
	fmt.Println(logMap, msg)
	return respBody, nil
}

// SendRequestDo 带超时的请求
func sendRequestDo(method, apiURL string, bodyData url.Values, timeout int) (resp *http.Response, err error) {
	request, _ := http.NewRequest(method, apiURL, strings.NewReader(bodyData.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 128
	// t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 256

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Millisecond,
		Transport: t,
	}

	return client.Do(request)
}

// DownloadImg 下载图片
func DownloadImg(imgURL, imageName, downDir string) error {
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
	resp, err := sendRequestDo(http.MethodGet, imgURL, nil, 5000)
	usageTime := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	fmt.Printf("download usage time:\t%.3fms\n", usageTime)

	_, err = io.Copy(imageFd, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
