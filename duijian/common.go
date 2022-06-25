package duijian

import (
	"github.com/shopspring/decimal"
)

//对用户增加金额
//money 变动金额,type 账变金额
func (d *Dispose) ChangeMoney(Money decimal.Decimal, Type int) {
	if Money.Equal(decimal.NewFromFloat(0)) {
		return
	}
	sql := "update user_captial set use_money = use_money + ? where u_id = ?"
	d.db.Exec(sql, Money, d.u_id)
	//进行账变
}
