package db
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
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

    err = c.Insert(&user)
    if err != nil {
        log.Fatal(err)
        return err, user
    }

    return err, user
}


func StartSession(user UserEntity) (UserSession, error) {
    c := session.DB("letitbe").C("session")
    var user_session UserSession
    user_session.User = user
    err := c.Insert(&user_session)
    return user_session, err
}