// Code generated by "stringer -type=Language"; DO NOT EDIT

package gen

import "fmt"

const _Language_name = "AndroidObjCSwift"

var _Language_index = [...]uint8{0, 7, 11, 16}

func (i Language) String() string {
	if i < 0 || i >= Language(len(_Language_index)-1) {
		return fmt.Sprintf("Language(%d)", i)
	}
	return _Language_name[_Language_index[i]:_Language_index[i+1]]
}

var _LanguageNameToValue_map = map[string]Language{
	_Language_name[0:7]:   0,
	_Language_name[7:11]:  1,
	_Language_name[11:16]: 2,
}

func LanguageString(s string) (Language, error) {
	if val, ok := _LanguageNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Language values", s)
}
