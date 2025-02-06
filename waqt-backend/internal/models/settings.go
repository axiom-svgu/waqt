package models

import "time"

type Setting struct {
	ThemePreference             string    `json:"theme_preference"`
	TimeZone                    string    `json:"time_zone"`
	Language                    string    `json:"language"`
	DefaultTaskDuration         time.Time `json:"default_task_duration"`
	TaskReminderIntervalMinutes int       `json:"task_reminder_interval_minutes"`

	// Preferences
	WorkStartTime string `json:"work_start_time" gorm:"default:'09:00'"`
	WorkEndTime   string `json:"work_end_time" gorm:"default:'17:00'"`
	WorkDays      string `json:"work_days" gorm:"default:'1,2,3,4,5'"` // 0=Sunday, 1=Monday, etc.

}
