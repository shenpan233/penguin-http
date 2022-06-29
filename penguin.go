/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 9:42
*/

package penguin

import (
	"bytes"
	"net/http"
	"strings"
)

func (this *PenguinHttp) SetOnResponse(on OnResponse) *PenguinHttp {
	this.onResponse = on
	return this
}

//SetOnComplete
//  如需协程控制,建议在此处加入WaitGroup.Done
func (this *PenguinHttp) SetOnComplete(on OnComplete) *PenguinHttp {
	this.onComplete = on
	return this
}

func (this *PenguinHttp) makeReq(url string) (req *http.Request, err error) {
	var (
		param strings.Builder
		body  *bytes.Buffer
	)
	//param build
	if len(this.param) > 0 {
		param.WriteString(this.build.baseurl)
		param.WriteString("/?")
		for k, v := range this.param {
			param.WriteString("&" + k + "=" + v)
			//param += "&" + k + "=" + v
		}
		url = param.String()
	} else {
		url = this.build.baseurl + url
	}

	//header build

	this.checkAndSetHeader("Accept", "*/*")
	this.checkAndSetHeader("Cache-Control", "no-cache")
	this.checkAndSetHeader("User-Agent", UAMacPCSafari)
	this.checkAndSetHeader("Connection", "keep-alive")
	this.checkAndSetHeader("Referer", url)

	//post参数提交处理
	if this.bodyUpload != nil {
		body = bytes.NewBuffer(this.bodyUpload)
	} else {
		if this.bodyReq.Len() > 0 {
			sendBody := this.bodyReq.String()
			body = bytes.NewBufferString(sendBody[:len(sendBody)-1])
		}
	}

	//Builder release
	if body == nil {
		req, err = http.NewRequest(this.method, url, nil)
	} else {
		req, err = http.NewRequest(this.method, url, body)
	}

	if err != nil {
		return
	} else {
		req.Header = this.header
		return
	}
}
