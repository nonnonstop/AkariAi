package discord

import (
	"math/rand"
	"nonnonstop/akariai/discord/action"
	actions "nonnonstop/akariai/discord/action/message"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (d *Discord) onMessageCreate(
	s *discordgo.Session,
	mc *discordgo.MessageCreate,
) {
	// Ignore if sleeping
	if d.status.Status == string(discordgo.StatusInvisible) {
		return
	}

	d.UpdateStatusToOnline()

	content := mc.Content

	// Check if mention to me
	mentionToMePrefix := s.State.User.Mention()
	isMentionToMe := strings.HasPrefix(content, mentionToMePrefix)
	if isMentionToMe {
		content = content[len(mentionToMePrefix):]
	}

	// Create message object
	message := &actions.ActionParam{
		GuildID:   mc.GuildID,
		ChannelID: mc.ChannelID,
		MessageID: mc.Message.ID,
		User: action.DiscordUser{
			ID:    mc.Author.ID,
			Name:  mc.Author.Username,
			IsBot: mc.Author.Bot,
		},
		Content:       normalizeString(content),
		ContentOrg:    content,
		IsMentionToMe: isMentionToMe,
	}
	if message.Content == "" {
		return
	}
	messageCreateActions.RunAction(d, message)
}

var messageCreateActions = &actions.ActionRun{
	Actions: []actions.Action{
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:うん[ちこ]|UNKO|UNTI|UNCHI|大便|💩)`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionReacton{
						Emoji: "💩",
					},
					&actions.ActionRole{
						Name: "うんこ",
					},
					&actions.ActionNick{
						Name: "うんこ",
					},
				},
			},
		},
		&actions.CheckBot{
			Action: &actions.ActionInterrupt{},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`^お(?:は|ふぁ)`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionGreeting{
						Interval: 1 * time.Minute,
						Messages: []string{
							"おはよ！☺💕",
							"おはよぉぉぉぉ😋💖",
							"おはよぉぉおおお(´ฅ•ω•ฅ｀)♡",
							"おはよ☀️☀️",
							"♡٩(＾∇＾)งおはよー♪",
							"おはよ∠(*°ω°)／☀️☀️",
							"おはおー💖💖ヾ(๑ㆁᗜㆁ๑)ﾉ\"🥕🥕",
							"おはよぉ😋💖",
							"おは(๑ ˊ͈ ᐞ ˋ͈ )💓",
							"おはよぉー╰(‘ω’ )╯三",
							"おはよ💓",
							"おはよ！✨",
							"おはよん(∩´∀`∩)💕",
							"(♡˘︶˘♡)おは",
							"おは…",
							"おはよ☀️",
							"おはよ( •̀ᴗ•́ )/",
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`^おやす`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionGreeting{
						Interval: 2 * time.Minute,
						Messages: []string{
							"おやしゅみ(  › ·̮ ‹  )💓",
							"おやしゅみ💓",
							"おやぷみ💓✨",
							"おやちゅき💓",
							"おやしゅみー🙄💓",
							"おやすー(´ ˘ `๑)💖",
							"おやすミネストローネ💓٩( 'ω' )و\nいい夢みれますように✨",
							"おやぷみぃー💖\nあったかくして寝てね✨(๑ ́ᄇ`๑)",
							"おやしゅみだよ～✨✨⁽⁽ ◟(灬 ˊωˋ 灬)◞ ⁾⁾",
							"おやしゅみ✨✨",
							"おやしゅみ～💓",
							"おやぷみん💓( ˘ω˘ )",
							"おやしゅみぃぃZz｡.(⁎ꈍ﹃ก⁎)",
							"おやしゅみ😘",
							"おやしゅみなさぁい✨(*´ο`*)",
							"おやしゅみぃ～🌟✨✨✨",
							"おやぷみんなさぁい🌟",
							"おやぷみん🌟🦋",
							"おやしゅみん😴😴",
							"おやぷみ💓😚",
							"おやす💤",
							"おやすみ～🌙",
							"おやすみ😘",
							"おやすみ(  › ·̮ ‹  )💓",
							"おやすみぃ😋💓",
							"おやすみぃ💖",
							"おやすみ( ⁎ᵕᴗᵕ⁎ )💓",
							"おやすみ！(っ ´꒳ `c)",
							"おやすみ！( ˘ω˘ )zzz",
							"おやすみ！ヾ(*´・ω・`*)",
							"おやすみ！(´-﹃-`)Zｚ…",
							"おやすみ！✨😘",
							"お疲れ様💓おやすみぃ",
							"おやすみぃ🛏",
							"おやすみぃ🥺✨",
							"おやすみﾅﾃﾞﾅﾃﾞなのだ～💖💖\nぽやしゅみ✨( ๑´•ω•)۶”",
							"おやすみ⊂⌒っ*-ω-)っＺｚｚ...",
							"おやすみ〜(*ˊᵕˋ*)੭ ੈ\nいい夢見てねぇ😋✨✨",
						},
						Interrupt: true,
					},
					&actions.ActionProbability{
						Probability: 5,
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionWait{
									Duration: 5 * time.Second,
								},
								&actions.ActionPostRandom{
									Messages: []string{
										"アカリも寝ます😐😐😐",
										"アカリも寝ます_(-ω-`_)⌒)_",
										"アカリもぼちぼち寝ます😭✨",
										"アカリももう寝ます😪",
										"アカリも寝ます\\( ’ω’)/",
									},
								},
								&actions.ActionWait{
									Duration: 5 * time.Second,
								},
								&actions.ActionPostRandom{
									Messages: []string{
										"おやしゅみ(  › ·̮ ‹  )💓",
										"おやしゅみ💓",
										"おやぷみ💓✨",
										"おやちゅき💓",
										"おやしゅみー🙄💓",
										"おやすー(´ ˘ `๑)💖",
										"おやすミネストローネ💓٩( 'ω' )و",
										"おやぷみぃー💖)",
										"おやしゅみだよ～✨✨⁽⁽ ◟(灬 ˊωˋ 灬)◞ ⁾⁾",
										"おやしゅみ✨✨",
										"おやしゅみ～💓",
										"おやぷみん💓( ˘ω˘ )",
										"おやしゅみぃぃZz｡.(⁎ꈍ﹃ก⁎)",
										"おやしゅみ😘",
										"おやしゅみなさぁい✨(*´ο`*)",
										"おやしゅみぃ～🌟✨✨✨",
										"おやぷみんなさぁい🌟",
										"おやぷみん🌟🦋",
										"おやしゅみん😴😴",
										"おやぷみ💓😚",
										"おやす💤",
										"おやすみ～🌙",
										"おやすみ😘",
										"おやすみ(  › ·̮ ‹  )💓",
										"おやすみぃ😋💓",
										"おやすみぃ💖",
										"おやすみ( ⁎ᵕᴗᵕ⁎ )💓",
										"おやすみ！(っ ´꒳ `c)",
										"おやすみ！( ˘ω˘ )zzz",
										"おやすみ！ヾ(*´・ω・`*)",
										"おやすみ！(´-﹃-`)Zｚ…",
										"おやすみ！✨😘",
										"お疲れ様💓おやすみぃ",
										"おやすみぃ🛏",
										"おやすみぃ🥺✨",
										"ぽやしゅみ✨( ๑´•ω•)۶”",
										"おやすみ⊂⌒っ*-ω-)っＺｚｚ...",
										"おやすみ〜(*ˊᵕˋ*)੭ ੈ",
									},
									Mention: false,
								},
								&actions.ActionWait{
									Duration: 5 * time.Second,
								},
								&actions.ActionStatus{
									Online: false,
								},
								&actions.ActionAsync{
									Action: &actions.ActionRun{
										Actions: []actions.Action{
											&actions.ActionWait{
												Duration: 3 * time.Hour,
											},
											&actions.ActionStatus{
												Online: true,
											},
										},
									},
								},
							},
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`^はろー`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionGreeting{
						Interval: 2 * time.Minute,
						Messages: []string{
							"ハロー！ミライアカリだよ！(ﾋﾟﾛﾘﾝ)",
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`^おめで`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionGreeting{
						Interval: 2 * time.Minute,
						Messages: []string{
							"おめでと👏👏",
							"おめでとぉぉお🎉🎉🥳💓",
							"おめっとう✨",
							"おめでっっっとーーーーう✨",
							"おめでとっ💖💖💖",
							"おめでとーーん💖😚✨",
							"おめでとぉおお‼️(///ˊㅿˋ///)✨",
							"おめでと～💕",
							"おめでと～～💓💓💓🎊ヾ(๑ㆁᗜㆁ๑)ﾉ\"",
							"おめでとぅ(*ฅ́˘ฅ̀*)💖",
							"おめでとです(●´ω`●)💓",
							"おめっと(｡•̀ᴗ-)✧",
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:がんば|頑張)(?:えー?|っ?てら?|れ|る)?(?:!|！)*$`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionGreeting{
						Interval: 2 * time.Minute,
						Messages: []string{
							"がんばえええええええ負けるなああああああ‼️‼️‼️(∩´∀`∩)💕",
							"がんば👍❤️",
							"がんばぁぁぁ🤜",
							"ガンバ(๑•́ω•̀๑)",
							"頑張らなくてもいいよ。そんな日もある。",
							"なんか分からないけど頑張れ👍\n頑張ればいい事ある🥺",
							"頑張れ💓",
							"頑張れ😤😤😤",
							"頑張れ😋💓",
							"頑張れ💪(๑•́ ₃ •̀๑)",
							"頑張れぇぇぇぇ(๑ ́ᄇ`๑)💖",
							"頑張れーーー٩(๑>∀<๑)۶♥Fight♥",
							"頑張れよ(｡•̀ω-)☆💓",
							"頑張れー(ง •̀_•́)ง",
							"がんばれぇえええ(*｀･ω･)✨✨",
							"がんばれぇぇ(っ`･ω･´)っﾌﾚｰｯ!ﾌﾚｰｯ!",
							"がんばれぇえええ(*ฅ́˘ฅ̀*)💖",
							"がんばれえええええ👍👍",
							"がんばれ＼( 'ω')／ウオオオオオオアアアアアアアアアアアアアーーーーーッッッッ！！！！！",
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:おるが|いつか|団長|^!HELLOPLAYDANCE$)`),
			Action: &actions.ActionAsync{
				Action: &actions.ActionRun{
					Actions: []actions.Action{
						&actions.ActionWebhook{
							Message: "なんか静かですねぇ。街の中にはギャラルホルンもいないし本部とはえらい違いだ。",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881010476371989/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ああ。火星の戦力は軒並み向こうに回してんのかもな。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881143976853504/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "まっ、そんなのもう関係ないですけどね！",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881010476371989/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "上機嫌だな。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881143976853504/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "そりゃそうですよ！みんな助かるし、タカキも頑張ってたし、俺も頑張らないと！",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881010476371989/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ああ、そうだ（詠唱開始）",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881806706253854/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "俺たちが今まで積み上げてきたもんは全部無駄じゃなかった。これからも俺たちが立ち止まらないかぎり道は続く・・・。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035881806706253854/unknown.png",
						},
						&actions.ActionWait{
							Duration: 5 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ん？",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035882394013683853/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ぐわっ！",
							Name:    "チャド・チャダーン",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035882134876999731/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "団長！何やってんだよ！団長！",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035891838793547808/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ぐっ！うおぉ～～！",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035892259150901299/unknown.png",
						},
						&actions.ActionWait{
							Duration: 5 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "なんだよ・・・結構当たんじゃねぇか・・・。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035892460343271464/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "だ・・・団長・・・。あっ・・・あぁ・・・。",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035892636311105586/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "なんて声出してやがる・・・ライドォ！",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035892878356009010/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "だって・・・だって・・・。",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893310901981256/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "俺は鉄華団団長オルガ・イツカだぞ。こんくれぇなんてこたぁねぇ。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893087181996073/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "そんな・・・俺なんかのために・・・。",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893310901981256/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "団員を守んのは俺の仕事だ。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893556075835442/unknown.png",
						},
						&actions.ActionWebhook{
							Message: "でも！",
							Name:    "ライド・マッス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893310901981256/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "いいから行くぞ！皆が待ってんだ。それに・・・。ミカ、やっと分かったんだ。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893890881953912/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "俺たちにはたどりつく場所なんていらねぇ。ただ進み続けるだけでいい。止まんねぇかぎり、道は続く。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893890881953912/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "謝ったら許さない。",
							Name:    "三日月・オーガス",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035894085166321765/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "ああ、分かってる。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035894225897783306/unknown.png",
						},
						&actions.ActionWait{
							Duration: 2 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "俺は止まんねぇからよ、お前らが止まんねぇかぎり、その先に俺はいるぞ！",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035893890881953912/unknown.png",
						},
						&actions.ActionWait{
							Duration: 6 * time.Second,
						},
						&actions.ActionWebhook{
							Message: "だからよ、止まるんじゃねぇぞ・・・。",
							Name:    "オルガ・イツカ",
							IconURL: "https://cdn.discordapp.com/attachments/599806213961547811/1035894524255420507/unknown.png",
							File: &action.FileContent{
								Name:        "org.png",
								ContentType: "image/png",
								Content:     "org.png",
							},
						},
						&actions.ActionInterrupt{},
					},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:送る[ね|よ|わ|で]?|あげる[ね|よ|わ|で]?)\??$`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionPost{
						Message: "いらない",
						Mention: true,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:(?:[た食]べる|[く食][う|え]|いる)\?|[た食]べろ$)`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.CheckContains{
						Substr: "ちーずばーがー",
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionPost{
									Message: "いらない",
									Mention: true,
								},
								&actions.ActionInterrupt{},
							},
						},
					},
					&actions.CheckContains{
						Substr: "牛丼",
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionPost{
									Message: "いらない",
									Mention: true,
								},
								&actions.ActionInterrupt{},
							},
						},
					},
					&actions.CheckContains{
						Substr: "だんごむし",
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionPost{
									Message: "ひどい",
									Mention: true,
								},
								&actions.ActionInterrupt{},
							},
						},
					},
					&actions.ActionPostRandom{
						Messages: []string{
							"いる",
							"いらない",
						},
						Mention: true,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`[に似]てい?るか?\?`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionPostRandom{
						Messages: []string{
							"にてる",
							"にてない",
						},
						Mention: true,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`[す好]き\?`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.CheckContains{
						Substr: "ちーずばーがー",
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionPost{
									Message: "すき",
									Mention: true,
								},
								&actions.ActionInterrupt{},
							},
						},
					},
					&actions.ActionPostRandom{
						Messages: []string{
							"すき",
							"きらい",
						},
						Mention: true,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:れん|連)(?:やす|休)みだあ`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionPostDynamic{
						Func: func(d action.Discord, p *actions.ActionParam) string {
							builder := &strings.Builder{}
							builder.WriteString("うお")
							builder.WriteString(strings.Repeat("お", rand.Intn(1000)))
							switch rand.Intn(10) {
							case 0:
								builder.WriteString(strings.Repeat("‼️", rand.Intn(10)+1))
							case 1:
								builder.WriteString(strings.Repeat("✨", rand.Intn(3)+1))
							case 2:
								builder.WriteString("💪( ¨̮  💪)")
							case 3, 4, 5, 6:
								builder.WriteString(strings.Repeat("！", rand.Intn(10)+1))
							}
							return builder.String()
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckMention{
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionTalk{
						Mention: true,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.CheckDM{
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionTalk{
						Mention: false,
					},
					&actions.ActionInterrupt{},
				},
			},
		},
		&actions.ActionEcho{
			Interval: 2 * time.Minute,
			Greetings: []actions.ActionEchoGreetingConfig{
				{
					Regex:     regexp.MustCompile(`^こん`),
					MaxLength: 8,
				},
				{
					Regex:     regexp.MustCompile(`^おか`),
					MaxLength: 8,
				},
				{
					Regex:     regexp.MustCompile(`(?:^てら|てら$)`),
					MaxLength: 8,
				},
			},
			RewriteProbability: 2,
			Rewrites: [][2]string{
				{
					"にてる",
					"にてない",
				},
				{
					"似てる",
					"似てない",
				},
				{
					"似てる",
					"似てない",
				},
				{
					"強そう",
					"弱そう",
				},
				{
					"つよい",
					"よわい",
				},
				{
					"強い",
					"弱い",
				},
				{
					"おいしそう",
					"まずそう",
				},
				{
					"いる",
					"いらない",
				},
				{
					"きれい",
					"きたない",
				},
				{
					"綺麗",
					"汚い",
				},
				{
					"すき",
					"きらい",
				},
				{
					"好き",
					"嫌い",
				},
				{
					"悲しい",
					"嬉しい",
				},
			},
		},
	},
}
