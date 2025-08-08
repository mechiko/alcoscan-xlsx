package utility

// var name, link string
// unpack(strings.Split("foo:bar", ":"), &name, &link)

func Unpack(s []string, vars ...*string) {
	for i, str := range s {
		// защита от паники по индексу
		if i >= len(vars) {
			return
		}
		*vars[i] = str
	}
}
