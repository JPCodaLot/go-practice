package main

import (
  "io"
  "net"
  "os"
  "flag"
  "log"
)

var hostFlag = flag.String("h", "localhost:8000", "host and port number")

func main() {
  flag.Parse()
  conn, err := net.Dial("tcp", *hostFlag)
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()
  go mustCopy(os.Stdout, conn)
  mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
  if _, err := io.Copy(dst, src); err != nil {
    log.Fatal(err)
  }
}
