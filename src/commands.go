package src

const (
	OPENING_BRACE = "{"
	CLOSING_BRACE = "}"
	BRACES        = "{}"

	OPENING_BRACKET = "["
	CLOSING_BRACKET = "]"
	BRACKETS        = "[]"

	OPENING_PARENTHESIS = "("
	CLOSING_PARENTHESIS = ")"
	PARENTHESES         = "()"

	SUBSUP      = "_^"
	SUBSCRIPT   = "_"
	SUPERSCRIPT = "^"
	APOSTROPHE  = "'"
	PRIME       = "\\prime"
	DPRIME      = "\\dprime"

	LEFT   = "\\left"
	MIDDLE = "\\middle"
	RIGHT  = "\\right"

	ABOVE           = "\\above"
	ABOVEWITHDELIMS = "\\abovewithdelims"
	ATOP            = "\\atop"
	ATOPWITHDELIMS  = "\\atopwithdelims"
	BINOM           = "\\binom"
	BRACE           = "\\brace"
	BRACK           = "\\brack"
	CFRAC           = "\\cfrac"
	CHOOSE          = "\\choose"
	DBINOM          = "\\dbinom"
	DFRAC           = "\\dfrac"
	FRAC            = "\\frac"
	GENFRAC         = "\\genfrac"
	OVER            = "\\over"
	TBINOM          = "\\tbinom"
	TFRAC           = "\\tfrac"

	ROOT = "\\root"
	SQRT = "\\sqrt"

	OVERSET  = "\\overset"
	UNDERSET = "\\underset"

	ACUTE               = "\\acute"
	BAR                 = "\\bar"
	BREVE               = "\\breve"
	CHECK               = "\\check"
	DOT                 = "\\dot"
	DDOT                = "\\ddot"
	DDDOT               = "\\dddot"
	DDDDOT              = "\\ddddot"
	GRAVE               = "\\grave"
	HAT                 = "\\hat"
	MATHRING            = "\\mathring"
	OVERBRACE           = "\\overbrace"
	OVERLEFTARROW       = "\\overleftarrow"
	OVERLEFTRIGHTARROW  = "\\overleftrightarrow"
	OVERLINE            = "\\overline"
	OVERPAREN           = "\\overparen"
	OVERRIGHTARROW      = "\\overrightarrow"
	TILDE               = "\\tilde"
	UNDERBRACE          = "\\underbrace"
	UNDERLEFTARROW      = "\\underleftarrow"
	UNDERLINE           = "\\underline"
	UNDERPAREN          = "\\underparen"
	UNDERRIGHTARROW     = "\\underrightarrow"
	UNDERLEFTRIGHTARROW = "\\underleftrightarrow"
	VEC                 = "\\vec"
	WIDEHAT             = "\\widehat"
	WIDETILDE           = "\\widetilde"
	XLEFTARROW          = "\\xleftarrow"
	XRIGHTARROW         = "\\xrightarrow"

	HREF   = "\\href"
	TEXT   = "\\text"
	TEXTBF = "\\textbf"
	TEXTIT = "\\textit"
	TEXTRM = "\\textrm"
	TEXTSF = "\\textsf"
	TEXTTT = "\\texttt"

	BEGIN = "\\begin"
	END   = "\\end"

	LIMITS    = "\\limits"
	INTEGRAL  = "\\int"
	SUMMATION = "\\sum"
	// LIMIT 相关的常量

	LIM = "\\lim"
	SUP = "\\sup"
	INF = "\\inf"
	MAX = "\\max"
	MIN = "\\min"

	OPERATORNAME = "\\operatorname"

	LBRACE = "\\{"
)

var LIMIT = []string{
	LIM,
	SUP,
	INF,
	MAX,
	MIN,
}

