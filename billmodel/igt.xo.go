// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Igt represents a row from 'equinox.igt'.
type Igt struct {
	Igtfullname     sql.NullString  `json:"igtfullname"`     // igtfullname
	Igtidno         sql.NullInt64   `json:"igtidno"`         // igtidno
	Igtvalidfrom    pq.NullTime     `json:"igtvalidfrom"`    // igtvalidfrom
	Igtvalidto      pq.NullTime     `json:"igtvalidto"`      // igtvalidto
	Igtrangestart   sql.NullInt64   `json:"igtrangestart"`   // igtrangestart
	Igtrangeend     sql.NullInt64   `json:"igtrangeend"`     // igtrangeend
	Igtemail        sql.NullString  `json:"igtemail"`        // igtemail
	Igttelno        sql.NullString  `json:"igttelno"`        // igttelno
	Igtadd1         sql.NullString  `json:"igtadd1"`         // igtadd1
	Igtadd2         sql.NullString  `json:"igtadd2"`         // igtadd2
	Igtadd3         sql.NullString  `json:"igtadd3"`         // igtadd3
	Igtadd4         sql.NullString  `json:"igtadd4"`         // igtadd4
	Igtcounty       sql.NullString  `json:"igtcounty"`       // igtcounty
	Igtpostcode     sql.NullString  `json:"igtpostcode"`     // igtpostcode
	Igtshortname    sql.NullString  `json:"igtshortname"`    // igtshortname
	Igtsurcharge    sql.NullInt64   `json:"igtsurcharge"`    // igtsurcharge
	Igtsurchargeamt sql.NullFloat64 `json:"igtsurchargeamt"` // igtsurchargeamt
	Igtcnfmethod    sql.NullInt64   `json:"igtcnfmethod"`    // igtcnfmethod
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

// IgtByEquinoxLrn retrieves a row from 'equinox.igt' as a Igt.
//
// Generated from index 'igt_pkey'.
func IgtByEquinoxLrn(db XODB, equinoxLrn int64) (*Igt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`igtfullname, igtidno, igtvalidfrom, igtvalidto, igtrangestart, igtrangeend, igtemail, igttelno, igtadd1, igtadd2, igtadd3, igtadd4, igtcounty, igtpostcode, igtshortname, igtsurcharge, igtsurchargeamt, igtcnfmethod, equinox_lrn, equinox_sec ` +
		`FROM equinox.igt ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	i := Igt{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&i.Igtfullname, &i.Igtidno, &i.Igtvalidfrom, &i.Igtvalidto, &i.Igtrangestart, &i.Igtrangeend, &i.Igtemail, &i.Igttelno, &i.Igtadd1, &i.Igtadd2, &i.Igtadd3, &i.Igtadd4, &i.Igtcounty, &i.Igtpostcode, &i.Igtshortname, &i.Igtsurcharge, &i.Igtsurchargeamt, &i.Igtcnfmethod, &i.EquinoxLrn, &i.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
