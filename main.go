///////////////////////////////////////////////////////////////////////////////
//
// simple "hello world" program
// John Simpson <jms1@jms1.net> 2023-07-02

package main

import (
    "fmt"
    "runtime"
)

///////////////////////////////////////////////////////////////////////////////
//
// Actual values will be filled by -X options in Makefile, these values are
// here in case somebody uses 'go run .'.

var (
    prog_name     string = "hello"
    prog_version  string = "(unset)"
    prog_date     string = "(unset)"
    prog_hash     string = ""
    prog_desc     string = ""
)

///////////////////////////////////////////////////////////////////////////////
//
// Show version info

func do_version( args ...string ) {
    fmt.Printf( "%s\nVersion %s for %s/%s\n" ,
        prog_name , prog_version , runtime.GOOS , runtime.GOARCH )

    if prog_desc != "" {
        fmt.Printf( "Built %s from %s\n" , prog_date , prog_desc )
    } else if prog_hash != "" {
        fmt.Printf( "Built %s from commit %s\n" , prog_date , prog_hash )
    } else {
        fmt.Printf( "Built %s\n" , prog_date )
    }
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func main() {
    do_version() ;
}
