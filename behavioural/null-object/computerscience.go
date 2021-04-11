package main

type computerscience struct {
	numberOfProfessors int
}

func (c *computerscience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerscience) getName() string {
	return "computerscience"
}
