package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeNotSigning(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		output string
	}{
		{
			name: "offerCreate",
			json: `{
  "Account": "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
  "Expiration": 595640108,
  "Fee": "10",
  "Flags": 524288,
  "OfferSequence": 1752791,
  "Sequence": 1752792,
  "SigningPubKey": "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
  "TakerGets": "15000000000",
  "TakerPays": {
    "currency": "USD",
    "issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
    "value": "7072.8"
  },
  "TransactionType": "OfferCreate",
  "TxnSignature": "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
  "hash": "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"
}`,
			output: "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
		},
		{
			name: "payment",
			json: `{
      "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
      "Amount": {
        "currency": "USD",
        "issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
        "value": "1"
      },
      "Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
      "Fee": "10",
      "Flags": 2147483648,
      "Sequence": 1,
      "SigningPubKey": "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
      "TransactionType": "Payment",
      "TxnSignature": "3045022100F8A650C1D58325FE8D74634C1DC0802BB2271EB84773793EF34085CFC7E32B1302206ECE43AFE94B7F9F0359D53E6B195C2D526DFDFBBBF328D6FE3A598F1D51DEBA"
    }`,
			output: "1200002280000000240000000161D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA968400000000000000A732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB74473045022100F8A650C1D58325FE8D74634C1DC0802BB2271EB84773793EF34085CFC7E32B1302206ECE43AFE94B7F9F0359D53E6B195C2D526DFDFBBBF328D6FE3A598F1D51DEBA81144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
		},
		{
			name: "memo",
			json: `{
      "TakerPays": "223174650",
      "Account": "rPk2dXr27rMw9G5Ej9ad2Tt7RJzGy8ycBp",
      "TransactionType": "OfferCreate",
      "Memos": [{
        "Memo": {
          "MemoType": "584D4D2076616C7565",
          "MemoData": "322E3230393635"
                }}],
                "Fee": "15",
                "OfferSequence": 1002,
                "TakerGets": {
                    "currency": "XMM",
                    "value": "100",
                    "issuer": "rExAPEZvbkZqYPuNcZ7XEBLENEshsWDQc8"
                },
                "Flags": 524288,
                "Sequence": 1003,
                "LastLedgerSequence": 6220135
            }`,
			output: "120007220008000024000003EB2019000003EA201B005EE96764400000000D4D5FFA65D5038D7EA4C68000000000000000000000000000584D4D0000000000A426093A78AA86EB2B878E5C2E33FEC224A0184968400000000000000F8114F990B9E746546554A7B50A5E013BCB57095C6BB8F9EA7C09584D4D2076616C75657D07322E3230393635E1F1",
		},
		{
			name: "path",
			json: `{
                "Account": "rHXUjUtk5eiPFYpg27izxHeZ1t4x835Ecn",
                "Destination": "r45dBj4S3VvMMYXxr9vHX4Z4Ma6ifPMCkK",
                "TransactionType": "Payment",
                "Amount": {
                    "currency": "CNY",
                    "value": "5000",
                    "issuer": "r45dBj4S3VvMMYXxr9vHX4Z4Ma6ifPMCkK"
                },
                "Fee": "12",
                "SendMax": {
                    "currency": "CNY",
                    "value": "5050",
                    "issuer": "rHXUjUtk5eiPFYpg27izxHeZ1t4x835Ecn"
                },
                "Flags": 0,
                "Sequence": 6,
                "Paths": [[{
                    "account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"
                }]],
                "DestinationTag": 736049272
            }`,
			output: "120000220000000024000000062E2BDF387861D551C37937E08000000000000000000000000000434E590000000000EE39E6D05CFD6A90DAB700A1D70149ECEE29DFEC68400000000000000C69D551F0F2C01DA000000000000000000000000000434E590000000000B53847FA45E828BF9A52E38F7FB39E363493CE8B8114B53847FA45E828BF9A52E38F7FB39E363493CE8B8314EE39E6D05CFD6A90DAB700A1D70149ECEE29DFEC01120141C8BE2C0A6AA17471B9F6D0AF92AAB1C94D5A2500",
		},
		{
			name: "destinationTag",
			json: `{
                "Account": "r3ZDv3hLmTKwkgAqcXtX2yaMfnhRD3Grjc",
                "Destination": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
                "TransactionType": "Payment",
                "Amount": {
                    "currency": "BTC",
                    "value": "0.04",
                    "issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
                },
                "Fee": "106",
                "SendMax": "3267350000",
                "Flags": 0,
                "Sequence": 10,
                "Paths": [[{
                    "currency": "BTC",
                    "issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
                }]],
                "InvoiceID": "342B8D16BEE494D169034AFF0908FDE35874A38E548D4CEC8DFC5C49E9A33B76",
                "DestinationTag": 1403334172
            }`,
			output: "1200002200000000240000000A2E53A52E1C5011342B8D16BEE494D169034AFF0908FDE35874A38E548D4CEC8DFC5C49E9A33B7661D40E35FA931A00000000000000000000000000004254430000000000DD39C650A96EDA48334E70CC4A85B8B2E8502CD368400000000000006A6940000000C2BFCDF0811452E0F910686FB449A23BC78C3D4CE564C988C6C08314DD39C650A96EDA48334E70CC4A85B8B2E8502CD30112300000000000000000000000004254430000000000DD39C650A96EDA48334E70CC4A85B8B2E8502CD300",
		}, {
			name: "payment channel",
			json: `{
  "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
  "TransactionType": "PaymentChannelFund",
  "Channel": "C1AE6DDDEEC05CF2978C0BAD6FE302948E9533691DC749DCDD3B9E5992CA6198",
  "Amount": "200000",
  "Expiration": 543171558
}`,
			output: "12000E2A206023E65016C1AE6DDDEEC05CF2978C0BAD6FE302948E9533691DC749DCDD3B9E5992CA6198614000000000030D4081144B4E9C06F24296074F7BC48F92A97916C6DC5EA9",
		},
	}

	for _, test := range tests {
		blob, err := hex.DecodeString(test.output)
		require.NoError(t, err, test.name)

		var parsedJSON any

		err = json.Unmarshal([]byte(test.json), &parsedJSON)
		require.NoError(t, err, test.name)

		bytes, err := Encode(parsedJSON, false)
		require.NoError(t, err, test.name)

		require.Equal(t, blob, bytes, test.name)
	}
}

