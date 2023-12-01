package web

import (
	"fmt"
	"reflect"
)

func RequestJsonValidation(request interface{}) string {
	requestCamps := reflect.ValueOf(request)
	for i := 0; i < requestCamps.NumField(); i++ {
		campName := requestCamps.Type().Field(i).Name
		campValue := requestCamps.Field(i).Interface()
		campType := fmt.Sprint(reflect.TypeOf(campValue).Kind())
		switch campType {
		case "string":
			if campValue == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "int":
			if campValue.(int) <= 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "float64":
			if campValue.(float64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "boolean":
			if !campValue.(bool) {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}

		}
	}
	return ""
}