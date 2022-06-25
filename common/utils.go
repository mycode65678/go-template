package common

import (
	"os"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检测两面类型的下注是否合法
// location 下注位置
// content 下注内容
// detail 下注详情进行对比 冠军@大
func CheckBet(location []string, content []string, detail string) bool {
	t := false
	for _, v := range location {
		for _, v1 := range content {
			if detail == v+"@"+v1 {
				t = true
			}
		}
	}
	return t
}

func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}
