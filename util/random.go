package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// init creates a seed for time calculations
func init() {
	rand.Seed(time.Now().Unix())
}

// RandomInt provides a random int64 value
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString provides a random non-numeric string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomSentence creates a sentence 
func RandomSentence(n, o int) string {
	sentence := []string{}
	for i := 0; i < n; i++ {
		word := RandomString(o)
		sentence = append(sentence, word)
	}
	res := strings.Join(sentence, " ")
	return res
}

// RandomTitle creates a title 
func RandomTitle() string {
	return RandomString(6)
}

// RandomDescription creates a description 
func RandomDescription() string {
	return RandomSentence(6, 4)
}

// RandomCookingTime creates a cooking time 
func RandomCookingTime() string {
	timeLength := RandomInt(0, 60)
	timeUnits := []string{"seconds", "minutes", "hours"}
	n := len(timeUnits)
	timeUnit := timeUnits[rand.Intn(n)]

	return strconv.Itoa(int(timeLength)) + " " + timeUnit
}

// RandomIngredients creates a slice of ingredients 
func RandomIngredients() []string {
	ingredients := []string{}

	for i := 0; i < 5; i++ {
		ing := RandomString(4)
		ingredients = append(ingredients, ing)
	}
	return ingredients
}

// RandomInstructions creates instructions  
func RandomInstructions() string {
	return RandomSentence(15, 4)
}

