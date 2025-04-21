package questions

func ModifyValue(p *int) {
	*p += 10 // 修改指针指向的值为10
}

func ModifySlice(s *[]int) {
	for _, v := range *s { // 遍历切片中的每个元素
		v *= 2 // 修改每个元素的值为10
	}
}
