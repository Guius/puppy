package puppy

import (
	dog "github.com/Guius/Dog"
)

func Bark() string {
	return "Bark"
}

func Barks() string {
	return "Bark Bark Bark"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}