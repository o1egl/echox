package log

import (
	"bytes"
	"github.com/Sirupsen/logrus"
	elog "github.com/labstack/echo/log"
	glog "github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

var u = make(glog.JSON)
var ustr = "{\\\"age\\\":25,\\\"name\\\":\\\"Joe\\\"}"

func init() {
	u["name"] = "Joe"
	u["age"] = 25
}

func TestLogrusInit(t *testing.T) {
	log := Logrus()
	assert.Implements(t, (*elog.Logger)(nil), log)

	ls := logrus.New()
	ls.Level = logrus.FatalLevel
	log = LogrusFromLogger(ls)
	assert.Equal(t, logrus.FatalLevel, log.logger.Level)
}

func TestLogrus(t *testing.T) {

	log := Logrus()

	lvls := map[glog.Lvl]logrus.Level{
		glog.DEBUG: logrus.DebugLevel,
		glog.INFO:  logrus.InfoLevel,
		glog.WARN:  logrus.WarnLevel,
		glog.ERROR: logrus.ErrorLevel,
		glog.FATAL: logrus.FatalLevel,
	}

	for gl, ll := range lvls {
		log.SetLevel(gl)
		assert.Equal(t, ll, log.logger.Level)
	}

	log.SetLevel(glog.DEBUG)

	switch os.Getenv("TEST_LOGGER_FATAL") {
	case "fatal":
		log.Fatal("Fatal")
		return
	case "fatalf":
		log.Fatalf("Fatal%s", "f")
		return
	case "fatalj":
		log.Fatalj(u)
		return
	case "":

	default:
		t.Fatalf("Incorrect TEST_LOGGER_FATAL param %s", os.Getenv("TEST_LOGGER_FATAL"))

	}

	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	log.Print("Print")
	testLogMsg(t, "Print", buf)
	log.Printf("Print%s", "f")
	testLogMsg(t, "Printf", buf)
	log.Printj(u)
	testLogMsg(t, ustr, buf)

	log.Debug("Debug")
	testLogMsg(t, "Debug", buf)
	log.Debugf("Debug%s", "f")
	testLogMsg(t, "Debugf", buf)
	log.Debugj(u)
	testLogMsg(t, ustr, buf)

	log.Info("Info")
	testLogMsg(t, "Info", buf)
	log.Infof("Info%s", "f")
	testLogMsg(t, "Infof", buf)
	log.Infoj(u)
	testLogMsg(t, ustr, buf)

	log.Warn("Warn")
	testLogMsg(t, "Warn", buf)
	log.Warnf("Warn%s", "f")
	testLogMsg(t, "Warnf", buf)
	log.Warnj(u)
	testLogMsg(t, ustr, buf)

	log.Error("Error")
	testLogMsg(t, "Error", buf)
	log.Errorf("Error%s", "f")
	testLogMsg(t, "Errorf", buf)
	log.Errorj(u)
	testLogMsg(t, ustr, buf)

	log.Error("Error")
	testLogMsg(t, "Error", buf)
	log.Errorf("Error%s", "f")
	testLogMsg(t, "Errorf", buf)
	log.Errorj(u)
	testLogMsg(t, ustr, buf)

	fatalLogrusTest(t, "fatal", "Fatal")
	fatalLogrusTest(t, "fatalf", "Fatalf")
	fatalLogrusTest(t, "fatalj", ustr)

}

func TestLogrusOff(t *testing.T) {
	log := Logrus()
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetLevel(glog.OFF)

	log.Print("Print")
	log.Printf("Print%s", "f")
	log.Printj(u)

	log.Debug("Debug")
	log.Debugf("Debug%s", "f")
	log.Debugj(u)

	log.Info("Info")
	log.Infof("Info%s", "f")
	log.Infoj(u)

	log.Warn("Warn")
	log.Warnf("Warn%s", "f")
	log.Warnj(u)

	log.Error("Error")
	log.Errorf("Error%s", "f")
	log.Errorj(u)

	log.Error("Error")
	log.Errorf("Error%s", "f")
	log.Errorj(u)

	assert.Equal(t, buf.String(), "")

	fatalLogrusTest(t, "fatal", "")
	fatalLogrusTest(t, "fatalf", "")
	fatalLogrusTest(t, "fatalj", "")
}

func fatalLogrusTest(t *testing.T, env string, contains string) {
	buf := new(bytes.Buffer)
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus")
	cmd.Env = append(os.Environ(), "TEST_LOGGER_FATAL="+env)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		assert.Contains(t, buf.String(), contains)
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func testLogMsg(t *testing.T, msg string, buf *bytes.Buffer) {
	assert.Contains(t, buf.String(), msg)
	buf.Reset()
}
