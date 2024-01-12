package insert

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	SharedCode = []string{}
	sharedUID  = 0
	Wg         sync.WaitGroup
	mu         sync.Mutex
)

func InsertTest(codes []string) {
	log.Println(codes)

	i := 0

	for i < len(codes) {
		// for i < 3 {
		Wg.Add(1)
		go Insert(codes[i])
		i++
	}

	Wg.Wait()
}

// func Insert(code string) {
// 	mu.Lock()
// 	uid := sharedUID
// 	sharedUID += 300
// 	mu.Unlock()

// 	users := []interface{}{}
// 	n := 0

// 	for n < 100 {
// 		userID := uid + n
// 		users = append(users, UserTestSchema{
// 			Uid:        userID,
// 			SchoolCode: code,
// 			Name:       "테스트유저" + strconv.Itoa(userID),
// 			Grade:      1,
// 		})
// 		n++
// 	}
// 	for n < 200 {
// 		userID := uid + n
// 		users = append(users, UserTestSchema{
// 			Uid:        userID,
// 			SchoolCode: code,
// 			Name:       "테스트유저" + strconv.Itoa(userID),
// 			Grade:      2,
// 		})
// 		n++
// 	}
// 	for n < 300 {
// 		userID := uid + n
// 		users = append(users, UserTestSchema{
// 			Uid:        userID,
// 			SchoolCode: code,
// 			Name:       "테스트유저" + strconv.Itoa(userID),
// 			Grade:      3,
// 		})
// 		n++
// 	}

// 	_, err := mongodb.UserColl.InsertMany(context.TODO(), users)
// 	if err != nil {
// 		log.Println("err: | ", err.Error())
// 	}
// 	Wg.Done()
// }

func Insert(code string) {
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

	for n < 100 {
		userID := uid + n
		users = append(users, User{
			Uid:  primitive.NewObjectID(),
			Name: "테스트유저" + strconv.Itoa(userID),
			Chip: Chip{
				SchoolName: code,
				Grade:      1,
			},
			PhoneNumber: code,
		})
		n++
	}
	for n < 200 {
		userID := uid + n
		users = append(users, User{
			Uid:  primitive.NewObjectID(),
			Name: "테스트유저" + strconv.Itoa(userID),
			Chip: Chip{
				SchoolName: code,
				Grade:      2,
			},
			PhoneNumber: code,
		})
		n++
	}
	for n < 300 {
		userID := uid + n
		users = append(users, User{
			Uid:  primitive.NewObjectID(),
			Name: "테스트유저" + strconv.Itoa(userID),
			Chip: Chip{
				SchoolName: code,
				Grade:      3,
			},
			PhoneNumber: code,
		})
		n++
	}

	_, err := mongodb.UserColl.InsertMany(context.TODO(), users)
	if err != nil {
		log.Println("err: | ", err.Error())
	}
	Wg.Done()
}
