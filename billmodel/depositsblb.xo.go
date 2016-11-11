// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// DepositsBlb represents a row from 'equinox.deposits_blb'.
type DepositsBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllDepositsBlb(db XODB, callback func(x DepositsBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.deposits_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		dbVal := DepositsBlb{}

		// scan
		err = q.Scan(&dbVal.BlbLrn, &dbVal.BlbData, &dbVal.BlbText)
		if err != nil {
			return err
		}
		if !callback(dbVal) {
			return nil
		}
	}

	return nil
}

// DepositsBlbByBlbLrn retrieves a row from 'equinox.deposits_blb' as a DepositsBlb.
//
// Generated from index 'deposits_blb_pkey'.
func DepositsBlbByBlbLrn(db XODB, blbLrn int64) (*DepositsBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.deposits_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	dbVal := DepositsBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&dbVal.BlbLrn, &dbVal.BlbData, &dbVal.BlbText)
	if err != nil {
		return nil, err
	}

	return &dbVal, nil
}
