package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	InsertOne() (any, error)
	InsertMany() ([]any, error)
	updateOne() (any, error)
	updateMany() (any, error)
	deleteOne(string) (any, error)
	deleteMany() (any, error)
	find() ([]Fact, error)
}

type Mongodb struct {
	Client *mongo.Client
}

type Fact struct {
	Fact   string
	Length int64
}

func NewStorage() (Storage, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &Mongodb{
		Client: client,
	}, nil
}

func (s *Mongodb) InsertOne() (any, error) {
	coll := s.Client.Database("catfact").Collection("facts")
	result, err := coll.InsertOne(context.TODO(), Fact{"test fact", 8})
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (s *Mongodb) InsertMany() ([]any, error) {

	coll := s.Client.Database("catfact").Collection("facts")
	factList := []any{
		Fact{"test fact 2", 9},
		Fact{"test fact 2", 10},
	}
	result, err := coll.InsertMany(context.TODO(), factList)

	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (s *Mongodb) updateOne() (any, error) {
	coll := s.Client.Database("catfact").Collection("facts")
	result, err := coll.UpdateOne(context.TODO(), bson.M{"fact": "test face updated 2"}, bson.M{
		"$set": bson.M{"fact": "test face updated 3"},
	})
	if err != nil {
		return nil, err
	}
	return result.ModifiedCount, nil
}

func (s *Mongodb) updateMany() (any, error) {
	coll := s.Client.Database("catfact").Collection("facts")
	result, err := coll.UpdateMany(context.TODO(), bson.M{"length": bson.M{"$lte": 10}}, bson.M{
		"$set": bson.M{"fact": "test face one title for secound time"},
	})
	if err != nil {
		return nil, err
	}
	return result.ModifiedCount, nil
}

func (s *Mongodb) deleteOne(id string) (any, error) {
	OID, err := primitive.ObjectIDFromHex(id) //convert id to ObjectID(id)
	coll := s.Client.Database("catfact").Collection("facts")
	result, err := coll.DeleteOne(context.TODO(), bson.M{"_id": OID})
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}

func (s *Mongodb) deleteMany() (any, error) {
	coll := s.Client.Database("catfact").Collection("facts")
	result, err := coll.DeleteMany(context.TODO(), bson.M{"length": bson.M{
		"$lte": 10,
	}})
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}

func (s *Mongodb) find() ([]Fact, error) {
	Facts := []Fact{}
	coll := s.Client.Database("catfact").Collection("facts")
	//filter := bson.M{"$text": bson.M{"$search": "diesl"}} //search { $text: { $search: "coffee shop" }// it fails beacuse theere is no text index
	filter := bson.M{"length": bson.M{
		"$lte": 25,
	}}
	//filter = bson.M{}

	curs, err := coll.Find(context.TODO(), filter, options.Find().SetLimit(10)) //show only 10 coument
	defer curs.Close(context.TODO())

	err = curs.All(context.TODO(), &Facts)
	if err != nil {
		return nil, err
	}
	return Facts, nil
}

func main() {
	s, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	// insertOne, err := s.InsertOne()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(insertOne)

	// insertManyids, err := s.InsertMany()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(insertManyids)

	// updatedCount, err := s.updateOne()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v is number of updated doc", updatedCount)

	// updatedCount, err := s.updateMany()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v is number of updated doc", updatedCount)

	// DeletedCount, err := s.deleteOne("65338d45272e036c19721cd6")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v is number of deleted docs\n", DeletedCount)

	// DeletedCount, err := s.deleteMany()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v is number of deleted docs\n", DeletedCount)

	facs, err := s.find()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(facs)

}
