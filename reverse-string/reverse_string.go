package reverse

func Reverse(input string) string {
	rev := make([]byte, len(input))
	prevPos, revPos := 0, len(input)
	for pos := range input {
		revPos -= pos - prevPos
		copy(rev[revPos:], input[prevPos:pos])
		prevPos = pos
	}
	copy(rev[0:], input[prevPos:])
	return string(rev)
}
