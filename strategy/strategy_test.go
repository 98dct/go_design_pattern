package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}
	operator.setStrategy(&add{})
	res1 := operator.Execute(1, 2)
	fmt.Println(res1)

	operator.setStrategy(&reduce{})
	res2 := operator.Execute(1, 2)
	fmt.Println(res2)
}
