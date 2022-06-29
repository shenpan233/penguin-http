/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 15:33
  @Notice:  notice
*/

package demo

import (
	penguin "github.com/shenpan233/PenguinHttp"
	"testing"
)

func TestDemo(t *testing.T) {
	demo()
}

var http = *penguin.Builder().
	BaseUrl("http://localhost").
	Build().
	SetHeaderFromMap(map[string]string{
		"header": "hello world",
	})

func demo() {
	penguinHttp := http.GET()
	penguinHttp.Sync("")
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo()
	}
}
