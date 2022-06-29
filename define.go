package penguin

const (
	GET    = "GET"
	POST   = "POST"
	HEAD   = "HEAD"
	PUT    = "PUT"
	DELETE = "DELETE"
)

const (
	UAMacPCSafari = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15"
)

const (
	TypeJSON = ContentType("json")
	TypeForm = ContentType("form")
)

var Types = map[ContentType]string{
	TypeJSON: "application/json",
	TypeForm: "application/x-www-form-urlencoded",
	//TypeXML:        "application/xml",
	//TypeHTML:       "text/html",
	//TypeText:       "text/plain",
	//TypeMultipart:  "multipart/form-data",
}
