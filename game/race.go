package game

import (
	"sort"
)

// RaceGrade はレースのグレード
type RaceGrade int

const (
	GradeG3 RaceGrade = iota
	GradeG2
	GradeG1
)

func (g RaceGrade) String() string {
	switch g {
	case GradeG3:
		return "G3"
	case GradeG2:
		return "G2"
	case GradeG1:
		return "G1"
	}
	return ""
}

// Prize はグレードに応じた賞金を返す
func (g RaceGrade) Prize() int {
	switch g {
	case GradeG3:
		return 3000
	case GradeG2:
		return 8000
	case GradeG1:
		return 20000
	}
	return 0
}

// Race はレースを表す構造体
type Race struct {
	Name     string
	Grade    RaceGrade
	Distance int // メートル
	Horses   []*Horse
}

// RaceEntry はレース中の馬の状態
type RaceEntry struct {
	Horse    *Horse
	Position float64 // 現在位置（0-100%）
	Speed    float64 // 現在の速度
	Finished bool
	Rank     int
}

// RaceState はレースの進行状態
type RaceState struct {
	Race       *Race
	Entries    []*RaceEntry
	IsRunning  bool
	IsFinished bool
	Tick       int
	Results    []*RaceEntry
}

// AvailableRaces は出走可能なレース一覧
var AvailableRaces = []Race{
	{Name: "新馬戦", Grade: GradeG3, Distance: 1600},
	{Name: "若駒ステークス", Grade: GradeG3, Distance: 1800},
	{Name: "弥生賞", Grade: GradeG2, Distance: 2000},
	{Name: "皐月賞", Grade: GradeG1, Distance: 2000},
	{Name: "日本ダービー", Grade: GradeG1, Distance: 2400},
	{Name: "菊花賞", Grade: GradeG1, Distance: 3000},
	{Name: "天皇賞・春", Grade: GradeG1, Distance: 3200},
	{Name: "天皇賞・秋", Grade: GradeG1, Distance: 2000},
	{Name: "有馬記念", Grade: GradeG1, Distance: 2500},
	{Name: "ジャパンカップ", Grade: GradeG1, Distance: 2400},
}

// NewRaceState は新しいレース状態を作成する
func NewRaceState(race *Race, playerHorse *Horse, numRivals int) *RaceState {
	entries := make([]*RaceEntry, 0, numRivals+1)

	// プレイヤーの馬を追加
	entries = append(entries, &RaceEntry{
		Horse:    playerHorse,
		Position: 0,
		Speed:    0,
		Finished: false,
	})

	// ライバル馬を追加
	for i := 0; i < numRivals; i++ {
		rival := GenerateRivalHorse()
		// グレードに応じてライバルの強さを調整
		switch race.Grade {
		case GradeG2:
			rival.Speed += 10
			rival.Stamina += 10
		case GradeG1:
			rival.Speed += 20
			rival.Stamina += 20
			rival.Power += 15
		}
		entries = append(entries, &RaceEntry{
			Horse:    rival,
			Position: 0,
			Speed:    0,
			Finished: false,
		})
	}

	return &RaceState{
		Race:       race,
		Entries:    entries,
		IsRunning:  false,
		IsFinished: false,
		Tick:       0,
		Results:    make([]*RaceEntry, 0),
	}
}

// Start はレースを開始する
func (rs *RaceState) Start() {
	rs.IsRunning = true
	rs.IsFinished = false
	rs.Tick = 0
	rs.Results = make([]*RaceEntry, 0)

	// 初期速度を設定
	for _, entry := range rs.Entries {
		entry.Position = 0
		entry.Speed = entry.Horse.RacePower() * 0.8
		entry.Finished = false
	}
}

// Update はレースを1フレーム進める
func (rs *RaceState) Update() {
	if !rs.IsRunning || rs.IsFinished {
		return
	}

	rs.Tick++
	finishCount := 0

	for _, entry := range rs.Entries {
		if entry.Finished {
			finishCount++
			continue
		}

		// 速度の変動（スタミナと根性の影響）
		speedVariation := 0.98 + (float64(entry.Horse.Stamina)/1000.0)
		gutsBoost := 1.0
		if entry.Position > 80 {
			// 終盤の根性発動
			gutsBoost = 1.0 + (float64(entry.Horse.Guts) / 500.0)
		}

		entry.Speed = entry.Horse.RacePower() * speedVariation * gutsBoost * 0.15

		// 位置を更新
		entry.Position += entry.Speed
		if entry.Position >= 100 {
			entry.Position = 100
			entry.Finished = true
			entry.Rank = len(rs.Results) + 1
			rs.Results = append(rs.Results, entry)
			finishCount++
		}
	}

	// 全馬がゴールしたらレース終了
	if finishCount == len(rs.Entries) {
		rs.IsRunning = false
		rs.IsFinished = true

		// 結果を順位でソート
		sort.Slice(rs.Results, func(i, j int) bool {
			return rs.Results[i].Rank < rs.Results[j].Rank
		})

		// プレイヤーの馬の戦績を更新
		for _, entry := range rs.Entries {
			if entry.Horse.IsPlayer {
				entry.Horse.Races++
				if entry.Rank == 1 {
					entry.Horse.Wins++
				}
				// レース後の疲労
				entry.Horse.Fatigue += 30
				if entry.Horse.Fatigue > 100 {
					entry.Horse.Fatigue = 100
				}
				break
			}
		}
	}
}

// GetPlayerRank はプレイヤーの馬の現在順位を返す
func (rs *RaceState) GetPlayerRank() int {
	// 現在位置でソート
	positions := make([]struct {
		Entry    *RaceEntry
		Position float64
	}, len(rs.Entries))

	for i, entry := range rs.Entries {
		positions[i].Entry = entry
		positions[i].Position = entry.Position
	}

	sort.Slice(positions, func(i, j int) bool {
		return positions[i].Position > positions[j].Position
	})

	for i, p := range positions {
		if p.Entry.Horse.IsPlayer {
			return i + 1
		}
	}
	return 0
}

// GetPlayerEntry はプレイヤーの馬のエントリを返す
func (rs *RaceState) GetPlayerEntry() *RaceEntry {
	for _, entry := range rs.Entries {
		if entry.Horse.IsPlayer {
			return entry
		}
	}
	return nil
}
