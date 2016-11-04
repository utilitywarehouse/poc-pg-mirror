// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Call represents a row from 'equinox.calls'.
type Call struct {
	Calldate         pq.NullTime     `json:"calldate"`         // calldate
	Calltime         sql.NullString  `json:"calltime"`         // calltime
	Calldestination  sql.NullString  `json:"calldestination"`  // calldestination
	Calldestnumber   sql.NullString  `json:"calldestnumber"`   // calldestnumber
	Callratingperiod sql.NullString  `json:"callratingperiod"` // callratingperiod
	Callduration     sql.NullFloat64 `json:"callduration"`     // callduration
	Callcarrierband  sql.NullString  `json:"callcarrierband"`  // callcarrierband
	Callwholesale    sql.NullInt64   `json:"callwholesale"`    // callwholesale
	Callretail       sql.NullString  `json:"callretail"`       // callretail
	Callcalcwholesal sql.NullInt64   `json:"callcalcwholesal"` // callcalcwholesal
	Callduplicate    sql.NullInt64   `json:"callduplicate"`    // callduplicate
	Callfileid       sql.NullString  `json:"callfileid"`       // callfileid
	Callcarrier      sql.NullString  `json:"callcarrier"`      // callcarrier
	Callactretail    sql.NullFloat64 `json:"callactretail"`    // callactretail
	Callactretailpre sql.NullFloat64 `json:"callactretailpre"` // callactretailpre
	Calldesttype     sql.NullInt64   `json:"calldesttype"`     // calldesttype
	Callourband      sql.NullString  `json:"callourband"`      // callourband
	Callcustbillingg sql.NullInt64   `json:"callcustbillingg"` // callcustbillingg
	Callpromocode    sql.NullString  `json:"callpromocode"`    // callpromocode
	Callextn         sql.NullString  `json:"callextn"`         // callextn
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Call exists in the database.
func (c *Call) Exists() bool {
	return c._exists
}

// Deleted provides information if the Call has been deleted from the database.
func (c *Call) Deleted() bool {
	return c._deleted
}

// Insert inserts the Call to the database.
func (c *Call) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.calls (` +
		`calldate, calltime, calldestination, calldestnumber, callratingperiod, callduration, callcarrierband, callwholesale, callretail, callcalcwholesal, callduplicate, callfileid, callcarrier, callactretail, callactretailpre, calldesttype, callourband, callcustbillingg, callpromocode, callextn, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Call in the database.
func (c *Call) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.calls SET (` +
		`calldate, calltime, calldestination, calldestnumber, callratingperiod, callduration, callcarrierband, callwholesale, callretail, callcalcwholesal, callduplicate, callfileid, callcarrier, callactretail, callactretailpre, calldesttype, callourband, callcustbillingg, callpromocode, callextn, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22` +
		`) WHERE equinox_lrn = $23`

	// run query
	XOLog(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Call to the database.
func (c *Call) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Call.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Call) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.calls (` +
		`calldate, calltime, calldestination, calldestnumber, callratingperiod, callduration, callcarrierband, callwholesale, callretail, callcalcwholesal, callduplicate, callfileid, callcarrier, callactretail, callactretailpre, calldesttype, callourband, callcustbillingg, callpromocode, callextn, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`calldate, calltime, calldestination, calldestnumber, callratingperiod, callduration, callcarrierband, callwholesale, callretail, callcalcwholesal, callduplicate, callfileid, callcarrier, callactretail, callactretailpre, calldesttype, callourband, callcustbillingg, callpromocode, callextn, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.calldate, EXCLUDED.calltime, EXCLUDED.calldestination, EXCLUDED.calldestnumber, EXCLUDED.callratingperiod, EXCLUDED.callduration, EXCLUDED.callcarrierband, EXCLUDED.callwholesale, EXCLUDED.callretail, EXCLUDED.callcalcwholesal, EXCLUDED.callduplicate, EXCLUDED.callfileid, EXCLUDED.callcarrier, EXCLUDED.callactretail, EXCLUDED.callactretailpre, EXCLUDED.calldesttype, EXCLUDED.callourband, EXCLUDED.callcustbillingg, EXCLUDED.callpromocode, EXCLUDED.callextn, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Calldate, c.Calltime, c.Calldestination, c.Calldestnumber, c.Callratingperiod, c.Callduration, c.Callcarrierband, c.Callwholesale, c.Callretail, c.Callcalcwholesal, c.Callduplicate, c.Callfileid, c.Callcarrier, c.Callactretail, c.Callactretailpre, c.Calldesttype, c.Callourband, c.Callcustbillingg, c.Callpromocode, c.Callextn, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Call from the database.
func (c *Call) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.calls WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CallByEquinoxLrn retrieves a row from 'equinox.calls' as a Call.
//
// Generated from index 'calls_pkey'.
func CallByEquinoxLrn(db XODB, equinoxLrn int64) (*Call, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`calldate, calltime, calldestination, calldestnumber, callratingperiod, callduration, callcarrierband, callwholesale, callretail, callcalcwholesal, callduplicate, callfileid, callcarrier, callactretail, callactretailpre, calldesttype, callourband, callcustbillingg, callpromocode, callextn, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.calls ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Call{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Calldate, &c.Calltime, &c.Calldestination, &c.Calldestnumber, &c.Callratingperiod, &c.Callduration, &c.Callcarrierband, &c.Callwholesale, &c.Callretail, &c.Callcalcwholesal, &c.Callduplicate, &c.Callfileid, &c.Callcarrier, &c.Callactretail, &c.Callactretailpre, &c.Calldesttype, &c.Callourband, &c.Callcustbillingg, &c.Callpromocode, &c.Callextn, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}