func TestLengthEncodeDecode(t *testing.T) {
	inputs := []int{0, 1, 100, 192, 193, 240, 241, 12480, 12481, 918744}

	for _, input := range inputs {
		prefix, err := lengthEncode(input)
		require.NoError(t, err)

		b := bytes.NewBuffer(prefix)
		length, err := lengthDecode(b)
		require.NoError(t, err)

		require.Equal(t, input, length)

		require.Equal(t, b.Len(), 0)
	}
}

func TestLengthEncodeFail(t *testing.T) {
	inputs := []int{-1, 918745}

	for _, input := range inputs {
		_, err := lengthEncode(input)
		require.Error(t, err)
	}
}

func TestLengthDecodeEncode(t *testing.T) {
	inputs := [][]byte{{0}, {1}, {192}, {193, 0}, {193, 1}, {193, 255}, {194, 0}, {241, 0, 0}, {254, 212, 23}}

	for _, input := range inputs {
		b := bytes.NewBuffer(input)
		length, err := lengthDecode(b)
		require.NoError(t, err)

		require.Equal(t, b.Len(), 0, input)

		prefix, err := lengthEncode(length)
		require.NoError(t, err)

		require.Equal(t, input, prefix)
	}
}

func TestLengthDecodeFail(t *testing.T) {
	inputs := [][]byte{{255}, {193}, {241, 0}, {254, 212, 24}}

	for _, input := range inputs {
		b := bytes.NewBuffer(input)
		_, err := lengthDecode(b)
		require.Error(t, err)
	}
}
