package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// GameScene はゲームのシーン
type GameScene int

const (
	SceneTitle GameScene = iota
	SceneTutorial
	SceneNaming
	SceneMain
	SceneTraining
	SceneRaceSelect
	SceneRace
	SceneRaceResult
	SceneGameOver
)

// 色の定義（UIテーマ）
var (
	ColorGold       = color.RGBA{255, 215, 0, 255}
	ColorDarkGold   = color.RGBA{184, 134, 11, 255}
	ColorBrown      = color.RGBA{139, 69, 19, 255}
	ColorDarkBrown  = color.RGBA{101, 67, 33, 255}
	ColorCream      = color.RGBA{255, 253, 208, 255}
	ColorGreen      = color.RGBA{34, 139, 34, 255}
	ColorDarkGreen  = color.RGBA{0, 100, 0, 255}
	ColorRed        = color.RGBA{178, 34, 34, 255}
	ColorBlue       = color.RGBA{70, 130, 180, 255}
	ColorPanelBg    = color.RGBA{45, 25, 15, 240}
	ColorPanelFrame = color.RGBA{139, 90, 43, 255}
	ColorWhite      = color.RGBA{255, 255, 255, 255}
	ColorLightGray  = color.RGBA{200, 200, 200, 255}
	ColorOrange     = color.RGBA{255, 165, 0, 255}
	ColorPurple     = color.RGBA{148, 0, 211, 255}
)

// Game はメインのゲーム構造体
type Game struct {
	scene         GameScene
	playerHorse   *Horse
	money         int
	week          int
	year          int
	message       string
	raceState     *RaceState
	selectedIdx   int
	nameInput     string
	initialized   bool
	tutorialPage  int
	animationTick int
}

// NewGame は新しいゲームを作成する
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		scene:         SceneTitle,
		money:         50000,
		week:          1,
		year:          1,
		message:       "",
		selectedIdx:   0,
		nameInput:     "",
		initialized:   true,
		tutorialPage:  0,
		animationTick: 0,
	}
}

