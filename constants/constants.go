package constants

import "os"

var GcOff = os.Getenv("GC") == "off"
