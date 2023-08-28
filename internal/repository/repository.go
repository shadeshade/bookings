package repository

import (
	"github.com/shadeshade/bookings-api/internal/models"
	"time"
)

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	GetRooms() ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	GetAllAvailableRooms(start, end time.Time) ([]models.Room, error)
}
