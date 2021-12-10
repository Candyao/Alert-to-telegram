package model

import (
	"alert/request"
	"alert/util"
)

func AddPromAlert(this *request.Alerts) error {
	data:=RangePromAlert(this)
	if err := Alertdb.database.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func RangePromAlert(alerts *request.Alerts) []Alert {
	var alertList=alerts.Alerts
	var data []Alert
	var timeInt=&util.SqlTime{}

	for _,alert:=range alertList{
		summary, _ :=alert.Annotations["summary"].(string)
		severity,_:=alert.Labels["severity"].(string)
		desc,_:=alert.Annotations["description"].(string)

		d:=Alert{
			StartAt:     util.ToUnix(alert.StartsAt),
			CreatedAt:   util.ToUnix(timeInt.GetUTC()),
			UpdatedAt:   util.ToUnix(timeInt.GetUTC()),
			AlertName:   alert.Labels["alertname"].(string),
			ChannelID:   getService(alert.Labels["service"]),
			Severity:    severity,
			Summary:     summary,
			Description: desc,
			Type:        "prom",
		}
		data=append(data,d)
	}
	return data
}

func getService(s interface{}) (id int) {
	value,ok:=s.(string)
	if ok {
		id=GetIDByService(value)
		return
	}
	return
}

func (*Alert) TableName() string {
	return "alert"
}