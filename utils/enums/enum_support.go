package enums

import (
	"fmt"
	"utils"
)

type EnumHelper interface {
	ToString(enum any) string
	IsValid(it any) (val int, ok bool)
}

func Init(enumName string, geInvalidEnum any) *Builder {
	err := utils.StringVisibleAsciiOrNonAsciiUTF8(enumName)
	if err != nil {
		panic(fmt.Errorf("enum error 0: enumName '%v' was invalid: %w", enumName, err))
	}
	return &Builder{name: enumName, geInvalidEnum: geInvalidEnum}
}

type Builder struct {
	name          string
	geInvalidEnum any
	options       []any
	optionsText   []string
	built         bool
}

func (r *Builder) Add(option any, optionText string) *Builder {
	if r.built {
		panic(fmt.Errorf("enum error 1: option Add on '%v' after Build() called", r.name))
	}
	ndx := len(r.options)
	err := utils.StringVisibleAsciiOrNonAsciiUTF8(optionText)
	if err != nil {
		panic(fmt.Errorf("enum error 2: option Add[%d] (text) '%v' on '%v' -- %w", ndx, optionText, r.name, err))
	}
	next := ndx + 1
	_, err = validateEnumInt(option, next, next)
	if err != nil {
		panic(fmt.Errorf("enum error 3: option Add[%d] '%v' w/ '%v' on '%v' -- %w", ndx, option, optionText, r.name, err))
	}
	r.options = append(r.options, option)
	r.optionsText = append(r.optionsText, optionText)
	return r
}

func (r *Builder) Build() EnumHelper {
	if r.built {
		panic(fmt.Errorf("enum error 4: called Build() on '%v' after already built", r.name))
	}
	maxValidEnum := len(r.options)
	if maxValidEnum == 0 {
		panic(fmt.Errorf("error 5: no options Add(ed) on '%v'", r.name))
	}
	next := maxValidEnum + 1
	_, err := validateEnumInt(r.geInvalidEnum, next, next)
	if err != nil {
		panic(fmt.Errorf("enum error 6: 'geInvalid%v' %w options Add(ed) on '%v'", r.name, err, r.name))
	}
	return &enumHelper{name: r.name, maxValidEnum: maxValidEnum, options: r.options, optionsText: r.optionsText}
}

type enumHelper struct {
	name         string
	maxValidEnum int
	options      []any
	optionsText  []string
}

func (r *enumHelper) ToString(enumIt any) string {
	val, err := r.validateEnumIt(enumIt)
	if err != nil {
		panic(fmt.Errorf("enum '%v' -- %w", r.name, err))
	}
	return r.optionsText[val-1]
}

func (r *enumHelper) IsValid(enumIt any) (val int, ok bool) {
	var err error
	val, err = r.validateEnumIt(enumIt)
	ok = err == nil
	return
}

func (r *enumHelper) validateEnumIt(enumIt any) (int, error) {
	return validateEnumInt(enumIt, 1, r.maxValidEnum)
}

func validateEnumInt(enumIt any, expectedMin, expectedMax int) (val int, err error) {
	val, err = utils.ToInt(enumIt)
	if err == nil {
		if (val < expectedMin) || (expectedMax < val) {
			if expectedMin == expectedMax {
				err = fmt.Errorf("expected %d, but was %d", expectedMin, val)
			} else {
				err = fmt.Errorf("expected %d - %d, but was: %d", expectedMin, expectedMax, val)
			}
		}
	}
	return
}
