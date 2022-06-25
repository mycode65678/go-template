package a

import (
	"encoding/json"
	"fmt"
	validator "gopkg.in/go-playground/validator.v9"
)

var responseLists = map[int]string{
	0: "错误",
	1: "成功",
}

type ResponseStruct struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// WithMsg 自定义响应信息
func (res *ResponseStruct) WithMsg(message string) ResponseStruct {
	return ResponseStruct{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// WithData 追加响应数据
func (res *ResponseStruct) WithData(data interface{}) ResponseStruct {
	return ResponseStruct{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *ResponseStruct) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// Response 构造函数
func Response(code int) *ResponseStruct {
	msg := responseLists[code]
	return &ResponseStruct{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// ResponseMsg 构造函数
func ResponseMsg(code int, msg string) *ResponseStruct {
	return &ResponseStruct{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// ResponseData 构造函数
func ResponseData(code int, Data interface{}) *ResponseStruct {
	return &ResponseStruct{
		Code: code,
		Msg:  "",
		Data: Data,
	}
}

// ResponseError 返回错误消息
func ResponseError(err error) *ResponseStruct {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := T(fmt.Sprintf("Field.%s", e.Field))
			tag := T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return ResponseData(1, fmt.Sprintf("%s%s", field, tag))
			//return serializer.ParamErr(
			//	fmt.Sprintf("%s%s", field, tag),
			//	err,
			//)
		}
	}
	//if _, ok := err.(*json.UnmarshalTypeError); ok {
	//	return serializer.ParamErr("JSON类型不匹配", err)
	//}
	return ResponseMsg(99, "参数错误")
}
