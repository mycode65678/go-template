package duijian

import (
	"strconv"
	"strings"
)

func (d *Dispose) pk() int {
	switch d.bet.PlayKindId {
	//	两面盘
	case 458:
		lane := make(map[string]int)
		lane["冠军"] = 0
		lane["亚军"] = 1
		lane["第三名"] = 2
		lane["第四名"] = 3
		lane["第五名"] = 4
		lane["第六名"] = 5
		lane["第七名"] = 6
		lane["第八名"] = 7
		lane["第九名"] = 8
		lane["第十名"] = 9

		arr := strings.Split(d.bet.Detail, "@")
		num, _ := strconv.Atoi(d.arr[lane[arr[0]]])

		if num > 5 && arr[1] == "大" {
			return 1
		}
		if num < 6 && arr[1] == "小" {
			return 1
		}

		if num%2 == 1 && arr[1] == "单" {
			return 1
		}
		if num%2 == 0 && arr[1] == "双" {
			return 1
		}
		// 龙虎判断
		n2, _ := strconv.Atoi(d.arr[9-lane[arr[0]]])
		//fmt.Println("id",d.bet.Id,"龙",num,n2,arr[1],num > n2,num < n2,"arr[1]",arr[1],d.arr)
		if num > n2 && arr[1] == "龙" {
			return 1
		}
		if num < n2 && arr[1] == "虎" {
			return 1
		}

	//	冠亚和
	case 459, 460:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		//str := strconv.Itoa(n1 + n2)
		arr := strings.Split(d.bet.Detail, "@")
		sum := n1 + n2
		if sum%2 == 1 && arr[1] == "单" {
			return 1
		}
		if sum%2 == 0 && arr[1] == "双" {
			return 1
		}
		if sum > 11 && arr[1] == "大" {
			return 1
		}
		if sum < 12 && arr[1] == "小" {
			return 1
		}
	//	单号1-10
	case 461:
		lane := make(map[string]int)
		lane["冠军"] = 0
		lane["亚军"] = 1
		lane["第三名"] = 2
		lane["第四名"] = 3
		lane["第五名"] = 4
		lane["第六名"] = 5
		lane["第七名"] = 6
		lane["第八名"] = 7
		lane["第九名"] = 8
		lane["第十名"] = 9

		arr := strings.Split(d.bet.Detail, "@")
		//num,_ := strconv.Atoi(d.arr[lane[arr[0]]])
		n1, _ := strconv.Atoi(d.arr[lane[arr[0]]])
		n2, _ := strconv.Atoi(arr[1])
		//fmt.Println("id", d.bet.Id, d.arr[lane[arr[0]]], arr[1], d.arr, arr, n1, n2)
		if n1 == n2 {
			return 1
		}
	//	和值
	case 463, 464, 465, 466, 467:
		n1, _ := strconv.Atoi(d.arr[0])
		n2, _ := strconv.Atoi(d.arr[1])
		str := strconv.Itoa(n1 + n2)
		if "冠亚和@"+str == d.bet.Detail {
			return 1
		}
	}
	return 0
}
