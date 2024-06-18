package notification

import "capstone/entities"

type NotificationRepository interface {
	GetNotificationByUserID(userID int, metadata *entities.Metadata) (*[]UserNotification, error)
	CreateUserNotification(notification *UserNotification) error
	DeleteUserNotification(notificationID int) error
	UpdateStatusNotification(notificationID int) error

	GetNotificationByDoctorID(doctorID int, metadata *entities.Metadata) (*[]DoctorNotification, error)
	CreateDoctorNotification(notification *DoctorNotification) error
	DeleteDoctorNotification(notificationID int) error
	UpdateStatusDoctorNotification(notificationID int) error
}

type NotificationUseCase interface {
	GetNotificationByUserID(userID int, metadata *entities.Metadata) (*[]UserNotification, error)
	CreateUserNotification(notification *UserNotification) error
	DeleteUserNotification(notificationID int) error
	UpdateStatusNotification(notificationID int) error

	GetNotificationByDoctorID(doctorID int, metadata *entities.Metadata) (*[]DoctorNotification, error)
	CreateDoctorNotification(notification *DoctorNotification) error
	DeleteDoctorNotification(notificationID int) error
	UpdateStatusDoctorNotification(notificationID int) error
}
