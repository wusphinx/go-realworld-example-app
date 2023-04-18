// generate by openai3.5
package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Record struct {
	Op    string
	Key   string
	Value string
}

const (
	RecordHeaderSize = 8 // key size (4 bytes) + value size (4 bytes)
)

func main() {
	walFile, err := os.OpenFile("wal.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open WAL file failed: %v", err))
	}
	defer walFile.Close()

	data := make([]*Record, 0)

	// 插入数据记录
	data = append(data, &Record{
		Op:    "set",
		Key:   "k1",
		Value: "v1",
	})

	// 追加写入 WAL 日志
	for _, record := range data {
		// 计算 RecordHeaderSize 的值
		keySize := uint32(len(record.Key))
		valueSize := uint32(len(record.Value))
		recordHeaderSize := keySize + valueSize

		// 写入 RecordHeaderSize
		if err := binary.Write(walFile, binary.LittleEndian, recordHeaderSize); err != nil {
		}

		// 写入 key 和 value
		if _, err := walFile.WriteString(record.Key); err != nil {
			panic(fmt.Errorf("write key failed: %v", err))
		}
		if _, err := walFile.WriteString(record.Value); err != nil {
			panic(fmt.Errorf("write value failed: %v", err))
		}
	}

	// 将 WAL 日志数据 flush 到磁盘
	if err := walFile.Sync(); err != nil {
		panic(fmt.Errorf("flush WAL file to disk failed: %v", err))
	}

	fmt.Println("Write-ahead Log has been written successfully.")
}
