package schema

import (
	"encoding/xml"
)

// O2BANK ...
type O2BANK struct {
	XMLName  xml.Name `xml:"O2_BANK"`
	O2BLZ    []string `xml:"O2_BLZ"`
	O2BNAM   []string `xml:"O2_B_NAM"`
	O2BKTO   []string `xml:"O2_B_KTO"`
	O2BLZSP  []string `xml:"O2_BLZ_SP"`
	O2BNAMSP []string `xml:"O2_B_NAM_SP"`
	O2BKTOSP []string `xml:"O2_B_KTO_SP"`
	O2CREDID []string `xml:"O2_CRED_ID"`
	O2IBAN   []string `xml:"O2_IBAN"`
	O2BIC    []string `xml:"O2_BIC"`
}

// SPA ...
type SPA struct {
	IDAttr int `xml:"ID,attr,omitempty"`
	*FOXType
}

// ABSCHNITT ...
type ABSCHNITT struct {
	IDAttr string      `xml:"ID,attr,omitempty"`
	KOPF   []*INFOType `xml:"KOPF"`
	SPA    []*SPA      `xml:"SPA"`
	FUSS   []*INFOType `xml:"FUSS"`
	TPC    []string    `xml:"TPC"`
}

// KZST ...
type KZST struct {
	XMLName xml.Name `xml:"K_ZST"`
	IDAttr  int      `xml:"ID,attr"`
	Value   string   `xml:",chardata"`
}

// KUNZUS ...
type KUNZUS struct {
	XMLName xml.Name `xml:"KUN_ZUS"`
	FDAT    []string `xml:"F_DAT"`
	ANUM    []string `xml:"A_NUM"`
	KNUM    []string `xml:"K_NUM"`
	MNUM    []string `xml:"M_NUM"`
	KZST    []*KZST  `xml:"K_ZST"`
	ANZNUM  []string `xml:"ANZ_NUM"`
	RVNUM   []string `xml:"RV_NUM"`
	USTNUM  []string `xml:"UST_NUM"`
	SUBNUM  []string `xml:"SUB_NUM"`
}

// ZAHZUS ...
type ZAHZUS struct {
	XMLName xml.Name `xml:"ZAH_ZUS"`
	ZIEDAT  []string `xml:"ZIE_DAT"`
}

// KOPF ...
type KOPF struct {
	ABSCHNITTAttr int `xml:"ABSCHNITT,attr"`
	*INFOType
}

// GES ...
type GES struct {
	ABSCHNITTAttr int `xml:"ABSCHNITT,attr"`
	*INFOType
}

// GUT ...
type GUT struct {
	ABSCHNITTAttr int `xml:"ABSCHNITT,attr,omitempty"`
	*FOXType
}

// VORMONAT ...
type VORMONAT struct {
	KOPF      []*KOPF    `xml:"KOPF"`
	GES       []*GES     `xml:"GES"`
	GUT       []*GUT     `xml:"GUT"`
	VORGES    []*FOXType `xml:"VOR_GES"`
	EINZAHGES []*FOXType `xml:"EIN_ZAH_GES"`
	VORGUNGES []*FOXType `xml:"VOR_GUN_GES"`
	VORMWS    []*FOXType `xml:"VOR_MWS"`
	VOROFF    []*FOXType `xml:"VOR_OFF"`
}

// AKTUELL ...
type AKTUELL struct {
	KOPF      []*FOXType      `xml:"KOPF"`
	GES       []*INFOType     `xml:"GES"`
	AKTGGB    []*BELTypeBesch `xml:"AKT_GGB"`
	AKTZSL    []*BELTypeBesch `xml:"AKT_ZSL"`
	AKTVER    []*BELTypeBesch `xml:"AKT_VER"`
	AKTSONRAB []*FOXType      `xml:"AKT_SON_RAB"`
	AKTGUNGES []*FOXType      `xml:"AKT_GUN_GES"`
	AKTGESN   []*FOXType      `xml:"AKT_GES_N"`
	AKTMWS    []*FOXType      `xml:"AKT_MWS"`
	AKTGESB   []*FOXType      `xml:"AKT_GES_B"`
	AKTRAB    []*BELTextType  `xml:"AKT_RAB"`
	AKTGESTPC []*FOXType      `xml:"AKT_GES_TPC"`
}

