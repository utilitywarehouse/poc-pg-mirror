// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Failpay represents a row from 'equinox.failpay'.
type Failpay struct {
	Fpaccount       sql.NullString  `json:"fpaccount"`       // fpaccount
	Fpadd1          sql.NullString  `json:"fpadd1"`          // fpadd1
	Fpadd2          sql.NullString  `json:"fpadd2"`          // fpadd2
	Fpadd3          sql.NullString  `json:"fpadd3"`          // fpadd3
	Fpadd4          sql.NullString  `json:"fpadd4"`          // fpadd4
	Fpcounty        sql.NullString  `json:"fpcounty"`        // fpcounty
	Fppostcode      sql.NullString  `json:"fppostcode"`      // fppostcode
	Fpsqlid         sql.NullString  `json:"fpsqlid"`         // fpsqlid
	Fpccnumber      sql.NullString  `json:"fpccnumber"`      // fpccnumber
	Fpccstart       sql.NullString  `json:"fpccstart"`       // fpccstart
	Fpccend         sql.NullString  `json:"fpccend"`         // fpccend
	Fpccissue       sql.NullString  `json:"fpccissue"`       // fpccissue
	Fpccname        sql.NullString  `json:"fpccname"`        // fpccname
	Fpccpostcode    sql.NullString  `json:"fpccpostcode"`    // fpccpostcode
	Fpccpciuniqueno sql.NullString  `json:"fpccpciuniqueno"` // fpccpciuniqueno
	Fpcctotal       sql.NullFloat64 `json:"fpcctotal"`       // fpcctotal
	Fpcccbc         sql.NullFloat64 `json:"fpcccbc"`         // fpcccbc
	Fpccreddeposit  sql.NullFloat64 `json:"fpccreddeposit"`  // fpccreddeposit
	Fpcchandsetdep  sql.NullFloat64 `json:"fpcchandsetdep"`  // fpcchandsetdep
	Fpcchsetoneoff  sql.NullFloat64 `json:"fpcchsetoneoff"`  // fpcchsetoneoff
	Fpcconlyverify  sql.NullFloat64 `json:"fpcconlyverify"`  // fpcconlyverify
	Fpcctranserrcde sql.NullString  `json:"fpcctranserrcde"` // fpcctranserrcde
	Fpccprepayhdset sql.NullFloat64 `json:"fpccprepayhdset"` // fpccprepayhdset
	Fpccauthcode    sql.NullString  `json:"fpccauthcode"`    // fpccauthcode
	Fpcctype        sql.NullString  `json:"fpcctype"`        // fpcctype
	Fpreexported    sql.NullInt64   `json:"fpreexported"`    // fpreexported
	Fpdate          pq.NullTime     `json:"fpdate"`          // fpdate
	Fpcctranserrmsg sql.NullString  `json:"fpcctranserrmsg"` // fpcctranserrmsg
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllFailpay(db XODB, callback func(x Failpay) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_lrn, equinox_sec ` +
		`FROM equinox.failpay `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		f := Failpay{}

		// scan
		err = q.Scan(&f.Fpaccount, &f.Fpadd1, &f.Fpadd2, &f.Fpadd3, &f.Fpadd4, &f.Fpcounty, &f.Fppostcode, &f.Fpsqlid, &f.Fpccnumber, &f.Fpccstart, &f.Fpccend, &f.Fpccissue, &f.Fpccname, &f.Fpccpostcode, &f.Fpccpciuniqueno, &f.Fpcctotal, &f.Fpcccbc, &f.Fpccreddeposit, &f.Fpcchandsetdep, &f.Fpcchsetoneoff, &f.Fpcconlyverify, &f.Fpcctranserrcde, &f.Fpccprepayhdset, &f.Fpccauthcode, &f.Fpcctype, &f.Fpreexported, &f.Fpdate, &f.Fpcctranserrmsg, &f.EquinoxLrn, &f.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(f) {
			return nil
		}
	}

	return nil
}

// FailpayByEquinoxLrn retrieves a row from 'equinox.failpay' as a Failpay.
//
// Generated from index 'failpay_pkey'.
func FailpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Failpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_lrn, equinox_sec ` +
		`FROM equinox.failpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	f := Failpay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&f.Fpaccount, &f.Fpadd1, &f.Fpadd2, &f.Fpadd3, &f.Fpadd4, &f.Fpcounty, &f.Fppostcode, &f.Fpsqlid, &f.Fpccnumber, &f.Fpccstart, &f.Fpccend, &f.Fpccissue, &f.Fpccname, &f.Fpccpostcode, &f.Fpccpciuniqueno, &f.Fpcctotal, &f.Fpcccbc, &f.Fpccreddeposit, &f.Fpcchandsetdep, &f.Fpcchsetoneoff, &f.Fpcconlyverify, &f.Fpcctranserrcde, &f.Fpccprepayhdset, &f.Fpccauthcode, &f.Fpcctype, &f.Fpreexported, &f.Fpdate, &f.Fpcctranserrmsg, &f.EquinoxLrn, &f.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
