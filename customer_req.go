package main

import (
	"encoding/json"
	"fmt"
	_ "go/token"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

const (
	APP_NAME = "MY-API-GORM"
)


type CustomerHandlerFunction interface{
	getHello() http.HandlerFunc
}

func (shop ServerShop)getHallo() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Hallo to yo and welcomme ")
	}
	
}

func (shop *ServerShop) getCustomers() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		//Execute select Query
	db := shop.DB

	//Read daten
	var customers []Customer
	db.Find(&customers)
	
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	//Encode  slice as Json and write to response
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		return
	}
	}

}

func (shop *ServerShop) getCustomer() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		db := shop.DB

	//Get ID from URL Parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//Create a struct of customers
	var customer = Customer{}
	// through row and add to struct
	db.First(&customer,id)
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Encode  slice as Json and write to response
	if err = json.NewEncoder(w).Encode(customer); err != nil {
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		return
	}
	}

}

func (shop *ServerShop) createCustomers()http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
			//Daten encodet in Json objekten
	customer := Customer{}
	//Encoder les donnees
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Etablir dabor une connection avec la base de donnee
	db := shop.DB
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//Enregistrer les donnees en BD

	db.Create(&Customer{FirstName: 
		customer.FirstName,
		LastName: customer.LastName,
		Email: customer.Email, 
		Phone: customer.Phone, 
		City: customer.City,
		State: customer.State})

	// Set ID field of user struct and encode as JSON
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	
	if err := json.NewEncoder(w).Encode(customer); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	}
}

func (shop *ServerShop) updateCustomer() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		db := shop.DB

	//Get ID from URL Parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var customer Customer;
	//Daten encodet in Json objekten
	Newcustomer := Customer{}
	//Encoder les donnees
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	db.First(&customer,id)

	customer.FirstName = Newcustomer.FirstName
	db.Save(&customer)

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,customer)



}
}

func (shop *ServerShop) deleteCustomer() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		db := shop.DB

	//Get ID from URL Parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//Check if customers exist
	var customer Customer
	db.Delete(&customer,id)
	}
}

const MY_SECRET_KEY = "my-secret-key"

func (shop *ServerShop) conn()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request)  {
		err := shop.openDBCon()
		if err !=nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Erreur lors de la connexion a la base de donnees %v", err)
			fmt.Fprint(w,msg)
		}
		var customers Customer
		db := shop.DB

		//reception des donnees
		customer := &Customer{}
		//Decodages des donnees de utilisateur qui veux se connecter
		err = json.NewDecoder(r.Body).Decode(customer)
		if err != nil {
			fmt.Println(err)
			w.Header().Add("Content-Type","application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"Erreur lors de execution de decodage %v",err)
		}
		//Verifier que utilisateur existe deja 

		err = db.Where("first_name",customer.FirstName).Where("email",customer.Email).First(&customers).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				w.WriteHeader(http.StatusUnauthorized)
				msg := fmt.Sprintf("Aucun utilisateur ne corresond a cette utilisateur %v", err)
				fmt.Fprintln(w,msg)
			}
			return
		
		}
		//faire le test pour verifier existance de utilisateur en Base de donnees
		if customer.FirstName == customers.FirstName && customer.Email == customers.Email {
			w.Header().Add("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w,"Welcomme to my api \n")

			//Creation de objet Claims
			claims := jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Issuer: APP_NAME,
			}
			customClaims := struct{
				FirstName string `gorm:"first_name"`
				Email string `gorm:"email"`
				jwt.StandardClaims
			}{
				FirstName: customer.FirstName,
				Email: customer.Email,
				StandardClaims: claims,
			}
			// Appel de la methode NewWithClaims() --> qui prend en entre une methode de signature et  un objet jwt.Claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256,customClaims)
			tokenF, err := token.SignedString([]byte(MY_SECRET_KEY))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				msg := fmt.Sprintf("Erreur lors de ncodage du Token en byte %v", err)
				fmt.Fprint(w, msg)
			} 

			fmt.Println(time.Now().Add(time.Hour * 24).Unix())
			fmt.Fprint(w, tokenF)

			//Enregistrer le token en Base de donnees
			
			db.Model(&customers).Where("first_name",customer.FirstName).Update("token",tokenF)
			

		}else{
			w.Header().Add("Content-Type","application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"Erreur de connexion mauvais mot de passe ou nom d'utilisateur")
		}
		
	}
	
	
}

