package cliente
 
import (
	"fmt"
	"net"
	"net/rpc"
	"log"
	"os"
	"time"
	"strings"
)
const BUFFER_SIZE = 1024
type Reply struct {
Data []byte
N, EOF int
}
type Args struct {
BufferSize int
FileName string
FilePath string
CurrentByte int64
}
func Run(id string, filename string, filepath string){
	ti := time.Now()
conn, err := net.Dial("tcp", "127.0.0.1:0666")
if err != nil {
log.Fatal("Errou ao conectar:", err)
}
client := rpc.NewClient(conn)
//file to write to
file, err := os.Create(strings.TrimSpace("src/cliente/"+filename+id+".txt"))
if err != nil {
log.Fatal(err)
}
var reply Reply
args := &Args{BUFFER_SIZE,filename,filepath,0}
for {
err = client.Call("FileTransfer.GetFile", args, &reply)
if err != nil {
log.Println("Erro na transferência: ", err)
break
}
//fmt.Println(reply.Data)
_,err = file.WriteAt(reply.Data[:reply.N], args.CurrentByte)
if err != nil {
log.Println("Erro na cópia:", err)
break
}
args.CurrentByte+=BUFFER_SIZE
if reply.EOF == 1 {
break
}
}
file.Close()
conn.Close()
tempofinal := time.Now().Sub(ti)
fmt.Printf("%v\n",tempofinal)

}

