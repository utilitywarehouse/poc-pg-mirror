// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Afftran represents a row from 'equinox.afftrans'.
type Afftran struct {
	Affttransactid  sql.NullFloat64 `json:"affttransactid"`  // affttransactid
	Afftdate        pq.NullTime     `json:"afftdate"`        // afftdate
	Afftdescription sql.NullString  `json:"afftdescription"` // afftdescription
	Afftvalue       sql.NullFloat64 `json:"afftvalue"`       // afftvalue
	Afftcommis      sql.NullFloat64 `json:"afftcommis"`      // afftcommis
	Afftsaving      sql.NullFloat64 `json:"afftsaving"`      // afftsaving
	Afftbillno      sql.NullFloat64 `json:"afftbillno"`      // afftbillno
	Afftafftype     sql.NullString  `json:"afftafftype"`     // afftafftype
	Afftaffid       sql.NullString  `json:"afftaffid"`       // afftaffid
	Afftafftranref  sql.NullString  `json:"afftafftranref"`  // afftafftranref
	Affturlaccount  sql.NullString  `json:"affturlaccount"`  // affturlaccount
	Afftaffmerid    sql.NullString  `json:"afftaffmerid"`    // afftaffmerid
	Afftsparec1     sql.NullString  `json:"afftsparec1"`     // afftsparec1
	Afftsparec2     sql.NullString  `json:"afftsparec2"`     // afftsparec2
	Afftspared1     pq.NullTime     `json:"afftspared1"`     // afftspared1
	Afftspared2     pq.NullTime     `json:"afftspared2"`     // afftspared2
	Afftspared3     pq.NullTime     `json:"afftspared3"`     // afftspared3
	Afftsparen1     sql.NullFloat64 `json:"afftsparen1"`     // afftsparen1
	Afftsparen2     sql.NullFloat64 `json:"afftsparen2"`     // afftsparen2
	Afftsparen3     sql.NullFloat64 `json:"afftsparen3"`     // afftsparen3
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Afftran exists in the database.
func (a *Afftran) Exists() bool {
	return a._exists
}

// Deleted provides information if the Afftran has been deleted from the database.
func (a *Afftran) Deleted() bool {
	return a._deleted
}

// Insert inserts the Afftran to the database.
func (a *Afftran) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.afftrans (` +
		`affttransactid, afftdate, afftdescription, afftvalue, afftcommis, afftsaving, afftbillno, afftafftype, afftaffid, afftafftranref, affturlaccount, afftaffmerid, afftsparec1, afftsparec2, afftspared1, afftspared2, afftspared3, afftsparen1, afftsparen2, afftsparen3, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Afftran in the database.
func (a *Afftran) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.afftrans SET (` +
		`affttransactid, afftdate, afftdescription, afftvalue, afftcommis, afftsaving, afftbillno, afftafftype, afftaffid, afftafftranref, affturlaccount, afftaffmerid, afftsparec1, afftsparec2, afftspared1, afftspared2, afftspared3, afftsparen1, afftsparen2, afftsparen3, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22` +
		`) WHERE equinox_lrn = $23`

	// run query
	XOLog(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Afftran to the database.
func (a *Afftran) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Afftran.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Afftran) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.afftrans (` +
		`affttransactid, afftdate, afftdescription, afftvalue, afftcommis, afftsaving, afftbillno, afftafftype, afftaffid, afftafftranref, affturlaccount, afftaffmerid, afftsparec1, afftsparec2, afftspared1, afftspared2, afftspared3, afftsparen1, afftsparen2, afftsparen3, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`affttransactid, afftdate, afftdescription, afftvalue, afftcommis, afftsaving, afftbillno, afftafftype, afftaffid, afftafftranref, affturlaccount, afftaffmerid, afftsparec1, afftsparec2, afftspared1, afftspared2, afftspared3, afftsparen1, afftsparen2, afftsparen3, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.affttransactid, EXCLUDED.afftdate, EXCLUDED.afftdescription, EXCLUDED.afftvalue, EXCLUDED.afftcommis, EXCLUDED.afftsaving, EXCLUDED.afftbillno, EXCLUDED.afftafftype, EXCLUDED.afftaffid, EXCLUDED.afftafftranref, EXCLUDED.affturlaccount, EXCLUDED.afftaffmerid, EXCLUDED.afftsparec1, EXCLUDED.afftsparec2, EXCLUDED.afftspared1, EXCLUDED.afftspared2, EXCLUDED.afftspared3, EXCLUDED.afftsparen1, EXCLUDED.afftsparen2, EXCLUDED.afftsparen3, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.Affttransactid, a.Afftdate, a.Afftdescription, a.Afftvalue, a.Afftcommis, a.Afftsaving, a.Afftbillno, a.Afftafftype, a.Afftaffid, a.Afftafftranref, a.Affturlaccount, a.Afftaffmerid, a.Afftsparec1, a.Afftsparec2, a.Afftspared1, a.Afftspared2, a.Afftspared3, a.Afftsparen1, a.Afftsparen2, a.Afftsparen3, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Afftran from the database.
func (a *Afftran) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.afftrans WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AfftranByEquinoxLrn retrieves a row from 'equinox.afftrans' as a Afftran.
//
// Generated from index 'afftrans_pkey'.
func AfftranByEquinoxLrn(db XODB, equinoxLrn int64) (*Afftran, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`affttransactid, afftdate, afftdescription, afftvalue, afftcommis, afftsaving, afftbillno, afftafftype, afftaffid, afftafftranref, affturlaccount, afftaffmerid, afftsparec1, afftsparec2, afftspared1, afftspared2, afftspared3, afftsparen1, afftsparen2, afftsparen3, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.afftrans ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Afftran{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Affttransactid, &a.Afftdate, &a.Afftdescription, &a.Afftvalue, &a.Afftcommis, &a.Afftsaving, &a.Afftbillno, &a.Afftafftype, &a.Afftaffid, &a.Afftafftranref, &a.Affturlaccount, &a.Afftaffmerid, &a.Afftsparec1, &a.Afftsparec2, &a.Afftspared1, &a.Afftspared2, &a.Afftspared3, &a.Afftsparen1, &a.Afftsparen2, &a.Afftsparen3, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