// Update はゲームの更新処理
func (g *Game) Update() error {
	g.animationTick++
	switch g.scene {
	case SceneTitle:
		g.updateTitle()
	case SceneTutorial:
		g.updateTutorial()
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
	switch g.scene {
	case SceneTitle:
		g.drawTitle(screen)
	case SceneTutorial:
		g.drawTutorial(screen)
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
		g.scene = SceneTutorial
		g.tutorialPage = 0
	}
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	// 背景
	screen.Fill(ColorDarkGreen)
	vector.DrawFilledRect(screen, 0, 0, 800, 200, color.RGBA{0, 60, 0, 255}, false)
	vector.DrawFilledRect(screen, 0, 400, 800, 200, color.RGBA{0, 80, 0, 255}, false)

	// 競馬場のトラック
	vector.DrawFilledRect(screen, 50, 180, 700, 120, ColorBrown, false)
	vector.DrawFilledRect(screen, 60, 190, 680, 100, ColorDarkBrown, false)
	for i := 0; i < 7; i++ {
		vector.DrawFilledRect(screen, float32(100+i*100), 235, 50, 3, ColorWhite, false)
	}

	// 走る馬のアニメーション
	horseX := float32((g.animationTick * 3) % 900)
	g.drawHorseSprite(screen, horseX-100, 200, ColorGold)
	g.drawHorseSprite(screen, horseX-250, 210, ColorBrown)
	g.drawHorseSprite(screen, horseX-180, 220, ColorDarkBrown)

	// メインタイトルパネル
	DrawPanel(screen, 150, 320, 500, 250)

	// タイトル
	DrawTextCentered(screen, "DERBY DREAM", 400, 360, 32, ColorGold)
	DrawTextCentered(screen, "ダービードリーム", 400, 400, 24, ColorWhite)
	DrawTextCentered(screen, "～ 競走馬育成シミュレーション ～", 400, 440, 16, ColorLightGray)

	// スタートメッセージ
	if (g.animationTick/30)%2 == 0 {
		DrawTextCentered(screen, ">>> PRESS ENTER OR SPACE TO START <<<", 400, 520, 16, ColorGold)
	}
}

// drawHorseSprite は馬のスプライトを描画
func (g *Game) drawHorseSprite(screen *ebiten.Image, x, y float32, c color.RGBA) {
	vector.DrawFilledRect(screen, x, y, 40, 20, c, false)
	vector.DrawFilledRect(screen, x+35, y-5, 15, 15, c, false)
	legOffset := float32((g.animationTick / 5) % 4)
	vector.DrawFilledRect(screen, x+5, y+20+legOffset, 5, 12-legOffset, c, false)
	vector.DrawFilledRect(screen, x+30, y+20-legOffset, 5, 12+legOffset, c, false)
}

// === チュートリアル画面 ===
func (g *Game) updateTutorial() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.tutorialPage++
		if g.tutorialPage >= 4 {
			g.scene = SceneNaming
			g.selectedIdx = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.tutorialPage--
		if g.tutorialPage < 0 {
			g.tutorialPage = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.scene = SceneNaming
		g.selectedIdx = 0
	}
}

func (g *Game) drawTutorial(screen *ebiten.Image) {
	screen.Fill(ColorDarkGreen)
	DrawPanel(screen, 50, 50, 700, 500)

	DrawTextCentered(screen, fmt.Sprintf("【 ゲームの遊び方 】 %d / 4", g.tutorialPage+1), 400, 80, 20, ColorGold)

	// 進捗バー
	for i := 0; i < 4; i++ {
		c := color.RGBA{60, 60, 60, 255}
		if i <= g.tutorialPage {
			c = ColorGold
		}
		vector.DrawFilledRect(screen, float32(300+i*50), 105, 40, 8, c, false)
	}

	switch g.tutorialPage {
	case 0:
		g.drawTutorialPage1(screen)
	case 1:
		g.drawTutorialPage2(screen)
	case 2:
		g.drawTutorialPage3(screen)
	case 3:
		g.drawTutorialPage4(screen)
	}

	DrawTextCentered(screen, "← →: ページ移動  |  Enter/Space: 次へ  |  Esc: スキップ", 400, 530, 14, ColorLightGray)
}

func (g *Game) drawTutorialPage1(screen *ebiten.Image) {
	DrawTextCentered(screen, "◆ ゲームの目的 ◆", 400, 140, 20, ColorGold)

	DrawText(screen, "あなたは新人馬主です。", 150, 180, ColorWhite)
	DrawText(screen, "3年間で愛馬を育て、G1レースでの勝利を目指しましょう！", 150, 210, ColorWhite)

	DrawText(screen, "【 勝利条件 】", 150, 260, ColorGold)
	DrawText(screen, "・できるだけ多くのレースに勝利する", 170, 295, ColorWhite)
	DrawText(screen, "・賞金を稼いで資金を増やす", 170, 325, ColorWhite)
	DrawText(screen, "・G1レース（最高峰）での勝利を目指す", 170, 355, ColorWhite)

	DrawText(screen, "期間: 3年間（1年 = 52週）", 150, 410, ColorLightGray)
	DrawText(screen, "毎週1つのアクションを選んで進めます", 150, 440, ColorLightGray)
}

func (g *Game) drawTutorialPage2(screen *ebiten.Image) {
	DrawTextCentered(screen, "◆ 育成システム ◆", 400, 140, 20, ColorGold)
	DrawText(screen, "馬には5つのステータスがあります：", 150, 180, ColorWhite)

	stats := []struct {
		name  string
		desc  string
		color color.RGBA
	}{
		{"スピード", "最高速度に影響。レースの基本となる能力", ColorRed},
		{"スタミナ", "持久力。長距離レースで重要", ColorGreen},
		{"パワー  ", "加速力。スタートダッシュに影響", ColorBlue},
		{"根性    ", "終盤の粘り強さ。逆転に必要", ColorOrange},
		{"賢さ    ", "レース運びの安定性", ColorPurple},
	}

	for i, s := range stats {
		y := float32(220 + i*50)
		vector.DrawFilledRect(screen, 150, y, 18, 18, s.color, false)
		DrawText(screen, s.name, 180, float64(y), ColorWhite)
		DrawText(screen, s.desc, 290, float64(y), ColorLightGray)
	}

	DrawText(screen, "※ トレーニングで各ステータスを上げられます", 150, 490, ColorGold)
}

func (g *Game) drawTutorialPage3(screen *ebiten.Image) {
	DrawTextCentered(screen, "◆ 毎週のアクション ◆", 400, 140, 20, ColorGold)
	DrawText(screen, "毎週、以下の3つから1つを選びます：", 150, 180, ColorWhite)

	DrawText(screen, "[1] トレーニング", 150, 230, ColorGold)
	DrawText(screen, "    → ステータスを上げる（疲労が溜まる）", 150, 260, ColorLightGray)

	DrawText(screen, "[2] レース出走", 150, 310, ColorGold)
	DrawText(screen, "    → 勝てば賞金獲得！（疲労が溜まる）", 150, 340, ColorLightGray)

	DrawText(screen, "[3] 休養", 150, 390, ColorGold)
	DrawText(screen, "    → 疲労を回復、調子を整える", 150, 420, ColorLightGray)

	DrawText(screen, "※ 疲労が高いと調子が下がりやすくなります！", 150, 470, ColorOrange)
	DrawText(screen, "   バランスよく育成と休養を行いましょう", 150, 495, ColorLightGray)
}

func (g *Game) drawTutorialPage4(screen *ebiten.Image) {
	DrawTextCentered(screen, "◆ レースについて ◆", 400, 140, 20, ColorGold)
	DrawText(screen, "レースには3つのグレードがあります：", 150, 180, ColorWhite)

	// G3
	vector.DrawFilledRect(screen, 170, 220, 40, 22, color.RGBA{205, 127, 50, 255}, false)
	DrawText(screen, "G3", 178, 222, ColorWhite)
	DrawText(screen, "入門レベル   賞金: 3,000円", 230, 222, ColorWhite)

	// G2
	vector.DrawFilledRect(screen, 170, 255, 40, 22, color.RGBA{192, 192, 192, 255}, false)
	DrawText(screen, "G2", 178, 257, ColorPanelBg)
	DrawText(screen, "中級レベル   賞金: 8,000円", 230, 257, ColorWhite)

	// G1
	vector.DrawFilledRect(screen, 170, 290, 40, 22, ColorGold, false)
	DrawText(screen, "G1", 178, 292, ColorPanelBg)
	DrawText(screen, "最高峰！     賞金: 20,000円", 230, 292, ColorGold)

	DrawText(screen, "【 賞金配分 】", 150, 350, ColorGold)
	DrawText(screen, "1位: 100%  |  2位: 50%  |  3位: 25%", 170, 385, ColorWhite)

	DrawText(screen, "準備はいいですか？ 愛馬の名前を決めましょう！", 150, 450, ColorGold)
}

// === 名前入力画面 ===
func (g *Game) updateNaming() {
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
		g.message = "愛馬「" + g.playerHorse.Name + "」との旅が始まります！まずはトレーニングで能力を上げましょう。"
	}
}

