package test_service

import (
	"github.com/IanVzs/Snowflakes/models"
)

type TestData struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *TestData) GetAll() ([]models.Test, error) {
	model1 := models.Test{
		Name:      "test_1",
		State:     1,
		CreatedBy: "code",
	}
	model2 := models.Test{
		Name:      "test_2",
		State:     2,
		CreatedBy: "code",
	}
	tests := []models.Test{model1, model2}
	return tests, nil
}
