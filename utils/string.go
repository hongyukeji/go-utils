package utils

import "os"

// TextTemplateFormat 文本模板格式化 将text字符串中的${example}变量, 通过params参数map数组替换
func TextTemplateFormat(text string, params map[string]string) string {
	return os.Expand(text, func(key string) string { return params[key] })
}
