/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 14:29
  @Notice:  Result封装
*/

package penguin

import (
	"crypto/tls"
	json "github.com/json-iterator/go"
	"net/http"
)

type (
	PenguinResult struct {
		bin              []byte //返回的原始数据
		Cookies          map[string]string
		Status           string // e.g. "200 OK"
		StatusCode       int    // e.g. 200
		Proto            string // e.g. "HTTP/1.0"
		ProtoMajor       int    // e.g. 1
		ProtoMinor       int    // e.g. 0
		Header           http.Header
		ContentLength    int64
		TransferEncoding []string
		Close            bool
		Uncompressed     bool
		Trailer          http.Header
		//Request *Request
		TLS *tls.ConnectionState
	}
)

func (this *PenguinResult) Json(val interface{}) {
	_ = json.Unmarshal(this.bin, val)
}

func (this *PenguinResult) String() string {
	if this == nil {
		return ""
	}
	return string(this.bin)
}

func (this *PenguinResult) Bytes() []byte {
	if this == nil {
		return nil
	}
	return this.bin
}

func (this *PenguinResult) Clear() {
	this.bin = nil
	this.Cookies = map[string]string{}
	this.Status = ""
	this.StatusCode = 0
	this.Proto = ""
	this.ProtoMajor = 0
	this.ProtoMinor = 0
	this.Header = http.Header{}
	this.ContentLength = 0
	this.TransferEncoding = make([]string, 0)
	this.Close = false
	this.Uncompressed = false
	this.Trailer = http.Header{}
	this.TLS = nil

}
