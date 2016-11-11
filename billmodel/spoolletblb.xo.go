// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// SpoolletBlb represents a row from 'equinox.spoollet_blb'.
type SpoolletBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllSpoolletBlb(db XODB, callback func(x SpoolletBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.spoollet_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		sb := SpoolletBlb{}

		// scan
		err = q.Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
		if err != nil {
			return err
		}
		if !callback(sb) {
			return nil
		}
	}

	return nil
}

// SpoolletBlbByBlbLrn retrieves a row from 'equinox.spoollet_blb' as a SpoolletBlb.
//
// Generated from index 'spoollet_blb_pkey'.
func SpoolletBlbByBlbLrn(db XODB, blbLrn int64) (*SpoolletBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.spoollet_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := SpoolletBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}
