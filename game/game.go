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

// === UI描画ヘルパー ===

// drawPanel は装飾付きパネルを描画
func drawPanel(screen *ebiten.Image, x, y, w, h float32) {
	// 外枠（ゴールド）
	vector.DrawFilledRect(screen, x, y, w, h, ColorPanelFrame, false)
	// 内側の背景
	vector.DrawFilledRect(screen, x+3, y+3, w-6, h-6, ColorPanelBg, false)
	// 内側の装飾線
	vector.DrawFilledRect(screen, x+6, y+6, w-12, 2, ColorDarkGold, false)
	vector.DrawFilledRect(screen, x+6, y+h-8, w-12, 2, ColorDarkGold, false)
}

// drawButton は選択可能なボタンを描画
func drawButton(screen *ebiten.Image, x, y, w, h float32, text string, selected bool) {
	if selected {
		vector.DrawFilledRect(screen, x, y, w, h, ColorGold, false)
		vector.DrawFilledRect(screen, x+2, y+2, w-4, h-4, ColorDarkBrown, false)
		ebitenutil.DebugPrintAt(screen, "> "+text, int(x)+10, int(y)+8)
	} else {
		vector.DrawFilledRect(screen, x, y, w, h, ColorBrown, false)
		vector.DrawFilledRect(screen, x+2, y+2, w-4, h-4, ColorPanelBg, false)
		ebitenutil.DebugPrintAt(screen, "  "+text, int(x)+10, int(y)+8)
	}
}

// drawStatusBar はステータスバーを描画
func drawStatusBar(screen *ebiten.Image, x, y, w float32, value int, maxValue int, barColor color.RGBA, label string) {
	// ラベル
	ebitenutil.DebugPrintAt(screen, label, int(x), int(y)-2)
	// 背景バー
	vector.DrawFilledRect(screen, x+70, y, w, 14, color.RGBA{30, 30, 30, 255}, false)
	// 値バー
	fillWidth := (float32(value) / float32(maxValue)) * w
	vector.DrawFilledRect(screen, x+70, y, fillWidth, 14, barColor, false)
	// 数値
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", value), int(x+w+80), int(y)-2)
}

// === タイトル画面 ===
func (g *Game) updateTitle() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.scene = SceneTutorial
		g.tutorialPage = 0
	}
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	// 背景グラデーション風
	screen.Fill(ColorDarkGreen)
	vector.DrawFilledRect(screen, 0, 0, 800, 200, color.RGBA{0, 60, 0, 255}, false)
	vector.DrawFilledRect(screen, 0, 400, 800, 200, color.RGBA{0, 80, 0, 255}, false)

	// 競馬場のトラック（装飾）
	vector.DrawFilledRect(screen, 50, 180, 700, 120, ColorBrown, false)
	vector.DrawFilledRect(screen, 60, 190, 680, 100, ColorDarkBrown, false)
	// 白線
	for i := 0; i < 7; i++ {
		vector.DrawFilledRect(screen, float32(100+i*100), 235, 50, 3, ColorWhite, false)
	}

	// 走る馬のアニメーション
	horseX := float32((g.animationTick * 3) % 900)
	g.drawHorseSprite(screen, horseX-100, 200, ColorGold)
	g.drawHorseSprite(screen, horseX-250, 210, ColorBrown)
	g.drawHorseSprite(screen, horseX-180, 220, ColorDarkBrown)

	// メインタイトルパネル
	drawPanel(screen, 150, 320, 500, 250)

	// タイトルロゴ
	ebitenutil.DebugPrintAt(screen, "╔═══════════════════════════════════════════╗", 170, 340)
	ebitenutil.DebugPrintAt(screen, "║                                           ║", 170, 355)
	ebitenutil.DebugPrintAt(screen, "║      DERBY DREAM  ダービードリーム        ║", 170, 370)
	ebitenutil.DebugPrintAt(screen, "║                                           ║", 170, 385)
	ebitenutil.DebugPrintAt(screen, "║        ～ 競走馬育成シミュレーション ～     ║", 170, 400)
	ebitenutil.DebugPrintAt(screen, "║                                           ║", 170, 415)
	ebitenutil.DebugPrintAt(screen, "╚═══════════════════════════════════════════╝", 170, 430)

	// 馬のASCIIアート
	ebitenutil.DebugPrintAt(screen, "       ,--,", 200, 460)
	ebitenutil.DebugPrintAt(screen, "      (o o)", 200, 475)
	ebitenutil.DebugPrintAt(screen, " /---( )--/\\", 200, 490)
	ebitenutil.DebugPrintAt(screen, " `===(_)=(_)~", 200, 505)

	// スタートメッセージ（点滅）
	if (g.animationTick/30)%2 == 0 {
		ebitenutil.DebugPrintAt(screen, ">>> PRESS ENTER OR SPACE TO START <<<", 250, 540)
	}
}

