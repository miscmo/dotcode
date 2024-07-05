package main

import (
    "encoding/json"
    "fmt"
    "time"
)

func test() (i int) {
    defer func() {
        i++
    }()
    return 1
}


func test2() int {
    i := 1
    defer func() {
        i++
    }()

    return i
}

func main() {
    //fmt.Println(test(), test2())
    //cur := time.Now().Unix()

    //bodyAppid := 3000013
    //bodysec := "HpZ~k2zi^EM}3iRP"
    //cur := 1699517236
    //
    //bodysignstr := fmt.Sprintf("%d%s%d", bodyAppid, bodysec, cur)
    //bodysignData := []byte(bodysignstr)
    //bodyMd5 := md5.Sum(bodysignData)
    //bodysignmd5Str := fmt.Sprintf("%x", bodyMd5)
    //
    //fmt.Printf("md5(%s)=%s", bodysignstr, bodysignmd5Str)


    str := `
        {"code":0,"msg":"success","result":"[{\"Fbili_id\":\"0\",\"Fchannels_id\":\"\",\"Fdy_id\":\"6473973488431402254\",\"Fid\":\"123\",\"Finsert_time\":\"2021-06-28T19:57:27+08:00\",\"Fkg_id\":\"0\",\"Fkge_id\":\"0\",\"Fkw_id\":\"0\",\"Fkwbd_id\":\"0\",\"Fmodify_time\":\"2022-07-08T12:14:40+08:00\",\"Fmusician_id\":\"0\",\"Fmv_id\":\"0\",\"Fpcg_id\":\"\",\"Fqeh_id\":\"\",\"Fqq_id\":\"14792439\",\"Fum_id\":\"0\",\"Fvideo_id\":\"123\",\"Fwangyi_id\":\"\",\"Fweibo_id\":\"\",\"Fxhs_id\":\"0\",\"Fyoutube_id\":\"\"}]"}
    `


    type VideoExchageRsp struct {
        Code   int      `json:"code"`
        Msg    string   `json:"msg"`
        Result string `json:"result"`
    }


    var results  []struct {
        FbiliId     string    `json:"Fbili_id"`
        FchannelsId string    `json:"Fchannels_id"`
        FdyId       string    `json:"Fdy_id"`
        Fid         string    `json:"Fid"`
        FinsertTime time.Time `json:"Finsert_time"`
        FkgId       string    `json:"Fkg_id"`
        FkgeId      string    `json:"Fkge_id"`
        FkwId       string    `json:"Fkw_id"`
        FkwbdId     string    `json:"Fkwbd_id"`
        FmodifyTime time.Time `json:"Fmodify_time"`
        FmusicianId string    `json:"Fmusician_id"`
        FmvId       string    `json:"Fmv_id"`
        FpcgId      string    `json:"Fpcg_id"`
        FqehId      string    `json:"Fqeh_id"`
        FqqId       string    `json:"Fqq_id"`
        FumId       string    `json:"Fum_id"`
        FvideoId    string    `json:"Fvideo_id"`
        FwangyiId   string    `json:"Fwangyi_id"`
        FweiboId    string    `json:"Fweibo_id"`
        FxhsId      string    `json:"Fxhs_id"`
        FyoutubeId  string    `json:"Fyoutube_id"`
    }

    type Results struct {
        Items []map[string]string
    }

    rsp := &VideoExchageRsp{}

    if err := json.Unmarshal([]byte(str), rsp); err != nil {
        fmt.Printf("json unmarshal failed, err: %s", err.Error())
        return
    }

    fmt.Printf("json unmarshal succ, rsp: %+v\n", rsp)

    if err := json.Unmarshal([]byte(rsp.Result), &results); err != nil {
        fmt.Printf("json unamrshal failed, err: %s", err.Error())
        return
    }

    fmt.Printf("json unmarshal succ, items: %+v\n", results)
}
