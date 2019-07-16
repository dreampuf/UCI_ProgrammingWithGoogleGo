package main

import "fmt"

type AnimalI interface {
	Eaten() string
	Locomotion() string
	Spoken() string
}

type Animal struct {
	name, food, locomotion, noise string
}
func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}


var animalMap = map[string]Animal{
	"cow": {"cow", "grass", "walk", "moo"},
	"bird": {"bird", "worms", "fly", "peep"},
	"snake": {"snake", "mice", "slither", "hsss"},
}

func main() {
	for {
		action, tp, name := "", "", ""
		fmt.Print("> ")
		if _, err := fmt.Scanf("%s %s %s", &action, &tp, &name); err != nil {
			break
		}
		switch action {
		case "newanimal":
			if a, ok := animalMap[name]; !ok {
				fmt.Println("Doesn't exist animal type")
				break
			} else {
				animalMap[tp] = a
				fmt.Println()
			}
		case "query":
			if a, ok := animalMap[tp]; !ok {
				fmt.Println("Doesn't exist animal type")
				break
			} else {
				switch name {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}
		default:
			fmt.Println("Doesn't support command: ", action)
		}
	}
}
