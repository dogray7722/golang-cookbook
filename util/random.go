package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomSentence(n, o int) string {
	sentence := []string{}
	for i := 0; i < n; i++ {
		word := RandomString(o)
		sentence = append(sentence, word)
	}
	res := strings.Join(sentence, " ")
	return res
}

func RandomTitle() string {
	return RandomString(6)
}

func RandomDescription() string {
	return RandomSentence(6, 4)
}

func RandomCookingTime() string {
	timeLength := RandomInt(0, 60)
	timeUnits := []string{"seconds", "minutes", "hours"}
	n := len(timeUnits)
	timeUnit := timeUnits[rand.Intn(n)]

	return strconv.Itoa(int(timeLength)) + " " + timeUnit
}

func RandomIngredients() []string {
	ingredients := []string{}

	for i := 0; i < 5; i++ {
		ing := RandomString(4)
		ingredients = append(ingredients, ing)
	}
	return ingredients
}

func RandomInstructions() string {
	return RandomSentence(15, 4)
}

