package complaint

import (
	consultationModel "capstone/repositories/mysql/consultation"
	userModel "capstone/repositories/mysql/user"
	"gorm.io/gorm"
)

type Complaint struct {
	gorm.Model
	ConsultationID     uint                           `gorm:"type:int;column:complaint_id"`
	Consultation       consultationModel.Consultation `gorm:"foreignKey:ConsultationID"`
	UserID             int                            `gorm:"type:int;column:user_id"`
	User               userModel.User                 `gorm:"foreignKey:UserID"`
	Name               string                         `gorm:"type:varchar(255);column:name"`
	Age                int                            `gorm:"type:int;column:age"`
	Gender             string                         `gorm:"type:Enum('pria', 'wanita')"`
	Message            string                         `gorm:"type:text;column:message"`
	MedicalHistory     string                         `gorm:"type:text;column:medical_history"`
	DoctorNotification string                         `gorm:"type:varchar(255);column:doctor_notification"`
}
