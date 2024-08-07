package entity

type TimeSheet struct {
	Base
	ConfirmedManagerBy string
	ConfirmedBenefitBy string
	StatusTimeSheetID  string
	UserID             string
	TimeSheetDetails   []TimeSheetDetail
}
