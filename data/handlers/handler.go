package handlers

import (
	"data_manager/data"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func AddNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var site data.Site

	if r.FormValue("site-id") != "" {
		id := r.FormValue("site-id")

		site = data.GetSite(id)

	} else {
		sites := data.GetSites()
		site = sites[len(sites)-1]
	}

	number_id := fmt.Sprint(time.Now().Unix())
	customer_number := r.FormValue("customer-number")

	number := data.Number{
		Id:      number_id,
		Number:  customer_number,
		Site_id: site.Id,
	}

	number.Add()
}

func AddSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var customer data.Customer

	if r.FormValue("customer-id") != "" {
		id, _ := strconv.ParseInt(r.FormValue("customer-id"), 10, 64)

		customer = data.GetCustomer(int(id))

	} else {
		customers := data.GetCustomers()
		customer = customers[len(customers)-1]
	}

	id := fmt.Sprint(time.Now().Unix())
	name := r.FormValue("site-name")

	site := data.Site{
		Id:          id,
		Name:        name,
		Customer_id: fmt.Sprint(customer.Id),
	}

	site.Add()
}

func AddRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number
	d := r.FormValue("data")
	id64, _ := strconv.ParseInt(d, 10, 64)

	volume_id := int(id64)

	if r.FormValue("number-id") != "" {
		number = data.GetNumber(r.FormValue("number-id"))
	} else {
		numbers := data.GetNumbers()
		number = numbers[len(numbers)-1]
	}

	id := fmt.Sprint(time.Now().Unix())

	volume := data.GetVolume(volume_id)

	date_re, _ := time.Parse("2006-01-02", r.FormValue("date-re"))
	date_exp := date_re.AddDate(0, 1, 1)
	auto_re, _ := strconv.ParseBool(r.FormValue("auto-re"))

	recharge := data.Recharge{
		Id:        id,
		Volume:    volume.Volume,
		Date_re:   date_re,
		Date_exp:  date_exp,
		Auto_re:   auto_re,
		Number_id: number.Id,
		Volume_id: volume_id,
	}

	recharge.Add()
}

func DeleteRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recharge data.Recharge
	recharge.Id = r.FormValue("recharge-id")
	v, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	recharge.Volume = int(v)
	recharge.Date_re, _ = time.Parse("2006-01-02", r.FormValue("date-re"))
	recharge.Date_exp = recharge.Date_re.AddDate(0, 1, 1)

	data.DeleteRecharge(recharge.Id)
}

func UpdateRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recharge data.Recharge
	recharge.Id = r.FormValue("recharge-id")
	v, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	recharge.Volume = int(v)
	recharge.Date_re, _ = time.Parse("2006-01-02", r.FormValue("date-re"))
	recharge.Date_exp = recharge.Date_re.AddDate(0, 1, 1)

	recharge.Update()
}

func DeleteIncident(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var incident data.Incident
	incident.Id = r.FormValue("incident-id")

	data.DeleteIncident(incident.Id)

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}

func DeleteNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number
	number.Id = r.FormValue("number-id")
	number.Number = r.FormValue("customer-number")

	data.DeleteNumber(number.Id)

}

func UpdateNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number
	number.Id = r.FormValue("number-id")
	number.Number = r.FormValue("customer-number")

	number.Update()
}

func DeleteSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var site data.Site
	site.Id = r.FormValue("site-id")
	site.Name = r.FormValue("site-name")

	data.DeleteSite(site.Id)
}
func UpdateSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var site data.Site
	site.Id = r.FormValue("site-id")
	site.Name = r.FormValue("site-name")

	site.Update()
}

func CustomerForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/customer_form.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

func ShowCustomer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/customer.html")

	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	customer := data.GetCustomer(int(id))

	t.Execute(w, customer)
}

func ShowCustomers(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/dashboard.html")

	if err != nil {
		panic(err)
	}
	data := data.GetCustomersDetails()

	t.Execute(w, data)
}

func UpdateIncident(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := r.FormValue("incident-id")
	date, _ := time.Parse("2006-01-02", r.FormValue("incident-date"))

	customer := data.GetCustomerByName(r.FormValue("incident-origin"))
	commercial := data.GetCommercialByName(r.FormValue("incident-intervenant"))

	site := data.GetSiteByName(r.FormValue("incident-site"))
	description := r.FormValue("incident-description")
	responsability := r.FormValue("incident-responsability")

	incident := data.Incident{
		Id:             id,
		Description:    description,
		Origin:         customer.Id,
		Responsability: responsability,
		Date:           date,
		CustomerId:     customer.Id,
		AgentId:        commercial.Id,
		Site_id:        site.Id,
	}

	incident.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddIncident(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	customer, _ := strconv.ParseInt(r.FormValue("customer-name"), 10, 64)
	customerId := int(customer)

	site := r.FormValue("site-name")
	description := r.FormValue("description")
	responsability := r.FormValue("responsability")
	agent64, _ := strconv.ParseInt(r.FormValue("agent-name"), 10, 42)
	agent := int(agent64)

	incident := data.Incident{
		Id:             `DEFAULT`,
		Description:    description,
		Origin:         customerId,
		Responsability: responsability,
		Date:           date,
		CustomerId:     customerId,
		AgentId:        agent,
		Site_id:        site,
	}

	incident.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)

}

