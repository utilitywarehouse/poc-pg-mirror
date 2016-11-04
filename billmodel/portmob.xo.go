// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Portmob represents a row from 'equinox.portmob'.
type Portmob struct {
	Portmobcli       sql.NullString  `json:"portmobcli"`       // portmobcli
	Portmobsim       sql.NullString  `json:"portmobsim"`       // portmobsim
	Portmobcover     sql.NullInt64   `json:"portmobcover"`     // portmobcover
	Portmobvmail     sql.NullInt64   `json:"portmobvmail"`     // portmobvmail
	Portmobsecurity  sql.NullString  `json:"portmobsecurity"`  // portmobsecurity
	Portmobpackageno sql.NullString  `json:"portmobpackageno"` // portmobpackageno
	Portdspreject    sql.NullString  `json:"portdspreject"`    // portdspreject
	Portmobequip     pq.NullTime     `json:"portmobequip"`     // portmobequip
	Portmodel        sql.NullString  `json:"portmodel"`        // portmodel
	Portmobpromocode sql.NullString  `json:"portmobpromocode"` // portmobpromocode
	Portpackagecost  sql.NullFloat64 `json:"portpackagecost"`  // portpackagecost
	Portmoboutso     pq.NullTime     `json:"portmoboutso"`     // portmoboutso
	Portoriginalcli  sql.NullString  `json:"portoriginalcli"`  // portoriginalcli
	Portcharge       sql.NullInt64   `json:"portcharge"`       // portcharge
	Portoptrpc       sql.NullInt64   `json:"portoptrpc"`       // portoptrpc
	Portoptsms       sql.NullInt64   `json:"portoptsms"`       // portoptsms
	Portoptmms       sql.NullInt64   `json:"portoptmms"`       // portoptmms
	Portoptgprs      sql.NullInt64   `json:"portoptgprs"`      // portoptgprs
	Portoptlocal     sql.NullInt64   `json:"portoptlocal"`     // portoptlocal
	Portoptintern    sql.NullInt64   `json:"portoptintern"`    // portoptintern
	Portunlock       sql.NullString  `json:"portunlock"`       // portunlock
	Portmobsparen1   sql.NullFloat64 `json:"portmobsparen1"`   // portmobsparen1
	Portmobsparen2   sql.NullFloat64 `json:"portmobsparen2"`   // portmobsparen2
	Portrpccode      sql.NullString  `json:"portrpccode"`      // portrpccode
	Portsmscode      sql.NullString  `json:"portsmscode"`      // portsmscode
	Portmmscode      sql.NullString  `json:"portmmscode"`      // portmmscode
	Portgprscode     sql.NullString  `json:"portgprscode"`     // portgprscode
	Portlocalcode    sql.NullString  `json:"portlocalcode"`    // portlocalcode
	Portintlcode     sql.NullString  `json:"portintlcode"`     // portintlcode
	Portcontract     sql.NullInt64   `json:"portcontract"`     // portcontract
	Portdatabundle   sql.NullString  `json:"portdatabundle"`   // portdatabundle
	Portbolton       sql.NullString  `json:"portbolton"`       // portbolton
	Portsmartphone   sql.NullFloat64 `json:"portsmartphone"`   // portsmartphone
	Portspecialdeal  sql.NullInt64   `json:"portspecialdeal"`  // portspecialdeal
	Portcampaign     sql.NullString  `json:"portcampaign"`     // portcampaign
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Portmob exists in the database.
func (p *Portmob) Exists() bool {
	return p._exists
}

// Deleted provides information if the Portmob has been deleted from the database.
func (p *Portmob) Deleted() bool {
	return p._deleted
}

// Insert inserts the Portmob to the database.
func (p *Portmob) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.portmob (` +
		`portmobcli, portmobsim, portmobcover, portmobvmail, portmobsecurity, portmobpackageno, portdspreject, portmobequip, portmodel, portmobpromocode, portpackagecost, portmoboutso, portoriginalcli, portcharge, portoptrpc, portoptsms, portoptmms, portoptgprs, portoptlocal, portoptintern, portunlock, portmobsparen1, portmobsparen2, portrpccode, portsmscode, portmmscode, portgprscode, portlocalcode, portintlcode, portcontract, portdatabundle, portbolton, portsmartphone, portspecialdeal, portcampaign, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Portmob in the database.
func (p *Portmob) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.portmob SET (` +
		`portmobcli, portmobsim, portmobcover, portmobvmail, portmobsecurity, portmobpackageno, portdspreject, portmobequip, portmodel, portmobpromocode, portpackagecost, portmoboutso, portoriginalcli, portcharge, portoptrpc, portoptsms, portoptmms, portoptgprs, portoptlocal, portoptintern, portunlock, portmobsparen1, portmobsparen2, portrpccode, portsmscode, portmmscode, portgprscode, portlocalcode, portintlcode, portcontract, portdatabundle, portbolton, portsmartphone, portspecialdeal, portcampaign, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37` +
		`) WHERE equinox_lrn = $38`

	// run query
	XOLog(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Portmob to the database.
func (p *Portmob) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Portmob.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Portmob) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.portmob (` +
		`portmobcli, portmobsim, portmobcover, portmobvmail, portmobsecurity, portmobpackageno, portdspreject, portmobequip, portmodel, portmobpromocode, portpackagecost, portmoboutso, portoriginalcli, portcharge, portoptrpc, portoptsms, portoptmms, portoptgprs, portoptlocal, portoptintern, portunlock, portmobsparen1, portmobsparen2, portrpccode, portsmscode, portmmscode, portgprscode, portlocalcode, portintlcode, portcontract, portdatabundle, portbolton, portsmartphone, portspecialdeal, portcampaign, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`portmobcli, portmobsim, portmobcover, portmobvmail, portmobsecurity, portmobpackageno, portdspreject, portmobequip, portmodel, portmobpromocode, portpackagecost, portmoboutso, portoriginalcli, portcharge, portoptrpc, portoptsms, portoptmms, portoptgprs, portoptlocal, portoptintern, portunlock, portmobsparen1, portmobsparen2, portrpccode, portsmscode, portmmscode, portgprscode, portlocalcode, portintlcode, portcontract, portdatabundle, portbolton, portsmartphone, portspecialdeal, portcampaign, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.portmobcli, EXCLUDED.portmobsim, EXCLUDED.portmobcover, EXCLUDED.portmobvmail, EXCLUDED.portmobsecurity, EXCLUDED.portmobpackageno, EXCLUDED.portdspreject, EXCLUDED.portmobequip, EXCLUDED.portmodel, EXCLUDED.portmobpromocode, EXCLUDED.portpackagecost, EXCLUDED.portmoboutso, EXCLUDED.portoriginalcli, EXCLUDED.portcharge, EXCLUDED.portoptrpc, EXCLUDED.portoptsms, EXCLUDED.portoptmms, EXCLUDED.portoptgprs, EXCLUDED.portoptlocal, EXCLUDED.portoptintern, EXCLUDED.portunlock, EXCLUDED.portmobsparen1, EXCLUDED.portmobsparen2, EXCLUDED.portrpccode, EXCLUDED.portsmscode, EXCLUDED.portmmscode, EXCLUDED.portgprscode, EXCLUDED.portlocalcode, EXCLUDED.portintlcode, EXCLUDED.portcontract, EXCLUDED.portdatabundle, EXCLUDED.portbolton, EXCLUDED.portsmartphone, EXCLUDED.portspecialdeal, EXCLUDED.portcampaign, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Portmobcli, p.Portmobsim, p.Portmobcover, p.Portmobvmail, p.Portmobsecurity, p.Portmobpackageno, p.Portdspreject, p.Portmobequip, p.Portmodel, p.Portmobpromocode, p.Portpackagecost, p.Portmoboutso, p.Portoriginalcli, p.Portcharge, p.Portoptrpc, p.Portoptsms, p.Portoptmms, p.Portoptgprs, p.Portoptlocal, p.Portoptintern, p.Portunlock, p.Portmobsparen1, p.Portmobsparen2, p.Portrpccode, p.Portsmscode, p.Portmmscode, p.Portgprscode, p.Portlocalcode, p.Portintlcode, p.Portcontract, p.Portdatabundle, p.Portbolton, p.Portsmartphone, p.Portspecialdeal, p.Portcampaign, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Portmob from the database.
func (p *Portmob) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.portmob WHERE equinox_lrn = $1`

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

// PortmobByEquinoxLrn retrieves a row from 'equinox.portmob' as a Portmob.
//
// Generated from index 'portmob_pkey'.
func PortmobByEquinoxLrn(db XODB, equinoxLrn int64) (*Portmob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`portmobcli, portmobsim, portmobcover, portmobvmail, portmobsecurity, portmobpackageno, portdspreject, portmobequip, portmodel, portmobpromocode, portpackagecost, portmoboutso, portoriginalcli, portcharge, portoptrpc, portoptsms, portoptmms, portoptgprs, portoptlocal, portoptintern, portunlock, portmobsparen1, portmobsparen2, portrpccode, portsmscode, portmmscode, portgprscode, portlocalcode, portintlcode, portcontract, portdatabundle, portbolton, portsmartphone, portspecialdeal, portcampaign, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.portmob ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Portmob{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Portmobcli, &p.Portmobsim, &p.Portmobcover, &p.Portmobvmail, &p.Portmobsecurity, &p.Portmobpackageno, &p.Portdspreject, &p.Portmobequip, &p.Portmodel, &p.Portmobpromocode, &p.Portpackagecost, &p.Portmoboutso, &p.Portoriginalcli, &p.Portcharge, &p.Portoptrpc, &p.Portoptsms, &p.Portoptmms, &p.Portoptgprs, &p.Portoptlocal, &p.Portoptintern, &p.Portunlock, &p.Portmobsparen1, &p.Portmobsparen2, &p.Portrpccode, &p.Portsmscode, &p.Portmmscode, &p.Portgprscode, &p.Portlocalcode, &p.Portintlcode, &p.Portcontract, &p.Portdatabundle, &p.Portbolton, &p.Portsmartphone, &p.Portspecialdeal, &p.Portcampaign, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}