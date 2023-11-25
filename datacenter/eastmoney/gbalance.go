// 获取资产负债表数据

package eastmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	"go.uber.org/zap"
)

// balanceData 资产负债表数据
type BalanceData struct {
	SECUCODE                string      `json:"SECUCODE"`
	SECURITYCODE            string      `json:"SECURITY_CODE"`
	SECURITYNAMEABBR        string      `json:"SECURITY_NAME_ABBR"`
	ORGCODE                 string      `json:"ORG_CODE"`
	ORGTYPE                 string      `json:"ORG_TYPE"`
	REPORTDATE              string      `json:"REPORT_DATE"`
	REPORTTYPE              string      `json:"REPORT_TYPE"`
	REPORTDATENAME          string      `json:"REPORT_DATE_NAME"`
	SECURITYTYPECODE        string      `json:"SECURITY_TYPE_CODE"`
	NOTICEDATE              string      `json:"NOTICE_DATE"`
	UPDATEDATE              string      `json:"UPDATE_DATE"`
	CURRENCY                string      `json:"CURRENCY"`
	ACCEPTDEPOSITINTERBANK  interface{} `json:"ACCEPT_DEPOSIT_INTERBANK"`
	ACCOUNTSPAYABLE         float64     `json:"ACCOUNTS_PAYABLE"`
	ACCOUNTSRECE            float64     `json:"ACCOUNTS_RECE"`
	ACCRUEDEXPENSE          interface{} `json:"ACCRUED_EXPENSE"`
	ADVANCERECEIVABLES      float64     `json:"ADVANCE_RECEIVABLES"`
	AGENTTRADESECURITY      interface{} `json:"AGENT_TRADE_SECURITY"`
	AGENTUNDERWRITESECURITY interface{} `json:"AGENT_UNDERWRITE_SECURITY"`
	AMORTIZECOSTFINASSET    interface{} `json:"AMORTIZE_COST_FINASSET"`
	AMORTIZECOSTFINLIAB     interface{} `json:"AMORTIZE_COST_FINLIAB"`
	AMORTIZECOSTNCFINASSET  interface{} `json:"AMORTIZE_COST_NCFINASSET"`
	AMORTIZECOSTNCFINLIAB   interface{} `json:"AMORTIZE_COST_NCFINLIAB"`
	APPOINTFVTPLFINASSET    interface{} `json:"APPOINT_FVTPL_FINASSET"`
	APPOINTFVTPLFINLIAB     interface{} `json:"APPOINT_FVTPL_FINLIAB"`
	ASSETBALANCE            int         `json:"ASSET_BALANCE"`
	ASSETOTHER              interface{} `json:"ASSET_OTHER"`
	ASSIGNCASHDIVIDEND      interface{} `json:"ASSIGN_CASH_DIVIDEND"`
	AVAILABLESALEFINASSET   interface{} `json:"AVAILABLE_SALE_FINASSET"`
	BONDPAYABLE             interface{} `json:"BOND_PAYABLE"`
	BORROWFUND              interface{} `json:"BORROW_FUND"`
	BUYRESALEFINASSET       interface{} `json:"BUY_RESALE_FINASSET"`
	CAPITALRESERVE          float64     `json:"CAPITAL_RESERVE"`
	// 在建工程
	CIP                        float64     `json:"CIP"`
	CONSUMPTIVEBIOLOGICALASSET interface{} `json:"CONSUMPTIVE_BIOLOGICAL_ASSET"`
	CONTRACTASSET              interface{} `json:"CONTRACT_ASSET"`
	CONTRACTLIAB               float64     `json:"CONTRACT_LIAB"`
	CONVERTDIFF                interface{} `json:"CONVERT_DIFF"`
	CREDITORINVEST             interface{} `json:"CREDITOR_INVEST"`
	CURRENTASSETBALANCE        int         `json:"CURRENT_ASSET_BALANCE"`
	CURRENTASSETOTHER          interface{} `json:"CURRENT_ASSET_OTHER"`
	CURRENTLIABBALANCE         int         `json:"CURRENT_LIAB_BALANCE"`
	CURRENTLIABOTHER           interface{} `json:"CURRENT_LIAB_OTHER"`
	DEFERINCOME                float64     `json:"DEFER_INCOME"`
	DEFERINCOME1YEAR           interface{} `json:"DEFER_INCOME_1YEAR"`
	DEFERTAXASSET              float64     `json:"DEFER_TAX_ASSET"`
	DEFERTAXLIAB               float64     `json:"DEFER_TAX_LIAB"`
	DERIVEFINASSET             interface{} `json:"DERIVE_FINASSET"`
	DERIVEFINLIAB              interface{} `json:"DERIVE_FINLIAB"`
	DEVELOPEXPENSE             interface{} `json:"DEVELOP_EXPENSE"`
	DIVHOLDSALEASSET           interface{} `json:"DIV_HOLDSALE_ASSET"`
	DIVHOLDSALELIAB            interface{} `json:"DIV_HOLDSALE_LIAB"`
	DIVIDENDPAYABLE            float64     `json:"DIVIDEND_PAYABLE"`
	DIVIDENDRECE               interface{} `json:"DIVIDEND_RECE"`
	EQUITYBALANCE              int         `json:"EQUITY_BALANCE"`
	EQUITYOTHER                interface{} `json:"EQUITY_OTHER"`
	EXPORTREFUNDRECE           interface{} `json:"EXPORT_REFUND_RECE"`
	FEECOMMISSIONPAYABLE       interface{} `json:"FEE_COMMISSION_PAYABLE"`
	FINFUND                    interface{} `json:"FIN_FUND"`
	FINANCERECE                float64     `json:"FINANCE_RECE"`
	// 固定资产
	FIXEDASSET               float64     `json:"FIXED_ASSET"`
	FIXEDASSETDISPOSAL       interface{} `json:"FIXED_ASSET_DISPOSAL"`
	FVTOCIFINASSET           interface{} `json:"FVTOCI_FINASSET"`
	FVTOCINCFINASSET         interface{} `json:"FVTOCI_NCFINASSET"`
	FVTPLFINASSET            interface{} `json:"FVTPL_FINASSET"`
	FVTPLFINLIAB             interface{} `json:"FVTPL_FINLIAB"`
	GENERALRISKRESERVE       interface{} `json:"GENERAL_RISK_RESERVE"`
	GOODWILL                 interface{} `json:"GOODWILL"`
	HOLDMATURITYINVEST       interface{} `json:"HOLD_MATURITY_INVEST"`
	HOLDSALEASSET            interface{} `json:"HOLDSALE_ASSET"`
	HOLDSALELIAB             interface{} `json:"HOLDSALE_LIAB"`
	INSURANCECONTRACTRESERVE interface{} `json:"INSURANCE_CONTRACT_RESERVE"`
	INTANGIBLEASSET          float64     `json:"INTANGIBLE_ASSET"`
	INTERESTPAYABLE          interface{} `json:"INTEREST_PAYABLE"`
	INTERESTRECE             interface{} `json:"INTEREST_RECE"`
	INTERNALPAYABLE          interface{} `json:"INTERNAL_PAYABLE"`
	INTERNALRECE             interface{} `json:"INTERNAL_RECE"`
	// 存货
	INVENTORY         float64     `json:"INVENTORY"`
	INVESTREALESTATE  float64     `json:"INVEST_REALESTATE"`
	LEASELIAB         float64     `json:"LEASE_LIAB"`
	LENDFUND          interface{} `json:"LEND_FUND"`
	LIABBALANCE       int         `json:"LIAB_BALANCE"`
	LIABEQUITYBALANCE interface{} `json:"LIAB_EQUITY_BALANCE"`
	LIABEQUITYOTHER   interface{} `json:"LIAB_EQUITY_OTHER"`
	LIABOTHER         interface{} `json:"LIAB_OTHER"`
	LOANADVANCE       interface{} `json:"LOAN_ADVANCE"`
	LOANPBC           interface{} `json:"LOAN_PBC"`
	// 长期股权投资
	LONGEQUITYINVEST       int         `json:"LONG_EQUITY_INVEST"`
	LONGLOAN               interface{} `json:"LONG_LOAN"`
	LONGPAYABLE            interface{} `json:"LONG_PAYABLE"`
	LONGPREPAIDEXPENSE     interface{} `json:"LONG_PREPAID_EXPENSE"`
	LONGRECE               interface{} `json:"LONG_RECE"`
	LONGSTAFFSALARYPAYABLE float64     `json:"LONG_STAFFSALARY_PAYABLE"`
	MINORITYEQUITY         float64     `json:"MINORITY_EQUITY"`
	// 货币资金
	MONETARYFUNDS          float64     `json:"MONETARYFUNDS"`
	NONCURRENTASSET1YEAR   interface{} `json:"NONCURRENT_ASSET_1YEAR"`
	NONCURRENTASSETBALANCE int         `json:"NONCURRENT_ASSET_BALANCE"`
	NONCURRENTASSETOTHER   interface{} `json:"NONCURRENT_ASSET_OTHER"`
	NONCURRENTLIAB1YEAR    float64     `json:"NONCURRENT_LIAB_1YEAR"`
	NONCURRENTLIABBALANCE  int         `json:"NONCURRENT_LIAB_BALANCE"`
	NONCURRENTLIABOTHER    interface{} `json:"NONCURRENT_LIAB_OTHER"`
	NOTEACCOUNTSPAYABLE    float64     `json:"NOTE_ACCOUNTS_PAYABLE"`
	// 应收账款
	NOTEACCOUNTSRECE        float64     `json:"NOTE_ACCOUNTS_RECE"`
	NOTEPAYABLE             float64     `json:"NOTE_PAYABLE"`
	NOTERECE                interface{} `json:"NOTE_RECE"`
	OILGASASSET             interface{} `json:"OIL_GAS_ASSET"`
	OTHERCOMPREINCOME       float64     `json:"OTHER_COMPRE_INCOME"`
	OTHERCREDITORINVEST     interface{} `json:"OTHER_CREDITOR_INVEST"`
	OTHERCURRENTASSET       float64     `json:"OTHER_CURRENT_ASSET"`
	OTHERCURRENTLIAB        float64     `json:"OTHER_CURRENT_LIAB"`
	OTHEREQUITYINVEST       float64     `json:"OTHER_EQUITY_INVEST"`
	OTHEREQUITYOTHER        interface{} `json:"OTHER_EQUITY_OTHER"`
	OTHEREQUITYTOOL         interface{} `json:"OTHER_EQUITY_TOOL"`
	OTHERNONCURRENTASSET    float64     `json:"OTHER_NONCURRENT_ASSET"`
	OTHERNONCURRENTFINASSET interface{} `json:"OTHER_NONCURRENT_FINASSET"`
	OTHERNONCURRENTLIAB     interface{} `json:"OTHER_NONCURRENT_LIAB"`
	OTHERPAYABLE            interface{} `json:"OTHER_PAYABLE"`
	OTHERRECE               interface{} `json:"OTHER_RECE"`
	PARENTEQUITYBALANCE     int         `json:"PARENT_EQUITY_BALANCE"`
	PARENTEQUITYOTHER       interface{} `json:"PARENT_EQUITY_OTHER"`
	PERPETUALBOND           interface{} `json:"PERPETUAL_BOND"`
	PERPETUALBONDPAYBALE    interface{} `json:"PERPETUAL_BOND_PAYBALE"`
	PREDICTCURRENTLIAB      interface{} `json:"PREDICT_CURRENT_LIAB"`
	PREDICTLIAB             interface{} `json:"PREDICT_LIAB"`
	PREFERREDSHARES         interface{} `json:"PREFERRED_SHARES"`
	PREFERREDSHARESPAYBALE  interface{} `json:"PREFERRED_SHARES_PAYBALE"`
	PREMIUMRECE             interface{} `json:"PREMIUM_RECE"`
	// 预付
	PREPAYMENT                    float64     `json:"PREPAYMENT"`
	PRODUCTIVEBIOLOGYASSET        interface{} `json:"PRODUCTIVE_BIOLOGY_ASSET"`
	PROJECTMATERIAL               interface{} `json:"PROJECT_MATERIAL"`
	RCRESERVERECE                 interface{} `json:"RC_RESERVE_RECE"`
	REINSUREPAYABLE               interface{} `json:"REINSURE_PAYABLE"`
	REINSURERECE                  interface{} `json:"REINSURE_RECE"`
	SELLREPOFINASSET              interface{} `json:"SELL_REPO_FINASSET"`
	SETTLEEXCESSRESERVE           interface{} `json:"SETTLE_EXCESS_RESERVE"`
	SHARECAPITAL                  int         `json:"SHARE_CAPITAL"`
	SHORTBONDPAYABLE              interface{} `json:"SHORT_BOND_PAYABLE"`
	SHORTFINPAYABLE               interface{} `json:"SHORT_FIN_PAYABLE"`
	SHORTLOAN                     interface{} `json:"SHORT_LOAN"`
	SPECIALPAYABLE                interface{} `json:"SPECIAL_PAYABLE"`
	SPECIALRESERVE                float64     `json:"SPECIAL_RESERVE"`
	STAFFSALARYPAYABLE            float64     `json:"STAFF_SALARY_PAYABLE"`
	SUBSIDYRECE                   interface{} `json:"SUBSIDY_RECE"`
	SURPLUSRESERVE                float64     `json:"SURPLUS_RESERVE"`
	TAXPAYABLE                    float64     `json:"TAX_PAYABLE"`
	TOTALASSETS                   float64     `json:"TOTAL_ASSETS"`
	TOTALCURRENTASSETS            float64     `json:"TOTAL_CURRENT_ASSETS"`
	TOTALCURRENTLIAB              float64     `json:"TOTAL_CURRENT_LIAB"`
	TOTALEQUITY                   float64     `json:"TOTAL_EQUITY"`
	TOTALLIABEQUITY               float64     `json:"TOTAL_LIAB_EQUITY"`
	TOTALLIABILITIES              float64     `json:"TOTAL_LIABILITIES"`
	TOTALNONCURRENTASSETS         float64     `json:"TOTAL_NONCURRENT_ASSETS"`
	TOTALNONCURRENTLIAB           float64     `json:"TOTAL_NONCURRENT_LIAB"`
	TOTALOTHERPAYABLE             float64     `json:"TOTAL_OTHER_PAYABLE"`
	TOTALOTHERRECE                float64     `json:"TOTAL_OTHER_RECE"`
	TOTALPARENTEQUITY             float64     `json:"TOTAL_PARENT_EQUITY"`
	TRADEFINASSET                 interface{} `json:"TRADE_FINASSET"`
	TRADEFINASSETNOTFVTPL         float64     `json:"TRADE_FINASSET_NOTFVTPL"`
	TRADEFINLIAB                  interface{} `json:"TRADE_FINLIAB"`
	TRADEFINLIABNOTFVTPL          interface{} `json:"TRADE_FINLIAB_NOTFVTPL"`
	TREASURYSHARES                int         `json:"TREASURY_SHARES"`
	UNASSIGNRPOFIT                float64     `json:"UNASSIGN_RPOFIT"`
	UNCONFIRMINVESTLOSS           interface{} `json:"UNCONFIRM_INVEST_LOSS"`
	USERIGHTASSET                 float64     `json:"USERIGHT_ASSET"`
	ACCEPTDEPOSITINTERBANKYOY     interface{} `json:"ACCEPT_DEPOSIT_INTERBANK_YOY"`
	ACCOUNTSPAYABLEYOY            float64     `json:"ACCOUNTS_PAYABLE_YOY"`
	ACCOUNTSRECEYOY               float64     `json:"ACCOUNTS_RECE_YOY"`
	ACCRUEDEXPENSEYOY             interface{} `json:"ACCRUED_EXPENSE_YOY"`
	ADVANCERECEIVABLESYOY         float64     `json:"ADVANCE_RECEIVABLES_YOY"`
	AGENTTRADESECURITYYOY         interface{} `json:"AGENT_TRADE_SECURITY_YOY"`
	AGENTUNDERWRITESECURITYYOY    interface{} `json:"AGENT_UNDERWRITE_SECURITY_YOY"`
	AMORTIZECOSTFINASSETYOY       interface{} `json:"AMORTIZE_COST_FINASSET_YOY"`
	AMORTIZECOSTFINLIABYOY        interface{} `json:"AMORTIZE_COST_FINLIAB_YOY"`
	AMORTIZECOSTNCFINASSETYOY     interface{} `json:"AMORTIZE_COST_NCFINASSET_YOY"`
	AMORTIZECOSTNCFINLIABYOY      interface{} `json:"AMORTIZE_COST_NCFINLIAB_YOY"`
	APPOINTFVTPLFINASSETYOY       interface{} `json:"APPOINT_FVTPL_FINASSET_YOY"`
	APPOINTFVTPLFINLIABYOY        interface{} `json:"APPOINT_FVTPL_FINLIAB_YOY"`
	ASSETBALANCEYOY               interface{} `json:"ASSET_BALANCE_YOY"`
	ASSETOTHERYOY                 interface{} `json:"ASSET_OTHER_YOY"`
	ASSIGNCASHDIVIDENDYOY         interface{} `json:"ASSIGN_CASH_DIVIDEND_YOY"`
	AVAILABLESALEFINASSETYOY      interface{} `json:"AVAILABLE_SALE_FINASSET_YOY"`
	BONDPAYABLEYOY                interface{} `json:"BOND_PAYABLE_YOY"`
	BORROWFUNDYOY                 interface{} `json:"BORROW_FUND_YOY"`
	BUYRESALEFINASSETYOY          interface{} `json:"BUY_RESALE_FINASSET_YOY"`
	CAPITALRESERVEYOY             float64     `json:"CAPITAL_RESERVE_YOY"`
	CIPYOY                        float64     `json:"CIP_YOY"`
	CONSUMPTIVEBIOLOGICALASSETYOY interface{} `json:"CONSUMPTIVE_BIOLOGICAL_ASSET_YOY"`
	CONTRACTASSETYOY              interface{} `json:"CONTRACT_ASSET_YOY"`
	CONTRACTLIABYOY               float64     `json:"CONTRACT_LIAB_YOY"`
	CONVERTDIFFYOY                interface{} `json:"CONVERT_DIFF_YOY"`
	CREDITORINVESTYOY             interface{} `json:"CREDITOR_INVEST_YOY"`
	CURRENTASSETBALANCEYOY        interface{} `json:"CURRENT_ASSET_BALANCE_YOY"`
	CURRENTASSETOTHERYOY          interface{} `json:"CURRENT_ASSET_OTHER_YOY"`
	CURRENTLIABBALANCEYOY         interface{} `json:"CURRENT_LIAB_BALANCE_YOY"`
	CURRENTLIABOTHERYOY           interface{} `json:"CURRENT_LIAB_OTHER_YOY"`
	DEFERINCOME1YEARYOY           interface{} `json:"DEFER_INCOME_1YEAR_YOY"`
	DEFERINCOMEYOY                float64     `json:"DEFER_INCOME_YOY"`
	DEFERTAXASSETYOY              float64     `json:"DEFER_TAX_ASSET_YOY"`
	DEFERTAXLIABYOY               float64     `json:"DEFER_TAX_LIAB_YOY"`
	DERIVEFINASSETYOY             interface{} `json:"DERIVE_FINASSET_YOY"`
	DERIVEFINLIABYOY              interface{} `json:"DERIVE_FINLIAB_YOY"`
	DEVELOPEXPENSEYOY             interface{} `json:"DEVELOP_EXPENSE_YOY"`
	DIVHOLDSALEASSETYOY           interface{} `json:"DIV_HOLDSALE_ASSET_YOY"`
	DIVHOLDSALELIABYOY            interface{} `json:"DIV_HOLDSALE_LIAB_YOY"`
	DIVIDENDPAYABLEYOY            int         `json:"DIVIDEND_PAYABLE_YOY"`
	DIVIDENDRECEYOY               interface{} `json:"DIVIDEND_RECE_YOY"`
	EQUITYBALANCEYOY              interface{} `json:"EQUITY_BALANCE_YOY"`
	EQUITYOTHERYOY                interface{} `json:"EQUITY_OTHER_YOY"`
	EXPORTREFUNDRECEYOY           interface{} `json:"EXPORT_REFUND_RECE_YOY"`
	FEECOMMISSIONPAYABLEYOY       interface{} `json:"FEE_COMMISSION_PAYABLE_YOY"`
	FINFUNDYOY                    interface{} `json:"FIN_FUND_YOY"`
	FINANCERECEYOY                float64     `json:"FINANCE_RECE_YOY"`
	FIXEDASSETDISPOSALYOY         interface{} `json:"FIXED_ASSET_DISPOSAL_YOY"`
	FIXEDASSETYOY                 float64     `json:"FIXED_ASSET_YOY"`
	FVTOCIFINASSETYOY             interface{} `json:"FVTOCI_FINASSET_YOY"`
	FVTOCINCFINASSETYOY           interface{} `json:"FVTOCI_NCFINASSET_YOY"`
	FVTPLFINASSETYOY              interface{} `json:"FVTPL_FINASSET_YOY"`
	FVTPLFINLIABYOY               interface{} `json:"FVTPL_FINLIAB_YOY"`
	GENERALRISKRESERVEYOY         interface{} `json:"GENERAL_RISK_RESERVE_YOY"`
	GOODWILLYOY                   interface{} `json:"GOODWILL_YOY"`
	HOLDMATURITYINVESTYOY         interface{} `json:"HOLD_MATURITY_INVEST_YOY"`
	HOLDSALEASSETYOY              interface{} `json:"HOLDSALE_ASSET_YOY"`
	HOLDSALELIABYOY               interface{} `json:"HOLDSALE_LIAB_YOY"`
	INSURANCECONTRACTRESERVEYOY   interface{} `json:"INSURANCE_CONTRACT_RESERVE_YOY"`
	INTANGIBLEASSETYOY            float64     `json:"INTANGIBLE_ASSET_YOY"`
	INTERESTPAYABLEYOY            interface{} `json:"INTEREST_PAYABLE_YOY"`
	INTERESTRECEYOY               interface{} `json:"INTEREST_RECE_YOY"`
	INTERNALPAYABLEYOY            interface{} `json:"INTERNAL_PAYABLE_YOY"`
	INTERNALRECEYOY               interface{} `json:"INTERNAL_RECE_YOY"`
	INVENTORYYOY                  float64     `json:"INVENTORY_YOY"`
	INVESTREALESTATEYOY           float64     `json:"INVEST_REALESTATE_YOY"`
	LEASELIABYOY                  float64     `json:"LEASE_LIAB_YOY"`
	LENDFUNDYOY                   interface{} `json:"LEND_FUND_YOY"`
	LIABBALANCEYOY                interface{} `json:"LIAB_BALANCE_YOY"`
	LIABEQUITYBALANCEYOY          interface{} `json:"LIAB_EQUITY_BALANCE_YOY"`
	LIABEQUITYOTHERYOY            interface{} `json:"LIAB_EQUITY_OTHER_YOY"`
	LIABOTHERYOY                  interface{} `json:"LIAB_OTHER_YOY"`
	LOANADVANCEYOY                interface{} `json:"LOAN_ADVANCE_YOY"`
	LOANPBCYOY                    interface{} `json:"LOAN_PBC_YOY"`
	LONGEQUITYINVESTYOY           float64     `json:"LONG_EQUITY_INVEST_YOY"`
	LONGLOANYOY                   interface{} `json:"LONG_LOAN_YOY"`
	LONGPAYABLEYOY                interface{} `json:"LONG_PAYABLE_YOY"`
	LONGPREPAIDEXPENSEYOY         interface{} `json:"LONG_PREPAID_EXPENSE_YOY"`
	LONGRECEYOY                   interface{} `json:"LONG_RECE_YOY"`
	LONGSTAFFSALARYPAYABLEYOY     float64     `json:"LONG_STAFFSALARY_PAYABLE_YOY"`
	MINORITYEQUITYYOY             float64     `json:"MINORITY_EQUITY_YOY"`
	MONETARYFUNDSYOY              float64     `json:"MONETARYFUNDS_YOY"`
	NONCURRENTASSET1YEARYOY       interface{} `json:"NONCURRENT_ASSET_1YEAR_YOY"`
	NONCURRENTASSETBALANCEYOY     interface{} `json:"NONCURRENT_ASSET_BALANCE_YOY"`
	NONCURRENTASSETOTHERYOY       interface{} `json:"NONCURRENT_ASSET_OTHER_YOY"`
	NONCURRENTLIAB1YEARYOY        float64     `json:"NONCURRENT_LIAB_1YEAR_YOY"`
	NONCURRENTLIABBALANCEYOY      interface{} `json:"NONCURRENT_LIAB_BALANCE_YOY"`
	NONCURRENTLIABOTHERYOY        interface{} `json:"NONCURRENT_LIAB_OTHER_YOY"`
	NOTEACCOUNTSPAYABLEYOY        float64     `json:"NOTE_ACCOUNTS_PAYABLE_YOY"`
	NOTEACCOUNTSRECEYOY           float64     `json:"NOTE_ACCOUNTS_RECE_YOY"`
	NOTEPAYABLEYOY                float64     `json:"NOTE_PAYABLE_YOY"`
	NOTERECEYOY                   interface{} `json:"NOTE_RECE_YOY"`
	OILGASASSETYOY                interface{} `json:"OIL_GAS_ASSET_YOY"`
	OTHERCOMPREINCOMEYOY          float64     `json:"OTHER_COMPRE_INCOME_YOY"`
	OTHERCREDITORINVESTYOY        interface{} `json:"OTHER_CREDITOR_INVEST_YOY"`
	OTHERCURRENTASSETYOY          float64     `json:"OTHER_CURRENT_ASSET_YOY"`
	OTHERCURRENTLIABYOY           float64     `json:"OTHER_CURRENT_LIAB_YOY"`
	OTHEREQUITYINVESTYOY          float64     `json:"OTHER_EQUITY_INVEST_YOY"`
	OTHEREQUITYOTHERYOY           interface{} `json:"OTHER_EQUITY_OTHER_YOY"`
	OTHEREQUITYTOOLYOY            interface{} `json:"OTHER_EQUITY_TOOL_YOY"`
	OTHERNONCURRENTASSETYOY       float64     `json:"OTHER_NONCURRENT_ASSET_YOY"`
	OTHERNONCURRENTFINASSETYOY    interface{} `json:"OTHER_NONCURRENT_FINASSET_YOY"`
	OTHERNONCURRENTLIABYOY        interface{} `json:"OTHER_NONCURRENT_LIAB_YOY"`
	OTHERPAYABLEYOY               interface{} `json:"OTHER_PAYABLE_YOY"`
	OTHERRECEYOY                  interface{} `json:"OTHER_RECE_YOY"`
	PARENTEQUITYBALANCEYOY        interface{} `json:"PARENT_EQUITY_BALANCE_YOY"`
	PARENTEQUITYOTHERYOY          interface{} `json:"PARENT_EQUITY_OTHER_YOY"`
	PERPETUALBONDPAYBALEYOY       interface{} `json:"PERPETUAL_BOND_PAYBALE_YOY"`
	PERPETUALBONDYOY              interface{} `json:"PERPETUAL_BOND_YOY"`
	PREDICTCURRENTLIABYOY         interface{} `json:"PREDICT_CURRENT_LIAB_YOY"`
	PREDICTLIABYOY                interface{} `json:"PREDICT_LIAB_YOY"`
	PREFERREDSHARESPAYBALEYOY     interface{} `json:"PREFERRED_SHARES_PAYBALE_YOY"`
	PREFERREDSHARESYOY            interface{} `json:"PREFERRED_SHARES_YOY"`
	PREMIUMRECEYOY                interface{} `json:"PREMIUM_RECE_YOY"`
	PREPAYMENTYOY                 float64     `json:"PREPAYMENT_YOY"`
	PRODUCTIVEBIOLOGYASSETYOY     interface{} `json:"PRODUCTIVE_BIOLOGY_ASSET_YOY"`
	PROJECTMATERIALYOY            interface{} `json:"PROJECT_MATERIAL_YOY"`
	RCRESERVERECEYOY              interface{} `json:"RC_RESERVE_RECE_YOY"`
	REINSUREPAYABLEYOY            interface{} `json:"REINSURE_PAYABLE_YOY"`
	REINSURERECEYOY               interface{} `json:"REINSURE_RECE_YOY"`
	SELLREPOFINASSETYOY           interface{} `json:"SELL_REPO_FINASSET_YOY"`
	SETTLEEXCESSRESERVEYOY        interface{} `json:"SETTLE_EXCESS_RESERVE_YOY"`
	SHARECAPITALYOY               float64     `json:"SHARE_CAPITAL_YOY"`
	SHORTBONDPAYABLEYOY           interface{} `json:"SHORT_BOND_PAYABLE_YOY"`
	SHORTFINPAYABLEYOY            interface{} `json:"SHORT_FIN_PAYABLE_YOY"`
	SHORTLOANYOY                  interface{} `json:"SHORT_LOAN_YOY"`
	SPECIALPAYABLEYOY             interface{} `json:"SPECIAL_PAYABLE_YOY"`
	SPECIALRESERVEYOY             float64     `json:"SPECIAL_RESERVE_YOY"`
	STAFFSALARYPAYABLEYOY         float64     `json:"STAFF_SALARY_PAYABLE_YOY"`
	SUBSIDYRECEYOY                interface{} `json:"SUBSIDY_RECE_YOY"`
	SURPLUSRESERVEYOY             int         `json:"SURPLUS_RESERVE_YOY"`
	TAXPAYABLEYOY                 float64     `json:"TAX_PAYABLE_YOY"`
	TOTALASSETSYOY                float64     `json:"TOTAL_ASSETS_YOY"`
	TOTALCURRENTASSETSYOY         float64     `json:"TOTAL_CURRENT_ASSETS_YOY"`
	TOTALCURRENTLIABYOY           float64     `json:"TOTAL_CURRENT_LIAB_YOY"`
	TOTALEQUITYYOY                float64     `json:"TOTAL_EQUITY_YOY"`
	TOTALLIABEQUITYYOY            float64     `json:"TOTAL_LIAB_EQUITY_YOY"`
	TOTALLIABILITIESYOY           float64     `json:"TOTAL_LIABILITIES_YOY"`
	TOTALNONCURRENTASSETSYOY      float64     `json:"TOTAL_NONCURRENT_ASSETS_YOY"`
	TOTALNONCURRENTLIABYOY        float64     `json:"TOTAL_NONCURRENT_LIAB_YOY"`
	TOTALOTHERPAYABLEYOY          float64     `json:"TOTAL_OTHER_PAYABLE_YOY"`
	TOTALOTHERRECEYOY             float64     `json:"TOTAL_OTHER_RECE_YOY"`
	TOTALPARENTEQUITYYOY          float64     `json:"TOTAL_PARENT_EQUITY_YOY"`
	TRADEFINASSETNOTFVTPLYOY      float64     `json:"TRADE_FINASSET_NOTFVTPL_YOY"`
	TRADEFINASSETYOY              interface{} `json:"TRADE_FINASSET_YOY"`
	TRADEFINLIABNOTFVTPLYOY       interface{} `json:"TRADE_FINLIAB_NOTFVTPL_YOY"`
	TRADEFINLIABYOY               interface{} `json:"TRADE_FINLIAB_YOY"`
	TREASURYSHARESYOY             float64     `json:"TREASURY_SHARES_YOY"`
	UNASSIGNRPOFITYOY             float64     `json:"UNASSIGN_RPOFIT_YOY"`
	UNCONFIRMINVESTLOSSYOY        interface{} `json:"UNCONFIRM_INVEST_LOSS_YOY"`
	USERIGHTASSETYOY              float64     `json:"USERIGHT_ASSET_YOY"`
	OPINIONTYPE                   string      `json:"OPINION_TYPE"`
	OSOPINIONTYPE                 interface{} `json:"OSOPINION_TYPE"`
	LISTINGSTATE                  string      `json:"LISTING_STATE"`
}

