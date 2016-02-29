package iDhelper //import "go.iondynamics.net/iDhelper"

import (
	"fmt"
	"os"
	"time"

	idl "go.iondynamics.net/iDlogger"
	"go.iondynamics.net/iDlogger/priority"
	"go.iondynamics.net/iDslackLog"
)

func LoggerQuickSlack(prefix, prioThreshold, slackurl string) {
	if slackurl != "" {

		idl.AddHook(&iDslackLog.SlackLogHook{
			AcceptedPriorities: priority.Threshold(priority.Atos(prioThreshold)),
			HookURL:            slackurl,
			IconURL:            "",
			Channel:            "",
			IconEmoji:          "",
			Username:           prefix + " Log",
		})
	}

	idl.StandardLogger().Async = true
	idl.SetPrefix(prefix)
	idl.SetErrCallback(func(err error) {
		fmt.Fprintln(os.Stderr, err)
		panic("Logger caught an internal error")
	})
}
