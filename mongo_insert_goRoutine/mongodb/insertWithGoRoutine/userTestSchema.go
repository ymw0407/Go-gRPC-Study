package insert

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTestSchema struct {
	Uid        int    `bson:"uid"`
	SchoolCode string `bson:"schoolCode"`
	Name       string `bson:"name"`
	Grade      int    `bson:"grade"`
}

type Config struct {
	HeartLock      bool `bson:"heartLock"`
	SystemAlarm    bool `bson:"systemAlarm"`
	QuestionAlarm  bool `bson:"questionAlarm"`
	BluetoothAlarm bool `bson:"bluetoothAlarm"`
	LikeAlarm      bool `bson:"likeAlarm"`
}

type Chip struct {
	SchoolName    string   `bson:"schoolName"`
	Grade         int      `bson:"grade"`
	InstaAccount  string   `bson:"instaAccount"`
	TiktokAccount string   `bson:"tiktokAccount"`
	Mbti          string   `bson:"mbti"`
	Others        []string `bson:"others"`
}

// TODO: Friend 타입 생성하기
// TODO: 차단친구 목록 생성하기
type User struct {
	Uid         primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Gender      bool               `bson:"gender"`
	Birth       string             `bson:"birth"`
	PhoneNumber string             `bson:"phoneNumber"`
	ProfileImg  []byte             `bson:"profileImg"` // TODO: 배열이 맞는 지 확인
	Description string             `bson:"description"`
	Chip        Chip               `bson:"chip"`
	Crush       string             `bson:"crush"` // 짝사랑 상대
	BestFriends []string           `bson:"bestFriends"`
	Config      Config             `bson:"config"`
	TimeStamp   TimeStamp          `bson:"timeStamp"`
}

type TimeStamp struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type Term struct {
	Uid       string    `bson:"_id"`
	Agree     bool      `bson:"agree"`
	TimeStamp TimeStamp `bson:"timeStamp"`
}

type Phone struct {
	PhoneNumber   string `bson:"phoneNumber"`
	CountryNumber string `bson:"country"`
}