// drawHorseSprite は簡易的な馬のスプライトを描画
func (g *Game) drawHorseSprite(screen *ebiten.Image, x, y float32, c color.RGBA) {
	// 胴体
	vector.DrawFilledRect(screen, x, y, 40, 20, c, false)
	// 頭
	vector.DrawFilledRect(screen, x+35, y-5, 15, 15, c, false)
	// 足（アニメーション）
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
	drawPanel(screen, 50, 50, 700, 500)

	// ページインジケーター
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 ゲームの遊び方 】 %d / 4", g.tutorialPage+1), 300, 70)

	// 進捗バー
	for i := 0; i < 4; i++ {
		c := color.RGBA{60, 60, 60, 255}
		if i <= g.tutorialPage {
			c = ColorGold
		}
		vector.DrawFilledRect(screen, float32(300+i*50), 95, 40, 8, c, false)
	}

	// ページ内容
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

	// ナビゲーション
	ebitenutil.DebugPrintAt(screen, "← →: ページ移動  |  Enter/Space: 次へ  |  Esc: スキップ", 200, 520)
}

func (g *Game) drawTutorialPage1(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "◆ ゲームの目的 ◆", 330, 130)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 150)

	ebitenutil.DebugPrintAt(screen, "あなたは新人馬主です。", 200, 180)
	ebitenutil.DebugPrintAt(screen, "3年間で愛馬を育て、G1レースでの勝利を目指しましょう！", 200, 210)

	ebitenutil.DebugPrintAt(screen, "【 勝利条件 】", 200, 260)
	ebitenutil.DebugPrintAt(screen, "・できるだけ多くのレースに勝利する", 220, 290)
	ebitenutil.DebugPrintAt(screen, "・賞金を稼いで資金を増やす", 220, 320)
	ebitenutil.DebugPrintAt(screen, "・G1レース（最高峰）での勝利を目指す", 220, 350)

	ebitenutil.DebugPrintAt(screen, "期間: 3年間（1年 = 52週）", 200, 400)
	ebitenutil.DebugPrintAt(screen, "毎週1つのアクションを選んで進めます", 200, 430)
}

func (g *Game) drawTutorialPage2(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "◆ 育成システム ◆", 330, 130)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 150)

	ebitenutil.DebugPrintAt(screen, "馬には5つのステータスがあります：", 200, 180)

	stats := []struct {
		name  string
		desc  string
		color color.RGBA
	}{
		{"スピード", "最高速度に影響。レースの基本となる能力", ColorRed},
		{"スタミナ", "持久力。長距離レースで重要", ColorGreen},
		{"パワー", "加速力。スタートダッシュに影響", ColorBlue},
		{"根性", "終盤の粘り強さ。逆転に必要", color.RGBA{255, 165, 0, 255}},
		{"賢さ", "レース運びの安定性", color.RGBA{148, 0, 211, 255}},
	}

	for i, s := range stats {
		y := float32(220 + i*50)
		vector.DrawFilledRect(screen, 200, y, 15, 15, s.color, false)
		ebitenutil.DebugPrintAt(screen, s.name, 225, int(y)-2)
		ebitenutil.DebugPrintAt(screen, s.desc, 320, int(y)-2)
	}

	ebitenutil.DebugPrintAt(screen, "※ トレーニングで各ステータスを上げられます", 200, 480)
}

