package main

import "fmt"

type validationError struct {
	Massage string
}

func (v *validationError) Error() string {
	return v.Massage

}

type notFoundError struct {
	Massage string
}

func (n *notFoundError) Error() string {
	return n.Massage
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"valdiation error"}

	}

	if id != "ilham" {
		return &validationError{"id not found"}

	}

	return nil
}

func main() {
	err := SaveData("ilham", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Error())
		} else {
			fmt.Println("unkwnown error:", err.Error())
		}

		// switch finalError := err.(type) {
		// case *validationError:
		// 	fmt.Println("validation error:", validationErr.Error())
		// case *notFoundError:
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// default:
		// 	fmt.Println("unkwnown error:", err.Error())

		// }

	} else {
		fmt.Println("id Benar")
	}
}
