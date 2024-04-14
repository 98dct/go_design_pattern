package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := &PaymentProxy{realPay: &ALiPay{}}
	res := proxy.pay("阿里订单")
	fmt.Println(res)
}