var FUNCTIONS = []string{
	"\\arccos",
	"\\arcsin",
	"\\arctan",
	"\\cos",
	"\\cosh",
	"\\cot",
	"\\coth",
	"\\csc",
	"\\deg",
	"\\dim",
	"\\exp",
	"\\hom",
	"\\ker",
	"\\ln",
	"\\lg",
	"\\log",
	"\\sec",
	"\\sin",
	"\\sinh",
	"\\tan",
	"\\tanh",
}

const (
	DETERMINANT = "\\det"
	GCD         = "\\gcd"
	INTOP       = "\\intop"
	INJLIM      = "\\injlim"
	LIMINF      = "\\liminf"
	LIMSUP      = "\\limsup"
	PR          = "\\Pr"
	PROJLIM     = "\\projlim"
	MOD         = "\\mod"
	PMOD        = "\\pmod"
	BMOD        = "\\bmod"

	HDASHLINE = "\\hdashline"
	HLINE     = "\\hline"
	HFIL      = "\\hfil"

	CASES        = "\\cases"
	DISPLAYLINES = "\\displaylines"
	SMALLMATRIX  = "\\smallmatrix"
	SUBSTACK     = "\\substack"
	SPLIT        = "\\split"
	ALIGN        = "\\align*"
)

var MATRICES = []string{
	"\\matrix",
	"\\matrix*",
	"\\pmatrix",
	"\\pmatrix*",
	"\\bmatrix",
	"\\bmatrix*",
	"\\Bmatrix",
	"\\Bmatrix*",
	"\\vmatrix",
	"\\vmatrix*",
	"\\Vmatrix",
	"\\Vmatrix*",
	"\\array",
	SUBSTACK,
	CASES,
	DISPLAYLINES,
	SMALLMATRIX,
	SPLIT,
	ALIGN,
}

const (
	BACKSLASH       = "\\"
	CARRIAGE_RETURN = "\\cr"

	COLON           = "\\:"
	COMMA           = "\\,"
	DOUBLEBACKSLASH = "\\\\"
	ENSPACE         = "\\enspace"
	EXCLAMATION     = "\\!"
	GREATER_THAN    = "\\>"
	HSKIP           = "\\hskip"
	HSPACE          = "\\hspace"
	KERN            = "\\kern"
	MKERN           = "\\mkern"
	MSKIP           = "\\mskip"
	MSPACE          = "\\mspace"
	NEGTHINSPACE    = "\\negthinspace"
	NEGMEDSPACE     = "\\negmedspace"
	NEGTHICKSPACE   = "\\negthickspace"
	NOBREAKSPACE    = "\\nobreakspace"
	SPACE           = "\\space"
	THINSPACE       = "\\thinspace"
	QQUAD           = "\\qquad"
	QUAD            = "\\quad"
	SEMICOLON       = "\\;"

	BLACKBOARD_BOLD = "\\Bbb"
	BOLD_SYMBOL     = "\\boldsymbol"
	MIT             = "\\mit"
	OLDSTYLE        = "\\oldstyle"
	SCR             = "\\scr"
	TT              = "\\tt"

	MATH     = "\\math"
	MATHBB   = "\\mathbb"
	MATHBF   = "\\mathbf"
	MATHCAL  = "\\mathcal"
	MATHFRAK = "\\mathfrak"
	MATHIT   = "\\mathit"
	MATHRM   = "\\mathrm"
	MATHSCR  = "\\mathscr"
	MATHSF   = "\\mathsf"
	MATHTT   = "\\mathtt"

	BOXED = "\\boxed"
	FBOX  = "\\fbox"
	HBOX  = "\\hbox"
	MBOX  = "\\mbox"

	COLOR             = "\\color"
	DISPLAYSTYLE      = "\\displaystyle"
	TEXTSTYLE         = "\\textstyle"
	SCRIPTSTYLE       = "\\scriptstyle"
	SCRIPTSCRIPTSTYLE = "\\scriptscriptstyle"
	STYLE             = "\\style"

	HPHANTOM = "\\hphantom"
	PHANTOM  = "\\phantom"
	VPHANTOM = "\\vphantom"

	IDOTSINT = "\\idotsint"
	LATEX    = "\\LaTeX"
	TEX      = "\\TeX"

	SIDESET = "\\sideset"

	SKEW = "\\skew"
	NOT  = "\\not"
)

