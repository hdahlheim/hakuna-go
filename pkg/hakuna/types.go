/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package hakuna

import "time"

type Absence struct {
	ID                   int         `json:"id"`
	StartDate            time.Time   `json:"start_date"`
	EndDate              time.Time   `json:"end_date"`
	FirstHalfDay         bool        `json:"first_half_day"`
	SecondHalfDay        bool        `json:"second_half_day"`
	IsRecurring          bool        `json:"is_recurring"`
	WeeklyRepeatInterval int         `json:"weekly_repeat_interval"`
	User                 User        `json:"user"`
	AbsenceType          AbsenceType `json:"absence_type"`
}

type AbsenceType struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	GrantsWorkTime bool   `json:"grants_work_time"`
	IsVacation     bool   `json:"is_vacation"`
	Archived       bool   `json:"archived"`
}

type Company struct {
	CompanyName            string `json:"company_name"`
	DurationFormat         string `json:"duration_format"`
	AbsenceRequestsEnabled bool   `json:"absence_requests_enabled"`
	ProjectsEnabled        bool   `json:"projects_enabled"`
	GroupsEnabled          bool   `json:"groups_enabled"`
}

type Pong struct {
	Pong time.Time `json:"pong"`
}

type Project struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
}

type Vacation struct {
	RedeemedDays  float64 `json:"redeemed_days"`
	RemainingDays float64 `json:"remaining_days"`
}

type Overview struct {
	Overtime          string   `json:"overtime"`
	OvertimeInSeconds float64  `json:"overtime_in_seconds"`
	Vacation          Vacation `json:"vacation"`
}

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
}

type TimeEntry struct {
	ID                int     `json:"id"`
	Note              string  `json:"note"`
	Date              string  `json:"date"`
	Duration          string  `json:"duration"`
	DurationInSeconds float64 `json:"duration_in_seconds"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	User              User    `json:"user"`
	Task              Task    `json:"task"`
	Project           Project `json:"project"`
}

type Timer struct {
	Note              string  `json:"note"`
	Date              string  `json:"date"`
	Duration          string  `json:"duration"`
	DurationInSeconds float64 `json:"duration_in_seconds"`
	StartTime         string  `json:"start_time"`
	User              User    `json:"user"`
	Task              Task    `json:"task"`
	Project           Project `json:"project"`
}

type User struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Groups []string `json:"groups"`
}

type StartTimerReq struct {
	TaskId    int    `json:"task_id"`
	StartTime string `json:"start_time,omitempty"`
	ProjectId int    `json:"project_id,omitempty"`
	Note      string `json:"note,omitempty"`
}

type StopTimerReq struct {
	EndTime string `json:"end_time"`
}

type ResponeError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CreatTimeEntryReq struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Note      string `json:"note"`
	ProjectID int    `json:"project_id"`
	TaskID    int    `json:"task_id"`
}

// type ListTimeEntriesReq struct {
// 	StartDate string `json:"start_date"`
// 	EndDate   string `json:"end_date"`
// }
