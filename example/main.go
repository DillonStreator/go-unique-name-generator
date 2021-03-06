package main

import (
	"fmt"
	"strings"

	ung "github.com/DillonStreator/go-unique-name-generator"
	"github.com/DillonStreator/go-unique-name-generator/dictionaries"
)

func main() {
	defaultGenerator := ung.NewUniqueNameGenerator(ung.UNGOpts{})
	generator1 := ung.NewUniqueNameGenerator(ung.UNGOpts{
		Dictionaries: [][]string{
			dictionaries.Adjectives,
			dictionaries.Adjectives,
			dictionaries.Colors,
			dictionaries.Animals,
		},
		Separator: ".",
		Style:     ung.Upper,
	})
	generator2 := ung.NewUniqueNameGenerator(ung.UNGOpts{
		Dictionaries: [][]string{
			dictionaries.Colors,
			dictionaries.Animals,
			dictionaries.Names,
		},
		Separator: "-",
		Style:     ung.Capital,
	})
	generator3 := ung.NewUniqueNameGenerator(ung.UNGOpts{
		Dictionaries: [][]string{
			dictionaries.Colors,
			dictionaries.Adjectives,
			dictionaries.Drinks,
		},
		Sanitizer: func(str string) string {
			return strings.Replace(str, " ", "", -1)
		},
	})

	fmt.Printf("defaultGenerator possible unique names: %d\n", defaultGenerator.UniquenessCount())
	fmt.Printf("defaultGenerator name: %s\n", defaultGenerator.Generate())
	fmt.Printf("generator1 possible unique names: %d\n", generator1.UniquenessCount())
	fmt.Printf("generator1 name: %s\n", generator1.Generate())
	fmt.Printf("generator2 possible unique names: %d\n", generator2.UniquenessCount())
	fmt.Printf("generator2 name: %s\n", generator2.Generate())

	fmt.Printf("generator3 name: %s\n", generator3.Generate())
}
