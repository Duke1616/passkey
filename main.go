package main

func main() {
	server := InitWebServer()

	server.web.Run(":8100")
}
