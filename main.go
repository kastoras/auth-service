package main

func main() {

	address := ":3030"
	server := NewAPIServer(address)
	server.Run()
}
