// Package types defines types that encode values in SOAP requests and responses.
package types

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	// localLoc acts like time.Local for this package, but is faked out by the
	// unit tests to ensure that things stay constant (especially when running
	// this test in a place where local time is UTC which might mask bugs).
	localLoc = time.Local
)

type SOAPValue interface {
	Marshal() (string, error)
	Unmarshal(s string) error
}

type UI1 uint8

var _ SOAPValue = new(UI1)

func NewUI1(v uint8) *UI1 {
	v2 := UI1(v)
	return &v2
}

func (v *UI1) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI1) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI1) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 8)
	*v = UI1(v2)
	return err
}

type UI2 uint16

var _ SOAPValue = new(UI2)

func NewUI2(v uint16) *UI2 {
	v2 := UI2(v)
	return &v2
}

func (v *UI2) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI2) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI2) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 16)
	*v = UI2(v2)
	return err
}

type UI4 uint32

var _ SOAPValue = new(UI4)

func NewUI4(v uint32) *UI4 {
	v2 := UI4(v)
	return &v2
}

func (v *UI4) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI4) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 32)
	*v = UI4(v2)
	return err
}

type UI8 uint64

var _ SOAPValue = new(UI8)

func NewUI8(v uint64) *UI8 {
	v2 := UI8(v)
	return &v2
}

func (v *UI8) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UI8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *UI8) Unmarshal(s string) error {
	v2, err := strconv.ParseUint(s, 10, 64)
	*v = UI8(v2)
	return err
}

type I1 int8

var _ SOAPValue = new(I1)

func NewI1(v int8) *I1 {
	v2 := I1(v)
	return &v2
}

func (v *I1) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I1) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I1) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 8)
	*v = I1(v2)
	return err
}

type I2 int16

var _ SOAPValue = new(I2)

func NewI2(v int16) *I2 {
	v2 := I2(v)
	return &v2
}

func (v *I2) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I2) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I2) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 16)
	*v = I2(v2)
	return err
}

type I4 int32

var _ SOAPValue = new(I4)

func NewI4(v int32) *I4 {
	v2 := I4(v)
	return &v2
}

func (v *I4) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I4) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 32)
	*v = I4(v2)
	return err
}

type I8 int64

var _ SOAPValue = new(I8)

func NewI8(v int64) *I8 {
	v2 := I8(v)
	return &v2
}

func (v *I8) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *I8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *I8) Unmarshal(s string) error {
	v2, err := strconv.ParseInt(s, 10, 64)
	*v = I8(v2)
	return err
}

type R4 float32

var _ SOAPValue = new(R4)

func NewR4(v float32) *R4 {
	v2 := R4(v)
	return &v2
}

func (v *R4) String() string {
	return strconv.FormatFloat(float64(*v), 'G', -1, 32)
}

func (v *R4) Marshal() (string, error) {
	return v.String(), nil
}

func (v *R4) Unmarshal(s string) error {
	v2, err := strconv.ParseFloat(s, 32)
	*v = R4(v2)
	return err
}

type R8 float64

var _ SOAPValue = new(R8)

func NewR8(v float64) *R8 {
	v2 := R8(v)
	return &v2
}

func (v *R8) String() string {
	return strconv.FormatFloat(float64(*v), 'G', -1, 64)
}

func (v *R8) Marshal() (string, error) {
	return v.String(), nil
}

func (v *R8) Unmarshal(s string) error {
	v2, err := strconv.ParseFloat(s, 64)
	*v = R8(v2)
	return err
}

// FloatFixed14_4 maps a float64 to the SOAP "fixed.14.4" type.
type FloatFixed14_4 float64

var _ SOAPValue = new(FloatFixed14_4)

func NewFloatFixed14_4(v float64) *FloatFixed14_4 {
	v2 := FloatFixed14_4(v)
	return &v2
}

func (v *FloatFixed14_4) String() string {
	return strconv.FormatFloat(float64(*v), 'f', 4, 64)
}

