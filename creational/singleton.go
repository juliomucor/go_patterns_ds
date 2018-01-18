package creational

// this implementation is not thread safe!
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// initizalize a pointer to a struct as nil
// but cannot initialize a strucuture to nil
var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
		// with the keyword new, we are creating a pointer
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
