// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Comcode represents a row from 'equinox.comcodes'.
type Comcode struct {
	Comcodeid       sql.NullInt64   `json:"comcodeid"`       // comcodeid
	Cccode          sql.NullString  `json:"cccode"`          // cccode
	Ccpart1         sql.NullString  `json:"ccpart1"`         // ccpart1
	Ccpart2         sql.NullString  `json:"ccpart2"`         // ccpart2
	Ccdescription   sql.NullString  `json:"ccdescription"`   // ccdescription
	Ccdisplaytext   sql.NullString  `json:"ccdisplaytext"`   // ccdisplaytext
	Cccommistype    sql.NullString  `json:"cccommistype"`    // cccommistype
	Cctotalingtype  sql.NullString  `json:"cctotalingtype"`  // cctotalingtype
	Cctotalto       sql.NullInt64   `json:"cctotalto"`       // cctotalto
	Cccomheadfield  sql.NullString  `json:"cccomheadfield"`  // cccomheadfield
	Ccdisplayorder  sql.NullInt64   `json:"ccdisplayorder"`  // ccdisplayorder
	Ccsubtariffband sql.NullInt64   `json:"ccsubtariffband"` // ccsubtariffband
	Ccunique        sql.NullString  `json:"ccunique"`        // ccunique
	Ccuniquefield   sql.NullString  `json:"ccuniquefield"`   // ccuniquefield
	Ccactive        sql.NullInt64   `json:"ccactive"`        // ccactive
	Ccusage         sql.NullString  `json:"ccusage"`         // ccusage
	Ccapplyvat      sql.NullInt64   `json:"ccapplyvat"`      // ccapplyvat
	Ccdowns         sql.NullInt64   `json:"ccdowns"`         // ccdowns
	Ccdowne         sql.NullInt64   `json:"ccdowne"`         // ccdowne
	Cclevel         sql.NullInt64   `json:"cclevel"`         // cclevel
	Ccrootentry     sql.NullString  `json:"ccrootentry"`     // ccrootentry
	Ccsparechar1    sql.NullString  `json:"ccsparechar1"`    // ccsparechar1
	Ccsparechar2    sql.NullString  `json:"ccsparechar2"`    // ccsparechar2
	Ccsparechar3    sql.NullString  `json:"ccsparechar3"`    // ccsparechar3
	Ccsparenum1     sql.NullFloat64 `json:"ccsparenum1"`     // ccsparenum1
	Ccsparenum2     sql.NullFloat64 `json:"ccsparenum2"`     // ccsparenum2
	Ccsparenum3     sql.NullFloat64 `json:"ccsparenum3"`     // ccsparenum3
	Ccsparedate1    pq.NullTime     `json:"ccsparedate1"`    // ccsparedate1
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Comcode exists in the database.
func (c *Comcode) Exists() bool {
	return c._exists
}

// Deleted provides information if the Comcode has been deleted from the database.
func (c *Comcode) Deleted() bool {
	return c._deleted
}

// Insert inserts the Comcode to the database.
func (c *Comcode) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.comcodes (` +
		`comcodeid, cccode, ccpart1, ccpart2, ccdescription, ccdisplaytext, cccommistype, cctotalingtype, cctotalto, cccomheadfield, ccdisplayorder, ccsubtariffband, ccunique, ccuniquefield, ccactive, ccusage, ccapplyvat, ccdowns, ccdowne, cclevel, ccrootentry, ccsparechar1, ccsparechar2, ccsparechar3, ccsparenum1, ccsparenum2, ccsparenum3, ccsparedate1, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Comcode in the database.
func (c *Comcode) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.comcodes SET (` +
		`comcodeid, cccode, ccpart1, ccpart2, ccdescription, ccdisplaytext, cccommistype, cctotalingtype, cctotalto, cccomheadfield, ccdisplayorder, ccsubtariffband, ccunique, ccuniquefield, ccactive, ccusage, ccapplyvat, ccdowns, ccdowne, cclevel, ccrootentry, ccsparechar1, ccsparechar2, ccsparechar3, ccsparenum1, ccsparenum2, ccsparenum3, ccsparedate1, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29` +
		`) WHERE equinox_lrn = $30`

	// run query
	XOLog(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Comcode to the database.
func (c *Comcode) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Comcode.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Comcode) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.comcodes (` +
		`comcodeid, cccode, ccpart1, ccpart2, ccdescription, ccdisplaytext, cccommistype, cctotalingtype, cctotalto, cccomheadfield, ccdisplayorder, ccsubtariffband, ccunique, ccuniquefield, ccactive, ccusage, ccapplyvat, ccdowns, ccdowne, cclevel, ccrootentry, ccsparechar1, ccsparechar2, ccsparechar3, ccsparenum1, ccsparenum2, ccsparenum3, ccsparedate1, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`comcodeid, cccode, ccpart1, ccpart2, ccdescription, ccdisplaytext, cccommistype, cctotalingtype, cctotalto, cccomheadfield, ccdisplayorder, ccsubtariffband, ccunique, ccuniquefield, ccactive, ccusage, ccapplyvat, ccdowns, ccdowne, cclevel, ccrootentry, ccsparechar1, ccsparechar2, ccsparechar3, ccsparenum1, ccsparenum2, ccsparenum3, ccsparedate1, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.comcodeid, EXCLUDED.cccode, EXCLUDED.ccpart1, EXCLUDED.ccpart2, EXCLUDED.ccdescription, EXCLUDED.ccdisplaytext, EXCLUDED.cccommistype, EXCLUDED.cctotalingtype, EXCLUDED.cctotalto, EXCLUDED.cccomheadfield, EXCLUDED.ccdisplayorder, EXCLUDED.ccsubtariffband, EXCLUDED.ccunique, EXCLUDED.ccuniquefield, EXCLUDED.ccactive, EXCLUDED.ccusage, EXCLUDED.ccapplyvat, EXCLUDED.ccdowns, EXCLUDED.ccdowne, EXCLUDED.cclevel, EXCLUDED.ccrootentry, EXCLUDED.ccsparechar1, EXCLUDED.ccsparechar2, EXCLUDED.ccsparechar3, EXCLUDED.ccsparenum1, EXCLUDED.ccsparenum2, EXCLUDED.ccsparenum3, EXCLUDED.ccsparedate1, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Comcodeid, c.Cccode, c.Ccpart1, c.Ccpart2, c.Ccdescription, c.Ccdisplaytext, c.Cccommistype, c.Cctotalingtype, c.Cctotalto, c.Cccomheadfield, c.Ccdisplayorder, c.Ccsubtariffband, c.Ccunique, c.Ccuniquefield, c.Ccactive, c.Ccusage, c.Ccapplyvat, c.Ccdowns, c.Ccdowne, c.Cclevel, c.Ccrootentry, c.Ccsparechar1, c.Ccsparechar2, c.Ccsparechar3, c.Ccsparenum1, c.Ccsparenum2, c.Ccsparenum3, c.Ccsparedate1, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Comcode from the database.
func (c *Comcode) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.comcodes WHERE equinox_lrn = $1`

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

// ComcodeByEquinoxLrn retrieves a row from 'equinox.comcodes' as a Comcode.
//
// Generated from index 'comcodes_pkey'.
func ComcodeByEquinoxLrn(db XODB, equinoxLrn int64) (*Comcode, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`comcodeid, cccode, ccpart1, ccpart2, ccdescription, ccdisplaytext, cccommistype, cctotalingtype, cctotalto, cccomheadfield, ccdisplayorder, ccsubtariffband, ccunique, ccuniquefield, ccactive, ccusage, ccapplyvat, ccdowns, ccdowne, cclevel, ccrootentry, ccsparechar1, ccsparechar2, ccsparechar3, ccsparenum1, ccsparenum2, ccsparenum3, ccsparedate1, equinox_lrn, equinox_sec ` +
		`FROM equinox.comcodes ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Comcode{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Comcodeid, &c.Cccode, &c.Ccpart1, &c.Ccpart2, &c.Ccdescription, &c.Ccdisplaytext, &c.Cccommistype, &c.Cctotalingtype, &c.Cctotalto, &c.Cccomheadfield, &c.Ccdisplayorder, &c.Ccsubtariffband, &c.Ccunique, &c.Ccuniquefield, &c.Ccactive, &c.Ccusage, &c.Ccapplyvat, &c.Ccdowns, &c.Ccdowne, &c.Cclevel, &c.Ccrootentry, &c.Ccsparechar1, &c.Ccsparechar2, &c.Ccsparechar3, &c.Ccsparenum1, &c.Ccsparenum2, &c.Ccsparenum3, &c.Ccsparedate1, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
