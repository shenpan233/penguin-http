/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 10:19
  @Notice:  Method
*/

package penguin

import (
	"reflect"
)

func (p *PenguinHttp) GET() (http *PenguinHttp) {
	p.method = GET
	return p
}

func (p *PenguinHttp) POST() *PenguinHttp {
	p.method = POST
	return p
}
func copyPoint(m *PenguinHttp) *PenguinHttp {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface().(*PenguinHttp)
}
