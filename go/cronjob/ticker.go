package main

import (
    "fmt"
    "time"
)

type jobFunc2 func(j *job)

type job struct {
    jf     jobFunc2
    params map[string]interface{}
    ch     chan int
}

func NewJob() *job {
    return &job{
        params: make(map[string]interface{}),
        ch:     make(chan int),
    }
}

func (j *job) Run(t time.Duration) {
    ticker := time.NewTicker(time.Second * t)
    go func() {
        for {
            select {
            case <-ticker.C:
                j.jf(j)
            case <-j.ch:
                fmt.Println("收到结束指令")
                ticker.Stop()
                break
            }
        }
    }()

}

func main() {
    j := NewJob()
    j.jf = func(jj *job) {
        fmt.Println("定时任务执行...", time.Now().Format("2006-02-01 15:04:05"), jj.params)
        fmt.Println("-5min", time.Now().Add(-1 * time.Minute * 5).Format("2006-02-01 15:04:05"))
    }
    j.params["p1"] = "第一个参数"
    j.params["p2"] = 100
    j.Run(1)

    // 阻塞一下，等待主进程结束
    tt := time.NewTimer(time.Second * 10)
    <-tt.C
    fmt.Println("over.")

    <-time.After(time.Second * 4)
    fmt.Println("再等待4秒退出。tt 没有终止，打印出 over 后会看见在继续执行...")
    tt.Stop()
    <-time.After(time.Second * 2)
    fmt.Println("tt.Stop()后， tt 仍继续执行，只是关闭了 tt.C 通道。")
}
