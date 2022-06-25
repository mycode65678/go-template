package duijian

import (
	"errors"
	"fmt"
	"github.com/gocraft/work"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"hash2/a"
	"hash2/orm"
	"regexp"
	"strconv"
	"time"
)

//兑奖任务开始
type Dj struct {
}

//实际处理的
type Dispose struct {
	db     *gorm.DB
	u_id   int64
	bet    orm.UserBet
	user   orm.User
	time   time.Time
	l_id   int64
	result map[string]string
	arr    []string
	// 13-14特殊规则
	Rule orm.BackwaterRule
}

//兑奖任务
func (d *Dj) DuiJian(job *work.Job) error {
	// Extract arguments:
	l_id := job.ArgInt64("l_id")
	a.NewLog.WithField("name", "duijian").Info("兑奖开始进行中,本次兑奖彩种ID:", l_id, "兑奖开始时间:", time.Now().Local().Format("2006-01-02 15:04:05.9999"))
	//查找未兑奖的
	lotteryResult := make([]orm.LotteryResult, 0)
	orm.Db.Debug().Where("l_id = ?", l_id).Where("dispose = 0").Find(&lotteryResult)
	expect := make([]string, 0)
	result := make(map[string]string)
	for _, v := range lotteryResult {
		//获取未处理的开奖号码
		expect = append(expect, v.Expect)
		result[v.Expect] = v.Result
	}
	BackwaterRule := new(orm.BackwaterRule).GetMap()
	//fmt.Println("BackwaterRule",BackwaterRule)
	userBet := make([]orm.UserBet, 0)
	orm.Db.Debug().Preload("User").Preload("Lottery").
		Where("expect in (?)", expect).
		Where("l_id = ?", l_id).
		Where("status = 0").
		Find(&userBet)
	//开始进行兑奖操作
	for _, bet := range userBet {
		dispose := new(Dispose)
		dispose.result = make(map[string]string)
		reg, _ := regexp.Compile(`\d+`)
		dispose.arr = reg.FindAllString(result[bet.Expect], -1)
		//dispose.result[bet.Expect] = result[bet.Expect]
		dispose.bet = bet
		dispose.user = bet.User
		dispose.time = time.Now().Local()
		dispose.l_id = l_id
		// 判断是否有规则添加
		if _, ok := BackwaterRule[l_id]; ok {
			//fmt.Println("BackwaterRule[l_id][bet.Pan]",BackwaterRule[l_id][bet.Pan])
			//logrus.Debug("BackwaterRule[l_id][bet.Pan]",BackwaterRule[l_id][bet.Pan])
			//if _, ok2 := BackwaterRule[l_id][bet.Pan]; ok2 {
			//	dispose.Rule = BackwaterRule[l_id][bet.Pan]
			//}
		}
		dispose.start()
	}

	for _, v := range expect {
		ub := make([]orm.UserBet, 0)
		e := orm.Db.Where("expect = ?", v).Where("status = 0").Where("l_id = ?", l_id).Find(&ub)
		if e.RowsAffected == 0 {
			orm.Db.Model(&lotteryResult).Where("l_id = ?", l_id).Where("expect = ?", v).Where("dispose = 0").Update("dispose", 1)
		} else {
			fmt.Println("expect 未结算完成期号:", expect, e.RowsAffected, ub[0].Id, time.Now().Format("2006-01-02 15:04:05"))
			//enqueuer := work.NewEnqueuer(tasks.GO_DUIJIAN, orm.RedisClient)
			//enqueuer.Enqueue(tasks.GO_DUIJIAN, work.Q{"l_id": l_id})
			//job.
			return errors.New(fmt.Sprintf("期号：%v,未兑奖数量:%v,兑奖游戏ID：%v", expect, e.RowsAffected, l_id))
		}
	}
	fmt.Println("job.ArgError", job.ArgError())
	if err := job.ArgError(); err != nil {
		fmt.Println("job.ArgError", job.ArgError())
		return err
	}
	fmt.Println("job.ArgError() end", nil)
	// Go ahead and send the email...
	// sendEmailTo(addr, subject)

	return nil
}

