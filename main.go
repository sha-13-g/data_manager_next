package main

import (
	"context"
	"data_manager/data"
	"data_manager/data/handlers"
	"fmt"
	"text/template"
	"time"

	"github.com/go-session/session/v3"
	"github.com/gocarina/gocsv"

	//"encoding/csv"
	//"github.com/mohae/struct2csv"
	"net/http"
	"os"
)

type Excel struct {
	Number string `csv:"number"`
	Volume string `csv:"volume"`
	Price  string `csv:"price"`
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, _ := template.ParseFiles("views/login.html")

		file.Execute(w, map[int]string{
			0: "",
		})
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		key, _ := store.Get("user")
		k := fmt.Sprintf("%s", key)
		store.Delete(k)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		users := data.GetUsers()

		for _, user := range users {
			if username == user.Email && password == user.Password {
				duration := 24 * time.Hour
				ctx, c := context.WithTimeout(context.Background(), duration)
				defer c()

				session, err := session.Start(ctx, w, r)

				if err != nil {
					fmt.Fprintln(w, err)
					return
				}

				session.Set("user", user.Commercial_id)
				err = session.Save()

				if err != nil {
					fmt.Fprintln(w, err)
					return
				}

				http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
			}
		}
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)

	})

	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/data/update-incident", handlers.UpdateIncident)
	http.HandleFunc("/data/delete-incident", handlers.DeleteIncident)

	http.HandleFunc("/data/delete-prospect", handlers.DeleteProspect)
	http.HandleFunc("/data/update-prospect", handlers.UpdateProspect)

	http.HandleFunc("/data/customer-form", handlers.CustomerForm)
	http.HandleFunc("/data/add-prospect", handlers.AddProspect)
	http.HandleFunc("/data/add-data-customer", func(w http.ResponseWriter, r *http.Request) {

		handlers.AddSite(w, r)
		handlers.AddNumber(w, r)
		handlers.AddRecharge(w, r)

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/data/show-customer/{id}", handlers.ShowCustomer)
	http.HandleFunc("/data/update-data-customer", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		handlers.UpdateCustomer(w, r)
		handlers.UpdateSite(w, r)
		handlers.UpdateNumber(w, r)
		handlers.UpdateRecharge(w, r)

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/data/delete-data-customer", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		handlers.DeleteRecharge(w, r)
		handlers.DeleteNumber(w, r)
		handlers.DeleteSite(w, r)

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/data/add-incident", handlers.AddIncident)

	http.HandleFunc("/data/add-commercial", handlers.AddCommercial)
	http.HandleFunc("/data/add-user", handlers.AddUser)
	http.HandleFunc("/data/add-service", handlers.AddService)
	http.HandleFunc("/data/add-customer", handlers.AddCustomer)
	http.HandleFunc("/data/add-project", handlers.AddProject)

	http.HandleFunc("/data/delete-customer", handlers.DeleteCustomer)
	http.HandleFunc("/data/delete-project", handlers.DeleteProject)
	http.HandleFunc("/data/delete-service", handlers.DeleteService)
	http.HandleFunc("/data/delete-commercial", handlers.DeleteCommercial)
	http.HandleFunc("/data/delete-user", handlers.DeleteUser)

	http.HandleFunc("/data/update-customer", handlers.UpdateCustomer)
	http.HandleFunc("/data/update-project", handlers.UpdateProject)
	http.HandleFunc("/data/update-service", handlers.UpdateService)
	http.HandleFunc("/data/update-commercial", handlers.UpdateCommercial)
	http.HandleFunc("/data/update-user", handlers.UpdateUser)

	http.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		db := data.InitConn()
		defer db.Close()

		query := fmt.Sprint(`select n.number, v.volume, v.price from number n, volume v, recharge 
		where n.id = recharge.number_id and v.id = recharge.volume_id 
		and floor(extract(epoch FROM (date_exp - now()))/3600) <=58;`)

		rows, err := db.Query(query)

		var excel Excel
		var data []Excel

		if err != nil {
			panic(err)
		}

		for rows.Next() {

			rows.Scan(&excel.Number, &excel.Volume, &excel.Price)
			data = append(data, excel)

		}

		rechargeFile, err := os.OpenFile("recharges.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

		defer rechargeFile.Close()

		if err != nil {
			panic(err)
		}

		e := gocsv.MarshalFile(&data, rechargeFile) // Get all clients as CSV string

		if e != nil {
			panic(err)
		}

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)

	})

	http.ListenAndServe(":5000", nil)
}
