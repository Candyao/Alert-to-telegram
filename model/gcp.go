package model

import (
	"alert/request"
	"alert/util"
)

type GcpData struct {
	Notify *request.Notification
}

func (d *GcpData) AddGcpAlert () error {
	var timeInt=&util.SqlTime{}
	data := Alert{
		StartAt:   util.ToUnix(d.Notify.StartedAt),
		CreatedAt: util.ToUnix(timeInt.GetUTC()),
		UpdatedAt: util.ToUnix(timeInt.GetUTC()),
		AlertName: d.Notify.PolicyName,
		Resource:  d.Notify.ResourceName,
		Severity:  "warning",
		Summary:   d.Notify.Summary,
		Type:      "gcp",
	}
	if err := Alertdb.database.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