func (g *Game) drawTutorialPage3(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "◆ 毎週のアクション ◆", 320, 130)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 150)

	ebitenutil.DebugPrintAt(screen, "毎週、以下の3つから1つを選びます：", 200, 180)

	ebitenutil.DebugPrintAt(screen, "[1] トレーニング", 200, 230)
	ebitenutil.DebugPrintAt(screen, "    → ステータスを上げる（疲労が溜まる）", 200, 255)

	ebitenutil.DebugPrintAt(screen, "[2] レース出走", 200, 300)
	ebitenutil.DebugPrintAt(screen, "    → 勝てば賞金獲得！（疲労が溜まる）", 200, 325)

	ebitenutil.DebugPrintAt(screen, "[3] 休養", 200, 370)
	ebitenutil.DebugPrintAt(screen, "    → 疲労を回復、調子を整える", 200, 395)

	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 430)
	ebitenutil.DebugPrintAt(screen, "※ 疲労が高いと調子が下がりやすくなります！", 200, 460)
	ebitenutil.DebugPrintAt(screen, "   バランスよく育成と休養を行いましょう", 200, 485)
}

func (g *Game) drawTutorialPage4(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "◆ レースについて ◆", 325, 130)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 150)

	ebitenutil.DebugPrintAt(screen, "レースには3つのグレードがあります：", 200, 180)

	ebitenutil.DebugPrintAt(screen, "【G3】 入門レベル   賞金: 3,000円", 220, 220)
	ebitenutil.DebugPrintAt(screen, "【G2】 中級レベル   賞金: 8,000円", 220, 250)
	ebitenutil.DebugPrintAt(screen, "【G1】 最高峰！     賞金: 20,000円", 220, 280)

	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 320)
	ebitenutil.DebugPrintAt(screen, "【 賞金 】", 200, 350)
	ebitenutil.DebugPrintAt(screen, "1位: 100%  |  2位: 50%  |  3位: 25%", 220, 380)

	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 150, 420)
	ebitenutil.DebugPrintAt(screen, "準備はいいですか？ 愛馬の名前を決めましょう！", 200, 460)
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
	drawPanel(screen, 150, 80, 500, 440)

	ebitenutil.DebugPrintAt(screen, "╔════════════════════════════════════╗", 180, 100)
	ebitenutil.DebugPrintAt(screen, "║      愛馬の名前を選んでください      ║", 180, 115)
	ebitenutil.DebugPrintAt(screen, "╚════════════════════════════════════╝", 180, 130)

	ebitenutil.DebugPrintAt(screen, "この馬と3年間の旅を共にします", 250, 165)

	names := []string{"ゴールドスター", "サンダーボルト", "ミラクルホープ", "ドリームチェイサー", "フェニックス"}

	for i, name := range names {
		drawButton(screen, 250, float32(210+i*45), 300, 35, name, i == g.selectedIdx)
	}

	// 選択中の馬のイメージ
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━", 250, 440)
	ebitenutil.DebugPrintAt(screen, "あなたの愛馬:", 280, 460)
	ebitenutil.DebugPrintAt(screen, "「"+names[g.selectedIdx]+"」", 380, 460)

	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  |  Enter: 決定", 290, 500)
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

	// ヘッダー情報バー
	vector.DrawFilledRect(screen, 0, 0, 800, 45, ColorPanelBg, false)
	vector.DrawFilledRect(screen, 0, 43, 800, 2, ColorGold, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %d年目 第%d週 】", g.year, g.week), 20, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("資金: %d円", g.money), 250, 15)
	// 残り週数表示
	remainingWeeks := (3-g.year)*52 + (52 - g.week)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("残り: %d週", remainingWeeks), 450, 15)

	// 愛馬ステータスパネル
	drawPanel(screen, 20, 60, 450, 280)
	ebitenutil.DebugPrintAt(screen, "━━━ 愛馬ステータス ━━━", 140, 75)

	if g.playerHorse != nil {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %s 】  %d歳", g.playerHorse.Name, g.playerHorse.Age), 50, 105)

		// ステータスバー
		drawStatusBar(screen, 50, 140, 120, g.playerHorse.Speed, 100, ColorRed, "スピード")
		drawStatusBar(screen, 50, 165, 120, g.playerHorse.Stamina, 100, ColorGreen, "スタミナ")
		drawStatusBar(screen, 50, 190, 120, g.playerHorse.Power, 100, ColorBlue, "パワー  ")
		drawStatusBar(screen, 50, 215, 120, g.playerHorse.Guts, 100, color.RGBA{255, 165, 0, 255}, "根性    ")
		drawStatusBar(screen, 50, 240, 120, g.playerHorse.Wisdom, 100, color.RGBA{148, 0, 211, 255}, "賢さ    ")

		// 総合値
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("総合値: %d", g.playerHorse.TotalStats()), 280, 140)

		// コンディション
		ebitenutil.DebugPrintAt(screen, "──────────────", 280, 165)
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
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("調子: %s (%d%%)", moodText, g.playerHorse.Mood), 280, 185)
		_ = moodColor // 色は将来的に使用

		// 疲労警告
		fatigueText := "元気"
		if g.playerHorse.Fatigue >= 70 {
			fatigueText = "疲労困憊！休養推奨"
		} else if g.playerHorse.Fatigue >= 50 {
			fatigueText = "やや疲れ気味"
		} else if g.playerHorse.Fatigue >= 30 {
			fatigueText = "少し疲労"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("疲労: %s (%d%%)", fatigueText, g.playerHorse.Fatigue), 280, 210)

		// 戦績
		ebitenutil.DebugPrintAt(screen, "──────────────", 280, 235)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("戦績: %d戦 %d勝", g.playerHorse.Races, g.playerHorse.Wins), 280, 255)
		if g.playerHorse.Races > 0 {
			winRate := float64(g.playerHorse.Wins) / float64(g.playerHorse.Races) * 100
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("勝率: %.1f%%", winRate), 280, 280)
		}
	}

	// アクションパネル
	drawPanel(screen, 490, 60, 290, 280)
	ebitenutil.DebugPrintAt(screen, "━━━ アクション ━━━", 550, 75)

	ebitenutil.DebugPrintAt(screen, "[1] トレーニング", 510, 115)
	ebitenutil.DebugPrintAt(screen, "    能力を上げる", 510, 135)

	ebitenutil.DebugPrintAt(screen, "[2] レースに出走", 510, 175)
	ebitenutil.DebugPrintAt(screen, "    賞金を獲得", 510, 195)

	ebitenutil.DebugPrintAt(screen, "[3] 休養する", 510, 235)
	ebitenutil.DebugPrintAt(screen, "    疲労回復・調子UP", 510, 255)

	ebitenutil.DebugPrintAt(screen, "──────────────────", 510, 285)
	ebitenutil.DebugPrintAt(screen, "数字キーで選択", 550, 305)

	// メッセージパネル
	if g.message != "" {
		drawPanel(screen, 20, 360, 760, 80)
		ebitenutil.DebugPrintAt(screen, "【 メッセージ 】", 350, 375)
		ebitenutil.DebugPrintAt(screen, g.message, 40, 405)
	}

	// ヒント
	drawPanel(screen, 20, 460, 760, 120)
	ebitenutil.DebugPrintAt(screen, "【 ヒント 】", 365, 475)
	if g.playerHorse != nil {
		if g.playerHorse.Fatigue >= 70 {
			ebitenutil.DebugPrintAt(screen, "疲労が溜まっています。[3]休養 で回復しましょう！", 150, 505)
		} else if g.playerHorse.TotalStats() < 200 {
			ebitenutil.DebugPrintAt(screen, "まだ能力が低いです。[1]トレーニング で鍛えましょう！", 150, 505)
		} else if g.playerHorse.Races == 0 {
			ebitenutil.DebugPrintAt(screen, "そろそろレースに挑戦してみましょう！[2]レース出走", 150, 505)
		} else {
			ebitenutil.DebugPrintAt(screen, "G1レースでの勝利を目指して頑張りましょう！", 180, 505)
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("目標: G1レース制覇！ (有馬記念、日本ダービーなど)"), 180, 535)
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
	drawPanel(screen, 100, 50, 600, 500)

	ebitenutil.DebugPrintAt(screen, "╔════════════════════════════════════════╗", 150, 70)
	ebitenutil.DebugPrintAt(screen, "║          トレーニング選択              ║", 150, 85)
	ebitenutil.DebugPrintAt(screen, "╚════════════════════════════════════════╝", 150, 100)

	trainings := []struct {
		name   string
		desc   string
		stat   int
		color  color.RGBA
	}{
		{"スピード", "最高速度UP → レースの基本能力", g.playerHorse.Speed, ColorRed},
		{"スタミナ", "持久力UP   → 長距離で有利", g.playerHorse.Stamina, ColorGreen},
		{"パワー", "加速力UP   → スタート有利", g.playerHorse.Power, ColorBlue},
		{"根性", "粘り強さUP → 終盤の追い込み", g.playerHorse.Guts, color.RGBA{255, 165, 0, 255}},
		{"賢さ", "安定性UP   → 実力を発揮しやすい", g.playerHorse.Wisdom, color.RGBA{148, 0, 211, 255}},
	}

	for i, t := range trainings {
		y := float32(140 + i*60)
		selected := i == g.selectedIdx

		if selected {
			vector.DrawFilledRect(screen, 130, y, 540, 50, ColorGold, false)
			vector.DrawFilledRect(screen, 133, y+3, 534, 44, ColorDarkBrown, false)
		} else {
			vector.DrawFilledRect(screen, 130, y, 540, 50, ColorBrown, false)
			vector.DrawFilledRect(screen, 133, y+3, 534, 44, ColorPanelBg, false)
		}

		prefix := "  "
		if selected {
			prefix = "> "
		}

		// ステータス名と色インジケーター
		vector.DrawFilledRect(screen, 145, y+10, 12, 12, t.color, false)
		ebitenutil.DebugPrintAt(screen, prefix+t.name, 165, int(y)+8)
		ebitenutil.DebugPrintAt(screen, t.desc, 280, int(y)+8)

		// 現在値バー
		barWidth := float32(t.stat) * 1.5
		vector.DrawFilledRect(screen, 145, y+30, 150, 10, color.RGBA{30, 30, 30, 255}, false)
		vector.DrawFilledRect(screen, 145, y+30, barWidth, 10, t.color, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", t.stat), 305, int(y)+26)
	}

	// 戻るボタン
	drawButton(screen, 300, 450, 200, 35, "戻る", g.selectedIdx == 5)

	// コンディション表示
	vector.DrawFilledRect(screen, 130, 500, 540, 40, color.RGBA{0, 0, 0, 100}, false)
	fatigueWarning := ""
	if g.playerHorse.Fatigue >= 70 {
		fatigueWarning = " ⚠ 疲労が高いです！休養を推奨"
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("現在の疲労: %d%%  調子: %d%%%s",
		g.playerHorse.Fatigue, g.playerHorse.Mood, fatigueWarning), 150, 510)

	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  |  Enter: 決定  |  Esc: 戻る", 260, 555)
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
	drawPanel(screen, 50, 30, 700, 540)

	ebitenutil.DebugPrintAt(screen, "╔════════════════════════════════════════════════════╗", 100, 50)
	ebitenutil.DebugPrintAt(screen, "║              レ ー ス 選 択                        ║", 100, 65)
	ebitenutil.DebugPrintAt(screen, "╚════════════════════════════════════════════════════╝", 100, 80)

	// ヘッダー
	ebitenutil.DebugPrintAt(screen, "グレード    レース名          距離     賞金", 120, 110)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 100, 125)

	for i, race := range AvailableRaces {
		y := float32(140 + i*35)
		selected := i == g.selectedIdx

		// グレードに応じた色
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
			vector.DrawFilledRect(screen, 80, y, 640, 30, ColorGold, false)
			vector.DrawFilledRect(screen, 83, y+2, 634, 26, ColorDarkBrown, false)
		}

		prefix := "  "
		if selected {
			prefix = "> "
		}

		// グレードバッジ
		vector.DrawFilledRect(screen, 95, y+5, 35, 18, gradeColor, false)
		ebitenutil.DebugPrintAt(screen, race.Grade.String(), 100, int(y)+6)

		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%-14s  %4dm   %6d円",
			prefix, race.Name, race.Distance, race.Grade.Prize()), 140, int(y)+6)
	}

	// 戻るボタン
	y := float32(140 + len(AvailableRaces)*35 + 10)
	drawButton(screen, 300, y, 200, 35, "戻る", g.selectedIdx == len(AvailableRaces))

	// 選択中のレース情報
	if g.selectedIdx < len(AvailableRaces) {
		race := AvailableRaces[g.selectedIdx]
		vector.DrawFilledRect(screen, 80, 515, 640, 45, color.RGBA{0, 0, 0, 150}, false)
		difficulty := "普通"
		if race.Grade == GradeG1 {
			difficulty = "最高難度"
		} else if race.Grade == GradeG2 {
			difficulty = "やや難しい"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【%s】 難易度: %s | 賞金: 1位 %d円 / 2位 %d円 / 3位 %d円",
			race.Name, difficulty, race.Grade.Prize(), race.Grade.Prize()/2, race.Grade.Prize()/4), 100, 525)
	}

	ebitenutil.DebugPrintAt(screen, "↑↓: 選択  |  Enter: 出走  |  Esc: 戻る", 270, 570)
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
	// 背景（観客席風）
	screen.Fill(color.RGBA{100, 149, 237, 255})
	vector.DrawFilledRect(screen, 0, 100, 800, 350, ColorGreen, false)

	// コース
	vector.DrawFilledRect(screen, 30, 130, 740, 300, ColorDarkBrown, false)
	vector.DrawFilledRect(screen, 40, 140, 720, 280, ColorBrown, false)

	// レーン区切り線
	for i := 0; i < 9; i++ {
		y := float32(140 + i*35)
		vector.DrawFilledRect(screen, 40, y, 720, 1, color.RGBA{255, 255, 255, 100}, false)
	}

	// レース情報ヘッダー
	vector.DrawFilledRect(screen, 0, 0, 800, 90, ColorPanelBg, false)
	vector.DrawFilledRect(screen, 0, 88, 800, 2, ColorGold, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("【 %s 】", g.raceState.Race.Name), 320, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("グレード: %s  |  距離: %dm  |  賞金: %d円",
		g.raceState.Race.Grade, g.raceState.Race.Distance, g.raceState.Race.Grade.Prize()), 200, 45)

	// 現在順位
	playerRank := g.raceState.GetPlayerRank()
	rankText := fmt.Sprintf("%d位", playerRank)
	if playerRank <= 3 {
		rankText += " 入賞圏内！"
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("あなたの馬: %s", rankText), 300, 70)

	// 馬を描画
	for i, entry := range g.raceState.Entries {
		y := float32(150 + i*35)
		x := float32(50 + entry.Position*6.5)

		// 馬の色
		horseColor := ColorBrown
		if entry.Horse.IsPlayer {
			horseColor = ColorGold
		}

		// 馬を描画
		g.drawHorseSprite(screen, x, y, horseColor)

		// 名前（プレイヤーの馬は強調）
		nameStr := entry.Horse.Name
		if entry.Horse.IsPlayer {
			nameStr = "★" + nameStr
		}
		// 名前の背景
		vector.DrawFilledRect(screen, x-5, y-18, float32(len(nameStr)*7+10), 14, color.RGBA{0, 0, 0, 150}, false)
		ebitenutil.DebugPrintAt(screen, nameStr, int(x), int(y)-17)
	}

	// ゴールライン
	vector.DrawFilledRect(screen, 720, 130, 8, 300, ColorWhite, false)
	vector.DrawFilledRect(screen, 720, 130, 8, 10, ColorRed, false)
	vector.DrawFilledRect(screen, 720, 150, 8, 10, ColorRed, false)
	vector.DrawFilledRect(screen, 720, 170, 8, 10, ColorRed, false)
	ebitenutil.DebugPrintAt(screen, "GOAL", 715, 115)

	// フッター
	vector.DrawFilledRect(screen, 0, 450, 800, 150, ColorPanelBg, false)
	ebitenutil.DebugPrintAt(screen, "実況: レースが進行中です！各馬のポジション争いが激化！", 150, 470)

	if g.raceState.IsFinished {
		// 終了オーバーレイ
		vector.DrawFilledRect(screen, 0, 0, 800, 600, color.RGBA{0, 0, 0, 150}, false)
		drawPanel(screen, 200, 200, 400, 150)
		ebitenutil.DebugPrintAt(screen, "━━━ レース終了！ ━━━", 305, 230)

		playerEntry := g.raceState.GetPlayerEntry()
		if playerEntry != nil {
			resultText := fmt.Sprintf("%s は %d位！", g.playerHorse.Name, playerEntry.Rank)
			if playerEntry.Rank == 1 {
				resultText += " おめでとう！"
			} else if playerEntry.Rank <= 3 {
				resultText += " 入賞！"
			}
			ebitenutil.DebugPrintAt(screen, resultText, 250, 270)
		}
		ebitenutil.DebugPrintAt(screen, "[ ENTER で結果確認 ]", 300, 320)
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
	drawPanel(screen, 100, 50, 600, 500)

	ebitenutil.DebugPrintAt(screen, "╔════════════════════════════════════════════╗", 140, 70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("║          %s  結果発表          ║", g.raceState.Race.Name), 140, 85)
	ebitenutil.DebugPrintAt(screen, "╚════════════════════════════════════════════╝", 140, 100)

	// 上位入賞を強調
	ebitenutil.DebugPrintAt(screen, "順位    馬名                  賞金", 180, 130)
	ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 160, 145)

	for i, entry := range g.raceState.Results {
		y := float32(160 + i*35)

		// 入賞は背景色を付ける
		if entry.Rank <= 3 {
			bgColor := color.RGBA{80, 60, 40, 200}
			if entry.Rank == 1 {
				bgColor = color.RGBA{255, 215, 0, 100}
			} else if entry.Rank == 2 {
				bgColor = color.RGBA{192, 192, 192, 100}
			} else if entry.Rank == 3 {
				bgColor = color.RGBA{205, 127, 50, 100}
			}
			vector.DrawFilledRect(screen, 150, y, 500, 30, bgColor, false)
		}

		// プレイヤーの馬を強調
		if entry.Horse.IsPlayer {
			vector.DrawFilledRect(screen, 148, y-2, 504, 34, ColorGold, false)
			vector.DrawFilledRect(screen, 150, y, 500, 30, ColorDarkBrown, false)
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

		rankStr := fmt.Sprintf("%d位", entry.Rank)
		if entry.Rank == 1 {
			rankStr = "1位"
		}

		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(" %s  %s%-16s  %s",
			rankStr, marker, entry.Horse.Name, prize), 160, int(y)+5)
	}

	// プレイヤーの結果サマリー
	playerEntry := g.raceState.GetPlayerEntry()
	if playerEntry != nil {
		vector.DrawFilledRect(screen, 150, 480, 500, 50, color.RGBA{0, 0, 0, 150}, false)
		resultMsg := fmt.Sprintf("あなたの馬「%s」は %d位でした！", g.playerHorse.Name, playerEntry.Rank)
		if playerEntry.Rank == 1 {
			resultMsg += " 優勝おめでとう！"
		} else if playerEntry.Rank <= 3 {
			resultMsg += " 入賞！賞金獲得！"
		}
		ebitenutil.DebugPrintAt(screen, resultMsg, 170, 495)
	}

	ebitenutil.DebugPrintAt(screen, "[ ENTER でメインに戻る ]", 300, 540)
}