// FontFactory 是一个存储默认字体和替代字体的结构体
type FontFactory struct {
	Default     string
	Replacement map[string]string
}

// NewFontFactory 创建并返回一个新的FontFactory实例
func NewFontFactory(defaultFont string, replacement map[string]string) *FontFactory {
	return &FontFactory{
		Default:     defaultFont,
		Replacement: replacement,
	}
}

// Get 返回与给定键关联的字体，如果键不存在，则返回默认字体
func (f *FontFactory) Get(key string) string {
	if font, ok := f.Replacement[key]; ok {
		return font
	}
	return f.Default
}

// 初始化字体工厂的函数
func initFontFactory(defaultFont string, replacement map[string]string) *FontFactory {
	return NewFontFactory(defaultFont, replacement)
}

var localFonts = map[string]*FontFactory{
	BLACKBOARD_BOLD: initFontFactory("double-struck", map[string]string{"fence": ""}),
	BOLD_SYMBOL:     initFontFactory("bold", map[string]string{"mi": "bold-italic", "mtext": ""}),
	MATHBB:          initFontFactory("double-struck", map[string]string{"fence": ""}),
	MATHBF:          initFontFactory("bold", map[string]string{"fence": ""}),
	MATHCAL:         initFontFactory("script", map[string]string{"fence": ""}),
	MATHFRAK:        initFontFactory("fraktur", map[string]string{"fence": ""}),
	MATHIT:          initFontFactory("italic", map[string]string{"fence": ""}),
	MATHRM:          initFontFactory("", map[string]string{"mi": "normal"}),
	MATHSCR:         initFontFactory("script", map[string]string{"fence": ""}),
	MATHSF:          initFontFactory("", map[string]string{"mi": "sans-serif"}),
	MATHTT:          initFontFactory("monospace", map[string]string{"fence": ""}),
	MIT:             initFontFactory("italic", map[string]string{"fence": "", "mi": ""}),
	OLDSTYLE:        initFontFactory("normal", map[string]string{"fence": ""}),
	SCR:             initFontFactory("script", map[string]string{"fence": ""}),
	TT:              initFontFactory("monospace", map[string]string{"fence": ""}),
}

// 初始化 OLD_STYLE_FONTS
var oldStyleFonts = map[string]*FontFactory{
	"\\rm": initFontFactory("", map[string]string{"mi": "normal"}),
	"\\bf": initFontFactory("", map[string]string{"mi": "bold"}),
	"\\it": initFontFactory("", map[string]string{"mi": "italic"}),
	"\\sf": initFontFactory("", map[string]string{"mi": "sans-serif"}),
	"\\tt": initFontFactory("", map[string]string{"mi": "monospace"}),
}

// 初始化 GLOBAL_FONTS
var globalFonts = map[string]*FontFactory{
	"\\cal":  initFontFactory("script", map[string]string{"fence": ""}),
	"\\frak": initFontFactory("fraktur", map[string]string{"fence": ""}),
}

// initializeGlobalFonts 合并 OLD_STYLE_FONTS 到 GLOBAL_FONTS
func initializeGlobalFonts() {
	for k, v := range oldStyleFonts {
		globalFonts[k] = v
	}
}

