package main

import "github.com/liangyt123/requests"


func main (){

        resp,_ := requests.Get("http://go.xiulian.net.cn")
        println(resp.Text())
}
