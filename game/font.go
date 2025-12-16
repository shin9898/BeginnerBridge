package game

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// フォント
	GameFont     *text.GoTextFaceSource
	FontFace12   *text.GoTextFace
	FontFace16   *text.GoTextFace
	FontFace20   *text.GoTextFace
	FontFace24   *text.GoTextFace
	FontFace32   *text.GoTextFace
	FontLoaded   bool
)

const (
	// M PLUS 1p フォント（Google Fonts）
	FontURL  = "https://github.com/coz-m/MPLUS_FONTS/raw/master/fonts/ttf/MPLUS1p-Regular.ttf"
	FontPath = "assets/fonts/MPLUS1p-Regular.ttf"
)

// InitFont はフォントを初期化する
func InitFont() error {
	// まずローカルファイルを試す
	fontData, err := loadFontFromFile()
	if err != nil {
		log.Println("ローカルフォントが見つかりません。ダウンロードを試みます...")
		fontData, err = downloadFont()
		if err != nil {
			log.Printf("フォントのダウンロードに失敗: %v", err)
			log.Println("日本語フォントなしで続行します（英語表示のみ）")
			FontLoaded = false
			return nil
		}
		// ダウンロードしたフォントを保存
		if err := saveFontToFile(fontData); err != nil {
			log.Printf("フォントの保存に失敗: %v", err)
		}
	}

	// フォントを解析
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		log.Printf("フォントの解析に失敗: %v", err)
		FontLoaded = false
		return nil
	}

	GameFont = source
	FontLoaded = true

	// 各サイズのフォントフェイスを作成
	FontFace12 = &text.GoTextFace{Source: GameFont, Size: 12}
	FontFace16 = &text.GoTextFace{Source: GameFont, Size: 16}
	FontFace20 = &text.GoTextFace{Source: GameFont, Size: 20}
	FontFace24 = &text.GoTextFace{Source: GameFont, Size: 24}
	FontFace32 = &text.GoTextFace{Source: GameFont, Size: 32}

	log.Println("フォントを正常に読み込みました")
	return nil
}

// loadFontFromFile はローカルファイルからフォントを読み込む
func loadFontFromFile() ([]byte, error) {
	// 実行ファイルのディレクトリを基準にパスを解決
	execPath, err := os.Executable()
	if err == nil {
		execDir := filepath.Dir(execPath)
		fullPath := filepath.Join(execDir, FontPath)
		if data, err := os.ReadFile(fullPath); err == nil {
			return data, nil
		}
	}

	// カレントディレクトリを試す
	if data, err := os.ReadFile(FontPath); err == nil {
		return data, nil
	}

	return nil, os.ErrNotExist
}

// downloadFont はフォントをダウンロードする
func downloadFont() ([]byte, error) {
	resp, err := http.Get(FontURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

// saveFontToFile はフォントをファイルに保存する
func saveFontToFile(data []byte) error {
	// ディレクトリを作成
	dir := filepath.Dir(FontPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(FontPath, data, 0644)
}