// COMMANDS_WITH_ONE_PARAMETER 是代表带有一个参数的LaTeX命令的字符串切片
var COMMANDS_WITH_ONE_PARAMETER = []string{
	ACUTE,
	BAR,
	BLACKBOARD_BOLD,
	BOLD_SYMBOL,
	BOXED,
	BREVE,
	CHECK,
	DOT,
	DDOT,
	DDDOT,
	DDDDOT,
	GRAVE,
	HAT,
	HPHANTOM,
	MATHRING,
	MIT,
	MOD,
	OLDSTYLE,
	OVERBRACE,
	OVERLEFTARROW,
	OVERLEFTRIGHTARROW,
	OVERLINE,
	OVERPAREN,
	OVERRIGHTARROW,
	PHANTOM,
	PMOD,
	SCR,
	TILDE,
	TT,
	UNDERBRACE,
	UNDERLEFTARROW,
	UNDERLINE,
	UNDERPAREN,
	UNDERRIGHTARROW,
	UNDERLEFTRIGHTARROW,
	VEC,
	VPHANTOM,
	WIDEHAT,
	WIDETILDE,
}

var COMMANDS_WITH_TWO_PARAMETERS = []string{
	BINOM,
	CFRAC,
	DBINOM,
	DFRAC,
	FRAC,
	OVERSET,
	TBINOM,
	TFRAC,
	UNDERSET,
}

// MathMLElement 用于表示MathML元素及其属性
type MathMLElement struct {
	Element    string
	Attributes map[string]string
}

// BIG 存储与LaTeX命令相关的MathML元素及其属性
var BIG = map[string]MathMLElement{
	"\\Bigg": {Element: "mo", Attributes: map[string]string{"minsize": "2.470em", "maxsize": "2.470em"}},
	"\\bigg": {Element: "mo", Attributes: map[string]string{"minsize": "2.047em", "maxsize": "2.047em"}},
	"\\Big":  {Element: "mo", Attributes: map[string]string{"minsize": "1.623em", "maxsize": "1.623em"}},
	"\\big":  {Element: "mo", Attributes: map[string]string{"minsize": "1.2em", "maxsize": "1.2em"}},
}

// BIG_OPEN_CLOSE 基于BIG生成的映射
var BIG_OPEN_CLOSE = make(map[string]MathMLElement)

func initBigOpenClose() {
	// 构建 BIG_OPEN_CLOSE
	for command, element := range BIG {
		for _, postfix := range []string{"l", "m", "r"} {
			key := command + postfix
			BIG_OPEN_CLOSE[key] = MathMLElement{
				Element: element.Element,
				Attributes: map[string]string{
					"stretchy": "true",
					"fence":    "true",
				},
			}
			// 将 BIG 中的属性也添加到 BIG_OPEN_CLOSE
			for attrKey, attrValue := range element.Attributes {
				BIG_OPEN_CLOSE[key].Attributes[attrKey] = attrValue
			}
		}
	}
}

// MSTYLE_SIZES 存储与大小相关的LaTeX命令及其MathML等效项和属性
var MSTYLE_SIZES = map[string]MathMLElement{
	"\\Huge":       {Element: "mstyle", Attributes: map[string]string{"mathsize": "2.49em"}},
	"\\huge":       {Element: "mstyle", Attributes: map[string]string{"mathsize": "2.07em"}},
	"\\LARGE":      {Element: "mstyle", Attributes: map[string]string{"mathsize": "1.73em"}},
	"\\Large":      {Element: "mstyle", Attributes: map[string]string{"mathsize": "1.44em"}},
	"\\large":      {Element: "mstyle", Attributes: map[string]string{"mathsize": "1.2em"}},
	"\\normalsize": {Element: "mstyle", Attributes: map[string]string{"mathsize": "1em"}},
	"\\scriptsize": {Element: "mstyle", Attributes: map[string]string{"mathsize": "0.7em"}},
	"\\small":      {Element: "mstyle", Attributes: map[string]string{"mathsize": "0.85em"}},
	"\\tiny":       {Element: "mstyle", Attributes: map[string]string{"mathsize": "0.5em"}},
	"\\Tiny":       {Element: "mstyle", Attributes: map[string]string{"mathsize": "0.6em"}},
}

