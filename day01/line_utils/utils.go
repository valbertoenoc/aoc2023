package line_utils

func InsertIntoString(index int, old, input string) string {
	return old[:index] + input + old[index:]
}
