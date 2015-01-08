package main
import (
	"cliente"
	"fmt"
	"os"
)
func test (name string) {
	fmt.Println(name)
}
func main() {
	FilePath := "arquivos/10kb/"
	Filename := "Arquivo"
	go cliente.Run("1",Filename,FilePath)
	go cliente.Run("2",Filename,FilePath)
	go cliente.Run("3",Filename,FilePath)
	go cliente.Run("4",Filename,FilePath)
	go cliente.Run("5",Filename,FilePath)
	var input string 
	fmt.Scanln(&input)
	os.Exit(0);
}
