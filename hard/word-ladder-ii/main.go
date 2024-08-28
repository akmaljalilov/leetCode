package main

import (
	"fmt"
	"slices"
)

type Kucha struct {
	Name      string
	endWord   string
	length    int
	Kuchahosh []*Kucha
	Parent    []*Kucha
}

func (receiver Kucha) FindPr() [][]string {
	var parentStr = make([][]string, 0)
	var qadam = []string{receiver.Name}
	var quchaho = make(map[string][]string)
	receiver.findPr(qadam, &parentStr, &quchaho)
	return parentStr
}
func (receiver Kucha) findPr(qadam []string, parentStr *[][]string, kuchaho *map[string][]string) {
	if len(receiver.Parent) == 0 {
		return
	}
	minParents := []*Kucha{receiver.Parent[0]}
	for i := 1; i < len(receiver.Parent); i++ {
		if minParents[0].length > receiver.Parent[i].length {
			minParents = []*Kucha{receiver.Parent[i]}
		} else if receiver.Parent[i].length == minParents[0].length {
			minParents = append(minParents, receiver.Parent[i])
		}
	}
	for _, parent := range minParents {
		idx := slices.Index(qadam, parent.Name)
		if idx == -1 {
			qadam = append(qadam, parent.Name)
			if parentStr != nil && len(*parentStr) > 0 && len((*parentStr)[len(*parentStr)-1]) < len(qadam) {
				qadam = qadam[:len(qadam)-1]
				break
			}
		}
		if parentStr != nil && len(*parentStr) > 0 && len((*parentStr)[len(*parentStr)-1]) < len(qadam) {
			qadam = qadam[:len(qadam)-1]
			break
		}
		if idx > -1 {
			continue
		}
		if len(parent.Parent) == 0 && parentStr != nil {
			if len(*parentStr) == 0 {
				clone := make([]string, len(qadam))
				copy(clone, qadam)
				slices.Reverse(clone)
				*parentStr = append(*parentStr, clone)
			} else if len((*parentStr)[len(*parentStr)-1]) > len(qadam) {
				clone := make([]string, len(qadam))
				copy(clone, qadam)
				slices.Reverse(clone)
				*parentStr = [][]string{clone}
			} else if len((*parentStr)[len(*parentStr)-1]) == len(qadam) {
				clone := make([]string, len(qadam))
				copy(clone, qadam)
				slices.Reverse(clone)
				*parentStr = append(*parentStr, clone)
			} else {
				qadam = qadam[:len(qadam)-1]
				break
			}

		}
		parent.findPr(qadam, parentStr, kuchaho)
		qadam = qadam[:len(qadam)-1]
	}
}

var find1 func(kucha *Kucha)
var address = make(map[string]*Kucha)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	address = make(map[string]*Kucha)
	words := append(wordList, beginWord)
	yakumKucha := Kucha{Name: beginWord, Kuchahosh: make([]*Kucha, 0), Parent: make([]*Kucha, 0), endWord: endWord}
	address[beginWord] = &yakumKucha
	find1 = func(kucha *Kucha) {
		if kucha.Name == endWord {
			return
		}
		for i := 0; i < len(words); i++ {
			if kucha.Name == words[i] || beginWord == words[i] {
				continue
			}
			if contains(kucha.Name, words[i]) {
				if address[words[i]] != nil {
					address[words[i]].Parent = append(address[words[i]].Parent, kucha)
				} else {
					address[words[i]] = &Kucha{
						Name:      words[i],
						Kuchahosh: make([]*Kucha, 0),
						Parent:    []*Kucha{kucha}, endWord: endWord,
					}
					find1(address[words[i]])
				}
				kucha.Kuchahosh = append(kucha.Kuchahosh, address[words[i]])
			}
		}
	}
	find1(&yakumKucha)
	if address[endWord] == nil {
		return make([][]string, 0)
	}
	queues := []*Kucha{address[beginWord]}
	visit := make(map[string]bool)
	for len(queues) > 0 {
		queue := queues[0]
		for _, kucha := range queue.Kuchahosh {
			if !visit[kucha.Name] {
				kucha.length = queue.length + 1
				visit[kucha.Name] = true
				queues = append(queues, kucha)
			}
		}
		queues = queues[1:]
	}

	return address[endWord].FindPr()
}
func contains(start, end string) bool {
	if len(start) != len(end) {
		return false
	}
	j := 0
	for i, _ := range start {
		if start[i] == end[i] {
			j++
		}
	}
	return len(start)-1 == j
}

