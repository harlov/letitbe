package db
import (
    "time"
)

type CounterEntity struct {
    Id string `bson:"_id"`
    Seq int64 `bson:"seq"`
}

type UserEntity struct {
    Id int64 "_id"
    Username string `bson:"username"`
    PassHash string `bson:"pass_hash"`
    Email string `bson:"email"`
    RegisteredAt time.Time `bson:"registered_at"`
}


type UserSession struct {
    User int64 `bson:"user" json:"user" `
    SessionToken string `bson:"session_token" json:"session_token"`
    StartedAt time.Time `bson:"started_at" json:"started_at"`
    ActiveTo time.Time `bson:"active_to" json:"active_to"`
}