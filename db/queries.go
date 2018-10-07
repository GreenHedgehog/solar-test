package db

const (
	qAdd    = "insert into vacancy (name, salary, experience, place) values ($1, $2, $3, $4) returning id;"
	qDelete = "delete from vacancy where id = $1"
	qGet    = "select * from vacancy where id = $1"
	qGetAll = "select * from get_all()"
)
