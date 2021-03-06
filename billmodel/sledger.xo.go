// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Sledger represents a row from 'equinox.sledger'.
type Sledger struct {
	Slaccountno    sql.NullString  `json:"slaccountno"`    // slaccountno
	Sloutstanding  sql.NullFloat64 `json:"sloutstanding"`  // sloutstanding
	Slonledger     sql.NullFloat64 `json:"slonledger"`     // slonledger
	Slonprepay     sql.NullFloat64 `json:"slonprepay"`     // slonprepay
	Slarchivedate  pq.NullTime     `json:"slarchivedate"`  // slarchivedate
	Slcomments     sql.NullInt64   `json:"slcomments"`     // slcomments
	Sldebits       sql.NullFloat64 `json:"sldebits"`       // sldebits
	Sldebititems   sql.NullInt64   `json:"sldebititems"`   // sldebititems
	Slcredits      sql.NullFloat64 `json:"slcredits"`      // slcredits
	Slcredititems  sql.NullInt64   `json:"slcredititems"`  // slcredititems
	Sldeposits     sql.NullFloat64 `json:"sldeposits"`     // sldeposits
	Slrecalc       pq.NullTime     `json:"slrecalc"`       // slrecalc
	Slearliestitem pq.NullTime     `json:"slearliestitem"` // slearliestitem
	Sllatestitem   pq.NullTime     `json:"sllatestitem"`   // sllatestitem
	Slosinvoices   sql.NullInt64   `json:"slosinvoices"`   // slosinvoices
	Slnettp        sql.NullFloat64 `json:"slnettp"`        // slnettp
	Slvattp        sql.NullFloat64 `json:"slvattp"`        // slvattp
	Slnetgp        sql.NullFloat64 `json:"slnetgp"`        // slnetgp
	Slvatgp        sql.NullFloat64 `json:"slvatgp"`        // slvatgp
	Slnetep        sql.NullFloat64 `json:"slnetep"`        // slnetep
	Slvatep        sql.NullFloat64 `json:"slvatep"`        // slvatep
	Slnettotal     sql.NullFloat64 `json:"slnettotal"`     // slnettotal
	Slvattotal     sql.NullFloat64 `json:"slvattotal"`     // slvattotal
	Slgpover3      sql.NullFloat64 `json:"slgpover3"`      // slgpover3
	Slepover3      sql.NullFloat64 `json:"slepover3"`      // slepover3
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

func AllSledger(db XODB, callback func(x Sledger) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_lrn, equinox_sec ` +
		`FROM equinox.sledger `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		s := Sledger{}

		// scan
		err = q.Scan(&s.Slaccountno, &s.Sloutstanding, &s.Slonledger, &s.Slonprepay, &s.Slarchivedate, &s.Slcomments, &s.Sldebits, &s.Sldebititems, &s.Slcredits, &s.Slcredititems, &s.Sldeposits, &s.Slrecalc, &s.Slearliestitem, &s.Sllatestitem, &s.Slosinvoices, &s.Slnettp, &s.Slvattp, &s.Slnetgp, &s.Slvatgp, &s.Slnetep, &s.Slvatep, &s.Slnettotal, &s.Slvattotal, &s.Slgpover3, &s.Slepover3, &s.EquinoxLrn, &s.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(s) {
			return nil
		}
	}

	return nil
}

// SledgerByEquinoxLrn retrieves a row from 'equinox.sledger' as a Sledger.
//
// Generated from index 'sledger_pkey'.
func SledgerByEquinoxLrn(db XODB, equinoxLrn int64) (*Sledger, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`slaccountno, sloutstanding, slonledger, slonprepay, slarchivedate, slcomments, sldebits, sldebititems, slcredits, slcredititems, sldeposits, slrecalc, slearliestitem, sllatestitem, slosinvoices, slnettp, slvattp, slnetgp, slvatgp, slnetep, slvatep, slnettotal, slvattotal, slgpover3, slepover3, equinox_lrn, equinox_sec ` +
		`FROM equinox.sledger ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Sledger{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Slaccountno, &s.Sloutstanding, &s.Slonledger, &s.Slonprepay, &s.Slarchivedate, &s.Slcomments, &s.Sldebits, &s.Sldebititems, &s.Slcredits, &s.Slcredititems, &s.Sldeposits, &s.Slrecalc, &s.Slearliestitem, &s.Sllatestitem, &s.Slosinvoices, &s.Slnettp, &s.Slvattp, &s.Slnetgp, &s.Slvatgp, &s.Slnetep, &s.Slvatep, &s.Slnettotal, &s.Slvattotal, &s.Slgpover3, &s.Slepover3, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
