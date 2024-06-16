package cronjob

import (
	"capstone/constants"
	consultationEntities "capstone/entities/consultation"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"log"
	"time"
)

type Cronjob struct {
	consultationRepository consultationEntities.ConsultationRepository
	scheduler              gocron.Scheduler
}

func NewCronJob(scheduler gocron.Scheduler, consultationRepository consultationEntities.ConsultationRepository) *Cronjob {
	return &Cronjob{
		scheduler:              scheduler,
		consultationRepository: consultationRepository,
	}
}

func (c *Cronjob) InitCronJob() {
	job, err := c.scheduler.NewJob(
		gocron.DurationJob(10*time.Minute),
		gocron.NewTask(c.UpdateStatusConsultation),
	)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Cronjob Init, Job ID:", job.ID())
}

func (c *Cronjob) UpdateStatusConsultation() {
	fmt.Println("Cronjob Update Status Consultation Running")
	consultations := c.consultationRepository.GetAllConsultation()
	for _, consultation := range *consultations {
		if consultation.Status != constants.DONE && consultation.Status != constants.REJECTED {
			// Update Status Pending to Rejected When Date is After Now
			consultationDate := time.Date(consultation.Date.Year(), consultation.Date.Month(), consultation.Date.Day(), consultation.Time.Hour(), consultation.Time.Minute(), consultation.Time.Second(), 0, time.UTC)
			currentTime := time.Now().UTC()

			// Check if consultationDate is after or equal to currentTime
			if (consultationDate.Before(currentTime) || consultationDate.Equal(currentTime)) && consultation.Status == constants.PENDING {
				fmt.Println("Update Status Pending to Rejected")
				consultation.Status = "rejected"
				_, err := c.consultationRepository.UpdateStatusConsultation(&consultation)
				if err != nil {
					log.Println(err.Error())
				}
				continue
			}

			if (consultationDate.After(currentTime) || consultationDate.Equal(currentTime)) && consultation.Status == constants.INCOMING {
				fmt.Println("Update Status Incoming to Active")
				consultation.Status = "active"
				_, err := c.consultationRepository.UpdateStatusConsultation(&consultation)
				if err != nil {
					log.Println(err.Error())
				}
				continue
			}
		}

	}
}
