package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/cpu"
	"time"
	"net/http"
	"os"
	"encoding/json"
	//"log"
)

type Memory struct {
	Amount int
	AmountUsed int
	AmountFree int
	Name string
	NameUsed string
	NameFree string
}

type Cpu struct {
	NameCpuuser string
	CPUuser float64
	NameCpusys string
	Cpusys float64

}

func main() {

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	mer := int(memory.Total)
    mer2 := int(memory.Used)
    mer3 := int(memory.Free)
	memories := []Memory{
		Memory {
			Amount: mer,
			Name: "Memory Total Bytes",
			AmountUsed: mer2,
			NameUsed: "Memory Used Bytes",
			AmountFree: mer3,
			NameFree: "Memory Free Bytes",

		},
	}

	bytes, _ := json.Marshal(memories)
	fmt.Println(string(bytes))

	r := mux.NewRouter()
	r.HandleFunc("/memory", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(bytes))
	})

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	total := float64(after.Total - before.Total)




	cpuuser:= float64(after.User-before.User)/total*100
	cpusystem := float64(after.System-before.System)/total*100

	cpus := []Cpu{
		Cpu {
			CPUuser: cpuuser,
			NameCpuuser: "Cpu User",
			Cpusys: cpusystem,
			NameCpusys: "CPU system",


		},
	}

	bytess, _ := json.Marshal(cpus)
	fmt.Println(string(bytess))


	r.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(bytess))
	})

	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	srv.ListenAndServe()

}
