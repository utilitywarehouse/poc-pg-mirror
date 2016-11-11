// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// CrmsblokBlb represents a row from 'equinox.crmsblok_blb'.
type CrmsblokBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllCrmsblokBlb(db XODB, callback func(x CrmsblokBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.crmsblok_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cb := CrmsblokBlb{}

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

// CrmsblokBlbByBlbLrn retrieves a row from 'equinox.crmsblok_blb' as a CrmsblokBlb.
//
// Generated from index 'crmsblok_blb_pkey'.
func CrmsblokBlbByBlbLrn(db XODB, blbLrn int64) (*CrmsblokBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.crmsblok_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CrmsblokBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
