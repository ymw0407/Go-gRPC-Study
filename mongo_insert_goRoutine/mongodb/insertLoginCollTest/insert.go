package insertLogin

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	sharedUID = 0
	Wg        sync.WaitGroup
	mu        sync.Mutex
)

func InsertTest() {
	i := 0

	for i < 6000 {
		Wg.Add(1)
		go Insert()
		i++
	}

	Wg.Wait()
}

func Insert() {
	mu.Lock()
	uid := sharedUID
	sharedUID += 300
	mu.Unlock()

	if sharedUID > 500000 {
		Wg.Done()
		return
	}

	users := []interface{}{}
	n := 0

	for n < 300 {
		userID := uid + n
		users = append(users, AuthUser{
			Uid: primitive.NewObjectID(),
			Phone: Phone{
				PhoneNumber:   strconv.Itoa(userID),
				CountryNumber: "82",
			},
		})
		n++
	}

	_, err := mongodb.LoginColl.InsertMany(context.TODO(), users)
	if err != nil {
		log.Println("err: | ", err.Error())
	}
	Wg.Done()
}
