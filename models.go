package penguin_http

import "net/http"

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
		bodyReq    string
		onResponse OnResponse
		onComplete OnComplete
	}
)

type ContentType string
