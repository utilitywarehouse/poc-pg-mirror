// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// DeladdrBlb represents a row from 'equinox.deladdr_blb'.
type DeladdrBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllDeladdrBlb(db XODB, callback func(x DeladdrBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.deladdr_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		dbVal := DeladdrBlb{}

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

// DeladdrBlbByBlbLrn retrieves a row from 'equinox.deladdr_blb' as a DeladdrBlb.
//
// Generated from index 'deladdr_blb_pkey'.
func DeladdrBlbByBlbLrn(db XODB, blbLrn int64) (*DeladdrBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.deladdr_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	dbVal := DeladdrBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&dbVal.BlbLrn, &dbVal.BlbData, &dbVal.BlbText)
	if err != nil {
		return nil, err
	}

	return &dbVal, nil
}