func (v *FloatFixed14_4) Marshal() (string, error) {
	if *v >= 1e14 || *v <= -1e14 {
		return "", fmt.Errorf("soap fixed14.4: value %v out of bounds", v)
	}
	return v.String(), nil
}

func (v *FloatFixed14_4) Unmarshal(s string) error {
	v2, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	if v2 >= 1e14 || v2 <= -1e14 {
		return fmt.Errorf("soap fixed14.4: value %q out of bounds", s)
	}
	*v = FloatFixed14_4(v2)
	return nil
}

// Char maps rune to SOAP "char" type.
type Char rune

var _ SOAPValue = new(Char)

func NewChar(v rune) *Char {
	v2 := Char(v)
	return &v2
}

func (v *Char) String() string {
	return string(*v)
}

func (v *Char) Marshal() (string, error) {
	if *v == 0 {
		return "", errors.New("soap char: rune 0 is not allowed")
	}
	return v.String(), nil
}

func (v *Char) Unmarshal(s string) error {
	if len(s) == 0 {
		return errors.New("soap char: got empty string")
	}
	v2, n := utf8.DecodeRune([]byte(s))
	if n != len(s) {
		return fmt.Errorf("soap char: value %q is not a single rune", s)
	}
	*v = Char(v2)
	return nil
}

type String string

var _ SOAPValue = new(String)

func NewString(v string) *String {
	v2 := String(v)
	return &v2
}

func (v *String) Marshal() (string, error) {
	return string(*v), nil
}

func (v *String) Unmarshal(s string) error {
	*v = String(s)
	return nil
}

func parseInt(s string, err *error) int {
	v, parseErr := strconv.ParseInt(s, 10, 64)
	if parseErr != nil {
		*err = parseErr
	}
	return int(v)
}

var dateRegexps = []*regexp.Regexp{
	// yyyy[-mm[-dd]]
	regexp.MustCompile(`^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`),
	// yyyy[mm[dd]]
	regexp.MustCompile(`^(\d{4})(?:(\d{2})(?:(\d{2}))?)?$`),
}

func parseDateParts(s string) (year, month, day int, err error) {
	var parts []string
	for _, re := range dateRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		err = fmt.Errorf("soap date: value %q is not in a recognized ISO8601 date format", s)
		return
	}

	year = parseInt(parts[1], &err)
	month = 1
	day = 1
	if len(parts[2]) != 0 {
		month = parseInt(parts[2], &err)
		if len(parts[3]) != 0 {
			day = parseInt(parts[3], &err)
		}
	}

	if err != nil {
		err = fmt.Errorf("soap date: %q: %v", s, err)
	}

	return
}

var timeRegexps = []*regexp.Regexp{
	// hh[:mm[:ss]]
	regexp.MustCompile(`^(\d{2})(?::(\d{2})(?::(\d{2}))?)?$`),
	// hh[mm[ss]]
	regexp.MustCompile(`^(\d{2})(?:(\d{2})(?:(\d{2}))?)?$`),
}

func parseTimeParts(s string) (TimeOfDay, error) {
	var parts []string
	for _, re := range timeRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		return TimeOfDay{}, fmt.Errorf("soap time: value %q is not in ISO8601 time format", s)
	}

	var err error
	var hour, minute, second int8
	hour = int8(parseInt(parts[1], &err))
	if len(parts[2]) != 0 {
		minute = int8(parseInt(parts[2], &err))
		if len(parts[3]) != 0 {
			second = int8(parseInt(parts[3], &err))
		}
	}

	if err != nil {
		return TimeOfDay{}, fmt.Errorf("soap time: %q: %v", s, err)
	}

	return TimeOfDay{hour, minute, second}, nil
}

// (+|-)hh[[:]mm]
var timezoneRegexp = regexp.MustCompile(`^([+-])(\d{2})(?::?(\d{2}))?$`)

