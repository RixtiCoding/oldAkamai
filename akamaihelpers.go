package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Bmak struct {
	Ver               float64       `json:"ver"`
	KeCntLmt          int           `json:"ke_cnt_lmt"`
	MmeCntLmt         int           `json:"mme_cnt_lmt"`
	MduceCntLmt       int           `json:"mduce_cnt_lmt"`
	PmeCntLmt         int           `json:"pme_cnt_lmt"`
	PduceCntLmt       int           `json:"pduce_cnt_lmt"`
	TmeCntLmt         int           `json:"tme_cnt_lmt"`
	TduceCntLmt       int           `json:"tduce_cnt_lmt"`
	DoeCntLmt         int           `json:"doe_cnt_lmt"`
	DmeCntLmt         int           `json:"dme_cnt_lmt"`
	VcCntLmt          int           `json:"vc_cnt_lmt"`
	DoaThrottle       int           `json:"doa_throttle"`
	DmaThrottle       int           `json:"dma_throttle"`
	SessionID         string        `json:"session_id"`
	JsPost            bool          `json:"js_post"`
	Loc               string        `json:"loc"`
	CfURL             string        `json:"cf_url"`
	ParamsURL         string        `json:"params_url"`
	Auth              string        `json:"auth"`
	APIPublicKey      string        `json:"api_public_key"`
	AjLmtDoact        int           `json:"aj_lmt_doact"`
	AjLmtDmact        int           `json:"aj_lmt_dmact"`
	AjLmtTact         int           `json:"aj_lmt_tact"`
	CeJsPost          int           `json:"ce_js_post"`
	InitTime          int           `json:"init_time"`
	Informinfo        string        `json:"informinfo"`
	Prevfid           int           `json:"prevfid"`
	Fidcnt            int           `json:"fidcnt"`
	SensorData        string        `json:"sensor_data"`
	Ins               string        `json:"ins"`
	Cns               string        `json:"cns"`
	Engetloc          int           `json:"enGetLoc"`
	Enreaddocurl      int           `json:"enReadDocUrl"`
	Disfpcalontimeout int           `json:"disFpCalOnTimeout"`
	Xagg              int           `json:"xagg"`
	Pen               int           `json:"pen"`
	Brow              string        `json:"brow"`
	Browver           string        `json:"browver"`
	Psub              string        `json:"psub"`
	Lang              string        `json:"lang"`
	Prod              string        `json:"prod"`
	Plen              int           `json:"plen"`
	DoadmaEn          int           `json:"doadma_en"`
	Sdfn              []interface{} `json:"sdfn"`
	D2                int64         `json:"d2"`
	D3                int           `json:"d3"`
	Thr               int           `json:"thr"`
	Cs                string        `json:"cs"`
	Hn                string        `json:"hn"`
	Z1                int64         `json:"z1"`
	O9                int           `json:"o9"`
	Vc                string        `json:"vc"`
	Y1                int           `json:"y1"`
	Ta                int           `json:"ta"`
	Tst               int           `json:"tst"`
	TTst              int64         `json:"t_tst"`
	Ckie              string        `json:"ckie"`
	NCk               string        `json:"n_ck"`
	Ckurl             int           `json:"ckurl"`
	Bm                bool          `json:"bm"`
	Mr                string        `json:"mr"`
	Altfonts          bool          `json:"altFonts"`
	Rst               bool          `json:"rst"`
	Runfonts          bool          `json:"runFonts"`
	Fsp               bool          `json:"fsp"`
	Firstload         bool          `json:"firstLoad"`
	Pstate            bool          `json:"pstate"`
	MnMcLmt           int           `json:"mn_mc_lmt"`
	MnState           int           `json:"mn_state"`
	MnMcIndx          int           `json:"mn_mc_indx"`
	MnSen             int           `json:"mn_sen"`
	MnTout            int           `json:"mn_tout"`
	MnStout           int           `json:"mn_stout"`
	MnCt              int           `json:"mn_ct"`
	MnCc              string        `json:"mn_cc"`
	MnCd              int           `json:"mn_cd"`
	MnLc              []interface{} `json:"mn_lc"`
	MnLd              []interface{} `json:"mn_ld"`
	MnLcl             int           `json:"mn_lcl"`
	MnAl              []interface{} `json:"mn_al"`
	MnIl              []interface{} `json:"mn_il"`
	MnTcl             []interface{} `json:"mn_tcl"`
	MnR               []interface{} `json:"mn_r"`
	MnRt              int           `json:"mn_rt"`
	MnWt              int           `json:"mn_wt"`
	MnAbck            string        `json:"mn_abck"`
	MnPsn             string        `json:"mn_psn"`
	MnTs              string        `json:"mn_ts"`
	MnLg              []interface{} `json:"mn_lg"`
	Loap              int           `json:"loap"`
	Dcs               int           `json:"dcs"`
	Listfunctions     struct {
	} `json:"listFunctions"`
	StartTs     int64  `json:"start_ts"`
	Kact        string `json:"kact"`
	KeCnt       int    `json:"ke_cnt"`
	KeVel       int    `json:"ke_vel"`
	Mact        string `json:"mact"`
	MmeCnt      int    `json:"mme_cnt"`
	MduceCnt    int    `json:"mduce_cnt"`
	MeVel       int    `json:"me_vel"`
	Pact        string `json:"pact"`
	PmeCnt      int    `json:"pme_cnt"`
	PduceCnt    int    `json:"pduce_cnt"`
	PeVel       int    `json:"pe_vel"`
	Tact        string `json:"tact"`
	TmeCnt      int    `json:"tme_cnt"`
	TduceCnt    int    `json:"tduce_cnt"`
	TeVel       int    `json:"te_vel"`
	Doact       string `json:"doact"`
	DoeCnt      int    `json:"doe_cnt"`
	DoeVel      int    `json:"doe_vel"`
	Dmact       string `json:"dmact"`
	DmeCnt      int    `json:"dme_cnt"`
	DmeVel      int    `json:"dme_vel"`
	Vcact       string `json:"vcact"`
	VcCnt       int    `json:"vc_cnt"`
	AjIndx      int    `json:"aj_indx"`
	AjSs        int    `json:"aj_ss"`
	AjType      int    `json:"aj_type"`
	AjIndxDoact int    `json:"aj_indx_doact"`
	AjIndxDmact int    `json:"aj_indx_dmact"`
	AjIndxTact  int    `json:"aj_indx_tact"`
	MeCnt       int    `json:"me_cnt"`
	PeCnt       int    `json:"pe_cnt"`
	TeCnt       int    `json:"te_cnt"`
	NavPerm     string `json:"nav_perm"`
	Brv         int    `json:"brv"`
	Hbcalc      bool   `json:"hbCalc"`
	Fmh         string `json:"fmh"`
	Fmz         string `json:"fmz"`
	SSH         string `json:"ssh"`
	Wv          string `json:"wv"`
	Wr          string `json:"wr"`
	Weh         string `json:"weh"`
	Wl          int    `json:"wl"`
	Wen         int    `json:"wen"`
	Den         int    `json:"den"`
}

// returns a random Int
func randomNumberInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// returns a random Float64
func randomNumberFloat() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

// converts a Float64 ==> String and returns the String
func floatToString(flt float64, precision int) string {
	return strconv.FormatFloat(flt, 'f', precision, 64)
}
