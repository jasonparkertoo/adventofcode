package day09

import "adventofcode.dev/utils"

// Block represents a contiguous region on the hard drive.
// FileId < 0 denotes a free block, otherwise it is a file with that ID.
type Block struct {
	FileId int
	Len    int
}

// IsFree reports whether the block represents free space.
// A block is considered free if its FileId is negative.
func (b Block) IsFree() bool {
	return b.FileId < 0
}

type Harddrive struct {
	DataMap string
}

const (
	CompactLeft   string = "L"
	CompactNormal string = "N"
)

// ChecksumHarddrive calculates the checksum of a hard drive represented by the given *utils.Data.
// The method string determines which compaction algorithm to use:
//
//	"L" – leftward compaction (compactLeft),
//	"N" – normal compaction (compact).  Any other value causes a panic.
//
// The checksum is computed over the compacted block sequence.
func ChecksumHarddrive(d *utils.Data, method string) int {
	dataMap := d.Lines()[0]

	var blocks []Block
	switch method {
	case "L":
		blocks = compactLeft(dataMap)
	case "N":
		blocks = compact(dataMap)
	default:
		panic("unknown compact method requested")
	}
	checksum := Checksum(blocks)
	return checksum
}

// blocks converts a data map string into a slice of Block. Each character in the
// string represents a block length; even indices correspond to files, odd
// indices to free space. Zero-length blocks are ignored.
func blocks(dataMap string) []Block {
	fileId := 0

	blocks := make([]Block, 0)
	for i, r := range dataMap {
		length := r - '0'
		if length == 0 {
			continue
		}
		if i%2 == 0 {
			blocks = append(blocks, Block{fileId, int(length)})
			fileId++
		} else {
			blocks = append(blocks, Block{-1, int(length)})
		}
	}
	return blocks
}

// compact performs the normal compaction algorithm on the data map string. It
// first expands the blocks into a linear representation, then moves free
// blocks towards the right while preserving file order. The result is a slice
// of compressed blocks.
func compact(dataMap string) []Block {
	b := blocks(dataMap)

	ds := make([]int, 0)
	for i := range b {
		for range b[i].Len {
			ds = append(ds, b[i].FileId)
		}
	}

	left, right := 0, len(ds)-1
	for {
		for left < right && ds[left] != -1 {
			left++
		}
		for left < right && ds[right] == -1 {
			right--
		}
		if left >= right {
			break
		}

		ds[left] = ds[right]
		ds[right] = -1
		left++
		right--
	}
	result := make([]Block, 0)
	if len(ds) == 0 {
		return []Block{}
	}

	count := 1
	current := ds[0]
	for i := 1; i < len(ds); i++ {
		next := ds[i]
		if next == current {
			count++
		} else {
			result = append(result, Block{current, count})
			current = next
			count = 1
		}
	}
	return append(result, Block{current, count})
}

// compactLeft performs a leftward compaction on the data map string. It moves
// files towards the beginning of the map by repeatedly locating the next
// suitable free block and shifting the file into that position, merging
// adjacent free spaces when necessary.
func compactLeft(dataMap string) []Block {
	// Start from parsed blocks
	blocks := append([]Block(nil), blocks(dataMap)...)

	// Determine highest file ID
	maxID := -1
	for _, b := range blocks {
		if !b.IsFree() && b.FileId > maxID {
			maxID = b.FileId
		}
	}

	// Move files leftward in reverse ID order
	for moveID := maxID; moveID >= 0; moveID-- {
		fileIndex := -1
		var fileBlock Block

		// Locate file
		for i, b := range blocks {
			if b.FileId == moveID {
				fileIndex = i
				fileBlock = b
				break
			}
		}
		if fileIndex == -1 {
			continue
		}

		// Find earliest suitable free block
		targetIndex := -1
		for i := 0; i < fileIndex; i++ {
			b := blocks[i]
			if b.IsFree() && b.Len >= fileBlock.Len {
				targetIndex = i
				break
			}
		}
		if targetIndex == -1 {
			continue
		}

		freeBlock := blocks[targetIndex]
		var updated []Block

		// Copy before target
		updated = append(updated, blocks[:targetIndex]...)

		// Insert moved file
		updated = append(updated, Block{FileId: fileBlock.FileId, Len: fileBlock.Len})

		// Remaining free space (if any)
		if remaining := freeBlock.Len - fileBlock.Len; remaining > 0 {
			updated = append(updated, Block{FileId: -1, Len: remaining})
		}

		// Copy between target and file
		updated = append(updated, blocks[targetIndex+1:fileIndex]...)

		// Replace file with free space
		updated = append(updated, Block{FileId: -1, Len: fileBlock.Len})

		// Copy after file
		updated = append(updated, blocks[fileIndex+1:]...)

		// Merge adjacent free blocks
		var merged []Block
		for _, b := range updated {
			if len(merged) > 0 && merged[len(merged)-1].IsFree() && b.IsFree() {
				last := merged[len(merged)-1]
				merged[len(merged)-1] = Block{FileId: -1, Len: last.Len + b.Len}
			} else {
				merged = append(merged, b)
			}
		}

		blocks = merged
	}

	return blocks
}

// Checksum calculates the sum of (position * file ID) for each byte in all
// non‑free blocks of the slice. Free space is skipped, but its length is
// accumulated to offset positions of subsequent blocks.
func Checksum(blocks []Block) int {
	sum, pos := 0, 0
	for _, b := range blocks {
		if b.IsFree() {
			pos += b.Len
			continue
		} else {
			start := pos
			pos += b.Len
			for i := range b.Len {
				sum += (start + i) * b.FileId
			}
		}
	}
	return sum
}
