package tcid

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Validate validates turkish citizen ID number. Returns false if invalid, otherwise returns true.
func Validate(id string) bool {
	if len(id) != 11 {
		return false
	}
	nums, err := breakString(id)
	if err != nil {
		return false
	}
	if (*nums)[0] < 1 || (*nums)[len(*nums)-1]%2 != 0 || len(*nums) != 11 {
		return false
	}
	even := (*nums)[1] + (*nums)[3] + (*nums)[5] + (*nums)[7]
	odd := (*nums)[0] + (*nums)[2] + (*nums)[4] + (*nums)[6] + (*nums)[8]
	d10 := (odd*7 - even) % 10
	total := (odd + even + (*nums)[9]) % 10
	if d10 != (*nums)[9] || total != (*nums)[10] {
		return false
	}
	return true
}

func breakString(s string) (*[]int64, error) {
	var nums []int64
	ss := strings.Split(s, "")
	for _, n := range ss {
		num, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return &nums, nil
}

// Generate generates 11 digit valid turkish identity number.
func Generate() string {
	var tc string
	rand.Seed(time.Now().Unix())
	nums := make([]int, 11)
	num1 := rand.Intn(10)
	if num1 == 0 {
		num1++
	}
	nums[0] = num1
	for i := 1; i < 9; i++ {
		nums[i] = rand.Intn(10)
	}
	nums[9] = ((nums[0]+nums[2]+nums[4]+nums[6]+nums[8])*7 - (nums[1] + nums[3] + nums[5] + nums[7])) % 10
	nums[10] = (nums[0] + nums[2] + nums[4] + nums[6] + nums[8] + nums[1] + nums[3] + nums[5] + nums[7] + nums[9]) % 10
	for i := 0; i < len(nums); i++ {
		tc += strconv.Itoa(nums[i])
	}
	return tc
}
