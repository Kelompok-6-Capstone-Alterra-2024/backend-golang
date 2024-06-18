package notification

import (
	"capstone/entities/notification"
	"gorm.io/gorm"
)

type UserNotification struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
	IsRead  bool   `gorm:"column:is_read"`
}

type DoctorNotification struct {
	gorm.Model
	DoctorID uint   `gorm:"column:doctor_id"`
	Content  string `gorm:"column:content"`
	IsRead   bool   `gorm:"column:is_read"`
}

func (n *UserNotification) ToUserEntities() *notification.UserNotification {
	return &notification.UserNotification{
		ID:      n.ID,
		UserID:  n.UserID,
		Content: n.Content,
		IsRead:  n.IsRead,
	}
}

func (n *DoctorNotification) ToDoctorEntities() *notification.DoctorNotification {
	return &notification.DoctorNotification{
		ID:       n.ID,
		DoctorID: n.DoctorID,
		Content:  n.Content,
		IsRead:   n.IsRead,
	}
}

func ToNotificationUserModel(notification *notification.UserNotification) *UserNotification {
	return &UserNotification{
		UserID:  notification.UserID,
		Content: notification.Content,
		IsRead:  notification.IsRead,
	}
}

func ToNotificationDoctorModel(notification *notification.DoctorNotification) *DoctorNotification {
	return &DoctorNotification{
		DoctorID: notification.DoctorID,
		Content:  notification.Content,
		IsRead:   notification.IsRead,
	}
}
