// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Express represents a row from 'equinox.express'.
type Express struct {
	ExprsID         sql.NullString  `json:"exprs_id"`         // exprs_id
	ExprsIssue      sql.NullInt64   `json:"exprs_issue"`      // exprs_issue
	ExprsStatus     sql.NullString  `json:"exprs_status"`     // exprs_status
	ExprsTicktssent pq.NullTime     `json:"exprs_ticktssent"` // exprs_ticktssent
	ExprsExecid     sql.NullString  `json:"exprs_execid"`     // exprs_execid
	ExprsOrderno    sql.NullString  `json:"exprs_orderno"`    // exprs_orderno
	ExprsForename   sql.NullString  `json:"exprs_forename"`   // exprs_forename
	ExprsSurname    sql.NullString  `json:"exprs_surname"`    // exprs_surname
	ExprsGuests     sql.NullInt64   `json:"exprs_guests"`     // exprs_guests
	ExprsEmail      sql.NullString  `json:"exprs_email"`      // exprs_email
	ExprsRegd       sql.NullInt64   `json:"exprs_regd"`       // exprs_regd
	ExprsRegtime    pq.NullTime     `json:"exprs_regtime"`    // exprs_regtime
	ExprsBadgeCode  sql.NullString  `json:"exprs_badge_code"` // exprs_badge_code
	ExprsMemo       sql.NullInt64   `json:"exprs_memo"`       // exprs_memo
	ExprsDate       pq.NullTime     `json:"exprs_date"`       // exprs_date
	ExprsPayMthd    sql.NullInt64   `json:"exprs_pay_mthd"`   // exprs_pay_mthd
	ExprsCcno       sql.NullString  `json:"exprs_ccno"`       // exprs_ccno
	ExprsExpdate    sql.NullString  `json:"exprs_expdate"`    // exprs_expdate
	ExprsStartdate  sql.NullString  `json:"exprs_startdate"`  // exprs_startdate
	ExprsIssueno    sql.NullString  `json:"exprs_issueno"`    // exprs_issueno
	ExprsChqno      sql.NullString  `json:"exprs_chqno"`      // exprs_chqno
	ExprsTotal      sql.NullFloat64 `json:"exprs_total"`      // exprs_total
	ExprsVat        sql.NullFloat64 `json:"exprs_vat"`        // exprs_vat
	ExprDisc        sql.NullFloat64 `json:"expr_disc"`        // expr_disc
	ExprTicketcode  sql.NullString  `json:"expr_ticketcode"`  // expr_ticketcode
	ExprNooftickets sql.NullInt64   `json:"expr_nooftickets"` // expr_nooftickets
	ExprVatrecpt    sql.NullInt64   `json:"expr_vatrecpt"`    // expr_vatrecpt
	ExprAdditional  sql.NullString  `json:"expr_additional"`  // expr_additional
	ExprGrpatickets sql.NullInt64   `json:"expr_grpatickets"` // expr_grpatickets
	ExprGrpbtickets sql.NullInt64   `json:"expr_grpbtickets"` // expr_grpbtickets
	ExprGrpctickets sql.NullInt64   `json:"expr_grpctickets"` // expr_grpctickets
	ExprGrpdtickets sql.NullInt64   `json:"expr_grpdtickets"` // expr_grpdtickets
	ExprSeminar     sql.NullString  `json:"expr_seminar"`     // expr_seminar
	ExprBlock       sql.NullString  `json:"expr_block"`       // expr_block
	ExprSparec1     sql.NullString  `json:"expr_sparec1"`     // expr_sparec1
	ExprSparec2     sql.NullString  `json:"expr_sparec2"`     // expr_sparec2
	ExprS1tickets   sql.NullInt64   `json:"expr_s1tickets"`   // expr_s1tickets
	ExprS2tickets   sql.NullInt64   `json:"expr_s2tickets"`   // expr_s2tickets
	ExprS3tickets   sql.NullInt64   `json:"expr_s3tickets"`   // expr_s3tickets
	ExprS4tickets   sql.NullInt64   `json:"expr_s4tickets"`   // expr_s4tickets
	ExprMainref     sql.NullString  `json:"expr_mainref"`     // expr_mainref
	ExprS5tickets   sql.NullInt64   `json:"expr_s5tickets"`   // expr_s5tickets
	ExprEventday    sql.NullString  `json:"expr_eventday"`    // expr_eventday
	ExprsDisplaynam sql.NullString  `json:"exprs_displaynam"` // exprs_displaynam
	EquinoxLrn      int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Express exists in the database.
func (e *Express) Exists() bool {
	return e._exists
}

// Deleted provides information if the Express has been deleted from the database.
func (e *Express) Deleted() bool {
	return e._deleted
}

// Insert inserts the Express to the database.
func (e *Express) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.express (` +
		`exprs_id, exprs_issue, exprs_status, exprs_ticktssent, exprs_execid, exprs_orderno, exprs_forename, exprs_surname, exprs_guests, exprs_email, exprs_regd, exprs_regtime, exprs_badge_code, exprs_memo, exprs_date, exprs_pay_mthd, exprs_ccno, exprs_expdate, exprs_startdate, exprs_issueno, exprs_chqno, exprs_total, exprs_vat, expr_disc, expr_ticketcode, expr_nooftickets, expr_vatrecpt, expr_additional, expr_grpatickets, expr_grpbtickets, expr_grpctickets, expr_grpdtickets, expr_seminar, expr_block, expr_sparec1, expr_sparec2, expr_s1tickets, expr_s2tickets, expr_s3tickets, expr_s4tickets, expr_mainref, expr_s5tickets, expr_eventday, exprs_displaynam, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Express in the database.
func (e *Express) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.express SET (` +
		`exprs_id, exprs_issue, exprs_status, exprs_ticktssent, exprs_execid, exprs_orderno, exprs_forename, exprs_surname, exprs_guests, exprs_email, exprs_regd, exprs_regtime, exprs_badge_code, exprs_memo, exprs_date, exprs_pay_mthd, exprs_ccno, exprs_expdate, exprs_startdate, exprs_issueno, exprs_chqno, exprs_total, exprs_vat, expr_disc, expr_ticketcode, expr_nooftickets, expr_vatrecpt, expr_additional, expr_grpatickets, expr_grpbtickets, expr_grpctickets, expr_grpdtickets, expr_seminar, expr_block, expr_sparec1, expr_sparec2, expr_s1tickets, expr_s2tickets, expr_s3tickets, expr_s4tickets, expr_mainref, expr_s5tickets, expr_eventday, exprs_displaynam, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45` +
		`) WHERE equinox_lrn = $46`

	// run query
	XOLog(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Express to the database.
func (e *Express) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Express.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Express) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.express (` +
		`exprs_id, exprs_issue, exprs_status, exprs_ticktssent, exprs_execid, exprs_orderno, exprs_forename, exprs_surname, exprs_guests, exprs_email, exprs_regd, exprs_regtime, exprs_badge_code, exprs_memo, exprs_date, exprs_pay_mthd, exprs_ccno, exprs_expdate, exprs_startdate, exprs_issueno, exprs_chqno, exprs_total, exprs_vat, expr_disc, expr_ticketcode, expr_nooftickets, expr_vatrecpt, expr_additional, expr_grpatickets, expr_grpbtickets, expr_grpctickets, expr_grpdtickets, expr_seminar, expr_block, expr_sparec1, expr_sparec2, expr_s1tickets, expr_s2tickets, expr_s3tickets, expr_s4tickets, expr_mainref, expr_s5tickets, expr_eventday, exprs_displaynam, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`exprs_id, exprs_issue, exprs_status, exprs_ticktssent, exprs_execid, exprs_orderno, exprs_forename, exprs_surname, exprs_guests, exprs_email, exprs_regd, exprs_regtime, exprs_badge_code, exprs_memo, exprs_date, exprs_pay_mthd, exprs_ccno, exprs_expdate, exprs_startdate, exprs_issueno, exprs_chqno, exprs_total, exprs_vat, expr_disc, expr_ticketcode, expr_nooftickets, expr_vatrecpt, expr_additional, expr_grpatickets, expr_grpbtickets, expr_grpctickets, expr_grpdtickets, expr_seminar, expr_block, expr_sparec1, expr_sparec2, expr_s1tickets, expr_s2tickets, expr_s3tickets, expr_s4tickets, expr_mainref, expr_s5tickets, expr_eventday, exprs_displaynam, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.exprs_id, EXCLUDED.exprs_issue, EXCLUDED.exprs_status, EXCLUDED.exprs_ticktssent, EXCLUDED.exprs_execid, EXCLUDED.exprs_orderno, EXCLUDED.exprs_forename, EXCLUDED.exprs_surname, EXCLUDED.exprs_guests, EXCLUDED.exprs_email, EXCLUDED.exprs_regd, EXCLUDED.exprs_regtime, EXCLUDED.exprs_badge_code, EXCLUDED.exprs_memo, EXCLUDED.exprs_date, EXCLUDED.exprs_pay_mthd, EXCLUDED.exprs_ccno, EXCLUDED.exprs_expdate, EXCLUDED.exprs_startdate, EXCLUDED.exprs_issueno, EXCLUDED.exprs_chqno, EXCLUDED.exprs_total, EXCLUDED.exprs_vat, EXCLUDED.expr_disc, EXCLUDED.expr_ticketcode, EXCLUDED.expr_nooftickets, EXCLUDED.expr_vatrecpt, EXCLUDED.expr_additional, EXCLUDED.expr_grpatickets, EXCLUDED.expr_grpbtickets, EXCLUDED.expr_grpctickets, EXCLUDED.expr_grpdtickets, EXCLUDED.expr_seminar, EXCLUDED.expr_block, EXCLUDED.expr_sparec1, EXCLUDED.expr_sparec2, EXCLUDED.expr_s1tickets, EXCLUDED.expr_s2tickets, EXCLUDED.expr_s3tickets, EXCLUDED.expr_s4tickets, EXCLUDED.expr_mainref, EXCLUDED.expr_s5tickets, EXCLUDED.expr_eventday, EXCLUDED.exprs_displaynam, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.ExprsID, e.ExprsIssue, e.ExprsStatus, e.ExprsTicktssent, e.ExprsExecid, e.ExprsOrderno, e.ExprsForename, e.ExprsSurname, e.ExprsGuests, e.ExprsEmail, e.ExprsRegd, e.ExprsRegtime, e.ExprsBadgeCode, e.ExprsMemo, e.ExprsDate, e.ExprsPayMthd, e.ExprsCcno, e.ExprsExpdate, e.ExprsStartdate, e.ExprsIssueno, e.ExprsChqno, e.ExprsTotal, e.ExprsVat, e.ExprDisc, e.ExprTicketcode, e.ExprNooftickets, e.ExprVatrecpt, e.ExprAdditional, e.ExprGrpatickets, e.ExprGrpbtickets, e.ExprGrpctickets, e.ExprGrpdtickets, e.ExprSeminar, e.ExprBlock, e.ExprSparec1, e.ExprSparec2, e.ExprS1tickets, e.ExprS2tickets, e.ExprS3tickets, e.ExprS4tickets, e.ExprMainref, e.ExprS5tickets, e.ExprEventday, e.ExprsDisplaynam, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Express from the database.
func (e *Express) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.express WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// ExpressByEquinoxLrn retrieves a row from 'equinox.express' as a Express.
//
// Generated from index 'express_pkey'.
func ExpressByEquinoxLrn(db XODB, equinoxLrn int64) (*Express, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`exprs_id, exprs_issue, exprs_status, exprs_ticktssent, exprs_execid, exprs_orderno, exprs_forename, exprs_surname, exprs_guests, exprs_email, exprs_regd, exprs_regtime, exprs_badge_code, exprs_memo, exprs_date, exprs_pay_mthd, exprs_ccno, exprs_expdate, exprs_startdate, exprs_issueno, exprs_chqno, exprs_total, exprs_vat, expr_disc, expr_ticketcode, expr_nooftickets, expr_vatrecpt, expr_additional, expr_grpatickets, expr_grpbtickets, expr_grpctickets, expr_grpdtickets, expr_seminar, expr_block, expr_sparec1, expr_sparec2, expr_s1tickets, expr_s2tickets, expr_s3tickets, expr_s4tickets, expr_mainref, expr_s5tickets, expr_eventday, exprs_displaynam, equinox_lrn, equinox_sec ` +
		`FROM equinox.express ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Express{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.ExprsID, &e.ExprsIssue, &e.ExprsStatus, &e.ExprsTicktssent, &e.ExprsExecid, &e.ExprsOrderno, &e.ExprsForename, &e.ExprsSurname, &e.ExprsGuests, &e.ExprsEmail, &e.ExprsRegd, &e.ExprsRegtime, &e.ExprsBadgeCode, &e.ExprsMemo, &e.ExprsDate, &e.ExprsPayMthd, &e.ExprsCcno, &e.ExprsExpdate, &e.ExprsStartdate, &e.ExprsIssueno, &e.ExprsChqno, &e.ExprsTotal, &e.ExprsVat, &e.ExprDisc, &e.ExprTicketcode, &e.ExprNooftickets, &e.ExprVatrecpt, &e.ExprAdditional, &e.ExprGrpatickets, &e.ExprGrpbtickets, &e.ExprGrpctickets, &e.ExprGrpdtickets, &e.ExprSeminar, &e.ExprBlock, &e.ExprSparec1, &e.ExprSparec2, &e.ExprS1tickets, &e.ExprS2tickets, &e.ExprS3tickets, &e.ExprS4tickets, &e.ExprMainref, &e.ExprS5tickets, &e.ExprEventday, &e.ExprsDisplaynam, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}