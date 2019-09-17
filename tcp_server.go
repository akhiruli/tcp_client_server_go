package main

import(
   "net"
   "fmt"
)

const(
   HOST="127.0.0.1"
   PORT= "34567"
   TYPE="tcp"
)

func HandleRequest(conn net.Conn){
   buffer := make([]byte, 1024)

   for{
      _, err := conn.Read(buffer)
      if err != nil{
         fmt.Printf("Error reading: %s\n", err.Error())
         conn.Close()
         break
      }

      conn.Write(buffer)
   }
}

func main(){
   listener, err := net.Listen(TYPE, ":"+PORT)
   if err != nil{
      fmt.Printf("Failed to listen on port %d with error: %s\n", PORT, err.Error())
      return
   }

   defer listener.Close()

   for{
      con, err := listener.Accept()
      if err != nil {
         fmt.Printf("Error in accepting client connection: %s\n", err.Error())
         break
      }

      go HandleRequest(con)
   }
}
