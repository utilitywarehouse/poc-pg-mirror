// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Acbufile represents a row from 'equinox.acbufile'.
type Acbufile struct {
	Gacbutranstype   sql.NullString `json:"gacbutranstype"`   // gacbutranstype
	Gacbushiprsortcd sql.NullString `json:"gacbushiprsortcd"` // gacbushiprsortcd
	Gacbuadjfromdate pq.NullTime    `json:"gacbuadjfromdate"` // gacbuadjfromdate
	Gacbuadjtodate   pq.NullTime    `json:"gacbuadjtodate"`   // gacbuadjtodate
	Gacbureasoncode  sql.NullString `json:"gacbureasoncode"`  // gacbureasoncode
	Gacbuadjusttype  sql.NullString `json:"gacbuadjusttype"`  // gacbuadjusttype
	Gacbudataitemchg sql.NullString `json:"gacbudataitemchg"` // gacbudataitemchg
	Gacbuvalue       sql.NullString `json:"gacbuvalue"`       // gacbuvalue
	Gacburvdvolorgun sql.NullString `json:"gacburvdvolorgun"` // gacburvdvolorgun
	Gacbureasnremrks sql.NullInt64  `json:"gacbureasnremrks"` // gacbureasnremrks
	Gacbudatecreated pq.NullTime    `json:"gacbudatecreated"` // gacbudatecreated
	Gacbutimecreated pq.NullTime    `json:"gacbutimecreated"` // gacbutimecreated
	Gacbufilecreated sql.NullString `json:"gacbufilecreated"` // gacbufilecreated
	Gacbureadrefrnce sql.NullString `json:"gacbureadrefrnce"` // gacbureadrefrnce
	Gacbuapprovaltxt sql.NullInt64  `json:"gacbuapprovaltxt"` // gacbuapprovaltxt
	EquinoxPrn       sql.NullInt64  `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllAcbufile(db XODB, callback func(x Acbufile) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`gacbutranstype, gacbushiprsortcd, gacbuadjfromdate, gacbuadjtodate, gacbureasoncode, gacbuadjusttype, gacbudataitemchg, gacbuvalue, gacburvdvolorgun, gacbureasnremrks, gacbudatecreated, gacbutimecreated, gacbufilecreated, gacbureadrefrnce, gacbuapprovaltxt, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.acbufile `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Acbufile{}

		// scan
		err = q.Scan(&a.Gacbutranstype, &a.Gacbushiprsortcd, &a.Gacbuadjfromdate, &a.Gacbuadjtodate, &a.Gacbureasoncode, &a.Gacbuadjusttype, &a.Gacbudataitemchg, &a.Gacbuvalue, &a.Gacburvdvolorgun, &a.Gacbureasnremrks, &a.Gacbudatecreated, &a.Gacbutimecreated, &a.Gacbufilecreated, &a.Gacbureadrefrnce, &a.Gacbuapprovaltxt, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// AcbufileByEquinoxLrn retrieves a row from 'equinox.acbufile' as a Acbufile.
//
// Generated from index 'acbufile_pkey'.
func AcbufileByEquinoxLrn(db XODB, equinoxLrn int64) (*Acbufile, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gacbutranstype, gacbushiprsortcd, gacbuadjfromdate, gacbuadjtodate, gacbureasoncode, gacbuadjusttype, gacbudataitemchg, gacbuvalue, gacburvdvolorgun, gacbureasnremrks, gacbudatecreated, gacbutimecreated, gacbufilecreated, gacbureadrefrnce, gacbuapprovaltxt, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.acbufile ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Acbufile{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Gacbutranstype, &a.Gacbushiprsortcd, &a.Gacbuadjfromdate, &a.Gacbuadjtodate, &a.Gacbureasoncode, &a.Gacbuadjusttype, &a.Gacbudataitemchg, &a.Gacbuvalue, &a.Gacburvdvolorgun, &a.Gacbureasnremrks, &a.Gacbudatecreated, &a.Gacbutimecreated, &a.Gacbufilecreated, &a.Gacbureadrefrnce, &a.Gacbuapprovaltxt, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
