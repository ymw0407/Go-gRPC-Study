package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string
	Email string
}

func MongoConnection(MONGODB_URI string) (client *mongo.Client) {
	if MONGODB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}
	return client
}

func MongoDisconnection(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func MongoUserInsert(user User, coll *mongo.Collection) {
	coll.InsertOne(context.TODO(), user)
}

// func (message string) {

// }

// func main() {
// 	if err := godotenv.Load(); err != nil { // dotenv load
// 		log.Println("No .env file found") // if err가 비어있는게 아니라면... -> log.Println 수행
// 	}

// 	uri := os.Getenv("MONGODB_URI") // 환경변수 가져와서 uri 변수에 저장
// 	if uri == "" {                  // 만약 MONGODB_URI가 환경 변수에 없다면 빈 스트링 ""를 내보냄
// 		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
// 	}
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri)) // Mongodb 연결
// 	if err != nil {                                                              // 에러가 발생한 시점에 에러를 출력하고 종료
// 		panic(err)
// 	}

// 	defer func() { // 함수를 바로 실행하지 않고 종료 시점에 실행 -> 종료 시점에 mongodb의 connection을 끝낸다.
// 		if err := client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	coll := client.Database("grpc").Collection("users") // sample_mflix DB에 연결, movies Collection에 접속
// 	user := User{name : }
// }
