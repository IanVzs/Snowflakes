package ws

func init() {
	go Manager.Start()
	go HManager.run()
}
