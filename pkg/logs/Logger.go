package logs

import (
	"log"
	"os"
)

const (
	perm = 0666
	flag = os.O_CREATE|os.O_APPEND|os.O_RDONLY
	file ="../pkg/logs/logs.txt"
)


type Logger struct{
	ErrorLogger *log.Logger
	WarningLogger *log.Logger
	InfoLogger *log.Logger
	file *os.File
}

func NewLogger()*Logger{
	file ,err := os.OpenFile(file,flag,perm)
	if err!=nil {
		log.Fatal(err)
	}
	return &Logger{
		ErrorLogger: nil,
		WarningLogger: nil,
		InfoLogger: nil,
		file: file,
	}
}

func (l *Logger) ErrorLoggerFunc(msg string) {
	l.ErrorLogger = log.New(NewLogger().file,"[ERROR] ",log.Ldate|log.Ltime|log.Lshortfile)
	l.ErrorLogger.Println(msg)
}

func (l *Logger) WarningLoggerFunc(msg string) {
	l.WarningLogger = log.New(NewLogger().file,"[WARNING] ",log.Ldate|log.Ltime|log.Lshortfile)
	l.WarningLogger.Println(msg)
}

func (l *Logger) InfoLoggerFunc(msg string) {
	l.InfoLogger = log.New(NewLogger().file,"[INFO] ",log.Ldate|log.Ltime|log.Lshortfile)
	l.InfoLogger.Println(msg)
}
