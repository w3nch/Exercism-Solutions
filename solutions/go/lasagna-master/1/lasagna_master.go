package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layer_minutes int) int {
    if layer_minutes == 0 {
        layer_minutes = 2
    }
    return len(layers) * layer_minutes
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodles := 0
    sauce := 0.0

    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        }
        if layer == "sauce" {
            sauce += 0.2
        }
    }

    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    secret := friendList[len(friendList)-1]
    myList[len(myList)-1] = secret
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portions int) []float64 {
    scale := []float64{}
    factor := float64(portions) / 2

    for _, q := range amount {
        scale = append(scale, q*factor)
    }

    return scale
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
