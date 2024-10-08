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
		tCREATE_ORDERED_BREEDS: nil,
		tCREATE_BREED_TO:       nil,
		tCREATE_BREEDS_TO:      nil,
		tCREATE_BREED_FROM:     nil,
		tCREATE_BREEDS_FROM:    nil,
		tCREATE_BREED_WITH:     nil,
		tCREATE_BREEDS_WITH:    nil,
		tCREATE_BREEDS:         nil,
		tIN_BREED_NEIGHBOR:     nil,
		tIN_BREED_NEIGHBORS:    nil,
		tIN_BREED_FROM:         nil,
		tMY_BREEDS:             nil,
		tMY_IN_BREEDS:          nil,
		tMY_OUT_BREEDS:         nil,
		tBREED_NEIGHBOR:        nil,
		tOUT_BREED_NEIGHBOR:    nil,
		tOUT_BREED_NEIGHBORS:   nil,
		tOUT_BREED_TO:          nil,
		tBREEDS_AT:             nil,
		tBREEDS_HERE:           nil,
		tBREEDS_ON:             nil,
		tBREEDS_OWN:            nil,
		tBREED_WITH:            nil,
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
		tNUMBER:  nil,
		tSTRING:  nil,
		tBOOLEAN: nil,
		tVAR:     nil,
		tNEWLINE: nil,
	}

	colorTokens = map[string]interface{}{
		tBLACK:     nil,
		tGRAY:      nil,
		tWHITE:     nil,
		tRED:       nil,
		tORANGE:    nil,
		tBROWN:     nil,
		tYELLOW:    nil,
		tGREEN:     nil,
		tLIME:      nil,
		tTURQUOISE: nil,
		tCYAN:      nil,
		tSKY:       nil,
		tBLUE:      nil,
		tVIOLET:    nil,
		tMAGENTA:   nil,
		tPINK:      nil,
	}

	// Keywords
	keywordToTokenType = map[string]string{
		"e":                             tE,
		"pi":                            tPI,
		"true":                          tTRUE,
		"false":                         tFALSE,
		"black":                         tBLACK,
		"gray":                          tGRAY,
		"white":                         tWHITE,
		"red":                           tRED,
		"orange":                        tORANGE,
		"brown":                         tBROWN,
		"yellow":                        tYELLOW,
		"green":                         tGREEN,
		"lime":                          tLIME,
		"turquoise":                     tTURQUOISE,
		"cyan":                          tCYAN,
		"sky":                           tSKY,
		"blue":                          tBLUE,
		"violet":                        tVIOLET,
		"magenta":                       tMAGENTA,
		"pink":                          tPINK,
		"abs":                           tABS,
		"acos":                          tACOS,
		"all?":                          tALL,
		"and":                           tAND,
		"any?":                          tANY,
		"approximate-hsb":               tAPPROXIMATE_HSB,
		"approximate-rgb":               tAPPROXIMATE_RGB,
		"+":                             tPLUS,
		"*":                             tMULTIPLY,
		"-":                             tMINUS,
		"/":                             tDIVIDE,
		"^":                             tPOWER,
		"<":                             tLESS,
		">":                             tGREATER,
		"<=":                            tLESS_EQUAL,
		">=":                            tGREATER_EQUAL,
		"= ":                            tEQUAL,
		"!= ":                           tNOT_EQUAL,
		"[":                             tOPEN_BRACKET,
		"]":                             tCLOSE_BRACKET,
		"(":                             tOPEN_PARENTHESIS,
		")":                             tCLOSE_PARENTHESIS,
		"asin":                          tASIN,
		"ask":                           tASK,
		"at-points":                     tAT_POINTS,
		"atan":                          tATAN,
		"back":                          tBACK,
		"bk":                            tBACK,
		"base-colors":                   tBASE_COLORS,
		"both-ends":                     tBOTH_ENDS,
		"breed":                         tBREED,
		"but-first":                     tBUT_FIRST,
		"butfirst":                      tBUT_FIRST,
		"bf":                            tBUT_FIRST,
		"but-last":                      tBUT_LAST,
		"butlast":                       tBUT_LAST,
		"bl":                            tBUT_LAST,
		"can-move?":                     tCAN_MOVE,
		"ceiling":                       tCEILING,
		"clear-all":                     tCLEAR_ALL,
		"clear-all-plots":               tCLEAR_ALL_PLOTS,
		"clear-drawing":                 tCLEAR_DRAWING,
		"cd":                            tCLEAR_DRAWING,
		"clear-globals":                 tCLEAR_GLOBALS,
		"clear-links":                   tCLEAR_LINKS,
		"clear-output":                  tCLEAR_OUTPUT,
		"clear-patches":                 tCLEAR_PATCHES,
		"cp":                            tCLEAR_PATCHES,
		"clear-plot":                    tCLEAR_PLOT,
		"clear-plots":                   tCLEAR_PLOTS,
		"clear-ticks":                   tCLEAR_TICKS,
		"ct":                            tCLEAR_TICKS,
		"color":                         tCOLOR,
		"cos":                           tCOS,
		"count":                         tCOUNT,
		"create-ordered-turtles":        tCREATE_ORDERED_TURTLES,
		"cro":                           tCREATE_ORDERED_TURTLES,
		"create-link-to":                tCREATE_LINK_TO,
		"create-links-to":               tCREATE_LINKS_TO,
		"create-link-from":              tCREATE_LINK_FROM,
		"create-links-from":             tCREATE_LINKS_FROM,
		"create-link-with":              tCREATE_LINK_WITH,
		"create-links-with":             tCREATE_LINKS_WITH,
		"create-turtles":                tCREATE_TURTLES,
		"crt":                           tCREATE_TURTLES,
		"date-and-time":                 tDATE_AND_TIME,
		"die":                           tDIE,
		"diffuse":                       tDIFFUSE,
		"diffuse4":                      tDIFFUSE4,
		"directed-link-breed":           tDIRECTED_LINK_BREED,
		"display":                       tDISPLAY,
		"distance":                      tDISTANCE,
		"distancexy":                    tDISTANCEXY,
		"downhill":                      tDOWNHILL,
		"downhill4":                     tDOWNHILL4,
		"dx":                            tDX,
		"dy":                            tDY,
		"empty?":                        tEMPTY,
		"end":                           tEND,
		"end1":                          tEND1,
		"end2":                          tEND2,
		"error":                         tERROR,
		"error-message":                 tERROR_MESSAGE,
		"every":                         tEVERY,
		"exp":                           tEXP,
		"extensions":                    tEXTENSIONS,
		"extract-hsb":                   tEXTRACT_HSB,
		"extract-rgb":                   tEXTRACT_RGB,
		"face":                          tFACE,
		"facexy":                        tFACEXY,
		"filter":                        tFILTER,
		"first":                         tFIRST,
		"floor":                         tFLOOR,
		"follow":                        tFOLLOW,
		"follow-me":                     tFOLLOW_ME,
		"foreach":                       tFOREACH,
		"forward":                       tFORWARD,
		"fd":                            tFORWARD,
		"fput":                          tFPUT,
		"globals":                       tGLOBALS,
		"hatch":                         tHATCH,
		"heading":                       tHEADING,
		"hidden?":                       tHIDDEN,
		"hide-link":                     tHIDE_LINK,
		"hide-turtle":                   tHIDE_TURTLE,
		"ht":                            tHIDE_TURTLE,
		"histogram":                     tHISTOGRAM,
		"home":                          tHOME,
		"hsb":                           tHSB,
		"if":                            tIF,
		"ifelse":                        tIFELSE,
		"ifelse-value":                  tIFELSE_VALUE,
		"import-drawing":                tIMPORT_DRAWING,
		"import-pcolors":                tIMPORT_PCOLORS,
		"import-pcolors-rgb":            tIMPORT_PCOLORS_RGB,
		"import-world":                  tIMPORT_WORLD,
		"in-cone":                       tIN_CONE,
		"in-link-neighbor?":             tIN_LINK_NEIGHBOR,
		"in-link-neighbors":             tIN_LINK_NEIGHBORS,
		"in-link-from":                  tIN_LINK_FROM,
		"in-radius":                     tIN_RADIUS,
		"insert-item":                   tINSERT_ITEM,
		"inspect":                       tINSPECT,
		"int":                           tINT,
		"is-agent?":                     tIS_AGENT,
		"is-agentset?":                  tIS_AGENTSET,
		"is-anonymous-command?":         tIS_ANONYMOUS_COMMAND,
		"is-anonymous-reporter?":        tIS_ANONYMOUS_REPORTER,
		"is-boolean?":                   tIS_BOOLEAN,
		"is-directed-link?":             tIS_DIRECTED_LINK,
		"is-link?":                      tIS_LINK,
		"is-linkset?":                   tIS_LINKSET,
		"is-list?":                      tIS_LIST,
		"is-number?":                    tIS_NUMBER,
		"is-patch?":                     tIS_PATCH,
		"is-patchset?":                  tIS_PATCHSET,
		"is-string?":                    tIS_STRING,
		"is-turtle?":                    tIS_TURTLE,
		"is-turtleset?":                 tIS_TURTLESET,
		"is-undirected-link?":           tIS_UNDIRECTED_LINK,
		"item":                          tITEM,
		"jump":                          tJUMP,
		"label":                         tLABEL,
		"label-color":                   tLABEL_COLOR,
		"last":                          tLAST,
		"layout-circle":                 tLAYOUT_CIRCLE,
		"layout-radial":                 tLAYOUT_RADIAL,
		"layout-spring":                 tLAYOUT_SPRING,
		"layout-tuttet":                 tLAYOUT_TUTTET,
		"left":                          tLEFT,
		"lt":                            tLEFT,
		"length":                        tLENGTH,
		"let":                           tLET,
		"link":                          tLINK,
		"link-heading":                  tLINK_HEADING,
		"link-length":                   tLINK_LENGTH,
		"link-set":                      tLINK_SET,
		"link-shapes":                   tLINK_SHAPES,
		"links":                         tLINKS,
		"links-own":                     tLINKS_OWN,
		"list":                          tLIST,
		"ln":                            tLN,
		"log":                           tLOG,
		"loop":                          tLOOP,
		"lput":                          tLPUT,
		"map":                           tMAP,
		"max":                           tMAX,
		"max-n-of":                      tMAX_N_OF,
		"max-one-of":                    tMAX_ONE_OF,
		"max-pxcor":                     tMAX_PXCOR,
		"max-pycor":                     tMAX_PYCOR,
		"mean":                          tMEAN,
		"median":                        tMEDIAN,
		"member?":                       tMEMBER,
		"min":                           tMIN,
		"min-n-of":                      tMIN_N_OF,
		"min-one-of":                    tMIN_ONE_OF,
		"min-pxcor":                     tMIN_PXCOR,
		"min-pycor":                     tMIN_PYCOR,
		"mod":                           tMOD,
		"modes":                         tMODES,
		"mouse-down?":                   tMOUSE_DOWN,
		"mouse-inside?":                 tMOUSE_INSIDE,
		"mouse-xcor":                    tMOUSE_XCOR,
		"mouse-ycor":                    tMOUSE_YCOR,
		"move-to":                       tMOVE_TO,
		"my-links":                      tMY_LINKS,
		"my-in-links":                   tMY_IN_LINKS,
		"my-out-links":                  tMY_OUT_LINKS,
		"myself":                        tMYSELF,
		"n-of":                          tN_OF,
		"n-values":                      tN_VALUES,
		"neighbors":                     tNEIGHBORS,
		"neighbors4":                    tNEIGHBORS4,
		"link-neighbor?":                tLINK_NEIGHBOR,
		"new-seed":                      tNEW_SEED,
		"no-display":                    tNO_DISPLAY,
		"nobody":                        tNOBODY,
		"no-links":                      tNO_LINKS,
		"no-patches":                    tNO_PATCHES,
		"not":                           tNOT,
		"no-turtles":                    tNO_TURTLES,
		"of":                            tOF,
		"one-of":                        tONE_OF,
		"or":                            tOR,
		"other":                         tOTHER,
		"other-end":                     tOTHER_END,
		"out-link-neighbor?":            tOUT_LINK_NEIGHBOR,
		"out-link-neighbors":            tOUT_LINK_NEIGHBORS,
		"out-link-to":                   tOUT_LINK_TO,
		"output-print":                  tOUTPUT_PRINT,
		"output-show":                   tOUTPUT_SHOW,
		"output-type":                   tOUTPUT_TYPE,
		"output-write":                  tOUTPUT_WRITE,
		"patch":                         tPATCH,
		"patch-ahead":                   tPATCH_AHEAD,
		"patch-at":                      tPATCH_AT,
		"patch-at-heading-and-distance": tPATCH_AT_HEADING_AND_DISTANCE,
		"patch-here":                    tPATCH_HERE,
		"patch-left-and-ahead":          tPATCH_LEFT_AND_AHEAD,
		"patch-right-and-ahead":         tPATCH_RIGHT_AND_AHEAD,
		"patch-set":                     tPATCH_SET,
		"patch-size":                    tPATCH_SIZE,
		"patches":                       tPATCHES,
		"patches-own":                   tPATCHES_OWN,
		"pcolor":                        tPCOLOR,
		"plabel":                        tPLABEL,
		"plabel-color":                  tPLABEL_COLOR,
		"position":                      tPOSITION,
		"precision":                     tPRECISION,
		"print":                         tPRINT,
		"pxcor":                         tPXCOR,
		"pycor":                         tPYCOR,
		"random":                        tRANDOM,
		"random-float":                  tRANDOM_FLOAT,
		"random-exponential":            tRANDOM_EXPONENTIAL,
		"random-gamma":                  tRANDOM_GAMMA,
		"random-normal":                 tRANDOM_NORMAL,
		"random-poisson":                tRANDOM_POISSON,
		"random-pxcor":                  tRANDOM_PXCOR,
		"random-pycor":                  tRANDOM_PYCOR,
		"random-seed":                   tRANDOM_SEED,
		"random-xcor":                   tRANDOM_XCOR,
		"random-ycor":                   tRANDOM_YCOR,
		"range":                         tRANGE,
		"read-from-string":              tREAD_FROM_STRING,
		"reduce":                        tREDUCE,
		"remainder":                     tREMAINDER,
		"remove":                        tREMOVE,
		"remove-duplicates":             tREMOVE_DUPLICATES,
		"remove-item":                   tREMOVE_ITEM,
		"repeat":                        tREPEAT,
		"replace-item":                  tREPLACE_ITEM,
		"report":                        tREPORT,
		"reset-perspective":             tRESET_PERSPECTIVE,
		"rp":                            tRESET_PERSPECTIVE,
		"reset-ticks":                   tRESET_TICKS,
		"reset-timer":                   tRESET_TIMER,
		"resize-world":                  tRESIZE_WORLD,
		"reverse":                       tREVERSE,
		"rgb":                           tRGB,
		"ride":                          tRIDE,
		"ride-me":                       tRIDE_ME,
		"right":                         tRIGHT,
		"rt":                            tRIGHT,
		"round":                         tROUND,
		"scale-color":                   tSCALE_COLOR,
		"self":                          tSELF,
		";":                             tSEMICOLON,
		"sentence":                      tSENTENCE,
		"se":                            tSENTENCE,
		"set":                           tSET,
		"set-default-shape":             tSET_DEFAULT_SHAPE,
		"set-histogram-num-bars":        tSET_HISTOGRAM_NUM_BARS,
		"set-patch-size":                tSET_PATCH_SIZE,
		"set-plot-background-color":     tSET_PLOT_BACKGROUND_COLOR,
		"set-plot-pen-color":            tSET_PLOT_PEN_COLOR,
		"set-plot-pen-interval":         tSET_PLOT_PEN_INTERVAL,
		"set-plot-pen-mode":             tSET_PLOT_PEN_MODE,
		"setup-plots":                   tSETUP_PLOTS,
		"set-plot-x-range":              tSET_PLOT_X_RANGE,
		"set-plot-y-range":              tSET_PLOT_Y_RANGE,
		"setxy":                         tSET_XY,
		"shade-of?":                     tSHADE_OF,
		"shape":                         tSHAPE,
		"shapes":                        tSHAPES,
		"show":                          tSHOW,
		"show-turtle":                   tSHOW_TURTLE,
		"st":                            tSHOW_TURTLE,
		"show-link":                     tSHOW_LINK,
		"shuffle":                       tSHUFFLE,
		"sin":                           tSIN,
		"size":                          tSIZE,
		"sort":                          tSORT,
		"sort-by":                       tSORT_BY,
		"sort-on":                       tSORT_ON,
		"sprout":                        tSPROUT,
		"sqrt":                          tSQRT,
		"stamp":                         tSTAMP,
		"stamp-erase":                   tSTAMP_ERASE,
		"standard-deviation":            tSTANDARD_DEVIATION,
		"startup":                       tSTARTUP,
		"stop":                          tSTOP,
		"subject":                       tSUBJECT,
		"sublist":                       tSUBLIST,
		"substring":                     tSUBSTRING,
		"subtract-headings":             tSUBTRACT_HEADINGS,
		"sum":                           tSUM,
		"tan":                           tTAN,
		"thickness":                     tTHICKNESS,
		"tick":                          tTICK,
		"tick-advance":                  tTICK_ADVANCE,
		"ticks":                         tTICKS,
		"tie":                           tTIE,
		"tie-mode":                      tTIE_MODE,
		"timer":                         tTIMER,
		"to":                            tTO,
		"to-report":                     tTO_REPORT,
		"towards":                       tTOWARDS,
		"towardsxy":                     tTOWARDSXY,
		"turtle":                        tTURTLE,
		"turtle-set":                    tTURTLE_SET,
		"turtles":                       tTURTLES,
		"turtles-at":                    tTURTLES_AT,
		"turtles-here":                  tTURTLES_HERE,
		"turtles-on":                    tTURTLES_ON,
		"turtles-own":                   tTURTLES_OWN,
		"type":                          tTYPE,
		"undirected-link-breed":         tUNDIRECTED_LINK_BREED,
		"untie":                         tUNTIE,
		"up-to-n-of":                    tUP_TO_N_OF,
		"update-plots":                  tUPDATE_PLOTS,
		"uphill":                        tUPHILL,
		"uphill4":                       tUPHILL4,
		"user-directory":                tUSER_DIRECTORY,
		"user-file":                     tUSER_FILE,
		"user-new-file":                 tUSER_NEW_FILE,
		"user-input":                    tUSER_INPUT,
		"user-message":                  tUSER_MESSAGE,
		"user-one-of":                   tUSER_ONE_OF,
		"user-yes-or-no?":               tUSER_YES_OR_NO,
		"variance":                      tVARIANCE,
		"wait":                          tWAIT,
		"while":                         tWHILE,
		"who":                           tWHO,
		"who-are-not":                   tWHO_ARE_NOT,
		"with":                          tWITH,
		"link-with":                     tLINK_WITH,
		"with-max":                      tWITH_MAX,
		"with-min":                      tWITH_MIN,
		"with-local-randomness":         tWITH_LOCAL_RANDOMNESS,
		"without-interruption":          tWITHOUT_INTERRUPTION,
		"word":                          tWORD,
		"world-width":                   tWORLD_WIDTH,
		"world-height":                  tWORLD_HEIGHT,
		"wrap-color":                    tWRAP_COLOR,
		"write":                         tWRITE,
		"xcor":                          tXCOR,
		"xor":                           tXOR,
		"ycor":                          tYCOR,
		"->":                            tARROW,
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
