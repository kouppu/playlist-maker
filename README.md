# playlist-maker
iTuensで作成されたプレイリストをWALKMAN用に変換して新規作成します。

## 使い方
### インストール
```
$ git clone https://github.com/suhrr/playlist-maker.git
```

### .envファイルを編集
```
// iTunesの曲の絶対パス。環境に合わせて編集
ITUNES_MUSIC_PATH=/iTunes/iTunes Media/music/
// WALKMANの曲の絶対パス。環境に合わせて編集
WALKMAN_MUSIC_PATH=/MUSIC/Music/
```

### iTunesで作成したm3uファイルをiTunesディレクトに移動

### 実行
```
$ ./main.go
```

### WALKMANディレクトに新たに保存される