// STYLES 存储与样式相关的LaTeX命令及其MathML等效项和属性
var STYLES = map[string]MathMLElement{
	DISPLAYSTYLE:      {Element: "mstyle", Attributes: map[string]string{"displaystyle": "true", "scriptlevel": "0"}},
	TEXTSTYLE:         {Element: "mstyle", Attributes: map[string]string{"displaystyle": "false", "scriptlevel": "0"}},
	SCRIPTSTYLE:       {Element: "mstyle", Attributes: map[string]string{"displaystyle": "false", "scriptlevel": "1"}},
	SCRIPTSCRIPTSTYLE: {Element: "mstyle", Attributes: map[string]string{"displaystyle": "false", "scriptlevel": "2"}},
}

// MathMLElement 结构体和相关映射（BIG, BIG_OPEN_CLOSE, MATRICES, MSTYLE_SIZES, STYLES）的定义...

// CONVERSION_MAP 存储LaTeX命令到MathML元素的映射
var CONVERSION_MAP = make(map[string]MathMLElement)

func initConversionMap() {
	// 添加 MATRICES 相关的映射
	for _, matrix := range MATRICES {
		CONVERSION_MAP[matrix] = MathMLElement{Element: "mtable", Attributes: map[string]string{}}
	}

	// 添加其他映射
	additionalMappings := map[string]MathMLElement{
		DISPLAYLINES: {"mtable", map[string]string{"rowspacing": "0.5em", "columnspacing": "1em", "displaystyle": "true"}},
		SMALLMATRIX:  {"mtable", map[string]string{"rowspacing": "0.1em", "columnspacing": "0.2778em"}},
		SPLIT:        {"mtable", map[string]string{"displaystyle": "true", "columnspacing": "0em", "rowspacing": "3pt"}},
		ALIGN:        {"mtable", map[string]string{"displaystyle": "true", "rowspacing": "3pt"}},
		SUBSCRIPT:    {"msub", map[string]string{}},
		SUPERSCRIPT:  {"msup", map[string]string{}},
		SUBSUP:       {"msubsup", map[string]string{}},
		BINOM:        {"mfrac", map[string]string{"linethickness": "0"}},
		CFRAC:        {"mfrac", map[string]string{}},
		DBINOM:       {"mfrac", map[string]string{"linethickness": "0"}},
		DFRAC:        {"mfrac", map[string]string{}},
		FRAC:         {"mfrac", map[string]string{}},
		GENFRAC:      {"mfrac", map[string]string{}},
		TBINOM:       {"mfrac", map[string]string{"linethickness": "0"}},
		TFRAC:        {"mfrac", map[string]string{}},
		//ACUTE: ("mover", {}),
		//BAR: ("mover", {}),
		//BREVE: ("mover", {}),
		//CHECK: ("mover", {}),
		//DOT: ("mover", {}),
		//DDOT: ("mover", {}),
		//DDDOT: ("mover", {}),
		//DDDDOT: ("mover", {}),
		ACUTE:  {"mover", map[string]string{}},
		BAR:    {"mover", map[string]string{}},
		BREVE:  {"mover", map[string]string{}},
		CHECK:  {"mover", map[string]string{}},
		DOT:    {"mover", map[string]string{}},
		DDOT:   {"mover", map[string]string{}},
		DDDOT:  {"mover", map[string]string{}},
		DDDDOT: {"mover", map[string]string{}},

		//GRAVE: ("mover", {}),
		//HAT: ("mover", {}),
		//LIMITS: ("munderover", {}),
		//MATHRING: ("mover", {}),
		//OVERBRACE: ("mover", {}),
		//OVERLEFTARROW: ("mover", {}),
		//OVERLEFTRIGHTARROW: ("mover", {}),
		//OVERLINE: ("mover", {}),
		//OVERPAREN: ("mover", {}),
		//OVERRIGHTARROW: ("mover", {}),
		//TILDE: ("mover", {}),
		//OVERSET: ("mover", {}),
		//UNDERBRACE: ("munder", {}),
		//UNDERLEFTARROW: ("munder", {}),
		//UNDERLINE: ("munder", {}),
		//UNDERPAREN: ("munder", {}),
		//UNDERRIGHTARROW: ("munder", {}),
		//UNDERLEFTRIGHTARROW: ("munder", {}),
		//UNDERSET: ("munder", {}),
		//VEC: ("mover", {}),
		//WIDEHAT: ("mover", {}),
		//WIDETILDE: ("mover", {}),

		GRAVE:               {"mover", map[string]string{}},
		HAT:                 {"mover", map[string]string{}},
		LIMITS:              {"munderover", map[string]string{}},
		MATHRING:            {"mover", map[string]string{}},
		OVERBRACE:           {"mover", map[string]string{}},
		OVERLEFTARROW:       {"mover", map[string]string{}},
		OVERLEFTRIGHTARROW:  {"mover", map[string]string{}},
		OVERLINE:            {"mover", map[string]string{}},
		OVERPAREN:           {"mover", map[string]string{}},
		OVERRIGHTARROW:      {"mover", map[string]string{}},
		TILDE:               {"mover", map[string]string{}},
		OVERSET:             {"mover", map[string]string{}},
		UNDERBRACE:          {"munder", map[string]string{}},
		UNDERLEFTARROW:      {"munder", map[string]string{}},
		UNDERLINE:           {"munder", map[string]string{}},
		UNDERPAREN:          {"munder", map[string]string{}},
		UNDERRIGHTARROW:     {"munder", map[string]string{}},
		UNDERLEFTRIGHTARROW: {"munder", map[string]string{}},
		UNDERSET:            {"munder", map[string]string{}},
		VEC:                 {"mover", map[string]string{}},
		WIDEHAT:             {"mover", map[string]string{}},
		WIDETILDE:           {"mover", map[string]string{}},

		//# spaces
		//COLON: ("mspace", {"width": "0.222em"}),
		//COMMA: ("mspace", {"width": "0.167em"}),
		//DOUBLEBACKSLASH: ("mspace", {"linebreak": "newline"}),
		//ENSPACE: ("mspace", {"width": "0.5em"}),
		//EXCLAMATION: ("mspace", {"width": "negativethinmathspace"}),
		//GREATER_THAN: ("mspace", {"width": "0.222em"}),
		//HSKIP: ("mspace", {}),
		//HSPACE: ("mspace", {}),
		//KERN: ("mspace", {}),
		//MKERN: ("mspace", {}),
		//MSKIP: ("mspace", {}),
		//MSPACE: ("mspace", {}),
		//NEGTHINSPACE: ("mspace", {"width": "negativethinmathspace"}),
		//NEGMEDSPACE: ("mspace", {"width": "negativemediummathspace"}),
		//NEGTHICKSPACE: ("mspace", {"width": "negativethickmathspace"}),
		//THINSPACE: ("mspace", {"width": "thinmathspace"}),
		//QQUAD: ("mspace", {"width": "2em"}),
		//QUAD: ("mspace", {"width": "1em"}),
		//SEMICOLON: ("mspace", {"width": "0.278em"}),

		COLON:           {"mspace", map[string]string{"width": "0.222em"}},
		COMMA:           {"mspace", map[string]string{"width": "0.167em"}},
		DOUBLEBACKSLASH: {"mspace", map[string]string{"linebreak": "newline"}},
		ENSPACE:         {"mspace", map[string]string{"width": "0.5em"}},
		EXCLAMATION:     {"mspace", map[string]string{"width": "negativethinmathspace"}},
		GREATER_THAN:    {"mspace", map[string]string{"width": "0.222em"}},
		HSKIP:           {"mspace", map[string]string{}},
		HSPACE:          {"mspace", map[string]string{}},
		KERN:            {"mspace", map[string]string{}},
		MKERN:           {"mspace", map[string]string{}},
		MSKIP:           {"mspace", map[string]string{}},
		MSPACE:          {"mspace", map[string]string{}},
		NEGTHINSPACE:    {"mspace", map[string]string{"width": "negativethinmathspace"}},
		NEGMEDSPACE:     {"mspace", map[string]string{"width": "negativemediummathspace"}},
		NEGTHICKSPACE:   {"mspace", map[string]string{"width": "negativethickmathspace"}},
		THINSPACE:       {"mspace", map[string]string{"width": "thinmathspace"}},
		QQUAD:           {"mspace", map[string]string{"width": "2em"}},
		QUAD:            {"mspace", map[string]string{"width": "1em"}},
		SEMICOLON:       {"mspace", map[string]string{"width": "0.278em"}},

		//# enclose
		//BOXED: ("menclose", {"notation": "box"}),
		//FBOX: ("menclose", {"notation": "box"}),

		BOXED: {"menclose", map[string]string{"notation": "box"}},
		FBOX:  {"menclose", map[string]string{"notation": "box"}},

		//# operators
		//**BIG,
		//**BIG_OPEN_CLOSE,
		//**MSTYLE_SIZES,
		//**{limit: ("mo", {}) for limit in LIMIT},

		//LEFT: ("mo", OrderedDict([("stretchy", "true"), ("fence", "true"), ("form", "prefix")])),
		//MIDDLE: ("mo", OrderedDict([("stretchy", "true"), ("fence", "true"), ("lspace", "0.05em"), ("rspace", "0.05em")])),
		//RIGHT: ("mo", OrderedDict([("stretchy", "true"), ("fence", "true"), ("form", "postfix")])),

		LEFT:   {"mo", map[string]string{"stretchy": "true", "fence": "true", "form": "prefix"}},
		MIDDLE: {"mo", map[string]string{"stretchy": "true", "fence": "true", "lspace": "0.05em", "rspace": "0.05em"}},
		RIGHT:  {"mo", map[string]string{"stretchy": "true", "fence": "true", "form": "postfix"}},

		//# styles
		//COLOR: ("mstyle", {}),

		COLOR: {"mstyle", map[string]string{}},

		//**STYLES,
		//# others
		//SQRT: ("msqrt", {}),
		//ROOT: ("mroot", {}),
		//HREF: ("mtext", {}),
		//TEXT: ("mtext", {}),
		//TEXTBF: ("mtext", {"mathvariant": "bold"}),
		//TEXTIT: ("mtext", {"mathvariant": "italic"}),
		//TEXTRM: ("mtext", {}),
		//TEXTSF: ("mtext", {"mathvariant": "sans-serif"}),
		//TEXTTT: ("mtext", {"mathvariant": "monospace"}),
		//HBOX: ("mtext", {}),
		//MBOX: ("mtext", {}),
		//HPHANTOM: ("mphantom", {}),
		//PHANTOM: ("mphantom", {}),
		//VPHANTOM: ("mphantom", {}),
		//SIDESET: ("mrow", {}),
		//SKEW: ("mrow", {}),
		//MOD: ("mi", {}),
		//PMOD: ("mi", {}),
		//BMOD: ("mo", {}),
		//XLEFTARROW: ("mover", {}),
		//XRIGHTARROW: ("mover", {}),

		SQRT:        {"msqrt", map[string]string{}},
		ROOT:        {"mroot", map[string]string{}},
		HREF:        {"mtext", map[string]string{}},
		TEXT:        {"mtext", map[string]string{}},
		TEXTBF:      {"mtext", map[string]string{"mathvariant": "bold"}},
		TEXTIT:      {"mtext", map[string]string{"mathvariant": "italic"}},
		TEXTRM:      {"mtext", map[string]string{}},
		TEXTSF:      {"mtext", map[string]string{"mathvariant": "sans-serif"}},
		TEXTTT:      {"mtext", map[string]string{"mathvariant": "monospace"}},
		HBOX:        {"mtext", map[string]string{}},
		MBOX:        {"mtext", map[string]string{}},
		HPHANTOM:    {"mphantom", map[string]string{}},
		PHANTOM:     {"mphantom", map[string]string{}},
		VPHANTOM:    {"mphantom", map[string]string{}},
		SIDESET:     {"mrow", map[string]string{}},
		SKEW:        {"mrow", map[string]string{}},
		MOD:         {"mi", map[string]string{}},
		PMOD:        {"mi", map[string]string{}},
		BMOD:        {"mo", map[string]string{}},
		XLEFTARROW:  {"mover", map[string]string{}},
		XRIGHTARROW: {"mover", map[string]string{}},

		// ... 添加更多映射 ...
	}

	for k, v := range additionalMappings {
		CONVERSION_MAP[k] = v
	}

	// 合并 BIG, BIG_OPEN_CLOSE, MSTYLE_SIZES, STYLES 等映射
	for k, v := range BIG {
		CONVERSION_MAP[k] = v
	}
	for k, v := range BIG_OPEN_CLOSE {
		CONVERSION_MAP[k] = v
	}
	for k, v := range MSTYLE_SIZES {
		CONVERSION_MAP[k] = v
	}
	for k, v := range STYLES {
		CONVERSION_MAP[k] = v
	}
	// 合并 LIMIT 相关的命令
	for _, limit := range LIMIT {
		CONVERSION_MAP[limit] = MathMLElement{Element: "mo", Attributes: map[string]string{}}
	}

}

