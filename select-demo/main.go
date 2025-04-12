package main

import "fmt"

func main() {

}

func workByPriority(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority // 优先执行ch1任务
				}
			}
			fmt.Println(job2) // 直到ch1任务完成间隙才有ch2任务执行的机会
		}
	}
}
