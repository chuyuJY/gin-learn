package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv := clientv3.NewKV(client)
	// clientv3.WithPrevKV() 是一个可选控制项，用于获取在设置当前键值对之前的该键的键值对
	// 有了该控制项后，putResp 才有 PrevKv 的属性，即获取之前的键值对。

	// context.TODO() 表示当前还不知道用哪个 context 控制该操作，先用该字段占位
	getResp, err := kv.Get(context.TODO(), "/user", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}
	for _, resp := range getResp.Kvs {
		fmt.Printf("key: %s, value:%s\n", string(resp.Key), string(resp.Value))
	}

}
