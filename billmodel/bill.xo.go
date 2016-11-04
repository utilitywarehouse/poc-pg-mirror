// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Bill represents a row from 'equinox.bills'.
type Bill struct {
	Billnumber      sql.NullFloat64 `json:"billnumber"`      // billnumber
	Billdate        pq.NullTime     `json:"billdate"`        // billdate
	Billtotretail   sql.NullFloat64 `json:"billtotretail"`   // billtotretail
	Billtotretail2  sql.NullFloat64 `json:"billtotretail2"`  // billtotretail2
	Billtotwhole    sql.NullFloat64 `json:"billtotwhole"`    // billtotwhole
	Billtotmnth     sql.NullFloat64 `json:"billtotmnth"`     // billtotmnth
	Billtotmnthw    sql.NullFloat64 `json:"billtotmnthw"`    // billtotmnthw
	Billtotoneoff   sql.NullFloat64 `json:"billtotoneoff"`   // billtotoneoff
	Billgreendeal   sql.NullFloat64 `json:"billgreendeal"`   // billgreendeal
	Billdebitos     sql.NullFloat64 `json:"billdebitos"`     // billdebitos
	Billcreditstd   sql.NullFloat64 `json:"billcreditstd"`   // billcreditstd
	Billcreditgood  sql.NullFloat64 `json:"billcreditgood"`  // billcreditgood
	Billdebitstd    sql.NullFloat64 `json:"billdebitstd"`    // billdebitstd
	Billsurcharge   sql.NullFloat64 `json:"billsurcharge"`   // billsurcharge
	Billdiscount    sql.NullFloat64 `json:"billdiscount"`    // billdiscount
	Billtotal       sql.NullFloat64 `json:"billtotal"`       // billtotal
	Billvattp       sql.NullFloat64 `json:"billvattp"`       // billvattp
	Billvatgp       sql.NullFloat64 `json:"billvatgp"`       // billvatgp
	Billvatep       sql.NullFloat64 `json:"billvatep"`       // billvatep
	Billnettp       sql.NullFloat64 `json:"billnettp"`       // billnettp
	Billnetgp       sql.NullFloat64 `json:"billnetgp"`       // billnetgp
	Billnetep       sql.NullFloat64 `json:"billnetep"`       // billnetep
	Billcomplete    sql.NullString  `json:"billcomplete"`    // billcomplete
	Billperiod      sql.NullString  `json:"billperiod"`      // billperiod
	Billddchange    sql.NullFloat64 `json:"billddchange"`    // billddchange
	Billcgadisc     sql.NullFloat64 `json:"billcgadisc"`     // billcgadisc
	Billcardcashbak sql.NullFloat64 `json:"billcardcashbak"` // billcardcashbak
	Billaffcashbak  sql.NullFloat64 `json:"billaffcashbak"`  // billaffcashbak
	Billclubfee     sql.NullFloat64 `json:"billclubfee"`     // billclubfee
	Billclubdisc    sql.NullFloat64 `json:"billclubdisc"`    // billclubdisc
	Billcommispaid  pq.NullTime     `json:"billcommispaid"`  // billcommispaid
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Bill exists in the database.
func (b *Bill) Exists() bool {
	return b._exists
}

// Deleted provides information if the Bill has been deleted from the database.
func (b *Bill) Deleted() bool {
	return b._deleted
}

// Insert inserts the Bill to the database.
func (b *Bill) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.bills (` +
		`billnumber, billdate, billtotretail, billtotretail2, billtotwhole, billtotmnth, billtotmnthw, billtotoneoff, billgreendeal, billdebitos, billcreditstd, billcreditgood, billdebitstd, billsurcharge, billdiscount, billtotal, billvattp, billvatgp, billvatep, billnettp, billnetgp, billnetep, billcomplete, billperiod, billddchange, billcgadisc, billcardcashbak, billaffcashbak, billclubfee, billclubdisc, billcommispaid, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxSec)
	err = db.QueryRow(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxSec).Scan(&b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Bill in the database.
func (b *Bill) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.bills SET (` +
		`billnumber, billdate, billtotretail, billtotretail2, billtotwhole, billtotmnth, billtotmnthw, billtotoneoff, billgreendeal, billdebitos, billcreditstd, billcreditgood, billdebitstd, billsurcharge, billdiscount, billtotal, billvattp, billvatgp, billvatep, billnettp, billnetgp, billnetep, billcomplete, billperiod, billddchange, billcgadisc, billcardcashbak, billaffcashbak, billclubfee, billclubdisc, billcommispaid, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33` +
		`) WHERE equinox_lrn = $34`

	// run query
	XOLog(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxSec, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxSec, b.EquinoxLrn)
	return err
}

// Save saves the Bill to the database.
func (b *Bill) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Bill.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Bill) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.bills (` +
		`billnumber, billdate, billtotretail, billtotretail2, billtotwhole, billtotmnth, billtotmnthw, billtotoneoff, billgreendeal, billdebitos, billcreditstd, billcreditgood, billdebitstd, billsurcharge, billdiscount, billtotal, billvattp, billvatgp, billvatep, billnettp, billnetgp, billnetep, billcomplete, billperiod, billddchange, billcgadisc, billcardcashbak, billaffcashbak, billclubfee, billclubdisc, billcommispaid, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`billnumber, billdate, billtotretail, billtotretail2, billtotwhole, billtotmnth, billtotmnthw, billtotoneoff, billgreendeal, billdebitos, billcreditstd, billcreditgood, billdebitstd, billsurcharge, billdiscount, billtotal, billvattp, billvatgp, billvatep, billnettp, billnetgp, billnetep, billcomplete, billperiod, billddchange, billcgadisc, billcardcashbak, billaffcashbak, billclubfee, billclubdisc, billcommispaid, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.billnumber, EXCLUDED.billdate, EXCLUDED.billtotretail, EXCLUDED.billtotretail2, EXCLUDED.billtotwhole, EXCLUDED.billtotmnth, EXCLUDED.billtotmnthw, EXCLUDED.billtotoneoff, EXCLUDED.billgreendeal, EXCLUDED.billdebitos, EXCLUDED.billcreditstd, EXCLUDED.billcreditgood, EXCLUDED.billdebitstd, EXCLUDED.billsurcharge, EXCLUDED.billdiscount, EXCLUDED.billtotal, EXCLUDED.billvattp, EXCLUDED.billvatgp, EXCLUDED.billvatep, EXCLUDED.billnettp, EXCLUDED.billnetgp, EXCLUDED.billnetep, EXCLUDED.billcomplete, EXCLUDED.billperiod, EXCLUDED.billddchange, EXCLUDED.billcgadisc, EXCLUDED.billcardcashbak, EXCLUDED.billaffcashbak, EXCLUDED.billclubfee, EXCLUDED.billclubdisc, EXCLUDED.billcommispaid, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxLrn, b.EquinoxSec)
	_, err = db.Exec(sqlstr, b.Billnumber, b.Billdate, b.Billtotretail, b.Billtotretail2, b.Billtotwhole, b.Billtotmnth, b.Billtotmnthw, b.Billtotoneoff, b.Billgreendeal, b.Billdebitos, b.Billcreditstd, b.Billcreditgood, b.Billdebitstd, b.Billsurcharge, b.Billdiscount, b.Billtotal, b.Billvattp, b.Billvatgp, b.Billvatep, b.Billnettp, b.Billnetgp, b.Billnetep, b.Billcomplete, b.Billperiod, b.Billddchange, b.Billcgadisc, b.Billcardcashbak, b.Billaffcashbak, b.Billclubfee, b.Billclubdisc, b.Billcommispaid, b.EquinoxPrn, b.EquinoxLrn, b.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Delete deletes the Bill from the database.
func (b *Bill) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.bills WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// BillByEquinoxLrn retrieves a row from 'equinox.bills' as a Bill.
//
// Generated from index 'bills_pkey'.
func BillByEquinoxLrn(db XODB, equinoxLrn int64) (*Bill, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`billnumber, billdate, billtotretail, billtotretail2, billtotwhole, billtotmnth, billtotmnthw, billtotoneoff, billgreendeal, billdebitos, billcreditstd, billcreditgood, billdebitstd, billsurcharge, billdiscount, billtotal, billvattp, billvatgp, billvatep, billnettp, billnetgp, billnetep, billcomplete, billperiod, billddchange, billcgadisc, billcardcashbak, billaffcashbak, billclubfee, billclubdisc, billcommispaid, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.bills ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Bill{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Billnumber, &b.Billdate, &b.Billtotretail, &b.Billtotretail2, &b.Billtotwhole, &b.Billtotmnth, &b.Billtotmnthw, &b.Billtotoneoff, &b.Billgreendeal, &b.Billdebitos, &b.Billcreditstd, &b.Billcreditgood, &b.Billdebitstd, &b.Billsurcharge, &b.Billdiscount, &b.Billtotal, &b.Billvattp, &b.Billvatgp, &b.Billvatep, &b.Billnettp, &b.Billnetgp, &b.Billnetep, &b.Billcomplete, &b.Billperiod, &b.Billddchange, &b.Billcgadisc, &b.Billcardcashbak, &b.Billaffcashbak, &b.Billclubfee, &b.Billclubdisc, &b.Billcommispaid, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}