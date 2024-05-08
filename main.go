package main

import "fmt"

func t(n int) int {
    
    if  (n == 1) {
        return 3;
    }
    if (n == 0) {
        return 1;
    }else{
        return (4 * t(n)) - (3 * t(n - 1) ) + 4;;
    }
}

func main() {
    n :=  2;
    var result int;
    result = (4 * t(n)) - (3 * t(n - 1) ) + 4;

    fmt.Println(result);
    
}
