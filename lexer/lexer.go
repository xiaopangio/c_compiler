package lexer

// 定义category
const (
	KEYWORD    = "keyword"    //关键字
	OPERATOR   = "operator"   //操作符
	SEPARATOR  = "delimiter"  //分隔符
	IDENTIFIER = "identifier" //标识符
	INTEGER    = "integer"    //常整数
	CHAR       = "char"       //字符常量
	STRING     = "string"     //字符串常量
	FLOAT      = "float"      //实型常量
)

// 关键字表
var keywords [20]string //种别码 1-20
// 操作符表
var operators [28]string //种别码21-48
// 分隔符表
var delimiter [12]string //种别码50-61
// 种别码对应表
var categoryCode map[string]int

func init() {
	//初始化keywords
	keywords[0] = "int"
	keywords[1] = "float"
	keywords[2] = "const"
	keywords[3] = "void"
	keywords[4] = "char"
	keywords[5] = "double"
	keywords[6] = "struct"
	keywords[7] = "return"
	keywords[8] = "if"
	keywords[9] = "else"
	keywords[10] = "switch"
	keywords[11] = "case"
	keywords[12] = "default"
	keywords[13] = "while"
	keywords[14] = "for"
	keywords[15] = "do"
	keywords[16] = "break"
	keywords[17] = "continue"
	keywords[18] = "main"
	keywords[19] = "include"
	//初始化operators
	operators[0] = "+"
	operators[1] = "-"
	operators[2] = "*"
	operators[3] = "/"
	operators[4] = "="
	operators[5] = ">"
	operators[6] = "<"
	operators[7] = "!"
	operators[8] = "&"
	operators[9] = "|"
	operators[10] = "^"
	operators[11] = "%"
	operators[12] = "++"
	operators[13] = "--"
	operators[14] = "+="
	operators[15] = "-="
	operators[16] = "*="
	operators[17] = "/="
	operators[18] = "=="
	operators[19] = ">="
	operators[20] = "<="
	operators[21] = "!="
	operators[22] = "&&"
	operators[23] = "||"
	operators[24] = "&="
	operators[25] = "|="
	operators[26] = "^="
	operators[27] = "%="
	//初始化separator
	delimiter[0] = "("
	delimiter[1] = ")"
	delimiter[2] = "["
	delimiter[3] = "]"
	delimiter[4] = "{"
	delimiter[5] = "}"
	delimiter[6] = ","
	delimiter[7] = ";"
	delimiter[8] = "."
	delimiter[9] = ":"
	delimiter[10] = "#"
	delimiter[11] = "?"
	//初始化categoryCode
	categoryCode = make(map[string]int)
	for i := 0; i < 20; i++ {
		categoryCode[keywords[i]] = i + 1
	}
	for i := 0; i < 28; i++ {
		categoryCode[operators[i]] = i + 21
	}
	for i := 0; i < 12; i++ {
		categoryCode[delimiter[i]] = i + 50
	}
	categoryCode[IDENTIFIER] = 62
	categoryCode[INTEGER] = 63
	categoryCode[CHAR] = 64
	categoryCode[STRING] = 65
	categoryCode[FLOAT] = 66
}

type Token struct {
	Type     int    //种别码
	Category string //种别码对应的大类类型
	Value    string //种别码对应具体类型
}

func (t Token) String() string {
	return t.Category + ":" + t.Value
}

type Lexer struct {
}

func (l *Lexer) Parse(rawString string) []Token {
	return nil
}
