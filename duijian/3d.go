package duijian

import (
	"strconv"
	"strings"
)

func (d *Dispose) sum(sum int, detail string) int {
	if sum%2 == 0 && detail == "双" {
		return 1
	}
	if sum%2 == 1 && detail == "单" {
		return 1
	}
	if sum > 13 && detail == "大" {
		return 1
	}
	if sum < 14 && detail == "小" {
		return 1
	}
	return 0
}

func (d *Dispose) d3() int {
	n1, _ := strconv.Atoi(d.arr[0])
	n2, _ := strconv.Atoi(d.arr[1])
	n3, _ := strconv.Atoi(d.arr[2])
	sum := n1 + n2 + n3
	//a.NewLog.WithFields(logrus.Fields{
	//	"game":         "xy28",
	//	"play_kind_id": d.bet.PlayKindId,
	//}).Debug("xy28", d.bet.Detail, "arr", d.arr)
	switch d.bet.PlayKindId {
	case 19:
		lane := make(map[string]int)
		lane["一球"] = 0
		lane["二球"] = 1
		lane["三球"] = 2
		//lane["第四球"] = 3
		//lane["第五球"] = 4
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
	case 20:
		lane := make(map[string]int)
		lane["一球"] = 0
		lane["二球"] = 1
		lane["三球"] = 2
		//lane["第四球"] = 3
		//lane["第五球"] = 4
		arr := strings.Split(d.bet.Detail, "@")
		if d.arr[lane[arr[0]]] == arr[1] {
			return 1
		}
	case 21:
		if n1 > n3 && d.bet.Detail == "龙" {
			return 1
		}
		if n1 < n3 && d.bet.Detail == "虎" {
			return 1
		}
	case 22:
		if n1 == n3 {
			return 1
		}
	case 23:
		if n1 == n2 && n2 == n3 {
			return 1
		}
	case 24:
		n := []int{n1, n2, n3}
		return d.Straight(n)
	case 25:
		n := []int{n1, n2, n3}
		return d.SemiShun(n)
	//	 对子
	case 26:
		n := []int{n1, n2, n3}
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
	case 27:
		n := []int{n1, n2, n3}
		return d.Miscellaneous(n)
	}
	return 0
}
