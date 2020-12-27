package models

type AppDomain struct {
	Name string   `bson:"name"`
	URLs []string `bson:"urls"`
}
