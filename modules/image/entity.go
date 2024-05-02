package image

type Image struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Src          string `json:"src"`
	Complaint_id int    `json:"complaint_id"`
}
