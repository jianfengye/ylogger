package ylogger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewYLogger(os.Stdout)
	logger.Debug("class1", "test", "debug", []string{"string1", "string2"}, logger)
	logger.Debug("class1", "testesfdasfdsa")
	logger.Warning("class1", "test")
	logger.Info("class1", "test")
	logger.Trace("class1", "test")
	logger.Error("class1", "test")

	logger.Disable("debug")
	logger.Debug("class2", "test", "debug", []string{"string1", "string2"}, logger)
	logger.Enable("debug")
	logger.Debug("class3", "test", "debug", []string{"string1", "string2"}, logger)

	Debug("class4", "test", "debug", []string{"string1", "string2"}, logger)
}
