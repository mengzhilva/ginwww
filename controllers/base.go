package controllers

type Base struct {
}

// 错误返回,不管正确错误，状态都返回200
func (c *Base) error(msg any, data any) map[string]any {
	return map[string]any{
		"msg":  msg,
		"data": data,
		"code": 1,
	}
}

// 成功返回
func (c *Base) success(msg any, data any) map[string]any {
	return map[string]any{
		"msg":  msg,
		"data": data,
		"code": 0,
	}
}
