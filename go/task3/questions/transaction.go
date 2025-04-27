package questions

import "fmt"

type Account struct {
	ID      uint64 `gorm:"primaryKey"`
	Balance int
}

type Transaction struct {
	ID              uint64 `gorm:"primaryKey"`
	From_account_id uint64
	To_account_id   uint64
	Amount          int
}

func TransactionRun(aid uint64, bid uint64, amount int) {
	tx := db.Begin()
	var accountA Account
	var accountB Account
	if err := tx.First(&accountA, aid).Error; err != nil {
		fmt.Println("查询账户A失败")
		tx.Rollback()
		return
	}
	if accountA.Balance < amount {
		fmt.Println("余额不足")
		tx.Rollback()
		return
	}
	if err := tx.Model(&Account{}).Where("id = ?", aid).Update("balance", accountA.Balance-amount).Error; err != nil {
		fmt.Println("账户A扣款失败")
		tx.Rollback()
		return
	}
	if err := tx.First(&accountB, bid).Error; err != nil {
		fmt.Println("查询账户B失败")
		tx.Rollback()
		return
	}
	if err := tx.Model(&Account{}).Where("id =?", bid).Update("balance", accountB.Balance+amount).Error; err != nil {
		fmt.Println("账户B存款失败")
		tx.Rollback()
		return
	}
	if err := tx.Create(&Transaction{From_account_id: aid, To_account_id: bid, Amount: amount}).Error; err != nil {
		fmt.Println("创建交易记录失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	fmt.Println("转账成功")
}
