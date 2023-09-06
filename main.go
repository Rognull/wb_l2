// package main

// import (
//     "bufio"
//     "strings"
//     "strconv"
//     "os"

// )

// type Ghost struct{
//     index int
//     changes int
// }

// type Band struct{
//     ghosts []Ghost
//     index int
// }

// func merge(first Band, second Band)Band{
// result:= Band{}
// result.ghosts = append(result.ghosts,first.ghosts...)
// result.ghosts = append(result.ghosts, second.ghosts...)
// return result
// }

// func main(){

//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Scan()
//     params := strings.Fields(scanner.Text())
//     numOfGhosts, _ := strconv.Atoi(params[0])
//     questions, _ := strconv.Atoi(params[1])
//     bandArr:= make([]Band,0,256)

//     for i:=0 i<numOfGhostsi++{
//         bandArr = append(bandArr,Band{ghosts:make([]Ghost,0,256),index:i+1})
//         bandArr[i].ghosts=append(bandArr[i].ghosts, Ghost{index:i+1,changes:1})
//     }

//     for i:=0i<questionsi++{
//         scanner.Scan()
//         params = strings.Fields(scanner.Text())
//          x,_:= strconv.Atoi(params[0])

//         if x == 1{
//             firstGhost := 0
//             for
//         }
//     }

// }

package main

import "fmt"

func newBand(bandId []int,idX int ,idY int,lastId int ) {
    for  i := 1 ;i < len(bandId); i++ {
        if bandId[i] == idX || bandId[i] == idY {
            bandId[i] = lastId
        }
    }
}

func plusIndex(cbi , bi []int, idx_x,idx_y int) {
    for i := 1 ;i < len(bi);  i++  {
        if bi[i] == idx_x || bi[i] == idx_y {
            cbi[i]++
        }  
    }
}

func unite_spirit( cbi,  bi []int,  x,  y int,  last_id *int) {
    if bi[x] == bi[y] {
        return
    }   
     id_x := bi[x]
     id_y := bi[y]

    plusIndex(cbi, bi, id_x, id_y)
    *last_id += 1

    newBand(bi, id_x, id_y, *last_id)
}

func main() {
    n := 0
     m := 0
    fmt.Scan(&n,&m)

    last_id := n
    count_band_id := make ([]int,n+1)
    for i:=0;i<len(count_band_id);i++{
        count_band_id[i]= 1
    }

   band_id := make([]int,n+1)
  
    for  i := 1 ;i < n+1; i++ {
        band_id[i] = i
    }
    for  i := 0 ;i < m; i++ {
        var nq, x int
        fmt.Scan(&nq,&x)
        if nq == 1  {
            y := 0
            fmt.Scan(&y)
            unite_spirit(count_band_id, band_id, x, y, &last_id)
        } else if nq == 2 {
             y := 0
             fmt.Scan(&y)
            if band_id[x] == band_id[y] {
                fmt.Println("YES")
                
            } else {
                fmt.Println("NO")
            }
        } else {
            fmt.Println(count_band_id[x])
           
        }
    }
    
}