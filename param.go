/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 9:47
  @Notice:  表单式提交数据
*/

package penguin

import (
	"net/url"
	"strconv"
)

//AddParam 添加一组Key-val
//	val 仅支持基本数据类型
func (this *PenguinHttp) AddParam(key string, val interface{}) *PenguinHttp {
	var (
		v string
	)

	switch val.(type) {
	case int:
		v = strconv.Itoa(val.(int))
	case string:
		v = val.(string)
	case bool:
		b := val.(bool)
		if this.build.isParamBoolBit {
			if b {
				v = "1"
			} else {
				v = "0"
			}
		} else {
			if b {
				v = "true"
			} else {
				v = "false"
			}
		}
	}
	this.param[key] = url.QueryEscape(v)
	return this
}

//AddParamFromMap AddParamFrom的Map版
func (this *PenguinHttp) AddParamFromMap(m map[string]interface{}) *PenguinHttp {
	for k, v := range m {
		this.AddParam(k, v)
	}
	return this
}
