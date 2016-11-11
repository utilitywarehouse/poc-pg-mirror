// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// UsergrpBlb represents a row from 'equinox.usergrp_blb'.
type UsergrpBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllUsergrpBlb(db XODB, callback func(x UsergrpBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.usergrp_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		ub := UsergrpBlb{}

		// scan
		err = q.Scan(&ub.BlbLrn, &ub.BlbData, &ub.BlbText)
		if err != nil {
			return err
		}
		if !callback(ub) {
			return nil
		}
	}

	return nil
}

// UsergrpBlbByBlbLrn retrieves a row from 'equinox.usergrp_blb' as a UsergrpBlb.
//
// Generated from index 'usergrp_blb_pkey'.
func UsergrpBlbByBlbLrn(db XODB, blbLrn int64) (*UsergrpBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.usergrp_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ub := UsergrpBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ub.BlbLrn, &ub.BlbData, &ub.BlbText)
	if err != nil {
		return nil, err
	}

	return &ub, nil
}
