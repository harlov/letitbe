package db
import "gopkg.in/mgo.v2/bson"

type CounterEntity struct {
    Id string `bson:"_id"`
    Seq int64 `bson:"seq"`
}

type UserEntity struct {
    Id int64 "_id"
    Username string `bson:"username"`
    PassHash string `bson:"pass_hash"`
    Email string `bson:"email"`
    RegisteredAt bson.MongoTimestamp `bson:"registered_at"`
}


type UserSession struct {
    User UserEntity `bson:"user"`
    StartedAt bson.MongoTimestamp `bson:"started_at"`
}