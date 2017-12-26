package main

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Article struct {
	Title  string `redis:"title"`
	Author string `redis:"author"`
	Body   string `redis:"body"`
}

func main() {
	conn, err := conn()
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	defer conn.Close()

	// for _, value := range os.Args {
	// 	if value == "string" {
	// 		testString(conn)
	// 	} else if value == "hmset" {
	// 		testHash(conn)
	// 	} else if value == "lpush" {
	// 		testList(conn)
	// 	} else if value == "sadd" {
	// 		testSet(conn)
	// 	}
	// }
	testString(conn)
	testHash(conn)
	testList(conn)
	testSet(conn)
}

func testString(conn redis.Conn) {
	writeString(conn, "stringkey", "stringvalue1")
	readString(conn, "stringkey")
}

func testHash(conn redis.Conn) {
	var p Article
	p.Author = "cdj"
	p.Body = "test body"
	p.Title = " test title"
	writeHash(conn, "hashkey", &p)
	readHash(conn, "hashkey")
}

func testList(conn redis.Conn) {
	writeList(conn, "keylist", "listvalue1", "listvalue2", "listvalue1", "listvlaue3", "listvlaue4")
	readList(conn, "keylist")
}

func testSet(conn redis.Conn) {
	writeSet(conn, "keyset", "setvalue1", "setvalue2", "setvalue1", "setvalue3")
	readSet(conn, "keyset")
}

func conn() (redis.Conn, error) {
	c, err := redis.Dial("tcp", "10.0.2.206:6379")
	if err != nil {
		fmt.Println("failed to connect redis server,", err.Error())
		return nil, errors.New("failed connect")
	}
	return c, nil
}

func writeString(conn redis.Conn, key string, value string) {
	_, err2 := conn.Do("set", key, value)
	if err2 != nil {
		fmt.Println("failed write do", err2.Error())
		return
	}
	fmt.Println("write string ok")
}

func readString(conn redis.Conn, key string) {
	a1, err2 := redis.String(conn.Do("get", key))
	if err2 != nil {
		fmt.Println("read error", err2.Error())
		return
	}
	fmt.Println("read string is:", a1)
}

func writeHash(conn redis.Conn, key string, article *Article) {
	_, err := conn.Do("hmset", redis.Args{}.Add(key).AddFlat(article)...)
	if err != nil {
		fmt.Println("error write hash", err.Error())
		return
	}
	fmt.Println("write hash ok")
}

func readHash(conn redis.Conn, key string) {
	var p2 Article
	value, _ := redis.Values(conn.Do("hgetall", key))
	redis.ScanStruct(value, &p2)
	fmt.Println("read hash is:", p2.Author, p2.Body, p2.Title)
}

func writeList(conn redis.Conn, key string, values ...interface{}) {
	// conn.Do("lpop", key)
	// v, _ := conn.Do("llen", key)
	// fmt.Println(v)

	_, err := conn.Do("lpush", key, values)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("write list ok")
}

func readList(conn redis.Conn, key string) {
	values, _ := redis.Values(conn.Do("lrange", key, "0", "100"))
	for _, v := range values {
		fmt.Println("read list is:", string(v.([]byte)))
	}
}

func writeSet(conn redis.Conn, key string, values ...interface{}) {
	_, err := conn.Do("sadd", key, values)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("write set ok")
}

func readSet(conn redis.Conn, key string) {
	values, _ := redis.Values(conn.Do("SMEMBERS", key))
	for _, v := range values {
		fmt.Println("read set is:", string(v.([]byte)))
	}
}
