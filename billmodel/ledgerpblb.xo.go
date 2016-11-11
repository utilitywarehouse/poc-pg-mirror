// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// LedgerpBlb represents a row from 'equinox.ledgerp_blb'.
type LedgerpBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllLedgerpBlb(db XODB, callback func(x LedgerpBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ledgerp_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		lb := LedgerpBlb{}

		// scan
		err = q.Scan(&lb.BlbLrn, &lb.BlbData, &lb.BlbText)
		if err != nil {
			return err
		}
		if !callback(lb) {
			return nil
		}
	}

	return nil
}

// LedgerpBlbByBlbLrn retrieves a row from 'equinox.ledgerp_blb' as a LedgerpBlb.
//
// Generated from index 'ledgerp_blb_pkey'.
func LedgerpBlbByBlbLrn(db XODB, blbLrn int64) (*LedgerpBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ledgerp_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	lb := LedgerpBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&lb.BlbLrn, &lb.BlbData, &lb.BlbText)
	if err != nil {
		return nil, err
	}

	return &lb, nil
}
