package serbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.Id)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "item"
	o = append(o, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Item)))
	for za0001 := range z.Item {
		o = msgp.AppendString(o, z.Item[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "item":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Item) >= int(zb0002) {
				z.Item = (z.Item)[:zb0002]
			} else {
				z.Item = make([]string, zb0002)
			}
			for za0001 := range z.Item {
				z.Item[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *User) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Item {
		s += msgp.StringPrefixSize + len(z.Item[za0001])
	}
	return
}
