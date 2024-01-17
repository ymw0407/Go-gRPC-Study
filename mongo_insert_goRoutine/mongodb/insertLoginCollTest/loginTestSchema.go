package insertLogin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUser struct {
	Uid                  primitive.ObjectID `bson:"_id"`
	Phone                Phone              `bson:"phone"`
	Password             string             `bson:"password"`
	CreatedAt            time.Time          `bson:"createdAt"`
	PhoneNumberUpdatedAt time.Time          `bson:"phoneNumberUpdatedAt"`
	PasswordUpdatedAt    time.Time          `bson:"passwordUpdatedAt"`
}

type Phone struct {
	PhoneNumber   string `bson:"phoneNumber"`
	CountryNumber string `bson:"country"`
}
