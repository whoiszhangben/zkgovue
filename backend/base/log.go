package base

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type BaseLog struct {
	Path      string
	Prefix    string
	IsOpen    int
	Index     int
	LogLevel  int
	TermLevel int
	LogTime   time.Time
	LogFile   *os.File
	FileName  string
	MaxLine   int
	LineCnt   int
}

var Logs *BaseLog

func Init_log() error {
	Logs = NewBaseLog()
	Logs.Path = beego.AppConfig.String("log_path")
	fmt.Println(Logs.Path)
	Logs.Prefix = beego.AppConfig.String("log_head")
	fmt.Println(Logs.Prefix)
	i, _ := beego.AppConfig.Int("log_level")
	Logs.LogLevel = i
	i, _ = beego.AppConfig.Int("term_level")
	Logs.TermLevel = i
	Logs.MaxLine, _ = beego.AppConfig.Int("logfile_max_line")
	//defer Logs.CloseLog()
	return Logs.CreateLog()
}

func NewBaseLog() *BaseLog {
	log := new(BaseLog)
	log.IsOpen = 0
	log.Index = 1
	return log
}

func (l *BaseLog) CreateLog() error {
	if l.IsOpen != 0 {
		l.CloseLog()
	}
	l.LogTime = time.Now()
	lt := l.LogTime
	filename := fmt.Sprintf("%s/%s_%d_%d%02d%02d_%d.log", l.Path, l.Prefix, os.Getpid(), lt.Year(), lt.Month(), lt.Day(), l.Index)
	l.FileName = filename
	tmpname := fmt.Sprintf("%s.tmp", filename)
	//fmt.Printf("filename:%s\n", filename)

	logfile, err := os.OpenFile(tmpname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open log %s failed %s\n", tmpname, err.Error())
		return err
	}

	l.LogFile = logfile
	l.Index++
	l.IsOpen = 1

	return nil
}

func (l *BaseLog) FileSwitch() {
	if l.LineCnt > l.MaxLine {
		l.CloseLog()
		l.CreateLog()
		l.LineCnt = 0
	}
	nt := time.Now()
	lt := l.LogTime
	if nt.Year() != lt.Year() || nt.Month() != lt.Month() || nt.Day() != lt.Day() {
		l.CloseLog()
		l.CreateLog()
		l.LineCnt = 0
	}
}

func (l *BaseLog) CloseLog() {
	file := l.LogFile
	oldname := file.Name()

	file.Close()

	os.Rename(oldname, l.FileName)
}

func (l *BaseLog) GetTimeString() string {
	t := time.Now()
	s := fmt.Sprintf("%d-%d-%d %02d:%02d:%02d.%06d", t.Year(), t.Month(),
		t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000)
	return s
}

func (l *BaseLog) Info(format string, a ...interface{}) {
	if l.LogLevel > 40 && l.TermLevel > 40 {
		return
	}
	l.FileSwitch()
	l.LineCnt++
	info := fmt.Sprintf(format, a...)
	log := fmt.Sprintf("[I] %s %s", l.GetTimeString(), info)
	if l.LogLevel <= 40 {
		l.LogFile.WriteString(log)
	}
	if l.TermLevel <= 40 {
		fmt.Printf("%s", log)
	}
}

func (l *BaseLog) Debug(format string, a ...interface{}) {
	if l.LogLevel > 0 && l.TermLevel > 0 {
		return
	}
	l.FileSwitch()
	l.LineCnt++
	info := fmt.Sprintf(format, a...)
	log := fmt.Sprintf("[D] %s %s", l.GetTimeString(), info)
	if l.LogLevel <= 0 {
		l.LogFile.WriteString(log)
	}
	if l.TermLevel <= 0 {
		fmt.Printf("%s", log)
	}
}

func (l *BaseLog) Error(format string, a ...interface{}) {
	if l.LogLevel > 80 && l.TermLevel > 80 {
		return
	}
	l.FileSwitch()
	l.LineCnt++
	info := fmt.Sprintf(format, a...)
	log := fmt.Sprintf("[E] %s %s", l.GetTimeString(), info)
	if l.LogLevel <= 80 {
		l.LogFile.WriteString(log)
	}
	if l.TermLevel <= 80 {
		fmt.Printf("%s", log)
	}
}
