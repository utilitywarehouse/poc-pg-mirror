// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Cnxmobil represents a row from 'equinox.cnxmobil'.
type Cnxmobil struct {
	Cnxmobcli        sql.NullString  `json:"cnxmobcli"`        // cnxmobcli
	Cnxmobsim        sql.NullString  `json:"cnxmobsim"`        // cnxmobsim
	Cnxmobiemi       sql.NullString  `json:"cnxmobiemi"`       // cnxmobiemi
	Cnxmobphonereqd  sql.NullInt64   `json:"cnxmobphonereqd"`  // cnxmobphonereqd
	Cnxmobcover      sql.NullInt64   `json:"cnxmobcover"`      // cnxmobcover
	Cnxmobvmail      sql.NullInt64   `json:"cnxmobvmail"`      // cnxmobvmail
	Cnxmobsecurity   sql.NullString  `json:"cnxmobsecurity"`   // cnxmobsecurity
	Cnxmobdateentrd  pq.NullTime     `json:"cnxmobdateentrd"`  // cnxmobdateentrd
	Cnxmobdateorder  pq.NullTime     `json:"cnxmobdateorder"`  // cnxmobdateorder
	Cnxmobdateivrreg pq.NullTime     `json:"cnxmobdateivrreg"` // cnxmobdateivrreg
	Cnxmobdatetompc  pq.NullTime     `json:"cnxmobdatetompc"`  // cnxmobdatetompc
	Cnxmobpp3pounds  pq.NullTime     `json:"cnxmobpp3pounds"`  // cnxmobpp3pounds
	Cnxmobdatefrmmpc pq.NullTime     `json:"cnxmobdatefrmmpc"` // cnxmobdatefrmmpc
	Cnxmobweborder   sql.NullInt64   `json:"cnxmobweborder"`   // cnxmobweborder
	Cnxmobpackageno  sql.NullString  `json:"cnxmobpackageno"`  // cnxmobpackageno
	Cnxmobdatetobill pq.NullTime     `json:"cnxmobdatetobill"` // cnxmobdatetobill
	Cnxmobdateletter pq.NullTime     `json:"cnxmobdateletter"` // cnxmobdateletter
	Cnxmobdateprepay pq.NullTime     `json:"cnxmobdateprepay"` // cnxmobdateprepay
	Cnxmobfreeaccess sql.NullInt64   `json:"cnxmobfreeaccess"` // cnxmobfreeaccess
	Cnxmobenteredby  sql.NullString  `json:"cnxmobenteredby"`  // cnxmobenteredby
	Cnxmobpromocd    sql.NullString  `json:"cnxmobpromocd"`    // cnxmobpromocd
	Cnxmobprice      sql.NullFloat64 `json:"cnxmobprice"`      // cnxmobprice
	Cnxmobprepaid    pq.NullTime     `json:"cnxmobprepaid"`    // cnxmobprepaid
	Cnxmobprepayment sql.NullFloat64 `json:"cnxmobprepayment"` // cnxmobprepayment
	Cnxmobloyaltylvl sql.NullString  `json:"cnxmobloyaltylvl"` // cnxmobloyaltylvl
	Cnxmobloyaltynxt pq.NullTime     `json:"cnxmobloyaltynxt"` // cnxmobloyaltynxt
	Cnxmobport       sql.NullInt64   `json:"cnxmobport"`       // cnxmobport
	Cnxmobcallquota  sql.NullFloat64 `json:"cnxmobcallquota"`  // cnxmobcallquota
	Cnxmobnettarrif  sql.NullString  `json:"cnxmobnettarrif"`  // cnxmobnettarrif
	Cnxmoboptrpc     sql.NullInt64   `json:"cnxmoboptrpc"`     // cnxmoboptrpc
	Cnxmoboptsms     sql.NullInt64   `json:"cnxmoboptsms"`     // cnxmoboptsms
	Cnxmoboptmms     sql.NullInt64   `json:"cnxmoboptmms"`     // cnxmoboptmms
	Cnxmoboptgprs    sql.NullInt64   `json:"cnxmoboptgprs"`    // cnxmoboptgprs
	Cnxmoboptlocal   sql.NullInt64   `json:"cnxmoboptlocal"`   // cnxmoboptlocal
	Cnxmoboptintern  sql.NullInt64   `json:"cnxmoboptintern"`  // cnxmoboptintern
	Cnxmoboptspare1  sql.NullInt64   `json:"cnxmoboptspare1"`  // cnxmoboptspare1
	Cnxmoboptspare2  sql.NullInt64   `json:"cnxmoboptspare2"`  // cnxmoboptspare2
	Cnxmobdiscount   sql.NullString  `json:"cnxmobdiscount"`   // cnxmobdiscount
	Cnxmobdeladdr    sql.NullInt64   `json:"cnxmobdeladdr"`    // cnxmobdeladdr
	Cnxmobdeladdref  sql.NullInt64   `json:"cnxmobdeladdref"`  // cnxmobdeladdref
	Cnxmobsubtariff  sql.NullString  `json:"cnxmobsubtariff"`  // cnxmobsubtariff
	Cnxmobservicelvl sql.NullString  `json:"cnxmobservicelvl"` // cnxmobservicelvl
	Cnxmobrpc        sql.NullString  `json:"cnxmobrpc"`        // cnxmobrpc
	Cnxmobsms        sql.NullString  `json:"cnxmobsms"`        // cnxmobsms
	Cnxmobmms        sql.NullString  `json:"cnxmobmms"`        // cnxmobmms
	Cnxmobgprs       sql.NullString  `json:"cnxmobgprs"`       // cnxmobgprs
	Cnxmoblocal      sql.NullString  `json:"cnxmoblocal"`      // cnxmoblocal
	Cnxmobinternat   sql.NullString  `json:"cnxmobinternat"`   // cnxmobinternat
	Cnxmoblinerental sql.NullString  `json:"cnxmoblinerental"` // cnxmoblinerental
	Cnxmobcgbid      sql.NullInt64   `json:"cnxmobcgbid"`      // cnxmobcgbid
	Cnxmobminterm    sql.NullInt64   `json:"cnxmobminterm"`    // cnxmobminterm
	Cnxmobunsubprice sql.NullFloat64 `json:"cnxmobunsubprice"` // cnxmobunsubprice
	Cnxmobordtype    sql.NullString  `json:"cnxmobordtype"`    // cnxmobordtype
	Cnxmobdatabundle sql.NullString  `json:"cnxmobdatabundle"` // cnxmobdatabundle
	Cnxmobbolton     sql.NullString  `json:"cnxmobbolton"`     // cnxmobbolton
	Cnxmobsparen1    sql.NullFloat64 `json:"cnxmobsparen1"`    // cnxmobsparen1
	Cnxmobsparen2    sql.NullFloat64 `json:"cnxmobsparen2"`    // cnxmobsparen2
	Cnxmobsparen3    sql.NullFloat64 `json:"cnxmobsparen3"`    // cnxmobsparen3
	Cnxmobportdate   pq.NullTime     `json:"cnxmobportdate"`   // cnxmobportdate
	Cnxmobcsigndate  pq.NullTime     `json:"cnxmobcsigndate"`  // cnxmobcsigndate
	Cnxmobspared3    pq.NullTime     `json:"cnxmobspared3"`    // cnxmobspared3
	Cnxmobtmlsalerep sql.NullInt64   `json:"cnxmobtmlsalerep"` // cnxmobtmlsalerep
	Cnxmobspecialoff sql.NullString  `json:"cnxmobspecialoff"` // cnxmobspecialoff
	Cnxmobdonoracc   sql.NullString  `json:"cnxmobdonoracc"`   // cnxmobdonoracc
	Cnxmobdoncliuni  sql.NullFloat64 `json:"cnxmobdoncliuni"`  // cnxmobdoncliuni
	Cnxmobholduntil  pq.NullTime     `json:"cnxmobholduntil"`  // cnxmobholduntil
	Cnxmobhmdate     pq.NullTime     `json:"cnxmobhmdate"`     // cnxmobhmdate
	Cnxmobiladdserv  sql.NullInt64   `json:"cnxmobiladdserv"`  // cnxmobiladdserv
	Cnxmobsmartphone sql.NullFloat64 `json:"cnxmobsmartphone"` // cnxmobsmartphone
	Cnxmobvaliddate  pq.NullTime     `json:"cnxmobvaliddate"`  // cnxmobvaliddate
	Cnxmobvalidby    sql.NullString  `json:"cnxmobvalidby"`    // cnxmobvalidby
	Cnxmobsold       sql.NullInt64   `json:"cnxmobsold"`       // cnxmobsold
	Cnxmobquicksim   sql.NullInt64   `json:"cnxmobquicksim"`   // cnxmobquicksim
	Cnxmobcampaign   sql.NullString  `json:"cnxmobcampaign"`   // cnxmobcampaign
	Cnxmobhscode     sql.NullString  `json:"cnxmobhscode"`     // cnxmobhscode
	Cnxmobpaccode    sql.NullString  `json:"cnxmobpaccode"`    // cnxmobpaccode
	Cnxmobpacexpiry  pq.NullTime     `json:"cnxmobpacexpiry"`  // cnxmobpacexpiry
	Cnxmobportingcli sql.NullString  `json:"cnxmobportingcli"` // cnxmobportingcli
	Cnxmobstatus     sql.NullString  `json:"cnxmobstatus"`     // cnxmobstatus
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cnxmobil exists in the database.
func (c *Cnxmobil) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cnxmobil has been deleted from the database.
func (c *Cnxmobil) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cnxmobil to the database.
func (c *Cnxmobil) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxmobil (` +
		`cnxmobcli, cnxmobsim, cnxmobiemi, cnxmobphonereqd, cnxmobcover, cnxmobvmail, cnxmobsecurity, cnxmobdateentrd, cnxmobdateorder, cnxmobdateivrreg, cnxmobdatetompc, cnxmobpp3pounds, cnxmobdatefrmmpc, cnxmobweborder, cnxmobpackageno, cnxmobdatetobill, cnxmobdateletter, cnxmobdateprepay, cnxmobfreeaccess, cnxmobenteredby, cnxmobpromocd, cnxmobprice, cnxmobprepaid, cnxmobprepayment, cnxmobloyaltylvl, cnxmobloyaltynxt, cnxmobport, cnxmobcallquota, cnxmobnettarrif, cnxmoboptrpc, cnxmoboptsms, cnxmoboptmms, cnxmoboptgprs, cnxmoboptlocal, cnxmoboptintern, cnxmoboptspare1, cnxmoboptspare2, cnxmobdiscount, cnxmobdeladdr, cnxmobdeladdref, cnxmobsubtariff, cnxmobservicelvl, cnxmobrpc, cnxmobsms, cnxmobmms, cnxmobgprs, cnxmoblocal, cnxmobinternat, cnxmoblinerental, cnxmobcgbid, cnxmobminterm, cnxmobunsubprice, cnxmobordtype, cnxmobdatabundle, cnxmobbolton, cnxmobsparen1, cnxmobsparen2, cnxmobsparen3, cnxmobportdate, cnxmobcsigndate, cnxmobspared3, cnxmobtmlsalerep, cnxmobspecialoff, cnxmobdonoracc, cnxmobdoncliuni, cnxmobholduntil, cnxmobhmdate, cnxmobiladdserv, cnxmobsmartphone, cnxmobvaliddate, cnxmobvalidby, cnxmobsold, cnxmobquicksim, cnxmobcampaign, cnxmobhscode, cnxmobpaccode, cnxmobpacexpiry, cnxmobportingcli, cnxmobstatus, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cnxmobil in the database.
func (c *Cnxmobil) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cnxmobil SET (` +
		`cnxmobcli, cnxmobsim, cnxmobiemi, cnxmobphonereqd, cnxmobcover, cnxmobvmail, cnxmobsecurity, cnxmobdateentrd, cnxmobdateorder, cnxmobdateivrreg, cnxmobdatetompc, cnxmobpp3pounds, cnxmobdatefrmmpc, cnxmobweborder, cnxmobpackageno, cnxmobdatetobill, cnxmobdateletter, cnxmobdateprepay, cnxmobfreeaccess, cnxmobenteredby, cnxmobpromocd, cnxmobprice, cnxmobprepaid, cnxmobprepayment, cnxmobloyaltylvl, cnxmobloyaltynxt, cnxmobport, cnxmobcallquota, cnxmobnettarrif, cnxmoboptrpc, cnxmoboptsms, cnxmoboptmms, cnxmoboptgprs, cnxmoboptlocal, cnxmoboptintern, cnxmoboptspare1, cnxmoboptspare2, cnxmobdiscount, cnxmobdeladdr, cnxmobdeladdref, cnxmobsubtariff, cnxmobservicelvl, cnxmobrpc, cnxmobsms, cnxmobmms, cnxmobgprs, cnxmoblocal, cnxmobinternat, cnxmoblinerental, cnxmobcgbid, cnxmobminterm, cnxmobunsubprice, cnxmobordtype, cnxmobdatabundle, cnxmobbolton, cnxmobsparen1, cnxmobsparen2, cnxmobsparen3, cnxmobportdate, cnxmobcsigndate, cnxmobspared3, cnxmobtmlsalerep, cnxmobspecialoff, cnxmobdonoracc, cnxmobdoncliuni, cnxmobholduntil, cnxmobhmdate, cnxmobiladdserv, cnxmobsmartphone, cnxmobvaliddate, cnxmobvalidby, cnxmobsold, cnxmobquicksim, cnxmobcampaign, cnxmobhscode, cnxmobpaccode, cnxmobpacexpiry, cnxmobportingcli, cnxmobstatus, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81` +
		`) WHERE equinox_lrn = $82`

	// run query
	XOLog(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cnxmobil to the database.
func (c *Cnxmobil) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cnxmobil.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cnxmobil) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxmobil (` +
		`cnxmobcli, cnxmobsim, cnxmobiemi, cnxmobphonereqd, cnxmobcover, cnxmobvmail, cnxmobsecurity, cnxmobdateentrd, cnxmobdateorder, cnxmobdateivrreg, cnxmobdatetompc, cnxmobpp3pounds, cnxmobdatefrmmpc, cnxmobweborder, cnxmobpackageno, cnxmobdatetobill, cnxmobdateletter, cnxmobdateprepay, cnxmobfreeaccess, cnxmobenteredby, cnxmobpromocd, cnxmobprice, cnxmobprepaid, cnxmobprepayment, cnxmobloyaltylvl, cnxmobloyaltynxt, cnxmobport, cnxmobcallquota, cnxmobnettarrif, cnxmoboptrpc, cnxmoboptsms, cnxmoboptmms, cnxmoboptgprs, cnxmoboptlocal, cnxmoboptintern, cnxmoboptspare1, cnxmoboptspare2, cnxmobdiscount, cnxmobdeladdr, cnxmobdeladdref, cnxmobsubtariff, cnxmobservicelvl, cnxmobrpc, cnxmobsms, cnxmobmms, cnxmobgprs, cnxmoblocal, cnxmobinternat, cnxmoblinerental, cnxmobcgbid, cnxmobminterm, cnxmobunsubprice, cnxmobordtype, cnxmobdatabundle, cnxmobbolton, cnxmobsparen1, cnxmobsparen2, cnxmobsparen3, cnxmobportdate, cnxmobcsigndate, cnxmobspared3, cnxmobtmlsalerep, cnxmobspecialoff, cnxmobdonoracc, cnxmobdoncliuni, cnxmobholduntil, cnxmobhmdate, cnxmobiladdserv, cnxmobsmartphone, cnxmobvaliddate, cnxmobvalidby, cnxmobsold, cnxmobquicksim, cnxmobcampaign, cnxmobhscode, cnxmobpaccode, cnxmobpacexpiry, cnxmobportingcli, cnxmobstatus, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cnxmobcli, cnxmobsim, cnxmobiemi, cnxmobphonereqd, cnxmobcover, cnxmobvmail, cnxmobsecurity, cnxmobdateentrd, cnxmobdateorder, cnxmobdateivrreg, cnxmobdatetompc, cnxmobpp3pounds, cnxmobdatefrmmpc, cnxmobweborder, cnxmobpackageno, cnxmobdatetobill, cnxmobdateletter, cnxmobdateprepay, cnxmobfreeaccess, cnxmobenteredby, cnxmobpromocd, cnxmobprice, cnxmobprepaid, cnxmobprepayment, cnxmobloyaltylvl, cnxmobloyaltynxt, cnxmobport, cnxmobcallquota, cnxmobnettarrif, cnxmoboptrpc, cnxmoboptsms, cnxmoboptmms, cnxmoboptgprs, cnxmoboptlocal, cnxmoboptintern, cnxmoboptspare1, cnxmoboptspare2, cnxmobdiscount, cnxmobdeladdr, cnxmobdeladdref, cnxmobsubtariff, cnxmobservicelvl, cnxmobrpc, cnxmobsms, cnxmobmms, cnxmobgprs, cnxmoblocal, cnxmobinternat, cnxmoblinerental, cnxmobcgbid, cnxmobminterm, cnxmobunsubprice, cnxmobordtype, cnxmobdatabundle, cnxmobbolton, cnxmobsparen1, cnxmobsparen2, cnxmobsparen3, cnxmobportdate, cnxmobcsigndate, cnxmobspared3, cnxmobtmlsalerep, cnxmobspecialoff, cnxmobdonoracc, cnxmobdoncliuni, cnxmobholduntil, cnxmobhmdate, cnxmobiladdserv, cnxmobsmartphone, cnxmobvaliddate, cnxmobvalidby, cnxmobsold, cnxmobquicksim, cnxmobcampaign, cnxmobhscode, cnxmobpaccode, cnxmobpacexpiry, cnxmobportingcli, cnxmobstatus, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cnxmobcli, EXCLUDED.cnxmobsim, EXCLUDED.cnxmobiemi, EXCLUDED.cnxmobphonereqd, EXCLUDED.cnxmobcover, EXCLUDED.cnxmobvmail, EXCLUDED.cnxmobsecurity, EXCLUDED.cnxmobdateentrd, EXCLUDED.cnxmobdateorder, EXCLUDED.cnxmobdateivrreg, EXCLUDED.cnxmobdatetompc, EXCLUDED.cnxmobpp3pounds, EXCLUDED.cnxmobdatefrmmpc, EXCLUDED.cnxmobweborder, EXCLUDED.cnxmobpackageno, EXCLUDED.cnxmobdatetobill, EXCLUDED.cnxmobdateletter, EXCLUDED.cnxmobdateprepay, EXCLUDED.cnxmobfreeaccess, EXCLUDED.cnxmobenteredby, EXCLUDED.cnxmobpromocd, EXCLUDED.cnxmobprice, EXCLUDED.cnxmobprepaid, EXCLUDED.cnxmobprepayment, EXCLUDED.cnxmobloyaltylvl, EXCLUDED.cnxmobloyaltynxt, EXCLUDED.cnxmobport, EXCLUDED.cnxmobcallquota, EXCLUDED.cnxmobnettarrif, EXCLUDED.cnxmoboptrpc, EXCLUDED.cnxmoboptsms, EXCLUDED.cnxmoboptmms, EXCLUDED.cnxmoboptgprs, EXCLUDED.cnxmoboptlocal, EXCLUDED.cnxmoboptintern, EXCLUDED.cnxmoboptspare1, EXCLUDED.cnxmoboptspare2, EXCLUDED.cnxmobdiscount, EXCLUDED.cnxmobdeladdr, EXCLUDED.cnxmobdeladdref, EXCLUDED.cnxmobsubtariff, EXCLUDED.cnxmobservicelvl, EXCLUDED.cnxmobrpc, EXCLUDED.cnxmobsms, EXCLUDED.cnxmobmms, EXCLUDED.cnxmobgprs, EXCLUDED.cnxmoblocal, EXCLUDED.cnxmobinternat, EXCLUDED.cnxmoblinerental, EXCLUDED.cnxmobcgbid, EXCLUDED.cnxmobminterm, EXCLUDED.cnxmobunsubprice, EXCLUDED.cnxmobordtype, EXCLUDED.cnxmobdatabundle, EXCLUDED.cnxmobbolton, EXCLUDED.cnxmobsparen1, EXCLUDED.cnxmobsparen2, EXCLUDED.cnxmobsparen3, EXCLUDED.cnxmobportdate, EXCLUDED.cnxmobcsigndate, EXCLUDED.cnxmobspared3, EXCLUDED.cnxmobtmlsalerep, EXCLUDED.cnxmobspecialoff, EXCLUDED.cnxmobdonoracc, EXCLUDED.cnxmobdoncliuni, EXCLUDED.cnxmobholduntil, EXCLUDED.cnxmobhmdate, EXCLUDED.cnxmobiladdserv, EXCLUDED.cnxmobsmartphone, EXCLUDED.cnxmobvaliddate, EXCLUDED.cnxmobvalidby, EXCLUDED.cnxmobsold, EXCLUDED.cnxmobquicksim, EXCLUDED.cnxmobcampaign, EXCLUDED.cnxmobhscode, EXCLUDED.cnxmobpaccode, EXCLUDED.cnxmobpacexpiry, EXCLUDED.cnxmobportingcli, EXCLUDED.cnxmobstatus, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cnxmobcli, c.Cnxmobsim, c.Cnxmobiemi, c.Cnxmobphonereqd, c.Cnxmobcover, c.Cnxmobvmail, c.Cnxmobsecurity, c.Cnxmobdateentrd, c.Cnxmobdateorder, c.Cnxmobdateivrreg, c.Cnxmobdatetompc, c.Cnxmobpp3pounds, c.Cnxmobdatefrmmpc, c.Cnxmobweborder, c.Cnxmobpackageno, c.Cnxmobdatetobill, c.Cnxmobdateletter, c.Cnxmobdateprepay, c.Cnxmobfreeaccess, c.Cnxmobenteredby, c.Cnxmobpromocd, c.Cnxmobprice, c.Cnxmobprepaid, c.Cnxmobprepayment, c.Cnxmobloyaltylvl, c.Cnxmobloyaltynxt, c.Cnxmobport, c.Cnxmobcallquota, c.Cnxmobnettarrif, c.Cnxmoboptrpc, c.Cnxmoboptsms, c.Cnxmoboptmms, c.Cnxmoboptgprs, c.Cnxmoboptlocal, c.Cnxmoboptintern, c.Cnxmoboptspare1, c.Cnxmoboptspare2, c.Cnxmobdiscount, c.Cnxmobdeladdr, c.Cnxmobdeladdref, c.Cnxmobsubtariff, c.Cnxmobservicelvl, c.Cnxmobrpc, c.Cnxmobsms, c.Cnxmobmms, c.Cnxmobgprs, c.Cnxmoblocal, c.Cnxmobinternat, c.Cnxmoblinerental, c.Cnxmobcgbid, c.Cnxmobminterm, c.Cnxmobunsubprice, c.Cnxmobordtype, c.Cnxmobdatabundle, c.Cnxmobbolton, c.Cnxmobsparen1, c.Cnxmobsparen2, c.Cnxmobsparen3, c.Cnxmobportdate, c.Cnxmobcsigndate, c.Cnxmobspared3, c.Cnxmobtmlsalerep, c.Cnxmobspecialoff, c.Cnxmobdonoracc, c.Cnxmobdoncliuni, c.Cnxmobholduntil, c.Cnxmobhmdate, c.Cnxmobiladdserv, c.Cnxmobsmartphone, c.Cnxmobvaliddate, c.Cnxmobvalidby, c.Cnxmobsold, c.Cnxmobquicksim, c.Cnxmobcampaign, c.Cnxmobhscode, c.Cnxmobpaccode, c.Cnxmobpacexpiry, c.Cnxmobportingcli, c.Cnxmobstatus, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cnxmobil from the database.
func (c *Cnxmobil) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cnxmobil WHERE equinox_lrn = $1`

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

// CnxmobilByEquinoxLrn retrieves a row from 'equinox.cnxmobil' as a Cnxmobil.
//
// Generated from index 'cnxmobil_pkey'.
func CnxmobilByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxmobil, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxmobcli, cnxmobsim, cnxmobiemi, cnxmobphonereqd, cnxmobcover, cnxmobvmail, cnxmobsecurity, cnxmobdateentrd, cnxmobdateorder, cnxmobdateivrreg, cnxmobdatetompc, cnxmobpp3pounds, cnxmobdatefrmmpc, cnxmobweborder, cnxmobpackageno, cnxmobdatetobill, cnxmobdateletter, cnxmobdateprepay, cnxmobfreeaccess, cnxmobenteredby, cnxmobpromocd, cnxmobprice, cnxmobprepaid, cnxmobprepayment, cnxmobloyaltylvl, cnxmobloyaltynxt, cnxmobport, cnxmobcallquota, cnxmobnettarrif, cnxmoboptrpc, cnxmoboptsms, cnxmoboptmms, cnxmoboptgprs, cnxmoboptlocal, cnxmoboptintern, cnxmoboptspare1, cnxmoboptspare2, cnxmobdiscount, cnxmobdeladdr, cnxmobdeladdref, cnxmobsubtariff, cnxmobservicelvl, cnxmobrpc, cnxmobsms, cnxmobmms, cnxmobgprs, cnxmoblocal, cnxmobinternat, cnxmoblinerental, cnxmobcgbid, cnxmobminterm, cnxmobunsubprice, cnxmobordtype, cnxmobdatabundle, cnxmobbolton, cnxmobsparen1, cnxmobsparen2, cnxmobsparen3, cnxmobportdate, cnxmobcsigndate, cnxmobspared3, cnxmobtmlsalerep, cnxmobspecialoff, cnxmobdonoracc, cnxmobdoncliuni, cnxmobholduntil, cnxmobhmdate, cnxmobiladdserv, cnxmobsmartphone, cnxmobvaliddate, cnxmobvalidby, cnxmobsold, cnxmobquicksim, cnxmobcampaign, cnxmobhscode, cnxmobpaccode, cnxmobpacexpiry, cnxmobportingcli, cnxmobstatus, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxmobil ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxmobil{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxmobcli, &c.Cnxmobsim, &c.Cnxmobiemi, &c.Cnxmobphonereqd, &c.Cnxmobcover, &c.Cnxmobvmail, &c.Cnxmobsecurity, &c.Cnxmobdateentrd, &c.Cnxmobdateorder, &c.Cnxmobdateivrreg, &c.Cnxmobdatetompc, &c.Cnxmobpp3pounds, &c.Cnxmobdatefrmmpc, &c.Cnxmobweborder, &c.Cnxmobpackageno, &c.Cnxmobdatetobill, &c.Cnxmobdateletter, &c.Cnxmobdateprepay, &c.Cnxmobfreeaccess, &c.Cnxmobenteredby, &c.Cnxmobpromocd, &c.Cnxmobprice, &c.Cnxmobprepaid, &c.Cnxmobprepayment, &c.Cnxmobloyaltylvl, &c.Cnxmobloyaltynxt, &c.Cnxmobport, &c.Cnxmobcallquota, &c.Cnxmobnettarrif, &c.Cnxmoboptrpc, &c.Cnxmoboptsms, &c.Cnxmoboptmms, &c.Cnxmoboptgprs, &c.Cnxmoboptlocal, &c.Cnxmoboptintern, &c.Cnxmoboptspare1, &c.Cnxmoboptspare2, &c.Cnxmobdiscount, &c.Cnxmobdeladdr, &c.Cnxmobdeladdref, &c.Cnxmobsubtariff, &c.Cnxmobservicelvl, &c.Cnxmobrpc, &c.Cnxmobsms, &c.Cnxmobmms, &c.Cnxmobgprs, &c.Cnxmoblocal, &c.Cnxmobinternat, &c.Cnxmoblinerental, &c.Cnxmobcgbid, &c.Cnxmobminterm, &c.Cnxmobunsubprice, &c.Cnxmobordtype, &c.Cnxmobdatabundle, &c.Cnxmobbolton, &c.Cnxmobsparen1, &c.Cnxmobsparen2, &c.Cnxmobsparen3, &c.Cnxmobportdate, &c.Cnxmobcsigndate, &c.Cnxmobspared3, &c.Cnxmobtmlsalerep, &c.Cnxmobspecialoff, &c.Cnxmobdonoracc, &c.Cnxmobdoncliuni, &c.Cnxmobholduntil, &c.Cnxmobhmdate, &c.Cnxmobiladdserv, &c.Cnxmobsmartphone, &c.Cnxmobvaliddate, &c.Cnxmobvalidby, &c.Cnxmobsold, &c.Cnxmobquicksim, &c.Cnxmobcampaign, &c.Cnxmobhscode, &c.Cnxmobpaccode, &c.Cnxmobpacexpiry, &c.Cnxmobportingcli, &c.Cnxmobstatus, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}