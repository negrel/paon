package widgets

//go:generate mockgen -source widget.go -destination widget_mock_test.go -package widgets . Widget
//go:generate mockgen -destination render_renderable_mock_test.go -package widgets github.com/negrel/paon/render Renderable
