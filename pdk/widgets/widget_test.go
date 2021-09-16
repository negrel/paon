package widgets

// type widgetTest struct {
// 	name string
// 	test func(t *testing.T)
// }

// func TestWidget(t *testing.T) {
// 	for _, methodTest := range generateWidgetsTests() {
// 		t.Run(methodTest.name, methodTest.test)
// 	}
// }

// func generateWidgetsTests() []widgetTest {
// 	tests := []widgetTest{
// 		{
// 			name: "ID",
// 			test: TestWidgetID,
// 		},
// 		{
// 			name: "LifeCycleStage",
// 			test: TestWidgetLifeCycleStage,
// 		},
// 		{
// 			name: "Node",
// 			test: TestWidgetNode,
// 		},
// 		{
// 			name: "Box",
// 			test: TestWidgetBox,
// 		},
// 		{
// 			name: "Theme",
// 			test: TestWidgetTheme,
// 		},
// 		{
// 			name: "Parent",
// 			test: TestWidgetParent,
// 		},
// 		// {
// 		// 	name: "Draw",
// 		// 	test: TestWidgetDraw,
// 		// },
// 		{
// 			name: "Siblings",
// 			test: TestWidgetSiblings,
// 		},
// 	}

// 	return tests
// }

// func TestWidgetID(t *testing.T) {
// 	widget := newWidget()
// 	other := newWidget()

// 	assert.NotEqual(t, widget, other)
// 	assert.NotEqual(t, widget.ID(), other.ID())

// 	assert.Equal(t, widget.ID(), widget.Node().ID())
// }

// func TestWidgetLifeCycleStage(t *testing.T) {
// 	for i := 0; i < int(_maxLifeCycle); i++ {
// 		lcs := LifeCycleStage(i)

// 		widget := newWidget(initialLCS(lcs))
// 		assert.Equal(t, lcs, widget.LifeCycleStage())
// 	}
// }

// func TestWidgetNode(t *testing.T) {
// 	type wrapper struct {
// 		*BaseWidget
// 	}

// 	var node tree.Node
// 	wWrapper := wrapper{}

// 	constructor := func(data interface{}) tree.Node {
// 		wWrapper.BaseWidget = data.(*BaseWidget)

// 		node = tree.NewNode(wWrapper)
// 		return node
// 	}
// 	widget := newWidget(NodeConstructor(constructor))

// 	assert.Equal(t, node.Unwrap(), wWrapper)
// 	assert.True(t, widget.Node().IsSame(node))
// }

// func TestWidgetBox(t *testing.T) {
// 	expectedBox := layout.NewBox(
// 		geometry.Rect(
// 			rand.Int(), rand.Int(), rand.Int(), rand.Int(),
// 		),
// 	)

// 	widget := newWidget(
// 		LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
// 			return expectedBox
// 		})),
// 	)

// 	constraint := layout.Constraint{
// 		MinSize: geometry.Rectangle{},
// 		MaxSize: expectedBox.ContentBox().
// 			GrowLeft(5).GrowRight(5).
// 			GrowTop(5).GrowBottom(5),
// 		ParentSize: geometry.Size{},
// 		RootSize:   geometry.Size{},
// 	}

// 	box := widget.Layout(constraint)

// 	assert.Equal(t, expectedBox, box)
// }

// func TestWidgetTheme(t *testing.T) {
// 	t.Run("Default", TestWidgetThemeDefault)
// }

// func TestWidgetThemeDefault(t *testing.T) {
// 	widget := newWidget()

// 	assert.NotNil(t, widget.Theme())
// }

// func TestWidgetParent(t *testing.T) {
// 	parent := newMountedLayout()
// 	widget := newWidget()

// 	err := parent.AppendChild(widget)
// 	assert.Nil(t, err)

// 	assert.Equal(t, parent, widget.Parent())
// }

// // func TestWidgetDraw(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	drawer := NewMockDrawer(ctrl)
// // 	context := NewMockContext(ctrl)

// // 	widget := newWidget(Drawer(drawer))
// // 	drawer.EXPECT().Draw(context).MinTimes(1)

// // 	widget.Draw(context)
// // }

// func TestWidgetSiblings(t *testing.T) {
// 	parent := newLayout()
// 	child := newWidget()
// 	prev := newWidget()
// 	next := newWidget()

// 	var err error
// 	err = parent.AppendChild(prev)
// 	assert.Nil(t, err)
// 	err = parent.AppendChild(child)
// 	assert.Nil(t, err)
// 	err = parent.AppendChild(next)
// 	assert.Nil(t, err)

// 	assert.Equal(t, parent, prev.Parent())
// 	assert.Equal(t, parent, child.Parent())
// 	assert.Equal(t, parent, next.Parent())

// 	assert.Nil(t, prev.PreviousSibling())
// 	assert.Equal(t, prev, child.PreviousSibling())
// 	assert.Equal(t, child, prev.NextSibling())
// 	assert.Equal(t, next, child.NextSibling())
// 	assert.Equal(t, child, next.PreviousSibling())
// 	assert.Nil(t, next.NextSibling())
// }
