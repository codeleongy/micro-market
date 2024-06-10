package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	cart "github.com/codeleongy/micro-market/service/cart/proto"

	pb "github.com/codeleongy/micro-market/api/cartApi/proto"
	"go-micro.dev/v4/logger"
)

type CartApi struct {
	CartService cart.CartService
}

func (c *CartApi) FindAll(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	logger.Info("接收到/cartApi/findAll 访问请求")
	if _, ok := req.Get["user_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}

	userIdStr := req.Get["user_id"].Values[0]

	logger.Info(userIdStr)

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return err
	}

	// 获取购物车所有商品
	cartAll, err := c.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	if err != nil {
		logger.Error(err)
		return err
	}

	// 数据类型转换
	b, err := json.Marshal(cartAll)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
