package domain

type VegetableRepositoryStub struct {
	vegetables []Vegetable
}

func (stub VegetableRepositoryStub) GetAll() ([]Vegetable, error){
	return stub.vegetables, nil;
}

func NewVegetableRepositoryStub() VegetableRepositoryStub {
	vegetables := []Vegetable{
		{1, "sawi", 4000},
		{2, "bayam", 2000},
		{3, "kangkung", 1000},
		{4, "kol", 5000},
		{5, "pare", 3000},
	}
	return VegetableRepositoryStub {vegetables}
}