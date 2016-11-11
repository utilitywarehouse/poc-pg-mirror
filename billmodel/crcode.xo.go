// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Crcode represents a row from 'equinox.crcodes'.
type Crcode struct {
	Crccodeid       sql.NullInt64   `json:"crccodeid"`       // crccodeid
	Crccode         sql.NullString  `json:"crccode"`         // crccode
	Crcsetdispflags sql.NullString  `json:"crcsetdispflags"` // crcsetdispflags
	Crcdisplay1     sql.NullString  `json:"crcdisplay1"`     // crcdisplay1
	Crcdisplay2     sql.NullString  `json:"crcdisplay2"`     // crcdisplay2
	Crcdisplay3     sql.NullString  `json:"crcdisplay3"`     // crcdisplay3
	Crcdisplay4     sql.NullString  `json:"crcdisplay4"`     // crcdisplay4
	Crcount         sql.NullInt64   `json:"crcount"`         // crcount
	Crtotal1        sql.NullFloat64 `json:"crtotal1"`        // crtotal1
	Crtotal2        sql.NullFloat64 `json:"crtotal2"`        // crtotal2
	Crcsparenum1    sql.NullFloat64 `json:"crcsparenum1"`    // crcsparenum1
	Crcsparenum2    sql.NullFloat64 `json:"crcsparenum2"`    // crcsparenum2
	Crcsparec1      sql.NullString  `json:"crcsparec1"`      // crcsparec1
	Crcsparec2      sql.NullString  `json:"crcsparec2"`      // crcsparec2
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllCrcode(db XODB, callback func(x Crcode) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`crccodeid, crccode, crcsetdispflags, crcdisplay1, crcdisplay2, crcdisplay3, crcdisplay4, crcount, crtotal1, crtotal2, crcsparenum1, crcsparenum2, crcsparec1, crcsparec2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crcodes `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Crcode{}

		// scan
		err = q.Scan(&c.Crccodeid, &c.Crccode, &c.Crcsetdispflags, &c.Crcdisplay1, &c.Crcdisplay2, &c.Crcdisplay3, &c.Crcdisplay4, &c.Crcount, &c.Crtotal1, &c.Crtotal2, &c.Crcsparenum1, &c.Crcsparenum2, &c.Crcsparec1, &c.Crcsparec2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CrcodeByEquinoxLrn retrieves a row from 'equinox.crcodes' as a Crcode.
//
// Generated from index 'crcodes_pkey'.
func CrcodeByEquinoxLrn(db XODB, equinoxLrn int64) (*Crcode, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`crccodeid, crccode, crcsetdispflags, crcdisplay1, crcdisplay2, crcdisplay3, crcdisplay4, crcount, crtotal1, crtotal2, crcsparenum1, crcsparenum2, crcsparec1, crcsparec2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crcodes ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Crcode{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Crccodeid, &c.Crccode, &c.Crcsetdispflags, &c.Crcdisplay1, &c.Crcdisplay2, &c.Crcdisplay3, &c.Crcdisplay4, &c.Crcount, &c.Crtotal1, &c.Crtotal2, &c.Crcsparenum1, &c.Crcsparenum2, &c.Crcsparec1, &c.Crcsparec2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
