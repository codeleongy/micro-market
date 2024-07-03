package handler

import (
	"context"
	"micro-market/common"
	"payment/domain/model"
	"payment/domain/service"
	pb "payment/proto"
)

var (
	logger = common.GetLogger()
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (p *Payment) AddPayment(ctx context.Context, in *pb.PaymentInfo, out *pb.PaymentID) error {
	paymentInfo := &model.Payment{}
	err := common.SwapTo(in, paymentInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	paymentId, err := p.PaymentDataService.AddPayment(paymentInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	out.PaymentId = paymentId
	return err
}

func (p *Payment) UpdatePayment(ctx context.Context, in *pb.PaymentInfo, out *pb.Response) error {
	paymentInfo := &model.Payment{}
	err := common.SwapTo(in, paymentInfo)
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := p.PaymentDataService.UpdatePayment(paymentInfo); err != nil {
		logger.Error(err)
		return err
	}
	out.Msg = "支付信息更新成功"

	return nil
}

func (p *Payment) DeletePaymentByID(ctx context.Context, in *pb.PaymentID, out *pb.Response) error {
	if err := p.PaymentDataService.DeletePayment(in.PaymentId); err != nil {
		logger.Error(err)
		return err
	}
	out.Msg = "支付信息删除成功"
	return nil
}

func (p *Payment) FindPaymentByID(ctx context.Context, in *pb.PaymentID, out *pb.PaymentInfo) error {
	paymentInfo, err := p.PaymentDataService.FindPaymentByID(in.PaymentId)

	if err != nil {
		logger.Error(err)
		return err
	}

	err = common.SwapTo(paymentInfo, out)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (p *Payment) FindAllPayment(ctx context.Context, in *pb.All, out *pb.AllPayment) error {
	payments, err := p.PaymentDataService.FindAllPayment()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, v := range payments {
		res := &pb.PaymentInfo{}
		if err := common.SwapTo(v, res); err != nil {
			logger.Error(err)
			return err
		}
		out.PaymentInfos = append(out.PaymentInfos, res)
	}
	return nil
}
