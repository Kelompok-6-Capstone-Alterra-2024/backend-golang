package notification

import (
	"capstone/entities"
	notificationEntities "capstone/entities/notification"
)

type NotificationUseCase struct {
	notificationRepository notificationEntities.NotificationRepository
}

func NewNotificationUseCase(notificationRepository notificationEntities.NotificationRepository) notificationEntities.NotificationUseCase {
	return &NotificationUseCase{notificationRepository}
}

func (usecase *NotificationUseCase) GetNotificationByUserID(userID int, metadata *entities.Metadata) (*[]notificationEntities.UserNotification, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *NotificationUseCase) CreateUserNotification(notification *notificationEntities.UserNotification) error {
	if err := usecase.notificationRepository.CreateUserNotification(notification); err != nil {
		return err
	}
	return nil
}

func (usecase *NotificationUseCase) DeleteUserNotification(notificationID int) error {
	//TODO implement me
	panic("implement me")
}

func (usecase *NotificationUseCase) UpdateStatusNotification(notificationID int) error {
	//TODO implement me
	panic("implement me")
}

func (usecase *NotificationUseCase) GetNotificationByDoctorID(doctorID int, metadata *entities.Metadata) (*[]notificationEntities.DoctorNotification, error) {
	notifications, err := usecase.notificationRepository.GetNotificationByDoctorID(doctorID, metadata)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (usecase *NotificationUseCase) CreateDoctorNotification(notification *notificationEntities.DoctorNotification) error {
	if err := usecase.notificationRepository.CreateDoctorNotification(notification); err != nil {
		return err
	}
	return nil
}

func (usecase *NotificationUseCase) DeleteDoctorNotification(notificationID int) error {
	if err := usecase.notificationRepository.DeleteDoctorNotification(notificationID); err != nil {
		return err
	}
	return nil
}

func (usecase *NotificationUseCase) UpdateStatusDoctorNotification(notificationID int) error {
	if err := usecase.notificationRepository.UpdateStatusDoctorNotification(notificationID); err != nil {
		return err
	}
	return nil
}