func parseTimezone(s string) (offset int, err error) {
	if s == "Z" {
		return 0, nil
	}
	parts := timezoneRegexp.FindStringSubmatch(s)
	if parts == nil {
		err = fmt.Errorf("soap timezone: value %q is not in ISO8601 timezone format", s)
		return
	}

	offset = parseInt(parts[2], &err) * 3600
	if len(parts[3]) != 0 {
		offset += parseInt(parts[3], &err) * 60
	}
	if parts[1] == "-" {
		offset = -offset
	}

	if err != nil {
		err = fmt.Errorf("soap timezone: %q: %v", s, err)
	}

	return
}

var completeDateTimeZoneRegexp = regexp.MustCompile(`^([^T]+)(?:T([^-+Z]+)(.+)?)?$`)

// splitCompleteDateTimeZone splits date, time and timezone apart from an
// ISO8601 string. It does not ensure that the contents of each part are
// correct, it merely splits on certain delimiters.
// e.g "2010-09-08T12:15:10+0700" => "2010-09-08", "12:15:10", "+0700".
// Timezone can only be present if time is also present.
func splitCompleteDateTimeZone(s string) (dateStr, timeStr, zoneStr string, err error) {
	parts := completeDateTimeZoneRegexp.FindStringSubmatch(s)
	if parts == nil {
		err = fmt.Errorf("soap date/time/zone: value %q is not in ISO8601 datetime format", s)
		return
	}
	dateStr = parts[1]
	timeStr = parts[2]
	zoneStr = parts[3]
	return
}

// TimeOfDay is used in cases where SOAP "time" or "time.tz" is used.
// It contains non-timezone aware components.
type TimeOfDay struct {
	Hour   int8
	Minute int8
	Second int8
}

var _ SOAPValue = &TimeOfDay{}

// Sets components based on duration since midnight.
func (v *TimeOfDay) SetFromDuration(d time.Duration) error {
	if d < 0 || d > 24*time.Hour {
		return fmt.Errorf("out of range of SOAP time type: %v", d)
	}
	v.Hour = int8(d / time.Hour)
	d = d % time.Hour
	v.Minute = int8(d / time.Minute)
	d = d % time.Minute
	v.Second = int8(d / time.Second)
	return nil
}

// Returns duration since midnight.
func (v *TimeOfDay) ToDuration() time.Duration {
	return time.Duration(v.Hour)*time.Hour +
		time.Duration(v.Minute)*time.Minute +
		time.Duration(v.Second)*time.Second
}

func (v *TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", v.Hour, v.Minute, v.Second)
}

func (v *TimeOfDay) Equal(o *TimeOfDay) bool {
	return v.Hour == o.Hour && v.Minute == o.Minute && v.Second == o.Second
}

// IsValid returns true iff v is positive and <= 24 hours.
// It allows equal to 24 hours as a special case as 24:00:00 is an allowed
// value by the SOAP type.
func (v *TimeOfDay) CheckValid() error {
	if (v.Hour < 0 || v.Minute < 0 || v.Second < 0) ||
		(v.Hour == 24 && (v.Minute > 0 || v.Second > 0)) ||
		v.Hour > 24 || v.Minute >= 60 || v.Second >= 60 {
		return fmt.Errorf("soap time: value %v has components(s) out of range", v)
	}
	return nil
}

func (v *TimeOfDay) Marshal() (string, error) {
	if err := v.CheckValid(); err != nil {
		return "", err
	}
	return v.String(), nil
}

func (v *TimeOfDay) Unmarshal(s string) error {
	var err error
	*v, err = parseTimeParts(s)
	if err != nil {
		return err
	}

	return v.CheckValid()
}

// TimeOfDayTZ is used in cases where SOAP "time.tz" is used.
type TimeOfDayTZ struct {
	// Components of the time of day.
	TimeOfDay TimeOfDay

	// Set to true if Offset is specified. If false, then the timezone is
	// unspecified (and by ISO8601 - implies some "local" time).
	HasOffset bool

	// Offset is non-zero only if time.tz is used. It is otherwise ignored. If
	// non-zero, then it is regarded as a UTC offset in seconds. Note that the
	// sub-minutes is ignored by the marshal function.
	Offset int
}

