package request

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/small-ek/ginp/conv"
	"io/ioutil"
	"log"
	"net/http"
)

//GetBody Get the requested data
func GetBody(this *gin.Context) map[string]interface{} {
	var request map[string]interface{}
	var body []byte
	if this.Request.Body != nil {
		body, _ = ioutil.ReadAll(this.Request.Body)
		this.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	var err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err.Error())
	}
	return request
}

//Input...
func Input(name string, this *gin.Context) interface{} {
	var request = GetBody(this)
	return request[name]
}

//GetString...
func GetString(name string, this *gin.Context) string {
	var request = GetBody(this)
	return String(request[name])
}

//GetBool...
func GetBool(name string, this *gin.Context) bool {
	var request = GetBody(this)
	return Bool(request[name])
}

//GetFloat32...
func GetFloat32(name string, this *gin.Context) float32 {
	var request = GetBody(this)
	return Float32(request[name])
}

//GetFloat64...
func GetFloat64(name string, this *gin.Context) float64 {
	var request = GetBody(this)
	return Float64(request[name])
}

//GetInt...
func GetInt(name string, this *gin.Context) int {
	var request = GetBody(this)
	return Int(request[name])
}

//Cors Cross-domain request
func Cors(this *gin.Context) {
	this.Header("Access-Control-Allow-Origin", "*")                                      //允许访问所有域
	this.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE,OPTIONS") //允许请求类型
	this.Header("Access-Control-Allow-Credentials", "true")                              //服务器是否接受浏览器发送的Cookie
	this.Header("Connection", "keep-alive")                                              //可以使一次TCP连接为同意用户的多次请求服务,提高了响应速度。
	this.Header("Access-Control-Max-Age", "3600")                                        //多少秒以后再次OPTIONS.默认60分钟
	this.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, If-Match, If-Modified-Since, If-None-Match, If-Unmodified-Since, X-Requested-With")
	this.Header("Strict-Transport-Security", " max-age=63072000; includeSubdomains; preload")
	//放行所有OPTIONS方法
	if this.Request.Method == "OPTIONS" {
		this.AbortWithStatus(http.StatusNoContent)
	}
	//处理请求
	this.Next()
}
