package abicoder

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

func TestEncodeUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	eight := uint8(8)

	packed, err := Encode(arg, eight)
	require.NoError(t, err)

	expected := make([]byte, 32)
	expected[31] = eight

	require.Equal(t, expected, packed)
}

func TestEncodeFailUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	notUint8 := uint16(8)

	_, err = Encode(arg, notUint8)
	require.Error(t, err)

	stillNotUint8 := 257
	_, err = Encode(arg, stillNotUint8)
	require.Error(t, err)
}

func TestEncodeStruct(t *testing.T) {
	abiJson := `{
          "components": [
            {
              "internalType": "address",
              "name": "A",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "B",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "C",
              "type": "uint8"
            }
          ],
          "internalType": "struct IEntityManager.VoterAddresses",
          "name": "Neki",
          "type": "tuple"
        }`

	arg := abi.Argument{}
	err := arg.UnmarshalJSON([]byte(abiJson))
	require.NoError(t, err)

	type neki struct {
		A uint8
		B uint8
		C uint8
	}

	n := neki{8, 8, 8}

	encoded, err := Encode(arg, n)
	require.NoError(t, err)

	expected := make([]byte, 96)
	nonzero := []int{31, 63, 95}

	for _, j := range nonzero {
		expected[j] = uint8(8)
	}

	require.Equal(t, expected, encoded)

	// Fail

	nd := struct {
		a string
	}{"a"}

	_, err = Encode(arg, nd)
	require.Error(t, err)
}

func TestEncodeString(t *testing.T) {
	stringType, err := abi.NewType("string", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "string",
		Type:    stringType,
		Indexed: false,
	}

	s := "abc"

	packed, err := Encode(arg, s)
	require.NoError(t, err)

	expected := make([]byte, 96)
	expected[31] = uint8(32)
	expected[63] = uint8(len(s))
	copy(expected[64:68], []byte("abc"))

	require.Equal(t, expected, packed)

	decoded, err := Decode[string](arg, packed)
	require.NoError(t, err)
	require.Equal(t, s, decoded)
}

// TestEncodeRecoversPanic verifies that Encode returns an error instead of
// panicking when abi.Arguments.Pack panics on a nil *big.Int, both directly
// and as a struct field.
func TestEncodeRecoversPanic(t *testing.T) {
	uint256Type, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	uint256Arg := abi.Argument{
		Name: "x",
		Type: uint256Type,
	}

	tupleType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "A", Type: "uint256"},
	})
	require.NoError(t, err)

	tupleArg := abi.Argument{
		Name: "t",
		Type: tupleType,
	}

	type withBigInt struct {
		A *big.Int
	}

	tests := []struct {
		name string
		arg  abi.Argument
		data any
	}{
		{
			name: "nil *big.Int",
			arg:  uint256Arg,
			data: (*big.Int)(nil),
		},
		{
			name: "nil *big.Int struct field",
			arg:  tupleArg,
			data: withBigInt{A: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Encode(test.arg, test.data)
			require.Error(t, err)
		})
	}
}

func TestEncodeBigInt(t *testing.T) {
	uint256Type, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint256",
		Type:    uint256Type,
		Indexed: false,
	}

	in := big.NewInt(8)

	packed, err := Encode(arg, in)
	require.NoError(t, err)

	expected := make([]byte, 32)
	expected[31] = 8

	require.Equal(t, expected, packed)

	x := big.NewInt(1231223123123333332)
	packed, err = Encode(arg, x)
	require.NoError(t, err)

	decoded, err := Decode[*big.Int](arg, packed)
	require.NoError(t, err)

	require.Equal(t, 0, x.Cmp(decoded))
}

func pow2(n uint) *big.Int {
	return new(big.Int).Lsh(big.NewInt(1), n)
}

// TestEncodeRefusesValuesOutOfType covers values go-ethereum packs without a
// range check into bytes Decode refuses.
func TestEncodeRefusesValuesOutOfType(t *testing.T) {
	var hash [32]byte
	hash[0], hash[31] = 0xaa, 1

	tests := []struct {
		name    string
		solType string
		value   any
	}{
		{"uint24", "uint24", pow2(24)},
		{"uint128", "uint128", pow2(128)},
		{"int24Above", "int24", pow2(23)},
		{"int128Below", "int128", new(big.Int).Sub(new(big.Int).Neg(pow2(127)), big.NewInt(1))},
		{"addressFromHash", "address", hash},
		{"functionFromHash", "function", hash},
		{"nested", "uint128[]", []*big.Int{big.NewInt(1), pow2(200)}},
		{"uint256Above", "uint256", pow2(256)},
		{"uint256Wrapping", "uint256", new(big.Int).Add(pow2(256), big.NewInt(5))},
		{"uint256Negative", "uint256", big.NewInt(-1)},
		{"int256Above", "int256", pow2(255)},
		{"int256Below", "int256", new(big.Int).Sub(new(big.Int).Neg(pow2(255)), big.NewInt(1))},
		{"arrayElem", "uint256[2]", [2]*big.Int{big.NewInt(1), pow2(256)}},
		{"pointer", "int256[]", []**big.Int{new(pow2(255))}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Encode(abi.Argument{Type: newType(t, tc.solType)}, tc.value)
			require.ErrorIs(t, err, ErrNonCanonical)
		})
	}
}

