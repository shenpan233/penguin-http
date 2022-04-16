/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 15:33
  @Notice:  notice
*/

package demo

import (
	penguin_http "penguin-http"
	"testing"
)

func TestDemo(t *testing.T) {
	demo()
}
func demo() {
	http := penguin_http.Builder().
		BaseUrl("http://github.com/shenpan233").
		Build()

	http.GET().
		SetCookie("cookies1", "2333").
		SetHeaderFromMap(map[string]string{
			"header": "hello world",
		}).
		Sync("/Penguin-http")

}
