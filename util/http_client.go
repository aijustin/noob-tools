/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2021/2/1 1:14 下午
 */
package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

const (
	MethodGet       = "GET"
	MethodPost      = "POST"
	HeadJson        = "application/json"
	HeadPostDefault = "application/x-www-form-urlencoded"
	HeadContentType = "Content-Type"
)

type HttpClient struct {
	method        string
	timeOut       time.Duration
	debug         bool
	header        map[string]string
	httpTransport *http.Transport
	client        *http.Client
	url           string
	requestParams interface{}
	responseData  *io.ReadCloser
	err           error
	files         map[string]*HttpClientUploadFile
	extFuc        func(r *http.Request)
	jsonParams    bool
}

type HttpClientUploadFile struct {
	FileName string
	File     io.Reader
}

func NewHttpClient() *HttpClient {
	hc := new(HttpClient)
	hc.requestParams = make(map[string]string)
	hc.header = make(map[string]string)
	hc.files = make(map[string]*HttpClientUploadFile)
	hc.httpTransport = &http.Transport{}
	return hc
}

func (hc *HttpClient) Debug() *HttpClient {
	hc.debug = true
	return hc
}

func (hc *HttpClient) GetLastErr() error {
	return hc.err
}

func (hc *HttpClient) GET(url string) *HttpClient {
	hc.method = MethodGet
	hc.url = url
	response, err := hc.request()
	if err != nil {
		hc.err = err
		return hc
	}
	hc.responseData = &response.Body

	return hc
}

func (hc *HttpClient) Params(params interface{}) *HttpClient {
	hc.requestParams = params
	return hc
}

func (hc *HttpClient) JsonParams(params interface{}) *HttpClient {
	hc.requestParams = params
	hc.jsonParams = true
	hc.SetHeader(HeadContentType, HeadJson)
	return hc
}

func (hc *HttpClient) POST(url string) *HttpClient {
	hc.method = MethodPost
	hc.url = url
	response, err := hc.request()
	if err != nil {
		hc.err = err
		return hc
	}
	hc.responseData = &response.Body
	return hc
}

func (hc *HttpClient) ToString() string {
	if hc.responseData == nil {
		return ""
	}
	resText, err := ioutil.ReadAll(*hc.responseData)
	if err != nil {
		hc.err = err
		return ""
	}
	return string(resText)
}

func (hc *HttpClient) ToBytes() []byte {
	resText, err := ioutil.ReadAll(*hc.responseData)
	if err != nil {
		hc.err = err
		return resText
	}
	return resText
}

func (hc *HttpClient) SetRequestOption(req func(r *http.Request)) *HttpClient {
	hc.extFuc = req
	return hc
}

func (hc *HttpClient) SetFileByFileName(filePath, uploadName, uploadFileName string) *HttpClient {
	file, err := os.Open(filePath)
	if err != nil {
		hc.err = err
		return hc
	}
	defer file.Close()

	fileB, err := ioutil.ReadAll(file)
	if err != nil {
		hc.err = err
		return hc
	}
	fileBytes := bytes.NewReader(fileB)
	hc.files[uploadName] = &HttpClientUploadFile{
		FileName: uploadFileName,
		File:     fileBytes,
	}
	return hc
}

func (hc *HttpClient) SetFileByIo(buf io.Reader, uploadName, uploadFileName string) *HttpClient {
	hc.files[uploadName] = &HttpClientUploadFile{
		FileName: uploadFileName,
		File:     buf,
	}
	return hc
}

func (hc *HttpClient) SetProxy(protocol string, addr string) error {
	if protocol == "tcp" {
		dial, err := proxy.SOCKS5(protocol, addr, nil, proxy.Direct)
		if err != nil {
			return err
		}
		hc.httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			return dial.Dial(network, addr)
		}
	} else if protocol == "http" {
		hc.httpTransport.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(addr)
		}
	}
	return nil
}

func (hc *HttpClient) GetClient() *http.Client {
	client := &http.Client{
		Transport: hc.httpTransport,
		Timeout:   hc.timeOut,
	}
	hc.client = client
	return hc.client
}

func (hc *HttpClient) SetHeader(k string, v string) *HttpClient {
	hc.header[k] = v
	return hc
}

func (hc *HttpClient) getRequestHeader(key string) string {
	if v, ok := hc.header[key]; !ok {
		return ""
	} else {
		return v
	}
}

func (hc *HttpClient) request() (*http.Response, error) {
	client := hc.GetClient()
	var buf io.Reader

	if hc.requestParams != nil && hc.method == MethodGet {
		uri, _ := url.Parse(hc.url)
		urlValuesStr := url.Values{}
		for k, v := range hc.getMapParams() {
			urlValuesStr.Add(k, v)
		}
		if uri.RawQuery == "" {
			hc.url += "?" + urlValuesStr.Encode()
		} else {
			hc.url += "&" + urlValuesStr.Encode()
		}
	}

	if hc.requestParams != nil && hc.method == MethodPost && hc.jsonParams == true {
		params, err := json.Marshal(hc.requestParams)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewReader(params)
	}

	if hc.requestParams != nil && hc.method == MethodPost && hc.jsonParams == false {
		if hc.method == MethodPost && len(hc.files) > 0 {
			bodyBuf := &bytes.Buffer{}
			bodyWriter := multipart.NewWriter(bodyBuf)
			for k, v := range hc.getMapParams() {
				_ = bodyWriter.WriteField(k, v)
			}
			for k, v := range hc.files {
				fileWriter, err := bodyWriter.CreateFormFile(k, v.FileName)
				if err != nil {
					hc.err = err
					return nil, err
				}
				_, err = io.Copy(fileWriter, v.File)
				if err != nil {
					hc.err = err
					return nil, err
				}
			}
			hc.SetHeader(HeadContentType, bodyWriter.FormDataContentType())
			bodyWriter.Close()
			buf = bodyBuf
		} else {
			form := url.Values{}
			for k, v := range hc.getMapParams() {
				form.Add(k, v)
			}
			buf = strings.NewReader(form.Encode())
			hc.SetHeader(HeadContentType, HeadPostDefault)
		}

	}

	req, err := http.NewRequest(hc.method, hc.url, buf)
	if err != nil {
		return nil, errors.New("NewHttpClient Request Error")
	}

	for k, v := range hc.header {
		req.Header.Set(k, v)
	}
	if hc.extFuc != nil {
		hc.extFuc(req)
	}
	return client.Do(req)
}

func (hc *HttpClient) structToMap(structData interface{}) (mapData map[string]interface{}, err error) {
	mapData = nil
	data, marErr := json.Marshal(structData)
	if marErr != nil {
		err = marErr
		return
	}

	unMarErr := json.Unmarshal(data, &mapData)
	if unMarErr != nil {
		err = unMarErr
		return
	}
	return
}

func (hc *HttpClient) getMapParams() map[string]string {
	mapParams := make(map[string]string)
	requestParamsMap, err := hc.structToMap(hc.requestParams)
	if err != nil {
		return mapParams
	}
	typeof := reflect.TypeOf(requestParamsMap)

	if typeof.Kind() != reflect.Map {
		return mapParams
	}
	if len(requestParamsMap) < 1 {
		return mapParams
	}

	for k, v := range requestParamsMap {
		mapParams[k] = fmt.Sprintf("%v", v)
	}
	fmt.Println(mapParams)
	return mapParams
}
