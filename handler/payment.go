package handler

import (
	"context"
	common "github.com/ljp-lachouchou/go-learn-micro-common"
	"github.com/ljp-lachouchou/go-micro-learn-payment/domain/model"
	"github.com/ljp-lachouchou/go-micro-learn-payment/domain/service"
	. "github.com/ljp-lachouchou/go-micro-learn-payment/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (p *Payment) AddPayment(ctx context.Context,
	req *PaymentInfo, res *PaymentID) error {
	pt := &model.Payment{}
	if err := common.SwapTo(req, pt); err != nil {
		common.Error(err)
		return err
	}
	paymentID, err := p.PaymentDataService.AddPayment(pt)
	if err != nil {
		common.Error(err)
		return err
	}
	res.PaymentId = paymentID
	return nil
}
func (p *Payment) UpdatePayment(ctx context.Context,
	req *PaymentInfo, res *Response) error {
	pt := &model.Payment{}
	if err := common.SwapTo(req, pt); err != nil {
		common.Error(err)
		return err
	}
	err := p.PaymentDataService.UpdatePayment(pt)
	if err != nil {
		common.Error(err)
		return err
	}
	return nil
}
func (p *Payment) DeletePaymentByID(ctx context.Context,
	req *PaymentID, res *Response) error {
	err := p.PaymentDataService.DeletePayment(req.PaymentId)
	if err != nil {
		common.Error(err)
		return err
	}
	res.Msg = "支付删除"
	return nil
}
func (p *Payment) FindPaymentByID(ctx context.Context,
	req *PaymentID, res *PaymentInfo) error {
	pt, err := p.PaymentDataService.FindPaymentByID(req.PaymentId)
	if err != nil {
		common.Error(err)
		return err
	}
	if err := common.SwapTo(pt, res); err != nil {
		common.Error(err)
		return err
	}
	return nil
}
func (p *Payment) FindAllPayment(ctx context.Context,
	req *All, res *PaymentAll) error {
	allPayment, err := p.PaymentDataService.FindAllPayment()
	if err != nil {
		common.Error(err)
		return err
	}
	for _, pt := range allPayment {
		info := &PaymentInfo{}
		if err := common.SwapTo(pt, info); err != nil {
			common.Error(err)
			return err
		}
		res.PaymentInfo = append(res.PaymentInfo, info)

	}
	return nil
}
