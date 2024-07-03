package service

import (
	"payment/domain/model"
	"payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindAllPayment() ([]model.Payment, error)
	FindPaymentByID(int64) (*model.Payment, error)
}

func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService {
	return &PaymentDataService{PaymentRepository: paymentRepository}
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}

// 添加订单
func (o *PaymentDataService) AddPayment(payment *model.Payment) (int64, error) {
	return o.PaymentRepository.CreatePayment(payment)
}

// 删除订单
func (o *PaymentDataService) DeletePayment(paymentID int64) error {
	return o.PaymentRepository.DeletePaymentByID(paymentID)
}

// 更新订单
func (o *PaymentDataService) UpdatePayment(payment *model.Payment) (err error) {
	return o.PaymentRepository.UpdatePayment(payment)
}

// 根据订单ID查找订单
func (o *PaymentDataService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return o.PaymentRepository.FindPaymentByID(paymentID)
}

// 查找所有订单
func (o *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return o.PaymentRepository.FindAll()
}
