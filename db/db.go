package db
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "time"
    "github.com/nu7hatch/gouuid"

)

var session *mgo.Session
var db_inst *mgo.Database
func Init() error {
    var err error
    session, err = mgo.Dial("localhost")
    if err != nil {
        return err
    }
    session.SetMode(mgo.Monotonic, true)
    db_inst = session.DB("letitbe")

    err = checkCounter("userid")
    checkNextSeqFunction()

    return nil
}

func checkCounter(name string) error {
    var err error
    var counter_entity CounterEntity
    var query_counter *mgo.Query
    var counter_count int
    query_counter = db_inst.C("counters").Find(bson.M{"_id":name})
    counter_count, err = query_counter.Count()

    if err != nil {
        return err
    }

    if counter_count > 0 {
        log.Println(name+" seq is exist. skiped")
        return err
    }

    counter_entity.Id = name
    counter_entity.Seq = 0

    err = db_inst.C("counters").Insert(&counter_entity)

    if err == nil {
        log.Println(name+" seq has created")
    }

    return err
}


func checkNextSeqFunction() {
    func_def := `db.system.js.save({
                    _id : "getNextSequence",
                    value: function (name) {
                   var ret = db.counters.findAndModify(
                          {
                            query: { _id: name },
                            update: { $inc: { seq: 1 } },
                            new: true,
                            upsert: true
                          }
                   );

                   return ret.seq;
                }
                    });`
    result := bson.M{}
    db_inst.Run(bson.M{"eval": func_def}, &result)
}


func FindUser(username string) (error, UserEntity) {
    c := session.DB("letitbe").C("user")
    var user_entity UserEntity
    err := c.Find(bson.M{"username": username}).One(&user_entity)
    return err, user_entity
}

func AddUser(user UserEntity) (error, UserEntity) {
    var err error
    c := db_inst.C("user")
    result_o := bson.M{}
    err = db_inst.Run(bson.M{"eval" : "getNextSequence(\"userid\")"}, &result_o)
    if err != nil {
        return err, user
    }

    log.Println(result_o)

    user.Id = int64(result_o["retval"].(float64))
    user.RegisteredAt = time.Now()
    err = c.Insert(&user)
    if err != nil {
        log.Fatal(err)
        return err, user
    }

    return err, user
}


func StartSession(user UserEntity, live_time int32) (UserSession, error) {
    c := session.DB("letitbe").C("session")
    var user_session UserSession
    var session_token *uuid.UUID
    var err error
    user_session.User = user.Id
    user_session.StartedAt = time.Now()
    user_session.ActiveTo = user_session.StartedAt.Add(time.Duration(live_time)*time.Second)
    session_token,err = uuid.NewV4()
    user_session.SessionToken = session_token.String()
    if err != nil {
        return user_session, err
    }

    err = c.Insert(&user_session)
    return user_session, err
}

func FindSession(session_token string)(error, UserSession) {
    var session UserSession
    var err error
    err = db_inst.C("session").Find(bson.M{"session_token": session_token}).One(&session)
    if err != nil {
        return err, session
    }
    return err, session
}