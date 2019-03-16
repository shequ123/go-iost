package vm

import "errors"

var errTxMessage = map[string]string{
	"HLxkjG6iRoeTviDDBG3jPT9kLobYR7zXnnvTpgu4UBPu": `error in data: invalid character 'h' looking for beginning of value, [huobipool]`,
	"EzjkEDtidB5YLUWTRtQBjgLQWV3TpGse6no3XMW4vupM": `error in data: invalid character 'h' looking for beginning of value, [huobipool]`,
	"7cBUeWKhu4W1LSKb9CjeVpyKNntfeeSb8oWN556dSLtp": `error in data: invalid character ']' after object key:value pair, [\"ycy\",\"lispccc\",\"1000\",{\"decimal\":1,\"full_name\":\"token for Yang Chaoyue\"]`,
	"6YW57thSdHnkW1CCsxxndnJjwYJ2GQS7p4ejrSWsq8B7": `error in data: invalid character 'i' looking for beginning of value, iostqed`,
	"5YQ1sqgkuJ4JnsZZ7TKFNtoxuD6fnNdP6szssn4DygV3": `error in data: invalid character 'i' looking for beginning of value, [iostqed]`,
	"4AicRFXMt2W4Z5RLB5Zb2A4ycUv7SfH2iHi5dbRAUBLf": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"A5ZDP5TwHRbQKDQXt9XyE4DbPtKED5ppUidu14kzuZ9o": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"BMCfLvQM5sJ3h1ZiwNj9MSSGScnNRpa3dhPQjN1CgVEj": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"9dv6kjTzoX9yu933p1p7SfLHWmDWjUycA2qGASz5UEv7": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"12ctz18wJVigGfLp9yGMnYkj6fQ2nx9GdAEPnghwEQBj": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"3g3YLu7HVaVcdM2pCWyP9TaMTPZSRgzFLLP89tmvUBCu": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"59cFgDDijMHxCS7F5rUjVYPDjL4UzEuLgwGJdx1NAYjh": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"8Uczn4EhTFDRsxWpYfoB942UAAgnRApube3mx92TkjNA": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"GJqwEkB2zMqsHhhhbMkfc15XD2kgCZubnybupYJdgymi": `error in data: invalid character 'h' looking for beginning of value, helloiost`,
	"FE4zHcs4aJtkYEXXdcS9VxhWf6rcgp1khjAjnwntHYmr": `error in data: invalid character 'h' looking for beginning of value, [helloiost]`,
	"HoRQd81KZFnEQ2NvN7AjBWTUYQHw4c1NjWcNmG3W6ASE": `error in data: invalid character 'h' looking for beginning of value, [helloiost]`,
}

func getTxError(txhash string, err error) error {
	if msg, ok := errTxMessage[txhash]; ok {
		return errors.New(msg)
	}
	return err
}
