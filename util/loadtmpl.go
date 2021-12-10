package util

import (
	"alert/config"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"
	"time"
)

func Str_FormatDate(toformat interface{}) string{
	t:=ToUnix(toformat)
	return t.Format(config.Configmanager.Config.TimeOutdata)
}

func ToUnix(toformat interface{}) time.Time {
	var t time.Time
	var err error
	value,ok:=toformat.(string)
	Outdata:=config.Configmanager.Config.TimeOutdata
	loc, _ := time.LoadLocation(config.Configmanager.Config.TimeZone)
	if !ok{
		t,err=time.ParseInLocation(Outdata,time.Unix(toformat.(int64),0).Format(Outdata),time.UTC)
	}else {
		t, err=time.Parse(time.RFC3339Nano, value)
	}
	if err != nil {
		log.Println(err)
	}
	return t.In(loc)
}

type SqlTime struct {}

func (*SqlTime) GetUTC() string {
	t:=time.Now()
	return t.In(time.UTC).Format(time.RFC3339Nano)
}

var funcMap=template.FuncMap{
	"str_FormatDate": Str_FormatDate,
	"str_UpperCase":  strings.ToUpper,
}

func LoadTemplate(tmplfile string) *template.Template {
	tmpH, err := template.New(tmplfile).Funcs(funcMap).ParseFiles(fmt.Sprintf("%s/%s",
		                                                config.Configmanager.Config.Templatepath,tmplfile))
	if err != nil {
		log.Fatalf("Problem reading parsing template file: %v", err)
	} else {
		log.Printf("Load template file:%s", tmplfile)
	}
	return tmpH
}

func AlertFormatTemplate(notification interface{},tmpH *template.Template) string {
	var bytesBuff bytes.Buffer
	var err error
	writer := io.Writer(&bytesBuff)

	err = tmpH.Execute(writer, notification)

	if err != nil {
		log.Fatalf("Problem with template execution: %v", err)
	}
	return bytesBuff.String()
}
