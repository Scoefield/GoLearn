package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (
	database   = "Test"
	collection = "TestModel"
)

//test data
//type Data struct {
//	//Id bson.ObjectId
//	Title string
//	Des string
//	Content string
//	Img string
//	Date time.Time
//}
// test
//func TestInsert(t *testing.T)  {
//	data := &Data{
//		//Id:      bson.NewObjectId().Hex(),
//		Title:   "标题2",
//		Des:     "博客描述信息2",
//		Content: "博客的内容信息2",
//		Img:     "https://upload-images.jianshu.io/upload_images/8679037-67456031925afca6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/700",
//		Date:    time.Now(),
//	}
//
//	err := Insert("Test", "TestModel", data)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("insert success.")
//}

//func TestFindOne(t *testing.T) {
//	var result Data
//	err := FindOne("Test", "TestModel", bson.M{"title": "标题2"}, bson.M{"_id":0}, &result)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("find one success:", result)
//}

//func TestFindAll(t *testing.T) {
//	var result []Data
//	err := FindAll("Test", "TestModel", bson.M{"title": "标题2"}, bson.M{"_id":0}, &result)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("find all success:", result)
//}

//func TestUpdate(t *testing.T) {
//	err := Update("Test", "TestModel", bson.M{"_id": "5cea46ca7eb409458018fe17"}, bson.M{"$set": bson.M{"title": "更新标题"}})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("update success:")
//}

func TestRemove(t *testing.T) {
	err := Remove(database, collection, bson.M{"_id": "5cea50d67eb409458018fedf"})
	if err != nil {
		panic(err)
	}
	fmt.Println("remove success")
}