// TestEncodeAllowsValuesAtTypeBounds checks the output check rejects nothing
// in range, and Decode reads it back.
func TestEncodeAllowsValuesAtTypeBounds(t *testing.T) {
	var high [20]byte
	for i := range high {
		high[i] = 0xff
	}

	tests := []struct {
		name    string
		solType string
		value   any
	}{
		{"uint24Max", "uint24", new(big.Int).Sub(pow2(24), big.NewInt(1))},
		{"int24Min", "int24", new(big.Int).Neg(pow2(23))},
		{"int24Max", "int24", new(big.Int).Sub(pow2(23), big.NewInt(1))},
		{"uint256Max", "uint256", new(big.Int).Sub(pow2(256), big.NewInt(1))},
		{"int256Min", "int256", new(big.Int).Neg(pow2(255))},
		{"int256Max", "int256", new(big.Int).Sub(pow2(255), big.NewInt(1))},
		{"uint256Zero", "uint256", big.NewInt(0)},
		{"addressMax", "address", high},
		{"addressShort", "address", [3]byte{1, 2, 3}},
		{"addressWord", "address", [32]byte{31: 1}},
		{"functionWord", "function", [32]byte{1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			arg := abi.Argument{Type: newType(t, tc.solType)}
			packed, err := Encode(arg, tc.value)
			require.NoError(t, err)

			_, err = Decode[any](arg, packed)
			require.NoError(t, err)
		})
	}
}

// TestEncodeRefusesLongAddress covers an address or function from more than 32
// bytes, which go-ethereum writes whole. Here the second word reads as the
// string's offset, so the output still parses, with another string.
func TestEncodeRefusesLongAddress(t *testing.T) {
	var long [64]byte
	long[63] = 0x40

	for _, solType := range []string{"address", "function"} {
		arg := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: solType}, {Name: "s", Type: "string"}})}
		_, err := Encode(arg, struct {
			A [64]byte
			S string
		}{long, "0123456789abcdef0123456789abcdef"})
		require.ErrorIs(t, err, ErrNonCanonical, solType)
	}
}

type cyclic *cyclic

// TestEncodeRefusesRecursivePointer covers a value of a recursive pointer type,
// which go-ethereum follows until a fatal stack overflow. A regression loops, so
// Encode runs under a timeout.
func TestEncodeRefusesRecursivePointer(t *testing.T) {
	var p cyclic
	p = &p

	tests := []struct {
		name string
		arg  abi.Argument
		data any
	}{
		{"direct", abi.Argument{Type: newType(t, "uint256")}, p},
		{"field", abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint256"}})}, struct{ A cyclic }{p}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			done := make(chan error, 1)
			go func() {
				_, err := Encode(tc.arg, tc.data)
				done <- err
			}()

			select {
			case err := <-done:
				require.ErrorContains(t, err, "recursive pointer type")
			case <-time.After(5 * time.Second):
				t.Fatal("Encode loops on a recursive pointer type")
			}
		})
	}
}

// TestEncodeRefusesUnsupportedTypes covers types Decode refuses whatever the
// data.
func TestEncodeRefusesUnsupportedTypes(t *testing.T) {
	_, err := Encode(abi.Argument{Type: newType(t, "uint256[0]")}, [0]*big.Int{})
	require.ErrorIs(t, err, errUnsupportedType)

	_, err = Encode(abi.Argument{Type: newType(t, "uint256[0][]")}, [][0]*big.Int{})
	require.ErrorIs(t, err, errUnsupportedType)
}

// TestEncodeRangeFollowsFieldNames checks each value against the component
// Encode writes it to, which it finds by name, not position.
func TestEncodeRangeFollowsFieldNames(t *testing.T) {
	arg := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
		{Name: "a", Type: "int256"},
		{Name: "b", Type: "uint256"},
	})}
	uintMax := new(big.Int).Sub(pow2(256), big.NewInt(1))
	minusOne := big.NewInt(-1)

	type swapped struct{ B, A *big.Int }
	_, err := Encode(arg, swapped{B: uintMax, A: minusOne})
	require.NoError(t, err)

	type tagged struct {
		X *big.Int `abi:"b"`
		Y *big.Int `abi:"a"`
	}
	_, err = Encode(arg, tagged{X: uintMax, Y: minusOne})
	require.NoError(t, err)

	_, err = Encode(arg, swapped{B: minusOne, A: minusOne})
	require.ErrorIs(t, err, ErrNonCanonical)
}
