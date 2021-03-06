// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Fsecure represents a row from 'equinox.fsecure'.
type Fsecure struct {
	Fskey        sql.NullString  `json:"fskey"`        // fskey
	Fsissuedate  pq.NullTime     `json:"fsissuedate"`  // fsissuedate
	Fslivedate   pq.NullTime     `json:"fslivedate"`   // fslivedate
	Fsbarreddate pq.NullTime     `json:"fsbarreddate"` // fsbarreddate
	Fsenddate    pq.NullTime     `json:"fsenddate"`    // fsenddate
	Fslicences   sql.NullString  `json:"fslicences"`   // fslicences
	Fsclinumber  sql.NullString  `json:"fsclinumber"`  // fsclinumber
	Fsmemo       sql.NullInt64   `json:"fsmemo"`       // fsmemo
	Fssparec1    sql.NullString  `json:"fssparec1"`    // fssparec1
	Fsspared1    pq.NullTime     `json:"fsspared1"`    // fsspared1
	Fssparen1    sql.NullFloat64 `json:"fssparen1"`    // fssparen1
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec
}

func AllFsecure(db XODB, callback func(x Fsecure) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`fskey, fsissuedate, fslivedate, fsbarreddate, fsenddate, fslicences, fsclinumber, fsmemo, fssparec1, fsspared1, fssparen1, equinox_lrn, equinox_sec ` +
		`FROM equinox.fsecure `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		f := Fsecure{}

		// scan
		err = q.Scan(&f.Fskey, &f.Fsissuedate, &f.Fslivedate, &f.Fsbarreddate, &f.Fsenddate, &f.Fslicences, &f.Fsclinumber, &f.Fsmemo, &f.Fssparec1, &f.Fsspared1, &f.Fssparen1, &f.EquinoxLrn, &f.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(f) {
			return nil
		}
	}

	return nil
}

// FsecureByEquinoxLrn retrieves a row from 'equinox.fsecure' as a Fsecure.
//
// Generated from index 'fsecure_pkey'.
func FsecureByEquinoxLrn(db XODB, equinoxLrn int64) (*Fsecure, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`fskey, fsissuedate, fslivedate, fsbarreddate, fsenddate, fslicences, fsclinumber, fsmemo, fssparec1, fsspared1, fssparen1, equinox_lrn, equinox_sec ` +
		`FROM equinox.fsecure ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	f := Fsecure{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&f.Fskey, &f.Fsissuedate, &f.Fslivedate, &f.Fsbarreddate, &f.Fsenddate, &f.Fslicences, &f.Fsclinumber, &f.Fsmemo, &f.Fssparec1, &f.Fsspared1, &f.Fssparen1, &f.EquinoxLrn, &f.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