// RECZUS ...
type RECZUS struct {
	XMLName  xml.Name    `xml:"REC_ZUS"`
	RDAT     []string    `xml:"R_DAT"`
	RNUM     []string    `xml:"R_NUM"`
	RECPER   []string    `xml:"REC_PER"`
	RBIS     []string    `xml:"R_BIS"`
	RVON     []string    `xml:"R_VON"`
	LZ       []string    `xml:"LZ"`
	VORMONAT []*VORMONAT `xml:"VORMONAT"`
	AKTUELL  []*AKTUELL  `xml:"AKTUELL"`
	RGES     []*INFOType `xml:"R_GES"`
	RGUT     []*INFOType `xml:"R_GUT"`
	RUEB     []string    `xml:"R_UEB"`
	GTAX     []string    `xml:"G_TAX"`
}

// KORZUS ...
type KORZUS struct {
	XMLName xml.Name   `xml:"KOR_ZUS"`
	VORGUN  []*ZUSType `xml:"VOR_GUN"`
	AKTGUN  []*ZUSType `xml:"AKT_GUN"`
}

// KTO ...
type KTO struct {
	NAM       []*FOXType `xml:"NAM"`
	EINZAH    []*ZUSType `xml:"EIN_ZAH"`
	KORZUS    []*KORZUS  `xml:"KOR_ZUS"`
	GGBZUS    []*ZUSType `xml:"GGB_ZUS"`
	ZSLZUS    []*ZUSType `xml:"ZSL_ZUS"`
	RABZUS    []*ZUSType `xml:"RAB_ZUS"`
	GESKTO    []*FOXType `xml:"GES_KTO"`
	GESKTOBRU []*FOXType `xml:"GES_KTO_BRU"`
	KOPF      []*FOXType `xml:"KOPF"`
	BET       []*FOXType `xml:"BET"`
	BETBRU    []*FOXType `xml:"BET_BRU"`
	RABBET    []*FOXType `xml:"RAB_BET"`
}

// KTOBEL ...
type KTOBEL struct {
	XMLName      xml.Name   `xml:"KTO_BEL"`
	KTO          []*KTO     `xml:"KTO"`
	GESKTOBEL    []*FOXType `xml:"GES_KTO_BEL"`
	GESKTOBELBRU []*FOXType `xml:"GES_KTO_BEL_BRU"`
}

// DW ...
type DW struct {
	NAM      []*FOXType `xml:"NAM"`
	EINZAH   []*ZUSType `xml:"EIN_ZAH"`
	KORZUS   []*KORZUS  `xml:"KOR_ZUS"`
	GGBZUS   []*ZUSType `xml:"GGB_ZUS"`
	ZSLZUS   []*ZUSType `xml:"ZSL_ZUS"`
	RABZUS   []*ZUSType `xml:"RAB_ZUS"`
	GESDW    []*FOXType `xml:"GES_DW"`
	GESDWBRU []*FOXType `xml:"GES_DW_BRU"`
	KOPF     []*FOXType `xml:"KOPF"`
	BET      []*FOXType `xml:"BET"`
	BETBRU   []*FOXType `xml:"BET_BRU"`
	RABBET   []*FOXType `xml:"RAB_BET"`
}

// DWBEL ...
type DWBEL struct {
	XMLName     xml.Name   `xml:"DW_BEL"`
	DW          []*DW      `xml:"DW"`
	GESDWBEL    []*FOXType `xml:"GES_DW_BEL"`
	GESDWBELBRU []*FOXType `xml:"GES_DW_BEL_BRU"`
}

// TPCZUS ...
type TPCZUS struct {
	XMLName xml.Name   `xml:"TPC_ZUS"`
	SERVICE []*FOXType `xml:"SERVICE"`
	DETAIL  []*FOXType `xml:"DETAIL"`
}

// ABBRNAM ...
type ABBRNAM struct {
	XMLName xml.Name    `xml:"ABBR_NAM"`
	INFO    []*INFOType `xml:"INFO"`
}

// ABBREVN ...
type ABBREVN struct {
	XMLName xml.Name   `xml:"ABBR_EVN"`
	ABBRNAM []*ABBRNAM `xml:"ABBR_NAM"`
}

