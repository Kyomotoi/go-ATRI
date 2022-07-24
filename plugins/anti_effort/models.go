package antieffort

import "time"

type AntiEffortModel struct {
	UpdateTime string                `json:"update_time"`
	Data       []AntiEffortUserModel `json:"data"`
}

type AntiEffortUserModel struct {
	UserID             int64   `json:"user_id"`
	UserNickname       string  `json:"user_nickname"`
	WakaUserID         string  `json:"w_user_id"`
	WakaURL            string  `json:"waka_url"`
	LastSevenDaysCount float64 `json:"last_7_days_count"`
	RecentCount        float64 `json:"recent_count"`
}

type AntiEffortUpdateDataMap struct {
	LastSevenDaysCount float64 `json:"last_7_days_count"`
	RecentCount        float64 `json:"recent_count"`
}

type WakatimeShareEmbadData struct {
	Data []struct {
		GrandTotal struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"grand_total"`
		Range struct {
			Date     string    `json:"date"`
			End      time.Time `json:"end"`
			Start    time.Time `json:"start"`
			Text     string    `json:"text"`
			Timezone string    `json:"timezone"`
		} `json:"range"`
	} `json:"data"`
}
