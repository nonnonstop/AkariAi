package message

import (
	"math/rand"
	"nonnonstop/akariai/discord/action"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type echoState struct {
	prevUserID     string
	prevContent    string
	prevContentOrg string
	lastMessage    string
	lastGreeting   time.Time
}

type ActionEchoGreetingConfig struct {
	Regex     *regexp.Regexp
	MaxLength int
}

type ActionEcho struct {
	Interval           time.Duration
	Greetings          []ActionEchoGreetingConfig
	RewriteProbability int
	Rewrites           [][2]string
	states             map[string]*echoState
}

func (a *ActionEcho) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if a.states == nil {
		a.states = make(map[string]*echoState)
	}

	state := a.states[p.ChannelID]
	if state == nil {
		state = &echoState{}
		a.states[p.ChannelID] = state
	}

	// Return if same user
	if state.prevUserID == p.User.ID {
		state.prevContent = p.Content
		state.prevContentOrg = p.ContentOrg
		return false
	}

	// Return if no prev messages
	prevContent := state.prevContent
	if prevContent == "" {
		state.prevContent = p.Content
		state.prevContentOrg = p.ContentOrg
		state.prevUserID = p.User.ID
		return false
	}
	now := time.Now()
	if state.lastGreeting.Add(a.Interval).Before(now) {
		for _, greeting := range a.Greetings {
			maxLength := greeting.MaxLength
			if utf8.RuneCountInString(prevContent) >= maxLength {
				continue
			}
			if utf8.RuneCountInString(p.Content) >= maxLength {
				continue
			}
			regex := greeting.Regex
			if !regex.MatchString(prevContent) {
				continue
			}
			if !regex.MatchString(p.Content) {
				continue
			}

			d.SendMessage(p.ChannelID, state.prevContentOrg)
			state.lastMessage = state.prevContentOrg
			state.lastGreeting = now
			state.prevContent = ""
			state.prevContentOrg = ""
			state.prevUserID = ""
			return false
		}
	} else {
		// Return if send same message in short time
		if state.lastMessage == p.ContentOrg {
			return false
		}
	}

	// Return if same messages
	if prevContent != p.Content {
		state.prevContent = p.Content
		state.prevContentOrg = p.ContentOrg
		state.prevUserID = p.User.ID
		return false
	}

	message := state.prevContentOrg
	if rand.Intn(a.RewriteProbability) == 0 {
		for _, rewrite := range a.Rewrites {
			if strings.HasSuffix(message, rewrite[0]) {
				message = message[:len(message)-len(rewrite[0])] + rewrite[1]
				break
			}
			if strings.HasSuffix(message, rewrite[1]) {
				message = message[:len(message)-len(rewrite[1])] + rewrite[0]
				break
			}
		}
	}
	d.SendMessage(p.ChannelID, message)
	state.lastMessage = state.prevContentOrg
	state.lastGreeting = now
	state.prevContent = ""
	state.prevContentOrg = ""
	state.prevUserID = ""
	return false
}
