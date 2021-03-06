// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Udccard represents a row from 'equinox.udccard'.
type Udccard struct {
	Udccenteredon    pq.NullTime     `json:"udccenteredon"`    // udccenteredon
	Udccenteredby    sql.NullString  `json:"udccenteredby"`    // udccenteredby
	Udcccardtype     sql.NullString  `json:"udcccardtype"`     // udcccardtype
	Udcccardnumber   sql.NullString  `json:"udcccardnumber"`   // udcccardnumber
	Udcccardname     sql.NullString  `json:"udcccardname"`     // udcccardname
	Udccshopperref   sql.NullString  `json:"udccshopperref"`   // udccshopperref
	Udccpostcode     sql.NullString  `json:"udccpostcode"`     // udccpostcode
	Udccvalid        sql.NullString  `json:"udccvalid"`        // udccvalid
	Udccexpiry       sql.NullString  `json:"udccexpiry"`       // udccexpiry
	Udcccardissue    sql.NullString  `json:"udcccardissue"`    // udcccardissue
	Udccdigits       sql.NullString  `json:"udccdigits"`       // udccdigits
	Udccaddress      sql.NullInt64   `json:"udccaddress"`      // udccaddress
	Udcccardunisys   sql.NullString  `json:"udcccardunisys"`   // udcccardunisys
	Udccfirstpayment sql.NullFloat64 `json:"udccfirstpayment"` // udccfirstpayment
	Udccfirstpaydate pq.NullTime     `json:"udccfirstpaydate"` // udccfirstpaydate
	Udccchangetime   pq.NullTime     `json:"udccchangetime"`   // udccchangetime
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllUdccard(db XODB, callback func(x Udccard) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`udccenteredon, udccenteredby, udcccardtype, udcccardnumber, udcccardname, udccshopperref, udccpostcode, udccvalid, udccexpiry, udcccardissue, udccdigits, udccaddress, udcccardunisys, udccfirstpayment, udccfirstpaydate, udccchangetime, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.udccard `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		u := Udccard{}

		// scan
		err = q.Scan(&u.Udccenteredon, &u.Udccenteredby, &u.Udcccardtype, &u.Udcccardnumber, &u.Udcccardname, &u.Udccshopperref, &u.Udccpostcode, &u.Udccvalid, &u.Udccexpiry, &u.Udcccardissue, &u.Udccdigits, &u.Udccaddress, &u.Udcccardunisys, &u.Udccfirstpayment, &u.Udccfirstpaydate, &u.Udccchangetime, &u.EquinoxPrn, &u.EquinoxLrn, &u.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(u) {
			return nil
		}
	}

	return nil
}

// UdccardByEquinoxLrn retrieves a row from 'equinox.udccard' as a Udccard.
//
// Generated from index 'udccard_pkey'.
func UdccardByEquinoxLrn(db XODB, equinoxLrn int64) (*Udccard, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`udccenteredon, udccenteredby, udcccardtype, udcccardnumber, udcccardname, udccshopperref, udccpostcode, udccvalid, udccexpiry, udcccardissue, udccdigits, udccaddress, udcccardunisys, udccfirstpayment, udccfirstpaydate, udccchangetime, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.udccard ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	u := Udccard{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&u.Udccenteredon, &u.Udccenteredby, &u.Udcccardtype, &u.Udcccardnumber, &u.Udcccardname, &u.Udccshopperref, &u.Udccpostcode, &u.Udccvalid, &u.Udccexpiry, &u.Udcccardissue, &u.Udccdigits, &u.Udccaddress, &u.Udcccardunisys, &u.Udccfirstpayment, &u.Udccfirstpaydate, &u.Udccchangetime, &u.EquinoxPrn, &u.EquinoxLrn, &u.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
