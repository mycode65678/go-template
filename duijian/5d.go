package duijian

import (
	"strconv"
	"strings"
)

func (d *Dispose) d5() int {
	n1, _ := strconv.Atoi(d.arr[0])
	n2, _ := strconv.Atoi(d.arr[1])
	n3, _ := strconv.Atoi(d.arr[2])
	n4, _ := strconv.Atoi(d.arr[3])
	n5, _ := strconv.Atoi(d.arr[4])
	sum := n1 + n2 + n3 + n4 + n5
	//a.NewLog.WithFields(logrus.Fields{
	//	"game":         "xy28",
	//	"play_kind_id": d.bet.PlayKindId,
	//}).Debug("xy28", d.bet.Detail, "arr", d.arr)
	switch d.bet.PlayKindId {
	case 28:
		lane := make(map[string]int)
		lane["一球"] = 0
		lane["二球"] = 1
		lane["三球"] = 2
		lane["四球"] = 3
		lane["五球"] = 4
		arr := strings.Split(d.bet.Detail, "@")
		if arr[0] == "和值" {
			return d.sum(sum, d.bet.Detail)
		}
		num, _ := strconv.Atoi(d.arr[lane[arr[0]]])
		if num > 4 && arr[1] == "大" {
			return 1
		}
		if num < 5 && arr[1] == "小" {
			return 1
		}
		if num%2 == 1 && arr[1] == "单" {
			return 1
		}
		if num%2 == 0 && arr[1] == "双" {
			return 1
		}
	case 29:
		lane := make(map[string]int)
		lane["一球"] = 0
		lane["二球"] = 1
		lane["三球"] = 2
		lane["四球"] = 3
		lane["五球"] = 4
		arr := strings.Split(d.bet.Detail, "@")
		if d.arr[lane[arr[0]]] == arr[1] {
			return 1
		}
	case 30:
		if n1 > n5 && d.bet.Detail == "龙" {
			return 1
		}
		if n1 < n5 && d.bet.Detail == "虎" {
			return 1
		}
	case 31:
		if n1 == n5 {
			return 1
		}
	case 32:
		if n3 == n4 && n4 == n5 {
			return 1
		}
	//	顺子
	case 33:
		n := []int{n3, n4, n5}
		return d.Straight(n)
	// 半顺
	case 34:
		n := []int{n3, n4, n5}
		return d.SemiShun(n)
	//	 对子
	case 35:
		n := []int{n3, n4, n5}
		obj := make(map[int]int)
		for _, v := range n {
			if _, ok := obj[v]; !ok {
				obj[v] = 0
			} else {
				obj[v]++
			}
		}
		for _, v := range obj {
			if v == 2 {
				return 1
			}
		}
	// 杂六
	case 36:
		n := []int{n1, n2, n3}
		return d.Miscellaneous(n)
	// 牛牛
	case 37:
	case 38:
	case 39:
	}
	return 0
}
