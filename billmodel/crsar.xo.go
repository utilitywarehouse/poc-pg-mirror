// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Crsar represents a row from 'equinox.crsar'.
type Crsar struct {
	Crsarid        sql.NullInt64   `json:"crsarid"`        // crsarid
	Crsarexecid    sql.NullString  `json:"crsarexecid"`    // crsarexecid
	Crsaramount    sql.NullFloat64 `json:"crsaramount"`    // crsaramount
	Crsarcommitted pq.NullTime     `json:"crsarcommitted"` // crsarcommitted
	Crsarsparec1   sql.NullString  `json:"crsarsparec1"`   // crsarsparec1
	Crsarsparec2   sql.NullString  `json:"crsarsparec2"`   // crsarsparec2
	Crsarsparec3   sql.NullString  `json:"crsarsparec3"`   // crsarsparec3
	Crsarsparen1   sql.NullInt64   `json:"crsarsparen1"`   // crsarsparen1
	Crsarsparen2   sql.NullInt64   `json:"crsarsparen2"`   // crsarsparen2
	Crsarsparen3   sql.NullInt64   `json:"crsarsparen3"`   // crsarsparen3
	Crsarspared1   sql.NullInt64   `json:"crsarspared1"`   // crsarspared1
	Crsarspared2   sql.NullInt64   `json:"crsarspared2"`   // crsarspared2
	Crsarspared3   sql.NullInt64   `json:"crsarspared3"`   // crsarspared3
	EquinoxPrn     sql.NullInt64   `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64           `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`    // equinox_sec
}

func AllCrsar(db XODB, callback func(x Crsar) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`crsarid, crsarexecid, crsaramount, crsarcommitted, crsarsparec1, crsarsparec2, crsarsparec3, crsarsparen1, crsarsparen2, crsarsparen3, crsarspared1, crsarspared2, crsarspared3, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crsar `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Crsar{}

		// scan
		err = q.Scan(&c.Crsarid, &c.Crsarexecid, &c.Crsaramount, &c.Crsarcommitted, &c.Crsarsparec1, &c.Crsarsparec2, &c.Crsarsparec3, &c.Crsarsparen1, &c.Crsarsparen2, &c.Crsarsparen3, &c.Crsarspared1, &c.Crsarspared2, &c.Crsarspared3, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CrsarByEquinoxLrn retrieves a row from 'equinox.crsar' as a Crsar.
//
// Generated from index 'crsar_pkey'.
func CrsarByEquinoxLrn(db XODB, equinoxLrn int64) (*Crsar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`crsarid, crsarexecid, crsaramount, crsarcommitted, crsarsparec1, crsarsparec2, crsarsparec3, crsarsparen1, crsarsparen2, crsarsparen3, crsarspared1, crsarspared2, crsarspared3, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crsar ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Crsar{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Crsarid, &c.Crsarexecid, &c.Crsaramount, &c.Crsarcommitted, &c.Crsarsparec1, &c.Crsarsparec2, &c.Crsarsparec3, &c.Crsarsparen1, &c.Crsarsparen2, &c.Crsarsparen3, &c.Crsarspared1, &c.Crsarspared2, &c.Crsarspared3, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
