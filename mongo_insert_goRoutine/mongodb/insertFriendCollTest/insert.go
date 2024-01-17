package insertfriend

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb"
	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb/find"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	sharedUID = 0
	Wg        sync.WaitGroup
	mu        sync.Mutex
)

func InsertTest(schools []find.SchoolTest) {
	i := 0

	for i < 5000 {
		Wg.Add(1)
		go Insert(schools[i])
		i++
	}

	Wg.Wait()
}

func Insert(school find.SchoolTest) {
	mu.Lock()
	uid := sharedUID
	sharedUID += 300
	mu.Unlock()

	if sharedUID > 500000 {
		Wg.Done()
		return
	}

	users := []interface{}{}
	friends := []interface{}{}
	n := 0

	for n < 100 {
		userID := uid + n
		uid := primitive.NewObjectID()
		users = append(users, User{
			Uid:         uid,
			Name:        "테스트유저" + strconv.Itoa(userID),
			Gender:      false,
			Birth:       strconv.Itoa(userID),
			PhoneNumber: strconv.Itoa(userID),
			SchoolCode:  school.Code,
			Chip: Chip{
				SchoolName:      school.Name,
				SchoolShortName: school.Alias,
				Grade:           1,
			},
			TimeStamp: TimeStamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
		friends = append(friends, Friends{
			Uid: uid,
			Crush: Crush{
				Reason:    "그냥",
				CreatedAt: time.Now(),
			},
		})

		n++
	}

	for n < 200 {
		userID := uid + n
		uid := primitive.NewObjectID()
		users = append(users, User{
			Uid:         uid,
			Name:        "테스트유저" + strconv.Itoa(userID),
			Gender:      false,
			Birth:       strconv.Itoa(userID),
			PhoneNumber: strconv.Itoa(userID),
			SchoolCode:  school.Code,
			Chip: Chip{
				SchoolName:      school.Name,
				SchoolShortName: school.Alias,
				Grade:           2,
			},
			TimeStamp: TimeStamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
		friends = append(friends, Friends{
			Uid: uid,
			Crush: Crush{
				Reason:    "그냥",
				CreatedAt: time.Now(),
			},
		})

		n++
	}

	for n < 300 {
		userID := uid + n
		uid := primitive.NewObjectID()
		users = append(users, User{
			Uid:         uid,
			Name:        "테스트유저" + strconv.Itoa(userID),
			Gender:      false,
			Birth:       strconv.Itoa(userID),
			PhoneNumber: strconv.Itoa(userID),
			SchoolCode:  school.Code,
			Chip: Chip{
				SchoolName:      school.Name,
				SchoolShortName: school.Alias,
				Grade:           3,
			},
			TimeStamp: TimeStamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
		friends = append(friends, Friends{
			Uid: uid,
			Crush: Crush{
				Reason:    "그냥",
				CreatedAt: time.Now(),
			},
		})

		n++
	}

	// log.Println(users)
	// log.Println(friends)

	ctx := context.TODO()
	_, err := mongodb.FriendColl.InsertMany(ctx, friends)
	if err != nil {
		log.Println("err: | ", err.Error())
	}
	// log.Println(test)
	_, err = mongodb.UserColl.InsertMany(ctx, users)
	if err != nil {
		log.Println("err: | ", err.Error())
	}
	Wg.Done()
}
