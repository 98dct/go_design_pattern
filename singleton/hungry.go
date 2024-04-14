package singleton

type Hungry struct {
}

var InsHunry *Hungry = &Hungry{}

func GetInsHungry() *Hungry {
	return InsHunry
}
