package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsPatna   bool     `json:"is_patna" pg:"is_patna"`
	Color     string   `json:"color"  pg:"color"`
}

// FindAllCats Получить список собак.
func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

// CreateCat Создать собаку.
func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// FindCatById Получить собаки по id.
func FindDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// DeleteCatById Удалить собаку по id.
func DeleteDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.IsPatna = dog.IsPatna
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldDog
}
