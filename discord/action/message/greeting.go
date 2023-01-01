package message

import (
	"math/rand"
	"nonnonstop/akariai/discord/action"
	"time"
)

type ActionGreeting struct {
	Interval  time.Duration
	Messages  []string
	Interrupt bool
	lastTime  map[string]time.Time
}

func (a *ActionGreeting) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	channelID := p.ChannelID
	if !p.IsMentionToMe && p.GuildID != "" {
		now := time.Now()
		if a.lastTime[channelID].Add(a.Interval).After(now) {
			return a.Interrupt
		}
		if a.lastTime == nil {
			a.lastTime = make(map[string]time.Time)
		}
		a.lastTime[channelID] = now
	}

	messages := a.Messages
	message := messages[rand.Intn(len(messages))]
	d.SendMessage(channelID, message)
	return false
}
