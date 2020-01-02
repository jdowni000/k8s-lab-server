package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func main() {
	http.HandleFunc("/", simulateDog)
	log.Println(http.ListenAndServe(":8000", nil))
}

func simulateDog(res http.ResponseWriter, req *http.Request) {

	// make a new dog
	dog := NewDog("Ace", res)
	res.Write([]byte("Bruce Wayne goes to the pet store and comes home with a dog named " + dog.Name + "<br>\n"))

	res.Write([]byte("<html>"))

	// begin simulating dog
	for {

		// let the dog die when its too old
		dog.Birthday()
		if dog.Age > 22 {
			break
		}

		// if 100 hungry, the dog has to eat
		if dog.Hunger >= 100 {
			res.Write([]byte(dog.Name + " is incredibly hungry and stops to eat.<br>\n"))
			dog.Eat()
			continue
		}

		// if 100 tired, the dog has to sleep
		if dog.Tiredness >= 100 {
			res.Write([]byte(dog.Name + " is incredibly tired and stops to sleep.<br>\n"))
			dog.Sleep()
			continue
		}

		// pick a random thing to do and do it
		choiceNumber := rand.Intn(7) // 7 being the number of things the dog can possibly do
		switch choiceNumber + 1 {
		case 1:
			dog.Bark()
		case 2:
			dog.Run()
		case 3:
			dog.Eat()
		case 4:
			dog.Sleep()
		case 5:
			dog.Play()
		case 6:
			dog.ChaseBird()
		case 7:
			dog.StopsCrime()
		}
	}

	res.Write([]byte(dog.Name + " had a great life, he came, he saw, he conquered, now he is dead and Bruce is sad."))
}
