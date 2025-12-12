package main

func main() {
	bridge := InitBridge()
	err := bridge.Start(":9999")
	if err != nil {
		panic(err)
	}
}
