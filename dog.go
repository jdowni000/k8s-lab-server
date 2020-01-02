package main

import (
	"io"
)

// Dog represents a four legged friend
type Dog struct {
	Name      string
	Hunger    int
	Tiredness int
	Age       int
	Output    io.Writer
}

// NewDog creates a new dog and returns a pointer to it
func NewDog(name string, output io.Writer) *Dog {
	return &Dog{
		Name:      name,
		Hunger:    0,
		Tiredness: 0,
		Output:    output,
	}
}

// Birthday makes the dog get older
func (d *Dog) Birthday() {
	d.Age++
}

// Sleep makes the dog take a rest. resets tiredness to 0
func (d *Dog) Sleep() {
	d.Output.Write([]byte(d.Name + " sleeps...<br>\n"))
	d.Tiredness = 0
}

// Bark makes the dog bark
func (d *Dog) Bark() {
	d.Output.Write([]byte(d.Name + " goes BARK BARK BARK<br>\n"))
	d.Tiredness = d.Tiredness + 20
}

// Eat makes the dog eat. returns hunger to 0
func (d *Dog) Eat() {
	d.Output.Write([]byte(d.Name + " eats some food<br>\n"))
	d.Hunger = 0
}

// Play makes the dog play
func (d *Dog) Play() {
	d.Output.Write([]byte(d.Name + " plays with toys<br>\n"))
	d.Tiredness = d.Tiredness + 30
}

// Run makes the dog run around for awhile
func (d *Dog) Run() {
	d.Output.Write([]byte(d.Name + " runs around<br>\n"))
	d.Tiredness = d.Tiredness + 40
}

//ChaseBird makes the dog chase a bird
func (d *Dog) ChaseBird() {
	d.Output.Write([]byte(d.Name + " chases a bird<br>\n"))
	d.Tiredness = d.Tiredness + 50

}

//StopsCrime makes Ace stop a crime
func (d *Dog) StopsCrime() {
	d.Output.Write([]byte(d.Name + " stops a crime<br>\n"))
	d.Hunger = d.Hunger + 50
}
