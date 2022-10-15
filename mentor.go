package mementor_back

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message string `json:"message"`
}

type Auth struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	ValidProfile bool   `json:"-,omitempty" bson:"validProfile,omitempty"`
}

type Mentor struct {
	Id                  *primitive.ObjectID `json:"_id" bson:"_id"`
	Name                string              `json:"name" bson:"name"  validate:"required"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required"`
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required"`
	Grade               string              `json:"grade" bson:"grade" validate:"required"`
	Language            []string            `json:"language" bson:"language"`
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required,len=3"`
}

type MentorFullInfo struct {
	Mentor
	ExperienceSince uint        `json:"experienceSince" bson:"experienceSince"  validate:"required"`
	Email           string      `json:"email,omitempty" bson:"email,omitempty" validate:"email"`
	Description     string      `json:"description" bson:"description"`
	ClassesDone     uint        `json:"classesDone" bson:"classesDone"`
	Education       []Education `json:"education" bson:"education"`
	Technology      []string    `json:"technology" bson:"technology"  validate:"required"`
	CanHelpWith     []string    `json:"canHelpWith" bson:"canHelpWith"`
	ValidProfile    bool        `json:"validProfile" bson:"validProfile"`
}

type Education struct {
	Place      string `json:"place" bson:"place"  validate:"required"`
	Department string `json:"department" bson:"department"  validate:"required"`
}

type Tariff struct {
	Price       uint   `json:"price" bson:"price"  validate:"required"`
	Name        string `json:"name" bson:"name"  validate:"required"`
	Description string `json:"description" bson:"description"  validate:"required,max=255"`
}
