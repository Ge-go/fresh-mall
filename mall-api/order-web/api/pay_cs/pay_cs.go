package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

// 支付宝支付代码实验
func main() {
	appID := ""
	privateKey := ""
	aliPubKey := ""
	var client, err = alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}

	err = client.LoadAliPayPublicKey(aliPubKey)
	if err != nil {
		panic(err)
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "https://www.baidu.com"
	p.ReturnURL = "https://www.baidu.com"
	p.Subject = "mall-fresh订单支付"
	p.OutTradeNo = "wS-mall"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}
