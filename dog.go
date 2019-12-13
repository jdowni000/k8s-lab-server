package main

import "io"

// Dog represents a four legged friend
type Dog struct {
	Name      string
	Hunger    int
	Tiredness int
	Age       int
}

// NewDog creates a new dog and returns a pointer to it
func NewDog(name string) *Dog {
	return &Dog{
		Name:      name,
		Hunger:    0,
		Tiredness: 0,
	}
}

// Birthday makes the dog get older
func (d *Dog) Birthday() {
	d.Age++
}

// Sleep makes the dog take a rest. resets tiredness to 0
func (d *Dog) Sleep(res io.Writer) {
	res.Write([]byte(d.Name + " sleeps...\n"))
	d.Tiredness = 0
}

// Bark makes the dog bark
func (d *Dog) Bark(res io.Writer) {
	res.Write([]byte(d.Name + " goes BARK BARK BARK\n"))
	d.Tiredness = d.Tiredness + 20
}

// Eat makes the dog eat. returns hunger to 0
func (d *Dog) Eat(res io.Writer) {
	res.Write([]byte(d.Name + " eats some food\n"))
	d.Hunger = 0
}

// Play makes the dog play
func (d *Dog) Play(res io.Writer) {
	res.Write([]byte(d.Name + " plays with toys\n"))
	d.Tiredness = d.Tiredness + 30
}

// Run makes the dog run around for awhile
func (d *Dog) Run(res io.Writer) {
	res.Write([]byte(d.Name + " runs around\n"))
	d.Tiredness = d.Tiredness + 40
}

//ChaseBird makes the dog chase a bird
func (d *Dog) ChaseBird(res io.Writer) {
	res.Write([]byte(d.Name + " chases a bird\n"))
	d.Tiredness = d.Tiredness + 50

}

//StopsCrime makes Ace stop a crime
func (d *Dog) StopsCrime(res io.Writer) {
	res.Write([]byte(d.Name + " stops a crime\n"))
	d.Hunger = d.Hunger + 50
}
