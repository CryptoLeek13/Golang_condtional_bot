package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	failFace := '\U0001F602'
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step")
	} else {
		isHeistOn = false
		fmt.Println("Plan better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else {
		isHeistOn = false
		fmt.Println("The vault can't be opened")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Printf("Better luck Next time %c", failFace)
		case 1:
			isHeistOn = false
			fmt.Printf("You suck at this %c", failFace)
		case 2:
			isHeistOn = false
			fmt.Printf("You're getting colder.....%c", failFace)
		case 3:
			isHeistOn = false
			fmt.Printf("You're a buster %c", failFace)
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			amtStolen := 1000 + rand.Intn(1000000)
			fmt.Println(amtStolen)
		}
	}
	fmt.Println(isHeistOn)
}
