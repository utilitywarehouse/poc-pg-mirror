// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// SlethistBlb represents a row from 'equinox.slethist_blb'.
type SlethistBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// SlethistBlbByBlbLrn retrieves a row from 'equinox.slethist_blb' as a SlethistBlb.
//
// Generated from index 'slethist_blb_pkey'.
func SlethistBlbByBlbLrn(db XODB, blbLrn int64) (*SlethistBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.slethist_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := SlethistBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}
