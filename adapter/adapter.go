package adapter

/**
  适配器模式用于将一种形态转换为另一种形态，使功能能够正常使用，达到适配的效果
  将一个类的接口转换为客户希望的另一个接口。
  解决的问题：系统需要使用现有的类，而此类的接口不符合系统当前的需要，主要解决一些现存的对象放到新环境中，而新环境要求的接口是现对象不能满足的
  解决方案：采用一个适配器对象，实现的是新的接口，实现新的接口，但是包含现有的类，在实现新接口时，实际调用的是现有的类的接口
*/

//typec接口
type TypeC interface {
	UseTypeC() string
}

//usb接口
type USB interface {
	UseUSB() string
}

type keyboard struct {
}

func NewKeyBoard() USB {
	return &keyboard{}
}

//keyboard实现了USB接口
func (a *keyboard) UseUSB() string {
	return "I use USB interface"
}

//适配器对象
type adapter struct {
	usb USB
}

func NewAdapter(keyboard USB) TypeC {
	return &adapter{usb: keyboard}
}

func (a *adapter) UseTypeC() string {
	return a.usb.UseUSB() + ", but now I use Type-C interface"
}
