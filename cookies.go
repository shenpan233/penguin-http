/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 12:42
  @Notice:  cookies
*/

package penguin_http

import (
	"net/http"
)

func (this *PenguinHttp) SetCookie(key, val string) *PenguinHttp {
	this.cookie = append(this.cookie, &http.Cookie{
		Name:  key,
		Value: val,
	})
	return this
}

func (this *PenguinHttp) SetCookieFromMap(m map[string]string) *PenguinHttp {
	for key, val := range m {
		this.cookie = append(this.cookie, &http.Cookie{
			Name:  key,
			Value: val,
		})
	}
	return this
}