func (d *Dispose) start() {
	d.db = orm.Db.Begin()
	orm.Block{
		Try: func() {
			// 进行资源锁定，防止重复派奖的情况
			lockBet := new(orm.UserBet)
			d.db.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", d.bet.Id).Find(&lockBet)
			if lockBet.Status != 0 {
				panic(fmt.Sprintf("已处理的订单,订单编号:%v", lockBet.Id))
			}

			d.u_id = d.bet.UId
			//进行兑奖,返回中奖的次数
			num := d.Control(int(d.l_id))
			total_bonus := decimal.NewFromFloat(0)
			if num == 1 {
				Money, _ := decimal.NewFromString(d.bet.TotalMoney)
				bonus, _ := decimal.NewFromString(d.bet.Bonus)
				total_bonus = Money.Mul(bonus)
			}

			userCaptial := new(orm.UserCaptial)
			e := d.db.Set("gorm:query_option", "FOR UPDATE").Where("u_id = ?", d.u_id).First(&userCaptial)
			if e.Error != nil {
				fmt.Println("no user_captial", e.Error)
				panic("no user_captial,error")
			}
			//fmt.Println("d=======", d.bet.Id, d.result)

			captialMap := make(map[string]interface{})
			captial := new(orm.UserCaptial)
			captial.UId = d.u_id
			//计算充值消费额
			captialMap["consume"] = gorm.Expr("consume - ?", d.bet.TotalMoney)

			if total_bonus.GreaterThan(decimal.NewFromFloat(0)) {
				//用户金额变动
				captialMap["use_money"] = gorm.Expr("use_money + ?", total_bonus.String())
				//账变更改
				debt := new(orm.UserDebt)
				debt.UId = d.u_id
				debt.Practical = total_bonus.String()
				debt.RecordId = strconv.FormatInt(d.bet.Id, 10)
				debt.Type = 4
				//tm, _ := decimal.NewFromString(d.bet.TotalMoney)
				//debt.Digest = d.bet.Lottery.Name + "," + d.bet.Period + "," + d.bet.Detail + ",下注" + tm.String() + "元宝,中奖" + total_bonus.String() + "元宝"
				e := debt.AddDebt(d.db)
				if e.(*orm.MyError).Code != 0 {
					panic(e)
				}
			}
			//fmt.Println("d=======11", d.bet.Id, d.result)

			ec := captial.Update(d.db, captialMap)
			if ec != nil {
				panic(ec)
			}
			fmt.Println("d=======2", d.bet.Id, d.result)

			//进行数据统计
			SysStatistics := new(orm.SysStatistics)
			SysStatistics.UId = d.bet.UId
			SysStatistics.Time = orm.Date(time.Now())
			SysStatistics.ActiveBet, _ = strconv.ParseFloat(d.bet.TotalMoney, 10)
			statis := make(map[string]interface{})
			statis["bet"] = gorm.Expr(" bet + ?", d.bet.TotalMoney)
			if num > 0 {
				statis["bonus"] = gorm.Expr(" bonus + ?", total_bonus.String())
			}

			//statis["num"] = gorm.Expr(" num + 1")
			err := SysStatistics.UpdateSysStatistics(d.db, statis)
			if err != nil {
				panic(err)
			}
			// 检测是否有活动的赠送金额

			//更新订单记录
			d.bet.TotalBonus = total_bonus.String()
			d.bet.Status = 1
			//计算本期盈利
			totalMoney, _ := decimal.NewFromString(d.bet.TotalMoney)
			d.bet.Profit = total_bonus.Sub(totalMoney).String()
			//d.bet.Result = d.result
			d.bet.UpdatedAt = orm.JSONTime{time.Now().Local()}
			// 分润计算
			//d.distribution()

			//d.bet.
			suc := d.db.Model(&d.bet).Where("status = 0").Save(&d.bet)
			//panic("test")
			fmt.Println("suc", suc.Error)
			if suc.RowsAffected == 0 {
				panic(suc.Error)
			}
			d.db.Commit()
		},
		Catch: func(e interface{}) {

		},
	}.Do()
}

//兑奖调度
func (d *Dispose) Control(l_id int) int {
	lottery := new(orm.Lottery)
	e := d.db.Where("id = ?", l_id).First(&lottery)
	if e.RowsAffected == 0 {
		panic("find lottery error")
	}
	//a.NewLog.Info("duijian Control", d.bet.Id, lottery.Type, d.bet.Expect, d.arr, d.bet.Detail)
	switch lottery.Type {
	//	28
	case "0":
		return d.xy28()
	case "1":
		return d.d3()
	case "2":
		return d.d5()
		//return d.ssc()
	}
	return 0
}
