package newtime

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

/** Location format **/

var ShortMonthNames_en1 = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var ShortMonthNames_ru1 = []string{
	"---",
	"январь",
	"февраль",
	"март",
	"апрель",
	"май",
	"июнь",
	"июль",
	"август",
	"сентябрь",
	"октябрь",
	"ноябрь",
	"декабрь",
}

var ShortMonthNames_ru2 = []string{
	"---",
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

var ShortMonthNames_ua1 = []string{
	"---",
	"січень",
	"лютий",
	"березень",
	"квітень",
	"травень",
	"червень",
	"липень",
	"серпень",
	"вересень",
	"жовтень",
	"листопад",
	"грудень",
}

var ShortMonthNames_ua2 = []string{
	"---",
	"січня",
	"лютого",
	"березня",
	"квітня",
	"травня",
	"червня",
	"липня",
	"серпня",
	"вересня",
	"жовтня",
	"листопада",
	"грудня",
}
var ShortMonths = [][]string{
	ShortMonthNames_en1,

	ShortMonthNames_ru1,
	ShortMonthNames_ru2,

	ShortMonthNames_ua1,
	ShortMonthNames_ua2,
}

var ShortMonths_ua = [][]string{
	ShortMonthNames_ua1,
	ShortMonthNames_ua2,
}
