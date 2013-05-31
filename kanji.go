package main

import (
	"image/color"
)

const (
	statusApprentice 	= 1
	statusGuru			= 2
	statusMaster		= 3
	statusEnlightened	= 4
	statusBurned		= 5
	statusUnknown		= 6
)

type KanjiStatus int32

type KanjiStats struct {
	srs string	
}

func NewKanji(character string) *Kanji {
	return &Kanji{ character, statusUnknown } 
}

func (s *KanjiStats) Status() KanjiStatus {
	switch s.srs {
	case "apprentice": 
		return statusApprentice
	case "guru":
		return statusGuru
	case "master":
		return statusMaster
	case "enlighten":
		return statusEnlightened
	case "burned":
		return statusBurned
	}

	panic("Unknown status")
}

type Kanji struct {
	character string
	stats KanjiStatus
}

func (k *Kanji) Color() color.RGBA {
	switch k.stats {
		case statusUnknown:
			return color.RGBA{ 40, 40, 40, 255 }
		case statusApprentice:
			return color.RGBA{ 221, 0, 147, 255 }
		case statusGuru:
			return color.RGBA{ 136, 45, 158, 255 }
		case statusMaster:
			return color.RGBA{ 41, 77, 219, 255 }
		case statusEnlightened:
			return color.RGBA{ 0, 147, 221, 255 }
		case statusBurned:
			return color.RGBA{ 240, 240, 240, 255 }
		default:
			return color.RGBA{ 255, 0, 0, 255 }
	}

	return color.RGBA{ 0, 0, 0, 0 }
}
