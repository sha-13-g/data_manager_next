package main

import (
	"context"
	"data_manager/data"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-session/session/v3"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	session, err := session.Start(context.Background(), w, r)
	if err != nil {
		fmt.Fprint(w, err)
	}

	key, ok := session.Get("user")

	if ok {
		user6 := fmt.Sprintf("%d", key)
		user64, err := strconv.ParseInt(user6, 10, 64)
		user := int(user64)

		if err != nil {
			fmt.Println(err)
		}

		file, _ := template.ParseFiles("views/dashboard.html")

		Username := data.GetCommercial(user)
		CustomersDetails := data.GetCustomersDetails()
		Users := data.GetUsers()
		Commercials := data.GetCommercials()
		Services := data.GetServices()
		Sites := data.GetSites()
		Customers := data.GetCustomers()
		Projects := data.GetProjectByAll()
		Volumes := data.GetVolumes()
		Incidents := data.GetIncidentTable()
		Prospect := data.GetProspects()

		file.Execute(w, map[string]any{
			"CustomersDetails": CustomersDetails,
			"Users":            Users,
			"Commercials":      Commercials,
			"Incidents":        Incidents,
			"Services":         Services,
			"Customers":        Customers,
			"Projects":         Projects,
			"Volumes":          Volumes,
			"Sites":            Sites,
			"Username":         Username,
			"Prospect":         Prospect,
		})
	} else {
		fmt.Println("Deconnecte")
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}
