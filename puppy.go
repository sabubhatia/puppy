package puppy

import "github.com/sabubhatia/dog"

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof! Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenHeGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenHeGrowsUp(Barks())
}

func From100() string {
	return "v1.0.0"
}

func From101() string {
	return "v1.0.1"
}

