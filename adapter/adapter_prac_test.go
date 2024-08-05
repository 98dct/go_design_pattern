package adapter

import (
	"fmt"
	"testing"
)

/*
	适配器模式是一种行为型模式，它允许不兼容的接口协同工作，该模式通过创建一个适配器类，封装不兼容的接口，并对外
	提供一个兼容的接口；
	将一个类的接口转换成客户希望的另外一个接口，adapter模式使得原本由于接口不兼容而不能在一起工作的那些类可以在一起工作
	举例：手机或者电脑插头通过适配器与插座连接在一起
	之所以叫适配器模式，作用必然是为了适配，更多的是一种事后补偿机制，并不是首选
*/

type PaymentProcessor interface {
	Pay(amount float64)
}

// 旧的支付系统
type OldPaymentSystem struct {
}

func (ops *OldPaymentSystem) Pay(amount float64) {
	fmt.Printf("Processing payment of %.2f using old payment system\n", amount)
}

// 新的支付系统
// 没有实现Pay方法，没有实现统一的PaymentProcessor接口
type NewPaymentSystem struct {
}

func (nps *NewPaymentSystem) MakePayment(amount float64) {
	fmt.Printf("Making payment of %.2f using new payment system\n", amount)
}

// 创建一个新支付系统的适配器
type NewPaymentAdapter struct {
	// 内部持有新支付系统
	NewSystem *NewPaymentSystem
}

func (npa *NewPaymentAdapter) Pay(amount float64) {
	npa.NewSystem.MakePayment(amount)
}

// 旧支付系统支付，不需要适配器
func Test1(t *testing.T) {
	var processor PaymentProcessor
	processor = &OldPaymentSystem{}
	processor.Pay(100)
}

// 新支付系统支付，适配器包装了下新支付系统
func Test2(t *testing.T) {
	newPayment := &NewPaymentSystem{}
	processor := &NewPaymentAdapter{NewSystem: newPayment}
	processor.Pay(100)
}
