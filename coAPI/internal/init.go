package coAPI

func Run() {

	clients := NewClients()
	go clients.Init()

}