// Description: https://leetcode.com/problems/word-ladder-ii/description/
func main() {
	//var words = []string{"a", "b", "c"}
	//var startWord = "a"
	//var endWord = "c"

	//var words = []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim"}
	//var startWord = "cet"
	//var endWord = "ism"
	//
	var words = []string{"aaaaa", "waaaa", "wbaaa", "xaaaa", "xbaaa", "bbaaa", "bbwaa", "bbwba", "bbxaa", "bbxba", "bbbba", "wbbba", "wbbbb", "xbbba", "xbbbb", "cbbbb", "cwbbb", "cwcbb", "cxbbb", "cxcbb", "cccbb", "cccwb", "cccwc", "cccxb", "cccxc", "ccccc", "wcccc", "wdccc", "xcccc", "xdccc", "ddccc", "ddwcc", "ddwdc", "ddxcc", "ddxdc", "ddddc", "wdddc", "wdddd", "xdddc", "xdddd", "edddd", "ewddd", "ewedd", "exddd", "exedd", "eeedd", "eeewd", "eeewe", "eeexd", "eeexe", "eeeee", "weeee", "wfeee", "xeeee", "xfeee", "ffeee", "ffwee", "ffwfe", "ffxee", "ffxfe", "ffffe", "wfffe", "wffff", "xfffe", "xffff", "gffff", "gwfff", "gwgff", "gxfff", "gxgff", "gggff", "gggwf", "gggwg", "gggxf", "gggxg", "ggggg", "wgggg", "whggg", "xgggg", "xhggg", "hhggg", "hhwgg", "hhwhg", "hhxgg", "hhxhg", "hhhhg", "whhhg", "whhhh", "xhhhg", "xhhhh", "ihhhh", "iwhhh", "iwihh", "ixhhh", "ixihh", "iiihh", "iiiwh", "iiiwi", "iiixh", "iiixi", "iiiii", "wiiii", "wjiii", "xiiii", "xjiii", "jjiii", "jjwii", "jjwji", "jjxii", "jjxji", "jjjji", "wjjji", "wjjjj", "xjjji", "xjjjj", "kjjjj", "kwjjj", "kwkjj", "kxjjj", "kxkjj", "kkkjj", "kkkwj", "kkkwk", "kkkxj", "kkkxk", "kkkkk", "wkkkk", "wlkkk", "xkkkk", "xlkkk", "llkkk", "llwkk", "llwlk", "llxkk", "llxlk", "llllk", "wlllk", "wllll", "xlllk", "xllll", "mllll", "mwlll", "mwmll", "mxlll", "mxmll", "mmmll", "mmmwl", "mmmwm", "mmmxl", "mmmxm", "mmmmm", "wmmmm", "wnmmm", "xmmmm", "xnmmm", "nnmmm", "nnwmm", "nnwnm", "nnxmm", "nnxnm", "nnnnm", "wnnnm", "wnnnn", "xnnnm", "xnnnn", "onnnn", "ownnn", "owonn", "oxnnn", "oxonn", "ooonn", "ooown", "ooowo", "oooxn", "oooxo", "ooooo", "woooo", "wpooo", "xoooo", "xpooo", "ppooo", "ppwoo", "ppwpo", "ppxoo", "ppxpo", "ppppo", "wpppo", "wpppp", "xpppo", "xpppp", "qpppp", "qwppp", "qwqpp", "qxppp", "qxqpp", "qqqpp", "qqqwp", "qqqwq", "qqqxp", "qqqxq", "qqqqq", "wqqqq", "wrqqq", "xqqqq", "xrqqq", "rrqqq", "rrwqq", "rrwrq", "rrxqq", "rrxrq", "rrrrq", "wrrrq", "wrrrr", "xrrrq", "xrrrr", "srrrr", "swrrr", "swsrr", "sxrrr", "sxsrr", "sssrr", "ssswr", "sssws", "sssxr", "sssxs", "sssss", "wssss", "wtsss", "xssss", "xtsss", "ttsss", "ttwss", "ttwts", "ttxss", "ttxts", "tttts", "wttts", "wtttt", "xttts", "xtttt", "utttt", "uwttt", "uwutt", "uxttt", "uxutt", "uuutt", "uuuwt", "uuuwu", "uuuxt", "uuuxu", "uuuuu", "zzzzz", "zzzzy", "zzzyy", "zzyyy", "zzyyx", "zzyxx", "zzxxx", "zzxxw", "zzxww", "zzwww", "zzwwv", "zzwvv", "zzvvv", "zzvvu", "zzvuu", "zzuuu", "zzuut", "zzutt", "zzttt", "zztts", "zztss", "zzsss", "zzssr", "zzsrr", "zzrrr", "zzrrq", "zzrqq", "zzqqq", "zzqqp", "zzqpp", "zzppp", "zzppo", "zzpoo", "zzooo", "zzoon", "zzonn", "zznnn", "zznnm", "zznmm", "zzmmm", "zzmml", "zzmll", "zzlll", "zzllk", "zzlkk", "zzkkk", "zzkkj", "zzkjj", "zzjjj", "zzjji", "zzjii", "zziii", "zziih", "zzihh", "zzhhh", "zzhhg", "zzhgg", "zzggg", "zzggf", "zzgff", "zzfff", "zzffe", "zzfee", "zzeee", "zzeed", "zzedd", "zzddd", "zzddc", "zzdcc", "zzccc", "zzccz", "azccz", "aaccz", "aaacz", "aaaaz", "uuuzu", "uuzzu", "uzzzu", "zzzzu"}
	var startWord = "aaaaa"
	var endWord = "uuuuu"

	//var words = []string{"hot", "dot", "dog", "lot", "log"}
	//var startWord = "hit"
	//var endWord = "cog"
	//var words = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	//var startWord = "hit"
	//var endWord = "cog"
	//var words = []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	//var startWord = "qa"
	//var beginWord = "sq"
	fmt.Println(findLadders(startWord, endWord, words))
}
