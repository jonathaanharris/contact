package controllers

import (
	"api-contact/config"
	"api-contact/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var rdb *redis.Client
var db *gorm.DB
var err error
var redisHost = os.Getenv("REDISHOST")
var redisPassword = os.Getenv("PASSWORDREDIS")

func GetContact(w http.ResponseWriter, r *http.Request) {
	var contact []models.Contact

	key := "contact"
	db = config.ConnectDatabase()
	rdb = newRedisClient(redisHost, redisPassword)

	contactRedis, _ := rdb.Get(context.Background(), key).Result()
	if contactRedis == "" {
		fmt.Println("from database")
		db.Find(&contact)
		data, _ := json.Marshal(contact)
		ttl := redis.KeepTTL
		op1 := rdb.Set(context.Background(), key, data, time.Duration(ttl))
		if err := op1.Err(); err != nil {
			fmt.Printf("unable to SET data. error: %v", err)
			return
		}
		json.NewEncoder(w).Encode(contact)
	} else {
		fmt.Println("from redis")
		op2, _ := rdb.Get(context.Background(), key).Result()
		var cont []models.Contact
		json.Unmarshal([]byte(op2), &cont)
		json.NewEncoder(w).Encode(cont)
	}
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db = config.ConnectDatabase()

	var contact models.Contact
	db.First(&contact, params["id"])

	if contact.ID == 0 {
		json.NewEncoder(w).Encode("DATA NOT FOUND")
	} else {
		json.NewEncoder(w).Encode(&contact)
	}
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	db = config.ConnectDatabase()
	rdb = newRedisClient(redisHost, redisPassword)

	headerContentTtype := r.Header.Get("Content-Type")
	rdb.Del(context.Background(), "contact")

	if headerContentTtype == "application/x-www-form-urlencoded" {
		r.ParseForm()
		contact.Name = r.FormValue("Name")
		contact.PhoneNumber = r.FormValue("PhoneNumber")
		if contact.Name == "" || contact.PhoneNumber == "" {
			json.NewEncoder(w).Encode("need fill the input!")
		} else {
			createdPerson := db.Create(&contact)
			err = createdPerson.Error
			if err != nil {
				json.NewEncoder(w).Encode(err)
			} else {
				json.NewEncoder(w).Encode(&contact)
			}
		}
	} else {
		// fmt.Println(r.Body)
		json.NewDecoder(r.Body).Decode(&contact)
		// json.NewDecoder(r.Form).Decode(&contacttemp)

		// fmt.Println(contact)

		if contact.Name == "" || contact.PhoneNumber == "" {
			json.NewEncoder(w).Encode("need fill the input!")
		} else {
			createdPerson := db.Create(&contact)
			err = createdPerson.Error
			if err != nil {
				json.NewEncoder(w).Encode(err)
			} else {
				json.NewEncoder(w).Encode(&contact)
			}
		}
	}
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rdb.Del(context.Background(), "contact")

	var contact models.Contact
	db.First(&contact, params["id"])
	if contact.ID == 0 {
		json.NewEncoder(w).Encode("DATA NOT FOUND")
	} else {

		db.Delete(&contact)
		json.NewEncoder(w).Encode(&contact)
	}
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rdb.Del(context.Background(), "contact")

	var contact models.Contact
	var contacttemp models.Contact

	r.ParseForm()
	if r.FormValue("Name") != "" {
		contacttemp.Name = r.FormValue("Name")
	}
	if r.FormValue("PhoneNumber") != "" {
		contacttemp.PhoneNumber = r.FormValue("PhoneNumber")
	}

	if r.FormValue("Name") == "" && r.FormValue("PhoneNumber") == "" {
		json.NewEncoder(w).Encode("need fill the input")
		return
	}

	db.First(&contact, params["id"])

	if contact.ID == 0 {
		json.NewEncoder(w).Encode("DATA NOT FOUND")
	} else {
		db.Model(&contact).Update(contacttemp)
		json.NewEncoder(w).Encode(&contact)
	}
}

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
