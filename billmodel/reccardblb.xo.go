// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// ReccardBlb represents a row from 'equinox.reccard_blb'.
type ReccardBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// ReccardBlbByBlbLrn retrieves a row from 'equinox.reccard_blb' as a ReccardBlb.
//
// Generated from index 'reccard_blb_pkey'.
func ReccardBlbByBlbLrn(db XODB, blbLrn int64) (*ReccardBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.reccard_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	rb := ReccardBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&rb.BlbLrn, &rb.BlbData, &rb.BlbText)
	if err != nil {
		return nil, err
	}

	return &rb, nil
}