// SI ...
type SI struct {
	SNUM           []*FOXType `xml:"S_NUM"`
	KZST           []*KZST    `xml:"K_ZST"`
	KORZUS         []*KORZUS  `xml:"KOR_ZUS"`
	GGBZUS         []*ZUSType `xml:"GGB_ZUS"`
	ZSLZUS         []*ZUSType `xml:"ZSL_ZUS"`
	VERZUS         []*ZUSType `xml:"VER_ZUS"`
	RABZUS         []*ZUSType `xml:"RAB_ZUS"`
	GESSI          []*FOXType `xml:"GES_SI"`
	GESSIMITVORGUN []*FOXType `xml:"GES_SI_MIT_VOR_GUN"`
	BET            []*FOXType `xml:"BET"`
	RABBET         []*FOXType `xml:"RAB_BET"`
	NAMEVNFLN      []*FOXType `xml:"NAM_EVN_FLN"`
	GESTPC         []string   `xml:"GES_TPC"`
	TPC            []string   `xml:"TPC"`
	TPSPEVN        []*FOXType `xml:"TPSP_EVN"`
	TPCZUS         []*TPCZUS  `xml:"TPC_ZUS"`
	GESKTORABBET   []string   `xml:"GES_KTO_RAB_BET"`
	GESSIMITKTORAB []*FOXType `xml:"GES_SI_MIT_KTO_RAB"`
	ABBREVN        []*ABBREVN `xml:"ABBR_EVN"`
}

// TRAV ...
type TRAV struct {
	IDAttr int    `xml:"ID,attr"`
	Value  string `xml:",chardata"`
}

// CONTDATBEL ...
type CONTDATBEL struct {
	XMLName xml.Name    `xml:"CONT_DAT_BEL"`
	IDAttr  interface{} `xml:"ID,attr,omitempty"`
	TRAV    []*TRAV     `xml:"TRAV"`
	KOPF    []*KOPF     `xml:"KOPF"`
}

// CONTPAKDAT ...
type CONTPAKDAT struct {
	XMLName xml.Name `xml:"CONT_PAK_DAT"`
	TRAV    []*TRAV  `xml:"TRAV"`
}

// NAM ...
type NAM struct {
	IDAttr int    `xml:"ID,attr"`
	Value  string `xml:",chardata"`
}

// PRODGRP ...
type PRODGRP struct {
	XMLName xml.Name `xml:"PROD_GRP"`
	NAM     []*NAM   `xml:"NAM"`
}

// SIBEL ...
type SIBEL struct {
	XMLName    xml.Name      `xml:"SI_BEL"`
	SI         []*SI         `xml:"SI"`
	GESSIBEL   []*FOXType    `xml:"GES_SI_BEL"`
	CONTDATBEL []*CONTDATBEL `xml:"CONT_DAT_BEL"`
	CONTPAKDAT []*CONTPAKDAT `xml:"CONT_PAK_DAT"`
	PRODGRP    []*PRODGRP    `xml:"PROD_GRP"`
}

// MWSZUS ...
type MWSZUS struct {
	XMLName xml.Name    `xml:"MWS_ZUS"`
	NAM     []*INFOType `xml:"NAM"`
	KOPF    []*KOPF     `xml:"KOPF"`
}

// ADRZUS ...
type ADRZUS struct {
	XMLName xml.Name `xml:"ADR_ZUS"`
	VADR    []string `xml:"V_ADR"`
	CORRADR []string `xml:"CORR_ADR"`
}

// FORMAT ...
type FORMAT struct {
	FORMLANGAttr int          `xml:"FORM_LANG,attr,omitempty"`
	ABSCHNITT    []*ABSCHNITT `xml:"ABSCHNITT"`
	KUNZUS       []*KUNZUS    `xml:"KUN_ZUS"`
	ZAHZUS       []*ZAHZUS    `xml:"ZAH_ZUS"`
	RECZUS       []*RECZUS    `xml:"REC_ZUS"`
	KTOBEL       []*KTOBEL    `xml:"KTO_BEL"`
	DWBEL        []*DWBEL     `xml:"DW_BEL"`
	SIBEL        []*SIBEL     `xml:"SI_BEL"`
	ABSZ         []*INFOType  `xml:"ABS_Z"`
	INFO         []*INFOType  `xml:"INFO"`
	MWSZUS       []*MWSZUS    `xml:"MWS_ZUS"`
	ADRZUS       []*ADRZUS    `xml:"ADR_ZUS"`
}

// ARCHIVE ...
type ARCHIVE struct {
	A1Attr      int         `xml:"A1,attr,omitempty"`
	A2AAttr     int         `xml:"A2A,attr,omitempty"`
	A2BAttr     int         `xml:"A2B,attr,omitempty"`
	A2CAttr     int         `xml:"A2C,attr,omitempty"`
	A2DAttr     interface{} `xml:"A2D,attr,omitempty"`
	A3Attr      int         `xml:"A3,attr,omitempty"`
	A3AAttr     interface{} `xml:"A3A,attr,omitempty"`
	A3BAttr     interface{} `xml:"A3B,attr,omitempty"`
	EVNAttr     int         `xml:"EVN,attr,omitempty"`
	CSVAttr     int         `xml:"CSV,attr,omitempty"`
	ONLINEAttr  int         `xml:"ONLINE,attr"`
	ZUGRIFFAttr int         `xml:"ZUGRIFF,attr,omitempty"`
	BMAttr      int         `xml:"BM,attr,omitempty"`
}

