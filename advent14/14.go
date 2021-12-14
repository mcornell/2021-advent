package advent14

import (
	"fmt"
	"strings"
)

type Rule struct {
	insert  string
	between string
}

func NewRule(ruleString string) Rule {
	rule_info := strings.Split(ruleString, " -> ")
	r := Rule{
		insert:  rule_info[1],
		between: rule_info[0],
	}
	return r
}

func (rule *Rule) String() string {
	return fmt.Sprintf("%s->%s", rule.between, rule.insert)
}

type Polymer struct {
	p string
}

func (p *Polymer) GeneratePairs() []string {
	pairs := make([]string, len(p.p)-1)
	for idx := 0; idx < len(p.p)-1; idx++ {
		pairs[idx] = (p.p)[idx : idx+2]
	}
	return pairs
}

func (p *Polymer) ExecuteRules(rules []Rule) {
	new_polyString := p.p
	insert_idx := 1
	for _, pair := range p.GeneratePairs() {
		for _, rule := range rules {
			if pair == rule.between {
				new_polyString = new_polyString[:insert_idx] + rule.insert + new_polyString[insert_idx:]
				insert_idx += 1
				break
			}
		}
		insert_idx += 1

	}
	p.p = new_polyString
}

func (p *Polymer) String() string {
	return p.p
}

func NewPolymer(polyString string) *Polymer {
	p := Polymer{
		p: polyString,
	}
	return &p
}

func ParseInput(data []string) (Polymer, []Rule) {
	p := NewPolymer(data[0])
	rules := []Rule{}
	for _, ruleString := range data[2:] {
		rules = append(rules, NewRule(string(ruleString)))
	}

	return *p, rules
}

func FindMostCommonElement(data []string) (int, int) {
	p, rules := ParseInput(data)
	for i := 0; i < 10; i++ {
		p.ExecuteRules(rules)
	}
	count := make(map[rune]int)
	min := 99999999999
	max := 0
	for _, element := range p.p {
		val := count[element]

		if val == 0 {
			num := strings.Count(p.p, string(element))
			count[element] = num
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	}
	return max, min
}
