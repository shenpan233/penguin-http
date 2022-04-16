/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 9:42
*/

package penguin_http

import (
	"bytes"
	"net/http"
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
		param string
		body  *bytes.Buffer
	)
	//param build
	for k, v := range this.param {
		param += "&" + k + "=" + v
	}
	if param != "" {
		url += "?" + param
	}
	url = this.build.baseurl + "/" + url
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
		if this.bodyReq != "" {
			body = bytes.NewBufferString(this.bodyReq)
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
