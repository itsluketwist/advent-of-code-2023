package utils

type BucketHeap[T any] struct {
	buckets map[int][]T
	minimum int
	used    bool
}

func (bh *BucketHeap[T]) Add(item T, value int) {
	if !bh.used {
		bh.buckets = make(map[int][]T)
		bh.minimum = value
		bh.used = true
	}

	bh.buckets[value] = append(bh.buckets[value], item)
	if value < bh.minimum {
		bh.minimum = value
	}
}

func (bh *BucketHeap[T]) Pop() (T, int) {
	item := bh.buckets[bh.minimum][0]
	value := bh.minimum
	bh.buckets[bh.minimum] = bh.buckets[bh.minimum][1:]

	if len(bh.buckets[bh.minimum]) == 0 {
		delete(bh.buckets, bh.minimum)
		bh.minimum = 999999999
		for bucketValue := range bh.buckets {
			if bucketValue < bh.minimum {
				bh.minimum = bucketValue
			}
		}
	}

	return item, value
}
