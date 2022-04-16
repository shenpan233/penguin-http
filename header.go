/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 11:18
  @Notice:  Header参数
*/

package penguin_http

//SetUserAgent 设置UA
func (this *PenguinHttp) SetUserAgent(ua string) *PenguinHttp {
	this.header.Add("User-Agent", ua)
	return this
}

//SetHeader 添加Header
func (this *PenguinHttp) SetHeader(key, val string) *PenguinHttp {
	this.header.Add(key, val)
	return this
}

//SetHeaderFromMap SetHeaderFromMap
func (this *PenguinHttp) SetHeaderFromMap(m map[string]string) *PenguinHttp {
	for key, val := range m {
		this.header.Add(key, val)
	}
	return this
}

func (this *PenguinHttp) checkAndSetHeader(name string, val string) {
	if this.header.Get(name) == "" {
		this.header.Set(name, val)
	}
}
