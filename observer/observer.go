package observer

import (
	"fmt"
)

/**
  观察者模式：监听模式、事件订阅模式  在对象间定义一对多的依赖关系，当这个对象的状态发生改变时，
  所有依赖他的对象都会被通知并自动更新
  观察者至少有三个方法：添加观察者、移除观察者、通知观察者
  被观察者至少有一个方法：更新方法，更新当前的内容
  观察者模式根据其侧重的功能分位推模型和拉模型
*/

//被观察者接口
type IGoods interface {
	Register(persion IPersion)
	Remove(persion IPersion)
	Notify()
}

type Goods struct {
	Name     string
	InStock  bool
	persions []IPersion
}

func NewGoods(name string) *Goods {
	return &Goods{
		Name: name,
	}
}

func (g *Goods) updateAvailability() {
	fmt.Printf("Goods %s is now in stock\n", g.Name)
	g.InStock = true
	//通知观察者
	g.Notify()
}

//添加观察者
func (g *Goods) Register(persion IPersion) {
	g.persions = append(g.persions, persion)
}

//移除观察者
func (g *Goods) Remove(persion IPersion) {
	g.persions = g.removePersion(g.persions, persion)
}

func (g *Goods) removePersion(persions []IPersion, perison IPersion) []IPersion {
	l := len(persions)
	for i, persion := range persions {
		if perison.getId() == persion.getId() {
			persions[l-1], persions[i] = persions[i], persions[l-1]
			return persions[:l-1]
		}
	}
	return persions
}

//通知观察者
func (g *Goods) Notify() {
	for _, persion := range g.persions {
		persion.Update(g.Name)
	}
}

//观察者接口
type IPersion interface {
	Update(msg string)
	getId() string
}

type Persion struct {
	Id string
}

func (p *Persion) getId() string {
	return p.Id
}

func (p *Persion) Update(gName string) {
	fmt.Printf("Sending email to persion %s for Goods %s\n", p.Id, gName)
}
