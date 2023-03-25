package topk

// HeavyKeeper: An Accurate Algorithm for Finding Top-k Elephant Flow (https://www.usenix.org/system/files/conference/atc18/atc18-gong.pdf)
// reference: https://github.com/go-kratos/aegis/blob/main/topk/heavykeeper.go

import (
	"container/heap"
	"hash/maphash"
	"math/rand"
)

const lookupSize int = 0x100 // 2^8=256
var generalHashSeed maphash.Seed = maphash.MakeSeed()

type HeapNode struct {
	freq int
	val  string
}

type KeeperNode struct {
	freq int
	fp   int
}

type HeavyKeeperHeap []HeapNode

var heapMap map[string]int

func (hkh HeavyKeeperHeap) Len() int           { return len(hkh) }
func (hkh HeavyKeeperHeap) Less(i, j int) bool { return hkh[i].freq < hkh[j].freq }

func (hkh HeavyKeeperHeap) Swap(i, j int) {
	heapMap[hkh[i].val] = j
	heapMap[hkh[j].val] = i
	hkh[i], hkh[j] = hkh[j], hkh[i]
}

func (hkh HeavyKeeperHeap) Top() HeapNode {
	return hkh[0]
}

func (hkh HeavyKeeperHeap) Find(target string) int { // should be O1
	if index, ok := heapMap[target]; ok {
		return index
	}
	return -1
}

func (hkh *HeavyKeeperHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*hkh = append(*hkh, HeapNode{0, x.(string)})
	heapMap[x.(string)] = len(*hkh) - 1
}

func (hkh *HeavyKeeperHeap) Pop() any {
	old := *hkh
	n := len(old)
	x := old[n-1]
	*hkh = old[0 : n-1]
	delete(heapMap, x.val)
	return x
}

func (hkh HeavyKeeperHeap) Get(i int) HeapNode {
	return hkh[i]
}

func (hkh *HeavyKeeperHeap) Set(i int, freq int) {
	(*hkh)[i].freq = freq
}

func (hkh *HeavyKeeperHeap) Fade() {
	for i := range *hkh {
		(*hkh)[i].freq = (*hkh)[i].freq >> 1
	}
}

type HeavyKeeper struct {
	minHeap *HeavyKeeperHeap
	k       int
	rows    uint64
	cols    uint64
	keeper  [][]KeeperNode
	seeds   []maphash.Seed
	table   []float64
	total   uint64
	minFreq int
}

func NewHeavyKeeper(k int, rows uint64, cols uint64, decay float64, minFreq int) *HeavyKeeper {
	heap.Init(&HeavyKeeperHeap{})
	keeper := make([][]KeeperNode, rows)
	seeds := make([]maphash.Seed, rows)
	for i := range keeper {
		keeper[i] = make([]KeeperNode, cols)
		seeds[i] = maphash.MakeSeed()
	}
	table := make([]float64, lookupSize)
	table[0] = decay
	for i := 1; i < lookupSize; i++ {
		table[i] = table[i-1] * decay // decay: 1.02-1.08
	}
	return &HeavyKeeper{
		minHeap: &HeavyKeeperHeap{},
		k:       k,
		rows:    rows,
		cols:    cols,
		keeper:  keeper,
		seeds:   seeds,
		table:   table,
		total:   0,
		minFreq: minFreq,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (hk *HeavyKeeper) Add(s string) (string, bool) {
	var flag bool = false
	var sIdx int
	if sIdx = hk.minHeap.Find(s); sIdx != -1 {
		flag = true
	}
	var maxFreq int = 0
	sBytes := []byte(s)
	hv := maphash.Bytes(generalHashSeed, sBytes)
	var fp int = int(hv >> 32)
	for i, row := range hk.keeper {
		var hi = maphash.Bytes(hk.seeds[i], sBytes) % hk.rows
		if row[hi].freq == 0 {
			row[hi].freq = 1
			row[hi].fp = fp
			if maxFreq < 1 {
				maxFreq = 1
			}
		} else if row[hi].fp == fp {
			if flag || row[hi].freq <= hk.minHeap.Top().freq {
				row[hi].freq++
				if row[hi].freq > maxFreq {
					maxFreq = row[hi].freq
				}
			}
		} else if rand.Int31()%int32(hk.table[max(row[hi].freq, lookupSize-1)]) == 0 {
			row[hi].freq--
			if row[hi].freq == 0 {
				row[hi].freq = 1
				row[hi].fp = fp
				if maxFreq < 1 {
					maxFreq = 1
				}
			}
		}
	}
	hk.total++
	if maxFreq < hk.minFreq || (hk.minHeap.Len() == hk.k && maxFreq < hk.minHeap.Top().freq) {
		return "", false
	}
	if flag {
		hk.minHeap.Set(sIdx, maxFreq)
		heap.Fix(hk.minHeap, sIdx)
		return "", true
	}
	heap.Push(hk.minHeap, HeapNode{maxFreq, s})
	var emit string
	if hk.minHeap.Len() > hk.k {
		emit = heap.Pop(hk.minHeap).(HeapNode).val
	}
	return emit, true
}

func (hk *HeavyKeeper) Fade() {
	for _, row := range hk.keeper {
		for i := range row {
			if row[i].freq > 0 {
				row[i].freq >>= 1
			}
		}
	}
	hk.minHeap.Fade()
	hk.total >>= 1
}

func (hk *HeavyKeeper) GetTopK() []HeapNode {
	var res []HeapNode
	for i := 0; i < hk.minHeap.Len(); i++ {
		res = append(res, hk.minHeap.Get(i))
	}
	return res
}
