package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
)

// ××××××××××× 定义常量 ×××××××××××××××××
const URL = "127.0.0.1:27017"

var (
	mgoSession *mgo.Session
)

// ×××××××××**×× 初始化init ×××××××××××××××××
func init() {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	mgoSession = session
}

// ××××××××××××××××**×× 获取连接 ×××××××××××××××××××××××
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := mgoSession.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

// ××××××××××××××××××××××**×× 插入 ××××××××××××××××××××××××××
func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(doc)
}

// ××××××××××××××××××××××××××**×× 查询 ××××××××××××××××××××××××××
/*
db:操作的数据库
collection:操作的文档(表)
query:查询条件
selector:需要过滤的数据(projection)
result:查询到的结果
*/
func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}
func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

// ××××××××××××××××××××××××××**×× 更新 ××××××××××××××××××××××××××
/**
db:操作的数据库
collection:操作的文档(表)
selector:更新条件
update:更新的操作
*/
func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

//更新，如果不存在就插入一个新的数据 `upsert:true`
func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

// `multi:true`
func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

// ××××××××××××××××××××××××××**×× 删除 ××××××××××××××××××××××××××
/**
db:操作的数据库
collection:操作的文档(表)
selector:删除条件
*/
func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

// ××××××××××××××××××××××××××**×× 分页查询 ××××××××××××××××××××××××××
/*
db:操作的数据库
collection:操作的文档(表)
page:当前页面
limit:每页的数量值
query:查询条件
selector:需要过滤的数据(projection)
result:查询到的结果
*/
func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}
