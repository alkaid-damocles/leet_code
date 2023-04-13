package code

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i := 0
	j := 0

	if len(nums1) == 0 {
		nums1 = append(nums1, 0, math.MaxInt)
	}

	if len(nums2) == 0 {
		nums2 = append(nums2, 0, math.MaxInt)
	}

	middle := (len(nums1) + len(nums2)) / 2
	for i != len(nums1) && j != len(nums2) {

		if i+j+1 == middle {
			break
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}

	}

	if (len(nums1)+len(nums2))%2 == 1 {
		if nums1[i] > nums2[j] {
			return float64(nums1[i])
		} else {
			return float64(nums2[j])
		}
	} else {
		return (float64(nums1[i]) + float64(nums2[j])) / 2
	}
}
