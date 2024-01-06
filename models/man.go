package models

type Man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	alive     bool
}

func (man *Man) Breathe() {
	man.breathing = true
}

func (man *Man) Think() {
	man.thinking = true
}

func (man *Man) Eat() {
	man.eating = true
}

func (man *Man) Sex() string {
	return "male"
}

func (man *Man) IsAlive() bool {
	return true
}
