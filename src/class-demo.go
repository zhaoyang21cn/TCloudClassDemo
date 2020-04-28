package src

import (
	"bytes"
	"classdemo/utils"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const source = "source"
const secretid="xxxxx"    		//填入你自己的secretid
const secretkey="xxxxx"   		//填入你自己的secretkey
const new_enter_id = 1075332  	//这里替换用创建学校接口返回的new_enter_id值
const admin = "xxxx"	  		//创建课堂需要传入user_id为老师或者管理员，这里填入自己注册的老师管理员id即可

const CreateNewEnterUrl = "https://iclass-open.api.qcloud.com/release/saas/v1/open/new_enter/create"					//创建学校URL
const UserRegisterUrl = "https://iclass-open.api.qcloud.com/release/saas/v1/open/user/register"							//用户注册URL
const UserOpenLoginUrl = "https://iclass-open.api.qcloud.com/release/saas/v1/open/login"								//用户登录(换取token，sig)URL
const CreateClassUrl = "https://iclass-open.api.qcloud.com/release/saas/v1/open/class/create"							//创建课堂URL
const CreateClassRoomCodeUrl = "https://iclass-open.api.qcloud.com/release/saas/v1/open/classroom_code/create"			//创建课堂URL

func calcAuthorization(source string, secretId string, secretKey string) (sign string, dateTime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	dateTime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	sign = fmt.Sprintf("x-date: %s\nsource: %s", dateTime, source)
	fmt.Println(sign)

	//hmac-sha1
	h := hmac.New(sha1.New, []byte(secretKey))
	io.WriteString(h, sign)
	sign = fmt.Sprintf("%x", h.Sum(nil))
	sign = string(h.Sum(nil))
	fmt.Println("sign:", fmt.Sprintf("%s", h.Sum(nil)))

	//base64
	sign = base64.StdEncoding.EncodeToString([]byte(sign))
	fmt.Println("sign:", sign)

	auth := fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date source\", signature=\"%s\"",
		secretId, sign)
	fmt.Println("auth:", auth)

	return auth, dateTime, nil
}
func Request(url string, bodyByte []byte, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return nil, err
	}

	sign, dateTime, err := calcAuthorization(source, secretid, secretkey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Source", source)
	req.Header.Set("X-Date", dateTime)
	req.Header.Set("Authorization", sign)

	client := &http.Client{
		Timeout: 7 * time.Second,//set timeout
	}
	rsp, err := client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
/***创建学习***/
func CreateNewEnterId(ctx *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		utils.Response(ctx, -1)
		return
	}
	reqUrl := fmt.Sprintf("%s?random=%d",CreateNewEnterUrl,utils.Random())
	rsp,err := Request(reqUrl,bodyBytes,http.MethodPost)
	if err != nil {
		utils.Response(ctx, -2, err.Error())
	}
	ctx.Data(ctx.Writer.Status(), "application/json; charset=utf-8", rsp)
}

/***注册用户***/
func UserRegister(ctx *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		utils.Response(ctx, -1)
		return
	}
	reqUrl := fmt.Sprintf("%s?new_enter_id=%d&random=%d",UserRegisterUrl,new_enter_id,utils.Random())
	rsp,err := Request(reqUrl,bodyBytes,http.MethodPost)
	if err != nil {
		utils.Response(ctx, -2, err.Error())
	}
	ctx.Data(ctx.Writer.Status(), "application/json; charset=utf-8", rsp)
}

/***用户登录***/
func UserOpenLogin(ctx *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		utils.Response(ctx, -1)
		return
	}
	reqUrl := fmt.Sprintf("%s?new_enter_id=%d&random=%d",UserOpenLoginUrl,new_enter_id,utils.Random())
	rsp,err := Request(reqUrl,bodyBytes,http.MethodPost)
	if err != nil {
		utils.Response(ctx, -2, err.Error())
	}
	ctx.Data(ctx.Writer.Status(), "application/json; charset=utf-8", rsp)
}

/***创建上课班号***/
func CreateClassRoomCode(ctx *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		utils.Response(ctx, -1)
		return
	}
	reqUrl := fmt.Sprintf("%s?new_enter_id=%d&random=%d&user_id=%s",CreateClassRoomCodeUrl,new_enter_id,utils.Random(),admin)
	rsp,err := Request(reqUrl,bodyBytes,http.MethodPost)
	if err != nil {
		utils.Response(ctx, -2, err.Error())
	}
	ctx.Data(ctx.Writer.Status(), "application/json; charset=utf-8", rsp)
}

/***创建课堂***/
func CreateClass(ctx *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		utils.Response(ctx, -1)
		return
	}
	reqUrl := fmt.Sprintf("%s?new_enter_id=%d&random=%d&user_id=%s",CreateClassUrl,new_enter_id,utils.Random(),admin)
	rsp,err := Request(reqUrl,bodyBytes,http.MethodPost)
	if err != nil {
		utils.Response(ctx, -2, err.Error())
	}
	ctx.Data(ctx.Writer.Status(), "application/json; charset=utf-8", rsp)
}

