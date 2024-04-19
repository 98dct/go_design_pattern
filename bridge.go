package bridge

import "fmt"

/**
  当不同的事物被联系到一起时，可以更换它们其中的任意一个不受影响。
  在上面的例子中，电器是一个抽象类，灯泡和风扇是它的具体实现类，
  开关是另一个抽象类，，圆形开关和方形开关是它的具体实现类，两个抽象类通过桥接的形式连接在一起（线路），
  在这种情况下，我们可以替换抽象类中的任意一个而让整体系统不受影响

  通俗的理解是：两种抽象a1, a2，这两种抽象之间彼此有关联，需要协同工作，每种抽象又有多重实现n1, n2，
  所以一种抽象的实现可以和另一种抽象的实现互相组合 n1 * n2中组合，设计的时候可以把后工作的抽象放到先
  工作的抽象里面
*/

type Appliances interface {
	Run()
	Name() string
}

//灯泡
type Bulb struct {
}

func (*Bulb) Name() string {
	return "Bulb"
}

func (*Bulb) Run() {
	fmt.Println("I am shining")
}

//风扇
type Fan struct {
}

func (*Fan) Name() string {
	return "Fan"
}

func (*Fan) Run() {
	fmt.Println("I am spining")
}

//开关接口
type Switch interface {
	TurnOn()
	TurnOff()
}

//圆形开关
type CircularSwitch struct {
	a Appliances
}

func (c *CircularSwitch) TurnOn() {
	fmt.Printf("turn on the CircularSwitch, %s will run\n", c.a.Name())
	c.a.Run()
}

func (c *CircularSwitch) TurnOff() {
	fmt.Printf("turn off the CircularSwitch, %s will stop\n", c.a.Name())
}

func NewCircularSwitch(a Appliances) *CircularSwitch {
	return &CircularSwitch{a: a}
}

//方形开关
type SquareSwitch struct {
	a Appliances
}

func (c *SquareSwitch) TurnOn() {
	fmt.Printf("turn on the SquareSwitch, %s will run\n", c.a.Name())
	c.a.Run()
}

func (c *SquareSwitch) TurnOff() {
	fmt.Printf("turn off the SquareSwitch, %s will stop\n", c.a.Name())
}

func NewSquareSwitch(a Appliances) *SquareSwitch {
	return &SquareSwitch{a: a}
}