// VADR ...
type VADR struct {
	XMLName xml.Name `xml:"V_ADR"`
	*PAdrType
}

// RADR ...
type RADR struct {
	XMLName   xml.Name `xml:"R_ADR"`
	HAUPTAttr bool     `xml:"HAUPT,attr,omitempty"`
	*PAdrType
}

// CDADR ...
type CDADR struct {
	XMLName xml.Name `xml:"CD_ADR"`
	*PAdrType
}

// EADR ...
type EADR struct {
	XMLName xml.Name `xml:"E_ADR"`
	*EADRType
}

// SMSADR ...
type SMSADR struct {
	XMLName xml.Name `xml:"SMS_ADR"`
	*SMSAdrType
}

// X400ADR ...
type X400ADR struct {
	XMLName xml.Name `xml:"X400_ADR"`
	*X400ADRType
}

// XRECHADR ...
type XRECHADR struct {
	XMLName xml.Name `xml:"XRECH_ADR"`
	*XRECHADRType
}

// KSEG ...
type KSEG struct {
	XMLName xml.Name `xml:"K_SEG"`
	Value   string   `xml:",chardata"`
}

// CONTENTPROVIDER ...
type CONTENTPROVIDER struct {
	XMLName    xml.Name `xml:"CONTENT_PROVIDER"`
	TARGETDESC []string `xml:"TARGET_DESC"`
	ANNOTATION []string `xml:"ANNOTATION"`
	AMT        []string `xml:"AMT"`
}

// RINFOTPC ...
type RINFOTPC struct {
	XMLName         xml.Name           `xml:"R_INFO_TPC"`
	PROVIDERCONTACT []string           `xml:"PROVIDER_CONTACT"`
	PROVIDERAMOUNT  []string           `xml:"PROVIDER_AMOUNT"`
	CONTENTPROVIDER []*CONTENTPROVIDER `xml:"CONTENT_PROVIDER"`
}

// INFOREF ...
type INFOREF struct {
	IDAttr  int `xml:"ID,attr"`
	POSAttr int `xml:"POS,attr"`
}

// EINZAH ...
type EINZAH struct {
	XMLName xml.Name    `xml:"EIN_ZAH"`
	GES     []string    `xml:"GES"`
	ZAH     []*RZahType `xml:"ZAH"`
}

// KORPLAZUS ...
type KORPLAZUS struct {
	XMLName xml.Name      `xml:"KOR_PLA_ZUS"`
	KORPLA  []*KORPLAType `xml:"KOR_PLA"`
}

// PADR ...
type PADR struct {
	XMLName xml.Name `xml:"P_ADR"`
	*PAdrType
}

// TPSPEVN ...
type TPSPEVN struct {
	XMLName xml.Name `xml:"TPSP_EVN"`
	TPSPNAM []string `xml:"TPSP_NAM"`
}

// GESSICONBEL ...
type GESSICONBEL struct {
	XMLName    xml.Name `xml:"GES_SI_CON_BEL"`
	CONTIDAttr string   `xml:"CONT_ID,attr"`
}

// CONTDATA ...
type CONTDATA struct {
	XMLName xml.Name `xml:"CONT_DATA"`
	TRAV    []*TRAV  `xml:"TRAV"`
}

// PAKDET ...
type PAKDET struct {
	XMLName xml.Name `xml:"PAK_DET"`
	TRAV    []*TRAV  `xml:"TRAV"`
}

// CONTBEL ...
type CONTBEL struct {
	XMLName    xml.Name      `xml:"CONT_BEL"`
	CONTDATA   []*CONTDATA   `xml:"CONT_DATA"`
	CONTPAKDAT []*CONTPAKDAT `xml:"CONT_PAK_DAT"`
}

// MWS ...
type MWS struct {
	*MWSType
}

// RECHNUNG ...
type RECHNUNG struct {
	IDAttr     int         `xml:"ID,attr,omitempty"`
	ANZSIAttr  int         `xml:"ANZ_SI,attr,omitempty"`
	ANZSMSAttr int         `xml:"ANZ_SMS,attr,omitempty"`
	ARCHIVE    []*ARCHIVE  `xml:"ARCHIVE"`
	ADRZUS     []*ADRZUS   `xml:"ADR_ZUS"`
	KUNZUS     []*KUNZUS   `xml:"KUN_ZUS"`
	RECZUS     []*RECZUS   `xml:"REC_ZUS"`
	RINFOTPC   []*RINFOTPC `xml:"R_INFO_TPC"`
	ZAHZUS     []*ZAHZUS   `xml:"ZAH_ZUS"`
	INFO       []*INFOType `xml:"INFO"`
	INFOREF    []*INFOREF  `xml:"INFOREF"`
	KTOBEL     []*KTOBEL   `xml:"KTO_BEL"`
	DWBEL      []*DWBEL    `xml:"DW_BEL"`
	SIBEL      []*SIBEL    `xml:"SI_BEL"`
	MWSZUS     []*MWSZUS   `xml:"MWS_ZUS"`
}

