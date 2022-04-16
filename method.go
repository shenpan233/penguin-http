/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 10:19
  @Notice:  Method
*/

package penguin_http

func (p *PenguinHttp) GET() *PenguinHttp {
	p.method = GET
	return p
}

func (p *PenguinHttp) POST() *PenguinHttp {
	p.method = POST
	return p
}
