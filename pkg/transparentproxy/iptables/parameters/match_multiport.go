package parameters

func Multiport() *MatchParameter {
	return &MatchParameter{
		name:       "multiport",
		parameters: []ParameterBuilder{},
	}
}
