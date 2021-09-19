package entities

type BirdGateway interface {
	Add(Bird) error
	List() ( []Bird,error )
	FindById( int ) (Bird, error )
	Update(Bird) error
	Delete ( Bird )  error
}