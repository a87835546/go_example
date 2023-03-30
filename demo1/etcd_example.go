package demo1

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var etcdClient = &clientv3.Client{}

func init() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"13.215.251.145:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal("err --->>> ", err.Error())
	}
	etcdClient = client
	log.Printf("init etcd .....")

	log.Printf(" etcd reading  .....")
	resp, err := etcdClient.Get(ctx, "test")
	if err != nil {
		log.Fatal("etcd get func err --->>> ", err.Error())
	}
	log.Printf("read resp:%v ,err:%v", resp.Kvs, err)
	log.Printf(" etcd end  .....")
}
func SetKey() {
	resp, err := etcdClient.Put(ctx, "test", "第一次测试etcd key")
	log.Printf("put resp:%v ,err:%v", resp, err)
}
func ReadKey() {
	log.Printf(" etcd reading  .....")
	resp, err := etcdClient.Get(ctx, "test")
	log.Printf("read resp:%v ,err:%v", resp.Kvs, err)
	log.Printf(" etcd end  .....")
}
