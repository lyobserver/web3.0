package questions

func ModifyValue(p *int) {
	*p += 10 // 修改指针指向的值为10
}

func ModifySlice(s *[]int) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] *= 2
	}
}
