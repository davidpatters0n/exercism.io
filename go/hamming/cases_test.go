package hamming

// Source: exercism/problem-specifications
// Commit: 4119671 Hamming: Add a tests to avoid wrong recursion solution (#1450)
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{ // empty strands
		"",
		"",
		0,
		false,
	},
	{ // single letter identical strands
		"A",
		"A",
		0,
		false,
	},
	{ // single letter different strands
		"G",
		"T",
		1,
		false,
	},
	{ // long identical strands
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		false,
	},
	{ // long different strands
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
		false,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		0,
		true,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		0,
		true,
	},
	{ // disallow left empty strand
		"",
		"G",
		0,
		true,
	},
	{ // disallow right empty strand
		"G",
		"",
		0,
		true,
	},
}

/*

1. Throw an error when:

	1. s1 || s2 are empty
	2. If s1 || s2 are not the same length

2. Take the string split it.
	2.1 Set a local variable count
	2.2 loop over the array of s1 with an index


	s1.chars.each_with_index{|c, idx| c != s2[idx] ? count += 1 : 0}
*/
