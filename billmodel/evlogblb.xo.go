// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// EvlogBlb represents a row from 'equinox.evlog_blb'.
type EvlogBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllEvlogBlb(db XODB, callback func(x EvlogBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.evlog_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		eb := EvlogBlb{}

		// scan
		err = q.Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
		if err != nil {
			return err
		}
		if !callback(eb) {
			return nil
		}
	}

	return nil
}

// EvlogBlbByBlbLrn retrieves a row from 'equinox.evlog_blb' as a EvlogBlb.
//
// Generated from index 'evlog_blb_pkey'.
func EvlogBlbByBlbLrn(db XODB, blbLrn int64) (*EvlogBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.evlog_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := EvlogBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
