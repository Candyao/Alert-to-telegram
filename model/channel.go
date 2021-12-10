package model

func GetIDByService(s string) int {
	var channel Channel
	Alertdb.database.Select("id").Where("service=?",s).First(&channel)
	return channel.ID
}

func GetChatIDByService(s string) string  {
	var channel Channel
	Alertdb.database.Select("chart_id").Where("service=?",s).First(&channel)
	return channel.ChartID
}

func (*Channel) TableName() string  {
	return "channel"
}