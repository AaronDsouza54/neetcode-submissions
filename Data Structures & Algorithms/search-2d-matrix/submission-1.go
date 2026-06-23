func searchMatrix(matrix [][]int, target int) bool {
	low, high := 0, len(matrix) - 1

	for low <= high {
		var mid int = low + (high - low) / 2

		arr := matrix[mid]
		lastPosOfArr := len(arr) - 1

		if arr[0] <= target && arr[lastPosOfArr] >= target {
			return binarySearch(arr, target)
		}

		if target > arr[lastPosOfArr] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func binarySearch(arr []int, target int) bool {
	low, high := 0, len(arr) - 1

	for low <= high {
		var mid int = low + (high - low) / 2

		if arr[mid] == target {
			return true
		}

		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
