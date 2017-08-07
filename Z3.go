/* Визуализация минимизированного автомата Мили.

   Входные данные: 
   1) Кол-во состояний автомата n.
   2) Размер входного алфавита m.
   3) Номер начального состояния q0.
   4) Матрица переходов n x m.
   5) Матрица  выходов n x m.
   
   Выходные данные:
   1) Описание минимизированного автомата Мили на языке DOT 
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
                j += 1
            }
        }
    }

    AufenkampHohn(&δ, &φ, &n, &m, &q0);

    visit := [50]int{}
    flag := [50]bool{}
    var start int
    reNumerate(δ, q0, m, &visit, &flag, &start)
    
    δ1 := [50][50]int{}
    φ1 := [50][50]byte{}
    for i := 0; i < n; i++ {
        if ! flag[i] {
            continue
        }
        for j := 0; j < m; j++ {
            δ1[visit[i]][j] = visit[δ[i][j]]
            φ1[visit[i]][j] = φ[i][j]
        }
    }

    fmt.Printf("digraph {\n")
    fmt.Printf("    rankdir = LR\n")
    fmt.Printf("    dummy [label = \"\", shape = none]\n")
    for i := 0; i < start; i++ {
        fmt.Printf("    %d [shape = circle]\n", i)
    }
    fmt.Printf("    dummy -> 0\n")
    for i := 0; i < start; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("    %d -> %d", i, δ1[i][j]);
            fmt.Printf(" [label = \"%c(%c)\"]\n", 97 + j, φ1[i][j]);
        }
    }
    fmt.Printf("}");
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

func AufenkampHohn(δ *[50][50]int, φ *[50][50]byte, n *int, m *int, q0 *int) {
    π:= [50]int{}
    var m1, m2 int
    m1 = -1
    m2 = -1
    Split1(&m1, &π, n, m, φ);
    for ; m1 != m2; {
        m2 = m1
        Split(&m1, &π, n, m, δ);
    }
    π1 := [50]int{}
    π2 := [50]int{}
    var a int
    for i := 0; i < *n; i++ {
        if π[i] == i {
            π2[a] = i
            π1[i] = a
            a++
        }
    }
    *n = m1
    *q0 = π1[π[*q0]]
    p := [50][50]byte{}
    for i := 0; i < *n; i++ {
        for j := 0; j < *m; j++ {
            δ[i][j] = π1[π[δ[π2[i]][j]]]
            p[i][j] = φ[π2[i]][j]
        }
    }
    for i := 0; i < *n; i++ {
        for j := 0; j < *m; j++ {
            φ[i][j] = p[i][j]
        }
    }
}

func Split1(q *int, π *[50]int, n *int, m *int, φ *[50][50]byte) {
    *q = *n
    q1q2 := [50]int{}
    for i := 0; i < *q; i++ {
        q1q2[i] = i
    }
    for i := 0; i < *n; i++ {
        for j := i + 1; j < *n; j++ {
            if Find(&q1q2, i) != Find(&q1q2, j) {
                var eq bool
                eq = true
                for k := 0; k < *m; k++ {
                    if φ[i][k] != φ[j][k] {
                        eq = false
                        break
                    }
                }
                if eq {
                    Union(&q1q2, i, j);
                    *q--
                }
            }
        }
    }
    for i := 0; i < *n; i++ {
        π[i] = Find(&q1q2, i);
    }
}

func Split(m *int, π *[50]int, n *int, m_ptr *int, δ *[50][50]int) {
    *m = *n
    q1q2 := [50]int{}
    for i := 0; i < *m; i++ {
        q1q2[i] = i
    }
    for i := 0; i < *n; i++ {
        for j := i + 1; j < *n; j++ {
            if (π[i] == π[j]) && (Find(&q1q2, i) != Find(&q1q2, j)) {
                var eq bool
                eq = true
                for k := 0; k < *m_ptr; k++ {
                    if π[δ[i][k]] != π[δ[j][k]] {
                        eq = false
                        break
                    }
                }
                if eq {
                    Union(&q1q2, i, j);
                    *m--
                }
            }
        }
    }
    for i := 0; i < *n; i++ {
        π[i] = Find(&q1q2, i);
    }
}

func Union(q1q2 *[50]int, x int, y int) {
    var root_x int
    root_x = Find(q1q2, x)
    var root_y int
    root_y = Find(q1q2, y)
    if root_x < root_y {
        q1q2[root_x] = root_y
    } else {
        q1q2[root_y] = root_x
        if (q1q2[root_x] == q1q2[root_y]) && (root_x != root_y) {
            root_x++
        }
    }
}

func Find(q1q2 *[50]int, x int) int {
    if q1q2[x] == x {
        return x
    } else {
        q1q2[x] = Find(q1q2, q1q2[x]);
        return q1q2[x]
    }
}
