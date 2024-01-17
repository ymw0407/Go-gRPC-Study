package find

import (
	"context"
	"log"

	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindSchool() []SchoolTest {

	filter := bson.D{}

	project := bson.D{{Key: "code", Value: 1}, {Key: "name", Value: 1}, {Key: "alias", Value: 1}}
	opts := options.Find().SetProjection(project)

	cursor, err := mongodb.SchoolColl.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err.Error())
	}

	var codes []SchoolTest

	for cursor.Next(context.TODO()) {
		var SchoolTest SchoolTest
		err := cursor.Decode(&SchoolTest)
		if err != nil {
			log.Println("user_find.go, MongoUserFind func : ", err)
		}

		codes = append(codes, SchoolTest)
	}

	return codes
}
