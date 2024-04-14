package template_pattern

import "fmt"

type Cook interface {
	fire()
	cook()
	outfire()
}

type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("点火")
}

func (CookMenu) cook() {

}

func (CookMenu) outfire() {
	fmt.Println("灭火")
}

func doCook(cook Cook) {
	cook.fire()
	cook.cook()
	cook.outfire()
}

type Xihongshi struct {
	CookMenu
}

func (*Xihongshi) cook() {
	fmt.Println("做西红柿")
}

type Chaojidan struct {
	CookMenu
}

func (*Chaojidan) cook() {
	fmt.Println("炒鸡蛋")
}
