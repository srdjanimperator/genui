package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/srdjanimperator/genui/db"
	"github.com/srdjanimperator/genui/model"
	"github.com/srdjanimperator/genui/model/mock"
)

const (
	AddrPort string = ":7001"
)

var (
	Db *db.Connection
)

func routesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto, r.Method, r.Host, r.URL.Port(), r.URL.Path, r.URL.Query())

	routeParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	fmt.Println(routeParts)
	route := ""
	if len(routeParts) > 0 {
		route = routeParts[0]
	}

	switch route {
	case "forms":
		switch r.Method {
		case http.MethodGet:
			if len(routeParts) == 2 {
				formCode := fmt.Sprintf("#%s", routeParts[1])
				mockForms := mock.GetMockForms()
				var requiredForm model.GuiForm
				foundForm := false
				for _, f := range mockForms {
					if f.Code == formCode {
						requiredForm = f
						foundForm = true
					}
				}
				if !foundForm {
					w.WriteHeader(http.StatusNotFound)
					w.Header().Add("Content-Type", "application/text")
					io.WriteString(w, fmt.Sprintf("Form with code '#%s' not found!", formCode))
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(requiredForm)
				return
			} else {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//json.NewEncoder(w).Encode(mock.GetMockForms())
				forms := Db.FormsList()
				json.NewEncoder(w).Encode(forms)
				return
			}
		case http.MethodPost:
			reqBody, _ := ioutil.ReadAll(r.Body)
			var newGuiFormReq model.GuiForm
			json.Unmarshal(reqBody, &newGuiFormReq)
			all := append(mock.GetMockForms(), newGuiFormReq)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(all)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Header().Add("Content-Type", "application/text")
			io.WriteString(w, fmt.Sprintf("Resource not found!"))
			return
		}
	}

}

func main() {

	cfg := &db.DbConnConfig{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		Pwd:    os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}

	Db = db.NewConnection(*cfg)

	http.HandleFunc("/", routesHandler)

	log.Fatal(http.ListenAndServe(AddrPort, nil))
}
