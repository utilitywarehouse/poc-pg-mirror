// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GcvalBlb represents a row from 'equinox.gcval_blb'.
type GcvalBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// GcvalBlbByBlbLrn retrieves a row from 'equinox.gcval_blb' as a GcvalBlb.
//
// Generated from index 'gcval_blb_pkey'.
func GcvalBlbByBlbLrn(db XODB, blbLrn int64) (*GcvalBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gcval_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GcvalBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