func (g *Game) drawNaming(screen *ebiten.Image) {
	screen.Fill(ColorDarkGreen)
	DrawPanel(screen, 150, 80, 500, 440)

	DrawTextCentered(screen, "愛馬の名前を選んでください", 400, 110, 20, ColorGold)
	DrawTextCentered(screen, "この馬と3年間の旅を共にします", 400, 145, 14, ColorLightGray)

	names := []string{"ゴールドスター", "サンダーボルト", "ミラクルホープ", "ドリームチェイサー", "フェニックス"}

	for i, name := range names {
		DrawButton(screen, 250, float32(190+i*50), 300, 40, name, i == g.selectedIdx)
	}

	DrawText(screen, "あなたの愛馬:", 260, 450, ColorLightGray)
	DrawText(screen, "「"+names[g.selectedIdx]+"」", 380, 450, ColorGold)

	DrawTextCentered(screen, "↑↓: 選択  |  Enter: 決定", 400, 500, 14, ColorLightGray)
}

// === メイン画面 ===
func (g *Game) updateMain() {
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
		g.message = g.playerHorse.Rest()
		g.advanceWeek()
	}
}

func (g *Game) drawMain(screen *ebiten.Image) {
	screen.Fill(ColorDarkGreen)

	// ヘッダー
	vector.DrawFilledRect(screen, 0, 0, 800, 50, ColorPanelBg, false)
	vector.DrawFilledRect(screen, 0, 48, 800, 2, ColorGold, false)

	DrawText(screen, fmt.Sprintf("【 %d年目 第%d週 】", g.year, g.week), 20, 18, ColorWhite)
	DrawText(screen, fmt.Sprintf("資金: %d円", g.money), 220, 18, ColorGold)
	remainingWeeks := (3-g.year)*52 + (52 - g.week)
	DrawText(screen, fmt.Sprintf("残り: %d週", remainingWeeks), 420, 18, ColorLightGray)

	// 愛馬ステータスパネル
	DrawPanel(screen, 20, 65, 460, 280)
	DrawTextCentered(screen, "愛馬ステータス", 250, 80, 18, ColorGold)

	if g.playerHorse != nil {
		DrawText(screen, fmt.Sprintf("【 %s 】  %d歳", g.playerHorse.Name, g.playerHorse.Age), 40, 110, ColorWhite)

		// ステータスバー
		DrawStatusBar(screen, 40, 145, 130, g.playerHorse.Speed, 100, ColorRed, "スピード")
		DrawStatusBar(screen, 40, 172, 130, g.playerHorse.Stamina, 100, ColorGreen, "スタミナ")
		DrawStatusBar(screen, 40, 199, 130, g.playerHorse.Power, 100, ColorBlue, "パワー  ")
		DrawStatusBar(screen, 40, 226, 130, g.playerHorse.Guts, 100, ColorOrange, "根性    ")
		DrawStatusBar(screen, 40, 253, 130, g.playerHorse.Wisdom, 100, ColorPurple, "賢さ    ")

		// 右側の情報
		DrawText(screen, fmt.Sprintf("総合値: %d", g.playerHorse.TotalStats()), 290, 145, ColorGold)

		moodText := "普通"
		moodColor := ColorWhite
		if g.playerHorse.Mood >= 80 {
			moodText = "絶好調！"
			moodColor = ColorGold
		} else if g.playerHorse.Mood >= 60 {
			moodText = "好調"
			moodColor = ColorGreen
		} else if g.playerHorse.Mood < 40 {
			moodText = "不調..."
			moodColor = ColorRed
		}
		DrawText(screen, fmt.Sprintf("調子: %s (%d%%)", moodText, g.playerHorse.Mood), 290, 180, moodColor)

		fatigueText := "元気"
		fatigueColor := ColorGreen
		if g.playerHorse.Fatigue >= 70 {
			fatigueText = "疲労困憊！"
			fatigueColor = ColorRed
		} else if g.playerHorse.Fatigue >= 50 {
			fatigueText = "やや疲れ"
			fatigueColor = ColorOrange
		} else if g.playerHorse.Fatigue >= 30 {
			fatigueText = "少し疲労"
			fatigueColor = ColorLightGray
		}
		DrawText(screen, fmt.Sprintf("疲労: %s (%d%%)", fatigueText, g.playerHorse.Fatigue), 290, 215, fatigueColor)

		DrawText(screen, fmt.Sprintf("戦績: %d戦 %d勝", g.playerHorse.Races, g.playerHorse.Wins), 290, 260, ColorWhite)
		if g.playerHorse.Races > 0 {
			winRate := float64(g.playerHorse.Wins) / float64(g.playerHorse.Races) * 100
			DrawText(screen, fmt.Sprintf("勝率: %.1f%%", winRate), 290, 290, ColorLightGray)
		}
	}

	// アクションパネル
	DrawPanel(screen, 500, 65, 280, 280)
	DrawTextCentered(screen, "アクション", 640, 80, 18, ColorGold)

	DrawText(screen, "[1] トレーニング", 520, 120, ColorWhite)
	DrawText(screen, "    能力を上げる", 520, 145, ColorLightGray)

	DrawText(screen, "[2] レースに出走", 520, 185, ColorWhite)
	DrawText(screen, "    賞金を獲得", 520, 210, ColorLightGray)

	DrawText(screen, "[3] 休養する", 520, 250, ColorWhite)
	DrawText(screen, "    疲労回復・調子UP", 520, 275, ColorLightGray)

	DrawTextCentered(screen, "数字キーで選択", 640, 310, 12, ColorLightGray)

	// メッセージパネル
	if g.message != "" {
		DrawPanel(screen, 20, 365, 760, 80)
		DrawTextCentered(screen, "【 メッセージ 】", 400, 380, 14, ColorGold)
		DrawText(screen, g.message, 40, 410, ColorWhite)
	}

	// ヒントパネル
	DrawPanel(screen, 20, 465, 760, 115)
	DrawTextCentered(screen, "【 ヒント 】", 400, 480, 14, ColorGold)
	if g.playerHorse != nil {
		var hint string
		if g.playerHorse.Fatigue >= 70 {
			hint = "疲労が溜まっています。[3]休養 で回復しましょう！"
		} else if g.playerHorse.TotalStats() < 200 {
			hint = "まだ能力が低いです。[1]トレーニング で鍛えましょう！"
		} else if g.playerHorse.Races == 0 {
			hint = "そろそろレースに挑戦してみましょう！[2]レース出走"
		} else {
			hint = "G1レースでの勝利を目指して頑張りましょう！"
		}
		DrawTextCentered(screen, hint, 400, 510, 16, ColorWhite)
		DrawTextCentered(screen, "目標: G1レース制覇！ (有馬記念、日本ダービーなど)", 400, 545, 14, ColorGold)
	}
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
	screen.Fill(ColorDarkGreen)
	DrawPanel(screen, 100, 50, 600, 500)

	DrawTextCentered(screen, "トレーニング選択", 400, 80, 24, ColorGold)

	trainings := []struct {
		name  string
		desc  string
		stat  int
		color color.RGBA
	}{
		{"スピード", "最高速度UP → レースの基本能力", g.playerHorse.Speed, ColorRed},
		{"スタミナ", "持久力UP   → 長距離で有利", g.playerHorse.Stamina, ColorGreen},
		{"パワー  ", "加速力UP   → スタート有利", g.playerHorse.Power, ColorBlue},
		{"根性    ", "粘り強さUP → 終盤の追い込み", g.playerHorse.Guts, ColorOrange},
		{"賢さ    ", "安定性UP   → 実力を発揮しやすい", g.playerHorse.Wisdom, ColorPurple},
	}

	for i, t := range trainings {
		y := float32(130 + i*65)
		selected := i == g.selectedIdx

		if selected {
			vector.DrawFilledRect(screen, 130, y, 540, 55, ColorGold, false)
			vector.DrawFilledRect(screen, 133, y+3, 534, 49, ColorDarkBrown, false)
		} else {
			vector.DrawFilledRect(screen, 130, y, 540, 55, ColorBrown, false)
			vector.DrawFilledRect(screen, 133, y+3, 534, 49, ColorPanelBg, false)
		}

		prefix := "  "
		if selected {
			prefix = "> "
		}

		vector.DrawFilledRect(screen, 150, y+12, 14, 14, t.color, false)
		DrawText(screen, prefix+t.name, 175, float64(y)+10, ColorWhite)
		DrawText(screen, t.desc, 300, float64(y)+10, ColorLightGray)

		// 現在値バー
		barWidth := float32(t.stat) * 1.8
		vector.DrawFilledRect(screen, 150, y+35, 180, 12, color.RGBA{30, 30, 30, 255}, false)
		vector.DrawFilledRect(screen, 150, y+35, barWidth, 12, t.color, false)
		DrawText(screen, fmt.Sprintf("%d", t.stat), 345, float64(y)+32, ColorWhite)
	}

	// 戻るボタン
	DrawButton(screen, 300, 460, 200, 35, "戻る", g.selectedIdx == 5)

	// コンディション
	fatigueWarning := ""
	if g.playerHorse.Fatigue >= 70 {
		fatigueWarning = " ⚠ 休養推奨！"
	}
	DrawTextCentered(screen, fmt.Sprintf("疲労: %d%%  調子: %d%%%s",
		g.playerHorse.Fatigue, g.playerHorse.Mood, fatigueWarning), 400, 510, 14, ColorOrange)

	DrawTextCentered(screen, "↑↓: 選択  |  Enter: 決定  |  Esc: 戻る", 400, 535, 12, ColorLightGray)
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
			g.raceState = NewRaceState(&race, g.playerHorse, 7)
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
	screen.Fill(ColorDarkGreen)
	DrawPanel(screen, 50, 30, 700, 540)

	DrawTextCentered(screen, "レ ー ス 選 択", 400, 55, 24, ColorGold)

	// ヘッダー
	DrawText(screen, "グレード      レース名              距離       賞金", 100, 90, ColorLightGray)

	for i, race := range AvailableRaces {
		y := float32(115 + i*38)
		selected := i == g.selectedIdx

		gradeColor := ColorWhite
		switch race.Grade {
		case GradeG1:
			gradeColor = ColorGold
		case GradeG2:
			gradeColor = color.RGBA{192, 192, 192, 255}
		case GradeG3:
			gradeColor = color.RGBA{205, 127, 50, 255}
		}

		if selected {
			vector.DrawFilledRect(screen, 80, y, 640, 33, ColorGold, false)
			vector.DrawFilledRect(screen, 83, y+2, 634, 29, ColorDarkBrown, false)
		}

		// グレードバッジ
		vector.DrawFilledRect(screen, 100, y+6, 40, 20, gradeColor, false)
		DrawText(screen, race.Grade.String(), 108, float64(y)+8, ColorPanelBg)

		prefix := "  "
		if selected {
			prefix = "> "
		}

		DrawText(screen, fmt.Sprintf("%s%-12s    %4dm    %6d円",
			prefix, race.Name, race.Distance, race.Grade.Prize()), 155, float64(y)+8, ColorWhite)
	}

	// 戻るボタン
	y := float32(115 + len(AvailableRaces)*38 + 10)
	DrawButton(screen, 300, y, 200, 35, "戻る", g.selectedIdx == len(AvailableRaces))

	// 選択中のレース情報
	if g.selectedIdx < len(AvailableRaces) {
		race := AvailableRaces[g.selectedIdx]
		vector.DrawFilledRect(screen, 80, 520, 640, 40, color.RGBA{0, 0, 0, 150}, false)
		difficulty := "普通"
		if race.Grade == GradeG1 {
			difficulty = "最高難度"
		} else if race.Grade == GradeG2 {
			difficulty = "やや難"
		}
		DrawText(screen, fmt.Sprintf("【%s】 難易度:%s | 1位:%d円 / 2位:%d円 / 3位:%d円",
			race.Name, difficulty, race.Grade.Prize(), race.Grade.Prize()/2, race.Grade.Prize()/4), 100, 530, ColorWhite)
	}

	DrawTextCentered(screen, "↑↓: 選択  |  Enter: 出走  |  Esc: 戻る", 400, 575, 12, ColorLightGray)
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
	// 背景
	screen.Fill(color.RGBA{100, 149, 237, 255})
	vector.DrawFilledRect(screen, 0, 100, 800, 350, ColorGreen, false)

	// コース
	vector.DrawFilledRect(screen, 30, 130, 740, 300, ColorDarkBrown, false)
	vector.DrawFilledRect(screen, 40, 140, 720, 280, ColorBrown, false)

	// レーン線
	for i := 0; i < 9; i++ {
		y := float32(140 + i*35)
		vector.DrawFilledRect(screen, 40, y, 720, 1, color.RGBA{255, 255, 255, 80}, false)
	}

	// ヘッダー
	vector.DrawFilledRect(screen, 0, 0, 800, 95, ColorPanelBg, false)
	vector.DrawFilledRect(screen, 0, 93, 800, 2, ColorGold, false)

	DrawTextCentered(screen, fmt.Sprintf("【 %s 】", g.raceState.Race.Name), 400, 15, 24, ColorGold)
	DrawTextCentered(screen, fmt.Sprintf("グレード: %s  |  距離: %dm  |  賞金: %d円",
		g.raceState.Race.Grade, g.raceState.Race.Distance, g.raceState.Race.Grade.Prize()), 400, 50, 14, ColorWhite)

	playerRank := g.raceState.GetPlayerRank()
	rankColor := ColorWhite
	rankText := fmt.Sprintf("%d位", playerRank)
	if playerRank <= 3 {
		rankText += " 入賞圏内！"
		rankColor = ColorGold
	}
	DrawTextCentered(screen, fmt.Sprintf("あなたの馬: %s", rankText), 400, 75, 16, rankColor)

	// 馬を描画
	for i, entry := range g.raceState.Entries {
		y := float32(150 + i*35)
		x := float32(50 + entry.Position*6.5)

		horseColor := ColorBrown
		if entry.Horse.IsPlayer {
			horseColor = ColorGold
		}

		g.drawHorseSprite(screen, x, y, horseColor)

		nameStr := entry.Horse.Name
		if entry.Horse.IsPlayer {
			nameStr = "★" + nameStr
		}
		vector.DrawFilledRect(screen, x-5, y-20, float32(len(nameStr)*8+10), 18, color.RGBA{0, 0, 0, 180}, false)
		DrawText(screen, nameStr, float64(x), float64(y)-18, ColorWhite)
	}

	// ゴールライン
	vector.DrawFilledRect(screen, 720, 130, 8, 300, ColorWhite, false)
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			vector.DrawFilledRect(screen, 720, float32(130+i*20), 8, 20, ColorRed, false)
		}
	}
	DrawText(screen, "GOAL", 715, 110, ColorWhite)

	// フッター
	vector.DrawFilledRect(screen, 0, 450, 800, 150, ColorPanelBg, false)
	DrawText(screen, "実況: レースが進行中です！各馬のポジション争いが激化！", 150, 475, ColorWhite)

	if g.raceState.IsFinished {
		vector.DrawFilledRect(screen, 0, 0, 800, 600, color.RGBA{0, 0, 0, 150}, false)
		DrawPanel(screen, 200, 200, 400, 150)
		DrawTextCentered(screen, "レース終了！", 400, 240, 24, ColorGold)

		playerEntry := g.raceState.GetPlayerEntry()
		if playerEntry != nil {
			resultText := fmt.Sprintf("%s は %d位！", g.playerHorse.Name, playerEntry.Rank)
			if playerEntry.Rank == 1 {
				resultText += " 優勝！"
			} else if playerEntry.Rank <= 3 {
				resultText += " 入賞！"
			}
			DrawTextCentered(screen, resultText, 400, 280, 18, ColorWhite)
		}
		DrawTextCentered(screen, "[ ENTER で結果確認 ]", 400, 320, 14, ColorLightGray)
	}
}

