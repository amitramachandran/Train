package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bt "github.com/amitramachandran/train-grpc/boardTrain"
	"github.com/amitramachandran/train-grpc/client/requests"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:8099"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	c := bt.NewBoardingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// To purchase a ticket
	purchaseticketresp, err := c.PurchaseTicket(ctx, requests.GetTicketRequest())
	if err != nil {
		log.Fatalf("Not able to purchase ticket, %v", err)
	}
	fmt.Println(purchaseticketresp)

	// To get receipt
	getReceiptResp, err := c.GetReceipt(ctx, requests.GetReceiptRequest())
	if err != nil {
		log.Fatalf("Not able to get receipt of ticket, %v", err)
	}
	fmt.Println(getReceiptResp)

	//To view Users in one section
	getPassengersBySectionResp, err := c.ViewSeats(ctx, requests.GetSectionRequest())
	if err != nil {
		log.Fatalf("Not able to get users of given section, %v", err)
	}
	fmt.Println(getPassengersBySectionResp)

	// To remove user from train
	removeUserResp, err := c.RemoveUser(ctx, requests.RemoveUserRequest())
	if err != nil {
		log.Fatalf("Not able to remove the user %v", err)
	}
	fmt.Println(removeUserResp)

	// To modify user seat in train
	modifySeatResp, err := c.ModifySeat(ctx, requests.ModifySeatRequest())
	if err != nil {
		log.Fatalf("Not able to modify user seat %v", err)
	}
	fmt.Println(modifySeatResp)

}
