// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gusedom represents a row from 'equinox.gusedom'.
type Gusedom struct {
	Guduniquesys     sql.NullFloat64 `json:"guduniquesys"`     // guduniquesys
	Guddate          pq.NullTime     `json:"guddate"`          // guddate
	Gudpercentaqused sql.NullFloat64 `json:"gudpercentaqused"` // gudpercentaqused
	Guderead         sql.NullFloat64 `json:"guderead"`         // guderead
	Gudaread         sql.NullFloat64 `json:"gudaread"`         // gudaread
	Gudcread         sql.NullFloat64 `json:"gudcread"`         // gudcread
	Gudreading       sql.NullFloat64 `json:"gudreading"`       // gudreading
	Gudreadtype      sql.NullString  `json:"gudreadtype"`      // gudreadtype
	Gudthroughzero   sql.NullFloat64 `json:"gudthroughzero"`   // gudthroughzero
	Gudmeterexchange sql.NullString  `json:"gudmeterexchange"` // gudmeterexchange
	Gudpercentaq     sql.NullFloat64 `json:"gudpercentaq"`     // gudpercentaq
	Gudkwhused       sql.NullFloat64 `json:"gudkwhused"`       // gudkwhused
	Gudregister      sql.NullString  `json:"gudregister"`      // gudregister
	Gudbilled        sql.NullFloat64 `json:"gudbilled"`        // gudbilled
	Gudtypeimpormet  sql.NullString  `json:"gudtypeimpormet"`  // gudtypeimpormet
	Gudmeterdials    sql.NullInt64   `json:"gudmeterdials"`    // gudmeterdials
	Gudmeterunits    sql.NullInt64   `json:"gudmeterunits"`    // gudmeterunits
	Gudsparen1       sql.NullFloat64 `json:"gudsparen1"`       // gudsparen1
	Gudmeterserialno sql.NullString  `json:"gudmeterserialno"` // gudmeterserialno
	Gudopusreading   sql.NullString  `json:"gudopusreading"`   // gudopusreading
	Gudsendtotransco pq.NullTime     `json:"gudsendtotransco"` // gudsendtotransco
	Gudsend          sql.NullInt64   `json:"gudsend"`          // gudsend
	Gudsuppressread  sql.NullString  `json:"gudsuppressread"`  // gudsuppressread
	Gudppkwh         sql.NullFloat64 `json:"gudppkwh"`         // gudppkwh
	Gudcv            sql.NullFloat64 `json:"gudcv"`            // gudcv
	Gudnett          sql.NullFloat64 `json:"gudnett"`          // gudnett
	Gudunits         sql.NullInt64   `json:"gudunits"`         // gudunits
	Gudsupreadarb    sql.NullString  `json:"gudsupreadarb"`    // gudsupreadarb
	Gudtariff        sql.NullString  `json:"gudtariff"`        // gudtariff
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// GusedomsByEquinoxPrnGuddate retrieves a row from 'equinox.gusedom' as a Gusedom.
//
// Generated from index 'gu_prm_guddate'.
func GusedomsByEquinoxPrnGuddate(db XODB, equinoxPrn sql.NullInt64, guddate pq.NullTime) ([]*Gusedom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`guduniquesys, guddate, gudpercentaqused, guderead, gudaread, gudcread, gudreading, gudreadtype, gudthroughzero, gudmeterexchange, gudpercentaq, gudkwhused, gudregister, gudbilled, gudtypeimpormet, gudmeterdials, gudmeterunits, gudsparen1, gudmeterserialno, gudopusreading, gudsendtotransco, gudsend, gudsuppressread, gudppkwh, gudcv, gudnett, gudunits, gudsupreadarb, gudtariff, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gusedom ` +
		`WHERE equinox_prn = $1 AND guddate = $2`

	// run query
	XOLog(sqlstr, equinoxPrn, guddate)
	q, err := db.Query(sqlstr, equinoxPrn, guddate)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Gusedom{}
	for q.Next() {
		g := Gusedom{}

		// scan
		err = q.Scan(&g.Guduniquesys, &g.Guddate, &g.Gudpercentaqused, &g.Guderead, &g.Gudaread, &g.Gudcread, &g.Gudreading, &g.Gudreadtype, &g.Gudthroughzero, &g.Gudmeterexchange, &g.Gudpercentaq, &g.Gudkwhused, &g.Gudregister, &g.Gudbilled, &g.Gudtypeimpormet, &g.Gudmeterdials, &g.Gudmeterunits, &g.Gudsparen1, &g.Gudmeterserialno, &g.Gudopusreading, &g.Gudsendtotransco, &g.Gudsend, &g.Gudsuppressread, &g.Gudppkwh, &g.Gudcv, &g.Gudnett, &g.Gudunits, &g.Gudsupreadarb, &g.Gudtariff, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GusedomsByEquinoxPrn retrieves a row from 'equinox.gusedom' as a Gusedom.
//
// Generated from index 'gu_prn'.
func GusedomsByEquinoxPrn(db XODB, equinoxPrn sql.NullInt64) ([]*Gusedom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`guduniquesys, guddate, gudpercentaqused, guderead, gudaread, gudcread, gudreading, gudreadtype, gudthroughzero, gudmeterexchange, gudpercentaq, gudkwhused, gudregister, gudbilled, gudtypeimpormet, gudmeterdials, gudmeterunits, gudsparen1, gudmeterserialno, gudopusreading, gudsendtotransco, gudsend, gudsuppressread, gudppkwh, gudcv, gudnett, gudunits, gudsupreadarb, gudtariff, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gusedom ` +
		`WHERE equinox_prn = $1`

	// run query
	XOLog(sqlstr, equinoxPrn)
	q, err := db.Query(sqlstr, equinoxPrn)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Gusedom{}
	for q.Next() {
		g := Gusedom{}

		// scan
		err = q.Scan(&g.Guduniquesys, &g.Guddate, &g.Gudpercentaqused, &g.Guderead, &g.Gudaread, &g.Gudcread, &g.Gudreading, &g.Gudreadtype, &g.Gudthroughzero, &g.Gudmeterexchange, &g.Gudpercentaq, &g.Gudkwhused, &g.Gudregister, &g.Gudbilled, &g.Gudtypeimpormet, &g.Gudmeterdials, &g.Gudmeterunits, &g.Gudsparen1, &g.Gudmeterserialno, &g.Gudopusreading, &g.Gudsendtotransco, &g.Gudsend, &g.Gudsuppressread, &g.Gudppkwh, &g.Gudcv, &g.Gudnett, &g.Gudunits, &g.Gudsupreadarb, &g.Gudtariff, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
		if err != nil {
			return nil, err
		}

		res = append(res, &g)
	}

	return res, nil
}

// GusedomByEquinoxLrn retrieves a row from 'equinox.gusedom' as a Gusedom.
//
// Generated from index 'gusedom_pkey'.
func GusedomByEquinoxLrn(db XODB, equinoxLrn int64) (*Gusedom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`guduniquesys, guddate, gudpercentaqused, guderead, gudaread, gudcread, gudreading, gudreadtype, gudthroughzero, gudmeterexchange, gudpercentaq, gudkwhused, gudregister, gudbilled, gudtypeimpormet, gudmeterdials, gudmeterunits, gudsparen1, gudmeterserialno, gudopusreading, gudsendtotransco, gudsend, gudsuppressread, gudppkwh, gudcv, gudnett, gudunits, gudsupreadarb, gudtariff, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gusedom ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gusedom{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Guduniquesys, &g.Guddate, &g.Gudpercentaqused, &g.Guderead, &g.Gudaread, &g.Gudcread, &g.Gudreading, &g.Gudreadtype, &g.Gudthroughzero, &g.Gudmeterexchange, &g.Gudpercentaq, &g.Gudkwhused, &g.Gudregister, &g.Gudbilled, &g.Gudtypeimpormet, &g.Gudmeterdials, &g.Gudmeterunits, &g.Gudsparen1, &g.Gudmeterserialno, &g.Gudopusreading, &g.Gudsendtotransco, &g.Gudsend, &g.Gudsuppressread, &g.Gudppkwh, &g.Gudcv, &g.Gudnett, &g.Gudunits, &g.Gudsupreadarb, &g.Gudtariff, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
