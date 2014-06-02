// Code generated by protoc-gen-gogo.
// source: example.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	A
	B
	C
	U
*/
package test

import testing "testing"
import math_rand "math/rand"
import time "time"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import encoding_json "encoding/json"
import testing2 "testing"
import math_rand2 "math/rand"
import time2 "time"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import fmt "fmt"
import math_rand5 "math/rand"
import time5 "time"
import testing5 "testing"
import code_google_com_p_gogoprotobuf_proto2 "code.google.com/p/gogoprotobuf/proto"
import math_rand6 "math/rand"
import time6 "time"
import testing6 "testing"
import fmt1 "fmt"
import go_parser "go/parser"
import math_rand7 "math/rand"
import time7 "time"
import testing7 "testing"
import math_rand8 "math/rand"
import time8 "time"
import testing8 "testing"
import code_google_com_p_gogoprotobuf_proto3 "code.google.com/p/gogoprotobuf/proto"
import testing9 "testing"

func TestAProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkAProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*A, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedA(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(NewPopulatedA(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &A{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestBProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestBMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkBProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*B, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedB(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkBProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(NewPopulatedB(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &B{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestCProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestCMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkCProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*C, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedC(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkCProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(NewPopulatedC(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &C{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestUProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkUProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*U, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedU(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto.Marshal(NewPopulatedU(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &U{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestAJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestBJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedB(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestCJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedC(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestUJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestAProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestBProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestBProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestCProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedC(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &C{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestCProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedC(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &C{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUOnlyOne(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	v := p.GetValue()
	msg := &U{}
	if !msg.SetValue(v) {
		t.Fatalf("OnlyOne: Could not set Value")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !OnlyOne Equal %#v", msg, p)
	}
}
func TestAStringer(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBStringer(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestCStringer(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestUStringer(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestASize(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkASize(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(616))
	total := 0
	pops := make([]*A, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedA(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestBSize(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedB(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkBSize(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(616))
	total := 0
	pops := make([]*B, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedB(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestCSize(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedC(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkCSize(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(616))
	total := 0
	pops := make([]*C, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedC(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestUSize(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkUSize(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(616))
	total := 0
	pops := make([]*U, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedU(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestAGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestBGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestCGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestUGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestAFace(t *testing7.T) {
	popr := math_rand7.New(math_rand7.NewSource(time7.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func TestAVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestBVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestCVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestUVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestExampleDescription(t *testing9.T) {
	ExampleDescription()
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
