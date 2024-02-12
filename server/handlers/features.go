package handlers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	bt "github.com/amitramachandran/train-grpc/boardTrain"
)

type BoardTrainServer struct {
	bt.UnimplementedBoardingServer
	mu           sync.Mutex
	tickets      map[string]*bt.TicketResponse
	seats        map[string]*bt.Seat
	routes       map[int]*bt.Transit
	routeCharges map[int]*bt.Price
	nextSeat     string
	l            *log.Logger
}

func NewServer() *BoardTrainServer {
	return &BoardTrainServer{
		mu:           sync.Mutex{},
		tickets:      make(map[string]*bt.TicketResponse),
		seats:        make(map[string]*bt.Seat),
		routes:       make(map[int]*bt.Transit),
		routeCharges: make(map[int]*bt.Price),
		nextSeat:     "",
		l:            log.New(io.Discard, "TrainBooking", 1),
	}
}

func (s *BoardTrainServer) AddRoutes(ctx context.Context, req *bt.RouteRequest) (*bt.StatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.routes[int(req.RouteID)] = &bt.Transit{
		From: req.Route.From,
		To:   req.Route.To,
	}

	s.routeCharges[int(req.RouteID)] = &bt.Price{
		Amount:   req.RouteCharge.Amount,
		Currency: req.RouteCharge.Currency,
	}
	s.l.Printf("Route added %v", s.routes)
	return &bt.StatusResponse{Message: "Route added successfully"}, nil
}

func (s *BoardTrainServer) PurchaseTicket(ctx context.Context, req *bt.TicketRequest) (*bt.TicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//assign a seat
	seat, err := s.assignSeat(req.Passenger.EmailAddress)
	if err != nil {
		s.l.Printf("Cannot assign seat for passenger %s : %v", req.Passenger.Firstname, err)
		return &bt.TicketResponse{}, err
	}
	s.l.Printf("Seat assigned for %s : section %s number %d", req.Passenger.Firstname, seat.Section, seat.Number)

	if len(s.routes) == 0 {
		return &bt.TicketResponse{}, errors.New("no routes are added, add and purchase tickets for added routes")
	}

	passenger_route := s.routes[int(req.GetRouteID())]
	passenger_routeCharge := s.routeCharges[int(req.GetRouteID())]

	receipt := &bt.TicketResponse{
		From: passenger_route.From,
		To:   passenger_route.To,
		UserDetails: &bt.User{
			EmailAddress: req.Passenger.EmailAddress,
		},
		PriceDetails: &bt.Price{
			Amount:   passenger_routeCharge.Amount,
			Currency: passenger_routeCharge.Currency,
		},
	}

	s.tickets[req.Passenger.EmailAddress] = receipt
	s.seats[req.Passenger.EmailAddress] = &seat

	return receipt, nil
}

func (s *BoardTrainServer) GetReceipt(ctx context.Context, req *bt.UserRequest) (*bt.TicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	resp, err := s.tickets[req.EmailAddress]
	if !err {
		s.l.Printf("The provided email has no history of ticket purchase")
		return &bt.TicketResponse{}, errors.New("user not found")
	}
	s.l.Printf("Receipt received %v", resp)
	return resp, nil
}

func (s *BoardTrainServer) ViewSeats(ctx context.Context, req *bt.SectionRequest) (*bt.SeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var UserInSection []*bt.User

	if len(s.tickets) == 0 {
		return &bt.SeatResponse{}, errors.New("no seats have been booked yet")
	}

	for email, receipt := range s.tickets {
		seat, ok := s.seats[email]
		if ok && seat.Section == req.Section {
			UserInSection = append(UserInSection, receipt.UserDetails)
		}
	}

	return &bt.SeatResponse{
		Users: UserInSection,
	}, nil
}

func (s *BoardTrainServer) RemoveUser(ctx context.Context, req *bt.UserRequest) (*bt.StatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.seats[req.EmailAddress]
	if !ok {
		s.l.Printf("User not found %s", req.EmailAddress)
		return &bt.StatusResponse{}, errors.New("user not found")
	}

	delete(s.seats, req.EmailAddress)
	delete(s.tickets, req.EmailAddress)

	return &bt.StatusResponse{
		Message: fmt.Sprintf("User %s removed successfully", req.EmailAddress),
	}, nil
}

func (s *BoardTrainServer) ModifySeat(ctx context.Context, req *bt.ModifySeatRequest) (*bt.StatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !checkSectionValid(req.NewSeat) {
		s.l.Printf("Provided seat number is not valid, valid example : A4")
		return &bt.StatusResponse{}, errors.New("user section is not valid in reqeust")
	}

	newSection, newSeatNumber, err := splitSeatNumberFromSection(req.NewSeat)
	if err != nil {
		s.l.Printf("Not able to convert the string to number")
		return &bt.StatusResponse{}, errors.New("seat request is not valid in request")

	}

	if newSeatNumber > 10 && newSection != "A" || newSection != "B" {
		return &bt.StatusResponse{}, errors.New("seat number should not exceed 10 and section should be either A or B")
	}

	s.seats[req.EmailAddress] = &bt.Seat{Section: newSection, Number: int32(newSeatNumber)}

	return &bt.StatusResponse{
		Message: fmt.Sprintf("User %s seat modified to %s successfully", req.EmailAddress, req.NewSeat),
	}, nil

}
