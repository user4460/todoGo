//todoリストを表示する
func List() {
	// ファイルを開く
	file, err := os.OpenFile("todos.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ファイルの内容を読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
