/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 13:31
  @Notice:  async异步请求
*/

package penguin_http

import (
	"io/ioutil"
	"net/http"
)

type (
	OnResponse func(result *PenguinResult) //请求成功后回调
	OnComplete func()
)

//Async
// onResponse 请求成功回调
func (this *PenguinHttp) Async(url string) {
	go func() {
		defer func() {
			if this.onComplete != nil {
				this.onComplete()
			}
		}()
		req, err := this.makeReq(url)
		if err != nil {
			return
		}
		//Cookies build
		for _, cookie := range this.cookie {
			req.AddCookie(cookie)
		}

		//发送请求
		do, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		//数据处理
		res := &PenguinResult{
			Status:           do.Status,
			StatusCode:       do.StatusCode,
			Proto:            do.Proto,
			ProtoMajor:       do.ProtoMajor,
			ProtoMinor:       do.ProtoMinor,
			Header:           do.Header,
			ContentLength:    do.ContentLength,
			TransferEncoding: do.TransferEncoding,
			Close:            do.Close,
			Uncompressed:     do.Uncompressed,
			Trailer:          do.Trailer,
			TLS:              do.TLS,
			Cookies:          make(map[string]string),
		}
		defer func() {
			do.Body.Close()
		}()
		res.bin, _ = ioutil.ReadAll(do.Body)
		cookies := do.Cookies()
		if len(cookies) > 0 {
			for _, cookie := range cookies {
				res.Cookies[cookie.Name] = cookie.Value
			}
		}

		//回调
		if this.onResponse != nil {
			this.onResponse(res)
		}
	}()
}
