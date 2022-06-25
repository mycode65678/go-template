package duijian

import (
	"strconv"
	"strings"
)

func (d *Dispose) hk6() int {

	var color = map[string]string{
		"01": "红波",
		"02": "红波",
		"07": "红波",
		"08": "红波",
		"12": "红波",
		"13": "红波",
		"18": "红波",
		"19": "红波",
		"23": "红波",
		"24": "红波",
		"29": "红波",
		"30": "红波",
		"34": "红波",
		"35": "红波",
		"40": "红波",
		"45": "红波",
		"46": "红波",
		"03": "蓝波",
		"04": "蓝波",
		"09": "蓝波",
		"10": "蓝波",
		"14": "蓝波",
		"15": "蓝波",
		"20": "蓝波",
		"25": "蓝波",
		"26": "蓝波",
		"31": "蓝波",
		"36": "蓝波",
		"37": "蓝波",
		"41": "蓝波",
		"42": "蓝波",
		"47": "蓝波",
		"48": "蓝波",
		"05": "绿波",
		"06": "绿波",
		"11": "绿波",
		"16": "绿波",
		"17": "绿波",
		"21": "绿波",
		"22": "绿波",
		"27": "绿波",
		"28": "绿波",
		"32": "绿波",
		"33": "绿波",
		"38": "绿波",
		"39": "绿波",
		"43": "绿波",
		"44": "绿波",
		"49": "绿波",
	}
	var zodiac = map[string]string{
		"01": "狗",
		"02": "鸡",
		"03": "猴",
		"04": "羊",
		"05": "马",
		"06": "蛇",
		"07": "龙",
		"08": "兔",
		"09": "虎",
		"10": "牛",
		"11": "鼠",
		"12": "猪",
		"13": "狗",
		"14": "鸡",
		"15": "猴",
		"16": "羊",
		"17": "马",
		"18": "蛇",
		"19": "龙",
		"20": "兔",
		"21": "虎",
		"22": "牛",
		"23": "鼠",
		"24": "猪",
		"25": "狗",
		"26": "鸡",
		"27": "猴",
		"28": "羊",
		"29": "马",
		"30": "蛇",
		"31": "龙",
		"32": "兔",
		"33": "虎",
		"34": "牛",
		"35": "鼠",
		"36": "猪",
		"37": "狗",
		"38": "鸡",
		"39": "猴",
		"40": "羊",
		"41": "马",
		"42": "蛇",
		"43": "龙",
		"44": "兔",
		"45": "虎",
		"46": "牛",
		"47": "鼠",
		"48": "猪",
		"49": "狗",
	}
	var lane = map[string]int{
		"正一特": 0,
		"正二特": 1,
		"正三特": 2,
		"正四特": 3,
		"正五特": 4,
		"正六特": 5,
		"特码":  6,
	}
	switch d.bet.PlayKindId {
	//特码
	case 435:

		arr := strings.Split(d.bet.Detail, "@")
		key := lane[arr[0]]
		if d.arr[key] == arr[1] {
			return 1
		}
		//两面
	case 443:
		arr := strings.Split(d.bet.Detail, "@")
		sum := 0
		for _, v := range d.arr {
			tempInt, _ := strconv.Atoi(v)
			sum += tempInt
		}
		if d.bet.Detail == "总大" && sum > 174 {
			return 1
		}
		if d.bet.Detail == "总小" && sum < 175 {
			return 1
		}
		if d.bet.Detail == "总单" && sum%2 == 1 {
			return 1
		}
		if d.bet.Detail == "总双" && sum%2 == 0 {
			return 1
		}
		if len(arr) == 1 {
			return 0
		}
		//合数
		key := lane[arr[0]]
		res := strings.Split(d.arr[key], "")
		resInt, _ := strconv.Atoi(d.arr[key])
		h1, _ := strconv.Atoi(res[0])
		h2, _ := strconv.Atoi(res[1])
		he := h1 + h2

		if arr[1] == "大" && resInt > 24 {
			if resInt == 49 {
				return -1
			}
			return 1
		}
		if arr[1] == "小" && resInt < 25 {
			return 1
		}
		if arr[1] == "单" && resInt%2 == 1 {
			if resInt == 49 {
				return -1
			}
			return 1
		}
		if arr[1] == "双" && resInt%2 == 0 {
			return 1
		}
		if arr[1] == "合大" && he > 6 {
			if resInt == 49 {
				return -1
			}
			return 1
		}
		if arr[1] == "合小" && he < 7 {
			return 1
		}

		if arr[1] == "尾大" && h2 > 4 {
			if resInt == 49 {
				return -1
			}
			return 1
		}
		if arr[1] == "尾小" && h2 < 5 {
			return 1
		}

		if arr[1] == "小单" && resInt < 25 && resInt%2 == 1 {
			return 1
		}

		if arr[1] == "大单" && resInt > 24 && resInt%2 == 1 {
			if resInt == 49 {
				return -1
			}
			return 1
		}

		if arr[1] == "小双" && resInt < 25 && resInt%2 == 0 {
			return 1
		}

		if arr[1] == "大双" && resInt > 24 && resInt%2 == 0 {
			return 1
		}

		//红波
	case 444, 445, 446:

		//判断开奖的波色
		arr := strings.Split(d.bet.Detail, "@")
		if len(arr) == 1 {
			if color[d.arr[6]] == d.bet.Detail {
				return 1
			} else {
				return 0
			}
		}
		key := lane[arr[0]]
		if color[d.arr[key]] == arr[1] {
			return 1
		}
		//红单
	case 483:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])

		if resInt%2 == 1 && color[d.arr[6]] == "红波" && d.bet.Detail == "红单" {
			return 1
		}
	case 484:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt%2 == 1 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝单" {
			return 1
		}
	case 485:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt%2 == 1 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿单" {
			return 1
		}

	case 486:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt%2 == 0 && color[d.arr[6]] == "红波" && d.bet.Detail == "红双" {
			return 1
		}
	case 487:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt%2 == 0 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝双" {
			return 1
		}
	case 488:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt%2 == 0 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿双" {
			return 1
		}
	case 489:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && color[d.arr[6]] == "红波" && d.bet.Detail == "红大" {
			return 1
		}
	case 490:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝大" {
			return 1
		}
	case 491:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿大" {
			return 1
		}
	case 492:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && color[d.arr[6]] == "红波" && d.bet.Detail == "红小" {
			return 1
		}
	case 493:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝小" {
			return 1
		}
	case 494:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿小" {
			return 1
		}
	case 495:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 1 && color[d.arr[6]] == "红波" && d.bet.Detail == "红大单" {
			return 1
		}
	case 496:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 0 && color[d.arr[6]] == "红波" && d.bet.Detail == "红大双" {
			return 1
		}
	case 497:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 1 && color[d.arr[6]] == "红波" && d.bet.Detail == "红小单" {
			return 1
		}
	case 498:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 0 && color[d.arr[6]] == "红波" && d.bet.Detail == "红小双" {
			return 1
		}
	case 499:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 1 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝大单" {
			return 1
		}
	case 500:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 0 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝大双" {
			return 1
		}
	case 501:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 1 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝小单" {
			return 1
		}
	case 502:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 0 && color[d.arr[6]] == "蓝波" && d.bet.Detail == "蓝小双" {
			return 1
		}
	case 503:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 1 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿大单" {
			return 1
		}
	case 504:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt > 24 && resInt%2 == 0 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿大双" {
			return 1
		}
	case 505:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 1 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿小单" {
			return 1
		}
	case 506:
		if d.arr[6] == "49" {
			return -1
		}
		resInt, _ := strconv.Atoi(d.arr[6])
		if resInt < 25 && resInt%2 == 0 && color[d.arr[6]] == "绿波" && d.bet.Detail == "绿小双" {
			return 1
		}
	case 507, 508:
		arr := strings.Split(d.bet.Detail, "@")
		if zodiac[d.arr[6]] == arr[1] {
			return 1
		}
	case 509, 510:
		arr := strings.Split(d.bet.Detail, "@")
		res := strings.Split(d.arr[6], "")
		if arr[1] == res[0] {
			return 1
		}
	case 511, 512:
		arr := strings.Split(d.bet.Detail, "@")
		res := strings.Split(d.arr[6], "")
		if arr[1] == res[1] {
			return 1
		}
		//正码
	case 513:
		arr := strings.Split(d.bet.Detail, "@")
		for i := 0; i < 6; i++ {
			if d.arr[i] == arr[1] {
				return 1
			}
		}
	}

	return 0
}
