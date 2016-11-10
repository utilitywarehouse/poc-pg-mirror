// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// MpmbBlb represents a row from 'equinox.mpmb_blb'.
type MpmbBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// MpmbBlbByBlbLrn retrieves a row from 'equinox.mpmb_blb' as a MpmbBlb.
//
// Generated from index 'mpmb_blb_pkey'.
func MpmbBlbByBlbLrn(db XODB, blbLrn int64) (*MpmbBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mpmb_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MpmbBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
