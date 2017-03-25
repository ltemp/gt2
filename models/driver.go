package models

type Driver struct {
	Id int
	Name string
	License_number string
}

func InsertDriver(driver Driver) (error) {
	stmt, err := DBCon.Prepare("INSERT INTO driver(id,name,License_number) VALUES($1,$2,$3);")
	_, err = stmt.Exec(driver.Id, driver.Name, driver.License_number)
	return err
}

func DeleteDriver(id int) (error, bool) {
	stmt, err := DBCon.Prepare("DELETE FROM driver where id = $1;")
	res, err := stmt.Exec(id)
	rows, _ := res.RowsAffected()
	return err, rows > 0
}

func UpdateDriver(driver Driver) (error, bool) {
	stmt, err := DBCon.Prepare("UPDATE driver SET name = $1, License_number=$2 WHERE id = $3")
	res, err := stmt.Exec(driver.Name, driver.License_number, driver.Id)
	rows, _ := res.RowsAffected()
	return err, rows > 0
}

func LoadDriver(id int) (Driver, error) {
	stmt, err := DBCon.Prepare("SELECT id, name FROM driver where id = $1")
	res := Driver{}
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Name)
	return res, err
}