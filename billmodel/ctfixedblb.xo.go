// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// CtfixedBlb represents a row from 'equinox.ctfixed_blb'.
type CtfixedBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllCtfixedBlb(db XODB, callback func(x CtfixedBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ctfixed_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cb := CtfixedBlb{}

		// scan
		err = q.Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
		if err != nil {
			return err
		}
		if !callback(cb) {
			return nil
		}
	}

	return nil
}

// CtfixedBlbByBlbLrn retrieves a row from 'equinox.ctfixed_blb' as a CtfixedBlb.
//
// Generated from index 'ctfixed_blb_pkey'.
func CtfixedBlbByBlbLrn(db XODB, blbLrn int64) (*CtfixedBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ctfixed_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CtfixedBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
