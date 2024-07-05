package main

import (
	"os"
	"reflect"
	"runtime"
	"runtime/pprof"
	"testing"
 _ "net/http/pprof"
)

type AuditInfo struct {

}

// Your data type
type DataItem struct {
	// add your fields here
	Id string
	//AuditInfo AuditInfo
	CreateTime string  `tag:"3"  required:"false"  json:"CreateTime"`
	UpdateTime string  `tag:"4"  required:"false"  json:"UpdateTime"`
	Platform int32  `tag:"5"  required:"false"  json:"Platform"`
	TaskId string  `tag:"6"  required:"false"  json:"TaskId"`
	Gtid string  `tag:"7"  required:"false"  json:"Gtid"`
	IsLongAudio int32  `tag:"8"  required:"false"  json:"IsLongAudio"`
	IsHistory int32  `tag:"9"  required:"false"  json:"IsHistory"`
	Appid string  `tag:"10"  required:"false"  json:"Appid"`
	Category string  `tag:"11"  required:"false"  json:"Category"`
	AllAlgoInfo string  `tag:"12"  required:"false"  json:"AllAlgoInfo"`
	SubPlatName string  `tag:"13"  required:"false"  json:"SubPlatName"`
	SpecialChannel int32  `tag:"14"  required:"false"  json:"SpecialChannel"`
	ChannelName string  `tag:"15"  required:"false"  json:"ChannelName"`
	InsertFrom int32  `tag:"16"  required:"false"  json:"InsertFrom"`
}



// mockIter simulates an iterator over a certain number of items.
type mockIter struct {
	count int
	index int
}

// Next simulates the iteration over items.
func (it *mockIter) Next(item interface{}) bool {
	if it.index >= it.count {
		return false
	}

	// Here you might want to fill in the item with actual data.
	it.index++

	return true
}

// Done checks whether the iteration is complete.
func (it *mockIter) Done() bool {
	return it.index >= it.count
}

func BenchmarkReflectAndIterate(b *testing.B) {
	dummyData := DataItem{}

		// ...

		cpuFile, err := os.Create("cpu.pprof")
		if err != nil {
			b.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(cpuFile); err != nil {
			b.Fatal("could not start CPU profile: ", err)
		}
		defer cpuFile.Close()
		defer pprof.StopCPUProfile()

	memFile, err := os.Create("mem.pprof")
	if err != nil {
		b.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // GC，以获得更为精确的内存分析数据
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		b.Fatal("could not write memory profile: ", err)
	}
	memFile.Close()

		// 执行基准测试代码

		// ...

	for n := 0; n < b.N; n++ {
		var list []interface{}
		it := &mockIter{count: 100000} // assuming we're iterating over 1000 items

		for !it.Done() {
			pItemInterface := reflect.New(reflect.TypeOf(dummyData)).Interface()
			if it.Next(pItemInterface) {
				list = append(list, pItemInterface)
			}
		}
	}
}

func BenchmarkDirectAllocationAndIterate(b *testing.B) {
	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		b.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		b.Fatal("could not start CPU profile: ", err)
	}
	defer cpuFile.Close()
	defer pprof.StopCPUProfile()

	memFile, err := os.Create("mem.pprof")
	if err != nil {
		b.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // GC，以获得更为精确的内存分析数据
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		b.Fatal("could not write memory profile: ", err)
	}
	memFile.Close()
	// 执行基准测试代码

	// ...
	for n := 0; n < b.N; n++ {
		var list []*DataItem
		it := &mockIter{count: 100000} // assuming we're iterating over 1000 items

		for !it.Done() {
			pItem := new(DataItem)
			if it.Next(pItem) {
				list = append(list, pItem)
			}
		}
	}
}