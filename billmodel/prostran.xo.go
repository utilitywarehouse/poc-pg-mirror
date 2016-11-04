// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Prostran represents a row from 'equinox.prostran'.
type Prostran struct {
	Pttransactionid  sql.NullInt64   `json:"pttransactionid"`  // pttransactionid
	Ptdate           pq.NullTime     `json:"ptdate"`           // ptdate
	Ptdatesettled    pq.NullTime     `json:"ptdatesettled"`    // ptdatesettled
	Ptdateprocessed  pq.NullTime     `json:"ptdateprocessed"`  // ptdateprocessed
	Pttranstype      sql.NullString  `json:"pttranstype"`      // pttranstype
	Ptlocaliso       sql.NullString  `json:"ptlocaliso"`       // ptlocaliso
	Ptvaluegbp       sql.NullFloat64 `json:"ptvaluegbp"`       // ptvaluegbp
	Ptvaluelocal     sql.NullFloat64 `json:"ptvaluelocal"`     // ptvaluelocal
	Ptauthcode       sql.NullString  `json:"ptauthcode"`       // ptauthcode
	Ptmerchantnum    sql.NullString  `json:"ptmerchantnum"`    // ptmerchantnum
	Ptdescription    sql.NullString  `json:"ptdescription"`    // ptdescription
	Ptmerchantcat    sql.NullString  `json:"ptmerchantcat"`    // ptmerchantcat
	Ptsaving         sql.NullFloat64 `json:"ptsaving"`         // ptsaving
	Ptbillnumber     sql.NullInt64   `json:"ptbillnumber"`     // ptbillnumber
	Ptexchangerate   sql.NullFloat64 `json:"ptexchangerate"`   // ptexchangerate
	Pttransactionchg sql.NullFloat64 `json:"pttransactionchg"` // pttransactionchg
	Ptpptreference   sql.NullString  `json:"ptpptreference"`   // ptpptreference
	Ptprossave       sql.NullFloat64 `json:"ptprossave"`       // ptprossave
	Ptprospercent    sql.NullFloat64 `json:"ptprospercent"`    // ptprospercent
	Ptaffmer         sql.NullInt64   `json:"ptaffmer"`         // ptaffmer
	Pteligiblespend  sql.NullFloat64 `json:"pteligiblespend"`  // pteligiblespend
	Ptdateclaimed    pq.NullTime     `json:"ptdateclaimed"`    // ptdateclaimed
	Ptdontpay        sql.NullString  `json:"ptdontpay"`        // ptdontpay
	Ptpaidfuel       sql.NullString  `json:"ptpaidfuel"`       // ptpaidfuel
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Prostran exists in the database.
func (p *Prostran) Exists() bool {
	return p._exists
}

// Deleted provides information if the Prostran has been deleted from the database.
func (p *Prostran) Deleted() bool {
	return p._deleted
}

// Insert inserts the Prostran to the database.
func (p *Prostran) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.prostran (` +
		`pttransactionid, ptdate, ptdatesettled, ptdateprocessed, pttranstype, ptlocaliso, ptvaluegbp, ptvaluelocal, ptauthcode, ptmerchantnum, ptdescription, ptmerchantcat, ptsaving, ptbillnumber, ptexchangerate, pttransactionchg, ptpptreference, ptprossave, ptprospercent, ptaffmer, pteligiblespend, ptdateclaimed, ptdontpay, ptpaidfuel, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Prostran in the database.
func (p *Prostran) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.prostran SET (` +
		`pttransactionid, ptdate, ptdatesettled, ptdateprocessed, pttranstype, ptlocaliso, ptvaluegbp, ptvaluelocal, ptauthcode, ptmerchantnum, ptdescription, ptmerchantcat, ptsaving, ptbillnumber, ptexchangerate, pttransactionchg, ptpptreference, ptprossave, ptprospercent, ptaffmer, pteligiblespend, ptdateclaimed, ptdontpay, ptpaidfuel, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`) WHERE equinox_lrn = $27`

	// run query
	XOLog(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Prostran to the database.
func (p *Prostran) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Prostran.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Prostran) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.prostran (` +
		`pttransactionid, ptdate, ptdatesettled, ptdateprocessed, pttranstype, ptlocaliso, ptvaluegbp, ptvaluelocal, ptauthcode, ptmerchantnum, ptdescription, ptmerchantcat, ptsaving, ptbillnumber, ptexchangerate, pttransactionchg, ptpptreference, ptprossave, ptprospercent, ptaffmer, pteligiblespend, ptdateclaimed, ptdontpay, ptpaidfuel, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`pttransactionid, ptdate, ptdatesettled, ptdateprocessed, pttranstype, ptlocaliso, ptvaluegbp, ptvaluelocal, ptauthcode, ptmerchantnum, ptdescription, ptmerchantcat, ptsaving, ptbillnumber, ptexchangerate, pttransactionchg, ptpptreference, ptprossave, ptprospercent, ptaffmer, pteligiblespend, ptdateclaimed, ptdontpay, ptpaidfuel, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.pttransactionid, EXCLUDED.ptdate, EXCLUDED.ptdatesettled, EXCLUDED.ptdateprocessed, EXCLUDED.pttranstype, EXCLUDED.ptlocaliso, EXCLUDED.ptvaluegbp, EXCLUDED.ptvaluelocal, EXCLUDED.ptauthcode, EXCLUDED.ptmerchantnum, EXCLUDED.ptdescription, EXCLUDED.ptmerchantcat, EXCLUDED.ptsaving, EXCLUDED.ptbillnumber, EXCLUDED.ptexchangerate, EXCLUDED.pttransactionchg, EXCLUDED.ptpptreference, EXCLUDED.ptprossave, EXCLUDED.ptprospercent, EXCLUDED.ptaffmer, EXCLUDED.pteligiblespend, EXCLUDED.ptdateclaimed, EXCLUDED.ptdontpay, EXCLUDED.ptpaidfuel, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Pttransactionid, p.Ptdate, p.Ptdatesettled, p.Ptdateprocessed, p.Pttranstype, p.Ptlocaliso, p.Ptvaluegbp, p.Ptvaluelocal, p.Ptauthcode, p.Ptmerchantnum, p.Ptdescription, p.Ptmerchantcat, p.Ptsaving, p.Ptbillnumber, p.Ptexchangerate, p.Pttransactionchg, p.Ptpptreference, p.Ptprossave, p.Ptprospercent, p.Ptaffmer, p.Pteligiblespend, p.Ptdateclaimed, p.Ptdontpay, p.Ptpaidfuel, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Prostran from the database.
func (p *Prostran) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.prostran WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// ProstranByEquinoxLrn retrieves a row from 'equinox.prostran' as a Prostran.
//
// Generated from index 'prostran_pkey'.
func ProstranByEquinoxLrn(db XODB, equinoxLrn int64) (*Prostran, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pttransactionid, ptdate, ptdatesettled, ptdateprocessed, pttranstype, ptlocaliso, ptvaluegbp, ptvaluelocal, ptauthcode, ptmerchantnum, ptdescription, ptmerchantcat, ptsaving, ptbillnumber, ptexchangerate, pttransactionchg, ptpptreference, ptprossave, ptprospercent, ptaffmer, pteligiblespend, ptdateclaimed, ptdontpay, ptpaidfuel, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.prostran ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Prostran{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Pttransactionid, &p.Ptdate, &p.Ptdatesettled, &p.Ptdateprocessed, &p.Pttranstype, &p.Ptlocaliso, &p.Ptvaluegbp, &p.Ptvaluelocal, &p.Ptauthcode, &p.Ptmerchantnum, &p.Ptdescription, &p.Ptmerchantcat, &p.Ptsaving, &p.Ptbillnumber, &p.Ptexchangerate, &p.Pttransactionchg, &p.Ptpptreference, &p.Ptprossave, &p.Ptprospercent, &p.Ptaffmer, &p.Pteligiblespend, &p.Ptdateclaimed, &p.Ptdontpay, &p.Ptpaidfuel, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
