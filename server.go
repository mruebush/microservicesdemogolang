package customerpref

import(
	// "fmt"
	"github.com/go-martini/martini"
)

type CustomerPref struct {
	Id string
	NotificationPref string
}

func StartServer(){
	m := martini.Classic()
  	m.Get("/CustomerPreferences", func() string {
    	return GetData()
    	//return "hello world"
  	})

  	m.Post("/CustomerPreferences", func() string {
  		cpf := CustomerPref{Id: "23476", NotificationPref: "pidgeon"}
  		AddCustomerPref(cpf)
  		SendUpdateCustomerPref(cpf.Id, cpf.NotificationPref)
  		return "Added"

  	})

  	m.Put("/CustomerPreferences/{id}", func() string {
  		cpf := CustomerPref{Id: "097687", NotificationPref: "telegraph"}
  		UpdateCustomerPref(cpf)
  		SendUpdateCustomerPref(cpf.Id, cpf.NotificationPref)
  		return "Updated"
  	})

  	m.Run()
}
