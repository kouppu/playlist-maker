package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/unicode/norm"
)

const (
	SRC_DIR_PATH       = "./iTunes/"
	DEST_DIR_PATH      = "./walkman/"
	ITUNES_MUSIC_PATH  = "/Volumes/HDD/iTunes/iTunes Media/music/"
	WALKMAN_MUSIC_PATH = "/MUSIC/Music/"
)

func main() {
	fmt.Println("createing...")

	files, _ := filepath.Glob(SRC_DIR_PATH + "*.m3u")
	if len(files) == 0 {
		fmt.Println("not found")
		os.Exit(0)
	}

	for _, f := range files {
		fmt.Println(filepath.Base(f))
		err := createPlayList(f)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("finish!")
}

// ウォークマン用のプレイリストを生成する
func createPlayList(srcPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(DEST_DIR_PATH + filepath.Base(srcPath))
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer destFile.Close()

	r := bufio.NewReader(srcFile)
	w := bufio.NewWriter(destFile)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}

		// パスを置換
		replaced := strings.Replace(line, ITUNES_MUSIC_PATH, WALKMAN_MUSIC_PATH, -1)
		// macはNFD形式なのでNFC形式に変換
		replaced = norm.NFC.String(replaced)

		w.WriteString(replaced)
	}

	w.Flush()

	return nil
}
