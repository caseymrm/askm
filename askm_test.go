package askm

import "testing"

type person struct {
	name string
	age  int
}

func TestArbitraryStringNotInSimpleMap(t *testing.T) {
	intMap := make(map[string]int)
	count := 0
	for count < 50000 {
		key := ArbitraryKeyNotInMap(intMap)
		if _, ok := intMap[key]; ok {
			t.Errorf("Got int collision %s in %v", key, intMap)
		}
		intMap[key] = 1
		count++
	}
}

func TestArbitraryStringNotInStructMap(t *testing.T) {
	personMap := make(map[string]person)
	count := 0
	for count < 50 {
		key := ArbitraryKeyNotInMap(personMap)
		if _, ok := personMap[key]; ok {
			t.Errorf("Got int collision %s in %v", key, personMap)
		}
		personMap[key] = person{name: RandomString(3), age: 25}
		count++
	}
}

func TestArbitraryStringNotInPointerMap(t *testing.T) {
	personMap := make(map[string]*person)
	count := 0
	for count < 50 {
		key := ArbitraryKeyNotInMap(personMap)
		if _, ok := personMap[key]; ok {
			t.Errorf("Got int collision %s in %v", key, personMap)
		}
		personMap[key] = &person{name: RandomString(3), age: 25}
		count++
	}
}