// Diacritic 表示一个字符实体引用和其属性
type Diacritic struct {
	Entity     string
	Attributes map[string]string
}

// DIACRITICS 存储LaTeX命令到字符实体引用和属性的映射
var DIACRITICS = map[string]Diacritic{
	ACUTE:               {"&#x000B4;", map[string]string{}},
	BAR:                 {"&#x000AF;", map[string]string{"stretchy": "true"}},
	BREVE:               {"&#x002D8;", map[string]string{}},
	CHECK:               {"&#x002C7;", map[string]string{}},
	DOT:                 {"&#x002D9;", map[string]string{}},
	DDOT:                {"&#x000A8;", map[string]string{}},
	DDDOT:               {"&#x020DB;", map[string]string{}},
	DDDDOT:              {"&#x020DC;", map[string]string{}},
	GRAVE:               {"&#x00060;", map[string]string{}},
	HAT:                 {"&#x0005E;", map[string]string{"stretchy": "false"}},
	MATHRING:            {"&#x002DA;", map[string]string{}},
	OVERBRACE:           {"&#x23DE;", map[string]string{}},
	OVERLEFTARROW:       {"&#x02190;", map[string]string{}},
	OVERLEFTRIGHTARROW:  {"&#x02194;", map[string]string{}},
	OVERLINE:            {"&#x02015;", map[string]string{"accent": "true"}},
	OVERPAREN:           {"&#x023DC;", map[string]string{}},
	OVERRIGHTARROW:      {"&#x02192;", map[string]string{}},
	TILDE:               {"&#x0007E;", map[string]string{"stretchy": "false"}},
	UNDERBRACE:          {"&#x23DF;", map[string]string{}},
	UNDERLEFTARROW:      {"&#x02190;", map[string]string{}},
	UNDERLEFTRIGHTARROW: {"&#x02194;", map[string]string{}},
	UNDERLINE:           {"&#x02015;", map[string]string{"accent": "true"}},
	UNDERPAREN:          {"&#x023DD;", map[string]string{}},
	UNDERRIGHTARROW:     {"&#x02192;", map[string]string{}},
	VEC:                 {"&#x02192;", map[string]string{"stretchy": "true"}},
	WIDEHAT:             {"&#x0005E;", map[string]string{}},
	WIDETILDE:           {"&#x0007E;", map[string]string{}},
}

func init() {
	initializeGlobalFonts()
	initBigOpenClose()
	initConversionMap()
}
