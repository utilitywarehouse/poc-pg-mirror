// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// VerbrteBlb represents a row from 'equinox.verbrte_blb'.
type VerbrteBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// VerbrteBlbByBlbLrn retrieves a row from 'equinox.verbrte_blb' as a VerbrteBlb.
//
// Generated from index 'verbrte_blb_pkey'.
func VerbrteBlbByBlbLrn(db XODB, blbLrn int64) (*VerbrteBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.verbrte_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	vb := VerbrteBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&vb.BlbLrn, &vb.BlbData, &vb.BlbText)
	if err != nil {
		return nil, err
	}

	return &vb, nil
}
