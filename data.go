package customerpref

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"bytes"
)

func AddCustomerPref(cp CustomerPref){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("Connecting to table")
	c := session.DB("customerpref").C("preferences")
	err = c.Insert(&CustomerPref{"1234", "email"},
					&CustomerPref{"2345", "phone"})
	if err != nil {
		fmt.Println(err)
	}

	result := CustomerPref{}
	err = c.Find(bson.M{"id": "2345"}).One(&result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Notification Preference:", result.NotificationPref)
}

func UpdateCustomerPref(cp CustomerPref){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("Connecting to table")
	c := session.DB("customerpref").C("preferences")
	err = c.Update(bson.M{"_id": cp.Id}, bson.M{"$inc": bson.M{"count": 1}})
  	if err != nil {
    	fmt.Printf("Can't update document %v\n", err)
  }
}

func GetData() string {
	var results = []CustomerPref{}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("Connecting to table")
	c := session.DB("customerpref").C("preferences")
	//result := CustomerPref{}
	err = c.Find(nil).All(&results)
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(results)

	fmt.Println(results[1].NotificationPref)
	n := bytes.Index(b, []byte{0})
	s := string(b[:n])
	return s

}
