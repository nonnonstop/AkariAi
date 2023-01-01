package discord

import "strings"

var normalizeRemoveCharacters = []rune{
	' ', '　', '"', '#', '\\', '$', '%', '&', '-', '^', '\'', '@', ';', ':',
	'.', '/', '=', '~', '|', '`', '+', '*', '_', '＂', '＃', '＄', '％', '＆',
	'＇', '－', '＾', '＼', '＠', '；', '：', '，', '．', '￥', '／', '＝', '～',
	'｜', '｀', '＋', '＊', '＿', '・', '\r', '\n', '\t',
}

var normalizeConvertCharactors = map[rune]rune{
	'ｦ': 'を', 'ｧ': 'ぁ', 'ｨ': 'ぃ', 'ｩ': 'ぅ', 'ｪ': 'ぇ', 'ｫ': 'ぉ', 'ｬ': 'ゃ',
	'ｭ': 'ゅ', 'ｮ': 'ょ', 'ｯ': 'っ', 'ｱ': 'あ', 'ｲ': 'い', 'ｳ': 'う', 'ｴ': 'え',
	'ｵ': 'お', 'ｶ': 'か', 'ｷ': 'き', 'ｸ': 'く', 'ｹ': 'け', 'ｺ': 'こ', 'ｻ': 'さ',
	'ｼ': 'し', 'ｽ': 'す', 'ｾ': 'せ', 'ｿ': 'そ', 'ﾀ': 'た', 'ﾁ': 'ち', 'ﾂ': 'つ',
	'ﾃ': 'て', 'ﾄ': 'と', 'ﾅ': 'な', 'ﾆ': 'に', 'ﾇ': 'ぬ', 'ﾈ': 'ね', 'ﾉ': 'の',
	'ﾊ': 'は', 'ﾋ': 'ひ', 'ﾌ': 'ふ', 'ﾍ': 'へ', 'ﾎ': 'ほ', 'ﾏ': 'ま', 'ﾐ': 'み',
	'ﾑ': 'む', 'ﾒ': 'め', 'ﾓ': 'も', 'ﾔ': 'や', 'ﾕ': 'ゆ', 'ﾖ': 'よ', 'ﾗ': 'ら',
	'ﾘ': 'り', 'ﾙ': 'る', 'ﾚ': 'れ', 'ﾛ': 'ろ', 'ﾜ': 'わ', 'ﾝ': 'ん',
}

func normalizeString(input string) string {
	var sb strings.Builder
InputLoop:
	for _, c := range input {
		// Remove some characters
		for _, tc := range normalizeRemoveCharacters {
			if c == tc {
				continue InputLoop
			}
		}
		switch {
		case c == '？':
			c = '?'
		case '\uFF21' <= c && c <= '\uFF3A': // full uppercase
			c -= '\uFF21' - '\u0041'
		case '\uFF41' <= c && c <= '\uFF5A': // full lowercase
			c -= '\uFF41' - '\u0041'
		case '\u0061' <= c && c <= '\u007A': // half lowercase
			c -= '\u0061' - '\u0041'
		case '\u30A1' <= c && c <= '\u30F6': // full katakana
			c -= '\u30A1' - '\u3041'
		case '\uFF66' <= c && c <= '\uFF9D': // half katakana
			c = normalizeConvertCharactors[c]
		}
		sb.WriteRune(c)
	}
	return sb.String()
}
