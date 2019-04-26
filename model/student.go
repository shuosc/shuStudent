package model

import "shuStudent/infrastructure"

type Student struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Mail        string `json:"mail"`
	PhoneNumber string `json:"phone_number"`
}

func Get(id string) (Student, error) {
	result := Student{Id: id}
	row := infrastructure.DB.QueryRow(`
	SELECT name, mail, phoneNumber 
	FROM student
	WHERE id=$1;
	`, id)
	err := row.Scan(&result.Name, &result.Mail, &result.PhoneNumber)
	return result, err
}

func Put(student Student) {
	_, _ = infrastructure.DB.Exec(`
	INSERT INTO student(id, name, mail, phoneNumber) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (id) do
	UPDATE SET
		name=$2,
		mail=$3,
	    phoneNumber=$4;
	`, student.Id, student.Name, student.Mail, student.PhoneNumber)
}
