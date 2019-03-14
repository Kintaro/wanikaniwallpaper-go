package main

import "io/ioutil"
import "container/list"
import "unicode"

type Order struct {
	intToChar map[int]string
	charToInt map[string]int
	charToKanji map[string]*Kanji
}

func NewOrder(filename string) *Order {
	o := Order { make(map[int]string), make(map[string]int), make(map[string]*Kanji)  }
	o.Load(filename)
	return &o
}

func (o *Order) Load(filename string) {
	bytes, _ := ioutil.ReadFile(filename)
	content := string(bytes)

	index := 0
	for _, c := range content {
		if unicode.IsSpace(c) {
			continue
		}
		s := string(c)
		kanji := NewKanji(s)
		o.intToChar[index] = s
		o.charToInt[s] = index
		o.charToKanji[s] = kanji
		index++
	}
}

func (o *Order) Update(l *list.List) {
	for i := l.Front(); i != nil; i = i.Next() {
		kanji := i.Value.(Kanji)
		o.charToKanji[kanji.character] = &kanji
	}
}

func (o *Order) Position(character string) int {
	return o.charToInt[character]
}

func (o *Order) Character(position int) string {
	return o.intToChar[position]
}

func (o *Order) KanjiForCharacter(character string) *Kanji {
	return o.charToKanji[character]
}

func (o *Order) KanjiForPosition(position int) *Kanji {
	return o.charToKanji[o.intToChar[position]]
}

func (o *Order) Size() int {
	return len(o.charToKanji)
}
