package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/text/unicode/norm"
)

// 環境変数
var envs = map[string]string{"SRC_DIR_PATH": "", "DEST_DIR_PATH": "", "ITUNES_MUSIC_PATH": "", "WALKMAN_MUSIC_PATH": ""}

func main() {
	fmt.Println("createing...")

	err := loadEnv()
	if err != nil {
		panic(err)
	}

	files, _ := filepath.Glob(envs["SRC_DIR_PATH"] + "*.m3u")
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

// .envファイルを読み込みグローバル変数にセットしていく
func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	for index := range envs {
		env := os.Getenv(index)
		if env == "" {
			return errors.New(index + " is empty")
		}
		envs[index] = env
	}

	return nil
}

// ウォークマン用のプレイリストを生成する
func createPlayList(srcPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(envs["DEST_DIR_PATH"] + filepath.Base(srcPath))
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
		replaced := strings.Replace(line, envs["ITUNES_MUSIC_PATH"], envs["WALKMAN_MUSIC_PATH"], -1)
		// macはNFD形式なのでNFC形式に変換
		replaced = norm.NFC.String(replaced)

		w.WriteString(replaced)
	}

	w.Flush()

	return nil
}
