package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type addr struct {
	Area    string `json:"area"`
	Country string `json:"country"`
}

type basicDetails struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address addr   `json:"address"`
}

type basicDetsArray struct {
	BasicDetailsArray []basicDetails `json:"basicDetailsArray"`
}

type techdet struct {
	Tech string  `json:"tech"`
	Exp  float64 `json:"exp"`
}

type techDetails struct {
	Id       int       `json:"id"`
	TechDets []techdet `json:"techDets"`
}

type techDetsArray struct {
	TechDetailsArray []techDetails `json:"techDetailsArray"`
}

type contactdet struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type contactDetails struct {
	Id          int        `json:"id"`
	ContactDets contactdet `json:"contactDets"`
}

type contactDetsArray struct {
	ContactDetailsArray []contactDetails `json:"contactDetailsArray"`
}

type totalDetails struct {
	name     string
	address  addr
	techDets []techdet
	email    string
	phone    string
}

func main() {

	// emp_obj := Employee{Name:"Rachel", Age:24, Salary :344444}
	// emp, _ := json.Marshal(emp_obj)
	// fmt.Println(string(emp))
	//op : {"Name":"Rachel","Age":24,"Salary":344444}

	basicjsonFile, err := os.Open("basicDetailsJSON.json")
	if err != nil {
		fmt.Println(err)
	}

	defer basicjsonFile.Close()

	basicbyteValue, _ := ioutil.ReadAll(basicjsonFile)

	// var result map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &result)
	//fmt.Println(result)

	var basicDetailsArray basicDetsArray
	json.Unmarshal(basicbyteValue, &basicDetailsArray)
	// fmt.Println(basicDetailsArray)

	techjsonFile, err := os.Open("techDetailsJSON.json")
	if err != nil {
		fmt.Println(err)
	}

	techbyteValue, _ := ioutil.ReadAll(techjsonFile)

	var techDetailsArray techDetsArray
	json.Unmarshal(techbyteValue, &techDetailsArray)
	// fmt.Println(techDetailsArray)

	contactjsonFile, err := os.Open("contactDetailsJSON.json")
	if err != nil {
		fmt.Println(err)
	}

	contactbyteValue, _ := ioutil.ReadAll(contactjsonFile)

	var contactDetailsArray contactDetsArray
	json.Unmarshal(contactbyteValue, &contactDetailsArray)
	// fmt.Println(contactDetailsArray)

	countryDetails := make(map[string]string)
	countryDetails["IND"] = "+91"
	countryDetails["US"] = "+41"

	totalDetailsArray := make(map[string]totalDetails, 3)

	for i := range basicDetailsArray.BasicDetailsArray {
		var temp totalDetails
		id := basicDetailsArray.BasicDetailsArray[i].Id
		temp.name = basicDetailsArray.BasicDetailsArray[i].Name
		temp.address = basicDetailsArray.BasicDetailsArray[i].Address
		country := basicDetailsArray.BasicDetailsArray[i].Address.Country
		for j := 0; j < len(techDetailsArray.TechDetailsArray); j++ {
			if techDetailsArray.TechDetailsArray[j].Id == id {
				temp.techDets = techDetailsArray.TechDetailsArray[j].TechDets
			}
		}
		for j := 0; j < len(contactDetailsArray.ContactDetailsArray); j++ {
			if contactDetailsArray.ContactDetailsArray[j].Id == id {
				temp.email = contactDetailsArray.ContactDetailsArray[j].ContactDets.Email
				temp.phone = contactDetailsArray.ContactDetailsArray[j].ContactDets.Phone
			}
		}
		for j := range countryDetails {
			if j == country {
				temp.phone = countryDetails[j] + " " + temp.phone
			}
		}
		totalDetailsArray[strconv.Itoa(id)] = temp
	}
	// fmt.Println(totalDetailsArray, "\n")
	for i := range totalDetailsArray {
		fmt.Print(i, ":", totalDetailsArray[i], "\n")
	}
}
