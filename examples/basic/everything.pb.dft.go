// Code generated by github.com/j2gg0s/protoc-gen-dft. DO NOT EDIT.
// source: basic/everything.proto

package basic

func (m *Everything) SetDefault() {
	if m == nil {
		return
	}

	if m.GetDoubleValue() == 0 {
		m.DoubleValue = 1.5
	}

	if m.GetInt64Value() == 0 {
		m.Int64Value = 3
	}

	if len(m.GetStringValue()) == 0 {
		m.StringValue = "hello"
	}

	if len(m.GetDoubleValus()) == 0 {
		m.DoubleValus = append(m.DoubleValus, 1.5, 2.5, 3.5)
	}

	if len(m.GetInt64Values()) == 0 {
		m.Int64Values = append(m.Int64Values, 1, 2, 3)
	}

	if len(m.GetStringValues()) == 0 {
		m.StringValues = append(m.StringValues, "A", "B", "C")
	}

	// NoneValue has not default.

}