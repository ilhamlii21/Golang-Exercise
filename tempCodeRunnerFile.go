if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Error())
		} else {
			fmt.Println("unkwnown error:", err.Error())
		}