// === レース結果画面 ===
func (g *Game) updateRaceResult() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
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
			g.message = fmt.Sprintf("【%s】%s が %d位！ 賞金 %d円 を獲得！",
				g.raceState.Race.Name, g.playerHorse.Name, playerEntry.Rank, prize)
		} else if playerEntry != nil {
			g.message = fmt.Sprintf("【%s】%s は %d位でした。次こそは！",
				g.raceState.Race.Name, g.playerHorse.Name, playerEntry.Rank)
		}

		g.advanceWeek()
		g.scene = SceneMain
	}
}

func (g *Game) drawRaceResult(screen *ebiten.Image) {
	screen.Fill(ColorDarkGreen)
	DrawPanel(screen, 100, 50, 600, 500)

	DrawTextCentered(screen, fmt.Sprintf("%s  結果発表", g.raceState.Race.Name), 400, 80, 24, ColorGold)

	DrawText(screen, "順位    馬名                    賞金", 160, 120, ColorLightGray)

	for i, entry := range g.raceState.Results {
		y := float32(150 + i*38)

		if entry.Rank <= 3 {
			bgColor := color.RGBA{80, 60, 40, 200}
			if entry.Rank == 1 {
				bgColor = color.RGBA{255, 215, 0, 80}
			} else if entry.Rank == 2 {
				bgColor = color.RGBA{192, 192, 192, 80}
			} else if entry.Rank == 3 {
				bgColor = color.RGBA{205, 127, 50, 80}
			}
			vector.DrawFilledRect(screen, 140, y, 520, 33, bgColor, false)
		}

		if entry.Horse.IsPlayer {
			vector.DrawFilledRect(screen, 138, y-2, 524, 37, ColorGold, false)
			vector.DrawFilledRect(screen, 140, y, 520, 33, ColorDarkBrown, false)
		}

		marker := ""
		if entry.Horse.IsPlayer {
			marker = "★"
		}

		prize := ""
		if entry.Rank == 1 {
			prize = fmt.Sprintf("+%d円", g.raceState.Race.Grade.Prize())
		} else if entry.Rank == 2 {
			prize = fmt.Sprintf("+%d円", g.raceState.Race.Grade.Prize()/2)
		} else if entry.Rank == 3 {
			prize = fmt.Sprintf("+%d円", g.raceState.Race.Grade.Prize()/4)
		}

		DrawText(screen, fmt.Sprintf(" %d位  %s%-14s    %s",
			entry.Rank, marker, entry.Horse.Name, prize), 160, float64(y)+8, ColorWhite)
	}

	// プレイヤー結果
	playerEntry := g.raceState.GetPlayerEntry()
	if playerEntry != nil {
		vector.DrawFilledRect(screen, 140, 485, 520, 45, color.RGBA{0, 0, 0, 150}, false)
		resultMsg := fmt.Sprintf("あなたの馬「%s」は %d位でした！", g.playerHorse.Name, playerEntry.Rank)
		if playerEntry.Rank == 1 {
			resultMsg += " 優勝おめでとう！"
		} else if playerEntry.Rank <= 3 {
			resultMsg += " 入賞！賞金獲得！"
		}
		DrawTextCentered(screen, resultMsg, 400, 500, 16, ColorWhite)
	}

	DrawTextCentered(screen, "[ ENTER でメインに戻る ]", 400, 540, 14, ColorLightGray)
}

