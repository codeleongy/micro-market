package handler

import (
	"context"

	"cart/domain/model"
	"cart/domain/service"
	pb "cart/proto"

	"micro-market/common"

	"go-micro.dev/v4/logger"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, req *pb.CartInfo, rsp *pb.ResponseAdd) error {
	cart := &model.Cart{}
	if err := common.SwapTo(req, cart); err != nil {
		logger.Error(err)
		return err
	}

	rsp.CartId = cart.ID
	rsp.Msg = "购物车添加成功"
	return nil
}

func (c *Cart) CleanCart(ctx context.Context, req *pb.Clean, rsp *pb.Response) error {
	if err := c.CartDataService.CleanCart(req.UserId); err != nil {
		logger.Error(err)
		return err
	}
	rsp.Msg = "购物车清空成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, req *pb.Item, rsp *pb.Response) error {
	if err := c.CartDataService.IncrNum(req.Id, req.ChangeNum); err != nil {
		logger.Error(err)
		return err
	}
	rsp.Msg = "购物车添加商品成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, req *pb.Item, rsp *pb.Response) error {
	if err := c.CartDataService.DecrNum(req.Id, req.ChangeNum); err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "购物车删除商品成功"
	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, req *pb.CartID, rsp *pb.Response) error {
	if err := c.CartDataService.DeleteCart(req.Id); err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "购物车删除成功"
	return nil

}

// 获取购物车所有条目
func (c *Cart) GetAll(ctx context.Context, req *pb.CartFindAll, rsp *pb.CartAll) error {
	carts, err := c.CartDataService.FindAllCart(req.UserId)

	if err != nil {
		logger.Error(err)
		return err
	}

	for _, v := range carts {
		cartRsp := &pb.CartInfo{}
		if err := common.SwapTo(v, cartRsp); err != nil {
			logger.Error(err)
			return err
		}

		rsp.CartInfos = append(rsp.CartInfos, cartRsp)
	}

	return nil
}