// CHECKSUM ...
type CHECKSUM struct {
	ANZRECAttr     int `xml:"ANZ_REC,attr,omitempty"`
	ANZPAPAttr     int `xml:"ANZ_PAP,attr,omitempty"`
	ANZCDAttr      int `xml:"ANZ_CD,attr,omitempty"`
	ANZX400Attr    int `xml:"ANZ_X400,attr,omitempty"`
	ANZXRECHAttr   int `xml:"ANZ_XRECH,attr,omitempty"`
	ANZEMAILAttr   int `xml:"ANZ_EMAIL,attr,omitempty"`
	ANZPAPEVNAttr  int `xml:"ANZ_PAP_EVN,attr,omitempty"`
	ANZCDEVNAttr   int `xml:"ANZ_CD_EVN,attr,omitempty"`
	ANZX400EVNAttr int `xml:"ANZ_X400_EVN,attr,omitempty"`
	ANZBMAttr      int `xml:"ANZ_BM,attr,omitempty"`
	ANZEVAttr      int `xml:"ANZ_EV,attr,omitempty"`
	ANZSIAttr      int `xml:"ANZ_SI,attr,omitempty"`
	ANZSMSAttr     int `xml:"ANZ_SMS,attr,omitempty"`
}

// PSI ...
type PSI struct {
	DATEAttr    string      `xml:"DATE,attr"`
	DESTAttr    string      `xml:"DEST,attr"`
	VERSIONAttr string      `xml:"VERSION,attr"`
	O2BANK      *O2BANK     `xml:"O2_BANK"`
	FORMAT      []string    `xml:"FORMAT"`
	RECHNUNG    []*RECHNUNG `xml:"RECHNUNG"`
	ANZ         []string    `xml:"ANZ"`
}

// ANZ is Version number of PSI file (for further processed at printshop currently MAGIC is the only valid value)
type ANZ []string

// SIZUS ...
type SIZUS struct {
	XMLName xml.Name `xml:"SI_ZUS"`
	IDAttr  string   `xml:"ID,attr"`
	GESN    []string `xml:"GES_N"`
	VERGESN []string `xml:"VER_GES_N"`
	ZSLGESN []string `xml:"ZSL_GES_N"`
	GESTPC  []string `xml:"GES_TPC"`
	GGBGESN []string `xml:"GGB_GES_N"`
}

// KZUO ...
type KZUO struct {
	XMLName xml.Name `xml:"K_ZUO"`
	IDAttr  string   `xml:"ID,attr,omitempty"`
	GESN    []string `xml:"GES_N"`
	VERGESN []string `xml:"VER_GES_N"`
	ZSLGESN []string `xml:"ZSL_GES_N"`
	GGBGESN []string `xml:"GGB_GES_N"`
	SIZUS   []*SIZUS `xml:"SI_ZUS"`
	GESTPC  []string `xml:"GES_TPC"`
}

// PAdrType ...
type PAdrType struct {
	XMLName      xml.Name `xml:"P_AdrType"`
	A1Attr       int      `xml:"A1,attr,omitempty"`
	A2AAttr      int      `xml:"A2A,attr,omitempty"`
	A2BAttr      int      `xml:"A2B,attr,omitempty"`
	A2CAttr      int      `xml:"A2C,attr,omitempty"`
	A2DAttr      int      `xml:"A2D,attr,omitempty"`
	A3Attr       int      `xml:"A3,attr,omitempty"`
	A3AAttr      int      `xml:"A3A,attr,omitempty"`
	A3BAttr      int      `xml:"A3B,attr,omitempty"`
	EVNAttr      int      `xml:"EVN,attr,omitempty"`
	ZVORAttr     int      `xml:"ZVOR,attr,omitempty"`
	BEILAGENAttr string   `xml:"BEILAGEN,attr,omitempty"`
	IDAttr       int      `xml:"ID,attr,omitempty"`
	ANR          string   `xml:"ANR"`
	TIT          string   `xml:"TIT"`
	VNAM         string   `xml:"V_NAM"`
	MNAM         string   `xml:"M_NAM"`
	NNAM         string   `xml:"N_NAM"`
	GNAM         string   `xml:"G_NAM"`
	NIE          string   `xml:"NIE"`
	ABT          string   `xml:"ABT"`
	ADR1         string   `xml:"ADR_1"`
	ADR2         string   `xml:"ADR_2"`
	ADR3         string   `xml:"ADR_3"`
	STA          string   `xml:"STA"`
	PLZ          string   `xml:"PLZ"`
	LAN          string   `xml:"LAN"`
	DIREKT       string   `xml:"DIREKT"`
	FIR          string   `xml:"FIR"`
}

