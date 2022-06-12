package main

import (
	"api-contact/config"
	"api-contact/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error
var rdb *redis.Client
var key string = "contact"

func main() {
	db = config.ConnectDatabase()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                              // All origins
		AllowedMethods: []string{"POST", "DELETE", "GET", "PATCH"}, // Allowing only get, just an example
	})

	var redisHost = os.Getenv("REDISHOST")
	var redisPassword = os.Getenv("PASSWORDREDIS")

	rdb = newRedisClient(redisHost, redisPassword)

	router := mux.NewRouter()

	router.HandleFunc("/contact", GetContact).Methods("GET")
	router.HandleFunc("/contact/{id}", GetContactById).Methods("GET")
	router.HandleFunc("/contact", AddContact).Methods("POST")
	router.HandleFunc("/contact/{id}", DeleteContact).Methods("DELETE")
	router.HandleFunc("/contact/{id}", UpdateContact).Methods("PATCH")
	router.Use(mux.CORSMethodMiddleware(router))

	// http.ListenAndServe(":8080", router)
	// c.Handler(router)

	http.ListenAndServe(":8080", c.Handler(router))
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	var contact []models.Contact
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var contact models.Contact
	db.First(&contact, params["id"])

	contactRedis, _ := rdb.Get(context.Background(), key).Result()
	if contactRedis == "" {
		fmt.Println("from database")
		if contact.ID == 0 {
			json.NewEncoder(w).Encode("DATA NOT FOUND")
		} else {
			json.NewEncoder(w).Encode(&contact)
		}
	} else {
		fmt.Println("from redis")
		op2, _ := rdb.Get(context.Background(), key).Result()
		var cont []models.Contact
		json.Unmarshal([]byte(op2), &cont)
		idParams, _ := strconv.Atoi(params["id"])
		for _, v := range cont {
			if v.ID == uint(idParams) {
				json.NewEncoder(w).Encode(v)
				return
			}
		}
		json.NewEncoder(w).Encode("DATA NOT FOUND")
	}
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	w.Header().Set("Access-Control-Allow-Origin", "*")

	headerContentType := r.Header.Get("Content-Type")
	rdb.Del(context.Background(), "contact")

	if headerContentType == "application/x-www-form-urlencoded" {
		r.ParseForm()
		contact.Name = r.FormValue("Name")
		contact.PhoneNumber = r.FormValue("PhoneNumber")
		contact.Email = r.FormValue("Email")
		if contact.Name == "" || contact.PhoneNumber == "" {
			json.NewEncoder(w).Encode("need fill the input!1")
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
			json.NewEncoder(w).Encode("need fill the input!2")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rdb.Del(context.Background(), "contact")

	var contact models.Contact
	var contacttemp models.Contact

	headerContentType := r.Header.Get("Content-Type")

	if headerContentType == "application/x-www-form-urlencoded" {
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
	} else {
		json.NewDecoder(r.Body).Decode(&contacttemp)
		if contacttemp.Name == "" || contacttemp.PhoneNumber == "" {
			json.NewEncoder(w).Encode("need fill the input!2")
			return
		}
		db.First(&contact, params["id"])
		if contact.ID == 0 {
			json.NewEncoder(w).Encode("DATA NOT FOUND")
			return
		}
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
