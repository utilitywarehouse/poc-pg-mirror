// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Affmer represents a row from 'equinox.affmer'.
type Affmer struct {
	Affmerid         sql.NullFloat64 `json:"affmerid"`         // affmerid
	Affmername       sql.NullString  `json:"affmername"`       // affmername
	Affmerdesc       sql.NullInt64   `json:"affmerdesc"`       // affmerdesc
	Affmerpreferred  sql.NullString  `json:"affmerpreferred"`  // affmerpreferred
	Affmerperpay     sql.NullFloat64 `json:"affmerperpay"`     // affmerperpay
	Affmerfixpay     sql.NullFloat64 `json:"affmerfixpay"`     // affmerfixpay
	Affmercashback   sql.NullFloat64 `json:"affmercashback"`   // affmercashback
	Affmerbuyressign sql.NullString  `json:"affmerbuyressign"` // affmerbuyressign
	Affmerspecial    sql.NullInt64   `json:"affmerspecial"`    // affmerspecial
	Affmerspecialcbc sql.NullInt64   `json:"affmerspecialcbc"` // affmerspecialcbc
	Affmersectors    sql.NullString  `json:"affmersectors"`    // affmersectors
	Affmerstatus     sql.NullString  `json:"affmerstatus"`     // affmerstatus
	Affmerrefresh    sql.NullFloat64 `json:"affmerrefresh"`    // affmerrefresh
	Affmerspecialurl sql.NullInt64   `json:"affmerspecialurl"` // affmerspecialurl
	Affmerdeeplinkur sql.NullInt64   `json:"affmerdeeplinkur"` // affmerdeeplinkur
	Affmerstoreurl   sql.NullInt64   `json:"affmerstoreurl"`   // affmerstoreurl
	Affmerdomain     sql.NullInt64   `json:"affmerdomain"`     // affmerdomain
	Affmerpref2      sql.NullString  `json:"affmerpref2"`      // affmerpref2
	Affmerchanged    pq.NullTime     `json:"affmerchanged"`    // affmerchanged
	Affmerpmmref     sql.NullFloat64 `json:"affmerpmmref"`     // affmerpmmref
	Affmerranking    sql.NullFloat64 `json:"affmerranking"`    // affmerranking
	Affmerbignames   sql.NullFloat64 `json:"affmerbignames"`   // affmerbignames
	Affmernotes      sql.NullInt64   `json:"affmernotes"`      // affmernotes
	Affmerstart      pq.NullTime     `json:"affmerstart"`      // affmerstart
	Affmerend        pq.NullTime     `json:"affmerend"`        // affmerend
	Affmernovouchers sql.NullString  `json:"affmernovouchers"` // affmernovouchers
	Affmersparec1    sql.NullString  `json:"affmersparec1"`    // affmersparec1
	Affmersparec2    sql.NullString  `json:"affmersparec2"`    // affmersparec2
	Affmersparen1    sql.NullFloat64 `json:"affmersparen1"`    // affmersparen1
	Affmersparen2    sql.NullFloat64 `json:"affmersparen2"`    // affmersparen2
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// AffmerByEquinoxLrn retrieves a row from 'equinox.affmer' as a Affmer.
//
// Generated from index 'affmer_pkey'.
func AffmerByEquinoxLrn(db XODB, equinoxLrn int64) (*Affmer, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`affmerid, affmername, affmerdesc, affmerpreferred, affmerperpay, affmerfixpay, affmercashback, affmerbuyressign, affmerspecial, affmerspecialcbc, affmersectors, affmerstatus, affmerrefresh, affmerspecialurl, affmerdeeplinkur, affmerstoreurl, affmerdomain, affmerpref2, affmerchanged, affmerpmmref, affmerranking, affmerbignames, affmernotes, affmerstart, affmerend, affmernovouchers, affmersparec1, affmersparec2, affmersparen1, affmersparen2, equinox_lrn, equinox_sec ` +
		`FROM equinox.affmer ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Affmer{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Affmerid, &a.Affmername, &a.Affmerdesc, &a.Affmerpreferred, &a.Affmerperpay, &a.Affmerfixpay, &a.Affmercashback, &a.Affmerbuyressign, &a.Affmerspecial, &a.Affmerspecialcbc, &a.Affmersectors, &a.Affmerstatus, &a.Affmerrefresh, &a.Affmerspecialurl, &a.Affmerdeeplinkur, &a.Affmerstoreurl, &a.Affmerdomain, &a.Affmerpref2, &a.Affmerchanged, &a.Affmerpmmref, &a.Affmerranking, &a.Affmerbignames, &a.Affmernotes, &a.Affmerstart, &a.Affmerend, &a.Affmernovouchers, &a.Affmersparec1, &a.Affmersparec2, &a.Affmersparen1, &a.Affmersparen2, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