// X400ADRType ...
type X400ADRType struct {
	XMLName xml.Name `xml:"X400_ADRType"`
	LAN     []string `xml:"LAN"`
	ADMD    []string `xml:"ADMD"`
	NNAM    []string `xml:"N_NAM"`
	VNAM    []string `xml:"V_NAM"`
	ORG     []string `xml:"ORG"`
	ORGE1   []string `xml:"ORG_E1"`
	ORGE2   []string `xml:"ORG_E2"`
	ORGE3   []string `xml:"ORG_E3"`
	ORGE4   []string `xml:"ORG_E4"`
	ANAM    []string `xml:"A_NAM"`
	LNAM    []*LNAM  `xml:"L_NAM"`
}

// LNAM ...
type LNAM []*LNAM

// XsString ...
type XsString string

// OU ...
type OU []*OU

// REC ...
type REC int

// EVN ...
type EVN int

// ID ...
type ID int

// XRECHADRType ...
type XRECHADRType struct {
	XMLName    xml.Name `xml:"XRECH_ADRType"`
	IDAttr     int      `xml:"ID,attr"`
	XLEITWEG   []string `xml:"X_LEITWEG"`
	XEMAIL     []string `xml:"X_EMAIL"`
	XAUFTRAGNR []string `xml:"X_AUFTRAGNR"`
	XLIEFERANT []string `xml:"X_LIEFERANT"`
	XVERSAND   []string `xml:"X_VERSAND"`
}

// SMSAdrType ...
type SMSAdrType struct {
	XMLName    xml.Name `xml:"SMS_AdrType"`
	IDAttr     int      `xml:"ID,attr,omitempty"`
	NOTIFYAttr bool     `xml:"NOTIFY,attr,omitempty"`
	MSISDN     string   `xml:"MSISDN"`
}

// EADRType ...
type EADRType struct {
	XMLName    xml.Name `xml:"E_ADRType"`
	IDAttr     int      `xml:"ID,attr"`
	A1Attr     int      `xml:"A1,attr,omitempty"`
	A2AAttr    int      `xml:"A2A,attr,omitempty"`
	A2BAttr    int      `xml:"A2B,attr,omitempty"`
	A2CAttr    int      `xml:"A2C,attr,omitempty"`
	A3Attr     int      `xml:"A3,attr,omitempty"`
	EVNAttr    int      `xml:"EVN,attr,omitempty"`
	NOTIFYAttr bool     `xml:"NOTIFY,attr,omitempty"`
	ADR        string   `xml:"ADR"`
}

// GGBZUSType ...
type GGBZUSType struct {
	XMLName    xml.Name    `xml:"GGB_ZUSType"`
	GES        []string    `xml:"GES"`
	GESBRU     []string    `xml:"GES_BRU"`
	GESUDP     []string    `xml:"GES_UDP"`
	GESUDPBRU  []string    `xml:"GES_UDP_BRU"`
	RABBET     []string    `xml:"RAB_BET"`
	GGB        []*RBelType `xml:"GGB"`
	INKLRABBET []string    `xml:"INKL_RAB_BET"`
	KTORABBET  []string    `xml:"KTO_RAB_BET"`
	GESN       []string    `xml:"GES_N"`
	GESB       []string    `xml:"GES_B"`
}

// ZSLZUSType ...
type ZSLZUSType struct {
	XMLName    xml.Name    `xml:"ZSL_ZUSType"`
	GES        []string    `xml:"GES"`
	GESBRU     []string    `xml:"GES_BRU"`
	GESUDP     []string    `xml:"GES_UDP"`
	GESUDPBRU  []string    `xml:"GES_UDP_BRU"`
	RABBET     []string    `xml:"RAB_BET"`
	ZSL        []*RBelType `xml:"ZSL"`
	INKLRABBET []string    `xml:"INKL_RAB_BET"`
	KTORABBET  []string    `xml:"KTO_RAB_BET"`
	GESN       []string    `xml:"GES_N"`
	GESB       []string    `xml:"GES_B"`
}

