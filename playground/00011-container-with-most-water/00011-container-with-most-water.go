package _00011_container_with_most_water

func maxArea(height []int) int {
	max, s, e := 0, 0, len(height)-1

	for s < e { // crash to true
		width := e - s
		high := 0
		if height[s] < height[e] { // min height s
			high = height[s]
			s++
		} else { // min height e
			high = height[e]
			e--
		}
		volume := width * high // now volume
		if volume > max {
			max = volume
		}
	}
	return max
}
