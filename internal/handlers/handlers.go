package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/shadeshade/bookings-api/internal/config"
	"github.com/shadeshade/bookings-api/internal/driver"
	"github.com/shadeshade/bookings-api/internal/helpers"
	"github.com/shadeshade/bookings-api/internal/models"
	"github.com/shadeshade/bookings-api/internal/repository"
	"github.com/shadeshade/bookings-api/internal/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)

const layout = "2006-01-02"
const reservationTypeID = 1 // reservationTypeID is a restriction in database

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(a, db.SQL),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// GetAllRooms returns all rooms
func (m *Repository) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := m.DB.GetRooms()
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(rooms)
}

// GetRoomByID returns room by ID
func (m *Repository) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		helpers.ClientError(w, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(room)
}

// GetRoomAvailability returns all available rooms in the specified date range
func (m *Repository) GetRoomAvailability(w http.ResponseWriter, r *http.Request) {
	sd := r.URL.Query().Get("startDate")
	ed := r.URL.Query().Get("endDate")

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		return
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		return
	}

	rooms, err := m.DB.GetAllAvailableRooms(startDate, endDate)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(rooms)
}

func (m *Repository) CreateReservation(w http.ResponseWriter, r *http.Request) {
	type Reservation struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		RoomID    int    `json:"room_id"`
	}

	var reqBody Reservation
	json.NewDecoder(r.Body).Decode(&reqBody)

	sd := reqBody.StartDate
	ed := reqBody.EndDate

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		return
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		return
	}

	// check if this room exists
	room, err := m.DB.GetRoomByID(reqBody.RoomID)
	if err != nil {
		return
	}

	reservation := models.Reservation{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Phone:     reqBody.Email,
		Email:     reqBody.Phone,
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    reqBody.RoomID,
		Room:      room,
	}

	reservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		return
	}
	reservation.ID = reservationID

	roomRestriction := models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        reqBody.RoomID,
		ReservationID: reservationID,
		RestrictionID: reservationTypeID,
	}

	err = m.DB.InsertRoomRestriction(roomRestriction)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(reservation)
}
