package entity


type TimeSheet struct {
	Base
	ConfirmedManagerBy string            `json:"id_manager"`
	ConfirmedBenefitBy string            `json:"id_benefit"`
	StatusTimeSheetID  string            `json:"status_time_sheet_id"`
	UserID             string            `json:"user_id"`
	TimeSheetDetails   []TimeSheetDetail `json:"time_sheet_details"`
}
