package req

/*
req extension by Leafney
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// *****Request********************

// Custom settings parameters
type XSets map[string]interface{}

func XReq(url string, xSets map[string]interface{}, v ...interface{}) (r *Resp, err error) {

	// default Request Types
	method_Type := "GET"

	header := make(http.Header)

	// Default configuration

	// Set default Content-Type
	// header.Set("Content-Type", "application/json;charset=utf-8")

	// Set default User-Agent
	// header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	// Set custom configuration parameters
	for key, value := range xSets {

		ky := strings.ToLower(key)
		switch ky {
		case "content-type":
			header.Set("Content-Type", value.(string))
		case "user-agent":
			header.Set("User-Agent", value.(string))
		case "time-out":
			SetTimeout(time.Duration(value.(int)) * time.Second)
		case "proxy":
			SetProxyUrl(value.(string))
		case "referer":
			header.Set("Referer", value.(string))
		case "debug":
			Debug = value.(bool)
		case "method":
			// Request Types
			method_Type = strings.ToUpper(value.(string))
		default:
			fmt.Printf("the %s-%s not default settings", key, value)
			// Other parameters are added to the request header by default.
			header.Set(key, value.(string))
		}
	}

	//Add custom parameters to default parameters
	v = append(v, header)

	// 根据请求类型，设置请求参数或请求体，发起请求
	switch method_Type {
	case "GET":
		return Get(url, v...)
	case "POST":
		return Post(url, v...)
	case "PUT":
		return Put(url, v...)
	case "DELETE":
		return Delete(url, v...)
	default:
		return
	}

}

// *****Response********************

// ToJSONfromJSONP convert jsonp response body string to json str and then to struct or map
func (r *Resp) ToJSONfromJSONP(v interface{}) error {
	data, err := r.ToString()
	if err != nil {
		return err
	}

	// Regular matching like `jQuery7955233 ();` jsonp format
	pat := `^[^(]*?\((.*)\)[^)]*$`
	reg := regexp.MustCompile(pat)
	resp := reg.FindStringSubmatch(data)
	// fmt.Println(resp) 解析不到内容时为 []
	if len(resp) == 2 {
		data = resp[1]
		// Try JSON format parsing
		return json.Unmarshal([]byte(data), v)
	} else {
		return fmt.Errorf("jsonp data parsing failure for: %s", data)
	}
}

// JSONfromJSONP convert jsonp response body string to json str
func (r *Resp) ToJSONStrfromJSONP() (string, error) {
	data, err := r.ToString()
	if err != nil {
		return "", err
	}

	// Regular matching like `jQuery7955233 ();` jsonp format
	pat := `^[^(]*?\((.*)\)[^)]*$`
	reg := regexp.MustCompile(pat)
	resp := reg.FindStringSubmatch(data)
	if len(resp) == 2 {
		data = resp[1]
		// Try JSON format parsing
		return data, nil
	} else {
		return "", fmt.Errorf("jsonp data parsing failure for: %s", data)
	}

}

// *************************
