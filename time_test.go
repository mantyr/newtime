package newtime

import (
    "testing"
)

func TestParseDateStringZero(t *testing.T) {
    date := NewTime()
    date.Parse("2014.05.12", "yyyy-dd-mm")

    val := date.Format("YYYY-mm-dd")
    if val != "0000-00-00" {
        t.Errorf("Error zero in Format, %q", val)
    }
}

func TestParseDateString(t *testing.T) {
    date := NewTime()
    date.Parse("2014.05.12", "yyyy.dd.mm")

    val := date.Format("YYYY-mm-dd")
    if val != "2014-12-05" {
        t.Errorf("Error date in Format, %q", val)
    }
}

func TestParseDateStringOther(t *testing.T) {
    date := NewTime()
    date.Parse("29 марта 2016 03:51", "dd mm yyyy HH:ii", ShortMonthNames_ru2)

    val := date.Format("YYYY-mm-dd HH:ii:ss")
    if val != "2016-03-29 03:51:00" {
        t.Errorf("Error date in Format, %q", val)
    }
}

/*
func TestDuration(t *testing.T) {
    date := NewTime()
    date.Parse("Jan 29 2016, 03:20:14", "mm dd yyyy, HH:ii:ss", ShortMonths...)

    d := date.LastDurationDate()
    t.Errorf("%q", d)
}
*/

func TestFormat(t *testing.T) {
    date := NewTime()
    date.Parse("29 марта 2016 03:51", "dd mm yyyy HH:ii", ShortMonthNames_ru2)

    val := date.Format("YYYY/_long_month_/dd")
    if val != "2016/March/29" {
        t.Errorf("Error text date in Format, %q", val)
    }
}

func TestFormatTimezone(t *testing.T) {
    date := NewTime()
    date.Parse("2016-04-01T00:12:36", "yyyy-mm-ddTHH:ii:ss")
    val := date.Format("YYYY-mm-dd HH:ii:ss")
    if val != "2016-04-01 00:12:36" {
        t.Errorf("Error text date in Format, %q", val)
    }
}

var formatTests = []FormatTest{
    {"ANSIC",       ANSIC,       "Wed Feb  4 21:00:57 2009"},
    {"UnixDate",    UnixDate,    "Wed Feb  4 21:00:57 PST 2009"},
    {"RubyDate",    RubyDate,    "Wed Feb 04 21:00:57 -0800 2009"},
    {"RFC822",      RFC822,      "04 Feb 09 21:00 PST"},
    {"RFC850",      RFC850,      "Wednesday, 04-Feb-09 21:00:57 PST"},
    {"RFC1123",     RFC1123,     "Wed, 04 Feb 2009 21:00:57 PST"},
    {"RFC1123Z",    RFC1123Z,    "Wed, 04 Feb 2009 21:00:57 -0800"},
    {"RFC3339",     RFC3339,     "2009-02-04T21:00:57-08:00"},
    {"RFC3339Nano", RFC3339Nano, "2009-02-04T21:00:57.0123456-08:00"},
    {"Kitchen",     Kitchen,     "9:00PM"},

    {"am/pm", "3pm", "9pm"},
    {"AM/PM", "3PM", "9PM"},
    {"two-digit year", "06 01 02", "09 02 04"},

    // Three-letter months and days must not be followed by lower-case letter.
    {"Janet", "Hi Janet, the Month is January", "Hi Janet, the Month is February"},
    // Time stamps, Fractional seconds.
    {"Stamp",      Stamp,      "Feb  4 21:00:57"},
    {"StampMilli", StampMilli, "Feb  4 21:00:57.012"},
    {"StampMicro", StampMicro, "Feb  4 21:00:57.012345"},
    {"StampNano",  StampNano,  "Feb  4 21:00:57.012345600"},
}

func TestFormatList(t *testing.T) {
    // The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2010
    date := NewTime()
    date.SetLocal("PST8PDT")

    date.SetUnix(0, 1233810057012345600)
    for _, test := range formatTests {
        result := date.FormatRaw(test.format)
        if result != test.result {
            t.Errorf("%s expected %q got %q", test.name, test.result, result)
        }
    }
}

