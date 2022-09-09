package controller

import (
	"context"
	"fmt"
	"micro_product/micro_common/utils"
	"micro_product/micro_proto/pc"
	"micro_product/models"
	"micro_product/services"
)

type DcProduct struct {
}

func (p *DcProduct) GetProduct(ctx context.Context, dto *pc.GetProductDto) (*pc.BaseResponse, error) {

	fmt.Println("GetProduct入参：", utils.JsonToString(dto))
	utils.MyLog.Info("GetProduct入参：", utils.JsonToString(dto))

	var res = &pc.BaseResponse{
		Code:  200,
		Msg:   "",
		Error: "",
		Data:  "",
	}

	conn := services.NewDbConn()

	var data []models.Product
	err := conn.Select("*").Limit(10).Find(&data)
	if err != nil {
		return res, err
	}

	fmt.Println("查询的商品数据： ", utils.JsonToString(data))
	utils.MyLog.Info("查询商品数据：", utils.JsonToString(data))

	res.Data = utils.JsonToString(data)
	return res, nil
}

func (p *DcProduct) TestProduct(ctx context.Context, dto *pc.GetProductDto) (*pc.BaseResponse, error) {

	fmt.Println("TestProduct：", utils.JsonToString(dto))

	var res = &pc.BaseResponse{
		Code:  200,
		Msg:   "",
		Error: "",
		Data:  "",
	}

	conn := services.NewDbConn()

	var data []models.Product
	err := conn.Select("*").Limit(1).Find(&data)
	if err != nil {
		return res, err
	}

	fmt.Println("TestProduct： ", utils.JsonToString(data))

	res.Data = utils.JsonToString(data)
	return res, nil
}
