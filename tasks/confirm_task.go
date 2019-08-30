package tasks

import (
	"go-weishan-shop-pay-server/global"
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/utils"
	"os"
	"strconv"
)

func buildOrder(data *global.OrderCache) *models.Order {
	return &models.Order{
		Goodname:    data.Goodname,
		GoodId:      data.GoodId,
		Realname:    data.Realname,
		Address:     data.Address,
		Phonenumber: data.Phonenumber,
		ExtInfo:     data.ExtInfo,
		BuyCount:    data.BuyCount,
		BuyPrice:    data.BuyPrice,
		Status:      false,
		SelfMention: data.SelfMention,
	}
}

func buildFee(data *global.OrderCache, to float64, fee float64) *models.Fee {
	return &models.Fee{
		TotalValue: data.BuyPrice,
		ToValue:    to,
		FeeValue:   fee,
	}
}

func createPayRecord(cachedata *global.OrderCache,
	to float64, fee float64) error {
	tx := models.DB.Begin()

	result := tx.Create(buildOrder(cachedata))
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	err := tx.Create(buildFee(cachedata, to, fee)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func ConfirmTask(data interface{}) error {
	d := data.(*global.OrderCache)

	feerate, err := strconv.ParseFloat(os.Getenv("FEE_RATE"), 64)

	if err != nil {
		panic(err)
	}

	to, fee := utils.CalcFee(d.BuyPrice, feerate)

	createPayRecord(d, to, fee)

	return nil
}
