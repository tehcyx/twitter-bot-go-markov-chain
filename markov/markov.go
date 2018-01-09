package markov

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// SubDictionary of follow up words and the corresponding factor.
type SubDictionary map[string]int

// Dictionary of start words with follow ups and the corresponding factor.
type Dictionary map[string]SubDictionary

// Train from string of words
func Train(text string, factor int) Dictionary {
	dict := make(Dictionary)

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		words[i] = strings.ToLower(words[i])
		if _, prefixAvail := dict[words[i]]; !prefixAvail {
			dict[words[i]] = make(SubDictionary)
		}
		if _, suffixAvail := dict[words[i]][words[i+1]]; !suffixAvail {
			dict[words[i]][words[i+1]] = factor
		} else {
			dict[words[i]][words[i+1]] = dict[words[i]][words[i+1]] + dict[words[i]][words[i+1]]*factor
		}
	}
	return dict
}

// TrainFromFile takes a file path, reads the file and passes the string to Train
func TrainFromFile(path string, factor int) Dictionary {
	buf, _ := ioutil.ReadFile(path)
	return Train(string(buf), factor)
}

// TrainFromFolder takes a path, and passes every text file it finds to TrainFromFile
func TrainFromFolder(path string, factor int) Dictionary {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dir + string(os.PathSeparator) + path)
	if err != nil {
		log.Fatal(err)
	}

	dict := make(Dictionary)

	for _, file := range files {
		match, _ := regexp.MatchString(".*\\.txt$", file.Name())
		if match {
			fmt.Println(file.Name())
			fmt.Println(dir + string(os.PathSeparator) + path + string(os.PathSeparator) + file.Name())
			tmp := TrainFromFile(path+string(os.PathSeparator)+file.Name(), factor)
			dict = mergeDict(dict, tmp)
		}
	}

	return dict
}

// Generate takes a dictionary, a maximum length and a startword to generate a text based on the inputs
func Generate(dict Dictionary, maxLength int, startWord string) string {
	var word = ""
	if startWord == "" {
		word = pickRandom(dict.keys())
	} else {
		word = startWord
	}

	sentence := word
	i := 0
	for maxLength == 0 || i < maxLength-1 {
		i++

		tmp := word

		for k := range dict[word] {
			if _, ok := dict[word][k]; ok {
				word = pickRandom(dict[word].keys())
			}
		}
		if word == tmp || word == "" {
			return sentence
		}
		tmp = word
		sentence = sentence + " " + word
	}
	return sentence
}

// adjustFactors: function(dict, maxLength = 2, f = fitnessFunc) {
// 	var extract = this.generate(dict, maxLength).split(' ');

// 	var pairs = [];
// 	for (var i = 0; i < extract.length; i++) {
// 		if (typeof extract[i + 1] == 'undefined') {
// 			continue;
// 		}
// 		pairs[i] = [extract[i], extract[i + 1]];
// 	}

// 	for (var p in pairs) {
// 		var fact = (f(dict, pairs[p]) - 0.5) * 2;

// 		dict = mergeDict(this.train(pairs[p][0] + " " + pairs[p][1], fact), dict);
// 	}
// 	return dict;
// },

// bulkAdjustFactors: function(dict, iterations = 1, f = [ undefined ]) {
// 	if (typeof f == 'undefined' || typeof f[0] == 'undefined') {
// 		return dict;
// 	}
// 	for (var i = 0; i < iterations; i++) {
// 		for (var j = 0; j < f.length; j++) {
// 			dict = this.adjustFactors(dict, 10, f[j]);
// 		}
// 	}

// 	return dict;
// }

// mergeDict, given two dictionaries merges them into one
func mergeDict(dict1, dict2 Dictionary) Dictionary {
	res := dict1
	for k := range dict2 {
		if _, ok := res[k]; ok {
			//do something here
			res[k] = dict2[k]
		} else {
			for sk := range dict2[k] {
				if _, ok := res[k][sk]; ok {
					res[k][sk] = dict2[k][sk]
				} else {
					res[k][sk] = res[k][sk] + dict2[k][sk]
				}
			}
		}
	}
	return res
}

func (m Dictionary) keys() []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m SubDictionary) keys() []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// pickRandom takes a dictionary and selects a random key
func pickRandom(keys []string) string {
	return keys[rand.Intn(len(keys))]
}

// Version with reflection (doesn't work just yet, would be more readable)
// // pickRandom takes a dictionary and selects a random key
// func pickRandom(dict map[string]interface{}) string { // see https://golang.org/pkg/reflect/
// 	keys := make([]string, len(dict))
// 	i := 0
// 	for k := range dict {
// 		keys[i] = k
// 		i++
// 	}
// 	return keys[rand.Intn(len(keys))]
// }
