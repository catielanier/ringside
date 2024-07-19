package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role int

const (
	Admin Role = iota
	Moderator
	Player
	Banned
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Role     Role               `json:"role,omitempty" bson:"role,omitempty"`
}

type Session struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Country   string             `json:"country,omitempty" bson:"country,omitempty"`
	IpAddress string             `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
}
