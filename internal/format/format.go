package format

func DepthAlign(depth int) string {
	var output string
	for i := 0; i < depth; i++ {
		output += "\t"
	}
	return output
}
