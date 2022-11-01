package repository

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mementor_back "mementor-back"
	"strings"
)

type MentorMongo struct {
	db *mongo.Collection
}

func (m *MentorMongo) GetMyMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error) {
	var mentor mementor_back.MentorFullInfo
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mementor_back.MentorFullInfo{}, err
	}
	_ = m.db.FindOne(ctx, primitive.M{"_id": hexedId}).Decode(&mentor)

	return mentor, nil
}

func (m *MentorMongo) GetMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error) {
	var mentor mementor_back.MentorFullInfo
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mementor_back.MentorFullInfo{}, err
	}
	_ = m.db.FindOne(ctx, primitive.M{"_id": hexedId, "validProfile": true}).Decode(&mentor)
	if mentor.Email == "" {
		return mementor_back.MentorFullInfo{}, errors.New("no such user")
	}

	return mentor, nil
}

func (m *MentorMongo) PutMentor(ctx context.Context, mentor mementor_back.MentorFullInfo) error {
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

func (m *MentorMongo) ListOfMentors(ctx context.Context, page uint, params mementor_back.SearchParameters) (mementor_back.ListOfMentorsResponse, error) {

	var cur *mongo.Cursor
	var err error

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSkip(int64(page) * 20)

	splitSearch := strings.Split(params.Search, " ")

	var response mementor_back.ListOfMentorsResponse

	baseRequest := bson.M{
		"grade":           bson.M{"$in": params.Grade},
		"experienceSince": bson.M{"$lte": params.ExperienceSince},
		"tariff.0.price":  bson.M{"$gte": params.MinPrice},
		"tariff.2.price":  bson.M{"$lte": params.MaxPrice},
		"validProfile":    params.ValidProfile,
	}

	requestWithSearch := bson.M{
		"$or": bson.A{
			bson.M{"description": bson.M{"$regex": primitive.Regex{
				Pattern: params.Search,
				Options: "im",
			}}},
			bson.M{"name": bson.M{"$regex": primitive.Regex{
				Pattern: params.Search,
				Options: "im",
			}}},
			bson.M{"surname": bson.M{"$regex": primitive.Regex{
				Pattern: params.Search,
				Options: "im",
			}}},
			bson.M{"programmingLanguage": bson.M{"$in": splitSearch}},
			bson.M{"technology": bson.M{"$in": splitSearch}},
		},
		"grade":           bson.M{"$in": params.Grade},
		"experienceSince": bson.M{"$lte": params.ExperienceSince},
		"tariff.0.price":  bson.M{"$gte": params.MinPrice},
		"tariff.2.price":  bson.M{"$lte": params.MaxPrice},
		"validProfile":    params.ValidProfile,
	}

	if params.Search == "" {
		cur, err = m.db.Find(ctx, baseRequest, opts)
	} else {
		cur, err = m.db.Find(ctx, requestWithSearch, opts)
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return mementor_back.ListOfMentorsResponse{}, nil
		}
		return mementor_back.ListOfMentorsResponse{}, err
	}
	err = cur.All(ctx, &response.Mentors)
	if err != nil {
		return mementor_back.ListOfMentorsResponse{}, err
	}

	number, err := m.db.CountDocuments(ctx, bson.M{"validProfile": true})
	if err != nil {
		return mementor_back.ListOfMentorsResponse{}, err
	}
	response.Pages = int(number / 20)
	return response, nil
}

func NewMentor(db *mongo.Database) *MentorMongo {
	return &MentorMongo{
		db: db.Collection(viper.GetString("db.collection")),
	}
}
