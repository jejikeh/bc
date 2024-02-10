package bytecode

// ../samples/hello.b
var helloKinds = []Kind{
	// ++++++++
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	// [
	MarkJump,
	// >
	Right,
	// ++++
	Increment,
	Increment,
	Increment,
	Increment,
	// [>++
	MarkJump,
	Right,
	Increment,
	Increment,
	// >+++
	Right,
	Increment,
	Increment,
	Increment,
	// >+++
	Right,
	Increment,
	Increment,
	Increment,
	// >+<<<<-]
	Right,
	Increment,
	Left,
	Left,
	Left,
	Left,
	Decrement,
	Jump,
	// >+>+>-
	Right,
	Increment,
	Right,
	Increment,
	Right,
	Decrement,
	// >>+
	Right,
	Right,
	Increment,
	// [<]
	MarkJump,
	Left,
	Jump,
	// <-]
	Left,
	Decrement,
	Jump,
	// >>.>
	Right,
	Right,
	Output,
	Right,
	// ---.
	Decrement,
	Decrement,
	Decrement,
	Output,
	// +++++++
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	Increment,
	// ..+++
	Output,
	Output,
	Increment,
	Increment,
	Increment,
	// .>>.<
	Output,
	Right,
	Right,
	Output,
	Left,
	// -.<.
	Decrement,
	Output,
	Left,
	Output,
	// +++.
	Increment,
	Increment,
	Increment,
	Output,
	// ------.
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Output,
	// --------.
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Decrement,
	Output,
	// >>+.>++.
	Right,
	Right,
	Increment,
	Output,
	Right,
	Increment,
	Increment,
	Output,
	End,
}

var helloIntermediateRepresentation = []IntermediateRepresentation{
	{
		Token: Token{Kind: Increment, Position: 0},
		Op:    8,
	},
	{
		Token: Token{Kind: MarkJump, Position: 8},
		Op:    0,
	},
	{
		Token: Token{Kind: Right, Position: 9},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 10},
		Op:    4,
	},
	{
		Token: Token{Kind: MarkJump, Position: 14},
		Op:    0,
	},
	{
		Token: Token{Kind: Right, Position: 15},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 16},
		Op:    2,
	},
	{
		Token: Token{Kind: Right, Position: 18},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 19},
		Op:    3,
	},
	{
		Token: Token{Kind: Right, Position: 22},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 23},
		Op:    3,
	},
	{
		Token: Token{Kind: Right, Position: 26},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 27},
		Op:    1,
	},
	{
		Token: Token{Kind: Left, Position: 28},
		Op:    4,
	},
	{
		Token: Token{Kind: Decrement, Position: 32},
		Op:    1,
	},
	{
		Token: Token{Kind: Jump, Position: 33},
		Op:    100,
	},
	{
		Token: Token{Kind: Right, Position: 34},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 35},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 36},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 37},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 38},
		Op:    1,
	},
	{
		Token: Token{Kind: Decrement, Position: 39},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 40},
		Op:    2,
	},
	{
		Token: Token{Kind: Increment, Position: 42},
		Op:    1,
	},
	{
		Token: Token{Kind: MarkJump, Position: 43},
		Op:    1,
	},
	{
		Token: Token{Kind: Left, Position: 44},
		Op:    1,
	},
	{
		Token: Token{Kind: Jump, Position: 45},
		Op:    1,
	},
	{
		Token: Token{Kind: Left, Position: 46},
		Op:    1,
	},
	{
		Token: Token{Kind: Decrement, Position: 47},
		Op:    1,
	},
	{
		Token: Token{Kind: Jump, Position: 48},
		Op:    1000,
	},
	{
		Token: Token{Kind: Right, Position: 49},
		Op:    2,
	},
	{
		Token: Token{Kind: Output, Position: 51},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 52},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 53},
		Op:    3,
	},
	{
		Token: Token{Kind: Output, Position: 56},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 57},
		Op:    7,
	},
	{
		Token: Token{Kind: Output, Position: 64},
		Op:    2,
	},
	{
		Token: Token{Kind: Increment, Position: 66},
		Op:    3,
	},
	{
		Token: Token{Kind: Output, Position: 69},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 70},
		Op:    2,
	},
	{
		Token: Token{Kind: Output, Position: 72},
		Op:    1,
	},
	{
		Token: Token{Kind: Left, Position: 73},
		Op:    1,
	},
	{
		Token: Token{Kind: Decrement, Position: 74},
		Op:    1,
	},
	{
		Token: Token{Kind: Output, Position: 75},
		Op:    1,
	},
	{
		Token: Token{Kind: Left, Position: 76},
		Op:    1,
	},
	{
		Token: Token{Kind: Output, Position: 77},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 78},
		Op:    3,
	},
	{
		Token: Token{Kind: Output, Position: 81},
		Op:    1,
	},
	{
		Token: Token{Kind: Decrement, Position: 82},
		Op:    6,
	},
	{
		Token: Token{Kind: Output, Position: 88},
		Op:    1,
	},
	{
		Token: Token{Kind: Decrement, Position: 89},
		Op:    8,
	},
	{
		Token: Token{Kind: Output, Position: 97},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 98},
		Op:    2,
	},
	{
		Token: Token{Kind: Increment, Position: 100},
		Op:    1,
	},
	{
		Token: Token{Kind: Output, Position: 101},
		Op:    1,
	},
	{
		Token: Token{Kind: Right, Position: 102},
		Op:    1,
	},
	{
		Token: Token{Kind: Increment, Position: 103},
		Op:    2,
	},
	{
		Token: Token{Kind: Output, Position: 105},
		Op:    1,
	},
	{
		Token: Token{Kind: End, Position: 107},
		Op:    0,
	},
}
