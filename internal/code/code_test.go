package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operants []int
		expected []byte
	}{
		{OpConstant, []int{65_534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
		{OpGetLocal, []int{255}, []byte{byte(OpGetLocal), 255}},
		{OpClosure, []int{65534, 255}, []byte{byte(OpClosure), 255, 254, 255}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operants...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d",
				len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d",
					i, b, instruction[i])
			}
		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpAdd),
		Make(OpGetLocal, 1),
		Make(OpConstant, 2),
		Make(OpConstant, 65_535),
		Make(OpClosure, 65_535, 255),
	}

	expected := `0000 OpAdd
0001 OpGetLocal 1
0003 OpConstant 2
0006 OpConstant 65535
0009 OpClosure 65535 255
`

	concatted := Instructions{}
	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot =%q",
			expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operants  []int
		bytesRead int
	}{
		{OpConstant, []int{65_535}, 2},
		{OpGetLocal, []int{255}, 1},
		{OpClosure, []int{65_535, 255}, 3},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operants...)

		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}

		operantsRead, n := ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.bytesRead, n)
		}

		for i, want := range tt.operants {
			if operantsRead[i] != want {
				t.Errorf(
					"operand wrong. want=%d, got=%d",
					want,
					operantsRead[i],
				)
			}
		}
	}
}
