// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// ValcomBlb represents a row from 'equinox.valcom_blb'.
type ValcomBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// ValcomBlbByBlbLrn retrieves a row from 'equinox.valcom_blb' as a ValcomBlb.
//
// Generated from index 'valcom_blb_pkey'.
func ValcomBlbByBlbLrn(db XODB, blbLrn int64) (*ValcomBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.valcom_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	vb := ValcomBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&vb.BlbLrn, &vb.BlbData, &vb.BlbText)
	if err != nil {
		return nil, err
	}

	return &vb, nil
}
