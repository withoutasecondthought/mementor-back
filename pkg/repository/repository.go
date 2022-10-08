package repository

import "go.mongodb.org/mongo-driver/mongo"

type Mentor interface {
}

type Repository struct {
	Mentor
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Mentor: NewMentor(db),
	}

}
