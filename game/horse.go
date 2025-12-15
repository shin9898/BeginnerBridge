package game

import (
	"fmt"
	"math/rand"
)

// Horse は競走馬を表す構造体
type Horse struct {
	Name      string
	Age       int // 2歳から5歳まで
	Speed     int // スピード (1-100)
	Stamina   int // スタミナ (1-100)
	Power     int // パワー (1-100)
	Guts      int // 根性 (1-100)
	Wisdom    int // 賢さ (1-100)
	Mood      int // 調子 (1-100)
	Fatigue   int // 疲労 (0-100)
	Wins      int
	Races     int
	IsPlayer  bool
}

// NewHorse は新しい馬を生成する
func NewHorse(name string, isPlayer bool) *Horse {
	baseStats := 30
	if !isPlayer {
		baseStats = rand.Intn(40) + 20 // ライバル馬は20-60のランダム
	}

	return &Horse{
		Name:     name,
		Age:      2,
		Speed:    baseStats + rand.Intn(20),
		Stamina:  baseStats + rand.Intn(20),
		Power:    baseStats + rand.Intn(20),
		Guts:     baseStats + rand.Intn(20),
		Wisdom:   baseStats + rand.Intn(20),
		Mood:     70,
		Fatigue:  0,
		Wins:     0,
		Races:    0,
		IsPlayer: isPlayer,
	}
}

// TotalStats は総合ステータスを返す
func (h *Horse) TotalStats() int {
	return h.Speed + h.Stamina + h.Power + h.Guts + h.Wisdom
}

// RacePower はレース時の実力を計算する
func (h *Horse) RacePower() float64 {
	base := float64(h.Speed*3+h.Stamina*2+h.Power*2+h.Guts*2+h.Wisdom) / 10.0
	moodFactor := float64(h.Mood) / 100.0
	fatigueFactor := 1.0 - float64(h.Fatigue)/200.0
	randomFactor := 0.9 + rand.Float64()*0.2 // 90%-110%のランダム要素

	return base * moodFactor * fatigueFactor * randomFactor
}

// StatusString はステータスの文字列表現を返す
func (h *Horse) StatusString() string {
	return fmt.Sprintf(
		"名前: %s (年齢: %d歳)\n"+
			"スピード: %d  スタミナ: %d\n"+
			"パワー: %d    根性: %d\n"+
			"賢さ: %d      総合: %d\n"+
			"調子: %d%%    疲労: %d%%\n"+
			"戦績: %d戦 %d勝",
		h.Name, h.Age,
		h.Speed, h.Stamina,
		h.Power, h.Guts,
		h.Wisdom, h.TotalStats(),
		h.Mood, h.Fatigue,
		h.Races, h.Wins,
	)
}

// Rest は休養して疲労を回復する
func (h *Horse) Rest() string {
	recovered := 30 + rand.Intn(20)
	h.Fatigue -= recovered
	if h.Fatigue < 0 {
		h.Fatigue = 0
	}
	h.Mood += 10
	if h.Mood > 100 {
		h.Mood = 100
	}
	return fmt.Sprintf("%sは休養した！疲労が%d回復した", h.Name, recovered)
}

// TrainingType はトレーニングの種類
type TrainingType int

const (
	TrainingSpeed TrainingType = iota
	TrainingStamina
	TrainingPower
	TrainingGuts
	TrainingWisdom
)

// Train はトレーニングを実行する
func (h *Horse) Train(trainingType TrainingType) string {
	gain := 3 + rand.Intn(5) // 3-7の成長
	fatigueGain := 10 + rand.Intn(10)

	var statName string
	switch trainingType {
	case TrainingSpeed:
		h.Speed += gain
		if h.Speed > 100 {
			h.Speed = 100
		}
		statName = "スピード"
	case TrainingStamina:
		h.Stamina += gain
		if h.Stamina > 100 {
			h.Stamina = 100
		}
		statName = "スタミナ"
	case TrainingPower:
		h.Power += gain
		if h.Power > 100 {
			h.Power = 100
		}
		statName = "パワー"
	case TrainingGuts:
		h.Guts += gain
		if h.Guts > 100 {
			h.Guts = 100
		}
		statName = "根性"
	case TrainingWisdom:
		h.Wisdom += gain
		if h.Wisdom > 100 {
			h.Wisdom = 100
		}
		statName = "賢さ"
	}

	h.Fatigue += fatigueGain
	if h.Fatigue > 100 {
		h.Fatigue = 100
	}

	// 疲労が高いと調子が下がる
	if h.Fatigue > 70 {
		h.Mood -= 5
		if h.Mood < 0 {
			h.Mood = 0
		}
	}

	return fmt.Sprintf("%sの%sが%d上がった！(疲労+%d)", h.Name, statName, gain, fatigueGain)
}

// 有名な競走馬の名前リスト
var famousHorseNames = []string{
	"サイレンススズカ", "トウカイテイオー", "オグリキャップ",
	"ナリタブライアン", "ディープインパクト", "オルフェーヴル",
	"ゴールドシップ", "キタサンブラック", "アーモンドアイ",
	"ウオッカ", "ブエナビスタ", "ジェンティルドンナ",
	"スペシャルウィーク", "グラスワンダー", "エルコンドルパサー",
	"テイエムオペラオー", "メジロマックイーン", "シンボリルドルフ",
}

// GenerateRivalHorse はライバル馬を生成する
func GenerateRivalHorse() *Horse {
	name := famousHorseNames[rand.Intn(len(famousHorseNames))]
	return NewHorse(name, false)
}
