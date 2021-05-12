package roaring64

// FastAnd computes the intersection between many bitmaps quickly
// Compared to the And function, it can take many bitmaps as input, thus saving the trouble
// of manually calling "And" many times.
func FastAnd(bitmaps ...*Bitmap) *Bitmap {
	if len(bitmaps) == 0 {
		return NewBitmap()
	} else if len(bitmaps) == 1 {
		return bitmaps[0].Clone()
	}
	answer := And(bitmaps[0], bitmaps[1])
	for _, bm := range bitmaps[2:] {
		answer.And(bm)
	}
	return answer
}

// FastOr computes the union between many bitmaps quickly, as opposed to having to call Or repeatedly.
// It might also be faster than calling Or repeatedly.
func FastOr(bitmaps ...*Bitmap) *Bitmap {
	if len(bitmaps) == 0 {
		return NewBitmap()
	} else if len(bitmaps) == 1 {
		return bitmaps[0].Clone()
	}
	//answer := lazyOR(bitmaps[0], bitmaps[1])
	answer := Or(bitmaps[0], bitmaps[1])
	for _, bm := range bitmaps[2:] {
		//answer = answer.lazyOR(bm)
		answer.Or(bm)
	}
	// here is where repairAfterLazy is called.
	//answer.repairAfterLazy()
	return answer
}

func FastOrSerial(bitmaps ...[]byte) ([]byte, error) {
	if len(bitmaps) == 0 {
		return NewBitmap().MarshalBinary()
	} else if len(bitmaps) == 1 {
		return bitmaps[0], nil
	}
	answer, err := OrSerial(bitmaps[0], bitmaps[1])
	if err != nil {
		return answer, err
	}
	for _, bm := range bitmaps[2:] {
		answer, err = OrSerial(answer, bm)
		if err != nil {
			return answer, err
		}
	}
	return answer, nil

}
