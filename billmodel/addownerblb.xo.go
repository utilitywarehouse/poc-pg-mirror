// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// AddownerBlb represents a row from 'equinox.addowner_blb'.
type AddownerBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllAddownerBlb(db XODB, callback func(x AddownerBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.addowner_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		ab := AddownerBlb{}

		// scan
		err = q.Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
		if err != nil {
			return err
		}
		if !callback(ab) {
			return nil
		}
	}

	return nil
}

// AddownerBlbByBlbLrn retrieves a row from 'equinox.addowner_blb' as a AddownerBlb.
//
// Generated from index 'addowner_blb_pkey'.
func AddownerBlbByBlbLrn(db XODB, blbLrn int64) (*AddownerBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.addowner_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AddownerBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
