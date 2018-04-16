//adding payment api code in GO for payment microservice in this file.	

package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/unrolled/render"
    "gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	
)

type Payment struct {
	_id bson.ObjectId `bson:"_id" json:"_id"`
	id int `bson:"id" json:"id"`
	amount int  `bson:"amount" json:"amount"`
}

func newParent(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {

	//parsing json file	
	req.ParseForm()
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	//code for database connection with mongo in local
	database := Database{"localhost", "cmpe281", nil}
	data := &database
	Connect(data)
	//collection
	c := data.db.C("payment")
	
	//need parent id for my api which i took from query string
	var p Payment 
	i, err := strconv.Atoi(req.FormValue("parentid"))
	err = c.Find(bson.M{"id":i}).One(&p) 
		formatter.JSON(w, http.StatusOK, p) //for checking expected output with curl
	if err==nil{
	fmt.Println("Already present in db")
	//handle error	

	} else {
	
	//handle error

	}
	//below code does not give output in struct as needed	
	//for unmarshalling data into JSON format and storing in struct	
	//data1, err := json.Marshal(req.Form)
    	//if err != nil {
         //Handle error
    	//}
    	//var user User
    	//if err = json.Unmarshal(data1, user); err != nil {
        // Handle error
    	//}
   	//fmt.Printf("%+v", user)

	}
}