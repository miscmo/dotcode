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
	//ReflectGC()
	//Base2()
	//Base3()
	//BaseStruct()
	//ForEachAudit(AuditInfo{})
	//
	//CreatAudit(AuditInfo{})

	SetVal()
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
	SpoolingInfo   SpoolingInfo `bson:"spooling_info"`
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

func Base2() {
	var s string = "AB"

	fmt.Println(reflect.TypeOf(s[0]))
	fmt.Println(reflect.TypeOf(s[0]).Kind())

	for _, v := range s {
		fmt.Println(reflect.TypeOf(v), reflect.TypeOf(v).Kind())
	}

	var v interface{} = 1
	var s2 uint8 = 1

	temp1 := int(s2)
	temp2 := v.(int)

	fmt.Println(temp1, temp2)
	fmt.Println(reflect.TypeOf(temp1).Kind(), reflect.TypeOf(temp2).Kind())
}

type MyM struct {
	i int64
}

type MyN struct {
	i int64
}
func Base3() {

	m := MyM{i: 12}

	var n MyN

	n = MyN(m)

	fmt.Println(n, m)
	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(m))
	fmt.Println(reflect.TypeOf(n).Kind(), reflect.TypeOf(m).Kind())

	var k interface{} = "s"
	l, ok := k.(int)


	fmt.Println(ok)
	fmt.Println(reflect.TypeOf(k).Kind(), reflect.TypeOf(l).Kind())
}

func BaseStruct() {
	kelu := Person{"kelu", 25}

	t := reflect.TypeOf(kelu)

	n := t.NumField()
	for i := 0; i < n; i++ {
		fmt.Println(t.Field(i).Name)
		fmt.Println(t.Field(i).Type)
	}
}

func ForEachAudit(info AuditInfo) {

	typ := reflect.TypeOf(info)
	val := reflect.ValueOf(info)

	for i := 0; i < typ.NumField(); i++ {
		fmt.Printf("%v: %v\n", typ.Field(i).Name, val.Field(i).Interface())
	}

}

func CreatAudit(info AuditInfo) AuditInfo {
	st := reflect.TypeOf(info)
	//st = st.Elem() 	// panic

	vl := reflect.New(st)

	fmt.Printf("%v", vl)

	return AuditInfo{}
}

func SetVal() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// panic: reflect: reflect.flag.mustBeAssignable using addressable value
	v.SetFloat(7.1)

	p := reflect.ValueOf(&x)
	// panic: reflect: reflect.flag.mustBeAssignable using addressable value
	p.SetFloat(7.1)

	e := reflect.ValueOf(&x).Elem()
	// OK
	e.SetFloat(7.1)
}