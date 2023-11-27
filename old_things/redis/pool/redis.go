package main

import (
	"context"
	"encoding/json"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	DeleteObjectWithKey(key string, ctx context.Context) error
	GetObjectByKey(key string, ctx context.Context) (interface{}, error)
	SaveObjectWithKey(object interface{}, key string, ctx context.Context) error
	GetConnection() (redigo.Conn, error)
}

type ServiceEntity struct {
	redisPool *redigo.Pool
	db        string
}

func NewRedisService(pool *redigo.Pool, db string) Service {
	conn := pool.Get()
	defer conn.Close()

	result, err := conn.Do("PING")
	if err != nil {
		pa := errors.Wrap(err, "[REDIS] Error when connect to redis")
		panic(pa)
	}
	log.Infof("[REDIS] response : %#v", result)
	return &ServiceEntity{
		redisPool: pool,
		db:        db,
	}
}

func (r *ServiceEntity) DeleteObjectWithKey(key string, ctx context.Context) error {
	conn, err := r.GetConnection()
	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when deleting key %v", key)
		return err
	}
	defer conn.Close()

	result, err := conn.Do("del", key)

	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when deleting key %v", key)
		return err
	}
	log.Infof("[REDIS] response : %#v", result)
	return nil
}

func (r *ServiceEntity) SaveObjectWithKey(object interface{}, key string, ctx context.Context) error {
	conn, err := r.GetConnection()
	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when insert object %v key %v", object, key)
		return err
	}
	defer conn.Close()

	jsonObject, err := json.Marshal(object)
	//str := base64.StdEncoding.EncodeToString(jsonObject)
	result, err := conn.Do("set", key, jsonObject)

	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when insert object %v key %v", object, key)
		return err
	}
	log.Infof("[REDIS] response : %#v", result)
	return nil
}

func (r *ServiceEntity) GetObjectByKey(key string, ctx context.Context) (interface{}, error) {
	conn, err := r.GetConnection()
	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when getting key %v", key)
		return nil, err
	}
	defer conn.Close()

	byteArray, err := conn.Do("get", key)

	if err != nil {
		err = errors.Wrap(err, "[REDIS] Error when connect to redis")
		log.WithError(err).Warnf("[REDIS] Error when getting key %v", key)
		return nil, err
	}
	log.Infof("[REDIS] response : %#v", byteArray)
	return byteArray, nil
}


func (r *ServiceEntity) GetConnection() (redigo.Conn, error) {
	conn := r.redisPool.Get()
	_, err := conn.Do("SELECT", r.db)

	return conn, err
}
