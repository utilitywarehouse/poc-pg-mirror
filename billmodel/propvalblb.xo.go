// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PropvalBlb represents a row from 'equinox.propval_blb'.
type PropvalBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// PropvalBlbByBlbLrn retrieves a row from 'equinox.propval_blb' as a PropvalBlb.
//
// Generated from index 'propval_blb_pkey'.
func PropvalBlbByBlbLrn(db XODB, blbLrn int64) (*PropvalBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.propval_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PropvalBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
