package main

import (
	"log"

	"github.com/asch/buse/lib/go/buse"
)

type driver struct {
}

const bt byte = 100

func (d driver) BuseRead(sector, length int64, chunk []byte) error {
	for i := int64(0); i < length; i++ {
		chunk[i] = bt
	}
	return nil
}

func (d driver) BuseWrite(writes int64, chunk []byte) error {
	return nil
}

func (d driver) BusePreRun() {
	log.Print("buse pre-run")
}

func (d driver) BusePostRemove() {
	log.Print("buse post-remove")
}

const BLOCK_SIZE = 512

func main() {
	dev, err := buse.New(driver{}, buse.Options{
		Durable:        false,
		WriteChunkSize: BLOCK_SIZE * 10,
		BlockSize:      BLOCK_SIZE,
		IOMin:          BLOCK_SIZE,
		IOOpt:          BLOCK_SIZE * 2,
		Threads:        1,
		Major:          0,
		WriteShmSize:   512 * 20,
		ReadShmSize:    512 * 20,
		Size:           512 * 1e4,
		CollisionArea:  512 * 100,
		QueueDepth:     int64(256 * 20), // todo добавить в бусе чтобы он чекал != 0; > много
		Scheduler:      false,
		CPUsPerNode:    1,
	})
	if err != nil {
		log.Fatalf("failed to init device: %s", err)
	}
	dev.Run()
}
