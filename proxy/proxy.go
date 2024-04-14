package proxy

import "fmt"

type PaymentService interface {
	pay(order string) string
}

type WXPay struct {
}

func (w *WXPay) pay(order string) string {
	return "从微信支付获取token"
}

type ALiPay struct {
}

func (a *ALiPay) pay(order string) string {
	return "从阿里支付获取token"
}

type PaymentProxy struct {
	realPay PaymentService
}

func (p *PaymentProxy) pay(order string) string {
	fmt.Println("处理" + order)
	fmt.Println("1校验签名")
	fmt.Println("2格式化订单数据")
	fmt.Println("3参数检查")
	fmt.Println("4记录请求日志")
	token := p.realPay.pay(order)
	return "http://组装" + token + "然后跳转到第三方支付"
}
