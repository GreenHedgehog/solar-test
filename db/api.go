package db

import "github.com/GreenHedgehog/solar-test/models"

func executor(query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	return err
}

// Add writes new vacancy to db
func Add(v *models.Vacancy) error {
	stmt, err := db.Prepare(qAdd)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(v.Name, v.Salary, v.Experience, v.Place).Scan(&v.ID)
}

// Delete removes vacancy with given id from db
func Delete(id int) error {
	return executor(qDelete, id)
}

// Get returns info for vacancy with given id
func Get(id int) (v models.Vacancy, err error) {
	err = db.QueryRow(qGet, id).Scan(&v.ID, &v.Name, &v.Salary, &v.Experience, &v.Place)
	return
}

// GetAll returns all vacancy sorted by name in json
func GetAll() (data []byte, err error) {
	err = db.QueryRow(qGetAll).Scan(&data)
	return
}
