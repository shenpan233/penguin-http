/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 9:53
  @Notice:  请求数据
*/

package penguin_http

import (
	json "github.com/json-iterator/go"
)

//SetUpload 上传[]byte非图片
func (this *PenguinHttp) SetUpload(bin []byte) *PenguinHttp {
	this.bodyUpload = bin
	return this
}

//SendString 增加一条请求数据 只是拼接不会进行其他处理
func (this *PenguinHttp) SendString(data string) *PenguinHttp {
	this.bodyReq += data
	return this

}

func (this *PenguinHttp) SendStruct(data interface{}) *PenguinHttp {
	marshal, _ := json.Marshal(data)
	this.bodyReq = string(marshal)
	return this
}
