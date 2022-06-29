package penguin

func (this *PenguinHttp) ResetSend() *PenguinHttp {
	this.bodyReq.Reset()
	return this
}
func (this *PenguinHttp) ResetCookies() *PenguinHttp {
	this.cookie = nil
	return this
}
