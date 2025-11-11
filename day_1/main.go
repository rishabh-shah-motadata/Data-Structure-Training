package day1

import (
	"fmt"
	"strings"
)

const (
	TotalSeats = 300
)

// The reason behind I used struct is because it keeps the code organized.
// Also creating a array and a array inside a struct will consume same memory as we have only one field
// Also in future if we need to add any other fields, we can easily add it
type SeatBookingSystem struct {
	seats [TotalSeats]bool
}

// The reason behind I have defined pointer is because if I dont do that then the whole 300 value array and mutex will be copied and returned
// Which will lead to memory consumption
func newSeatBookingSystem() *SeatBookingSystem {
	return &SeatBookingSystem{
		seats: [TotalSeats]bool{}, // Here taking bool because it will consume 1 byte for every seat
	}
}

// Now defining receiver as pointer receiver because we are modifying the value of the array which we want to be reflected in original struct.
// If I dont do that I need to return a copy of array which will lead to consume more memory.
func (sbs *SeatBookingSystem) bookSeat(seatNumber int) error {
	if seatNumber < 1 || seatNumber > TotalSeats {
		return fmt.Errorf("seat number must be between 1 and %d", TotalSeats)
	}

	if sbs.seats[seatNumber-1] {
		return fmt.Errorf("seat %d is already booked", seatNumber)
	}

	sbs.seats[seatNumber-1] = true
	fmt.Printf("Seat %d booked successfully!\n", seatNumber)
	return nil
}

func (sbs *SeatBookingSystem) displayAvailableSeats() {
	var availableSeats int

	for i := range TotalSeats {
		if !sbs.seats[i] {
			availableSeats++
		}
	}

	fmt.Printf("\nAvailable Seats (%d/%d):\n\n", availableSeats, TotalSeats)
}

func SeatBooking() {
	seatBookingSystem := newSeatBookingSystem()

	for {
		fmt.Print("\n" + strings.Repeat("=", 40))
		fmt.Print("\nSeat Booking System\n")
		fmt.Print(strings.Repeat("=", 40))
		fmt.Print("\n1. Book a Seat")
		fmt.Print("\n2. View Available Seats")
		fmt.Print("\n>2. Exit")
		fmt.Print("\nChoose option (1-3): ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var seatNumber int

			fmt.Print("Enter seat number to book: ")
			if _, err := fmt.Scanln(&seatNumber); err != nil {
				fmt.Printf("Invalid seat number %d", seatNumber)
			}

			if err := seatBookingSystem.bookSeat(seatNumber); err != nil {
				fmt.Printf("Booking failed: %v\n", err)
			}

		case 2:
			seatBookingSystem.displayAvailableSeats()

		default:
			fmt.Println("Thank you for using the seat booking system!")
			return
		}
	}
}
