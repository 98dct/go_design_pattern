package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	processor := &BaseProcessor{}
	encryptProcessor := NewEncryptDecorator(processor)
	formatProcessor := NewFormatDecorator(encryptProcessor)

	text := "hello world"
	res := formatProcessor.process(text) //这里相当于一个装饰器的递归调用链
	fmt.Println(res)
}
