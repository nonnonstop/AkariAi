package discord

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (d *Discord) startStatusThread() {
	go func() {
		idleDuration := 10 * time.Minute
		idleTimer := time.NewTimer(idleDuration)
		defer idleTimer.Stop()

		gameTimer := time.NewTicker(15 * time.Minute)
		defer gameTimer.Stop()

		for {
			select {
			case s := <-d.statusUpdateCh:
				d.updateStatus(s)
				switch s {
				case discordgo.StatusOnline:
					// Extend idle timer
					extendTimer(idleTimer, idleDuration)
				case discordgo.StatusIdle:
					// Extend idle timer
					extendTimer(idleTimer, idleDuration)
				}
			case <-d.statusOnlineCh:
				switch d.status.Status {
				case string(discordgo.StatusOnline):
					// Extend timer
					extendTimer(idleTimer, idleDuration)
				case string(discordgo.StatusIdle):
					// Update status to online
					d.updateStatus(discordgo.StatusOnline)
					// Extend idle timer
					extendTimer(idleTimer, idleDuration)
				}
			case <-d.statusReadyCh:
				d.updateStatusGame()
			case <-idleTimer.C:
				if d.status.Status == string(discordgo.StatusOnline) {
					// Update status to idle
					d.updateStatus(discordgo.StatusIdle)
				}
			case <-gameTimer.C:
				if d.status.Status == string(discordgo.StatusOnline) {
					// Update game name
					d.updateStatusGame()
					// Extend idle timer
					extendTimer(idleTimer, idleDuration)
				}
			}
		}
	}()
}

func (d *Discord) updateStatus(s discordgo.Status) {
	d.status.Status = string(s)
	err := d.session.UpdateStatusComplex(d.status)
	if err != nil {
		d.logger.Errorln("Failed to update status: ", err)
	}
}

func (d *Discord) updateStatusGame() {
	activity := activities[rand.Intn(len(activities))]
	d.status.Activities = []*discordgo.Activity{
		{
			Name: activity,
			Type: discordgo.ActivityTypeGame,
		},
	}
	err := d.session.UpdateStatusComplex(d.status)
	if err != nil {
		d.logger.Errorln("Failed to update status: ", err)
	}
}

func extendTimer(t *time.Timer, duration time.Duration) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	t.Reset(duration)
}

var activities = []string{
	"",
	"KITCHEN",
	"斎藤さん",
	"VRChat",
	"どうぶつタワーバトル",
	"Getting Over It with Bennett Foddy",
	"おえかきの森",
	"Undertale",
	"QUICK,DRAW!",
	"Airtone",
	"Until Dawn: Rush of Blood",
	"SEIYA",
	"マリオカート",
	"セガキャッチャーオンライン",
	"Beat Saber",
	"Dead Hungry",
	"ねこあつめ",
	"Trollface Quest",
	"ラジオ体操",
	"モンスター娘TD",
	"スプラトゥーン3",
	"PigeonSimulator",
	"Emily Wants To Play",
	"Panty Party",
	"GO HOME",
}
