package main

import (
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func main() {
	http.HandleFunc("/", simulateDog)
	http.ListenAndServe(":8000", nil)
}

func simulateDog(res http.ResponseWriter, req *http.Request) {

	// make a new dog
	dog := NewDog("Ace")
	res.Write([]byte("Bruce Wayne goes to the pet store and comes home with a dog named " + dog.Name + "\n"))

	// begin simulating dog
	for {

		// let the dog die when its too old
		dog.Birthday()
		if dog.Age > 22 {
			break
		}

		// if 100 hungry, the dog has to eat
		if dog.Hunger >= 100 {
			res.Write([]byte(dog.Name + " is incredibly hungry and stops to eat.\n"))
			dog.Eat(res)
			continue
		}

		// if 100 tired, the dog has to sleep
		if dog.Tiredness >= 100 {
			res.Write([]byte(dog.Name + " is incredibly tired and stops to sleep.\n"))
			dog.Sleep(res)
			continue
		}

		// pick a random thing to do and do it
		choiceNumber := rand.Intn(7) // 7 being the number of things the dog can possibly do
		switch choiceNumber + 1 {
		case 1:
			dog.Bark(res)
		case 2:
			dog.Run(res)
		case 3:
			dog.Eat(res)
		case 4:
			dog.Sleep(res)
		case 5:
			dog.Play(res)
		case 6:
			dog.ChaseBird(res)
		case 7:
			dog.StopsCrime(res)
		}
	}

	res.Write([]byte(dog.Name + " had a great life, he came, he saw, he conquered, now he is dead and Bruce is sad."))
}
