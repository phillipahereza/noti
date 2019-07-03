package command

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/variadico/noti/service/notifyicon"
	"github.com/variadico/noti/service/speechsynthesizer"
)

func getBanner(title, message string, v *viper.Viper) notification {
	title, message = prefixSuffixTitleMessage(title, message, v)
	return &notifyicon.Notification{
		BalloonTipTitle: title,
		BalloonTipText:  message,
		BalloonTipIcon:  notifyicon.BalloonTipIconInfo,
	}
}

func getSpeech(title, message string, v *viper.Viper) notification {
	return &speechsynthesizer.Notification{
		Text:  fmt.Sprintf("%s %s", title, message),
		Rate:  3,
		Voice: v.GetString("speechsynthesizer.voice"),
	}
}
