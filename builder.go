package penguin

import (
	"net/http"
	"strings"
	"sync"
)

func Builder() *PenguinBuilder {
	return new(PenguinBuilder)
}

func (build *PenguinBuilder) BaseUrl(url string) *PenguinBuilder {
	build.baseurl = url
	return build
}

//IsParamBoolBit
// 取反 -> 是否用0或1代替false或true
func (build *PenguinBuilder) IsParamBoolBit() *PenguinBuilder {
	build.isParamBoolBit = !build.isParamBoolBit
	return build
}

func (build *PenguinBuilder) Build() *PenguinHttp {
	return &PenguinHttp{
		build:      build,
		method:     "",
		param:      make(map[string]string),
		cookie:     nil,
		header:     http.Header{},
		bodyUpload: nil,
		bodyReq:    strings.Builder{},
		pool: &sync.Pool{New: func() interface{} {
			p := new(PenguinResult)
			p.Clear()
			return p
		}},
	}
}
