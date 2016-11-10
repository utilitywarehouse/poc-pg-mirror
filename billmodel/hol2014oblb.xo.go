// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Hol2014oBlb represents a row from 'equinox.hol2014o_blb'.
type Hol2014oBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// Hol2014oBlbByBlbLrn retrieves a row from 'equinox.hol2014o_blb' as a Hol2014oBlb.
//
// Generated from index 'hol2014o_blb_pkey'.
func Hol2014oBlbByBlbLrn(db XODB, blbLrn int64) (*Hol2014oBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.hol2014o_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	hb := Hol2014oBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&hb.BlbLrn, &hb.BlbData, &hb.BlbText)
	if err != nil {
		return nil, err
	}

	return &hb, nil
}
