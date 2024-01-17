package find

type SchoolTest struct {
	Code  string `bson:"_id"`
	Name  string `bson:"name"`
	Alias string `bson:"alias"`
}
