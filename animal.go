package animal

type Animal interface {
	Speaks() string
}

// implementation of Animal
type Dog struct{}
func (a Dog) Speaks() string { return "woof" }