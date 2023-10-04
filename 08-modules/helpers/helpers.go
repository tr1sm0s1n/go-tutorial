package helpers

import "errors"

type Student struct {
	Name, Course, Grade string
}

func ConvertMarks(marks uint) (string, error) {
	switch {
	case marks >= 80:
		return "A", nil
	case marks >= 60 && marks < 80:
		return "B", nil
	case marks >= 40 && marks < 60:
		return "C", nil
	case marks >= 0 && marks < 40:
		return "F", nil
	} 
	return "", errors.New("failed to convert")
}
