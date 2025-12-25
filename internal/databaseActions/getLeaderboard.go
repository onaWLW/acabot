package databaseActions

import (
	"acabot/internal/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func GetLeaderboard(db *gorm.DB, sId string) string {
	var leaderboard strings.Builder
	leaderboard.WriteString("# ~ ~ ~ 1️⃣ 3️⃣ : 1️⃣ 2️⃣  Leaderboard ~ ~ ~")

	var streaks []model.Score
	db.
		Where("server_id = ? AND streak >= ?", sId, 2).
		Order("streak DESC, acab_count DESC").
		Limit(20).
		Find(&streaks)

	var maxUsernameLength = 0
	for _, streak := range streaks {
		if len([]rune(streak.UserName)) > maxUsernameLength {
			maxUsernameLength = len([]rune(streak.UserName))
		}
	}

	var rank int = 1
	if len(streaks) != 0 {
		leaderboard.WriteString("\n## 🔥   Streak Champions")
		for i, streak := range streaks {
			if i > 0 && !(streaks[i-1].Streak == streak.Streak && streaks[i-1].AcabCount == streak.AcabCount) {
				rank += 1
				if rank == 2 {
					leaderboard.WriteString("\n-#   \n")
				}
			}

			for len([]rune(streak.UserName)) != maxUsernameLength {
				streak.UserName += " "
			}

			switch rank {
			case 1:
				leaderboard.WriteString("\n### 👑 🥇  `" + streak.UserName + "`     🔥  " + strconv.Itoa(streak.Streak) + "     (" + strconv.Itoa(streak.AcabCount) + " total)")
			case 2:
				leaderboard.WriteString("\n🥈  `" + streak.UserName + "`     🔥  " + strconv.Itoa(streak.Streak) + "     (" + strconv.Itoa(streak.AcabCount) + " total)")
			case 3:
				leaderboard.WriteString("\n🥉  `" + streak.UserName + "`     🔥  " + strconv.Itoa(streak.Streak) + "     (" + strconv.Itoa(streak.AcabCount) + " total)")
			default:
				leaderboard.WriteString("\n🤝  `" + streak.UserName + "`     🔥  " + strconv.Itoa(streak.Streak) + "     (" + strconv.Itoa(streak.AcabCount) + " total)")
			}
		}
	}

	var top []model.Score
	db.
		Where("server_id = ?", sId).
		Order("acab_count DESC").
		Limit(10).
		Find(&top)

	var rankIcons = [10]string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}

	maxUsernameLength = 0
	for _, t := range top {
		if len([]rune(t.UserName)) > maxUsernameLength {
			maxUsernameLength = len([]rune(t.UserName))
		}
	}

	rank = 0
	if len(top) != 0 {
		leaderboard.WriteString("\n\n## ⭐   Total Score Leaders")
		for i, t := range top {
			if i > 0 && top[i-1].AcabCount != t.AcabCount {
				rank += 1
			}

			for len([]rune(t.UserName)) != maxUsernameLength {
				t.UserName += " "
			}

			leaderboard.WriteString("\n" + rankIcons[rank] + "   `" + t.UserName + "`  🤜  " + strconv.Itoa(t.AcabCount))
		}
	}

	return leaderboard.String()
}
