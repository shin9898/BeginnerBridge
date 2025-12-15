package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// GameScene はゲームのシーン
type GameScene int

const (
	SceneTitle GameScene = iota
	SceneNaming
	SceneMain
	SceneTraining
	SceneRaceSelect
	SceneRace
	SceneRaceResult
	SceneGameOver
)

// Game はメインのゲーム構造体
type Game struct {
	scene       GameScene
	playerHorse *Horse
	money       int
	week        int
	year        int
	message     string
	raceState   *RaceState
	selectedIdx int
	nameInput   string
	initialized bool
}

// NewGame は新しいゲームを作成する
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		scene:       SceneTitle,
		money:       50000, // 初期資金5万円
		week:        1,
		year:        1,
		message:     "",
		selectedIdx: 0,
		nameInput:   "",
		initialized: true,
	}
}

// Update はゲームの更新処理
func (g *Game) Update() error {
	switch g.scene {
	case SceneTitle:
		g.updateTitle()
	case SceneNaming:
		g.updateNaming()
	case SceneMain:
		g.updateMain()
	case SceneTraining:
		g.updateTraining()
	case SceneRaceSelect:
		g.updateRaceSelect()
	case SceneRace:
		g.updateRace()
	case SceneRaceResult:
		g.updateRaceResult()
	case SceneGameOver:
		g.updateGameOver()
	}
	return nil
}

// Draw はゲームの描画処理
func (g *Game) Draw(screen *ebiten.Image) {
	// 背景を緑色で塗る（競馬場の芝をイメージ）
	screen.Fill(color.RGBA{34, 85, 51, 255})

	switch g.scene {
	case SceneTitle:
		g.drawTitle(screen)
	case SceneNaming:
		g.drawNaming(screen)
	case SceneMain:
		g.drawMain(screen)
	case SceneTraining:
		g.drawTraining(screen)
	case SceneRaceSelect:
		g.drawRaceSelect(screen)
	case SceneRace:
		g.drawRace(screen)
	case SceneRaceResult:
		g.drawRaceResult(screen)
	case SceneGameOver:
		g.drawGameOver(screen)
	}
}

// Layout はウィンドウサイズを返す
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

// === タイトル画面 ===
func (g *Game) updateTitle() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.scene = SceneNaming
	}
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	// タイトル背景
	vector.DrawFilledRect(screen, 150, 100, 500, 400, color.RGBA{50, 50, 80, 240}, false)
	vector.DrawFilledRect(screen, 155, 105, 490, 390, color.RGBA{30, 30, 50, 255}, false)

	ebitenutil.DebugPrintAt(screen, "========================================", 200, 150)
	ebitenutil.DebugPrintAt(screen, "    HORSE RACING TYCOON", 200, 180)
	ebitenutil.DebugPrintAt(screen, "    ~ 競走馬育成シミュレーション ~", 200, 210)
	ebitenutil.DebugPrintAt(screen, "========================================", 200, 240)

	ebitenutil.DebugPrintAt(screen, "あなたは新人馬主です。", 220, 300)
	ebitenutil.DebugPrintAt(screen, "愛馬を育て、G1レースで勝利を目指しましょう！", 220, 330)

	ebitenutil.DebugPrintAt(screen, "[ Press ENTER or SPACE to Start ]", 250, 420)
}

