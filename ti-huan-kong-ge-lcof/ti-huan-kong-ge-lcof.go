package ti_huan_kong_ge_lcof

func replaceSpace(s string) string {
	r := []byte("%20")
	buf := bytes.NewBufferString("")
	for _, v := range s {
		if string(v) != " " {
			buf.Write([]byte(string(v)))
		} else {
			buf.Write(r)
		}
	}

	return buf.String()
}