package puppy

import (
	"fmt"

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

func From10() {
	fmt.Println("I'm from v1.0.0")
}
