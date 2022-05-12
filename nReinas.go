package main
import (
    "fmt"
    "strings"
)

func nReina(n int, filas int, columnas, tablero []string) {
    if filas == n {
        *respuestas = append(*respuestas, append([]string{}, tablero...))
        return
    }
    for columna := 0; columna < n; columna++ {        
        if columnas[columna] || num1[filas+columna] || num2[filas-columna+(n-1)] { continue }

        columnas[columna] = true

        tablero[filas] = tablero[filas][:columna] + "R" + tablero[filas][columna+1:]
        nReina(n, filas+1, columnas, tablero)
        columnas[columna] = false
        tablero[filas] = tablero[filas][:columna] + tablero[filas][columna + 1:]
    }

}

func resolverNReinas(n int) [][]string {
    columna := make([]bool, n)
    num1 := make([]bool, 2 * n)
    num2 := make([]bool, 2 * n)
    tablero := make([]string, n)

    nReina(n, 0, columna, tablero)
}

func main() {
    resolverNReinas(4)
}