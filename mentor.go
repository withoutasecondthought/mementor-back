package mementor_back

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message string `json:"message" example:"some string here"`
} //@name BasicResponse

type Auth struct {
	Email        string `json:"email" validate:"required,email" example:"mrmarkeld@gmail.com"`
	Password     string `json:"password" validate:"required,min=6" example:"123456"`
	ValidProfile bool   `json:"-,omitempty" bson:"validProfile,omitempty" example:"false"`
} //@name PostAuthRequest

type Mentor struct {
	Id                  *primitive.ObjectID `json:"_id" bson:"_id" example:"634afbd6c7cc8190a74feb35"`
	Name                string              `json:"name" bson:"name"  validate:"required" example:"Test"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required" example:"Subject"`
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required" example:"python,js,trash"`
	Grade               string              `json:"grade" bson:"grade" validate:"required,oneof=junior middle senior" example:"junior"`
	Language            []string            `json:"language" bson:"language" example:"russian, english"`
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required,len=3"`
} //@name Mentor

type MentorFullInfo struct {
	Id                  *primitive.ObjectID `json:"_id" bson:"_id" example:"634afbd6c7cc8190a74feb35"`
	Name                string              `json:"name" bson:"name"  validate:"required" example:"Test"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required" example:"Subject"`
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required" example:"cpp, go, scala"`
	Grade               string              `json:"grade" bson:"grade" validate:"required,oneof=junior middle senior" example:"junior"`
	Language            []string            `json:"language" bson:"language" example:"ru, en"`
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required,len=3"`
	ExperienceSince     uint                `json:"experienceSince" bson:"experienceSince"  validate:"required" example:"2019"`
	Email               string              `json:"email,omitempty" bson:"email,omitempty" validate:"required,email" example:"mrmarkeld@gmail.com"`
	Description         string              `json:"description" bson:"description" example:"Im the best from the best"`
	ClassesDone         uint                `json:"classesDone" bson:"classesDone" example:"21"`
	Education           []Education         `json:"education" bson:"education"`
	Technology          []string            `json:"technology" bson:"technology"  validate:"required" example:"cpp, go,scala"`
	CanHelpWith         []string            `json:"canHelpWith" bson:"canHelpWith" example:"Your mother, Your sister"`
	ValidProfile        bool                `json:"validProfile" bson:"validProfile" example:"true"`
} //@name GetMentorResponse

type PostMentorRequest struct {
	MentorFullInfo
} //@name PostMentorRequest

type PutMentorRequest struct {
	MentorFullInfo
} //@name PutMentorRequest

type Education struct {
	Place      string `json:"place" bson:"place"  validate:"required" example:"MGU"`
	Department string `json:"department" bson:"department"  validate:"required" example:"computer science"`
} //@name Education

type Tariff struct {
	Price       uint   `json:"price" bson:"price"  validate:"required" example:"2000"`
	Name        string `json:"name" bson:"name"  validate:"required" example:"big boby"`
	Description string `json:"description" bson:"description"  validate:"required,max=255" example:"Free for you my little friend"`
} //@name Tariff

type ListOfMentorsResponse struct {
	Pages   int      `json:"pages" bson:"pages" validate:"required" example:"1"`
	Mentors []Mentor `json:"mentors" validate:"required" bson:"mentors"`
} //@name PostMentorResponse

type Booking struct {
	CustomerName     string `json:"customerName"  bson:"customerName" validate:"required"`
	CustomerTelegram string `json:"customerTelegram" bson:"customerTelegram" validate:"required"`
	MentorId         string `json:"mentorId" bson:"mentorId" validate:"required"`
	TariffIndex      *int   `json:"tariffIndex" bson:"tariffIndex" validate:"required,min=0,max=2"`
} //@name PostBookingRequest
