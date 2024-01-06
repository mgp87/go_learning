package models

type Woman struct {
	Man
	IsMother bool
}

func (man *Woman) Breathe() {
	man.breathing = true
}

func (man *Woman) Think() {
	man.thinking = true
}

func (man *Woman) Eat() {
	man.eating = true
}

func (man *Woman) Sex() string {
	return "female"
}
