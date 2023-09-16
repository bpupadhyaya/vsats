package main

func main() {
	bc := NewBlockchain("5")
	defer bc.db.Close()
	cli := CLI{}
	cli.Run()
}
