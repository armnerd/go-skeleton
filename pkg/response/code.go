package response

// ParamsLost 参数缺失
const ParamsLost = 4001

// RecordNotExsit 记录不存在
const RecordNotExsit = 4002

// RequestFail curl请求失败
const RequestFail = 4004

// InternalError 内部错误
const InternalError = 5000

// 错误码映射
var codeMap = map[int]string{
	4001: "参数缺失",
	4002: "记录不存在",
	4004: "curl请求失败",
	5000: "内部错误",
}

// GetMessageByCode 获取错误信息
func GetMessageByCode(code int) string {
	if message, ok := codeMap[code]; ok {
		return message
	}
	return "未知错误码"
}
