package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
)

func isPortAvailable(port int) bool {
    address := fmt.Sprintf("localhost:%d", port)
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return false
    }
    listener.Close()
    return true
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Invalid number of arguments, port number should be the only arg.");
        return
    }
    arg := os.Args[1];
    port, err := strconv.Atoi(arg);
    if err != nil {
        fmt.Println("Could not parse port number to int!")
        return
    }

    if isPortAvailable(port) {
        fmt.Printf("Port %d is available.\n", port)
    } else {
        fmt.Printf("Port %d is in use or unavailable.\n", port)
    }
}
