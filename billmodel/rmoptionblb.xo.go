// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// RmoptionBlb represents a row from 'equinox.rmoption_blb'.
type RmoptionBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// RmoptionBlbByBlbLrn retrieves a row from 'equinox.rmoption_blb' as a RmoptionBlb.
//
// Generated from index 'rmoption_blb_pkey'.
func RmoptionBlbByBlbLrn(db XODB, blbLrn int64) (*RmoptionBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.rmoption_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	rb := RmoptionBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&rb.BlbLrn, &rb.BlbData, &rb.BlbText)
	if err != nil {
		return nil, err
	}

	return &rb, nil
}
