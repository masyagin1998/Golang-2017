/* Визуализация автомата Мили.

   Входные данные: 
   1) Кол-во состояний автомата n.
   2) Размер входного алфавита m.
   3) Номер начального состояния q0.
   4) Матрица переходов n x m.
   5) Матрица  выходов n x m.
   
   Выходные данные:
   1) Описание автомата Мили на языке DOT 
      для последующей визуализации в GraphViz
*/

package main

import "fmt"

func main() {
    var n, m, q0 int
    fmt.Scan(&n, &m, &q0)
    δ := [50][50]int{}
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Scan(&δ[i][j]);
        }
    }
    φ := [50][50]byte{}
    for i := 0; i < n; i++ {
        for j := 0; j < m; {
            var c byte
            fmt.Scanf("%c", &c)
            if (c != ' ' && c != '\n') {
                φ[i][j] = c
                j = j + 1
            }
        }
    }
    fmt.Printf("digraph {\n")
    fmt.Printf("    rankdir = LR\n")
    fmt.Printf("    dummy [label = \"\", shape = none]\n")
    for i := 0; i < n; i++ {
        fmt.Printf("    %d [shape = circle]\n", i)
    }
    fmt.Printf("    dummy -> %d\n", q0)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("%d -> %d", i, δ[i][j])
            fmt.Printf(" [label = \"%c(%c)\"]\n", (97 + j), φ[i][j])
        }
    }
    fmt.Printf("}\n")
}