// === 名前入力画面 ===
func (g *Game) updateNaming() {
	// アルファベットと日本語入力は難しいので、簡易版として固定名から選択
	names := []string{"ゴールドスター", "サンダーボルト", "ミラクルホープ", "ドリームチェイサー", "フェニックス"}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.selectedIdx--
		if g.selectedIdx < 0 {
			g.selectedIdx = len(names) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.selectedIdx++
		if g.selectedIdx >= len(names) {
			g.selectedIdx = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.playerHorse = NewHorse(names[g.selectedIdx], true)
		g.scene = SceneMain
	}
}

func (g *Game) drawNaming(screen *ebiten.Image) {
	names := []string{"ゴールドスター", "サンダーボルト", "ミラクルホープ", "ドリームチェイサー", "フェニックス"}

	vector.DrawFilledRect(screen, 150, 100, 500, 400, color.RGBA{50, 50, 80, 240}, false)
	vector.DrawFilledRect(screen, 155, 105, 490, 390, color.RGBA{30, 30, 50, 255}, false)

	ebitenutil.DebugPrintAt(screen, "愛馬の名前を選んでください", 280, 130)
	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  Enter: 決定", 290, 160)

	for i, name := range names {
		prefix := "  "
		if i == g.selectedIdx {
			prefix = "> "
		}
		ebitenutil.DebugPrintAt(screen, prefix+name, 320, 220+i*40)
	}
}

// === メイン画面 ===
func (g *Game) updateMain() {
	// ゲームオーバー判定
	if g.year > 3 || (g.playerHorse != nil && g.playerHorse.Age > 5) {
		g.scene = SceneGameOver
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.scene = SceneTraining
		g.selectedIdx = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.scene = SceneRaceSelect
		g.selectedIdx = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		// 休養
		g.message = g.playerHorse.Rest()
		g.advanceWeek()
	}
}

func (g *Game) drawMain(screen *ebiten.Image) {
	// ステータスパネル
	vector.DrawFilledRect(screen, 20, 20, 350, 250, color.RGBA{40, 40, 60, 230}, false)
	vector.DrawFilledRect(screen, 25, 25, 340, 240, color.RGBA{20, 20, 40, 255}, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %d年目 第%d週 】", g.year, g.week), 40, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("資金: %d円", g.money), 40, 70)

	if g.playerHorse != nil {
		ebitenutil.DebugPrintAt(screen, "───────────────────", 40, 95)
		ebitenutil.DebugPrintAt(screen, g.playerHorse.StatusString(), 40, 110)
	}

	// アクションパネル
	vector.DrawFilledRect(screen, 400, 20, 380, 250, color.RGBA{40, 40, 60, 230}, false)
	vector.DrawFilledRect(screen, 405, 25, 370, 240, color.RGBA{20, 20, 40, 255}, false)

	ebitenutil.DebugPrintAt(screen, "【 アクション 】", 420, 40)
	ebitenutil.DebugPrintAt(screen, "───────────────────", 420, 65)
	ebitenutil.DebugPrintAt(screen, "[1] トレーニング", 420, 90)
	ebitenutil.DebugPrintAt(screen, "[2] レースに出走", 420, 120)
	ebitenutil.DebugPrintAt(screen, "[3] 休養する", 420, 150)

	// メッセージパネル
	if g.message != "" {
		vector.DrawFilledRect(screen, 20, 300, 760, 80, color.RGBA{60, 40, 40, 230}, false)
		vector.DrawFilledRect(screen, 25, 305, 750, 70, color.RGBA{40, 20, 20, 255}, false)
		ebitenutil.DebugPrintAt(screen, g.message, 40, 330)
	}

	// ヘルプ
	ebitenutil.DebugPrintAt(screen, "数字キーでアクションを選択", 300, 560)
}

// === トレーニング画面 ===
func (g *Game) updateTraining() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.selectedIdx--
		if g.selectedIdx < 0 {
			g.selectedIdx = 5
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.selectedIdx++
		if g.selectedIdx > 5 {
			g.selectedIdx = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.selectedIdx < 5 {
			g.message = g.playerHorse.Train(TrainingType(g.selectedIdx))
			g.advanceWeek()
			g.scene = SceneMain
		} else {
			g.scene = SceneMain
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.scene = SceneMain
	}
}

func (g *Game) drawTraining(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 150, 80, 500, 440, color.RGBA{50, 50, 80, 240}, false)
	vector.DrawFilledRect(screen, 155, 85, 490, 430, color.RGBA{30, 30, 50, 255}, false)

	ebitenutil.DebugPrintAt(screen, "【 トレーニング選択 】", 310, 100)
	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  Enter: 決定  Esc: 戻る", 250, 130)

	trainings := []struct {
		name string
		desc string
	}{
		{"スピード", "最高速度を上げる"},
		{"スタミナ", "持久力を上げる"},
		{"パワー", "加速力を上げる"},
		{"根性", "終盤の粘り強さを上げる"},
		{"賢さ", "レース運びを上げる"},
		{"戻る", "メニューに戻る"},
	}

	for i, t := range trainings {
		prefix := "  "
		if i == g.selectedIdx {
			prefix = "> "
		}
		y := 180 + i*50
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%s", prefix, t.name), 250, y)
		if i < 5 {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("   - %s", t.desc), 250, y+20)
		}
	}

	// 現在の疲労度を表示
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("現在の疲労度: %d%%", g.playerHorse.Fatigue), 250, 480)
	if g.playerHorse.Fatigue > 70 {
		ebitenutil.DebugPrintAt(screen, "※ 疲労が高いです。休養をおすすめします", 250, 500)
	}
}

