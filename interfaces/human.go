package interfaces

type Human interface { // Definition of functions that must be implemented by a type
	Breathe()
	Think()
	Eat()
	Sex() string
}
