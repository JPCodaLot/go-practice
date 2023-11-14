package main

import (
	"fmt"
)

type Gender uint

const (
	Male Gender = iota
	Female
	Neuter
)

type Number bool

const (
	Singular Number = true
	Plural   Number = false
)

type Case uint

const (
	// Subject, Predicate Nominative, or Predicate Adjective
	Nominative Case = iota
	// Possessive Adjective
	Genitive
	// Indirect Object
	Dative
	// Deirect Object or Object of Preposition
	Accusative
	Vocative
	// Object of Preposition (instument by which the action is accomplished)
	Ablative
)

type Ending struct {
	Runes  string
	Gender Gender
	Number Number
	Case   Case
}

var endings = [...]Ending{
	// 1st Declension
	{"a", Female, Singular, Nominative},
	{"ae", Female, Plural, Nominative},
	{"ae", Female, Singular, Genitive},
	{"arum", Female, Plural, Genitive},
	{"ae", Female, Singular, Dative},
	{"is", Female, Plural, Dative},
	{"am", Female, Singular, Accusative},
	{"is", Female, Plural, Accusative},
	{"a", Female, Singular, Ablative},
	{"is", Female, Plural, Ablative},
	// 2nd Declention
	{"us", Male, Singular, Nominative},
	{"i", Male, Plural, Nominative},
	{"i", Male, Singular, Genitive},
	{"orum", Male, Plural, Genitive},
	{"o", Male, Singular, Dative},
	{"is", Male, Plural, Dative},
	{"um", Male, Singular, Accusative},
	{"os", Male, Plural, Accusative},
	{"a", Male, Singular, Ablative},
	{"is", Male, Plural, Ablative},
}

type wordType uint

const (
	_ wordType = iota
	Noun
	Adjective
	Verb
	Adverb
)

func (w wordType) String() string {
	switch w {
	case 1:
		return "noun"
	case 2:
		return "adjective"
	case 3:
		return "verb"
	case 4:
		return "adverb"
	default:
		return "unknown"
	}
}

var wordBank = map[string]wordType{
	"prophet": Noun,
	"su":      Verb,
	"vir":     Noun,
	"sanct":   Adjective,
}

func main() {
	var word string
	for {
		fmt.Printf("Latin Word: ")
		fmt.Scanln(&word)

		wordType := findType(word)
		fmt.Printf("%v\n", wordType)

		ok, ending := findEnding(word)
		if ok != true {
			fmt.Println("Unreconized latin ending. Check your spelling.")
			continue
		}
		fmt.Printf("%+v\n", ending)
	}
}

func findType(word string) wordType {
	for base, v := range wordBank {
		if len(word) < len(base) {
			continue
		}
		if word[:len(base)] == base {
			return v
		}
	}
	return 0
}

func findEnding(word string) (bool, *Ending) {
	for _, ending := range endings {
		var size = len(word)
		if word[size-len(ending.Runes):size] == ending.Runes {
			return true, &ending
		}
	}
	return false, nil
}
