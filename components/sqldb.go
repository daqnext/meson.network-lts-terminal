package components

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/universe-30/ULog"
	"gorm.io/gorm/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

/*
db_host
db_port
db_name
db_username
db_password
*/
func NewDB() (*gorm.DB, *sql.DB, error) {

	db_host, db_host_err := basic.Config.GetString("db_host", "127.0.0.1")
	if db_host_err != nil {
		return nil, nil, errors.New("db_host [string] in config err," + db_host_err.Error())
	}

	db_port, db_port_err := basic.Config.GetInt("db_port", 3306)
	if db_port_err != nil {
		return nil, nil, errors.New("db_port [int] in config err," + db_port_err.Error())
	}

	db_name, db_name_err := basic.Config.GetString("db_name", "dbname")
	if db_name_err != nil {
		return nil, nil, errors.New("db_name [string] in config err," + db_name_err.Error())
	}

	db_username, db_username_err := basic.Config.GetString("db_username", "username")
	if db_username_err != nil {
		return nil, nil, errors.New("db_username [string] in config err," + db_username_err.Error())
	}

	db_password, db_password_err := basic.Config.GetString("db_password", "password")
	if db_password_err != nil {
		return nil, nil, errors.New("db_password [string] in config err," + db_password_err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	var GormDB *gorm.DB
	var errOpen error

	GormDB, errOpen = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: New_gormLocalLogger(basic.Logger),
	})

	if errOpen != nil {
		return nil, nil, errOpen
	}

	sqlDB, errsql := GormDB.DB()
	if errsql != nil {
		return nil, nil, errsql
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return GormDB, sqlDB, nil

}

///////////////////////////

type gormLocalLogger struct {
	LocalLogger           ULog.Logger
	SlowThreshold         time.Duration
	SkipErrRecordNotFound bool
}

func New_gormLocalLogger(localLogger ULog.Logger) *gormLocalLogger {
	return &gormLocalLogger{
		LocalLogger:           localLogger,
		SlowThreshold:         200 * time.Millisecond,
		SkipErrRecordNotFound: true,
	}
}

func (l *gormLocalLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLocalLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	//when no err
	if err == nil {
		//return when only interested in Error,Fatal,Panic
		if l.LocalLogger.GetLevel() < ULog.WarnLevel {
			return
		}
	}

	elapsed := time.Since(begin)
	elapsedStr := fmt.Sprintf("%fms", float64(elapsed.Nanoseconds())/1e6)
	if err == nil && l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		//slow log
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		l.LocalLogger.Warnln(utils.FileWithLineNum(), slowLog, elapsedStr, "-", sql)
		return
	}

	///errors , when error happens logs it at any loglevel
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		sql, rows := fc()
		if rows == -1 {
			l.LocalLogger.Errorln(utils.FileWithLineNum(), "err:", err, elapsedStr, "-", sql)
		} else {
			l.LocalLogger.Errorln(utils.FileWithLineNum(), "err:", err, elapsedStr, rows, sql)
		}
		return
	}

	//info
	if l.LocalLogger.GetLevel() == ULog.TraceLevel {
		sql, rows := fc()
		if rows == -1 {
			l.LocalLogger.Traceln(utils.FileWithLineNum(), "err:", err, elapsedStr, "-", sql)
		} else {
			l.LocalLogger.Traceln(utils.FileWithLineNum(), "err:", err, elapsedStr, rows, sql)
		}
		return
	}
}

func (l *gormLocalLogger) Info(ctx context.Context, s string, args ...interface{}) {
	//not use
	//l.LocalLogger.Infoln(s, append([]interface{}{utils.FileWithLineNum()}, args...))
}

func (l *gormLocalLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	//not use
	//l.LocalLogger.Warnln(s, append([]interface{}{utils.FileWithLineNum()}, args...))
}

func (l *gormLocalLogger) Error(ctx context.Context, s string, args ...interface{}) {
	//not use
	//l.LocalLogger.Errorln(s, append([]interface{}{utils.FileWithLineNum()}, args...))
}
