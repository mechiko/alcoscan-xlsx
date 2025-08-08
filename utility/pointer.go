package utility

import "reflect"

func IsPointer(i interface{}) bool {
	if i == nil {
		// fmt.Printf("Interface is nil: %v\n", i)
		return false
	}
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		// fmt.Printf("Interface contains a pointer: %v (Type: %T)\n", i, i)
		// Additionally, you can check if the pointer itself is nil
		if val.IsNil() {
			// fmt.Printf("  The contained pointer is nil.\n")
			return false
		} else {
			// fmt.Printf("  The contained pointer is not nil.\n")
			return true
		}
	} else {
		// fmt.Printf("Interface does not contain a pointer: %v (Type: %T)\n", i, i)
		return false
	}
}
