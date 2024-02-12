package requests

import bt "github.com/amitramachandran/train-grpc/boardTrain"

func AddRouteRequest() *bt.RouteRequest {
	routeReq := &bt.RouteRequest{
		RouteID:     1,
		Route:       &bt.Transit{From: "france", To: "italy"},
		RouteCharge: &bt.Price{Amount: 20.00, Currency: "dollars"},
	}
	return routeReq
}

func GetTicketRequest() *bt.TicketRequest {
	purchaseReq := &bt.TicketRequest{
		RouteID: 1,
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
