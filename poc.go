package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jawher/mow.cli"
	_ "github.com/lib/pq"
	"github.com/utilitywarehouse/poc-pg-mirror/billmodel"
)

func main() {
	app := cli.App("poc-pg-mirror", "postgres mirror demo")

	user := app.String(cli.StringOpt{
		Name:   "pguser",
		Desc:   "Postgres user",
		Value:  "postgres",
		EnvVar: "PG_USER",
	})
	pass := app.String(cli.StringOpt{
		Name:      "password",
		Desc:      "Postgres password",
		EnvVar:    "PG_PASS",
		HideValue: true,
	})
	host := app.String(cli.StringOpt{
		Name:   "pghost",
		Desc:   "Postgres host",
		Value:  "new-build.postgres.dev.uw.systems",
		EnvVar: "PG_HOST",
	})
	dbname := app.String(cli.StringOpt{
		Name:   "pgdb",
		Desc:   "Postgres database",
		Value:  "new_build",
		EnvVar: "PG_DB",
	})
	port := app.Int(cli.IntOpt{
		Name:   "pgport",
		Desc:   "Postgres port",
		Value:  5432,
		EnvVar: "PG_PORT",
	})

	app.Action = func() {
		runServer(*user, *pass, *host, *port, *dbname)
	}

	app.Run(os.Args)
}

func runServer(user, pass, host string, port int, dbname string) {

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", user, pass, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}

	h := handlers{db}

	r := mux.NewRouter()
	r.HandleFunc("/customers/{id}", h.custHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Listening on port 8080.  Try: curl http://localhost:8080/customers/0000001")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type handlers struct {
	db *sql.DB
}

func (h *handlers) custHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	cust, err := billmodel.CustomersByCustaccountno(h.db, sql.NullString{String: id, Valid: true})
	if err != nil {
		log.Fatal(err)
	}

	switch len(cust) {
	case 0:
		resp.WriteHeader(http.StatusNotFound)
		return
	case 1:
	default:
		resp.WriteHeader(http.StatusInternalServerError)
	}

	rc := cust[0]

	c := Customer{
		Title:     rc.Custconttitle.String,
		Initials:  rc.Custcontinitials.String,
		Name:      rc.Custname.String,
		Firstname: rc.Custcontfirstnam.String,
		Surname:   rc.Custcontsurname.String,
		DOB:       &rc.Custdob.Time,
		Address: &Address{
			Address:  []string{rc.Custadd1.String, rc.Custadd2.String, rc.Custadd3.String, rc.Custadd4.String, rc.Custcounty.String},
			Postcode: rc.Custpostcode.String,
		},
		Phone:  rc.Custphone.String,
		Fax:    rc.Custfax.String,
		Mobile: rc.Custmobile.String,
		Email:  rc.Custemail.String,
		//TODO: invoice & the rest
	}

	resp.Header().Add("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(resp)
	enc.SetIndent("  ", "  ")
	if err := enc.Encode(c); err != nil {
		log.Fatal(err)
	}
}
