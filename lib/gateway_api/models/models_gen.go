// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

// Booking has now been confirmed.
type Booking struct {
	// unique identifier of the booking
	ID string `json:"id"`
	// unique identifier of the venue
	VenueID string `json:"venueId"`
	// email of the customer
	Email string `json:"email"`
	// amount of people attending the booking
	People int `json:"people"`
	// start time of the booking (hh:mm)
	StartsAt time.Time `json:"startsAt"`
	// end time of the booking (hh:mm)
	EndsAt time.Time `json:"endsAt"`
	// duration of the booking in minutes
	Duration int `json:"duration"`
	// unique identifier of the booking table
	TableID string `json:"tableId"`
}

// Slot is a possible booking that has yet to be confirmed.
type BookingInput struct {
	// unique identifier of the venue
	VenueID string `json:"venueId"`
	// email of the customer
	Email string `json:"email"`
	// amount of people attending the booking
	People int `json:"people"`
	// start time of the booking (YYYY-MM-DDThh:mm:ssZ)
	StartsAt time.Time `json:"startsAt"`
	// duration of the booking in minutes
	Duration int `json:"duration"`
}

// Booking Enquiry Response.
type GetSlotResponse struct {
	// slot matching the given enquiy
	Match *Slot `json:"match"`
	// slots have match the enquiry but have different starting times
	OtherAvailableSlots []*Slot `json:"otherAvailableSlots"`
}

type IsAdminInput struct {
	VenueID string `json:"venueId"`
}

// Day specific operating hours.
type OpeningHoursSpecification struct {
	// the day of the week for which these opening hours are valid
	DayOfWeek DayOfWeek `json:"dayOfWeek"`
	// the opening time of the place or service on the given day(s) of the week
	Opens TimeOfDay `json:"opens"`
	// the closing time of the place or service on the given day(s) of the week
	Closes TimeOfDay `json:"closes"`
	// date the special opening hours starts at. only valid for special opening hours
	ValidFrom *time.Time `json:"validFrom"`
	// date the special opening hours ends at. only valid for special opening hours
	ValidThrough *time.Time `json:"validThrough"`
}

// Input to remove a venue table
type RemoveTableInput struct {
	// unique venue identifier the table belongs to
	VenueID string `json:"venueId"`
	// unique identifier of the table to be removed
	TableID string `json:"tableId"`
}

// Slot is a possible booking that has yet to be confirmed.
type Slot struct {
	// unique identifier of the venue
	VenueID string `json:"venueId"`
	// email of the customer
	Email string `json:"email"`
	// amount of people attending the booking
	People int `json:"people"`
	// desired start time of the booking (YYYY-MM-DDThh:mm:ssZ)
	StartsAt time.Time `json:"startsAt"`
	// potential ending time of the booking (YYYY-MM-DDThh:mm:ssZ)
	EndsAt time.Time `json:"endsAt"`
	// potential duration of the booking in minutes
	Duration int `json:"duration"`
}

// Slot Input is a booking enquiry.
type SlotInput struct {
	// unique identifier of the venue
	VenueID string `json:"venueId"`
	// email of the customer
	Email string `json:"email"`
	// amount of people attending the booking
	People int `json:"people"`
	// desired start time of the booking (YYYY-MM-DDThh:mm:ssZ)
	StartsAt time.Time `json:"startsAt"`
	// desired duration of the booking in minutes
	Duration int `json:"duration"`
}

// An individual table at a venue.
type Table struct {
	// unique identifier of the table
	ID string `json:"id"`
	// name of the table
	Name string `json:"name"`
	// maximum amount of people that can sit at table
	Capacity int `json:"capacity"`
}

// An individual table at a venue.
type TableInput struct {
	// unique venue identifier the table belongs to
	VenueID string `json:"venueId"`
	// name of the table
	Name string `json:"name"`
	// maximum amount of people that can sit at table
	Capacity int `json:"capacity"`
}

// Venue where a booking can take place.
type Venue struct {
	// unique identifier of the venue
	ID string `json:"id"`
	// name of the venue
	Name string `json:"name"`
	// operating hours of the venue
	OpeningHours []*OpeningHoursSpecification `json:"openingHours"`
	// special operating hours of the venue
	SpecialOpeningHours []*OpeningHoursSpecification `json:"specialOpeningHours"`
	// tables at the venue
	Tables []*Table `json:"tables"`
}
