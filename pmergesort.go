package main

import (
		"fmt"
		"log"
		"sync"
    "time"
    "math/rand"
		"sort"
		"os"
		"strconv"
)


func merge(left, right []int) (ret []int) {
    ret = []int{}
    for len(left) > 0 && len(right) > 0 {
        var x int
        // ソート済みのふたつのスライスからより小さいものを選んで追加していく (これがソート処理)
        if right[0] > left[0] {
            x, left = left[0], left[1:]
        } else {
            x, right = right[0], right[1:]
        }
        ret = append(ret, x)
    }
    // 片方のスライスから追加する要素がなくなったら残りは単純に連結できる (各スライスは既にソートされているため)
    ret = append(ret, left...)
    ret = append(ret, right...)
    return
}

func merge_sort(left, right []int, id int) (ret []int) {
    // ふたつのスライスをそれぞれ再帰的にソートする
		if id < 4 {
		var wg sync.WaitGroup
		wg.Add(2)
		go func(){
    if len(left) > 1 {
        l, r := split(left)
        left = merge_sort(l, r,id+1)
    }
		wg.Done()
	}()
	go func(){
    if len(right) > 1 {
        l, r := split(right)
        right = merge_sort(l, r,id+1)
    }
		wg.Done()
	}()
	wg.Wait()

    // ソート済みのふたつのスライスをひとつにマージする
    ret = merge(left, right)
    return
	} else {
		v := append(left,right...)
	 	sort.Ints(v)
		ret = v
		return
	}
}

func split(values []int) (left, right []int) {
    // スライスを真ん中でふたつに分割する
    left = values[:len(values) / 2]
    right = values[len(values) / 2:]
    return
}

func merge_Sort(values []int) (ret []int) {
    left, right := split(values)
    ret = merge_sort(left, right,1)
    return
}

func main() {
    // UNIX 時間をシードにして乱数生成器を用意する
    t := time.Now().Unix()
    s := rand.NewSource(t)
    r := rand.New(s)

    // ランダムな値の入った配列を作る
    N , _ := strconv.Atoi(os.Args[1])
		fmt.Println("N=",N)
    values := make([]int, N)
    for i := 0; i < N; i++ {
        values[i] = r.Intn(N)
    }

		log.Print("Start")

    // ソートして結果を出力する
		sortedValues := merge_Sort(values)
		//merge_Sort(values)
		log.Print("end")
		fmt.Println(sortedValues)
	}

func bubble_sort(values []int, start int, stop int) (ret []int){
	//fmt.Println("bubble")
	for i:=start; i<stop; i++ {
		for j:=i+1; j<stop; j++{
			if values[i] > values[j] {
				t := values[i]
				values[i] = values[j]
				values[j] = t
			}
		}
	}
	ret = values
	return
}
