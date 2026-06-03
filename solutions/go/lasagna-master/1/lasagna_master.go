package lasagnamaster

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }

    return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
    noodles := 0
    var sauce float64

    for _, v := range layers {
        switch v {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }

    return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaled := make([]float64, len(quantities))

    factor := float64(portions) / 2

    for i, q := range quantities {
        scaled[i] = q * factor
    }

    return scaled
}