var _ SOAPValue = &TimeOfDayTZ{}

func (v *TimeOfDayTZ) String() string {
	return fmt.Sprintf("%v %t %+03d:%02d:%02d", v.TimeOfDay, v.HasOffset, v.Offset/3600, (v.Offset%3600)/60, v.Offset%60)
}

func (v *TimeOfDayTZ) Equal(o *TimeOfDayTZ) bool {
	return v.TimeOfDay.Equal(&o.TimeOfDay) &&
		v.HasOffset == o.HasOffset && v.Offset == o.Offset
}

func (v *TimeOfDayTZ) Marshal() (string, error) {
	tod, err := v.TimeOfDay.Marshal()
	if err != nil {
		return "", err
	}

	tz := ""
	if v.HasOffset {
		if v.Offset == 0 {
			tz = "Z"
		} else {
			offsetMins := v.Offset / 60
			sign := '+'
			if offsetMins < 1 {
				offsetMins = -offsetMins
				sign = '-'
			}
			tz = fmt.Sprintf("%c%02d:%02d", sign, offsetMins/60, offsetMins%60)
		}
	}

	return tod + tz, nil
}

func (v *TimeOfDayTZ) Unmarshal(s string) error {
	zoneIndex := strings.IndexAny(s, "Z+-")
	var timePart string
	if zoneIndex == -1 {
		v.HasOffset = false
		v.Offset = 0
		timePart = s
	} else {
		v.HasOffset = true
		timePart = s[:zoneIndex]
		var err error
		v.Offset, err = parseTimezone(s[zoneIndex:])
		if err != nil {
			return err
		}
	}

	if err := v.TimeOfDay.Unmarshal(timePart); err != nil {
		return err
	}

	return nil
}

// DateLocal maps time.Time to the SOAP "date" type. Dates map to midnight in
// the local time zone. The time of day components are ignored when
// marshalling.
type DateLocal time.Time

var _ SOAPValue = &DateLocal{}

func NewDateLocal(v time.Time) *DateLocal {
	v2 := DateLocal(v)
	return &v2
}

func (v DateLocal) String() string {
	return v.ToTime().String()
}

func (v DateLocal) ToTime() time.Time {
	return time.Time(v)
}

func (v *DateLocal) Marshal() (string, error) {
	return time.Time(*v).In(localLoc).Format("2006-01-02"), nil
}

func (v *DateLocal) Unmarshal(s string) error {
	year, month, day, err := parseDateParts(s)
	if err != nil {
		return err
	}
	*v = DateLocal(time.Date(year, time.Month(month), day, 0, 0, 0, 0, localLoc))
	return nil
}

// MarshalDateTime maps time.Time to SOAP "dateTime" type, with the local timezone.
type DateTimeLocal time.Time

var _ SOAPValue = &DateTimeLocal{}

func NewDateTimeLocal(v time.Time) *DateTimeLocal {
	v2 := DateTimeLocal(v)
	return &v2
}

func (v DateTimeLocal) String() string {
	return v.ToTime().String()
}

func (v DateTimeLocal) ToTime() time.Time {
	return time.Time(v)
}

func (v *DateTimeLocal) Marshal() (string, error) {
	return v.ToTime().In(localLoc).Format("2006-01-02T15:04:05"), nil
}

func (v *DateTimeLocal) Unmarshal(s string) error {
	dateStr, timeStr, zoneStr, err := splitCompleteDateTimeZone(s)
	if err != nil {
		return err
	}

	if len(zoneStr) != 0 {
		return fmt.Errorf("soap datetime: unexpected timezone in %q", s)
	}

	year, month, day, err := parseDateParts(dateStr)
	if err != nil {
		return err
	}

	var tod TimeOfDay
	if len(timeStr) != 0 {
		tod, err = parseTimeParts(timeStr)
		if err != nil {
			return err
		}
	}

	*v = DateTimeLocal(time.Date(year, time.Month(month), day,
		int(tod.Hour), int(tod.Minute), int(tod.Second), 0,
		localLoc))
	return nil
}

