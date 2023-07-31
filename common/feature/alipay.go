package feature

import (
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"

	"github.com/smartwalle/alipay/v3"
)

const (
	AppId        = "9021000123613685"
	PrivateKey   = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDf4/19hCVbrsShf/zXCU+/5OEQbEd4YLyOe8BYBB8L+xjfzQ6GuylME/fTQYEpdIYcHU6OfKhgmLlsQllT2OZ5UZnFAZPhpMGuWF7DIuy+x+XdRfJE8h/eJh3BYGb58iIOOajOcwUkOVpgN/kXPU/mN3CN0WT+FITBKmiEGRjIHu5h3IbdFEedFyOvdym1eaSR1Ek7os3Xa+M8rYrpkIj6fgOB0FOQ2FJquHJCdQm7qYSzVK/5jf/h3xBxnGq/iBP2+tdXbfPWXy7Bcys/J9E3GGJSHMH6kapK23Gt+JnDYglQgUIw2aE4crwyvJT/BK7mkGy2hx8RimYNNKrhXs3xAgMBAAECggEBAMWeoSZEGRGG/uqqZQuNnYX8yafMW31mrah6lPlbkefqWDda8UJ9O2N6kJo4zIBB6Qox2CAu6hRxWeGz4tL+tdhJ7ZeV3+kgmxB/0g0d66guG7gnQEQZD4XvUP5aUCq4zdSOknC/1770nNAnN8eKh9bmAoQ7WpBmnhM+kohe+p/P/o/D4EOH4MYCTJtCe4Ea6W8o0dVG5zfYR0SQ+AlSYgLYXi0dh72PmOuahNUvS3I53/7BW17YZ3xojmK+liP+wUuwZ6k/1lOhLrk3qH9rYY0Ud//hp6BBrF/T+fafkVa3fICjRdqVVX6JRPa8U3Ag86eudnnIFuiBFRNDOr6gRAECgYEA+U8fd1qociCdu2eYP/N6qjy15NEv3fYHckH9VhfdNJVIjDADwvmIbexJ2hnHSFx3tMkAWXaQ86JCB2w47qLRgryI66O3MFoCgWKrRb2r1GN4XNgcks2YOqgU86bWbHpzzwF8ZmKznMVAagJ6wYQ3t56CllzaDk9/mr12K6wAILECgYEA5eY6ytr6hwBF77qB9i6RjkjOUoKL+6T6ZLMDLBzrFs4bw0xZI/IIXJG2Ojvtz6sRWtUYMNvhIEiNfCy/tAVcRM6vqu6xwF5B+zL8F3t1FqNsPRgjMSCxPryUJiu79ASg8eKmLZEAQSi8FWi70Vg34/80tt7xKnnl8xlu9l6G0UECgYA0L0Gw5AMaUIVzss2FpVtpwud4C+lvFo6cdf+nQ7uDxDO5VFuVNlI+YBVdV8QE+4X7C4/NKipRNQeJMjgAi4g7S2eFm3E+57haiOK86GTNQjzxgjMI74wLyx8HmUaM0lznWbJGZCagjBFdn0M+uoRHJSDEhI8IK8/T/hB8N7aXIQKBgHudswk2e7UxgWloyM69tYhjP4WAKpLR381bsB39Iq9tfeIiYGACmVplAy4G4VVjr34+zLYg1MlOGb2mNiIvK7DXzf9EP5GnXSUcAg8CVDArCV1EaE/XO1b3gVWQ+Iw6HOxTKXWg3RksNQ3x9eOX4t2WcRrPf6+OQgXYLCEPLMDBAoGAQZxHWuzGCW7fEUa8W6/CpDClq2x3nxN/7n8Geadqg7wTpvO3BYiJDsx+9KK/rwU/gEva95FTs0yNXh9vGa6Mt+7SsHJHvj4eVC9lKgTeaKZ32ySbXYeoPkQKzyPALVk7KJBl59Nrsb7/oU0E7fC81mZX72YVW/BTLtsRnyRpd0c="
	AliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwCOf6ASz1XTiswHCN+uLmwagdgoh0D6aEbztuFqRgmi5xi3S+JPWOvNJVpaNntsOGsVbH9KPMZR1pQsYZLJ52gD2+PUzVcPrCltTkBlR2QXamvSjwFCYRgOtbNr4FGtbgYE83Fh0TC1IoLGdlvpKrQDdItNi0nYC/yDow3xroRQR1Vplhz3F/DHDJ5efyZ/yuOUG7p0oc3s3RoUeH2KW9sIcSewN+0mW5NLgXSz/U1yIpSC5PDI1LDBZP3XEl3BfPQ+D5jYGoxC46KuZBq8ii3nKI2TUH4hnYtRLXw5o+CPABEVttywMYWWNpZ15K4lJtqPyQAsZfaV2aInGoCSpDwIDAQAB"
	IsProduction = false
	ProductCode  = "FAST_INSTANT_TRADE_PAY"
	Host         = "http://zbiy92.natappfree.cc"
	NotifyURL    = "https://" + Host + "/notify"
	ReturnURL    = "http://127.0.0.1:8999/return"
)

func GoPayV() {
	xlog.Info("GoPay Version: ", gopay.Version)

}

