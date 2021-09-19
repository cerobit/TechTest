package usecases

import (
	"LealTechTest/domain/entities"
)

type CrudUseCases struct {
	// Tech separation isolated
	BGateway entities.BirdGateway
}

// Add
func (crud CrudUseCases) Add(bird entities.Bird) (entities.Bird, error) {
	err := crud.BGateway.Add(bird)
	return bird,err
}

// Edit
func (crud CrudUseCases) Update(bird entities.Bird) (entities.Bird, error){
	err := crud.BGateway.Update(bird)
	return bird,err
}

//List
func (crud CrudUseCases) List() ([]entities.Bird, error){
	return crud.BGateway.List()
}


// Delete
func (crud CrudUseCases) Delete(bird entities.Bird) error {
	err := crud.BGateway.Delete( bird )
	return err
}


// FindById
func (crud CrudUseCases) FindById(i int ) (entities.Bird, error){
	bird, err := crud.BGateway.FindById(i)
	if err != nil {
		return bird, err
	}
	return bird,err
}

