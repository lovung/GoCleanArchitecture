package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Parallel()
	t.Run("Test logger", func(t *testing.T) {
		Init(true)
		SetLevel("debug")
		Debug("debug")
		Debugf("debugf %v", 1)
		Printf("printf %v", 1)
		Info("info")
		Infof("infof %v", 1)
		Warn("warn")
		Warnf("warnf %v", 1)
		Error("error")
		Errorf("errorf %v", 1)
		SetLevel("info")
		SetLevel("warn")
		SetLevel("error")
		SetLevel("default")
	})
}

func TestInitProduction(t *testing.T) {
	t.Parallel()
	t.Run("Test production", func(t *testing.T) {
		Init(false)
		Instance()
	})
}