func PayUrl() {

	appID := "9021000123613685"
	privateKey := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDf4/19hCVbrsShf/zXCU+/5OEQbEd4YLyOe8BYBB8L+xjfzQ6GuylME/fTQYEpdIYcHU6OfKhgmLlsQllT2OZ5UZnFAZPhpMGuWF7DIuy+x+XdRfJE8h/eJh3BYGb58iIOOajOcwUkOVpgN/kXPU/mN3CN0WT+FITBKmiEGRjIHu5h3IbdFEedFyOvdym1eaSR1Ek7os3Xa+M8rYrpkIj6fgOB0FOQ2FJquHJCdQm7qYSzVK/5jf/h3xBxnGq/iBP2+tdXbfPWXy7Bcys/J9E3GGJSHMH6kapK23Gt+JnDYglQgUIw2aE4crwyvJT/BK7mkGy2hx8RimYNNKrhXs3xAgMBAAECggEBAMWeoSZEGRGG/uqqZQuNnYX8yafMW31mrah6lPlbkefqWDda8UJ9O2N6kJo4zIBB6Qox2CAu6hRxWeGz4tL+tdhJ7ZeV3+kgmxB/0g0d66guG7gnQEQZD4XvUP5aUCq4zdSOknC/1770nNAnN8eKh9bmAoQ7WpBmnhM+kohe+p/P/o/D4EOH4MYCTJtCe4Ea6W8o0dVG5zfYR0SQ+AlSYgLYXi0dh72PmOuahNUvS3I53/7BW17YZ3xojmK+liP+wUuwZ6k/1lOhLrk3qH9rYY0Ud//hp6BBrF/T+fafkVa3fICjRdqVVX6JRPa8U3Ag86eudnnIFuiBFRNDOr6gRAECgYEA+U8fd1qociCdu2eYP/N6qjy15NEv3fYHckH9VhfdNJVIjDADwvmIbexJ2hnHSFx3tMkAWXaQ86JCB2w47qLRgryI66O3MFoCgWKrRb2r1GN4XNgcks2YOqgU86bWbHpzzwF8ZmKznMVAagJ6wYQ3t56CllzaDk9/mr12K6wAILECgYEA5eY6ytr6hwBF77qB9i6RjkjOUoKL+6T6ZLMDLBzrFs4bw0xZI/IIXJG2Ojvtz6sRWtUYMNvhIEiNfCy/tAVcRM6vqu6xwF5B+zL8F3t1FqNsPRgjMSCxPryUJiu79ASg8eKmLZEAQSi8FWi70Vg34/80tt7xKnnl8xlu9l6G0UECgYA0L0Gw5AMaUIVzss2FpVtpwud4C+lvFo6cdf+nQ7uDxDO5VFuVNlI+YBVdV8QE+4X7C4/NKipRNQeJMjgAi4g7S2eFm3E+57haiOK86GTNQjzxgjMI74wLyx8HmUaM0lznWbJGZCagjBFdn0M+uoRHJSDEhI8IK8/T/hB8N7aXIQKBgHudswk2e7UxgWloyM69tYhjP4WAKpLR381bsB39Iq9tfeIiYGACmVplAy4G4VVjr34+zLYg1MlOGb2mNiIvK7DXzf9EP5GnXSUcAg8CVDArCV1EaE/XO1b3gVWQ+Iw6HOxTKXWg3RksNQ3x9eOX4t2WcRrPf6+OQgXYLCEPLMDBAoGAQZxHWuzGCW7fEUa8W6/CpDClq2x3nxN/7n8Geadqg7wTpvO3BYiJDsx+9KK/rwU/gEva95FTs0yNXh9vGa6Mt+7SsHJHvj4eVC9lKgTeaKZ32ySbXYeoPkQKzyPALVk7KJBl59Nrsb7/oU0E7fC81mZX72YVW/BTLtsRnyRpd0c="
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwCOf6ASz1XTiswHCN+uLmwagdgoh0D6aEbztuFqRgmi5xi3S+JPWOvNJVpaNntsOGsVbH9KPMZR1pQsYZLJ52gD2+PUzVcPrCltTkBlR2QXamvSjwFCYRgOtbNr4FGtbgYE83Fh0TC1IoLGdlvpKrQDdItNi0nYC/yDow3xroRQR1Vplhz3F/DHDJ5efyZ/yuOUG7p0oc3s3RoUeH2KW9sIcSewN+0mW5NLgXSz/U1yIpSC5PDI1LDBZP3XEl3BfPQ+D5jYGoxC46KuZBq8ii3nKI2TUH4hnYtRLXw5o+CPABEVttywMYWWNpZ15K4lJtqPyQAsZfaV2aInGoCSpDwIDAQAB"

	var client, err = alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}

	var p = alipay.TradePagePay{}
	// page支付方式使用
	p.NotifyURL = "http://zbiy92.natappfree.cc/callback" // 支付结果回调的url，注意内网穿透问题
	p.ReturnURL = "http://zbiy92.natappfree.cc/retururl" // 支付成功后倒计时结束跳转的页面
	p.Subject = "标题"
	p.OutTradeNo = "sn12389342479" //传递一个唯一单号
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY" // page支付必须使用这个配置

	url, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}

	//fmt.Println(url.String())

	fmt.Println("支付地址：", url)
}
