/**
 * User: jackong
 * Date: 11/18/13
 * Time: 3:25 PM
 */
package protocol

import (
	"code.google.com/p/goprotobuf/proto"
	"testing"
	"time"
	"encoding/json"
)

var (
	user *User
)

func init() {
	user = &User{
		Password: proto.String("123456"),
		Time: proto.Int64(time.Now().Unix()),
		IsNice: proto.Bool(true),
		Info: &Info{Name: proto.String("jack"), Age: proto.Int32(22)},
	}
}

func BenchmarkUserEncode2PB(b *testing.B) {
	repeat(func() {
		_, err := proto.Marshal(user)
		if err != nil {
			b.Error(err)
		}
	}, b)
}

func BenchmarkUserEncode2Js(b *testing.B) {
	repeat(func () {
		_, err := json.Marshal(user)
		if err != nil {
			b.Error(err)
		}
	}, b)
}

func BenchmarkUserDecode2PB(b *testing.B) {
    data, err := proto.Marshal(user)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	repeat(func() {
		u := &User{}
		err = proto.Unmarshal(data, u)
		if err != nil {
			b.Error(err)
		}
	}, b)
}

func BenchmarkUserDecode2Js(b *testing.B) {
    data, err := json.Marshal(user)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	repeat(func(){
		u := &User{}
		err = json.Unmarshal(data, u)
		if err != nil {
			b.Error(err)
		}
	}, b)
}

func repeat(f func(), b *testing.B) {
	for	i := 0; i < b.N; i++ {
		f()
	}
}
