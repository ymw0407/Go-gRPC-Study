/*
source by : https://www.mongodb.com/docs/drivers/go/current/quick-start/#std-label-golang-quickstart
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil { // dotenv load
		log.Println("No .env file found") // if err가 비어있는게 아니라면... -> log.Println 수행
	}

	uri := os.Getenv("MONGODB_URI") // 환경변수 가져와서 uri 변수에 저장
	if uri == "" {                  // 만약 MONGODB_URI가 환경 변수에 없다면 빈 스트링 ""를 내보냄
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri)) // Mongodb 연결
	if err != nil {                                                              // 에러가 발생한 시점에 에러를 출력하고 종료
		panic(err)
	}

	defer func() { // 함수를 바로 실행하지 않고 종료 시점에 실행 -> 종료 시점에 mongodb의 connection을 끝낸다.
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies") // sample_mflix DB에 연결, movies Collection에 접속
	title := "Back to the Future"

	var result bson.M // 순서가 없는 map 형태. 순서를 유지하지 않는다는 점을 빼면 D와 같습니다.
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result) // 하나의 BSON 도큐멘트. MongoDB command 처럼 순서가 중요한 경우에 사용합니다. // FindOne
	curs, err := coll.Find(context.TODO(), bson.D{{"title", title}}) //.Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

    /* find */
    for curs.Next(context.TODO()) {
		var elem bson.M
		err := curs.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		// find 결과 print
		fmt.Println(elem)
	}
    /* find */

    /* findOne */
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
    /* findOne */
}
