package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}

func main() {

	//BaseReflect()
	ReflectGC()
}

type QukuAuditInfo struct {
	Id             string       `bson:"id"`
	AuditInfo      AuditInfo    `bson:"audit_info"`
	SpoolingInfo   SpoolingInfo `bson:"spooling_info"`
	Platform       int          `bson:"platform"`
	TaskId         string       `bson:"task_id"`
	CreateTime     string       `bson:"create_time"`
	UpdateTime     string       `bson:"update_time"`
	Gtid           string       `bson:"gtid"`
	IsLongAudio    int32        `bson:"is_long_audio"`
	IsHistory      int32        `bson:"is_history"`
	Appid          string       `bson:"appid"`
	Category       string       `bson:"category"`
	AllAlgoInfo    string       `bson:"all_algo_info"`
	SubPlatName    string       `bson:"sub_plat_name"`
	SpecialChannel int          `bson:"special_channel"`
	ChannelName    string       `bson:"channel_name"`
	InsertFrom     int          `bson:"insert_from"`
	BlackList      int          `bson:"black_list"`
}

type AuditInfo struct {
	Auditor   string `bson:"auditor"`
	Result    int    `bson:"result"`
	Reason    int    `bson:"reason"`
	Status    int    `bson:"status"`
	Channel   int    `bson:"channel"`
	Label     string `bson:"label"`
	AuditTime string `bson:"audit_time"`
	GotTime   string `bson:"got_time"`
	HitAlgo   string `bson:"hit_algo"`
	Marks     string `bson:"marks"`
}

type SpoolingInfo struct {
	Ids    []string `bson:"ids"`
	Allows []string `bson:"allows"`
	Blocks []string `bson:"blocks"`
}

func ReflectGC() {
	var vecI []interface{}
	for i := 0 ; i < 100000000; i++ {
		//inter := reflect.New(reflect.TypeOf(&QukuAuditInfo{}).Elem()).Interface()
		inter := &QukuAuditInfo{}
		vecI = append(vecI, inter)
		//start := time.Now()
		//runtime.GC()
		//fmt.Printf("GC took %s\n", time.Since(start))
	}
	//runtime.KeepAlive(vecI)
}



func BaseReflect() {

	s := &Student{
		Name: "zhangsan",
		Age:  18,
	}

	sv := reflect.ValueOf(s)
	st := reflect.TypeOf(s)

	svi := sv.Interface()

	// interface{}, reflect.Value, reflect.Type 跟原始类型的关系

	fmt.Printf("svi: %+v\n", svi)

	// 为什么 sv.Elem() 可以调用 CanSet() 方法，而 sv 不能？
	// 因为 sv 是指针类型，而 sv.Elem() 是指针指向的值类型，所以 sv.Elem() 可以调用 CanSet() 方法

	x := Person{
		Name: "zhangsan",
	}
	xv := reflect.ValueOf(x)

	fmt.Printf("x canset: %+v\n", reflect.ValueOf(x).CanSet()) // false
	fmt.Printf("xv canset: %+v\n", xv.CanSet())                // false

	fmt.Printf("canset: %+v\n", sv.CanSet())             // false
	fmt.Printf("elem canset: %+v\n", sv.Elem().CanSet()) // true

	sst := reflect.ValueOf(&s).Elem().Type()

	fmt.Printf("sv: %+v\nst: %+v\nsst: %+v\n", sv, st, sst)

	fmt.Printf("%+v\n", s)


}