// === レース選択画面 ===
func (g *Game) updateRaceSelect() {
	maxIdx := len(AvailableRaces)
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.selectedIdx--
		if g.selectedIdx < 0 {
			g.selectedIdx = maxIdx
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.selectedIdx++
		if g.selectedIdx > maxIdx {
			g.selectedIdx = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.selectedIdx < len(AvailableRaces) {
			race := AvailableRaces[g.selectedIdx]
			g.raceState = NewRaceState(&race, g.playerHorse, 7) // 7頭のライバル
			g.raceState.Start()
			g.scene = SceneRace
		} else {
			g.scene = SceneMain
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.scene = SceneMain
	}
}

func (g *Game) drawRaceSelect(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 100, 50, 600, 500, color.RGBA{50, 50, 80, 240}, false)
	vector.DrawFilledRect(screen, 105, 55, 590, 490, color.RGBA{30, 30, 50, 255}, false)

	ebitenutil.DebugPrintAt(screen, "【 レース選択 】", 340, 70)
	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  Enter: 出走  Esc: 戻る", 270, 95)

	for i, race := range AvailableRaces {
		prefix := "  "
		if i == g.selectedIdx {
			prefix = "> "
		}
		y := 130 + i*35
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s[%s] %s (%dm) 賞金:%d円",
			prefix, race.Grade, race.Name, race.Distance, race.Grade.Prize()), 150, y)
	}

	// 戻るオプション
	prefix := "  "
	if g.selectedIdx == len(AvailableRaces) {
		prefix = "> "
	}
	ebitenutil.DebugPrintAt(screen, prefix+"戻る", 150, 130+len(AvailableRaces)*35+20)
}

// === レース画面 ===
func (g *Game) updateRace() {
	if g.raceState.IsRunning {
		g.raceState.Update()
	}
	if g.raceState.IsFinished {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.scene = SceneRaceResult
		}
	}
}

func (g *Game) drawRace(screen *ebiten.Image) {
	// レース場の背景
	vector.DrawFilledRect(screen, 0, 0, 800, 600, color.RGBA{34, 100, 34, 255}, false)

	// コース
	vector.DrawFilledRect(screen, 50, 150, 700, 300, color.RGBA{139, 90, 43, 255}, false)

	// レース名
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %s 】 %dm",
		g.raceState.Race.Name, g.raceState.Race.Distance), 300, 30)

	// 馬を描画
	for i, entry := range g.raceState.Entries {
		y := float32(170 + i*35)
		x := float32(60 + entry.Position*6.5)

		// 馬を四角で表示
		horseColor := color.RGBA{139, 69, 19, 255} // 茶色
		if entry.Horse.IsPlayer {
			horseColor = color.RGBA{255, 215, 0, 255} // 金色
		}
		vector.DrawFilledRect(screen, x, y, 25, 20, horseColor, false)

		// 馬名
		nameColor := ""
		if entry.Horse.IsPlayer {
			nameColor = "★"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%s", nameColor, entry.Horse.Name), int(x)-20, int(y)-15)
	}

	// ゴールライン
	vector.DrawFilledRect(screen, 720, 150, 5, 300, color.RGBA{255, 255, 255, 255}, false)
	ebitenutil.DebugPrintAt(screen, "GOAL", 710, 130)

	// 現在順位
	playerRank := g.raceState.GetPlayerRank()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("現在順位: %d位 / %d頭",
		playerRank, len(g.raceState.Entries)), 50, 480)

	if g.raceState.IsFinished {
		vector.DrawFilledRect(screen, 250, 200, 300, 100, color.RGBA{0, 0, 0, 200}, false)
		ebitenutil.DebugPrintAt(screen, "レース終了！", 350, 230)
		ebitenutil.DebugPrintAt(screen, "[ Press ENTER to continue ]", 290, 270)
	}
}

