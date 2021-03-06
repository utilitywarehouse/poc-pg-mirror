// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Edsitem represents a row from 'equinox.edsitem'.
type Edsitem struct {
	Edsiinstance     sql.NullString  `json:"edsiinstance"`     // edsiinstance
	Edsibatchno      sql.NullString  `json:"edsibatchno"`      // edsibatchno
	Edsiimportbalnce sql.NullFloat64 `json:"edsiimportbalnce"` // edsiimportbalnce
	Edsiimported     pq.NullTime     `json:"edsiimported"`     // edsiimported
	Edsielecmpan     sql.NullString  `json:"edsielecmpan"`     // edsielecmpan
	Edsielecserialno sql.NullString  `json:"edsielecserialno"` // edsielecserialno
	Edsigasmpr       sql.NullString  `json:"edsigasmpr"`       // edsigasmpr
	Edsigasserialno  sql.NullString  `json:"edsigasserialno"`  // edsigasserialno
	Edsisent         pq.NullTime     `json:"edsisent"`         // edsisent
	Edsireturned     pq.NullTime     `json:"edsireturned"`     // edsireturned
	Edsipdvoutcome   sql.NullString  `json:"edsipdvoutcome"`   // edsipdvoutcome
	Edsisentwarrant  sql.NullInt64   `json:"edsisentwarrant"`  // edsisentwarrant
	Edsiagentcomment sql.NullInt64   `json:"edsiagentcomment"` // edsiagentcomment
	Edsippmrequested sql.NullString  `json:"edsippmrequested"` // edsippmrequested
	Edsiappointment  pq.NullTime     `json:"edsiappointment"`  // edsiappointment
	Edsiappointment2 pq.NullTime     `json:"edsiappointment2"` // edsiappointment2
	Edsippminstalled sql.NullString  `json:"edsippminstalled"` // edsippminstalled
	Edsidebtloadppme sql.NullFloat64 `json:"edsidebtloadppme"` // edsidebtloadppme
	Edsidebtloadppmg sql.NullFloat64 `json:"edsidebtloadppmg"` // edsidebtloadppmg
	Edsichangetenant sql.NullInt64   `json:"edsichangetenant"` // edsichangetenant
	Edsiproofrecd    sql.NullString  `json:"edsiproofrecd"`    // edsiproofrecd
	Edsitenantcomp   sql.NullInt64   `json:"edsitenantcomp"`   // edsitenantcomp
	Edsipdvreqconf   sql.NullInt64   `json:"edsipdvreqconf"`   // edsipdvreqconf
	Edsipdvcharge    sql.NullFloat64 `json:"edsipdvcharge"`    // edsipdvcharge
	Edsicourtdate    pq.NullTime     `json:"edsicourtdate"`    // edsicourtdate
	Edsiisoldate     pq.NullTime     `json:"edsiisoldate"`     // edsiisoldate
	Edsireplan       sql.NullString  `json:"edsireplan"`       // edsireplan
	Edsireplanreason sql.NullString  `json:"edsireplanreason"` // edsireplanreason
	Edsicourtdate2   pq.NullTime     `json:"edsicourtdate2"`   // edsicourtdate2
	Edsiisoldate2    pq.NullTime     `json:"edsiisoldate2"`    // edsiisoldate2
	Edsineedsdiscon  sql.NullInt64   `json:"edsineedsdiscon"`  // edsineedsdiscon
	Edsidatewarrant  pq.NullTime     `json:"edsidatewarrant"`  // edsidatewarrant
	Edsiagentname    sql.NullString  `json:"edsiagentname"`    // edsiagentname
	Edsiagentcontact sql.NullString  `json:"edsiagentcontact"` // edsiagentcontact
	Edsinewmsne      sql.NullFloat64 `json:"edsinewmsne"`      // edsinewmsne
	Edsinewmsng      sql.NullFloat64 `json:"edsinewmsng"`      // edsinewmsng
	Edsiclosingreadg sql.NullInt64   `json:"edsiclosingreadg"` // edsiclosingreadg
	Edsiclosingreade sql.NullInt64   `json:"edsiclosingreade"` // edsiclosingreade
	Edsidisconnected sql.NullString  `json:"edsidisconnected"` // edsidisconnected
	Edsifinaldebt    sql.NullFloat64 `json:"edsifinaldebt"`    // edsifinaldebt
	Edsidisconreadg  sql.NullInt64   `json:"edsidisconreadg"`  // edsidisconreadg
	Edsidisconreade  sql.NullInt64   `json:"edsidisconreade"`  // edsidisconreade
	Edsiwarrantbatch sql.NullString  `json:"edsiwarrantbatch"` // edsiwarrantbatch
	Edsilegalaction  sql.NullInt64   `json:"edsilegalaction"`  // edsilegalaction
	Edsippmwarrant   sql.NullString  `json:"edsippmwarrant"`   // edsippmwarrant
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllEdsitem(db XODB, callback func(x Edsitem) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`edsiinstance, edsibatchno, edsiimportbalnce, edsiimported, edsielecmpan, edsielecserialno, edsigasmpr, edsigasserialno, edsisent, edsireturned, edsipdvoutcome, edsisentwarrant, edsiagentcomment, edsippmrequested, edsiappointment, edsiappointment2, edsippminstalled, edsidebtloadppme, edsidebtloadppmg, edsichangetenant, edsiproofrecd, edsitenantcomp, edsipdvreqconf, edsipdvcharge, edsicourtdate, edsiisoldate, edsireplan, edsireplanreason, edsicourtdate2, edsiisoldate2, edsineedsdiscon, edsidatewarrant, edsiagentname, edsiagentcontact, edsinewmsne, edsinewmsng, edsiclosingreadg, edsiclosingreade, edsidisconnected, edsifinaldebt, edsidisconreadg, edsidisconreade, edsiwarrantbatch, edsilegalaction, edsippmwarrant, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.edsitem `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Edsitem{}

		// scan
		err = q.Scan(&e.Edsiinstance, &e.Edsibatchno, &e.Edsiimportbalnce, &e.Edsiimported, &e.Edsielecmpan, &e.Edsielecserialno, &e.Edsigasmpr, &e.Edsigasserialno, &e.Edsisent, &e.Edsireturned, &e.Edsipdvoutcome, &e.Edsisentwarrant, &e.Edsiagentcomment, &e.Edsippmrequested, &e.Edsiappointment, &e.Edsiappointment2, &e.Edsippminstalled, &e.Edsidebtloadppme, &e.Edsidebtloadppmg, &e.Edsichangetenant, &e.Edsiproofrecd, &e.Edsitenantcomp, &e.Edsipdvreqconf, &e.Edsipdvcharge, &e.Edsicourtdate, &e.Edsiisoldate, &e.Edsireplan, &e.Edsireplanreason, &e.Edsicourtdate2, &e.Edsiisoldate2, &e.Edsineedsdiscon, &e.Edsidatewarrant, &e.Edsiagentname, &e.Edsiagentcontact, &e.Edsinewmsne, &e.Edsinewmsng, &e.Edsiclosingreadg, &e.Edsiclosingreade, &e.Edsidisconnected, &e.Edsifinaldebt, &e.Edsidisconreadg, &e.Edsidisconreade, &e.Edsiwarrantbatch, &e.Edsilegalaction, &e.Edsippmwarrant, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// EdsitemByEquinoxLrn retrieves a row from 'equinox.edsitem' as a Edsitem.
//
// Generated from index 'edsitem_pkey'.
func EdsitemByEquinoxLrn(db XODB, equinoxLrn int64) (*Edsitem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`edsiinstance, edsibatchno, edsiimportbalnce, edsiimported, edsielecmpan, edsielecserialno, edsigasmpr, edsigasserialno, edsisent, edsireturned, edsipdvoutcome, edsisentwarrant, edsiagentcomment, edsippmrequested, edsiappointment, edsiappointment2, edsippminstalled, edsidebtloadppme, edsidebtloadppmg, edsichangetenant, edsiproofrecd, edsitenantcomp, edsipdvreqconf, edsipdvcharge, edsicourtdate, edsiisoldate, edsireplan, edsireplanreason, edsicourtdate2, edsiisoldate2, edsineedsdiscon, edsidatewarrant, edsiagentname, edsiagentcontact, edsinewmsne, edsinewmsng, edsiclosingreadg, edsiclosingreade, edsidisconnected, edsifinaldebt, edsidisconreadg, edsidisconreade, edsiwarrantbatch, edsilegalaction, edsippmwarrant, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.edsitem ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Edsitem{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Edsiinstance, &e.Edsibatchno, &e.Edsiimportbalnce, &e.Edsiimported, &e.Edsielecmpan, &e.Edsielecserialno, &e.Edsigasmpr, &e.Edsigasserialno, &e.Edsisent, &e.Edsireturned, &e.Edsipdvoutcome, &e.Edsisentwarrant, &e.Edsiagentcomment, &e.Edsippmrequested, &e.Edsiappointment, &e.Edsiappointment2, &e.Edsippminstalled, &e.Edsidebtloadppme, &e.Edsidebtloadppmg, &e.Edsichangetenant, &e.Edsiproofrecd, &e.Edsitenantcomp, &e.Edsipdvreqconf, &e.Edsipdvcharge, &e.Edsicourtdate, &e.Edsiisoldate, &e.Edsireplan, &e.Edsireplanreason, &e.Edsicourtdate2, &e.Edsiisoldate2, &e.Edsineedsdiscon, &e.Edsidatewarrant, &e.Edsiagentname, &e.Edsiagentcontact, &e.Edsinewmsne, &e.Edsinewmsng, &e.Edsiclosingreadg, &e.Edsiclosingreade, &e.Edsidisconnected, &e.Edsifinaldebt, &e.Edsidisconreadg, &e.Edsidisconreade, &e.Edsiwarrantbatch, &e.Edsilegalaction, &e.Edsippmwarrant, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
