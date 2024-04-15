package observer

import "testing"

func TestObserve(t *testing.T) {
	book := NewGoods("book")
	observer1 := &Persion{Id: "observer1@qq.com"}
	observer2 := &Persion{
		Id: "observer2@qq.com",
	}

	book.Register(observer1)
	book.Register(observer2)
	book.updateAvailability()
}
