package common

import (
	"fmt"

	"github.com/Sirupsen/logrus"

	"gitlab.com/gitlab-org/gitlab-runner/helpers"
	"gitlab.com/gitlab-org/gitlab-runner/helpers/url"
)

type BuildLogger struct {
	log   JobTrace
	entry *logrus.Entry
}

func (e *BuildLogger) SendRawLog(args ...interface{}) {
	if e.log != nil {
		fmt.Fprint(e.log, args...)
	}
}

func (e *BuildLogger) sendLog(logger func(args ...interface{}), logPrefix string, args ...interface{}) {
	if e.log != nil {
		logLine := url_helpers.ScrubSecrets(logPrefix + fmt.Sprintln(args...))
		e.SendRawLog(logLine)
		e.SendRawLog(helpers.ANSI_RESET)

		if e.log.IsStdout() {
			return
		}
	}

	if len(args) == 0 {
		return
	}

	logger(args...)
}

func (e *BuildLogger) Debugln(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.entry.Debugln(args...)
}

func (e *BuildLogger) Println(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.sendLog(e.entry.Debugln, helpers.ANSI_CLEAR, args...)
}

func (e *BuildLogger) Infoln(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.sendLog(e.entry.Println, helpers.ANSI_BOLD_GREEN, args...)
}

func (e *BuildLogger) Warningln(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.sendLog(e.entry.Warningln, helpers.ANSI_YELLOW+"WARNING: ", args...)
}

func (e *BuildLogger) SoftErrorln(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.sendLog(e.entry.Warningln, helpers.ANSI_BOLD_RED+"ERROR: ", args...)
}

func (e *BuildLogger) Errorln(args ...interface{}) {
	if e.entry == nil {
		return
	}
	e.sendLog(e.entry.Errorln, helpers.ANSI_BOLD_RED+"ERROR: ", args...)
}

func NewBuildLogger(log JobTrace, entry *logrus.Entry) BuildLogger {
	return BuildLogger{
		log:   log,
		entry: entry,
	}
}
