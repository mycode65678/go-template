package duijian

import (
	"sort"
	"strconv"
	"strings"
)

func (d *Dispose) ssc() int {
	//fmt.Println("d", d.bet.Result)
	//fmt.Println("result", d.result)
	switch d.bet.PlayKindId {
	//定位胆
	case 468:
		lane := make(map[string]int)
		lane["第一球"] = 0
		lane["第二球"] = 1
		lane["第三球"] = 2
		lane["第四球"] = 3
		lane["第五球"] = 4
		arr := strings.Split(d.bet.Detail, "@")
		if d.arr[lane[arr[0]]] == arr[1] {
			return 1
		}
	//	两面盘
	case 469:
		lane := make(map[string]int)
		lane["第一球"] = 0
		lane["第二球"] = 1
		lane["第三球"] = 2
		lane["第四球"] = 3
		lane["第五球"] = 4
		arr := strings.Split(d.bet.Detail, "@")
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
	//	总和
	case 470:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		n3, _ := strconv.Atoi(d.arr[2])
		n4, _ := strconv.Atoi(d.arr[3])
		n5, _ := strconv.Atoi(d.arr[4])
		sum := n1 + n2 + n3 + n4 + n5
		if sum%2 == 0 && d.bet.Detail == "总和@双" {
			return 1
		}
		if sum%2 == 1 && d.bet.Detail == "总和@单" {
			return 1
		}
		if sum > 22 && d.bet.Detail == "总和@大" {
			return 1
		}
		if sum < 23 && d.bet.Detail == "总和@小" {
			return 1
		}
	//	豹子
	case 532:
		if d.arr[0] == d.arr[1] && d.arr[0] == d.arr[2] && d.bet.Detail == "前三@豹子" {
			return 1
		}
		if d.arr[3] == d.arr[1] && d.arr[3] == d.arr[2] && d.bet.Detail == "中三@豹子" {
			return 1
		}
		if d.arr[3] == d.arr[2] && d.arr[3] == d.arr[4] && d.bet.Detail == "后三@豹子" {
			return 1
		}
	case 533:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		n3, _ := strconv.Atoi(d.arr[2])
		n4, _ := strconv.Atoi(d.arr[3])
		n5, _ := strconv.Atoi(d.arr[4])
		if d.bet.Detail == "前三@顺子" {
			n := []int{n1, n2, n3}
			return d.Straight(n)
		}
		if d.bet.Detail == "中三@顺子" {
			n := []int{n2, n3, n4}
			return d.Straight(n)
		}
		if d.bet.Detail == "后三@顺子" {
			n := []int{n3, n4, n5}
			return d.Straight(n)
		}
	case 534:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		n3, _ := strconv.Atoi(d.arr[2])
		n4, _ := strconv.Atoi(d.arr[3])
		n5, _ := strconv.Atoi(d.arr[4])
		if d.bet.Detail == "前三@对子" {
			n := []int{n1, n2, n3}
			if n[0] == n[1] && n[0] != n[2] {
				return 1
			}
			if n[1] == n[2] && n[2] != n[0] {
				return 1
			}
			if n[0] == n[2] && n[2] != n[1] {
				return 1
			}
		}
		if d.bet.Detail == "中三@对子" {
			n := []int{n2, n3, n4}
			if n[0] == n[1] && n[0] != n[2] {
				return 1
			}
			if n[1] == n[2] && n[2] != n[0] {
				return 1
			}
			if n[0] == n[2] && n[2] != n[1] {
				return 1
			}
		}
		if d.bet.Detail == "后三@对子" {
			n := []int{n3, n4, n5}
			if n[0] == n[1] && n[0] != n[2] {
				return 1
			}
			if n[1] == n[2] && n[2] != n[0] {
				return 1
			}
			if n[0] == n[2] && n[2] != n[1] {
				return 1
			}
		}
	//	半顺
	case 535:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		n3, _ := strconv.Atoi(d.arr[2])
		n4, _ := strconv.Atoi(d.arr[3])
		n5, _ := strconv.Atoi(d.arr[4])
		if d.bet.Detail == "前三@半顺" {
			n := []int{n1, n2, n3}
			return d.SemiShun(n)
		}
		if d.bet.Detail == "中三@半顺" {
			n := []int{n2, n3, n4}
			return d.SemiShun(n)
		}
		if d.bet.Detail == "后三@半顺" {
			n := []int{n3, n4, n5}
			return d.SemiShun(n)
		}
	case 536:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		n3, _ := strconv.Atoi(d.arr[2])
		n4, _ := strconv.Atoi(d.arr[3])
		n5, _ := strconv.Atoi(d.arr[4])
		if d.bet.Detail == "前三@杂六" {
			n := []int{n1, n2, n3}
			return d.Miscellaneous(n)
		}
		if d.bet.Detail == "中三@杂六" {
			n := []int{n2, n3, n4}
			return d.Miscellaneous(n)
		}
		if d.bet.Detail == "后三@杂六" {
			n := []int{n3, n4, n5}
			return d.Miscellaneous(n)
		}
	//	龙
	case 537:
		n1, _ := strconv.Atoi(d.arr[0])
		n5, _ := strconv.Atoi(d.arr[4])
		if n1 == n5 {
			d.bet.Bonus = d.bet.TotalMoney
			return 1
		}
		if n1 > n5 {
			return 1
		}
	case 538:
		n1, _ := strconv.Atoi(d.arr[0])
		n5, _ := strconv.Atoi(d.arr[4])
		if n1 == n5 {
			d.bet.Bonus = d.bet.TotalMoney
			return 1
		}
		if n1 < n5 {
			return 1
		}
	case 539:
		n1, _ := strconv.Atoi(d.arr[0])
		n5, _ := strconv.Atoi(d.arr[4])
		if n1 == n5 {
			return 1
		}
	}

	return 0
}

// 检测是否是顺子
func (d *Dispose) Straight(n []int) int {
	sort.Ints(n)
	if n[0] == 0 && n[2] == 9 && (n[1] == 8 || n[1] == 1) {
		return 1
	}
	if n[1]-n[0] == 1 && n[2]-n[1] == 1 {
		return 1
	}
	return 0
}

func (d *Dispose) SemiShun(n []int) int {
	sort.Ints(n)
	if n[1]-n[0] == 1 && n[2]-n[1] > 1 {
		return 1
	}
	if n[2]-n[1] == 1 && n[1]-n[0] > 1 {
		return 1
	}
	if n[0] == 0 && n[2] == 9 && (n[1] < 8 && n[1] > 1) {
		return 1
	}
	return 0
}

// 杂六
func (d *Dispose) Miscellaneous(n []int) int {
	sort.Ints(n)
	if n[0] == 0 && n[2] == 9 {
		return 0
	}
	if n[1]-n[0] > 1 && n[2]-n[1] > 1 {
		return 1
	}
	return 0
}
