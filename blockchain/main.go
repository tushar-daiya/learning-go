package main

func main() {
	bc := NewBlockChain() //create a new blockchain
	defer bc.db.Close()   //close the database when the program ends

	cli := CLI{bc} //create a new CLI
	cli.Run()      //run the CLI
}