// BELType ...
type BELType struct {
	BET    string `xml:"BET"`
	RABBET string `xml:"RAB_BET"`
	MWS    *MWS   `xml:"MWS"`
}

// RABZUSType ...
type RABZUSType struct {
	XMLName   xml.Name    `xml:"RAB_ZUSType"`
	GES       []string    `xml:"GES"`
	GESBRU    []string    `xml:"GES_BRU"`
	GESUDP    []string    `xml:"GES_UDP"`
	GESUDPBRU []string    `xml:"GES_UDP_BRU"`
	RAB       []*RBelType `xml:"RAB"`
}

// RBelType ...
type RBelType struct {
	XMLName   xml.Name    `xml:"R_BelType"`
	HAUPTAttr bool        `xml:"HAUPT,attr,omitempty"`
	TYPEAttr  string      `xml:"TYPE,attr,omitempty"`
	NAM       []*INFOType `xml:"NAM"`
	RNAM      []*INFOType `xml:"R_NAM"`
	LIN       []string    `xml:"LIN"`
	BET       []string    `xml:"BET"`
	BETBRU    []string    `xml:"BET_BRU"`
	RABBET    []string    `xml:"RAB_BET"`
	RES       []string    `xml:"RES"`
	VON       []string    `xml:"VON"`
	BIS       []string    `xml:"BIS"`
	DAT       []string    `xml:"DAT"`
	PER       []float64   `xml:"PER"`
	ID        []int       `xml:"ID"`
	BETN      []string    `xml:"BET_N"`
	BETB      []string    `xml:"BET_B"`
	RNAMB     []*INFOType `xml:"R_NAMB"`
	COUNT     []int       `xml:"COUNT"`
	KTORABBET []float32   `xml:"KTO_RAB_BET"`
	MWSSAT    []string    `xml:"MWS_SAT"`
}

// KORPLAType ...
type KORPLAType struct {
	XMLName xml.Name    `xml:"KOR_PLAType"`
	NAM     []*INFOType `xml:"NAM"`
	URS     []string    `xml:"URS"`
	ZIE     []string    `xml:"ZIE"`
	VON     []string    `xml:"VON"`
	BIS     []string    `xml:"BIS"`
}

// FELDREF ...
type FELDREF struct {
	NAMEAttr string `xml:"NAME,attr"`
}

// GESType ...
type GESType struct {
	ABSCHNITTAttr string     `xml:"ABSCHNITT,attr,omitempty"`
	FELDREF       []*FELDREF `xml:"FELDREF"`
}

// INFOType ...
type INFOType struct {
	IDAttr  int `xml:"ID,attr,omitempty"`
	POSAttr int `xml:"POS,attr,omitempty"`
	*FOXType
}

// KZSTType ...
type KZSTType struct {
	XMLName xml.Name `xml:"K_ZSTType"`
	NAM     string   `xml:"NAM"`
	GES     string   `xml:"GES"`
	RABGES  string   `xml:"RAB_GES"`
	GESN    string   `xml:"GES_N"`
}

// BELTypeBesch ...
type BELTypeBesch struct {
	BET    []*FOXType `xml:"BET"`
	RABBET []*FOXType `xml:"RAB_BET"`
}

// RZahType ...
type RZahType struct {
	XMLName xml.Name    `xml:"R_ZahType"`
	BET     []string    `xml:"BET"`
	NAM     []string    `xml:"NAM"`
	RNAM    []*INFOType `xml:"R_NAM"`
	DAT     []string    `xml:"DAT"`
}

// EV ...
type EV struct {
	SORTAttr     int    `xml:"SORT,attr,omitempty"`
	SUPPRESSAttr bool   `xml:"SUPPRESS,attr,omitempty"`
	SPA          []*SPA `xml:"SPA"`
}

// GRPREF ...
type GRPREF struct {
	IDAttr int         `xml:"ID,attr"`
	SEK    []int       `xml:"SEK"`
	GES    []string    `xml:"GES"`
	GESBRU []string    `xml:"GES_BRU"`
	RABBET []string    `xml:"RAB_BET"`
	NAM    []*INFOType `xml:"NAM"`
	EV     []*EV       `xml:"EV"`
	GESN   []*GESN     `xml:"GES_N"`
	GESB   []*GESB     `xml:"GES_B"`
	GESBP  []string    `xml:"GES_BP"`
	TYPE   []string    `xml:"TYPE"`
}

// GESZUS ...
type GESZUS struct {
	XMLName xml.Name `xml:"GES_ZUS"`
	GES     []*GES   `xml:"GES"`
}

