package logger

var (
    Debug   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func Init() {
    Debug=log.new(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    Info=log.new(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Warning=log.new(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error=log.new(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(message string) {
    Debug.Println(message)
}

func Info(message string) {
    Info.Println(message)
}

func Warning(message string) {
    Warning.Println(message)
}

func Error(message string) {
    Error.Println(message)
}

func Fatal(message string) {
    Error.Fatalln(message)
}

func Panic(message string) {
    Error.Panicln(message)
}
