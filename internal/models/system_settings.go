package models

type SystemSettings struct {
	ID              uint   `gorm:"primaryKey" json:"-"` // Singleton ID=1
	SiteName        string `gorm:"default:'Ayee Portal'" json:"siteName"`
	MaintenanceMode bool   `gorm:"default:false" json:"maintenanceMode"`
	DefaultTheme    string `gorm:"default:'light'" json:"defaultTheme"` // 'light' | 'dark'
}
