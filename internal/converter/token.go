package converter

type token struct {
	tokenType string
	lexeme    string
	line      int
	column    int
}

//token types for netlogo
var (

	// Keywords that are used with breeds
	breededTokens = map[string]interface{}{
		"CREATE-ORDERED_BREEDS": nil,
		"CREATE-BREED-TO":       nil,
		"CREATE-BREEDS-TO":      nil,
		"CREATE-BREED-FROM":     nil,
		"CREATE-BREEDS-FROM":    nil,
		"CREATE-BREED-WITH":     nil,
		"CREATE-BREEDS-WITH":    nil,
		"CREATE-BREEDS":         nil,
		"IN-BREED-NEIGHBOR":     nil,
		"IN-BREED-NEIGHBORS":    nil,
		"IN-BREED-FROM":         nil,
		"MY-BREEDS":             nil,
		"MY-IN-BREEDS":          nil,
		"MY-OUT-BREEDS":         nil,
		"BREED-NEIGHBOR":        nil,
		"OUT-BREED-NEIGHBOR":    nil,
		"OUT-BREED-NEIGHBORS":   nil,
		"OUT-BREED-TO":          nil,
		"BREEDS-AT":             nil,
		"BREEDS-HERE":           nil,
		"BREEDS-ON":             nil,
		"BREEDS-OWN":            nil,
		"BREED-WITH":            nil,
	}

	noSpaceTokens = map[string]interface{}{
		"(": nil,
		")": nil,
		"[": nil,
		"]": nil,
		"{": nil,
		"}": nil,
		",": nil,
		":": nil,
		";": nil,
		"=": nil,
		"+": nil,
		"*": nil,
		"/": nil,
		"^": nil,
		"<": nil,
		">": nil,
	}

	doubleNoSpaceTokens = map[string]interface{}{
		"!=": nil,
		"<=": nil,
		">=": nil,
		"->": nil,
	}

	valueTokens = map[string]interface{}{
		"NUMBER":  nil,
		"STRING":  nil,
		"BOOLEAN": nil,
		"VAR":     nil,
		"NEWLINE": nil,
	}

	// Keywords
	keywordToTokenType = map[string]string{
		"e":                             "E",
		"pi":                            "PI",
		"true":                          "TRUE",
		"false":                         "FALSE",
		"black":                         "BLACK",
		"gray":                          "GRAY",
		"white":                         "WHITE",
		"red":                           "RED",
		"orange":                        "ORANGE",
		"brown":                         "BROWN",
		"yellow":                        "YELLOW",
		"green":                         "GREEN",
		"lime":                          "LIME",
		"turquoise":                     "TURQUOISE",
		"cyan":                          "CYAN",
		"sky":                           "SKY",
		"blue":                          "BLUE",
		"violet":                        "VIOLET",
		"magenta":                       "MAGENTA",
		"pink":                          "PINK",
		"abs":                           "ABS",
		"acos":                          "ACOS",
		"all":                           "ALL",
		"and":                           "AND",
		"any?":                          "ANY",
		"approximate-hsb":               "APPROXIMATE-HSB",
		"approximate-rgb":               "APPROXIMATE-RGB",
		"+":                             "PLUS",
		"*":                             "MULTIPLY",
		"-":                             "MINUS",
		"/":                             "DIVIDE",
		"^":                             "POWER",
		"<":                             "LESS",
		">":                             "GREATER",
		"<=":                            "LESS-EQUAL",
		">=":                            "GREATER-EQUAL",
		"= ":                            "EQUAL",
		"!= ":                           "NOT-EQUAL",
		"[":                             "OPEN-BRACKET",
		"]":                             "CLOSE-BRACKET",
		"(":                             "OPEN-PARENTHESIS",
		")":                             "CLOSE-PARENTHESIS",
		"asin":                          "ASIN",
		"ask":                           "ASK",
		"at-points":                     "AT-POINTS",
		"atan":                          "ATAN",
		"back":                          "BACK",
		"bk":                            "BACK",
		"base-colors":                   "BASE-COLORS",
		"both-ends":                     "BOTH-ENDS",
		"breed":                         "BREED",
		"but-first":                     "BUT-FIRST",
		"butfirst":                      "BUT-FIRST",
		"bf":                            "BUT-FIRST",
		"but-last":                      "BUT-LAST",
		"butlast":                       "BUT-LAST",
		"bl":                            "BUT-LAST",
		"can-move?":                     "CAN-MOVE",
		"ceiling":                       "CEILING",
		"clear-all":                     "CLEAR-ALL",
		"clear-all-plots":               "CLEAR-ALL-PLOTS",
		"clear-drawing":                 "CLEAR-DRAWING",
		"cd":                            "CLEAR-DRAWING",
		"clear-globals":                 "CLEAR-GLOBALS",
		"clear-links":                   "CLEAR-LINKS",
		"clear-output":                  "CLEAR-OUTPUT",
		"clear-patches":                 "CLEAR-PATCHES",
		"cp":                            "CLEAR-PATCHES",
		"clear-plot":                    "CLEAR-PLOT",
		"clear-plots":                   "CLEAR-PLOTS",
		"clear-ticks":                   "CLEAR-TICKS",
		"ct":                            "CLEAR-TICKS",
		"color":                         "COLOR",
		"cos":                           "COS",
		"count":                         "COUNT",
		"create-ordered-turtles":        "CREATE-ORDERED-TURTLES",
		"cro":                           "CREATE-ORDERED-TURTLES",
		"create-link-to":                "CREATE-LINK-TO",
		"create-links-to":               "CREATE-LINKS-TO",
		"create-link-from":              "CREATE-LINK-FROM",
		"create-links-from":             "CREATE-LINKS-FROM",
		"create-link-with":              "CREATE-LINK-WITH",
		"create-links-with":             "CREATE-LINKS-WITH",
		"create-turtles":                "CREATE-TURTLES",
		"crt":                           "CREATE-TURTLES",
		"date-and-time":                 "DATE-AND-TIME",
		"die":                           "DIE",
		"diffuse":                       "DIFFUSE",
		"diffuse4":                      "DIFFUSE4",
		"directed-link-breed":           "DIRECTED-LINK-BREED",
		"display":                       "DISPLAY",
		"distance":                      "DISTANCE",
		"distancexy":                    "DISTANCEXY",
		"downhill":                      "DOWNHILL",
		"downhill4":                     "DOWNHILL4",
		"dx":                            "DX",
		"dy":                            "DY",
		"empty?":                        "EMPTY",
		"end":                           "END",
		"end1":                          "END1",
		"end2":                          "END2",
		"error":                         "ERROR",
		"error-message":                 "ERROR-MESSAGE",
		"every":                         "EVERY",
		"exp":                           "EXP",
		"extensions":                    "EXTENSIONS",
		"extract-hsb":                   "EXTRACT-HSB",
		"extract-rgb":                   "EXTRACT-RGB",
		"face":                          "FACE",
		"facexy":                        "FACEXY",
		"filter":                        "FILTER",
		"first":                         "FIRST",
		"floor":                         "FLOOR",
		"follow":                        "FOLLOW",
		"follow-me":                     "FOLLOW-ME",
		"foreach":                       "FOREACH",
		"forward":                       "FORWARD",
		"fd":                            "FORWARD",
		"fput":                          "FPUT",
		"globals":                       "GLOBALS",
		"hatch":                         "HATCH",
		"heading":                       "HEADING",
		"hidden?":                       "HIDDEN",
		"hide-link":                     "HIDE-LINK",
		"hide-turtle":                   "HIDE-TURTLE",
		"ht":                            "HIDE-TURTLE",
		"histogram":                     "HISTOGRAM",
		"home":                          "HOME",
		"hsb":                           "HSB",
		"if":                            "IF",
		"ifelse":                        "IFELSE",
		"ifelse-value":                  "IFELSE-VALUE",
		"import-drawing":                "IMPORT-DRAWING",
		"import-pcolors":                "IMPORT-PCOLORS",
		"import-pcolors-rgb":            "IMPORT-PCOLORS-RGB",
		"import-world":                  "IMPORT-WORLD",
		"in-cone":                       "IN-CONE",
		"in-link-neighbor?":             "IN-LINK-NEIGHBOR",
		"in-link-neighbors":             "IN-LINK-NEIGHBORS",
		"in-link-from":                  "IN-LINK-FROM",
		"in-radius":                     "IN-RADIUS",
		"insert-item":                   "INSERT-ITEM",
		"inspect":                       "INSPECT",
		"int":                           "INT",
		"is-agent?":                     "IS-AGENT",
		"is-agentset?":                  "IS-AGENTSET",
		"is-anonymous-command?":         "IS-ANONYMOUS-COMMAND",
		"is-anonymous-reporter?":        "IS-ANONYMOUS-REPORTER",
		"is-boolean?":                   "IS-BOOLEAN",
		"is-directed-link?":             "IS-DIRECTED-LINK",
		"is-link?":                      "IS-LINK",
		"is-linkset?":                   "IS-LINKSET",
		"is-list?":                      "IS-LIST",
		"is-number?":                    "IS-NUMBER",
		"is-patch?":                     "IS-PATCH",
		"is-patchset?":                  "IS-PATCHSET",
		"is-string?":                    "IS-STRING",
		"is-turtle?":                    "IS-TURTLE",
		"is-turtleset?":                 "IS-TURTLESET",
		"is-undirected-link?":           "IS-UNDIRECTED-LINK",
		"item":                          "ITEM",
		"jump":                          "JUMP",
		"label":                         "LABEL",
		"label-color":                   "LABEL-COLOR",
		"last":                          "LAST",
		"layout-circle":                 "LAYOUT-CIRCLE",
		"layout-radial":                 "LAYOUT-RADIAL",
		"layout-spring":                 "LAYOUT-SPRING",
		"layout-tuttet":                 "LAYOUT-TUTTET",
		"left":                          "LEFT",
		"lt":                            "LEFT",
		"length":                        "LENGTH",
		"let":                           "LET",
		"link":                          "LINK",
		"link-heading":                  "LINK-HEADING",
		"link-length":                   "LINK-LENGTH",
		"link-set":                      "LINK-SET",
		"link-shapes":                   "LINK-SHAPES",
		"links":                         "LINKS",
		"links-own":                     "LINKS-OWN",
		"list":                          "LIST",
		"ln":                            "LN",
		"log":                           "LOG",
		"loop":                          "LOOP",
		"lput":                          "LPUT",
		"map":                           "MAP",
		"max":                           "MAX",
		"max-n-of":                      "MAX-N-OF",
		"max-one-of":                    "MAX-ONE-OF",
		"max-pxcor":                     "MAX-PXCOR",
		"max-pycor":                     "MAX-PYCOR",
		"mean":                          "MEAN",
		"median":                        "MEDIAN",
		"member?":                       "MEMBER",
		"min":                           "MIN",
		"min-n-of":                      "MIN-N-OF",
		"min-one-of":                    "MIN-ONE-OF",
		"min-pxcor":                     "MIN-PXCOR",
		"min-pycor":                     "MIN-PYCOR",
		"mod":                           "MOD",
		"modes":                         "MODES",
		"mouse-down?":                   "MOUSE-DOWN",
		"mouse-inside?":                 "MOUSE-INSIDE",
		"mouse-xcor":                    "MOUSE-XCOR",
		"mouse-ycor":                    "MOUSE-YCOR",
		"move-to":                       "MOVE-TO",
		"my-links":                      "MY-LINKS",
		"my-in-links":                   "MY-IN-LINKS",
		"my-out-links":                  "MY-OUT-LINKS",
		"myself":                        "MYSELF",
		"n-of":                          "N-OF",
		"n-values":                      "N-VALUES",
		"neighbors":                     "NEIGHBORS",
		"neighbors4":                    "NEIGHBORS4",
		"link-neighbor?":                "LINK-NEIGHBOR",
		"new-seed":                      "NEW-SEED",
		"no-display":                    "NO-DISPLAY",
		"nobody":                        "NOBODY",
		"no-links":                      "NO-LINKS",
		"no-patches":                    "NO-PATCHES",
		"not":                           "NOT",
		"no-turtles":                    "NO-TURTLES",
		"of":                            "OF",
		"one-of":                        "ONE-OF",
		"or":                            "OR",
		"other":                         "OTHER",
		"other-end":                     "OTHER-END",
		"out-link-neighbor?":            "OUT-LINK-NEIGHBOR",
		"out-link-neighbors":            "OUT-LINK-NEIGHBORS",
		"out-link-to":                   "OUT-LINK-TO",
		"output-print":                  "OUTPUT-PRINT",
		"output-show":                   "OUTPUT-SHOW",
		"output-type":                   "OUTPUT-TYPE",
		"output-write":                  "OUTPUT-WRITE",
		"patch":                         "PATCH",
		"patch-ahead":                   "PATCH-AHEAD",
		"patch-at":                      "PATCH-AT",
		"patch-at-heading-and-distance": "PATCH-AT-HEADING-AND-DISTANCE",
		"patch-here":                    "PATCH-HERE",
		"patch-left-and-ahead":          "PATCH-LEFT-AND-AHEAD",
		"patch-right-and-ahead":         "PATCH-RIGHT-AND-AHEAD",
		"patch-set":                     "PATCH-SET",
		"patch-size":                    "PATCH-SIZE",
		"patches":                       "PATCHES",
		"patches-own":                   "PATCHES-OWN",
		"pcolor":                        "PCOLOR",
		"plabel":                        "PLABEL",
		"plabel-color":                  "PLABEL-COLOR",
		"position":                      "POSITION",
		"precision":                     "PRECISION",
		"print":                         "PRINT",
		"pxcor":                         "PXCOR",
		"pycor":                         "PYCOR",
		"random":                        "RANDOM",
		"random-float":                  "RANDOM-FLOAT",
		"random-exponential":            "RANDOM-EXPONENTIAL",
		"random-gamma":                  "RANDOM-GAMMA",
		"random-normal":                 "RANDOM-NORMAL",
		"random-poisson":                "RANDOM-POISSON",
		"random-pxcor":                  "RANDOM-PXCOR",
		"random-pycor":                  "RANDOM-PYCOR",
		"random-seed":                   "RANDOM-SEED",
		"random-xcor":                   "RANDOM-XCOR",
		"random-ycor":                   "RANDOM-YCOR",
		"range":                         "RANGE",
		"read-from-string":              "READ-FROM-STRING",
		"reduce":                        "REDUCE",
		"remainder":                     "REMAINDER",
		"remove":                        "REMOVE",
		"remove-duplicates":             "REMOVE-DUPLICATES",
		"remove-item":                   "REMOVE-ITEM",
		"repeat":                        "REPEAT",
		"replace-item":                  "REPLACE-ITEM",
		"report":                        "REPORT",
		"reset-perspective":             "RESET-PERSPECTIVE",
		"rp":                            "RESET-PERSPECTIVE",
		"reset-ticks":                   "RESET-TICKS",
		"reset-timer":                   "RESET-TIMER",
		"resize-world":                  "RESIZE-WORLD",
		"reverse":                       "REVERSE",
		"rgb":                           "RGB",
		"ride":                          "RIDE",
		"ride-me":                       "RIDE-ME",
		"right":                         "RIGHT",
		"rt":                            "RIGHT",
		"round":                         "ROUND",
		"scale-color":                   "SCALE-COLOR",
		"self":                          "SELF",
		";":                             "SEMICOLON",
		"sentence":                      "SENTENCE",
		"se":                            "SENTENCE",
		"set":                           "SET",
		"set-default-shape":             "SET-DEFAULT-SHAPE",
		"set-histogram-num-bars":        "SET-HISTOGRAM-NUM-BARS",
		"set-patch-size":                "SET-PATCH-SIZE",
		"set-plot-background-color":     "SET-PLOT-BACKGROUND-COLOR",
		"set-plot-pen-color":            "SET-PLOT-PEN-COLOR",
		"set-plot-pen-interval":         "SET-PLOT-PEN-INTERVAL",
		"set-plot-pen-mode":             "SET-PLOT-PEN-MODE",
		"setup-plots":                   "SETUP-PLOTS",
		"set-plot-x-range":              "SET-PLOT-X-RANGE",
		"set-plot-y-range":              "SET-PLOT-Y-RANGE",
		"set-xy":                        "SET-XY",
		"shade-of?":                     "SHADE-OF",
		"shape":                         "SHAPE",
		"shapes":                        "SHAPES",
		"show":                          "SHOW",
		"show-turtle":                   "SHOW-TURTLE",
		"st":                            "SHOW-TURTLE",
		"show-link":                     "SHOW-LINK",
		"shuffle":                       "SHUFFLE",
		"sin":                           "SIN",
		"size":                          "SIZE",
		"sort":                          "SORT",
		"sort-by":                       "SORT-BY",
		"sort-on":                       "SORT-ON",
		"sprout":                        "SPROUT",
		"sqrt":                          "SQRT",
		"stamp":                         "STAMP",
		"stamp-erase":                   "STAMP-ERASE",
		"standard-deviation":            "STANDARD-DEVIATION",
		"startup":                       "STARTUP",
		"stop":                          "STOP",
		"subject":                       "SUBJECT",
		"sublist":                       "SUBLIST",
		"substring":                     "SUBSTRING",
		"subtract-headings":             "SUBTRACT-HEADINGS",
		"sum":                           "SUM",
		"tan":                           "TAN",
		"thickness":                     "THICKNESS",
		"tick":                          "TICK",
		"tick-advance":                  "TICK-ADVANCE",
		"ticks":                         "TICKS",
		"tie":                           "TIE",
		"tie-mode":                      "TIE-MODE",
		"timer":                         "TIMER",
		"to":                            "TO",
		"to-report":                     "TO-REPORT",
		"towards":                       "TOWARDS",
		"towardsxy":                     "TOWARDSXY",
		"turtle":                        "TURTLE",
		"turtle-set":                    "TURTLE-SET",
		"turtles":                       "TURTLES",
		"turtles-at":                    "TURTLES-AT",
		"turtles-here":                  "TURTLES-HERE",
		"turtles-on":                    "TURTLES-ON",
		"turtles-own":                   "TURTLES-OWN",
		"type":                          "TYPE",
		"undirected-link-breed":         "UNDIRECTED-LINK-BREED",
		"untie":                         "UNTIE",
		"up-to-n-of":                    "UP-TO-N-OF",
		"update-plots":                  "UPDATE-PLOTS",
		"uphill":                        "UPHILL",
		"uphill4":                       "UPHILL4",
		"user-directory":                "USER-DIRECTORY",
		"user-file":                     "USER-FILE",
		"user-new-file":                 "USER-NEW-FILE",
		"user-input":                    "USER-INPUT",
		"user-message":                  "USER-MESSAGE",
		"user-one-of":                   "USER-ONE-OF",
		"user-yes-or-no?":               "USER-YES-OR-NO",
		"variance":                      "VARIANCE",
		"wait":                          "WAIT",
		"while":                         "WHILE",
		"who":                           "WHO",
		"who-are-not":                   "WHO-ARE-NOT",
		"with":                          "WITH",
		"link-with":                     "LINK-WITH",
		"with-max":                      "WITH-MAX",
		"with-min":                      "WITH-MIN",
		"with-local-randomness":         "WITH-LOCAL-RANDOMNESS",
		"without-interruption":          "WITHOUT-INTERRUPTION",
		"word":                          "WORD",
		"world-width":                   "WORLD-WIDTH",
		"world-height":                  "WORLD-HEIGHT",
		"wrap-color":                    "WRAP-COLOR",
		"write":                         "WRITE",
		"xcor":                          "XCOR",
		"xor":                           "XOR",
		"ycor":                          "YCOR",
		"->":                            "ARROW",
	}

	notImplementedTokens = map[string]string{
		"ask-concurrent":                "ASK-CONCURRENT",
		"autoplot?":                     "AUTOPLOT",
		"auto-plot-on":                  "AUTO-PLOT-ON",
		"auto-plot-off":                 "AUTO-PLOT-OFF",
		"beep":                          "BEEP",
		"behaviorspace-experiment-name": "BEHAVIORSPACE-EXPERIMENT-NAME",
		"behaviorspace-run-number":      "BEHAVIORSPACE-RUN-NUMBER",
		"ca":                            "CLEAR-ALL",
		"carefully":                     "CAREFULLY",
		"create-temporary-plot-pen":     "CREATE-TEMPORARY-PLOT-PEN",
		"export-view":                   "EXPORT-VIEW",
		"export-interface":              "EXPORT-INTERFACE",
		"export-output":                 "EXPORT-OUTPUT",
		"export-plot":                   "EXPORT-PLOT",
		"export-all-plots":              "EXPORT-ALL-PLOTS",
		"export-world":                  "EXPORT-WORLD",
		"file-at-end?":                  "FILE-AT-END",
		"file-close":                    "FILE-CLOSE",
		"file-close-all":                "FILE-CLOSE-ALL",
		"file-delete":                   "FILE-DELETE",
		"file-exists?":                  "FILE-EXISTS",
		"file-flush":                    "FILE-FLUSH",
		"file-open":                     "FILE-OPEN",
		"file-print":                    "FILE-PRINT",
		"file-read":                     "FILE-READ",
		"file-read-characters":          "FILE-READ-CHARACTERS",
		"file-read-line":                "FILE-READ-LINE",
		"file-show":                     "FILE-SHOW",
		"file-type":                     "FILE-TYPE",
		"file-write":                    "FILE-WRITE",
		"hubnet-broadcast":              "HUBNET-BROADCAST",
		"hubnet-broadcast-clear-output": "HUBNET-BROADCAST-CLEAR-OUTPUT",
		"hubnet-broadcast-message":      "HUBNET-BROADCAST-MESSAGE",
		"hubnet-clear-override":         "HUBNET-BROADCAST-OVERRIDE",
		"hubnet-clear-overrides":        "HUBNET-CLEAR-OVERRIDES",
		"hubnet-clients-list":           "HUBNET-CLIENTS-LIST",
		"hubnet-enter-message":          "HUBNET-ENTER-MESSAGE",
		"hubnet-exit-message":           "HUBNET-EXIT-MESSAGE",
		"hubnet-fetch-message":          "HUBNET-FETCH-MESSAGE",
		"hubnet-kick-all-clients":       "HUBNET-KICK-ALL-CLIENTS",
		"hubnet-message":                "HUBNET-MESSAGE",
		"hubnet-message-source":         "HUBNET-MESSAGE-SOURCE",
		"hubnet-message-tag":            "HUBNET-MESSAGE-TAG",
		"hubnet-message-waiting?":       "HUBNET-MESSAGE-WAITING",
		"hubnet-reset":                  "HUBNET-RESET",
		"hubnet-reset-perspective":      "HUBNET-RESET-PERSPECTIVE",
		"hubnet-send":                   "HUBNET-SEND",
		"hubnet-send-clear-output":      "HUBNET-SEND-CLEAR-OUTPUT",
		"hubnet-send-follow":            "HUBNET-SEND-FOLLOW",
		"hubnet-send-message":           "HUBNET-SEND-MESSAGE",
		"hubnet-send-override":          "HUBNET-SEND-OVERRIDE",
		"hubnet-send-watch":             "HUBNET-SEND-WATCH",
		"__includes":                    "INCLUDES",
		"netlogo-version":               "NETLOGO-VERSION",
		"netlogo-web?":                  "NETLOGO-WEB",
		"pen-down":                      "PEN-DOWN",
		"pd":                            "PEN-DOWN",
		"pen-erase":                     "PEN-ERASE",
		"pe":                            "PEN-ERASE",
		"pen-up":                        "PEN-UP",
		"pu":                            "PEN-UP",
		"pen-mode":                      "PEN-MODE",
		"pen-size":                      "PEN-SIZE",
		"plot":                          "PLOT",
		"plot-name":                     "PLOT-NAME",
		"plot-pen-exists?":              "PLOT-PEN-EXISTS",
		"plot-pen-down":                 "PLOT-PEN-DOWN",
		"plot-pen-up":                   "PLOT-PEN-UP",
		"plot-pen-reset":                "PLOT-PEN-RESET",
		"plotxy":                        "PLOTXY",
		"plot-x-min":                    "PLOT-X-MIN",
		"plot-x-max":                    "PLOT-X-MAX",
		"plot-y-min":                    "PLOT-Y-MIN",
		"plot-y-max":                    "PLOT-Y-MAX",
		"run":                           "RUN",
		"runresult":                     "RUNRESULT",
		"set-current-directory":         "SET-CURRENT-DIRECTORY",
		"set-current-plot":              "SET-CURRENT-PLOT",
		"set-current-plot-pen":          "SET-CURRENT-PLOT-PEN",
		"__set-line-thickness":          "SET-LINE-THICKNESS",
		"stop-inspecting":               "STOP-INSPECTING",
		"stop-inspecting-dead-agents":   "STOP-INSPECTING-DEAD-AGENTS",
		"watch":                         "WATCH",
		"watch-me":                      "WATCH-ME",
	}
)