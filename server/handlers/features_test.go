package handlers

import (
	"context"
	"fmt"
	"testing"

	bt "github.com/amitramachandran/train-grpc/boardTrain"
	"github.com/stretchr/testify/assert"
)

var mockserver = NewServer()

func TestPurchaseTicket(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mockTicketRequest := &bt.TicketRequest{
			Route: &bt.Transit{
				From: "France",
				To:   "Italy",
			},
			Passenger: &bt.User{
				Firstname:    "amit",
				Lastname:     "Kumar",
				EmailAddress: "amit@gmail.com",
				Id:           1,
			},
		}
		resp, _ := mockserver.PurchaseTicket(context.Background(), mockTicketRequest)
		assert.NotNil(t, resp)
	})

	t.Run("sad path when the section is completely booked", func(t *testing.T) {
		mockTicketRequest := &bt.TicketRequest{
			Route: &bt.Transit{
				From: "France",
				To:   "Germany",
			},
			Passenger: &bt.User{
				Firstname:    "amit",
				Lastname:     "Kumar",
				EmailAddress: "amit@gmail.com",
				Id:           2,
			},
		}
		mockserver.nextSeat = "B10"
		fmt.Printf("%v", mockserver)
		_, err := mockserver.PurchaseTicket(context.Background(), mockTicketRequest)
		assert.NotNil(t, err)
	})
}

func TestGetReceipt(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mockgetReceipt := &bt.UserRequest{
			EmailAddress: "amit@gmail.com",
		}
		mockserver.tickets[mockgetReceipt.EmailAddress] = &bt.TicketResponse{}
		resp, _ := mockserver.GetReceipt(context.Background(), mockgetReceipt)
		assert.NotNil(t, resp)
	})

	t.Run("sad path, get receipt for user not found", func(t *testing.T) {
		mockgetReceipt := &bt.UserRequest{
			EmailAddress: "john@gmail.com",
		}
		_, err := mockserver.GetReceipt(context.Background(), mockgetReceipt)
		assert.NotNil(t, err)
	})
}

func TestRemoveUser(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mockremoveuser := &bt.UserRequest{
			EmailAddress: "amit@gmail.com",
		}
		mockserver.tickets[mockremoveuser.EmailAddress] = &bt.TicketResponse{}
		mockserver.seats[mockremoveuser.EmailAddress] = &bt.Seat{}
		resp, _ := mockserver.RemoveUser(context.Background(), mockremoveuser)
		assert.NotNil(t, resp)
	})

	t.Run("sad path, get receipt for user not found", func(t *testing.T) {
		mockremoveuser := &bt.UserRequest{
			EmailAddress: "john@gmail.com",
		}
		_, err := mockserver.RemoveUser(context.Background(), mockremoveuser)
		assert.NotNil(t, err)
	})
}

func TestViewSeats(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mockViewSeat := &bt.SectionRequest{
			Section: "A",
		}
		mockserver.tickets[mockViewSeat.Section] = &bt.TicketResponse{}
		resp, _ := mockserver.ViewSeats(context.Background(), mockViewSeat)
		assert.NotNil(t, resp)
	})

	t.Run("sad path", func(t *testing.T) {
		mockViewSeat := &bt.SectionRequest{
			Section: "Z",
		}
		_, err := mockserver.ViewSeats(context.Background(), mockViewSeat)
		assert.NotNil(t, err)
	})
}

func TestModifySeat(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		mockModifySeat := &bt.ModifySeatRequest{
			EmailAddress: "amit@gmail.com",
			NewSeat:      "A5",
		}
		mockserver.seats[mockModifySeat.EmailAddress] = &bt.Seat{Section: "A", Number: 10}
		resp, _ := mockserver.ModifySeat(context.Background(), mockModifySeat)
		assert.NotNil(t, resp)
	})

	t.Run("Sad path - user not found", func(t *testing.T) {
		mockModifySeat := &bt.ModifySeatRequest{
			EmailAddress: "arumugam@gmail.com",
			NewSeat:      "A5",
		}
		//mockserver.seats[mockModifySeat.EmailAddress] = &bt.Seat{Section: "A", Number: 10}
		_, err := mockserver.ModifySeat(context.Background(), mockModifySeat)
		assert.NotNil(t, err)
	})

	t.Run("Sad path - seat number invalid", func(t *testing.T) {
		mockModifySeat := &bt.ModifySeatRequest{
			EmailAddress: "arumugam@gmail.com",
			NewSeat:      "A11",
		}
		//mockserver.seats[mockModifySeat.EmailAddress] = &bt.Seat{Section: "A", Number: 10}
		_, err := mockserver.ModifySeat(context.Background(), mockModifySeat)
		assert.NotNil(t, err)
	})

	t.Run("Sad path - section invalid", func(t *testing.T) {
		mockModifySeat := &bt.ModifySeatRequest{
			EmailAddress: "arumugam@gmail.com",
			NewSeat:      "Z5",
		}
		//mockserver.seats[mockModifySeat.EmailAddress] = &bt.Seat{Section: "A", Number: 10}
		_, err := mockserver.ModifySeat(context.Background(), mockModifySeat)
		assert.NotNil(t, err)
	})
}
