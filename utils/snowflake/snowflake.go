package snowflake

import (
	"errors"
	"sync"
	"time"
)

/*
	相关文档参看：https://segmentfault.com/a/1190000024499175
	twitter 的雪花算法，用来短时间内生成大量的唯一id，特点：
	1. 能满足高并发分布式系统环境下ID不重复
	2. 生成效率高
	3. 基于时间戳，可以保证基本有序递增
	4. 不依赖于第三方的库或者中间件
	5. 生成的id具有时序性和唯一性

	最大可容纳1024个分布式节点，每个节点一毫秒内可生成4095个不同的 id
*/

const (
	workerIDBits     = uint64(5) // 10 bit 工作机器ID中的 5bit workerID
	dataCenterIDBits = uint64(5) // 10 bit 工作机器ID中的 5bit dataCenterID
	sequenceBits     = uint64(12)

	maxWorkerID     = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxSequence     = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22) // timeLeft = workerIDBits + sequenceBits // 时间戳向左偏移量
	dataLeft = uint8(17) // dataLeft = dataCenterIDBits + sequenceBits
	workLeft = uint8(12) // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1589923200000) // 常量时间戳(毫秒)
)

var GlobalSnowFlakeWorker *SnowFlakeWorker

type SnowFlakeWorker struct {
	mu           sync.Mutex
	LastStamp    int64 // 记录上一次ID的时间戳
	WorkerID     int64 // 该节点的ID
	DataCenterID int64 // 该节点的 数据中心ID
	Sequence     int64 // 当前毫秒已经生成的ID序列号(从0 开始累加) 1毫秒内最多生成4096个ID
}

// SetUpSnowFlakeWorker 雪花算法支持最大 32 个服务器集群，单集群最大 32 台机器的部署方式，因此 worker，center 取值均为 0 - 31（5位整数）
func SetUpSnowFlakeWorker(worker,center int64) {
	GlobalSnowFlakeWorker = newSnowFlakeWorker(worker, center)
}

func newSnowFlakeWorker(workerID, dataCenterID int64) *SnowFlakeWorker {
	return &SnowFlakeWorker{
		WorkerID:     workerID,
		LastStamp:    0,
		Sequence:     0,
		DataCenterID: dataCenterID,
	}
}

func (w *SnowFlakeWorker) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *SnowFlakeWorker) NextID() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.nextID()
}

func (w *SnowFlakeWorker) nextID() (uint64, error) {
	timeStamp := w.getMilliSeconds()
	if timeStamp < w.LastStamp {
		return 0, errors.New("time is moving backwards,waiting until")
	}

	if w.LastStamp == timeStamp {

		w.Sequence = (w.Sequence + 1) & maxSequence

		if w.Sequence == 0 {
			for timeStamp <= w.LastStamp {
				timeStamp = w.getMilliSeconds()
			}
		}
	} else {
		w.Sequence = 0
	}

	w.LastStamp = timeStamp
	id := ((timeStamp - twepoch) << timeLeft) |
		(w.DataCenterID << dataLeft) |
		(w.WorkerID << workLeft) |
		w.Sequence

	return uint64(id), nil
}