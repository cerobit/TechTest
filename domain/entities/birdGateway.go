package entities

type BirdInterfce interface {
	Add(Bird) error
	List() ( []Bird,error )
	FindById( int ) Bird
	Update(Bird) error
	DeleteById( int )  error
}