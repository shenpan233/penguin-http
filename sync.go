/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 13:30
  @Notice:  同步请求
*/

package penguin

import (
	"io/ioutil"
	"net/http"
)

//Sync 同步请求
func (this *PenguinHttp) Sync(url string) (*PenguinResult, error) {
	req, err := this.makeReq(url)
	if err != nil {
		return nil, err
	}
	//Cookies build
	for _, cookie := range this.cookie {
		req.AddCookie(cookie)
	}
	do, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	//res := &PenguinResult{
	//	Status:           do.Status,
	//	StatusCode:       do.StatusCode,
	//	Proto:            do.Proto,
	//	ProtoMajor:       do.ProtoMajor,
	//	ProtoMinor:       do.ProtoMinor,
	//	Header:           do.Header,
	//	ContentLength:    do.ContentLength,
	//	TransferEncoding: do.TransferEncoding,
	//	Close:            do.Close,
	//	Uncompressed:     do.Uncompressed,
	//	Trailer:          do.Trailer,
	//	TLS:              do.TLS,
	//	Cookies:          make(map[string]string),
	//}
	res := this.pool.Get().(*PenguinResult)
	{
		res.Status = do.Status
		res.StatusCode = do.StatusCode
		res.Proto = do.Proto
		res.ProtoMajor = do.ProtoMajor
		res.ProtoMinor = do.ProtoMinor
		res.Header = do.Header
		res.TransferEncoding = do.TransferEncoding
		res.Close = do.Close
		res.Uncompressed = do.Uncompressed
		res.Trailer = do.Trailer
		res.TLS = do.TLS
	}
	defer func() {
		do.Body.Close()
		res.Clear()
		this.pool.Put(res)
	}()
	res.bin, _ = ioutil.ReadAll(do.Body)
	cookies := do.Cookies()
	if len(cookies) > 0 {
		for _, cookie := range cookies {
			res.Cookies[cookie.Name] = cookie.Value
		}
	}
	return res, nil
}
