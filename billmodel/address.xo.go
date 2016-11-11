// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Address represents a row from 'equinox.address'.
type Address struct {
	Addpostcode     sql.NullString `json:"addpostcode"`     // addpostcode
	Addroadnumber   sql.NullString `json:"addroadnumber"`   // addroadnumber
	Addsubpremises  sql.NullString `json:"addsubpremises"`  // addsubpremises
	Addroadname     sql.NullString `json:"addroadname"`     // addroadname
	Addlocality     sql.NullString `json:"addlocality"`     // addlocality
	Addposttown     sql.NullString `json:"addposttown"`     // addposttown
	Addcounty       sql.NullString `json:"addcounty"`       // addcounty
	Addalk          sql.NullString `json:"addalk"`          // addalk
	Addmpr          sql.NullString `json:"addmpr"`          // addmpr
	Addmpan         sql.NullString `json:"addmpan"`         // addmpan
	Addcli          sql.NullString `json:"addcli"`          // addcli
	Addaccountno    sql.NullString `json:"addaccountno"`    // addaccountno
	Addlastupdate   pq.NullTime    `json:"addlastupdate"`   // addlastupdate
	Addlastuptime   pq.NullTime    `json:"addlastuptime"`   // addlastuptime
	Addlastupdateby sql.NullString `json:"addlastupdateby"` // addlastupdateby
	Addnotes        sql.NullInt64  `json:"addnotes"`        // addnotes
	Addline1        sql.NullString `json:"addline1"`        // addline1
	Addline2        sql.NullString `json:"addline2"`        // addline2
	Addline3        sql.NullString `json:"addline3"`        // addline3
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllAddress(db XODB, callback func(x Address) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`addpostcode, addroadnumber, addsubpremises, addroadname, addlocality, addposttown, addcounty, addalk, addmpr, addmpan, addcli, addaccountno, addlastupdate, addlastuptime, addlastupdateby, addnotes, addline1, addline2, addline3, equinox_lrn, equinox_sec ` +
		`FROM equinox.address `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Address{}

		// scan
		err = q.Scan(&a.Addpostcode, &a.Addroadnumber, &a.Addsubpremises, &a.Addroadname, &a.Addlocality, &a.Addposttown, &a.Addcounty, &a.Addalk, &a.Addmpr, &a.Addmpan, &a.Addcli, &a.Addaccountno, &a.Addlastupdate, &a.Addlastuptime, &a.Addlastupdateby, &a.Addnotes, &a.Addline1, &a.Addline2, &a.Addline3, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// AddressByEquinoxLrn retrieves a row from 'equinox.address' as a Address.
//
// Generated from index 'address_pkey'.
func AddressByEquinoxLrn(db XODB, equinoxLrn int64) (*Address, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`addpostcode, addroadnumber, addsubpremises, addroadname, addlocality, addposttown, addcounty, addalk, addmpr, addmpan, addcli, addaccountno, addlastupdate, addlastuptime, addlastupdateby, addnotes, addline1, addline2, addline3, equinox_lrn, equinox_sec ` +
		`FROM equinox.address ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Address{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Addpostcode, &a.Addroadnumber, &a.Addsubpremises, &a.Addroadname, &a.Addlocality, &a.Addposttown, &a.Addcounty, &a.Addalk, &a.Addmpr, &a.Addmpan, &a.Addcli, &a.Addaccountno, &a.Addlastupdate, &a.Addlastuptime, &a.Addlastupdateby, &a.Addnotes, &a.Addline1, &a.Addline2, &a.Addline3, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
