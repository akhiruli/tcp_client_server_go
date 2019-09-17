package main

import(
   "fmt"
   "net"
   "bufio"
   "os"
)

func main(){
   conn, err := net.Dial("tcp", "127.0.0.1:34567")
   if err != nil {
      fmt.Printf("Failed to connect to server with error: %s\n", err.Error())
      return
   }

   for{
      fmt.Printf("Enter your message to be sent to client: ")
      stdread := bufio.NewReader(os.Stdin)
      text, _ := stdread.ReadString('\n')
      fmt.Fprintf(conn, text + "\n")
      message, _ := bufio.NewReader(conn).ReadString('\n')
      fmt.Print("Message from server: "+message)
   }
}
