package main

import (
	"net"
	"net/rpc"
	"log"
	"fmt"
	"io"
	"os"
	"strings"
)
const BUFFER_SIZE = 1024

var conn net.Conn

var maxClients int


type Reply struct {
	Data []byte
	N ,EOF int
}
type Args struct {
	BufferSize int
	FileName string
	FilePath string
	CurrentByte int64
}
type FileTransfer struct{
	
}
func (t *FileTransfer) GetFile(args *Args, reply *Reply) error {
	file, err := os.Open(strings.TrimSpace("src/"+args.FilePath+args.FileName+".txt"))
	if err != nil && err != io.EOF {
		log.Fatal(err)
		defer file.Close()
		return err
	}
	reply.Data = make([]byte, args.BufferSize)
	reply.N, err = file.ReadAt(reply.Data, args.CurrentByte)
	if err == io.EOF {
		reply.EOF = 1
		file.Close()
		return nil
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	reply.EOF = 0
	file.Close()
	return err
	}

func main(){
	fmt.Println("start listening")
	fileTransfer := new (FileTransfer)
	rpc.Register(fileTransfer)
	listener, e := net.Listen("tcp", ":0666")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	maxClients = 4
	
	sem := make(chan bool, maxClients)
	for {
		var err error
			conn, err = listener.Accept()
			if err != nil {
				log.Fatal("accept error: " + err.Error())
			} else {
				log.Printf("new connection established.")
				sem <- true
				go func (net.Conn) {
				defer func() { <-sem } ()
				rpc.ServeConn(conn)
				}(conn)
			}
	}
}
