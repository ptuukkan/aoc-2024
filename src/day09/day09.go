package day09

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Block struct {
	Id     int
	Length int
}

func (b *Block) Checksum(index int) (int, int) {
	if b.Id == -1 {
		return 0, index + b.Length
	}
	checksum := 0
	for i := 0; i < b.Length; i++ {
		checksum += b.Id * index
		index++
	}
	return checksum, index
}

func createDiskMap(input string) []int {
	id := 0
	var diskMap []int

	for i, c := range input {
		length, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("error converting number")
			fmt.Println(i)
		}
		blockId := -1
		if i%2 == 0 {
			blockId = id
			id++
		}
		block := slices.Repeat([]int{blockId}, length)
		diskMap = append(diskMap, block...)
	}

	return diskMap
}

func createBlockDiskMap(input string) ([]Block, int) {
	id := -1
	var diskMap []Block

	for i, c := range input {
		length, err := strconv.Atoi(string(c))
		if length == 0 {
			continue
		}
		if err != nil {
			fmt.Println("error converting number")
			fmt.Println(i)
		}
		blockId := -1
		if i%2 == 0 {
			id++
			blockId = id
		}
		block := Block{Id: blockId, Length: length}
		diskMap = append(diskMap, block)
	}

	return diskMap, id
}

func Part1(input string) string {
	input = strings.Trim(input, "\n")
	diskMap := createDiskMap(input)
	start := 0
	end := len(diskMap) - 1
	for start < end {
		for diskMap[start] != -1 {
			start++
		}
		for diskMap[end] == -1 {
			end--
		}
		for diskMap[start] == -1 && diskMap[end] != -1 && start < end {
			diskMap[start] = diskMap[end]
			diskMap[end] = -1
			start++
			end--
		}
	}

	checksum := 0

	for i, block := range diskMap {
		if block == -1 {
			continue
		}
		checksum += i * block
	}
	return strconv.Itoa(checksum)
}

func Part2(input string) string {
	input = strings.Trim(input, "\n")
	diskMap, currentId := createBlockDiskMap(input)

	for currentId != 0 {
		blockIndex := slices.IndexFunc(diskMap, func(b Block) bool {
			return b.Id == currentId
		})

		if blockIndex == -1 {
			fmt.Println("Cannot find current block")
			return ""
		}

		freeIndex := slices.IndexFunc(diskMap, func(b Block) bool {
			return b.Id == -1 && b.Length >= diskMap[blockIndex].Length
		})

		if freeIndex == -1 || freeIndex >= blockIndex {
			currentId--
			continue
		}

		diskMap[freeIndex].Id = diskMap[blockIndex].Id
		diskMap[blockIndex].Id = -1

		if diskMap[freeIndex].Length > diskMap[blockIndex].Length {
			overflow := diskMap[freeIndex].Length - diskMap[blockIndex].Length
			diskMap[freeIndex].Length -= overflow
			overflowBlock := Block{Id: -1, Length: overflow}
			diskMap = append(diskMap[:freeIndex+1], append([]Block{overflowBlock}, diskMap[freeIndex+1:]...)...)
		}
		currentId--
	}

	checksum := 0
	i := 0
	for _, block := range diskMap {
		var blockChecksum int
		blockChecksum, i = block.Checksum(i)
		checksum += blockChecksum
	}

	return strconv.Itoa(checksum)
}
