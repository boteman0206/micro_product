package controller

import (
	"context"
	"fmt"
	"micro_product/micro_proto/pc"
	"micro_product/utils"
)

type DcProduct struct {
}

func (p *DcProduct) GetProduct(ctx context.Context, dto *pc.GetProductDto) (*pc.BaseResponse, error) {

	fmt.Println("GetProduct入参：", utils.JsonToString(dto))

	var res = &pc.BaseResponse{
		Code:  200,
		Msg:   "",
		Error: "",
		Data:  "",
	}

	return res, nil
}
