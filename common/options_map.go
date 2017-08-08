package common

import (
	"image/color"
)

type MapOption func(*Map)


func Width(in int) MapOption {
	return func(m *Map) {
		if in < 1 {
			in = 1
		}
		m.Width = in
	}
}

func Height(in int) MapOption {
	return func(m *Map) {
		if in < 1 {
			in = 1
		}
		m.Height = in
	}
}

func TileWidth(in int) MapOption {
	return func(m *Map) {
		if in < 1 {
			in = 1
		}
		m.TileWidth = in
	}
}

func TileHeight(in int) MapOption {
	return func(m *Map) {
		if in < 1 {
			in = 1
		}
		m.TileHeight = in
	}
}

func Background(in *color.RGBA) MapOption {
	return func(m *Map) {
		m.BackgroundColor = in
	}
}

func RenderOrderRightDown() MapOption {
	return func(m *Map) {
		m.renderOrder = MapRenderOrderRightDown
	}
}

func RenderOrderRightUp() MapOption {
	return func(m *Map) {
		m.renderOrder = MapRenderOrderRightUp
	}
}

func RenderOrderLeftDown() MapOption {
	return func(m *Map) {
		m.renderOrder = MapRenderOrderLeftDown
	}
}

func RenderOrderLeftUp() MapOption {
	return func(m *Map) {
		m.renderOrder = MapRenderOrderLeftUp
	}
}

func OrientationOrthogonal() MapOption {
	return func(m *Map) {
		m.orientation = MapOrientationOrthogonal
	}
}

func OrientationIsometric() MapOption {
	return func(m *Map) {
		m.orientation = MapOrientationIsometric
	}
}

func OrientationStaggered(staggerAxis, staggerIndex string) MapOption {
	return func(m *Map) {
		m.orientation = MapOrientationStaggered

		// defaults
		m.staggerAxis = DefaultStaggerAxis
		m.staggerIndex = DefaultStaggerIndex

		if staggerAxis == MapStaggerAxisX || staggerAxis == MapStaggerAxisY {
			m.staggerAxis = staggerAxis
		}
		if staggerIndex == MapStaggerIndexEven || staggerIndex == MapStaggerIndexOdd {
			m.staggerIndex = staggerIndex
		}
	}
}
