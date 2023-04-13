package middleware

import (
	"bytes"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/constants"
	"encoding/json"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLoggerActivity func for logging
func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		writeLogReq(c)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		dur := time.Since(t)
		c.Set("Latency", strconv.Itoa(int(dur.Milliseconds())))
		writeLogResp(c, blw.body.String())
	}
}

func writeLogReq(c *gin.Context) {
	header := make(map[string][]string)
	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		re := regexp.MustCompile(`\r?\n`)
		var request = re.ReplaceAllString(readBody(rdr1), "")
		var bindReq interface{}
		err := json.Unmarshal([]byte(request), &bindReq)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}

		logger.Log.WithFields(logrus.Fields{
			constants.Url:           c.Request.URL.Path,
			constants.Method:        c.Request.Method,
			constants.RequestId:     requestid.Get(c),
			constants.RequestBody:   bindReq,
			constants.RequestHeader: header,
		}).Info(constants.Request)
		c.Request.Body = rdr2
	} else {
		if c.FullPath() != "/" {
			logger.Log.WithFields(logrus.Fields{
				constants.Url:           c.Request.URL.Path,
				constants.Method:        c.Request.Method,
				constants.RequestId:     requestid.Get(c),
				constants.UserAgent:     c.Request.UserAgent(),
				constants.RequestHeader: header,
			}).Info(constants.Request)
		}
	}
}

func writeLogResp(c *gin.Context, resp string) {
	var bindResp interface{}
	latency, _ := c.Get("Latency")
	rc, _ := c.Get("RespCode")
	if rc == nil {
		rc = "--"
	}
	err := json.Unmarshal([]byte(resp), &bindResp)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}

	if c.FullPath() != "/" {
		logger.Log.WithFields(logrus.Fields{
			constants.Url:          c.Request.URL.Path,
			constants.ResponseCode: rc.(string),
			constants.Method:       c.Request.Method,
			constants.RequestId:    requestid.Get(c),
			constants.UserAgent:    c.Request.UserAgent(),
			constants.Latency:      latency.(string),
			constants.ResponseBody: bindResp,
		}).Info(constants.Response)
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func SetupLogger() {
	logger.Log.Println("Setup Logger Start")
	lumberjackLogger := &lumberjack.Logger{
		// Log file absolute path, os agnostic
		Filename:   filepath.ToSlash("./log/log"),
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30,   // days
		Compress:   true, // disabled by default
	}
	multiWriter := io.MultiWriter(lumberjackLogger, os.Stderr)
	Formatter := new(logrus.JSONFormatter)
	//You can change the Timestamp format. But you have to use the same date and time.
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	logger.Log.SetFormatter(Formatter)
	logger.Log.SetOutput(multiWriter)
	logger.Log.Println("Setup Logger Finish")
}
