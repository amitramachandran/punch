package familytree

import (
	"regexp"
	"strings"
)

func LowerCase(in []string) (out []string) {

	for _, s := range in {
		o := strings.ToLower(s)
		out = append(out, o)
	}
	return
}

func JoinString(in []string) (out string) {
	return strings.Join(in, " ")
}

func remove(in []string, r string) []string {
	for i, v := range in {
		if v == r {
			return append(in[:i], in[i+1:]...)
		}
	}
	return in
}

func PreProcessName(in []string) (out string) {
	lc := LowerCase(in)
	removeOf := remove(lc, "of")
	return JoinString(removeOf)
}

// (from, to, relation string)
func PreProcessConnection(in []string) (from, relation, to string) {
	var res []string
	js := JoinString(in)
	pattern := regexp.MustCompile("as|of")
	split := pattern.Split(js, -1)
	for _, s := range split {
		if s != "" {
			res = append(res, strings.TrimSpace(strings.ToLower(s)))
		}
	}
	return res[0], res[1], res[2]
}
