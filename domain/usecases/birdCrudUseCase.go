package usecases

import (
	"LealTechTest/domain/entities"
)

type CrudUseCases struct {
	// Tech separation isolated
	BGateway entities.BirdInterfce
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
func (crud CrudUseCases) Delete(bird entities.Bird) (entities.Bird, error){
	err := crud.BGateway.DeleteById(bird.ID)
	return bird,err
}


