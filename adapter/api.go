package adapter

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var gCache = cache.New(60*time.Minute, 1*time.Minute)

// APIs ..
type APIs map[string]API

// Get API
func (apis APIs) Get(name string) (result API) {
	if adapter, ok := apis[name]; ok {
		result = adapter
		result.Name = name
	} else {
		panic("Không tìm thấy config API " + name)
	}
	return
}


// API ..
type API struct {
	Name   string
	URL    string `mapstructure:"url"`
	Method string `mapstructure:"method"`
	Type   string `mapstructure:"content-type"`
	Auth   struct {
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"auth"`
	Headers []header `mapstructure:"header"`
	Params  []param  `mapstructure:"params"`
	Cache   int64    `mapstructure:"cache"`
}

type header struct {
	Key   string `mapstructure:"key"`
	Value string `mapstructure:"value"`
}

type param struct {
	Key   string      `mapstructure:"key"`
	Value interface{} `mapstructure:"value"`
}

// SetParams ..
func (api *API) SetParams(params map[string]interface{}) {
	for k, v := range params {
		api.Params = append(api.Params, param{Key: k, Value: v})
	}
}

// SetRequestParams ..
func (api *API) SetRequestParams(params interface{}) {
	apiParams := structs.Map(params)
	api.SetParams(apiParams)
}

// SetParams ..
func (api *API) getParams() map[string]interface{} {
	params := make(map[string]interface{})
	for _, v := range api.Params {
		params[v.Key] = v.Value
	}
	return params
}

func hash(data API) string {
	arrBytes := []byte{}

	jsonBytes, _ := json.Marshal(data)
	arrBytes = append(arrBytes, jsonBytes...)

	return fmt.Sprintf("%x", md5.Sum(arrBytes))
}

// MakeRequest ..
func (api API) MakeRequest(requestID string) (body []byte, err error) {

	start := time.Now()

	key := hash(api)

	cacheData, ok := gCache.Get(key)
	if ok {
		body = cacheData.([]byte)
		return
	}

	req := new(http.Request)

	if api.Method == "" {
		api.Method = "GET"
	}

	switch api.Method {
	case http.MethodGet:
		req, err = http.NewRequest(api.Method, api.URL, nil)
		if err != nil {
			return
		}
		q := req.URL.Query()
		for _, param := range api.Params {
			q.Add(param.Key, fmt.Sprintf("%v", param.Value))
		}
		req.URL.RawQuery = q.Encode()
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		bodyRequest := &bytes.Buffer{}
		json.NewEncoder(bodyRequest).Encode(api.getParams())
		req, err = http.NewRequest(api.Method, api.URL, bodyRequest)
		if err != nil {
			return
		}
	}
	return api.doRequest(start, requestID, key, req)
}


func (api API) doRequest(start time.Time, requestID, key string, req *http.Request)(body []byte, err error) {

	if api.Auth.UserName != "" {
		req.SetBasicAuth(api.Auth.UserName, api.Auth.Password)
	}

	for _, header := range api.Headers {
		if header.Key != "" {
			req.Header.Set(header.Key, header.Value)
		}
	}

	//reset requestID
	req.Header.Set(echo.HeaderXRequestID, requestID)

	response, errDoRequest := http.DefaultClient.Do(req)
	if errDoRequest != nil {
		err = errDoRequest
		return
	}

	body, err = ioutil.ReadAll(response.Body)

	response.Body.Close()

	stop := time.Now()

	if api.Cache > 0 {
		if err == nil && response.StatusCode == http.StatusOK {
			gCache.Set(key, body, time.Second*time.Duration(api.Cache))
		} else {
			log.Printf("Can't cache request: %v\n", api.URL)
		}
	}

	log.Printf("MakeRequest: %v - %v - %v - %v\n", api.URL, api.Cache, stop.Sub(start).String(), response.Status)

	return
}


// MakeRequest ..
func (api API) MakeRequestWithAttachments(requestID string) (body []byte, err error) {
	start := time.Now()

	key := hash(api)

	cacheData, ok := gCache.Get(key)
	if ok {
		body = cacheData.([]byte)
		return
	}
	req := new(http.Request)
	bodyRequest := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyRequest)
	for key, val := range api.getParams() {
		valueType := reflect.TypeOf(val)
		if valueType.String() == "[]*multipart.FileHeader" {
			values := val.([]*multipart.FileHeader)
			for _, file := range values {
				rawFile, err := file.Open()
				if err != nil {
					return nil, err
				}
				defer rawFile.Close()

				fileWriter, err := bodyWriter.CreateFormFile(key, file.Filename)
				if err != nil {
					return nil, err
				}
				if _, err = io.Copy(fileWriter, rawFile); err != nil {
					return nil, err
				}
			}
		} else {
			if valueType.String() == "string" {
				newVal := val.(string)
				err = bodyWriter.WriteField(key, newVal)
				if err != nil {
					return
				}
			} else {
				var wt io.Writer
				value, _ := json.Marshal(val)
				wt, err = bodyWriter.CreateFormField(key)
				if err != nil {
					return
				}
				if _, err = io.Copy(wt, bytes.NewReader(value)); err != nil {
					return
				}
			}
		}
	}
	bodyWriter.Close()
	req, err = http.NewRequest(api.Method, api.URL, bodyRequest)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	return api.doRequest(start, requestID, key, req)
}


