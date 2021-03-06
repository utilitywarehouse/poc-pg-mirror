// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// CitylinkBlb represents a row from 'equinox.citylink_blb'.
type CitylinkBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllCitylinkBlb(db XODB, callback func(x CitylinkBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.citylink_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cb := CitylinkBlb{}

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

// CitylinkBlbByBlbLrn retrieves a row from 'equinox.citylink_blb' as a CitylinkBlb.
//
// Generated from index 'citylink_blb_pkey'.
func CitylinkBlbByBlbLrn(db XODB, blbLrn int64) (*CitylinkBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.citylink_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CitylinkBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
