package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	rn_len := len(ransomNote)
	m_len := len(magazine)

	map_test := make(map[string]int)

	if rn_len == 0 || m_len == 0 {
		return false
	}

	counter := 0
	for i := 0; i < rn_len; i++ {
		current_letter := string(ransomNote[i])
		map_test[current_letter] = 1
		counter = 0
		for j := 0; j < m_len; j++ {
			if ransomNote[i] == magazine[j] {
				counter++
				map_test[current_letter] = map_test[current_letter] + 1
			}
		}
		_, isPresent := map_test[string(ransomNote[i])]
		if counter == 0 || isPresent {
			return false
		}
	}

	return true
}
