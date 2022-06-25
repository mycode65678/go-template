package duijian

import (
	"strconv"
)

func (d *Dispose) xy28() int {
	n1, _ := strconv.Atoi(d.arr[0])
	n2, _ := strconv.Atoi(d.arr[1])
	n3, _ := strconv.Atoi(d.arr[2])
	sum := n1 + n2 + n3
	//a.NewLog.WithFields(logrus.Fields{
	//	"game":         "xy28",
	//	"play_kind_id": d.bet.PlayKindId,
	//}).Debug("xy28", d.bet.Detail, "arr", d.arr)
	switch d.bet.PlayKindId {
	case 1:
		if d.bet.Detail == "大" && sum >= 14 {
			return 1
		}
		if d.bet.Detail == "小" && sum < 14 {
			return 1
		}
		if d.bet.Detail == "单" && sum%2 == 1 {
			return 1
		}
		if d.bet.Detail == "双" && sum%2 == 0 {
			return 1
		}
	case 2, 3:
		if d.bet.Detail == "大单" && sum >= 14 && sum%2 == 1 {
			return 1
		}
		if d.bet.Detail == "大双" && sum >= 14 && sum%2 == 0 {
			return 1
		}
		if d.bet.Detail == "小单" && sum < 14 && sum%2 == 1 {
			return 1
		}
		if d.bet.Detail == "小双" && sum < 14 && sum%2 == 0 {
			return 1
		}
	case 4:
		if d.bet.Detail == "极大" && sum >= 22 {
			return 1
		}
		if d.bet.Detail == "极小" && sum < 6 {
			return 1
		}
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		if d.bet.Detail == strconv.Itoa(sum) {
			return 1
		}
		//	豹子
	case 18:
		if n1 == n2 && n1 == n3 {
			return 1
		}
	case 17:
		color := map[int]string{
			0:  "黄波",
			1:  "绿波",
			2:  "蓝波",
			3:  "红波",
			4:  "绿波",
			5:  "蓝波",
			6:  "红波",
			7:  "绿波",
			8:  "蓝波",
			9:  "红波",
			10: "绿波",
			11: "蓝波",
			12: "红波",
			13: "黄波",
			14: "黄波",
			15: "红波",
			16: "绿波",
			17: "蓝波",
			18: "红波",
			19: "绿波",
			20: "蓝波",
			21: "红波",
			22: "绿波",
			23: "蓝波",
			24: "红波",
			25: "绿波",
			26: "蓝波",
			27: "黄波",
		}
		if color[sum] == d.bet.Detail {
			return 1
		}
		return 0
	}
	return 0
}
