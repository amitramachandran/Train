package handlers

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"

	bt "github.com/amitramachandran/train-grpc/boardTrain"
)

func (s *BoardTrainServer) assignSeat(email string) (bt.Seat, error) {
	var section string
	var seatnumber int

	if strings.Contains(s.nextSeat, "A") {
		section = "B"
	} else {
		section = "A"
	}

	if s.nextSeat == "" {
		s.nextSeat = "A1"
		return bt.Seat{Section: section, Number: 1}, nil
	}

	if s.nextSeat == "B10" {
		return bt.Seat{}, errors.New("seats are not available")
	}

	return bt.Seat{Section: section, Number: int32(seatnumber) + 1}, nil

}

func splitSeatNumberFromSection(in string) (string, int, error) {
	section := in[:1]
	seatNumber, err := strconv.Atoi(in[1:])
	if err != nil {
		log.Fatalf("Could not convert the string to integer for seatnumber %v", err)
	}
	return section, seatNumber, nil
}

func checkSectionValid(in string) bool {
	section := []rune(in[:1])
	seatnumb := []rune(in[1:])

	switch len(in) {
	case 2:
		if unicode.IsLetter(section[0]) && unicode.IsNumber(seatnumb[0]) {
			return true
		}
	case 3:
		if unicode.IsLetter(section[0]) && unicode.IsNumber(seatnumb[0]) && unicode.IsNumber(seatnumb[1]) {
			return true
		}
	default:
		return false
	}
	return false

}
