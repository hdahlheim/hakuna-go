package lib

func Ternary(cond bool, truthy interface{}, falsy interface{}) interface{} {
	if cond {
		return truthy
	} else {
		return falsy
	}
}
