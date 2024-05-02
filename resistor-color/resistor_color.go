package resistorcolor

var colorCode = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{}
	for color := range colorCode {
		colors = append(colors, color)
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colorCode[color]
}
