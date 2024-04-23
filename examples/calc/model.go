package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ModelView interface {
	SetText(string)
}

type Model struct {
	left     string
	operator string
	right    string
	view     ModelView
}

func NewModel(mv ModelView) *Model {
	m := &Model{
		view: mv,
	}
	m.Clear()

	return m
}

func (m *Model) Clear() {
	m.left = ""
	m.operator = ""
	m.right = ""
	m.updateView()
}

func (m *Model) currentNumber() *string {
	if m.operator == "" {
		return &m.left
	}

	return &m.right
}

func (m *Model) AddDigit(d string) {
	*m.currentNumber() += d
	m.updateView()
}

func (m *Model) AddDecimal() {
	if !strings.Contains(*m.currentNumber(), ".") {
		*m.currentNumber() += "."
	}
}

func (m *Model) AddOperator(op string) {
	if m.left == "" {
		if op == "-" {
			m.left = "-"
			m.updateView()
		}
		return
	}

	if op == "-" && m.operator != "" && m.right == "" {
		m.right = "-"
		m.updateView()
		return
	}

	if m.right != "" && m.operator != "" {
		m.Compute()
	}

	m.operator = op
	m.updateView()
}

func (m *Model) Compute() {
	if m.left == "" || m.right == "" {
		return
	}

	left, err := strconv.ParseFloat(m.left, 64)
	if err != nil {
		panic(err)
	}

	right, err := strconv.ParseFloat(m.right, 64)
	if err != nil {
		panic(err)
	}

	var result float64 = 0

	switch m.operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "/":
		result = left / right
	case "*":
		result = left * right
	}

	m.left = fmt.Sprint(result)
	m.operator = ""
	m.right = ""

	m.updateView()
}

func (m *Model) updateView() {
	if m.operator == "" {
		left := m.left
		if left == "" {
			left = "0"
		}
		m.view.SetText(left)
	} else {
		right := m.right
		if right == "" {
			right = "0"
		}
		m.view.SetText(fmt.Sprintf("%v %v %v", m.left, m.operator, m.right))
	}
}