var parseTestsRaw = []ParseTest{
    {"ANSIC",    ANSIC,    "Thu Feb  4 21:00:57 2010",         "2010-02-04 21:00:57Z"},
    {"UnixDate", UnixDate, "Thu Feb  4 21:00:57 PST 2010",     "2010-02-04 21:00:57Z"},
    {"RubyDate", RubyDate, "Thu Feb 04 21:00:57 -0800 2010",   "2010-02-04 21:00:57-08:00"},
    {"RFC850",   RFC850,   "Thursday, 04-Feb-10 21:00:57 PST", "2010-02-04 21:00:57Z"},
    {"RFC1123",  RFC1123,  "Thu, 04 Feb 2010 21:00:57 PST",    "2010-02-04 21:00:57Z"},
    {"RFC1123",  RFC1123,  "Thu, 04 Feb 2010 22:00:57 PDT",    "2010-02-04 22:00:57Z"},
    {"RFC1123Z", RFC1123Z, "Thu, 04 Feb 2010 21:00:57 -0800",  "2010-02-04 21:00:57-08:00"},
    {"RFC3339",  RFC3339,  "2010-02-04T21:00:57-08:00",        "2010-02-04 21:00:57-08:00"},
    {"custom: \"2006-01-02 15:04:05-07\"", "2006-01-02 15:04:05-07", "2010-02-04 21:00:57-08", "2010-02-04 21:00:57-08:00"},
    // Optional fractional seconds.
    {"ANSIC",    ANSIC,    "Thu Feb  4 21:00:57.0 2010",            "2010-02-04 21:00:57Z"},
    {"UnixDate", UnixDate, "Thu Feb  4 21:00:57.01 PST 2010",       "2010-02-04 21:00:57Z"},
    {"RubyDate", RubyDate, "Thu Feb 04 21:00:57.012 -0800 2010",    "2010-02-04 21:00:57-08:00"},
    {"RFC850",   RFC850,   "Thursday, 04-Feb-10 21:00:57.0123 PST", "2010-02-04 21:00:57Z"},
    {"RFC1123",  RFC1123,  "Thu, 04 Feb 2010 21:00:57.01234 PST",   "2010-02-04 21:00:57Z"},
    {"RFC1123Z", RFC1123Z, "Thu, 04 Feb 2010 21:00:57.01234 -0800", "2010-02-04 21:00:57-08:00"},
    {"RFC3339",  RFC3339,  "2010-02-04T21:00:57.012345678-08:00",   "2010-02-04 21:00:57-08:00"},
    {"custom: \"2006-01-02 15:04:05\"", "2006-01-02 15:04:05", "2010-02-04 21:00:57.0", "2010-02-04 21:00:57Z"},
    // Amount of white space should not matter.
    {"ANSIC", ANSIC, "Thu Feb 4 21:00:57 2010",                  "2010-02-04 21:00:57Z"},
    {"ANSIC", ANSIC, "Thu      Feb     4     21:00:57     2010", "2010-02-04 21:00:57Z"},
    // Case should not matter
    {"ANSIC", ANSIC, "THU FEB 4 21:00:57 2010", "2010-02-04 21:00:57Z"},
    {"ANSIC", ANSIC, "thu feb 4 21:00:57 2010", "2010-02-04 21:00:57Z"},
    // Fractional seconds.
    {"millisecond", "Mon Jan _2 15:04:05.000 2006", "Thu Feb  4 21:00:57.012 2010",             "2010-02-04 21:00:57Z"},
    {"microsecond", "Mon Jan _2 15:04:05.000000 2006", "Thu Feb  4 21:00:57.012345 2010",       "2010-02-04 21:00:57Z"},
    {"nanosecond",  "Mon Jan _2 15:04:05.000000000 2006", "Thu Feb  4 21:00:57.012345678 2010", "2010-02-04 21:00:57Z"},
    // Leading zeros in other places should not be taken as fractional seconds.
    {"zero1", "2006.01.02.15.04.05.0", "2010.02.04.21.00.57.0",                                 "2010-02-04 21:00:57Z"},
    {"zero2", "2006.01.02.15.04.05.00", "2010.02.04.21.00.57.01",                               "2010-02-04 21:00:57Z"},
    // Month and day names only match when not followed by a lower-case letter.
    {"Janet", "Hi Janet, the Month is January: Jan _2 15:04:05 2006", "Hi Janet, the Month is February: Feb  4 21:00:57 2010", "2010-02-04 21:00:57Z"},

    // GMT with offset.
    {"GMT-8", UnixDate, "Fri Feb  5 05:00:57 GMT-8 2010",                                       "2010-02-04 21:00:57-08:00"},

    // Accept any number of fractional second digits (including none) for .999...
    // In Go 1, .999... was completely ignored in the format, meaning the first two
    // cases would succeed, but the next four would not. Go 1.1 accepts all six.
    {"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57 -0800 PST",                 "2010-02-04 21:00:57-08:00"},
    {"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57 -0800 PST",            "2010-02-04 21:00:57-08:00"},
    {"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST",            "2010-02-04 21:00:57-08:00"},
    {"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST",       "2010-02-04 21:00:57-08:00"},
    {"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST",       "2010-02-04 21:00:57-08:00"},
    {"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST",  "2010-02-04 21:00:57-08:00"},

    // issue 4502.
    {"", StampNano, "Feb  4 21:00:57.012345678",                                                "0000-02-04 21:00:57Z"},
    {"", "Jan _2 15:04:05.999", "Feb  4 21:00:57.012300000",                                    "0000-02-04 21:00:57Z"},
    {"", "Jan _2 15:04:05.999", "Feb  4 21:00:57.012345678",                                    "0000-02-04 21:00:57Z"},
    {"", "Jan _2 15:04:05.999999999", "Feb  4 21:00:57.0123",                                   "0000-02-04 21:00:57Z"},
    {"", "Jan _2 15:04:05.999999999", "Feb  4 21:00:57.012345678",                              "0000-02-04 21:00:57Z"},
}

