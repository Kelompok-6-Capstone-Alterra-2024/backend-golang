package complaint

import "capstone/repositories/mysql/doctor"

type Complaint struct {
	ID       int           `gorm:"primaryKey:autoIncrement"`
	DoctorID int           `gorm:"type:int"`
	Doctor   doctor.Doctor `gorm:"foreignKey:doctor_id;references:id"`
	UserID   int           `gorm:"type:int"`
	Content  string        `gorm:"type:text"`
}
