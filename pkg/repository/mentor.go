package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mementor_back "mementor-back"
)

type MentorMongo struct {
	db *mongo.Collection
}

func (m *MentorMongo) GetMentor(ctx context.Context, id string) (mementor_back.Mentor, error) {
	var mentor mementor_back.Mentor
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mementor_back.Mentor{}, err
	}
	_ = m.db.FindOne(ctx, primitive.M{"_id": hexedId}).Decode(&mentor)

	return mentor, nil
}

func (m *MentorMongo) PutMentor(ctx context.Context, mentor mementor_back.Mentor) error {
	_, err := m.db.UpdateOne(ctx, primitive.M{"_id": mentor.Id}, bson.M{"$set": mentor})
	return err
}

func (m *MentorMongo) DeleteMentor(ctx context.Context, id string) error {
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.db.DeleteOne(ctx, primitive.M{"_id": hexedId})
	return err
}

func (m *MentorMongo) ListOfMentors(ctx context.Context, page uint, params interface{}) ([]*mementor_back.Mentor, error) {
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSkip(int64(page) * 20)

	var mentors []*mementor_back.Mentor

	cur, err := m.db.Find(ctx, params, opts)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &mentors)
	if err != nil {
		return nil, err
	}

	return mentors, nil
}

func NewMentor(db *mongo.Database) *MentorMongo {
	return &MentorMongo{
		db: db.Collection(viper.GetString("db.collection")),
	}
}
