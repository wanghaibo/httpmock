package main

import "github.com/boltdb/bolt"
import "errors"

type KvStore interface {
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	Get(key []byte) (value []byte, err error)
	Close()
}

type kvStore struct {
	db         *bolt.DB
	bucketName []byte
}

func NewKvStore(options map[string]interface{}) (KvStore, error) {
	var path string
	var bucketName []byte
	if val, ok := options["datapath"].(string); ok {
		path = val
	} else {
		return nil, errors.New("require options datapath for KvStore")
	}
	if val, ok := options["bucketName"].([]byte); ok {
		bucketName = val
	} else {
		return nil, errors.New("require options bucketName for KvStore")
	}

	db, err := bolt.Open(path, 0600, nil)
	kv := &kvStore{db: db, bucketName: bucketName}
	kv.createBucket()
	return kv, err
}

func (b *kvStore) Put(key []byte, value []byte) error {
	bucketName := b.bucketName
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return errors.New("bucket does not exists:" + string(bucketName))
		}
		if err := bucket.Put(key, value); err != nil {
			return err
		}
		return nil
	})
}

func (b *kvStore) Get(key []byte) (value []byte, err error) {
	bucketName := b.bucketName
	err = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return errors.New("bucket does not exists:" + string(bucketName))
		}
		value = bucket.Get(key)
		return nil
	})
	return value, err
}

func (b *kvStore) Delete(key []byte) error {
	bucketName := b.bucketName
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		if bucket == nil {
			return errors.New("bucket does not exists:" + string(bucketName))
		}

		err := bucket.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
}

func (b *kvStore) createBucket() (err error) {
	bucketName := b.bucketName
	return b.db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)
		if err != nil && err != bolt.ErrBucketExists {
			return err
		}
		return nil
	})
}

func (b *kvStore) Close() {
	b.db.Close()
}
