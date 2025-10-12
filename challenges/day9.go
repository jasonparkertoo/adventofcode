package challenges

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

func ChecksumHarddrive(hd Harddrive) int {
	blocks := blocks(hd)
	compacted := compact(blocks)
	checksum := ChecksumBlocks(compacted)
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

func compact(b []Block) []Block {
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

func ChecksumBlocks(blocks []Block) int {
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
