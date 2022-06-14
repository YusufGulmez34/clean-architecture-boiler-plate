package helpers

func MethodName(name string) string {
	n := []byte(name)

	switch n[len(n)-1] {
	case 'y':
		n[len(n)-1] = 'i'
		return string(n) + "es"
	case 's':
		return name + "es"
	default:
		return name + "s"
	}
}
