// Package util 是我乱写的一个例子包
package util

import "os"

// GetEnvVar 获取环境变量
// 当 varName 表示的环境变量存在且不为空字符串时返回环境变量 varName 的值
// 当 varName 表示的环境变量不存在或为空时返回 defaultVal 的值
func GetEnvVar(varName string, defaultVal string) string {
	ret, exist := os.LookupEnv(varName)
	if !exist || ret == "" {
		return defaultVal
	}
	return ret
}

// GetEnvVarAlternative 获取环境变量
// 获取优先级为 varName[0] 到 varName[len(varName) - 1]，最后添加 defaultVal
func GetEnvVarAlternative(varName []string, defaultVal func() string) string {
	for _, it := range varName {
		ret, exist := os.LookupEnv(it)
		if !exist || ret == "" {
			continue
		}
		return ret
	}
	return defaultVal()
}