type BalanceDataList []BalanceData

type RespBalanceData struct {
	Version string `json:"version"`
	Result  struct {
		Pages int             `json:"pages"`
		Data  BalanceDataList `json:"data"`
		Count int             `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e EastMoney) QueryBlanceData(ctx context.Context, secuCode string, date string) (BalanceDataList, error) {
	apiurl := "https://datacenter.eastmoney.com/securities/api/data/get"
	params := map[string]string{
		"source": "HSF10",
		"client": "APP",
		"type":   "RPT_F10_FINANCE_GBALANCE",
		"sty":    "F10_FINANCE_GBALANCE",
		"filter": fmt.Sprintf(`(SECUCODE="%s")(REPORT_DATE in (%s))`, strings.ToUpper(secuCode), date),
		"ps":     "10",
		"sr":     "-1",
		"st":     "REPORT_DATE",
	}
	logrus.Debug(ctx, "EastMoney QueryBlanceData "+apiurl+" begin", zap.Any("params", params))
	beginTime := time.Now()

	balaceResp := RespBalanceData{}

	resp, err := e.HTTPClient.R().SetQueryParams(params).Get(apiurl)

	latency := time.Now().Sub(beginTime).Milliseconds()
	logrus.Debug(
		ctx,
		"EastMoney QueryBlanceData "+apiurl+" end",
		zap.Int64("latency(ms)", latency),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("%s %#v", secuCode, resp)
	}

	_ = json.Unmarshal(resp.Body(), &balaceResp)

	return balaceResp.Result.Data, nil
}
