package types

type Event struct {
	Name             string
	Tickets          uint
	RemainingTickets uint
	Attendances      []User
}
