/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 9:53
  @Notice:  请求数据
*/

package penguin

import (
	"fmt"
	json "github.com/json-iterator/go"
)

//SetUpload 上传[]byte非图片
func (this *PenguinHttp) SetUpload(bin []byte) *PenguinHttp {
	this.bodyUpload = bin
	return this
}

//SendString 增加一条请求数据 只是拼接不会进行其他处理
func (this *PenguinHttp) SendString(key, val string) *PenguinHttp {
	this.bodyReq.WriteString(key + "=" + val + "&")
	return this

}

//SendMap 通过map方式添加数据,性能低不建议使用
func (this *PenguinHttp) SendMap(data map[string]interface{}) *PenguinHttp {
	if len(data) > 0 {
		for key, val := range data {
			this.bodyReq.WriteString(fmt.Sprintf("%s=%v", key, val) + "&")
		}
	}
	return this

}

//SendStructJson 自动将struct转化为json数据
func (this *PenguinHttp) SendStructJson(data interface{}) *PenguinHttp {
	marshal, _ := json.Marshal(data)
	this.bodyReq.Write(marshal)
	return this
}
