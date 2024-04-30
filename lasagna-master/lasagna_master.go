package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer < 1 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += .2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	mult := float64(portions) / 2
	var scaled = make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaled[i] = quantities[i] * mult
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