func AddProspect(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	name := r.FormValue("prospect-name")
	address := r.FormValue("prospect-address")
	contact := r.FormValue("prospect-contact")
	referant := r.FormValue("referant-name")
	contactReferant := r.FormValue("referant-contact")
	mail_object := r.FormValue("mail-object")
	ddp, _ := time.Parse("2006-01-02", r.FormValue("ddp"))
	drc := ddp.AddDate(0, 0, 7)

	prospect := data.Prospect{
		Id:               0,
		Name:             name,
		Address:          address,
		Referant:         referant,
		Referant_contact: contactReferant,
		Contact:          contact,
		Mail_object:      mail_object,
		Ddp:              ddp.Format("2006-01-02"),
		Ddc:              drc.Format("2006-01-02"),
	}

	prospect.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddCommercial(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	commercial := data.Commercial{
		Id:      0,
		Name:    r.FormValue("name"),
		Role:    r.FormValue("role"),
		Address: r.FormValue("address"),
	}

	commercial.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddProject(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	commercial, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	service, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)

	service_id := int(service)
	commercial_id := int(commercial)

	project := data.Project{
		Id:            0,
		Name:          r.FormValue("project-name"),
		Description:   r.FormValue("project-description"),
		Customer_id:   r.FormValue("customer-id"),
		Commercial_id: commercial_id,
		Service_id:    service_id,
	}

	project.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	commercialId, _ := strconv.ParseInt(r.FormValue("commercial-id"), 10, 64)

	customer := data.Customer{
		Id:            0,
		Logo:          r.FormValue("customer-logo"),
		Name:          r.FormValue("customer-name"),
		Address:       r.FormValue("customer-address"),
		Contact:       r.FormValue("customer-contact"),
		Branch:        r.FormValue("customer-branch"),
		Commercial_id: int(commercialId),
	}
	customer.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddService(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	service := data.Service{
		Id:   0,
		Name: r.FormValue("service-name"),
	}

	service.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	idc, _ := strconv.ParseInt(r.FormValue("commercial-id"), 10, 64)
	commercial_id := int(idc)

	user := data.User{
		Id:            0,
		Email:         r.FormValue("user-email"),
		Password:      r.FormValue("user-password"),
		Commercial_id: commercial_id,
	}

	user.Add()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("user-id"))), 10, 64)
	id := int(id64)

	data.DeleteUser(id)

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteCommercial(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	id := int(id64)

	fmt.Println(r.PostForm)
	data.DeleteCommercial(id)

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)
	id := int(id64)

	data.DeleteService(id)

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("project-id"))), 10, 64)
	id := int(id64)

	data.DeleteProject(id)

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("customer-id"))), 10, 64)
	id := int(id64)
	var numbers [][]data.Number

	sites := data.GetSitesByCustomer(r.FormValue("customer-id"))

	for _, site := range sites {
		number := data.GetNumberBySite(site.Id)
		numbers = append(numbers, number)
	}

	for _, number := range numbers {
		for _, n := range number {
			data.DeleteRechargesByNumber(n.Id)
		}
	}

	for _, site := range sites {
		data.DeleteNumberBySite(site.Id)
	}

	data.DeleteSitesByCustomer(fmt.Sprint(id))
	data.DeleteCustomer(id)

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("customer-id"))), 10, 64)
	id := int(id64)

	commercialId, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)

	customer := data.Customer{
		Id:            id,
		Logo:          r.FormValue("customer-logo"),
		Name:          r.FormValue("customer-name"),
		Address:       r.FormValue("customer-address"),
		Contact:       r.FormValue("customer-contact"),
		Branch:        r.FormValue("customer-branch"),
		Commercial_id: int(commercialId),
	}

	customer.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("user-id"))), 10, 64)
	id := int(id64)

	user := data.User{
		Id:       id,
		Email:    r.FormValue("user-name"),
		Password: r.FormValue("user-password"),
	}

	user.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("project-id"))), 10, 64)
	commercial, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	service, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)

	id := int(id64)
	service_id := int(service)
	commercial_id := int(commercial)

	project := data.Project{
		Id:            id,
		Name:          r.FormValue("project-name"),
		Description:   r.FormValue("project-description"),
		Customer_id:   r.FormValue("customer-name"),
		Commercial_id: commercial_id,
		Service_id:    service_id,
	}

	data.UpdateCommercialId(commercial_id)
	project.Update()
	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)
	id := int(id64)

	service := data.Service{
		Id:   id,
		Name: r.FormValue("service-name"),
	}

	service.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateCommercial(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	id := int(id64)
	fmt.Print(id)

	commercial := data.Commercial{
		Id:      id,
		Name:    r.FormValue("commercial-name"),
		Role:    r.FormValue("commercial-role"),
		Address: r.FormValue("commercial-address"),
	}

	commercial.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func UpdateProspect(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	name := r.FormValue("prospect-name")
	id64, _ := strconv.ParseInt(r.FormValue("prospect-id"), 10, 64)
	id := int(id64)
	address := r.FormValue("prospect-address")
	contact := r.FormValue("prospect-contact")
	referant := r.FormValue("referant-name")
	contactReferant := r.FormValue("referant-contact")
	mail_object := r.FormValue("mail-object")
	ddp, _ := time.Parse("2006-01-02", r.FormValue("ddp"))
	drc := ddp.AddDate(0, 0, 7)

	prospect := data.Prospect{
		Id:               id,
		Name:             name,
		Address:          address,
		Referant:         referant,
		Referant_contact: contactReferant,
		Contact:          contact,
		Mail_object:      mail_object,
		Ddp:              ddp.Format("2006-01-02"),
		Ddc:              drc.Format("2006-01-02"),
	}

	prospect.Update()

	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}

func DeleteProspect(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(r.FormValue("prospect-id"), 10, 64)
	id := int(id64)

	data.DeleteProspect(id)

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}
