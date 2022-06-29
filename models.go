package penguin

import (
	"net/http"
	"strings"
	"sync"
)

type (
	PenguinBuilder struct {
		baseurl        string
		isParamBoolBit bool //是否用0或1代替false或true
		ContentType    ContentType
	}

	PenguinHttp struct {
		build      *PenguinBuilder
		method     string
		param      map[string]string
		cookie     []*http.Cookie
		header     http.Header
		bodyUpload []byte
		bodyReq    strings.Builder //POST的请求参
		onResponse OnResponse
		onComplete OnComplete
		pool       *sync.Pool
	}
)

type ContentType string
