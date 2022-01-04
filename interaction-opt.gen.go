// retrieve options from interaction data
package corde

// String returns the option with key k of type string
func (o OptionsInteractions) String(k string) string {
	var v string
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Int returns the option with key k of type int
func (o OptionsInteractions) Int(k string) int {
	var v int
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Int32 returns the option with key k of type int32
func (o OptionsInteractions) Int32(k string) int32 {
	var v int32
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Int64 returns the option with key k of type int64
func (o OptionsInteractions) Int64(k string) int64 {
	var v int64
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Uint returns the option with key k of type uint
func (o OptionsInteractions) Uint(k string) uint {
	var v uint
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Uint32 returns the option with key k of type uint32
func (o OptionsInteractions) Uint32(k string) uint32 {
	var v uint32
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Uint64 returns the option with key k of type uint64
func (o OptionsInteractions) Uint64(k string) uint64 {
	var v uint64
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Float32 returns the option with key k of type float32
func (o OptionsInteractions) Float32(k string) float32 {
	var v float32
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Float64 returns the option with key k of type float64
func (o OptionsInteractions) Float64(k string) float64 {
	var v float64
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Bool returns the option with key k of type bool
func (o OptionsInteractions) Bool(k string) bool {
	var v bool
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Snowflake returns the option with key k of type Snowflake
func (o OptionsInteractions) Snowflake(k string) Snowflake {
	var v Snowflake
	_ = o[k].UnmarshalTo(&v)
	return v
}

// Any returns the option with key k of type any
func (o OptionsInteractions) Any(k string) any {
	var v any
	_ = o[k].UnmarshalTo(&v)
	return v
}
