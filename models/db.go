package models

import (
	"github.com/15125505/zlog/log"
	"github.com/globalsign/mgo"
	"time"
)

const (
	DbName  = "enterbj"    // 数据库
	BjCard  = "bjcard"     // 表名
	BdToken = "baidutoken" // 表名
)

var globalS *mgo.Session

// 初始化并连接数据库
func LoadMongoDBInfo(host, user, password string) {
	/*
	   载入配置文件中对mongodb的配置信息
	*/
	// Todo 载入配置文件
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Username: user,
		Password: password,
		Timeout:  10 * time.Second, // 10秒超时
	}

	log.Info("[Loading MongoDB]_Waiting mongodb connect to...", host)
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error("[Loading MongoDB]_Mongodb connected Failed: ", err)
		// Todo 中断执行
		panic("[Loading MongoDB]_Please check mongodb connect info.")
	}
	log.Info("[Loading MongoDB]_Mongodb connected Success")
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

// 查询总数
func FindCount(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}

// 排序查询
func FindAllOrderBy(db, collection, orderby string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Sort(orderby).All(result)
}

// 分页排序查询
func FindAllLimitAndOrderBy(db, collection string, limit int, query, selector, result interface{}, orderbys ...string) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	//c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)
	return c.Find(query).Select(selector).Sort(orderbys...).Limit(limit).All(result)
}

func FindOneOrderBy(db, collection, orderby string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Sort(orderby).One(result)
}

// group
func FindAllPipe(db, collection string, pipeline, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.All(result)
}

// 作者：CoderMiner
// 链接：https://www.jianshu.com/p/a2f067cd49b3
// 來源：简书
// 简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
