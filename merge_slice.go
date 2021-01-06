package main

import (
	"fmt"
	"strconv"
)

type addr struct {
	area    string
	country string
}

type basicDetails struct {
	id      int
	name    string
	address addr
}

type techdet struct {
	tech string
	exp  float64
}

type techDetails struct {
	id       int
	techDets []techdet
}

type contactdet struct {
	email string
	phone string
}

type contactDetails struct {
	id          int
	contactDets contactdet
}

type totalDetails struct {
	name     string
	address  addr
	techDets []techdet
	email    string
	phone    string
}

func main() {
	basicDetailsArray := make([]basicDetails, 3)
	basicDetailsArray[0] = basicDetails{id: 1, name: "faizal", address: addr{area: "Vejalpur", country: "IND"}}
	basicDetailsArray[1] = basicDetails{id: 2, name: "nisarg", address: addr{area: "Jivrajpark", country: "US"}}
	basicDetailsArray[2] = basicDetails{id: 3, name: "parth", address: addr{area: "Satellite", country: "US"}}

	techDetailsArray := make([]techDetails, 3)
	techDetailsArray[0] = techDetails{id: 1, techDets: []techdet{{tech: "PHP", exp: 1}, {tech: "Node", exp: 0.5}}}
	techDetailsArray[1] = techDetails{id: 2, techDets: []techdet{{tech: "JS", exp: 0.5}}}
	techDetailsArray[2] = techDetails{id: 3, techDets: []techdet{{tech: "Flutter", exp: 0.7}, {tech: "React", exp: 0.3}}}

	contactDetailsArray := make([]contactDetails, 3)
	contactDetailsArray[0] = contactDetails{id: 1, contactDets: contactdet{email: "maulvifaizal@gmail.com", phone: "8488982986"}}
	contactDetailsArray[1] = contactDetails{id: 2, contactDets: contactdet{email: "nisargchokshi4000@gmail.com", phone: "9624904521"}}
	contactDetailsArray[2] = contactDetails{id: 3, contactDets: contactdet{email: "parthpv99@gmail.com", phone: "8429958754"}}

	countryDetails := make(map[string]string)
	countryDetails["IND"] = "+91"
	countryDetails["US"] = "+41"

	totalDetailsArray := make(map[string]totalDetails, 3)

	for i := range basicDetailsArray {
		var temp totalDetails
		id := basicDetailsArray[i].id
		temp.name = basicDetailsArray[i].name
		temp.address = basicDetailsArray[i].address
		country := basicDetailsArray[i].address.country
		for j := 0; j < len(techDetailsArray); j++ {
			if techDetailsArray[j].id == id {
				temp.techDets = techDetailsArray[j].techDets
			}
		}
		for j := 0; j < len(contactDetailsArray); j++ {
			if contactDetailsArray[j].id == id {
				temp.email = contactDetailsArray[j].contactDets.email
				temp.phone = contactDetailsArray[j].contactDets.phone
			}
		}
		for j := range countryDetails {
			if j == country {
				temp.phone = countryDetails[j] + " " + temp.phone
			}
		}
		totalDetailsArray[strconv.Itoa(id)] = temp
	}
	fmt.Println(totalDetailsArray, "\n")
	for i := range totalDetailsArray {
		fmt.Print(i, ":", totalDetailsArray[i], "\n")
	}
}
