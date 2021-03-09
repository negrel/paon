package flows

//// AddCustomRenderer adds the given draw.Painter constructor to the render map
//// and return its ID. This ID should be used for the Display property.
//func AddCustomAlgorithm(constructor func() flow) int {
//	return fMap.add(constructor)
//}
//
//// GetFor returns a new draw.Painter for the given draw.Object.
//func GetFor(style styles.Style) flow {
//	prop := style.Get(property.IDFlow)
//	if prop == nil {
//		return fMap.algorithms[Unset]()
//	}
//	flowID := prop.(property.Int)
//
//	makeFlow, ok := fMap.algorithms[flowID.Value]
//	if !ok {
//		return fMap.algorithms[Unset]()
//	}
//
//	return makeFlow()
//}
//
//// IsValidFlowID returns true if the given rendererID is valid.
//func IsValidFlowID(algorithmID int) bool {
//	return algorithmID > 0 && algorithmID < fMap.len()
//}
//
//type flowMap struct {
//	algorithms map[int]func() flow
//}
//
//const (
//	Unset = iota
//	Hidden
//	BlockFlow
//	Inline
//	Flex
//)
//
//var fMap = flowMap{
//	algorithms: map[int]func() flow{
//		Unset:  makeUnset,
//		Hidden: makeHidden,
//		Block:  makeBlock,
//		Inline: makeInline,
//		Flex:   makeFlex,
//	},
//}
//
//func (fm flowMap) add(renderer func() flow) int {
//	index := fm.len()
//	fm.algorithms[index] = renderer
//
//	return index
//}
//
//func (fm flowMap) len() int {
//	return len(fm.algorithms)
//}