func TestParseRaw(t *testing.T) {
    date := NewTime()
    date.SetUTC()
    for _, test := range parseTestsRaw {
        date.ParseRaw(test.value, test.format)
        val := date.Format("YYYY-mm-dd HH:ii:ssZ07:00")
        if val != test.result {
            t.Errorf("%s expected %q got %q", test.name, test.result, val)
        }
    }
}

var parseTests = []ParseTest{
    {"1_yyyy-mm-dd", "yyyy-mm-dd", "2015-02-18", "2015-02-18 00:00:00Z"},
    {"2_yyyy-dd-mm", "yyyy-dd-mm", "2015-18-20", "0000-00-00 00:00:00Z"},
    {"3_yyyy-mm-dd", "yyyy-dd-mm", "2015-18-02", "2015-02-18 00:00:00Z"},

    {"4_yyyy-mm-ddTHH:ii:ssZ0700", "yyyy-mm-ddTHH:ii:ssZ07:00", "2015-02-18T00:00:00-12:00", "2015-02-18 00:00:00-12:00"},
    {"5_yyyy-mm-dd HH:ii:ssZ0700", "yyyy-mm-dd HH:ii:ssZ07:00", "2015-02-18 00:00:00-12:00", "2015-02-18 00:00:00-12:00"},

    {"6_yyyy-mm-dd Z0700",  "yyyy-mm-dd Z0700",  "2015-02-18 -1200", "2015-02-18 00:00:00-12:00"},
    {"6_yyyy-mm-dd Z07:00", "yyyy-mm-dd Z07:00", "2015-02-18 -12:00", "2015-02-18 00:00:00-12:00"},
}

func TestParse(t *testing.T) {
    date := NewTime()
    date.SetUTC()
    for _, test := range parseTests {
        date.Parse(test.value, test.format)
        val := date.Format("YYYY-mm-dd HH:ii:ssZ07:00")
        if val != test.result {
            t.Errorf("%s expected %q got %q", test.name, test.result, val)
        }
    }
}