package _9

import (
	"fmt"
	"strconv"
)

type Amphipod struct {
	disk              []int
	currentPosition   int
	lastBlockPosition int
}

func (a *Amphipod) ParseInput(input string) error {
	// example `2323` represents
	// - two blocks with index 1
	// - three empty blocks
	// - two blocks with index 2
	// - three empty blocks
	// I need to iterate over the string in steps of 2 but only increment the index by 1

	index := 0
	for i, nrOfBlocks := range input {
		// parse nrOfBlocks into an int
		nr, err := strconv.Atoi(string(nrOfBlocks))
		if err != nil {
			return err
		}

		if i%2 == 0 {
			// we have a block
			for j := 0; j < nr; j++ {
				a.disk = append(a.disk, index)
			}
			// increase index for the next block
			index++
		} else {
			// we have an empty block, adding -1 to the disk
			for j := 0; j < nr; j++ {
				a.disk = append(a.disk, -1)
			}
		}
	}
	a.currentPosition = 0
	a.lastBlockPosition = len(a.disk) - 1
	return nil
}

func (a *Amphipod) Print() {
	for _, block := range a.disk {
		if block < 0 {
			// whitespace
			fmt.Print(".")
		} else {
			fmt.Print(block)
		}
	}
	fmt.Println()
}

// findNextSpace returns the index of the next empty space in the disk
// returns the startingPosition of the space and the length of the space block
func (a *Amphipod) findNextSpace() (int, int) {
	for i := a.currentPosition; i < len(a.disk); i++ {
		if a.disk[i] < 0 {
			// starting-position
			length := 0
			for j := i; j < len(a.disk); j++ {
				if a.disk[j] >= 0 {
					break
				}
				length++
			}
			return i, length
		}
	}
	return -1, 0
}

// findLastBlock returns the index of the last block in the disk that is not empty space
func (a *Amphipod) findLastBlock(advanced bool) (int, int) {
	// simple mode, just find the last written block
	if !advanced {
		for i := a.lastBlockPosition; i >= a.currentPosition; i-- {
			if a.disk[i] >= 0 {
				return i, 1
			}
		}
	}

	// advanced mode
	// find the last written block that has the requested block size
	for i := a.lastBlockPosition; i >= a.currentPosition; i-- {
		blockIndex := a.disk[i]
		if blockIndex >= 0 {
			// get the block size
			blockSize := 0
			for j := i; j >= 0; j-- {
				if a.disk[j] != blockIndex {
					break
				}
				blockSize++
				i = j
			}
			// switch block-index to the first block of the block
			return i, blockSize
		}
	}

	return -1, -1
}

// AdvancedFragment also defragments the disk but moving whole blockes instead of single blocks
func (a *Amphipod) AdvancedFragment() {
	// iterate over all blocks from the end
	for a.lastBlockPosition > a.currentPosition {
		// find the next block from the end
		blockIndex, size := a.findLastBlock(true)
		//log.Printf("find space for block id: %d, index: %d, size: %d", a.disk[blockIndex], blockIndex, size)
		// iterate through space from the beginning until you find a space that fits the block
		for i := 0; i < blockIndex; i++ {
			// find the next space
			emptySpaceIndex, emptySpaceLength := a.findNextSpace()
			if emptySpaceIndex > blockIndex {
				break
			}

			if emptySpaceLength >= size {
				// move block
				for j := 0; j < size; j++ {
					a.disk[emptySpaceIndex+j], a.disk[blockIndex+j] = a.disk[blockIndex+j], a.disk[emptySpaceIndex+j]
				}
				break
			} else {
				a.currentPosition = emptySpaceIndex + 1
			}
		}
		// set lastBlockPosition to the last block
		a.lastBlockPosition = blockIndex - 1

		// reset current position to the first empty space
		a.currentPosition = 0
		firstFreeSpaceIndex, _ := a.findNextSpace()
		a.currentPosition = firstFreeSpaceIndex

		//a.Print()
	}
}

func (a *Amphipod) Fragment() {
	// loop over the disk while the "currentPosition" is less than the "lastBlockPosition"
	for a.currentPosition < a.lastBlockPosition {
		// find first empty space index
		emptySpaceIndex, _ := a.findNextSpace()
		lastBlockIndex, _ := a.findLastBlock(false)
		//log.Printf("empty space index: %d, last block index: %d", emptySpaceIndex, lastBlockIndex)
		if emptySpaceIndex >= lastBlockIndex {
			// no more empty spaces or blocks to move
			break
		}
		// switch the contents of those two blocks
		a.disk[emptySpaceIndex], a.disk[lastBlockIndex] = a.disk[lastBlockIndex], a.disk[emptySpaceIndex]
		// update positions
		a.currentPosition = emptySpaceIndex
		a.lastBlockPosition = lastBlockIndex
	}
}

// Checksum calculates a checksum by multiplying the block-number with the index of the block and
// summing all those values
func (a *Amphipod) Checksum() int {
	sum := 0
	for i, block := range a.disk {
		// skip empty blocks
		if block < 0 {
			continue
		}
		sum += block * i
	}
	return sum
}

func NewAmphipod() *Amphipod {
	return &Amphipod{
		disk:              make([]int, 0),
		currentPosition:   0,
		lastBlockPosition: -1,
	}
}
