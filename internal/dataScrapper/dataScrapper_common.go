package datascrapper

type UptimeInfo struct {
	IdleTime  float64
	AllTime   float64
	TotalTime float64
}

func (u *UptimeInfo) CalculateTotalTime() {
	u.TotalTime = u.AllTime + u.IdleTime
}

type RAMInfo struct {
	Free  int
	Total int
}

type CPUInfo struct {
	TotalLoad float64
}