// DateTimeLocal maps time.Time to SOAP "dateTime.tz" type, using the local
// timezone when one is unspecified.
type DateTimeTZLocal time.Time

var _ SOAPValue = &DateTimeTZLocal{}

func NewDateTimeTZLocal(v time.Time) *DateTimeTZLocal {
	v2 := DateTimeTZLocal(v)
	return &v2
}

func (v DateTimeTZLocal) String() string {
	return v.ToTime().String()
}

func (v DateTimeTZLocal) ToTime() time.Time {
	return time.Time(v)
}

func (v *DateTimeTZLocal) Marshal() (string, error) {
	return time.Time(*v).Format("2006-01-02T15:04:05-07:00"), nil
}

func (v *DateTimeTZLocal) Unmarshal(s string) error {
	dateStr, timeStr, zoneStr, err := splitCompleteDateTimeZone(s)
	if err != nil {
		return err
	}

	year, month, day, err := parseDateParts(dateStr)
	if err != nil {
		return err
	}

	var tod TimeOfDay
	var location *time.Location = localLoc
	if len(timeStr) != 0 {
		tod, err = parseTimeParts(timeStr)
		if err != nil {
			return err
		}
		if len(zoneStr) != 0 {
			var offset int
			offset, err = parseTimezone(zoneStr)
			if err != nil {
				return err
			}
			if offset == 0 {
				location = time.UTC
			} else {
				location = time.FixedZone("", offset)
			}
		}
	}

	*v = DateTimeTZLocal(time.Date(year, time.Month(month), day,
		int(tod.Hour), int(tod.Minute), int(tod.Second), 0,
		location))
	return nil
}

type Boolean bool

var _ SOAPValue = new(Boolean)

func NewBoolean(v bool) *Boolean {
	v2 := Boolean(v)
	return &v2
}

func (v *Boolean) String() string {
	if *v {
		return "true"
	}
	return "false"
}

func (v *Boolean) Marshal() (string, error) {
	if *v {
		return "1", nil
	}
	return "0", nil
}

func (v *Boolean) Unmarshal(s string) error {
	switch s {
	case "0", "false", "no":
		*v = false
	case "1", "true", "yes":
		*v = true
	default:
		return fmt.Errorf("soap boolean: %q is not a valid boolean value", s)
	}
	return nil
}

// BinBase64 maps []byte to SOAP "bin.base64" type.
type BinBase64 []byte

var _ SOAPValue = new(BinBase64)

func NewBinBase64(v []byte) *BinBase64 {
	v2 := BinBase64(v)
	return &v2
}

func (v *BinBase64) String() string {
	return base64.StdEncoding.EncodeToString(*v)
}

func (v *BinBase64) Marshal() (string, error) {
	return v.String(), nil
}

func (v *BinBase64) Unmarshal(s string) error {
	v2, err := base64.StdEncoding.DecodeString(s)
	*v = v2
	return err
}

// BinHex maps []byte to SOAP "bin.hex" type.
type BinHex []byte

var _ SOAPValue = new(BinHex)

func NewBinHex(v []byte) *BinHex {
	v2 := BinHex(v)
	return &v2
}

func (v *BinHex) String() string {
	return hex.EncodeToString(*v)
}

func (v *BinHex) Marshal() (string, error) {
	return v.String(), nil
}

func (v *BinHex) Unmarshal(s string) error {
	v2, err := hex.DecodeString(s)
	*v = v2
	return err
}

// URI maps *url.URL to SOAP "uri" type.
type URI url.URL

var _ SOAPValue = new(URI)

func (v *URI) String() string {
	return v.ToURL().String()
}

func (v *URI) ToURL() *url.URL {
	return (*url.URL)(v)
}

func (v *URI) Marshal() (string, error) {
	return (*url.URL)(v).String(), nil
}

func (v *URI) Unmarshal(s string) error {
	v2, err := url.Parse(s)
	if err != nil {
		return err
	}
	*v = URI(*v2)
	return nil
}
