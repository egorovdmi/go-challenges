package strings

type StringBuilder struct {
	buffer []byte
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

func (sb *StringBuilder) Append(s string) {
	sb.buffer = append(sb.buffer, s...)
}

func (sb *StringBuilder) ToString() string {
	return string(sb.buffer)
}
