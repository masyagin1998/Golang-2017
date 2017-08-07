/* Каноническая нумерация состояний автомата Мили.

   Входные данные: 
   1) Кол-во состояний автомата n.
   2) Размер входного алфавита m.
   3) Номер начального состояния q0.
   4) Матрица переходов и выходов размера n x m.
   
   Выходные данные:
   1) Новое кол-во состояний автомата n1.
   2) Новый размер входного алфавита m1.
   3) Новый номер начального состояния q01.
   4) Новая матрица переходов и выходов размера n1 x m1.
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
                j += 1
            }
        }
    }
    
    var isSimple bool
    isSimple = true
    for i := 0; i < n; i++ {
        var needBreak bool
        for j := 0; j < m; j++ {
            if δ[i][j] != δ[i][0] {
                needBreak = true
                isSimple = false
                break
            }
        }
        if needBreak {
            break;
        }
    }
    
    if isSimple {
        fmt.Printf("1\n%d\n0\n", m);
        for i := 0; i < m; i++ {
            fmt.Printf("0 ");
        }
        fmt.Printf("\n");
        for i := 0; i < n; i++ {
            fmt.Printf("%c ", φ[q0][i]);
        }
        return
    }
    
    visit := [50]int{}
    flag := [50]bool{}
    var start int
    reNumerate(δ, q0, m, &visit, &flag, &start)
    
    δ1 := [50][50]int{}
    φ1 := [50][50]byte{}
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            δ1[visit[i]][j] = visit[δ[i][j]]
            φ1[visit[i]][j] = φ[i][j]
        }
    }
    
    fmt.Printf("%d\n%d\n0\n", start, m);
    for i := 0; i < start; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("%d ", δ1[i][j])
        }
        fmt.Printf("\n")
    }
    for i := 0; i < start; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("%c ", φ1[i][j])
        }
        fmt.Printf("\n")
    }
}

func reNumerate(δ [50][50]int, q0 int, length int, visit *[50]int, flag *[50]bool, start *int) {
    if flag[q0] {
        return
    }
    visit[q0] = *start
    flag[q0] = true
	(*start) += 1
	for i := 0; i < length; i++ {
        reNumerate(δ, δ[q0][i], length, visit, flag, start)
    }
}
