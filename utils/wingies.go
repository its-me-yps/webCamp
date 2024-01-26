package utils

import (
	"errors"
	model "wingiesOrNot/models"
)

func WingiesOrNot(id1 string, id2 string, raw model.Students) (bool, error) {
	var student1, student2 model.Student
	for _, student := range raw {
		if student.Id == id1 {
			student1 = student
		} else if student.Id == id2 {
			student2 = student
		}
	}

	if student1.Id == "" || student2.Id == "" {
		if student1.Id != "" {
			return false, errors.New("id2 not found")
		}
		if student2.Id != "" {
			return false, errors.New("id1 not found")
		}

		return false, errors.New("Both ids not found")
	}

	if student1.Room[0:3] == student2.Room[0:3] {
		return true, nil
	}

	return false, nil
}
