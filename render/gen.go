package render

//go:generate mockgen -destination draw_surface_mock_test.go -package render github.com/negrel/paon/draw Surface
//go:generate mockgen -source renderable.go -destination renderable_mock_test.go -package render . Renderable
