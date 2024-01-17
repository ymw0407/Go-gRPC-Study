package insertfriend

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	HeartLock      bool `bson:"heartLock"`
	SystemAlarm    bool `bson:"systemAlarm"`
	QuestionAlarm  bool `bson:"questionAlarm"`
	BluetoothAlarm bool `bson:"bluetoothAlarm"`
	LikeAlarm      bool `bson:"likeAlarm"`
}

type Chip struct {
	SchoolName      string   `bson:"schoolName"`
	SchoolShortName string   `bson:"schoolShortName"`
	Grade           int32    `bson:"grade"`
	InstaAccount    string   `bson:"instaAccount"`
	TiktokAccount   string   `bson:"tiktokAccount"`
	Mbti            string   `bson:"mbti"`
	Others          []string `bson:"others"`
}

type Friend struct {
	FriendId primitive.ObjectID `bson:"friendId"`
}

type BestFriend struct {
	FriendId  primitive.ObjectID `bson:"friendId"`
	PushAlarm bool               `bson:"pushAlarm"`
}

// 짝사랑 상대
type Crush struct {
	FriendId      primitive.ObjectID `bson:"friendId"`
	MyDisplayName string             `bson:"myDisplayName"` // 상대방에게 보여질 나의 이름
	Reason        string             `bson:"reason"`        // 상대방을 좋아하는 이유
	CreatedAt     time.Time          `bson:"createdAt"`
}

type User struct {
	Uid         primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Gender      bool               `bson:"gender"`
	Birth       string             `bson:"birth"`
	PhoneNumber string             `bson:"phoneNumber"`
	ProfileImg  []byte             `bson:"profileImg"`
	Description string             `bson:"description"`
	SchoolCode  string             `bson:"schoolCode"`
	Chip        Chip               `bson:"chip"`
	Config      Config             `bson:"config"`
	TimeStamp   TimeStamp          `bson:"timeStamp"`
}

type Friends struct {
	Uid                primitive.ObjectID `bson:"_id"`
	Crush              Crush              `bson:"crush"`
	BestFriends        []BestFriend       `bson:"bestFriends"`
	BlockedFriends     []Friend           `bson:"blockedFriends"`
	RecommendedFriends []Friend           `bson:"recommendedFriends"`
}

func NewConfig(config Config) Config {
	return Config{
		HeartLock:      config.HeartLock,
		SystemAlarm:    config.SystemAlarm,
		QuestionAlarm:  config.QuestionAlarm,
		BluetoothAlarm: config.BluetoothAlarm,
		LikeAlarm:      config.LikeAlarm,
	}
}

func NewChip(chip Chip) Chip {
	return Chip{
		SchoolName:      chip.SchoolName,
		SchoolShortName: chip.SchoolShortName,
		Grade:           chip.Grade,
		InstaAccount:    chip.InstaAccount,
		TiktokAccount:   chip.TiktokAccount,
		Mbti:            chip.Mbti,
		Others:          chip.Others,
	}
}

func NewUser(user User) User {
	return User{
		Uid:         user.Uid,
		Name:        user.Name,
		Gender:      user.Gender,
		Birth:       user.Birth,
		PhoneNumber: user.PhoneNumber,
		ProfileImg:  user.ProfileImg,
		Description: user.Description,
		Chip:        user.Chip,
		Config:      user.Config,
		TimeStamp:   user.TimeStamp,
	}
}

type TimeStamp struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
