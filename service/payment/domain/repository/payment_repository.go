package repository

import (
	"payment/domain/model"

	"github.com/jinzhu/gorm"
)

type IPaymentRepository interface {
	// 初始化数据表
	InitTable() error
	// 创建支付
	CreatePayment(*model.Payment) (int64, error)
	// 根据支付ID删除支付
	DeletePaymentByID(int64) error
	// 更新支付信息
	UpdatePayment(*model.Payment) error
	// 查找所有支付
	FindAll() ([]model.Payment, error)
	// 根据支付ID查找支付信息
	FindPaymentByID(int64) (*model.Payment, error)
}

// 创建PaymentRepository
func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{
		mysqlDB: db,
	}
}

type PaymentRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (p *PaymentRepository) InitTable() error {
	return p.mysqlDB.CreateTable(&model.Payment{}).Error
}

// 根据支付ID查找支付信息
func (p *PaymentRepository) FindPaymentByID(paymentID int64) (Payment *model.Payment, err error) {
	return Payment, p.mysqlDB.First(Payment, paymentID).Error
}

// 创建支付
func (p *PaymentRepository) CreatePayment(payment *model.Payment) (int64, error) {
	return payment.ID, p.mysqlDB.Create(payment).Error
}

// 根据支付ID删除支付
func (p *PaymentRepository) DeletePaymentByID(paymentID int64) error {
	return p.mysqlDB.Where("id = ?", paymentID).Delete(&model.Payment{}).Error
}

// 更新支付
func (p *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	return p.mysqlDB.Model(payment).Update(payment).Error
}

// 查找所有类目
func (p *PaymentRepository) FindAll() (payments []model.Payment, err error) {
	return payments, p.mysqlDB.Find(&payments).Error
}
