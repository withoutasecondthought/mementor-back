package mementor_back

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Message string `json:"message"`
} // @name BasicResponse

type Auth struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	ValidProfile bool   `json:"-,omitempty" bson:"validProfile,omitempty"`
} // @name PostAuthRequest

type Mentor struct {
	ID                  *primitive.ObjectID `json:"_id" bson:"_id" `
	Name                string              `json:"name" bson:"name"  validate:"required"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required" `
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required"`
	Grade               string              `json:"grade" bson:"grade" validate:"required,oneof=junior middle senior" `
	Language            []string            `json:"language" bson:"language" `
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required,len=3"`
	Image               *Image              `bson:"image" json:"image,omitempty"`
} // @name Mentor

type MentorFullInfo struct {
	ID                  *primitive.ObjectID `json:"_id" bson:"_id"`
	Name                string              `json:"name" bson:"name"  validate:"required"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required"`
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required"`
	Grade               string              `json:"grade" bson:"grade" validate:"required,oneof=junior middle senior"`
	Language            []string            `json:"language" bson:"language"`
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required,len=3"`
	ExperienceSince     uint                `json:"experienceSince" bson:"experienceSince"  validate:"required"`
	Email               string              `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Description         string              `json:"description" bson:"description" `
	ClassesDone         uint                `json:"classesDone" bson:"classesDone" `
	Education           []Education         `json:"education" bson:"education"`
	Technology          []string            `json:"technology" bson:"technology"  validate:"required" `
	CanHelpWith         []string            `json:"canHelpWith" bson:"canHelpWith" `
	ValidProfile        bool                `json:"validProfile" bson:"validProfile"`
	Image               *Image              `json:"image" bson:"image,omitempty"`
} // @name GetMentorResponse

type PutMentorRequest struct {
	MentorFullInfo
} // @name PutMentorRequest

type Education struct {
	Place      string `json:"place" bson:"place"  validate:"required"`
	Department string `json:"department" bson:"department"  validate:"required" `
} // @name Education

type Tariff struct {
	Price       uint   `json:"price" bson:"price"  validate:"required"`
	Name        string `json:"name" bson:"name"  validate:"required"`
	Description string `json:"description" bson:"description"  validate:"required,max=255"`
} // @name Tariff

type ListOfMentorsResponse struct {
	Pages   int64    `json:"pages" bson:"pages" validate:"required"`
	Mentors []Mentor `json:"mentors" validate:"required" bson:"mentors"`
} // @name PostMentorResponse

type Booking struct {
	CustomerName     string `json:"customerName"  bson:"customerName" validate:"required"`
	CustomerTelegram string `json:"customerTelegram" bson:"customerTelegram" validate:"required"`
	MentorID         string `json:"mentorId" bson:"mentorId" validate:"required"`
	TariffIndex      int8   `json:"tariffIndex" bson:"tariffIndex" validate:"min=0,max=2"`
} // @name PostBookingRequest

type SearchParameters struct {
	ValidProfile    bool     `json:"-" bson:"validProfile"`
	Search          string   `json:"search" `
	Grade           []string `json:"grade"`
	ExperienceSince int      `json:"experienceSince"`
	MinPrice        int      `json:"minPrice"`
	MaxPrice        int      `json:"maxPrice"`
} // @name PostMentorRequest

type PostImage struct {
	ID     *primitive.ObjectID `json:"-" bson:"_id"`
	Base64 string              `json:"base64"`
	Image  Image               `json:"-" bson:"image"`
} // @name PostImageRequest

type Image struct {
	BigImage   string `json:"512x512"`
	SmallImage string `json:"144x144"`
} // @name Image

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"userId" bson:"userId"`
}
