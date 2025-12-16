package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// DrawText はテキストを描画する
func DrawText(screen *ebiten.Image, str string, x, y float64, clr color.Color) {
	if !FontLoaded || FontFace16 == nil {
		return
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, str, FontFace16, op)
}

// DrawTextSize はサイズ指定でテキストを描画する
func DrawTextSize(screen *ebiten.Image, str string, x, y float64, size int, clr color.Color) {
	if !FontLoaded || GameFont == nil {
		return
	}

	var face *text.GoTextFace
	switch {
	case size <= 12:
		face = FontFace12
	case size <= 16:
		face = FontFace16
	case size <= 20:
		face = FontFace20
	case size <= 24:
		face = FontFace24
	default:
		face = FontFace32
	}

	if face == nil {
		return
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, str, face, op)
}

// DrawTextCentered は中央揃えでテキストを描画する
func DrawTextCentered(screen *ebiten.Image, str string, centerX, y float64, size int, clr color.Color) {
	if !FontLoaded || GameFont == nil {
		return
	}

	var face *text.GoTextFace
	switch {
	case size <= 12:
		face = FontFace12
	case size <= 16:
		face = FontFace16
	case size <= 20:
		face = FontFace20
	case size <= 24:
		face = FontFace24
	default:
		face = FontFace32
	}

	if face == nil {
		return
	}

	// テキストの幅を計算
	w, _ := text.Measure(str, face, 0)
	x := centerX - w/2

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, str, face, op)
}

// DrawPanel は装飾付きパネルを描画
func DrawPanel(screen *ebiten.Image, x, y, w, h float32) {
	// 外枠（ゴールド）
	vector.DrawFilledRect(screen, x, y, w, h, ColorPanelFrame, false)
	// 内側の背景
	vector.DrawFilledRect(screen, x+3, y+3, w-6, h-6, ColorPanelBg, false)
	// 内側の装飾線
	vector.DrawFilledRect(screen, x+6, y+6, w-12, 2, ColorDarkGold, false)
	vector.DrawFilledRect(screen, x+6, y+h-8, w-12, 2, ColorDarkGold, false)
}

// DrawButton は選択可能なボタンを描画
func DrawButton(screen *ebiten.Image, x, y, w, h float32, txt string, selected bool) {
	if selected {
		vector.DrawFilledRect(screen, x, y, w, h, ColorGold, false)
		vector.DrawFilledRect(screen, x+2, y+2, w-4, h-4, ColorDarkBrown, false)
		DrawText(screen, "> "+txt, float64(x)+10, float64(y)+8, ColorWhite)
	} else {
		vector.DrawFilledRect(screen, x, y, w, h, ColorBrown, false)
		vector.DrawFilledRect(screen, x+2, y+2, w-4, h-4, ColorPanelBg, false)
		DrawText(screen, "  "+txt, float64(x)+10, float64(y)+8, ColorLightGray)
	}
}

// DrawStatusBar はステータスバーを描画
func DrawStatusBar(screen *ebiten.Image, x, y, w float32, value int, maxValue int, barColor color.RGBA, label string) {
	// ラベル
	DrawText(screen, label, float64(x), float64(y)-2, ColorWhite)
	// 背景バー
	vector.DrawFilledRect(screen, x+80, y, w, 16, color.RGBA{30, 30, 30, 255}, false)
	// 値バー
	fillWidth := (float32(value) / float32(maxValue)) * w
	vector.DrawFilledRect(screen, x+80, y, fillWidth, 16, barColor, false)
	// 枠線
	vector.StrokeRect(screen, x+80, y, w, 16, 1, color.RGBA{60, 60, 60, 255}, false)
}
