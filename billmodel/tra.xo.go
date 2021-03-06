// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Tra represents a row from 'equinox.tras'.
type Tra struct {
	TrCustaccountno sql.NullString  `json:"tr_custaccountno"` // tr_custaccountno
	TrClinumber     sql.NullString  `json:"tr_clinumber"`     // tr_clinumber
	TrImportdate    pq.NullTime     `json:"tr_importdate"`    // tr_importdate
	TrFuelType      sql.NullString  `json:"tr_fuel_type"`     // tr_fuel_type
	TrSupplierRef   sql.NullString  `json:"tr_supplier_ref"`  // tr_supplier_ref
	TrAccountName   sql.NullString  `json:"tr_account_name"`  // tr_account_name
	TrSupplyAdd1    sql.NullString  `json:"tr_supply_add1"`   // tr_supply_add1
	TrSupplyAdd2    sql.NullString  `json:"tr_supply_add2"`   // tr_supply_add2
	TrSupplyAdd3    sql.NullString  `json:"tr_supply_add3"`   // tr_supply_add3
	TrSupplyAdd4    sql.NullString  `json:"tr_supply_add4"`   // tr_supply_add4
	TrSupplyAdd5    sql.NullString  `json:"tr_supply_add5"`   // tr_supply_add5
	TrSupplyPcode   sql.NullString  `json:"tr_supply_pcode"`  // tr_supply_pcode
	TrPaymentMeth   sql.NullString  `json:"tr_payment_meth"`  // tr_payment_meth
	TrArrearsFlag   sql.NullString  `json:"tr_arrears_flag"`  // tr_arrears_flag
	TrLongTermVac   sql.NullString  `json:"tr_long_term_vac"` // tr_long_term_vac
	TrMpanMprn      sql.NullString  `json:"tr_mpan_mprn"`     // tr_mpan_mprn
	TrMeterserialno sql.NullString  `json:"tr_meterserialno"` // tr_meterserialno
	TrMeterType     sql.NullString  `json:"tr_meter_type"`    // tr_meter_type
	TrMeterInstall  pq.NullTime     `json:"tr_meter_install"` // tr_meter_install
	TrMeterInspect  pq.NullTime     `json:"tr_meter_inspect"` // tr_meter_inspect
	TrMeterlocation sql.NullString  `json:"tr_meterlocation"` // tr_meterlocation
	TrMeterStatus   sql.NullString  `json:"tr_meter_status"`  // tr_meter_status
	TrAnnualconsD   pq.NullTime     `json:"tr_annualcons_d"`  // tr_annualcons_d
	TrAnnualconsN   sql.NullFloat64 `json:"tr_annualcons_n"`  // tr_annualcons_n
	TrSc3Score      sql.NullString  `json:"tr_sc3_score"`     // tr_sc3_score
	TrSc3Index      sql.NullString  `json:"tr_sc3_index"`     // tr_sc3_index
	TrReasonTab01   sql.NullString  `json:"tr_reason_tab01"`  // tr_reason_tab01
	TrReasonTab02   sql.NullString  `json:"tr_reason_tab02"`  // tr_reason_tab02
	TrReasonTab03   sql.NullString  `json:"tr_reason_tab03"`  // tr_reason_tab03
	TrReasonTab04   sql.NullString  `json:"tr_reason_tab04"`  // tr_reason_tab04
	TrReasonTab05   sql.NullString  `json:"tr_reason_tab05"`  // tr_reason_tab05
	TrReasonTab06   sql.NullString  `json:"tr_reason_tab06"`  // tr_reason_tab06
	TrReasonTab07   sql.NullString  `json:"tr_reason_tab07"`  // tr_reason_tab07
	TrReasonTab08   sql.NullString  `json:"tr_reason_tab08"`  // tr_reason_tab08
	TrReasonTab09   sql.NullString  `json:"tr_reason_tab09"`  // tr_reason_tab09
	TrReasonTab10   sql.NullString  `json:"tr_reason_tab10"`  // tr_reason_tab10
	TrReasonTab11   sql.NullString  `json:"tr_reason_tab11"`  // tr_reason_tab11
	TrReasonTab12   sql.NullString  `json:"tr_reason_tab12"`  // tr_reason_tab12
	TrReasonTab13   sql.NullString  `json:"tr_reason_tab13"`  // tr_reason_tab13
	TrReasonTab14   sql.NullString  `json:"tr_reason_tab14"`  // tr_reason_tab14
	TrReasonTab15   sql.NullString  `json:"tr_reason_tab15"`  // tr_reason_tab15
	TrReasonTab16   sql.NullString  `json:"tr_reason_tab16"`  // tr_reason_tab16
	TrReasonTab17   sql.NullString  `json:"tr_reason_tab17"`  // tr_reason_tab17
	TrReasonTab18   sql.NullString  `json:"tr_reason_tab18"`  // tr_reason_tab18
	TrReasonTab19   sql.NullString  `json:"tr_reason_tab19"`  // tr_reason_tab19
	TrReasonTab20   sql.NullString  `json:"tr_reason_tab20"`  // tr_reason_tab20
	TrReasonTab21   sql.NullString  `json:"tr_reason_tab21"`  // tr_reason_tab21
	TrReasonTab22   sql.NullString  `json:"tr_reason_tab22"`  // tr_reason_tab22
	TrReasonTab23   sql.NullString  `json:"tr_reason_tab23"`  // tr_reason_tab23
	TrReasonTab24   sql.NullString  `json:"tr_reason_tab24"`  // tr_reason_tab24
	TrReasonTab25   sql.NullString  `json:"tr_reason_tab25"`  // tr_reason_tab25
	TrReasonTab26   sql.NullString  `json:"tr_reason_tab26"`  // tr_reason_tab26
	TrReasonTab27   sql.NullString  `json:"tr_reason_tab27"`  // tr_reason_tab27
	TrReasonTab28   sql.NullString  `json:"tr_reason_tab28"`  // tr_reason_tab28
	TrReasonTab29   sql.NullString  `json:"tr_reason_tab29"`  // tr_reason_tab29
	TrReasonTab30   sql.NullString  `json:"tr_reason_tab30"`  // tr_reason_tab30
	TrSupInvestID   sql.NullString  `json:"tr_sup_invest_id"` // tr_sup_invest_id
	TrLeadSource    sql.NullString  `json:"tr_lead_source"`   // tr_lead_source
	TrCurinvestcode sql.NullString  `json:"tr_curinvestcode"` // tr_curinvestcode
	TrTampRepD      pq.NullTime     `json:"tr_tamp_rep_d"`    // tr_tamp_rep_d
	TrInvestStart   pq.NullTime     `json:"tr_invest_start"`  // tr_invest_start
	TrVisit1        pq.NullTime     `json:"tr_visit_1"`       // tr_visit_1
	TrVisit2        pq.NullTime     `json:"tr_visit_2"`       // tr_visit_2
	TrVisit3        pq.NullTime     `json:"tr_visit_3"`       // tr_visit_3
	TrCourtDate     pq.NullTime     `json:"tr_court_date"`    // tr_court_date
	TrWarrantD      pq.NullTime     `json:"tr_warrant_d"`     // tr_warrant_d
	TrWarrantA      sql.NullInt64   `json:"tr_warrant_a"`     // tr_warrant_a
	TrWarrantR      sql.NullInt64   `json:"tr_warrant_r"`     // tr_warrant_r
	TrD0239Rec      pq.NullTime     `json:"tr_d0239_rec"`     // tr_d0239_rec
	TrTampCode      sql.NullString  `json:"tr_tamp_code"`     // tr_tamp_code
	TrTampSource    sql.NullString  `json:"tr_tamp_source"`   // tr_tamp_source
	TrTheftType     sql.NullString  `json:"tr_theft_type"`    // tr_theft_type
	TrOutcome       sql.NullString  `json:"tr_outcome"`       // tr_outcome
	TrVulnerable    sql.NullString  `json:"tr_vulnerable"`    // tr_vulnerable
	TrCrimeRefno    sql.NullString  `json:"tr_crime_refno"`   // tr_crime_refno
	TrSecDevsFit    sql.NullString  `json:"tr_sec_devs_fit"`  // tr_sec_devs_fit
	TrDateClosed    pq.NullTime     `json:"tr_date_closed"`   // tr_date_closed
	TrDateRepIn     pq.NullTime     `json:"tr_date_rep_in"`   // tr_date_rep_in
	TrContRefno     sql.NullString  `json:"tr_cont_refno"`    // tr_cont_refno
	TrTheftTypecms  sql.NullString  `json:"tr_theft_typecms"` // tr_theft_typecms
	TrTogInvestout  sql.NullString  `json:"tr_tog_investout"` // tr_tog_investout
	TrTotalCosts    sql.NullFloat64 `json:"tr_total_costs"`   // tr_total_costs
	TrCustCharged   sql.NullString  `json:"tr_cust_charged"`  // tr_cust_charged
	TrChargeExvat   sql.NullFloat64 `json:"tr_charge_exvat"`  // tr_charge_exvat
	TrChargeVat     sql.NullFloat64 `json:"tr_charge_vat"`    // tr_charge_vat
	TrChargeIncvat  sql.NullFloat64 `json:"tr_charge_incvat"` // tr_charge_incvat
	TrRevRecovered  sql.NullFloat64 `json:"tr_rev_recovered"` // tr_rev_recovered
	TrTheftStarted  pq.NullTime     `json:"tr_theft_started"` // tr_theft_started
	TrTheftEnded    pq.NullTime     `json:"tr_theft_ended"`   // tr_theft_ended
	TrAssessedLoss  sql.NullFloat64 `json:"tr_assessed_loss"` // tr_assessed_loss
	TrDayUnReckwh   sql.NullInt64   `json:"tr_day_un_reckwh"` // tr_day_un_reckwh
	TrNightUreckwh  sql.NullInt64   `json:"tr_night_ureckwh"` // tr_night_ureckwh
	TrDayRate       sql.NullFloat64 `json:"tr_day_rate"`      // tr_day_rate
	TrNightRate     sql.NullFloat64 `json:"tr_night_rate"`    // tr_night_rate
	TrSentToDc      pq.NullTime     `json:"tr_sent_to_dc"`    // tr_sent_to_dc
	TrAaUpdated     pq.NullTime     `json:"tr_aa_updated"`    // tr_aa_updated
	TrConsumpExvat  sql.NullFloat64 `json:"tr_consump_exvat"` // tr_consump_exvat
	TrConsumpnrgvat sql.NullFloat64 `json:"tr_consumpnrgvat"` // tr_consumpnrgvat
	TrConsumpstdvat sql.NullFloat64 `json:"tr_consumpstdvat"` // tr_consumpstdvat
	TrConsumpTotal  sql.NullFloat64 `json:"tr_consump_total"` // tr_consump_total
	TrUnitrate      sql.NullFloat64 `json:"tr_unitrate"`      // tr_unitrate
	TrLastchangeD   pq.NullTime     `json:"tr_lastchange_d"`  // tr_lastchange_d
	TrLastchangeT   pq.NullTime     `json:"tr_lastchange_t"`  // tr_lastchange_t
	TrLastchangeU   sql.NullString  `json:"tr_lastchange_u"`  // tr_lastchange_u
	TrSentBack      pq.NullTime     `json:"tr_sent_back"`     // tr_sent_back
	TrStatuschanged pq.NullTime     `json:"tr_statuschanged"` // tr_statuschanged
	TrSparec1       sql.NullString  `json:"tr_sparec1"`       // tr_sparec1
	TrSparen1       sql.NullInt64   `json:"tr_sparen1"`       // tr_sparen1
	TrSparem1       sql.NullInt64   `json:"tr_sparem1"`       // tr_sparem1
	TrSpared1       pq.NullTime     `json:"tr_spared1"`       // tr_spared1
	EquinoxLrn      int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllTra(db XODB, callback func(x Tra) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`tr_custaccountno, tr_clinumber, tr_importdate, tr_fuel_type, tr_supplier_ref, tr_account_name, tr_supply_add1, tr_supply_add2, tr_supply_add3, tr_supply_add4, tr_supply_add5, tr_supply_pcode, tr_payment_meth, tr_arrears_flag, tr_long_term_vac, tr_mpan_mprn, tr_meterserialno, tr_meter_type, tr_meter_install, tr_meter_inspect, tr_meterlocation, tr_meter_status, tr_annualcons_d, tr_annualcons_n, tr_sc3_score, tr_sc3_index, tr_reason_tab01, tr_reason_tab02, tr_reason_tab03, tr_reason_tab04, tr_reason_tab05, tr_reason_tab06, tr_reason_tab07, tr_reason_tab08, tr_reason_tab09, tr_reason_tab10, tr_reason_tab11, tr_reason_tab12, tr_reason_tab13, tr_reason_tab14, tr_reason_tab15, tr_reason_tab16, tr_reason_tab17, tr_reason_tab18, tr_reason_tab19, tr_reason_tab20, tr_reason_tab21, tr_reason_tab22, tr_reason_tab23, tr_reason_tab24, tr_reason_tab25, tr_reason_tab26, tr_reason_tab27, tr_reason_tab28, tr_reason_tab29, tr_reason_tab30, tr_sup_invest_id, tr_lead_source, tr_curinvestcode, tr_tamp_rep_d, tr_invest_start, tr_visit_1, tr_visit_2, tr_visit_3, tr_court_date, tr_warrant_d, tr_warrant_a, tr_warrant_r, tr_d0239_rec, tr_tamp_code, tr_tamp_source, tr_theft_type, tr_outcome, tr_vulnerable, tr_crime_refno, tr_sec_devs_fit, tr_date_closed, tr_date_rep_in, tr_cont_refno, tr_theft_typecms, tr_tog_investout, tr_total_costs, tr_cust_charged, tr_charge_exvat, tr_charge_vat, tr_charge_incvat, tr_rev_recovered, tr_theft_started, tr_theft_ended, tr_assessed_loss, tr_day_un_reckwh, tr_night_ureckwh, tr_day_rate, tr_night_rate, tr_sent_to_dc, tr_aa_updated, tr_consump_exvat, tr_consumpnrgvat, tr_consumpstdvat, tr_consump_total, tr_unitrate, tr_lastchange_d, tr_lastchange_t, tr_lastchange_u, tr_sent_back, tr_statuschanged, tr_sparec1, tr_sparen1, tr_sparem1, tr_spared1, equinox_lrn, equinox_sec ` +
		`FROM equinox.tras `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		t := Tra{}

		// scan
		err = q.Scan(&t.TrCustaccountno, &t.TrClinumber, &t.TrImportdate, &t.TrFuelType, &t.TrSupplierRef, &t.TrAccountName, &t.TrSupplyAdd1, &t.TrSupplyAdd2, &t.TrSupplyAdd3, &t.TrSupplyAdd4, &t.TrSupplyAdd5, &t.TrSupplyPcode, &t.TrPaymentMeth, &t.TrArrearsFlag, &t.TrLongTermVac, &t.TrMpanMprn, &t.TrMeterserialno, &t.TrMeterType, &t.TrMeterInstall, &t.TrMeterInspect, &t.TrMeterlocation, &t.TrMeterStatus, &t.TrAnnualconsD, &t.TrAnnualconsN, &t.TrSc3Score, &t.TrSc3Index, &t.TrReasonTab01, &t.TrReasonTab02, &t.TrReasonTab03, &t.TrReasonTab04, &t.TrReasonTab05, &t.TrReasonTab06, &t.TrReasonTab07, &t.TrReasonTab08, &t.TrReasonTab09, &t.TrReasonTab10, &t.TrReasonTab11, &t.TrReasonTab12, &t.TrReasonTab13, &t.TrReasonTab14, &t.TrReasonTab15, &t.TrReasonTab16, &t.TrReasonTab17, &t.TrReasonTab18, &t.TrReasonTab19, &t.TrReasonTab20, &t.TrReasonTab21, &t.TrReasonTab22, &t.TrReasonTab23, &t.TrReasonTab24, &t.TrReasonTab25, &t.TrReasonTab26, &t.TrReasonTab27, &t.TrReasonTab28, &t.TrReasonTab29, &t.TrReasonTab30, &t.TrSupInvestID, &t.TrLeadSource, &t.TrCurinvestcode, &t.TrTampRepD, &t.TrInvestStart, &t.TrVisit1, &t.TrVisit2, &t.TrVisit3, &t.TrCourtDate, &t.TrWarrantD, &t.TrWarrantA, &t.TrWarrantR, &t.TrD0239Rec, &t.TrTampCode, &t.TrTampSource, &t.TrTheftType, &t.TrOutcome, &t.TrVulnerable, &t.TrCrimeRefno, &t.TrSecDevsFit, &t.TrDateClosed, &t.TrDateRepIn, &t.TrContRefno, &t.TrTheftTypecms, &t.TrTogInvestout, &t.TrTotalCosts, &t.TrCustCharged, &t.TrChargeExvat, &t.TrChargeVat, &t.TrChargeIncvat, &t.TrRevRecovered, &t.TrTheftStarted, &t.TrTheftEnded, &t.TrAssessedLoss, &t.TrDayUnReckwh, &t.TrNightUreckwh, &t.TrDayRate, &t.TrNightRate, &t.TrSentToDc, &t.TrAaUpdated, &t.TrConsumpExvat, &t.TrConsumpnrgvat, &t.TrConsumpstdvat, &t.TrConsumpTotal, &t.TrUnitrate, &t.TrLastchangeD, &t.TrLastchangeT, &t.TrLastchangeU, &t.TrSentBack, &t.TrStatuschanged, &t.TrSparec1, &t.TrSparen1, &t.TrSparem1, &t.TrSpared1, &t.EquinoxLrn, &t.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(t) {
			return nil
		}
	}

	return nil
}

// TraByEquinoxLrn retrieves a row from 'equinox.tras' as a Tra.
//
// Generated from index 'tras_pkey'.
func TraByEquinoxLrn(db XODB, equinoxLrn int64) (*Tra, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tr_custaccountno, tr_clinumber, tr_importdate, tr_fuel_type, tr_supplier_ref, tr_account_name, tr_supply_add1, tr_supply_add2, tr_supply_add3, tr_supply_add4, tr_supply_add5, tr_supply_pcode, tr_payment_meth, tr_arrears_flag, tr_long_term_vac, tr_mpan_mprn, tr_meterserialno, tr_meter_type, tr_meter_install, tr_meter_inspect, tr_meterlocation, tr_meter_status, tr_annualcons_d, tr_annualcons_n, tr_sc3_score, tr_sc3_index, tr_reason_tab01, tr_reason_tab02, tr_reason_tab03, tr_reason_tab04, tr_reason_tab05, tr_reason_tab06, tr_reason_tab07, tr_reason_tab08, tr_reason_tab09, tr_reason_tab10, tr_reason_tab11, tr_reason_tab12, tr_reason_tab13, tr_reason_tab14, tr_reason_tab15, tr_reason_tab16, tr_reason_tab17, tr_reason_tab18, tr_reason_tab19, tr_reason_tab20, tr_reason_tab21, tr_reason_tab22, tr_reason_tab23, tr_reason_tab24, tr_reason_tab25, tr_reason_tab26, tr_reason_tab27, tr_reason_tab28, tr_reason_tab29, tr_reason_tab30, tr_sup_invest_id, tr_lead_source, tr_curinvestcode, tr_tamp_rep_d, tr_invest_start, tr_visit_1, tr_visit_2, tr_visit_3, tr_court_date, tr_warrant_d, tr_warrant_a, tr_warrant_r, tr_d0239_rec, tr_tamp_code, tr_tamp_source, tr_theft_type, tr_outcome, tr_vulnerable, tr_crime_refno, tr_sec_devs_fit, tr_date_closed, tr_date_rep_in, tr_cont_refno, tr_theft_typecms, tr_tog_investout, tr_total_costs, tr_cust_charged, tr_charge_exvat, tr_charge_vat, tr_charge_incvat, tr_rev_recovered, tr_theft_started, tr_theft_ended, tr_assessed_loss, tr_day_un_reckwh, tr_night_ureckwh, tr_day_rate, tr_night_rate, tr_sent_to_dc, tr_aa_updated, tr_consump_exvat, tr_consumpnrgvat, tr_consumpstdvat, tr_consump_total, tr_unitrate, tr_lastchange_d, tr_lastchange_t, tr_lastchange_u, tr_sent_back, tr_statuschanged, tr_sparec1, tr_sparen1, tr_sparem1, tr_spared1, equinox_lrn, equinox_sec ` +
		`FROM equinox.tras ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Tra{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.TrCustaccountno, &t.TrClinumber, &t.TrImportdate, &t.TrFuelType, &t.TrSupplierRef, &t.TrAccountName, &t.TrSupplyAdd1, &t.TrSupplyAdd2, &t.TrSupplyAdd3, &t.TrSupplyAdd4, &t.TrSupplyAdd5, &t.TrSupplyPcode, &t.TrPaymentMeth, &t.TrArrearsFlag, &t.TrLongTermVac, &t.TrMpanMprn, &t.TrMeterserialno, &t.TrMeterType, &t.TrMeterInstall, &t.TrMeterInspect, &t.TrMeterlocation, &t.TrMeterStatus, &t.TrAnnualconsD, &t.TrAnnualconsN, &t.TrSc3Score, &t.TrSc3Index, &t.TrReasonTab01, &t.TrReasonTab02, &t.TrReasonTab03, &t.TrReasonTab04, &t.TrReasonTab05, &t.TrReasonTab06, &t.TrReasonTab07, &t.TrReasonTab08, &t.TrReasonTab09, &t.TrReasonTab10, &t.TrReasonTab11, &t.TrReasonTab12, &t.TrReasonTab13, &t.TrReasonTab14, &t.TrReasonTab15, &t.TrReasonTab16, &t.TrReasonTab17, &t.TrReasonTab18, &t.TrReasonTab19, &t.TrReasonTab20, &t.TrReasonTab21, &t.TrReasonTab22, &t.TrReasonTab23, &t.TrReasonTab24, &t.TrReasonTab25, &t.TrReasonTab26, &t.TrReasonTab27, &t.TrReasonTab28, &t.TrReasonTab29, &t.TrReasonTab30, &t.TrSupInvestID, &t.TrLeadSource, &t.TrCurinvestcode, &t.TrTampRepD, &t.TrInvestStart, &t.TrVisit1, &t.TrVisit2, &t.TrVisit3, &t.TrCourtDate, &t.TrWarrantD, &t.TrWarrantA, &t.TrWarrantR, &t.TrD0239Rec, &t.TrTampCode, &t.TrTampSource, &t.TrTheftType, &t.TrOutcome, &t.TrVulnerable, &t.TrCrimeRefno, &t.TrSecDevsFit, &t.TrDateClosed, &t.TrDateRepIn, &t.TrContRefno, &t.TrTheftTypecms, &t.TrTogInvestout, &t.TrTotalCosts, &t.TrCustCharged, &t.TrChargeExvat, &t.TrChargeVat, &t.TrChargeIncvat, &t.TrRevRecovered, &t.TrTheftStarted, &t.TrTheftEnded, &t.TrAssessedLoss, &t.TrDayUnReckwh, &t.TrNightUreckwh, &t.TrDayRate, &t.TrNightRate, &t.TrSentToDc, &t.TrAaUpdated, &t.TrConsumpExvat, &t.TrConsumpnrgvat, &t.TrConsumpstdvat, &t.TrConsumpTotal, &t.TrUnitrate, &t.TrLastchangeD, &t.TrLastchangeT, &t.TrLastchangeU, &t.TrSentBack, &t.TrStatuschanged, &t.TrSparec1, &t.TrSparen1, &t.TrSparem1, &t.TrSpared1, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
