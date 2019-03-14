package sender

import (
	"fmt"
	"time"

	"urlooker/alarm/cache"
	"urlooker/web/model"
)

func BuildMail(event *model.Event) string {
	strategy, _ := cache.StrategyMap.Get(event.StrategyId)
	respTime := fmt.Sprintf("%dms", event.RespTime)
	return fmt.Sprintf(
		"Status:%s\nUrl:%s\nIp:%s\nRespCode:%s\nRespTime:%s\nTimestamp:%s\nStep:%d\nNote:%s\n",
		event.Status,
		event.Url,
		event.Ip,
		event.RespCode,
		respTime,
		humanTime(event.EventTime),
		event.CurrentStep,
		strategy.Note,
	)
}

func BuildSms(event *model.Event) string {
	respTime := fmt.Sprintf("%dms", event.RespTime)
	return fmt.Sprintf(
		"[%s][%s %s][%s][%s][%s][O%d]",
		event.Status,
		showSubString(event.Url, 50),
		event.Ip,
		event.RespCode,
		respTime,
		humanTime(event.EventTime),
		event.CurrentStep,
	)
}

func humanTime(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func showSubString(str string, length int) string {
	runeStr := []rune(str)
	s := ""
	if length > len(runeStr) {
		length = len(runeStr)
	}

	for i := 0; i < length; i++ {
		s += string(runeStr[i])
	}
	return s
}
