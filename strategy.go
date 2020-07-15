package strategy
/*
题目
	使用 Go 实现一个 商场收银系统，要求:
		1、输入商品的单价，数量，输出总价
		2、商场会举行活动，给用户优惠，活动有两种:
			a.打折活动(打5，6，7，8，9折扣)
			b.满减活动(按照总价参与例如满700减200，满300减50) 此时这个收银系统也需要支持
要求
	用工厂or策略模式实现
*/

//收银总接口
 type ICash interface {
    Compute(price float64, amount float64) float64
 }

//不打折
type Nodiscount struct {
}

func (p *Nodiscount) Compute(price float64, amount float64) float64 {
	return price * amount
}


//打五折
type FiftyPercentOff struct {
}

func (p *FiftyPercentOff) Compute(price float64, amount float64) float64 {
	return price * amount *0.5
}

//打六折
type SixtyPercentOff struct {
}

func (p *SixtyPercentOff) Compute(price float64, amount float64) float64 {
	return price * amount *0.6
}
//打七折
type SeventyPercentOff struct {
}

func (p *SeventyPercentOff) Compute(price float64, amount float64) float64 {
	return price * amount *0.7
}
//打八折
type EightyPercentOff struct {
}

func (p *EightyPercentOff) Compute(price float64, amount float64) float64 {
	return price * amount *0.8
}
//打九折
type NinetyPercentOff struct {
}

func (p *NinetyPercentOff) Compute(price float64, amount float64) float64 {
	return price * amount *0.9
}


//满减活动
type FullReduction struct {
}

func (p *FullReduction) Compute(price float64, amount float64) float64 {
	TotalPrice:= price * amount
	if TotalPrice >700{
		TotalPrice-=200

	}else if TotalPrice <=700 && TotalPrice>300{
		TotalPrice-=500
	}
	return TotalPrice

}

//策略类
var cash ICash
type Context struct {
	A, B float64
}

func (p *Context) SetContext(o ICash) {
	cash = o
}

func (p *Context) Result() float64 {
	return cash.Compute(p.A, p.B)
}