// VERZUSType ...
type VERZUSType struct {
	XMLName    xml.Name    `xml:"VER_ZUSType"`
	GES        []*GES      `xml:"GES"`
	GESBRU     []string    `xml:"GES_BRU"`
	GESUDP     []string    `xml:"GES_UDP"`
	GESUDPBRU  []string    `xml:"GES_UDP_BRU"`
	RABBET     []string    `xml:"RAB_BET"`
	GRPREF     []*GRPREF   `xml:"GRPREF"`
	INFOREF    []*INFOType `xml:"INFOREF"`
	INKLRABBET []string    `xml:"INKL_RAB_BET"`
	KTORABBET  []string    `xml:"KTO_RAB_BET"`
	GESZUS     []*GESZUS   `xml:"GES_ZUS"`
	GESN       []*GESN     `xml:"GES_N"`
	GESB       []*GESB     `xml:"GES_B"`
	GESBP      []string    `xml:"GES_BP"`
	GESTPC     []string    `xml:"GES_TPC"`
	GESUDPTPC  []string    `xml:"GES_UDP_TPC"`
	GESEVN     []string    `xml:"GES_EVN"`
	GESTPCB    []string    `xml:"GES_TPC_B"`
	GESTPCN    []string    `xml:"GES_TPC_N"`
}

// ENTZUSType ...
type ENTZUSType struct {
	XMLName   xml.Name `xml:"ENT_ZUSType"`
	GES       []*GES   `xml:"GES"`
	GESBRU    []string `xml:"GES_BRU"`
	GESUDP    []string `xml:"GES_UDP"`
	GESUDPBRU []string `xml:"GES_UDP_BRU"`
}

// MWSType ...
type MWSType struct {
	BET  []string  `xml:"BET"`
	SAT  []float32 `xml:"SAT"`
	BETN []float64 `xml:"BET_N"`
}

// MWSZUSType ...
type MWSZUSType struct {
	XMLName xml.Name   `xml:"MWS_ZUSType"`
	GES     string     `xml:"GES"`
	MWS     []*MWSType `xml:"MWS"`
}

// GESN ...
type GESN struct {
	XMLName       xml.Name `xml:"GES_N"`
	ABSCHNITTAttr string   `xml:"ABSCHNITT,attr,omitempty"`
	*INFOType
}

// GESB ...
type GESB struct {
	XMLName       xml.Name `xml:"GES_B"`
	ABSCHNITTAttr string   `xml:"ABSCHNITT,attr,omitempty"`
	*INFOType
}

// GRP ...
type GRP struct {
	IDAttr   interface{} `xml:"ID,attr,omitempty"`
	TYPEAttr interface{} `xml:"TYPE,attr,omitempty"`
	SPA      []*SPA      `xml:"SPA"`
}

// ZUSType ...
type ZUSType struct {
	GES       []*GES       `xml:"GES"`
	KOPF      []*KOPF      `xml:"KOPF"`
	KTORABBET []string     `xml:"KTO_RAB_BET"`
	GESN      []*GESN      `xml:"GES_N"`
	GESB      []*GESB      `xml:"GES_B"`
	UEBZ      []*FOXType   `xml:"UEB_Z"`
	GESTPC    []*FOXType   `xml:"GES_TPC"`
	GRP       []*GRP       `xml:"GRP"`
	ABSCHNITT []*ABSCHNITT `xml:"ABSCHNITT"`
}

// VORGUN ...
type VORGUN struct {
	XMLName xml.Name    `xml:"VOR_GUN"`
	GES     []string    `xml:"GES"`
	GESBRU  []string    `xml:"GES_BRU"`
	KOR     []*RBelType `xml:"KOR"`
}

// AKTGUN ...
type AKTGUN struct {
	XMLName xml.Name    `xml:"AKT_GUN"`
	GES     []string    `xml:"GES"`
	GESBRU  []string    `xml:"GES_BRU"`
	KOR     []*RBelType `xml:"KOR"`
}

// KORZUSType ...
type KORZUSType struct {
	XMLName xml.Name  `xml:"KOR_ZUSType"`
	GES     []string  `xml:"GES"`
	GESBRU  []string  `xml:"GES_BRU"`
	VORGUN  []*VORGUN `xml:"VOR_GUN"`
	AKTGUN  []*AKTGUN `xml:"AKT_GUN"`
}

// BELTextType ...
type BELTextType struct {
	BET    []*FOXType `xml:"BET"`
	RABBET []*FOXType `xml:"RAB_BET"`
	MWS    []*FOXType `xml:"MWS"`
}
