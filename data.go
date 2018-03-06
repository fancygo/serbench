package serbench

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//go:generate protoc --go_out=. proto.proto
//go:generate protoc --gogofaster_out=. -I. -I. mygogo.proto

type User struct {
	Id   int      `json:"id" xml:"id,attr" msg:"id"`
	Name string   `json:"name" xml:"name" msg:"name"`
	Item []string `json:"item" xml:"item" msg:"item"`
}
