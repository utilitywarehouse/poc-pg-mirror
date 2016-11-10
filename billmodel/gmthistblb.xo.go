// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GmthistBlb represents a row from 'equinox.gmthist_blb'.
type GmthistBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// GmthistBlbByBlbLrn retrieves a row from 'equinox.gmthist_blb' as a GmthistBlb.
//
// Generated from index 'gmthist_blb_pkey'.
func GmthistBlbByBlbLrn(db XODB, blbLrn int64) (*GmthistBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gmthist_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GmthistBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
