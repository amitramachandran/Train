package requests

import bt "github.com/amitramachandran/train-grpc/boardTrain"

func GetTicketRequest() *bt.TicketRequest {
	purchaseReq := &bt.TicketRequest{
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
	return purchaseReq
}

func GetReceiptRequest() *bt.UserRequest {
	return &bt.UserRequest{
		EmailAddress: "amit@gmail.com",
	}
}

func GetSectionRequest() *bt.SectionRequest {
	return &bt.SectionRequest{
		Section: "A",
	}
}

func RemoveUserRequest() *bt.UserRequest {
	return &bt.UserRequest{
		EmailAddress: "amit@gmail.com",
	}
}

func ModifySeatRequest() *bt.ModifySeatRequest {
	return &bt.ModifySeatRequest{
		EmailAddress: "amit@gmail.com",
		NewSeat:      "B2",
	}
}
