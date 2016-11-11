// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Homedel represents a row from 'equinox.homedel'.
type Homedel struct {
	Hdcustref     sql.NullString `json:"hdcustref"`     // hdcustref
	Hdname        sql.NullString `json:"hdname"`        // hdname
	Hdaddress1    sql.NullString `json:"hdaddress1"`    // hdaddress1
	Hdaddress2    sql.NullString `json:"hdaddress2"`    // hdaddress2
	Hdaddress3    sql.NullString `json:"hdaddress3"`    // hdaddress3
	Hdaddress4    sql.NullString `json:"hdaddress4"`    // hdaddress4
	Hdaddress5    sql.NullString `json:"hdaddress5"`    // hdaddress5
	Hdpostcode    sql.NullString `json:"hdpostcode"`    // hdpostcode
	Hdcontact     sql.NullString `json:"hdcontact"`     // hdcontact
	Hdtelephone   sql.NullString `json:"hdtelephone"`   // hdtelephone
	Hddelinstruct sql.NullString `json:"hddelinstruct"` // hddelinstruct
	Hdsenderacc   sql.NullString `json:"hdsenderacc"`   // hdsenderacc
	Hdpacks       sql.NullString `json:"hdpacks"`       // hdpacks
	Hdweight      sql.NullString `json:"hdweight"`      // hdweight
	Hddespdate    pq.NullTime    `json:"hddespdate"`    // hddespdate
	Hdservicecode sql.NullString `json:"hdservicecode"` // hdservicecode
	Hdtype        sql.NullString `json:"hdtype"`        // hdtype
	Hdselect      sql.NullInt64  `json:"hdselect"`      // hdselect
	Hddeltype     sql.NullString `json:"hddeltype"`     // hddeltype
	Hdtypeofpack  sql.NullString `json:"hdtypeofpack"`  // hdtypeofpack
	Hdemail       sql.NullString `json:"hdemail"`       // hdemail
	Hdsparen1     sql.NullInt64  `json:"hdsparen1"`     // hdsparen1
	Hdsparen2     sql.NullInt64  `json:"hdsparen2"`     // hdsparen2
	Hdsparen3     sql.NullInt64  `json:"hdsparen3"`     // hdsparen3
	Hdspared1     pq.NullTime    `json:"hdspared1"`     // hdspared1
	Hdspared2     pq.NullTime    `json:"hdspared2"`     // hdspared2
	Hdspared3     pq.NullTime    `json:"hdspared3"`     // hdspared3
	Hdsparel1     sql.NullInt64  `json:"hdsparel1"`     // hdsparel1
	Hdsparel2     sql.NullInt64  `json:"hdsparel2"`     // hdsparel2
	Hdsparel3     sql.NullInt64  `json:"hdsparel3"`     // hdsparel3
	Hdprinted     pq.NullTime    `json:"hdprinted"`     // hdprinted
	Hdsparec4     sql.NullString `json:"hdsparec4"`     // hdsparec4
	Hdsparec5     sql.NullString `json:"hdsparec5"`     // hdsparec5
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

func AllHomedel(db XODB, callback func(x Homedel) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`hdcustref, hdname, hdaddress1, hdaddress2, hdaddress3, hdaddress4, hdaddress5, hdpostcode, hdcontact, hdtelephone, hddelinstruct, hdsenderacc, hdpacks, hdweight, hddespdate, hdservicecode, hdtype, hdselect, hddeltype, hdtypeofpack, hdemail, hdsparen1, hdsparen2, hdsparen3, hdspared1, hdspared2, hdspared3, hdsparel1, hdsparel2, hdsparel3, hdprinted, hdsparec4, hdsparec5, equinox_lrn, equinox_sec ` +
		`FROM equinox.homedel `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		h := Homedel{}

		// scan
		err = q.Scan(&h.Hdcustref, &h.Hdname, &h.Hdaddress1, &h.Hdaddress2, &h.Hdaddress3, &h.Hdaddress4, &h.Hdaddress5, &h.Hdpostcode, &h.Hdcontact, &h.Hdtelephone, &h.Hddelinstruct, &h.Hdsenderacc, &h.Hdpacks, &h.Hdweight, &h.Hddespdate, &h.Hdservicecode, &h.Hdtype, &h.Hdselect, &h.Hddeltype, &h.Hdtypeofpack, &h.Hdemail, &h.Hdsparen1, &h.Hdsparen2, &h.Hdsparen3, &h.Hdspared1, &h.Hdspared2, &h.Hdspared3, &h.Hdsparel1, &h.Hdsparel2, &h.Hdsparel3, &h.Hdprinted, &h.Hdsparec4, &h.Hdsparec5, &h.EquinoxLrn, &h.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(h) {
			return nil
		}
	}

	return nil
}

// HomedelByEquinoxLrn retrieves a row from 'equinox.homedel' as a Homedel.
//
// Generated from index 'homedel_pkey'.
func HomedelByEquinoxLrn(db XODB, equinoxLrn int64) (*Homedel, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`hdcustref, hdname, hdaddress1, hdaddress2, hdaddress3, hdaddress4, hdaddress5, hdpostcode, hdcontact, hdtelephone, hddelinstruct, hdsenderacc, hdpacks, hdweight, hddespdate, hdservicecode, hdtype, hdselect, hddeltype, hdtypeofpack, hdemail, hdsparen1, hdsparen2, hdsparen3, hdspared1, hdspared2, hdspared3, hdsparel1, hdsparel2, hdsparel3, hdprinted, hdsparec4, hdsparec5, equinox_lrn, equinox_sec ` +
		`FROM equinox.homedel ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	h := Homedel{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&h.Hdcustref, &h.Hdname, &h.Hdaddress1, &h.Hdaddress2, &h.Hdaddress3, &h.Hdaddress4, &h.Hdaddress5, &h.Hdpostcode, &h.Hdcontact, &h.Hdtelephone, &h.Hddelinstruct, &h.Hdsenderacc, &h.Hdpacks, &h.Hdweight, &h.Hddespdate, &h.Hdservicecode, &h.Hdtype, &h.Hdselect, &h.Hddeltype, &h.Hdtypeofpack, &h.Hdemail, &h.Hdsparen1, &h.Hdsparen2, &h.Hdsparen3, &h.Hdspared1, &h.Hdspared2, &h.Hdspared3, &h.Hdsparel1, &h.Hdsparel2, &h.Hdsparel3, &h.Hdprinted, &h.Hdsparec4, &h.Hdsparec5, &h.EquinoxLrn, &h.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
