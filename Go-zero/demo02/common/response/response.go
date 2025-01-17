package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response http返回
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := &Body{
			Code: 0,
			Msg:  "成功",
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	//错误返回
	errCode := uint32(10086)
	// 可以根据错误码，返回具体错误信息
	errMsg := "服务器错误"

	httpx.WriteJson(w, http.StatusBadRequest, &Body{
		Code: errCode,
		Msg:  errMsg,
		Data: nil,
	})

}
