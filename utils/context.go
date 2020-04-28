package utils

import (
	"github.com/gin-gonic/gin"
)

type comRsp struct {
	ErrCode int    `json:"error_code"`
	ErrMsg  string `json:"error_msg"`
}

func Response(ctx *gin.Context, code int, msg ...string) {
	errMsg := ""
	if len(msg) > 0 {
		errMsg = msg[0]
	} else {
		errMsg = "request fail"
	}
	rsp := comRsp{
		ErrCode: code,
		ErrMsg:  errMsg,
	}
	ctx.JSON(ctx.Writer.Status(), rsp)
}
