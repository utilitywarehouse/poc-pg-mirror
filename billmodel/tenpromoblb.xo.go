// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// TenpromoBlb represents a row from 'equinox.tenpromo_blb'.
type TenpromoBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// TenpromoBlbByBlbLrn retrieves a row from 'equinox.tenpromo_blb' as a TenpromoBlb.
//
// Generated from index 'tenpromo_blb_pkey'.
func TenpromoBlbByBlbLrn(db XODB, blbLrn int64) (*TenpromoBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.tenpromo_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	tb := TenpromoBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&tb.BlbLrn, &tb.BlbData, &tb.BlbText)
	if err != nil {
		return nil, err
	}

	return &tb, nil
}