// === ゲームオーバー画面 ===
func (g *Game) updateGameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		*g = *NewGame()
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	DrawPanel(screen, 100, 50, 600, 500)

	DrawTextCentered(screen, "ゲ ー ム 終 了", 400, 85, 28, ColorGold)
	DrawTextCentered(screen, "3年間の育成が終わりました", 400, 120, 16, ColorLightGray)

	if g.playerHorse != nil {
		DrawTextCentered(screen, fmt.Sprintf("愛馬: 【 %s 】", g.playerHorse.Name), 400, 170, 20, ColorWhite)

		DrawText(screen, fmt.Sprintf("最終年齢: %d歳", g.playerHorse.Age), 280, 210, ColorWhite)
		DrawText(screen, fmt.Sprintf("最終戦績: %d戦 %d勝", g.playerHorse.Races, g.playerHorse.Wins), 280, 245, ColorWhite)

		winRate := 0.0
		if g.playerHorse.Races > 0 {
			winRate = float64(g.playerHorse.Wins) / float64(g.playerHorse.Races) * 100
		}
		DrawText(screen, fmt.Sprintf("勝率: %.1f%%", winRate), 280, 280, ColorWhite)
		DrawText(screen, fmt.Sprintf("最終資金: %d円", g.money), 280, 315, ColorGold)

		// 評価
		var rating string
		var ratingDesc string
		ratingColor := ColorWhite
		if g.playerHorse.Wins >= 10 {
			rating = "【 伝説の名馬 】"
			ratingDesc = "歴史に名を刻む偉業を達成！"
			ratingColor = ColorGold
		} else if g.playerHorse.Wins >= 5 {
			rating = "【 名馬 】"
			ratingDesc = "素晴らしい活躍でした！"
			ratingColor = ColorGold
		} else if g.playerHorse.Wins >= 3 {
			rating = "【 優秀 】"
			ratingDesc = "立派な成績です！"
			ratingColor = ColorGreen
		} else if g.playerHorse.Wins >= 1 {
			rating = "【 及第点 】"
			ratingDesc = "勝利の喜びを味わえました"
			ratingColor = ColorLightGray
		} else {
			rating = "【 発展途上 】"
			ratingDesc = "次こそは頑張りましょう..."
			ratingColor = ColorRed
		}
		DrawTextCentered(screen, rating, 400, 380, 24, ratingColor)
		DrawTextCentered(screen, ratingDesc, 400, 420, 16, ColorWhite)
	}

	if (g.animationTick/30)%2 == 0 {
		DrawTextCentered(screen, "[ ENTER でもう一度プレイ ]", 400, 500, 16, ColorGold)
	}
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

	if g.playerHorse != nil {
		moodChange := rand.Intn(11) - 5
		g.playerHorse.Mood += moodChange
		if g.playerHorse.Mood > 100 {
			g.playerHorse.Mood = 100
		}
		if g.playerHorse.Mood < 20 {
			g.playerHorse.Mood = 20
		}
	}
}
