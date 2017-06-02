package main

import "./subpkg"

func main() {
    println("hello", "world")
    subpkg.Echo("girl");
}
