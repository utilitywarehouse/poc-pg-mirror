// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cash4cus represents a row from 'equinox.cash4cus'.
type Cash4cus struct {
	Cashexecid       sql.NullString  `json:"cashexecid"`       // cashexecid
	Cashcustaccno    sql.NullString  `json:"cashcustaccno"`    // cashcustaccno
	Cashstatus       sql.NullString  `json:"cashstatus"`       // cashstatus
	Cashinitialamoun sql.NullFloat64 `json:"cashinitialamoun"` // cashinitialamoun
	Cashinitialcrnum sql.NullInt64   `json:"cashinitialcrnum"` // cashinitialcrnum
	Cashinitialmonth sql.NullString  `json:"cashinitialmonth"` // cashinitialmonth
	Cashbalance      sql.NullFloat64 `json:"cashbalance"`      // cashbalance
	Cashclosedcrnum  sql.NullInt64   `json:"cashclosedcrnum"`  // cashclosedcrnum
	Cashloaded       pq.NullTime     `json:"cashloaded"`       // cashloaded
	Cashnotes        sql.NullString  `json:"cashnotes"`        // cashnotes
	Cashsparenum1    sql.NullFloat64 `json:"cashsparenum1"`    // cashsparenum1
	Cashsparenum2    sql.NullFloat64 `json:"cashsparenum2"`    // cashsparenum2
	Cashsparechar1   sql.NullString  `json:"cashsparechar1"`   // cashsparechar1
	Cashsparechar2   sql.NullString  `json:"cashsparechar2"`   // cashsparechar2
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// Cash4cusByEquinoxLrn retrieves a row from 'equinox.cash4cus' as a Cash4cus.
//
// Generated from index 'cash4cus_pkey'.
func Cash4cusByEquinoxLrn(db XODB, equinoxLrn int64) (*Cash4cus, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cashexecid, cashcustaccno, cashstatus, cashinitialamoun, cashinitialcrnum, cashinitialmonth, cashbalance, cashclosedcrnum, cashloaded, cashnotes, cashsparenum1, cashsparenum2, cashsparechar1, cashsparechar2, equinox_lrn, equinox_sec ` +
		`FROM equinox.cash4cus ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cash4cus{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cashexecid, &c.Cashcustaccno, &c.Cashstatus, &c.Cashinitialamoun, &c.Cashinitialcrnum, &c.Cashinitialmonth, &c.Cashbalance, &c.Cashclosedcrnum, &c.Cashloaded, &c.Cashnotes, &c.Cashsparenum1, &c.Cashsparenum2, &c.Cashsparechar1, &c.Cashsparechar2, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
