package mementor_back

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message string `json:"message"`
}

type Auth struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	ValidProfile bool   `json:"-" bson:"validProfile"`
}

type Mentor struct {
	Id                  *primitive.ObjectID `json:"_id" bson:"_id"`
	Email               string              `json:"email" bson:"email" validate:"email"`
	Name                string              `json:"name" bson:"name"  validate:"required"`
	Surname             string              `json:"surname" bson:"surname"  validate:"required"`
	ExperienceSince     uint                `json:"experienceSince" bson:"experienceSince"  validate:"required"`
	ProgrammingLanguage []string            `json:"programmingLanguage" bson:"programmingLanguage"  validate:"required"`
	Technology          []string            `json:"technology" bson:"technology"  validate:"required"`
	Grade               string              `json:"grade" bson:"grade" validate:"required"`
	Description         string              `json:"description" bson:"description"`
	ClassesDone         uint                `json:"classesDone" bson:"classesDone"`
	Education           []Education         `json:"education" bson:"education"`
	Tariff              []Tariff            `json:"tariff" bson:"tariff"  validate:"required"`
	CanHelpWith         []string            `json:"canHelpWith" bson:"canHelpWith"`
	Language            []string            `json:"language" bson:"language"`
	ValidProfile        bool                `json:"validProfile" bson:"validProfile"`
}

type Education struct {
	Place      string `json:"place" bson:"place"  validate:"required"`
	Department string `json:"department" bson:"department"  validate:"required"`
}

type Tariff struct {
	Price uint   `json:"price" bson:"price"  validate:"required"`
	Name  string `json:"name" bson:"name"  validate:"required"`
}