// === レース結果画面 ===
func (g *Game) updateRaceResult() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// 賞金を獲得
		playerEntry := g.raceState.GetPlayerEntry()
		if playerEntry != nil && playerEntry.Rank <= 3 {
			prize := 0
			switch playerEntry.Rank {
			case 1:
				prize = g.raceState.Race.Grade.Prize()
			case 2:
				prize = g.raceState.Race.Grade.Prize() / 2
			case 3:
				prize = g.raceState.Race.Grade.Prize() / 4
			}
			g.money += prize
			g.message = fmt.Sprintf("%s が %d位でゴール！ 賞金 %d円 を獲得！",
				g.playerHorse.Name, playerEntry.Rank, prize)
		} else if playerEntry != nil {
			g.message = fmt.Sprintf("%s は %d位でした...",
				g.playerHorse.Name, playerEntry.Rank)
		}

		g.advanceWeek()
		g.scene = SceneMain
	}
}

func (g *Game) drawRaceResult(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 150, 80, 500, 440, color.RGBA{50, 50, 80, 240}, false)
	vector.DrawFilledRect(screen, 155, 85, 490, 430, color.RGBA{30, 30, 50, 255}, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %s 結果 】", g.raceState.Race.Name), 300, 100)

	for i, entry := range g.raceState.Results {
		y := 150 + i*35
		marker := ""
		if entry.Horse.IsPlayer {
			marker = "★ "
		}
		prize := ""
		if i == 0 {
			prize = fmt.Sprintf(" (+%d円)", g.raceState.Race.Grade.Prize())
		} else if i == 1 {
			prize = fmt.Sprintf(" (+%d円)", g.raceState.Race.Grade.Prize()/2)
		} else if i == 2 {
			prize = fmt.Sprintf(" (+%d円)", g.raceState.Race.Grade.Prize()/4)
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d位: %s%s%s",
			entry.Rank, marker, entry.Horse.Name, prize), 200, y)
	}

	ebitenutil.DebugPrintAt(screen, "[ Press ENTER to continue ]", 280, 480)
}

// === ゲームオーバー画面 ===
func (g *Game) updateGameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// ゲームをリセット
		*g = *NewGame()
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 150, 100, 500, 400, color.RGBA{50, 30, 30, 240}, false)
	vector.DrawFilledRect(screen, 155, 105, 490, 390, color.RGBA{30, 20, 20, 255}, false)

	ebitenutil.DebugPrintAt(screen, "========================================", 200, 130)
	ebitenutil.DebugPrintAt(screen, "        GAME OVER", 200, 160)
	ebitenutil.DebugPrintAt(screen, "========================================", 200, 190)

	if g.playerHorse != nil {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("愛馬: %s", g.playerHorse.Name), 250, 250)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("最終戦績: %d戦 %d勝",
			g.playerHorse.Races, g.playerHorse.Wins), 250, 280)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("最終資金: %d円", g.money), 250, 310)

		// 評価
		var rating string
		if g.playerHorse.Wins >= 5 {
			rating = "伝説の名馬！"
		} else if g.playerHorse.Wins >= 3 {
			rating = "素晴らしい活躍でした！"
		} else if g.playerHorse.Wins >= 1 {
			rating = "勝利を経験できました"
		} else {
			rating = "次こそは頑張りましょう..."
		}
		ebitenutil.DebugPrintAt(screen, rating, 250, 350)
	}

	ebitenutil.DebugPrintAt(screen, "[ Press ENTER to Restart ]", 280, 430)
}

// advanceWeek は週を進める
func (g *Game) advanceWeek() {
	g.week++
	if g.week > 52 {
		g.week = 1
		g.year++
		if g.playerHorse != nil {
			g.playerHorse.Age++
		}
	}

	// 調子の自然変動
	if g.playerHorse != nil {
		moodChange := rand.Intn(11) - 5 // -5 ~ +5
		g.playerHorse.Mood += moodChange
		if g.playerHorse.Mood > 100 {
			g.playerHorse.Mood = 100
		}
		if g.playerHorse.Mood < 20 {
			g.playerHorse.Mood = 20
		}
	}
}
