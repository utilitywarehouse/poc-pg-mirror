// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// MportBlb represents a row from 'equinox.mport_blb'.
type MportBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllMportBlb(db XODB, callback func(x MportBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mport_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		mb := MportBlb{}

		// scan
		err = q.Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
		if err != nil {
			return err
		}
		if !callback(mb) {
			return nil
		}
	}

	return nil
}

// MportBlbByBlbLrn retrieves a row from 'equinox.mport_blb' as a MportBlb.
//
// Generated from index 'mport_blb_pkey'.
func MportBlbByBlbLrn(db XODB, blbLrn int64) (*MportBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mport_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MportBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
