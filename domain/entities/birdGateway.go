package entities

type birdInterfce interface {
	 list() []Bird
	 findById() (Bird, error)
	 add() Bird
	 delete() Bird
}