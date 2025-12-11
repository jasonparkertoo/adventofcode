package y2024

type Block struct {
	FileId int
	Len    int
}

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

func ChecksumHarddrive(hd Harddrive, method string) int {
	var blocks []Block
	switch method {
	case "L":
		blocks = compactLeft(hd)
	case "N":
		blocks = compact(hd)
	default:
		panic("unknown compact method requested")
	}
	checksum := Checksum(blocks)
	return checksum
}

func blocks(hd Harddrive) []Block {
	fileId := 0

	blocks := make([]Block, 0)
	for i, r := range hd.DataMap {
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

func compact(hd Harddrive) []Block {
	b := blocks(hd)

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

func compactLeft(hd Harddrive) []Block {
	// Start from parsed blocks
	blocks := append([]Block(nil), blocks(hd)...) // shallow copy

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
