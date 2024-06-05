package api

var codeMap = map[int]string{
	SELECT_SUCCESS: "讀取成功",

	SELECT_FAILED: "讀取失敗",
}

func GetCodeMsg(code int) string {
	if _, ok := codeMap[code]; !ok {
		return ""
	}
	return codeMap[code]
}
