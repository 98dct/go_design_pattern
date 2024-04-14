package strategy

type IStrategy interface {
	Do(int, int) int
}

type add struct {
}

func (*add) Do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) Do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

func (operator *Operator) Execute(a, b int) int {
	return operator.strategy.Do(a, b)
}
