type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	message := ""

	if len(strs) == 0 {
		return "<//{empty}{}/>"
	}

	for _, v := range strs {
		encodedLength := "<//{}{}/>"
		message += encodedLength + v
	}

	return message
}

func (s *Solution) Decode(encoded string) []string {
	// strs = []string
	pattern := "<//{}{}/>"

	if encoded == "<//{empty}{}/>" {
		return []string{}
	}

	if len(encoded) == 0 {
		return []string{""}
	}

	return strings.Split(encoded, pattern)[1:]

}