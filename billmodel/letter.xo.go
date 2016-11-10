// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Letter represents a row from 'equinox.letters'.
type Letter struct {
	Letterfilename   sql.NullString `json:"letterfilename"`   // letterfilename
	Letcategory      sql.NullString `json:"letcategory"`      // letcategory
	Letothercategory sql.NullString `json:"letothercategory"` // letothercategory
	Letsimplexduplex sql.NullString `json:"letsimplexduplex"` // letsimplexduplex
	Letsuspend       sql.NullString `json:"letsuspend"`       // letsuspend
	Letsuspendreason sql.NullString `json:"letsuspendreason"` // letsuspendreason
	Letsuspendreaso2 sql.NullInt64  `json:"letsuspendreaso2"` // letsuspendreaso2
	Letcatstrategy   sql.NullString `json:"letcatstrategy"`   // letcatstrategy
	Letsparechar2    sql.NullString `json:"letsparechar2"`    // letsparechar2
	Letpending       sql.NullInt64  `json:"letpending"`       // letpending
	Letsparenum1     sql.NullInt64  `json:"letsparenum1"`     // letsparenum1
	Letsparenum2     sql.NullInt64  `json:"letsparenum2"`     // letsparenum2
	Letdescription   sql.NullString `json:"letdescription"`   // letdescription
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

// LetterByEquinoxLrn retrieves a row from 'equinox.letters' as a Letter.
//
// Generated from index 'letters_pkey'.
func LetterByEquinoxLrn(db XODB, equinoxLrn int64) (*Letter, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`letterfilename, letcategory, letothercategory, letsimplexduplex, letsuspend, letsuspendreason, letsuspendreaso2, letcatstrategy, letsparechar2, letpending, letsparenum1, letsparenum2, letdescription, equinox_lrn, equinox_sec ` +
		`FROM equinox.letters ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	l := Letter{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&l.Letterfilename, &l.Letcategory, &l.Letothercategory, &l.Letsimplexduplex, &l.Letsuspend, &l.Letsuspendreason, &l.Letsuspendreaso2, &l.Letcatstrategy, &l.Letsparechar2, &l.Letpending, &l.Letsparenum1, &l.Letsparenum2, &l.Letdescription, &l.EquinoxLrn, &l.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