// === ゲームオーバー画面 ===
func (g *Game) updateGameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		*g = *NewGame()
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	drawPanel(screen, 100, 50, 600, 500)

	ebitenutil.DebugPrintAt(screen, "╔════════════════════════════════════════════╗", 140, 70)
	ebitenutil.DebugPrintAt(screen, "║            ゲ ー ム 終 了                  ║", 140, 85)
	ebitenutil.DebugPrintAt(screen, "║          3年間の育成が終わりました          ║", 140, 100)
	ebitenutil.DebugPrintAt(screen, "╚════════════════════════════════════════════╝", 140, 115)

	if g.playerHorse != nil {
		// 馬のASCIIアート
		ebitenutil.DebugPrintAt(screen, "       ,--,", 340, 150)
		ebitenutil.DebugPrintAt(screen, "      (o o)", 340, 165)
		ebitenutil.DebugPrintAt(screen, " /---( )--/\\", 340, 180)
		ebitenutil.DebugPrintAt(screen, " `===(_)=(_)~", 340, 195)

		ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━", 250, 220)

		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("愛馬: 【 %s 】", g.playerHorse.Name), 280, 250)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("最終年齢: %d歳", g.playerHorse.Age), 300, 280)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("最終戦績: %d戦 %d勝", g.playerHorse.Races, g.playerHorse.Wins), 300, 310)

		winRate := 0.0
		if g.playerHorse.Races > 0 {
			winRate = float64(g.playerHorse.Wins) / float64(g.playerHorse.Races) * 100
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("勝率: %.1f%%", winRate), 300, 340)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("最終資金: %d円", g.money), 300, 370)

		// 評価
		ebitenutil.DebugPrintAt(screen, "━━━━━━━━━━━━━━━━━━━━━━━━━━", 250, 400)
		var rating string
		var ratingDesc string
		if g.playerHorse.Wins >= 10 {
			rating = "【 伝説の名馬 】"
			ratingDesc = "歴史に名を刻む偉業を達成！"
		} else if g.playerHorse.Wins >= 5 {
			rating = "【 名馬 】"
			ratingDesc = "素晴らしい活躍でした！"
		} else if g.playerHorse.Wins >= 3 {
			rating = "【 優秀 】"
			ratingDesc = "立派な成績です！"
		} else if g.playerHorse.Wins >= 1 {
			rating = "【 及第点 】"
			ratingDesc = "勝利の喜びを味わえました"
		} else {
			rating = "【 発展途上 】"
			ratingDesc = "次こそは頑張りましょう..."
		}
		ebitenutil.DebugPrintAt(screen, rating, 330, 430)
		ebitenutil.DebugPrintAt(screen, ratingDesc, 290, 460)
	}

	// 点滅
	if (g.animationTick/30)%2 == 0 {
		ebitenutil.DebugPrintAt(screen, "[ ENTER でもう一度プレイ ]", 290, 520)
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
