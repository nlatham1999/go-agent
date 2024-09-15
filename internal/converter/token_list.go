package converter

var (
	tE                             = "E"
	tPI                            = "PI"
	tTRUE                          = "TRUE"
	tFALSE                         = "FALSE"
	tBLACK                         = "BLACK"
	tGRAY                          = "GRAY"
	tWHITE                         = "WHITE"
	tRED                           = "RED"
	tORANGE                        = "ORANGE"
	tBROWN                         = "BROWN"
	tYELLOW                        = "YELLOW"
	tGREEN                         = "GREEN"
	tLIME                          = "LIME"
	tTURQUOISE                     = "TURQUOISE"
	tCYAN                          = "CYAN"
	tSKY                           = "SKY"
	tBLUE                          = "BLUE"
	tVIOLET                        = "VIOLET"
	tMAGENTA                       = "MAGENTA"
	tPINK                          = "PINK"
	tABS                           = "ABS"
	tACOS                          = "ACOS"
	tALL                           = "ALL"
	tAND                           = "AND"
	tANY                           = "ANY"
	tAPPROXIMATE_HSB               = "APPROXIMATE-HSB"
	tAPPROXIMATE_RGB               = "APPROXIMATE-RGB"
	tPLUS                          = "PLUS"
	tMULTIPLY                      = "MULTIPLY"
	tMINUS                         = "MINUS"
	tDIVIDE                        = "DIVIDE"
	tPOWER                         = "POWER"
	tLESS                          = "LESS"
	tGREATER                       = "GREATER"
	tLESS_EQUAL                    = "LESS-EQUAL"
	tGREATER_EQUAL                 = "GREATER-EQUAL"
	tEQUAL                         = "EQUAL"
	tNOT_EQUAL                     = "NOT-EQUAL"
	tOPEN_BRACKET                  = "OPEN-BRACKET"
	tCLOSE_BRACKET                 = "CLOSE-BRACKET"
	tOPEN_PARENTHESIS              = "OPEN-PARENTHESIS"
	tCLOSE_PARENTHESIS             = "CLOSE-PARENTHESIS"
	tASIN                          = "ASIN"
	tASK                           = "ASK"
	tAT_POINTS                     = "AT-POINTS"
	tATAN                          = "ATAN"
	tBACK                          = "BACK"
	tBASE_COLORS                   = "BASE-COLORS"
	tBOTH_ENDS                     = "BOTH-ENDS"
	tBREED                         = "BREED"
	tBUT_FIRST                     = "BUT-FIRST"
	tBUT_LAST                      = "BUT-LAST"
	tCAN_MOVE                      = "CAN-MOVE"
	tCEILING                       = "CEILING"
	tCLEAR_ALL                     = "CLEAR-ALL"
	tCLEAR_ALL_PLOTS               = "CLEAR-ALL-PLOTS"
	tCLEAR_DRAWING                 = "CLEAR-DRAWING"
	tCLEAR_GLOBALS                 = "CLEAR-GLOBALS"
	tCLEAR_LINKS                   = "CLEAR-LINKS"
	tCLEAR_OUTPUT                  = "CLEAR-OUTPUT"
	tCLEAR_PATCHES                 = "CLEAR-PATCHES"
	tCLEAR_PLOT                    = "CLEAR-PLOT"
	tCLEAR_PLOTS                   = "CLEAR-PLOTS"
	tCLEAR_TICKS                   = "CLEAR-TICKS"
	tCOLOR                         = "COLOR"
	tCOS                           = "COS"
	tCOUNT                         = "COUNT"
	tCREATE_ORDERED_TURTLES        = "CREATE-ORDERED-TURTLES"
	tCREATE_LINK_TO                = "CREATE-LINK-TO"
	tCREATE_LINKS_TO               = "CREATE-LINKS-TO"
	tCREATE_LINK_FROM              = "CREATE-LINK-FROM"
	tCREATE_LINKS_FROM             = "CREATE-LINKS-FROM"
	tCREATE_LINK_WITH              = "CREATE-LINK-WITH"
	tCREATE_LINKS_WITH             = "CREATE-LINKS-WITH"
	tCREATE_TURTLES                = "CREATE-TURTLES"
	tDATE_AND_TIME                 = "DATE-AND-TIME"
	tDIE                           = "DIE"
	tDIFFUSE                       = "DIFFUSE"
	tDIFFUSE4                      = "DIFFUSE4"
	tDIRECTED_LINK_BREED           = "DIRECTED-LINK-BREED"
	tDISPLAY                       = "DISPLAY"
	tDISTANCE                      = "DISTANCE"
	tDISTANCEXY                    = "DISTANCEXY"
	tDOWNHILL                      = "DOWNHILL"
	tDOWNHILL4                     = "DOWNHILL4"
	tDX                            = "DX"
	tDY                            = "DY"
	tEMPTY                         = "EMPTY"
	tEND                           = "END"
	tEND1                          = "END1"
	tEND2                          = "END2"
	tERROR                         = "ERROR"
	tERROR_MESSAGE                 = "ERROR-MESSAGE"
	tEVERY                         = "EVERY"
	tEXP                           = "EXP"
	tEXTENSIONS                    = "EXTENSIONS"
	tEXTRACT_HSB                   = "EXTRACT-HSB"
	tEXTRACT_RGB                   = "EXTRACT-RGB"
	tFACE                          = "FACE"
	tFACEXY                        = "FACEXY"
	tFILTER                        = "FILTER"
	tFIRST                         = "FIRST"
	tFLOOR                         = "FLOOR"
	tFOLLOW                        = "FOLLOW"
	tFOLLOW_ME                     = "FOLLOW-ME"
	tFOREACH                       = "FOREACH"
	tFORWARD                       = "FORWARD"
	tFPUT                          = "FPUT"
	tGLOBALS                       = "GLOBALS"
	tHATCH                         = "HATCH"
	tHEADING                       = "HEADING"
	tHIDDEN                        = "HIDDEN"
	tHIDE_LINK                     = "HIDE-LINK"
	tHIDE_TURTLE                   = "HIDE-TURTLE"
	tHISTOGRAM                     = "HISTOGRAM"
	tHOME                          = "HOME"
	tHSB                           = "HSB"
	tIF                            = "IF"
	tIFELSE                        = "IFELSE"
	tIFELSE_VALUE                  = "IFELSE-VALUE"
	tIMPORT_DRAWING                = "IMPORT-DRAWING"
	tIMPORT_PCOLORS                = "IMPORT-PCOLORS"
	tIMPORT_PCOLORS_RGB            = "IMPORT-PCOLORS-RGB"
	tIMPORT_WORLD                  = "IMPORT-WORLD"
	tIN_CONE                       = "IN-CONE"
	tIN_LINK_NEIGHBOR              = "IN-LINK-NEIGHBOR"
	tIN_LINK_NEIGHBORS             = "IN-LINK-NEIGHBORS"
	tIN_LINK_FROM                  = "IN-LINK-FROM"
	tIN_RADIUS                     = "IN-RADIUS"
	tINSERT_ITEM                   = "INSERT-ITEM"
	tINSPECT                       = "INSPECT"
	tINT                           = "INT"
	tIS_AGENT                      = "IS-AGENT"
	tIS_AGENTSET                   = "IS-AGENTSET"
	tIS_ANONYMOUS_COMMAND          = "IS-ANONYMOUS-COMMAND"
	tIS_ANONYMOUS_REPORTER         = "IS-ANONYMOUS-REPORTER"
	tIS_BOOLEAN                    = "IS-BOOLEAN"
	tIS_DIRECTED_LINK              = "IS-DIRECTED-LINK"
	tIS_LINK                       = "IS-LINK"
	tIS_LINKSET                    = "IS-LINKSET"
	tIS_LIST                       = "IS-LIST"
	tIS_NUMBER                     = "IS-NUMBER"
	tIS_PATCH                      = "IS-PATCH"
	tIS_PATCHSET                   = "IS-PATCHSET"
	tIS_STRING                     = "IS-STRING"
	tIS_TURTLE                     = "IS-TURTLE"
	tIS_TURTLESET                  = "IS-TURTLESET"
	tIS_UNDIRECTED_LINK            = "IS-UNDIRECTED-LINK"
	tITEM                          = "ITEM"
	tJUMP                          = "JUMP"
	tLABEL                         = "LABEL"
	tLABEL_COLOR                   = "LABEL-COLOR"
	tLAST                          = "LAST"
	tLAYOUT_CIRCLE                 = "LAYOUT-CIRCLE"
	tLAYOUT_RADIAL                 = "LAYOUT-RADIAL"
	tLAYOUT_SPRING                 = "LAYOUT-SPRING"
	tLAYOUT_TUTTET                 = "LAYOUT-TUTTET"
	tLEFT                          = "LEFT"
	tLENGTH                        = "LENGTH"
	tLET                           = "LET"
	tLINK                          = "LINK"
	tLINK_HEADING                  = "LINK-HEADING"
	tLINK_LENGTH                   = "LINK-LENGTH"
	tLINK_SET                      = "LINK-SET"
	tLINK_SHAPES                   = "LINK-SHAPES"
	tLINKS                         = "LINKS"
	tLINKS_OWN                     = "LINKS-OWN"
	tLIST                          = "LIST"
	tLN                            = "LN"
	tLOG                           = "LOG"
	tLOOP                          = "LOOP"
	tLPUT                          = "LPUT"
	tMAP                           = "MAP"
	tMAX                           = "MAX"
	tMAX_N_OF                      = "MAX-N-OF"
	tMAX_ONE_OF                    = "MAX-ONE-OF"
	tMAX_PXCOR                     = "MAX-PXCOR"
	tMAX_PYCOR                     = "MAX-PYCOR"
	tMEAN                          = "MEAN"
	tMEDIAN                        = "MEDIAN"
	tMEMBER                        = "MEMBER"
	tMIN                           = "MIN"
	tMIN_N_OF                      = "MIN-N-OF"
	tMIN_ONE_OF                    = "MIN-ONE-OF"
	tMIN_PXCOR                     = "MIN-PXCOR"
	tMIN_PYCOR                     = "MIN-PYCOR"
	tMOD                           = "MOD"
	tMODES                         = "MODES"
	tMOUSE_DOWN                    = "MOUSE-DOWN"
	tMOUSE_INSIDE                  = "MOUSE-INSIDE"
	tMOUSE_XCOR                    = "MOUSE-XCOR"
	tMOUSE_YCOR                    = "MOUSE-YCOR"
	tMOVE_TO                       = "MOVE-TO"
	tMY_LINKS                      = "MY-LINKS"
	tMY_IN_LINKS                   = "MY-IN-LINKS"
	tMY_OUT_LINKS                  = "MY-OUT-LINKS"
	tMYSELF                        = "MYSELF"
	tN_OF                          = "N-OF"
	tN_VALUES                      = "N-VALUES"
	tNEIGHBORS                     = "NEIGHBORS"
	tNEIGHBORS4                    = "NEIGHBORS4"
	tLINK_NEIGHBOR                 = "LINK-NEIGHBOR"
	tNEW_SEED                      = "NEW-SEED"
	tNO_DISPLAY                    = "NO-DISPLAY"
	tNOBODY                        = "NOBODY"
	tNO_LINKS                      = "NO-LINKS"
	tNO_PATCHES                    = "NO-PATCHES"
	tNOT                           = "NOT"
	tNO_TURTLES                    = "NO-TURTLES"
	tOF                            = "OF"
	tONE_OF                        = "ONE-OF"
	tOR                            = "OR"
	tOTHER                         = "OTHER"
	tOTHER_END                     = "OTHER-END"
	tOUT_LINK_NEIGHBOR             = "OUT-LINK-NEIGHBOR"
	tOUT_LINK_NEIGHBORS            = "OUT-LINK-NEIGHBORS"
	tOUT_LINK_TO                   = "OUT-LINK-TO"
	tOUTPUT_PRINT                  = "OUTPUT-PRINT"
	tOUTPUT_SHOW                   = "OUTPUT-SHOW"
	tOUTPUT_TYPE                   = "OUTPUT-TYPE"
	tOUTPUT_WRITE                  = "OUTPUT-WRITE"
	tPATCH                         = "PATCH"
	tPATCH_AHEAD                   = "PATCH-AHEAD"
	tPATCH_AT                      = "PATCH-AT"
	tPATCH_AT_HEADING_AND_DISTANCE = "PATCH-AT-HEADING-AND-DISTANCE"
	tPATCH_HERE                    = "PATCH-HERE"
	tPATCH_LEFT_AND_AHEAD          = "PATCH-LEFT-AND-AHEAD"
	tPATCH_RIGHT_AND_AHEAD         = "PATCH-RIGHT-AND-AHEAD"
	tPATCH_SET                     = "PATCH-SET"
	tPATCH_SIZE                    = "PATCH-SIZE"
	tPATCHES                       = "PATCHES"
	tPATCHES_OWN                   = "PATCHES-OWN"
	tPCOLOR                        = "PCOLOR"
	tPLABEL                        = "PLABEL"
	tPLABEL_COLOR                  = "PLABEL-COLOR"
	tPOSITION                      = "POSITION"
	tPRECISION                     = "PRECISION"
	tPRINT                         = "PRINT"
	tPXCOR                         = "PXCOR"
	tPYCOR                         = "PYCOR"
	tRANDOM                        = "RANDOM"
	tRANDOM_FLOAT                  = "RANDOM-FLOAT"
	tRANDOM_EXPONENTIAL            = "RANDOM-EXPONENTIAL"
	tRANDOM_GAMMA                  = "RANDOM-GAMMA"
	tRANDOM_NORMAL                 = "RANDOM-NORMAL"
	tRANDOM_POISSON                = "RANDOM-POISSON"
	tRANDOM_PXCOR                  = "RANDOM-PXCOR"
	tRANDOM_PYCOR                  = "RANDOM-PYCOR"
	tRANDOM_SEED                   = "RANDOM-SEED"
	tRANDOM_XCOR                   = "RANDOM-XCOR"
	tRANDOM_YCOR                   = "RANDOM-YCOR"
	tRANGE                         = "RANGE"
	tREAD_FROM_STRING              = "READ-FROM-STRING"
	tREDUCE                        = "REDUCE"
	tREMAINDER                     = "REMAINDER"
	tREMOVE                        = "REMOVE"
	tREMOVE_DUPLICATES             = "REMOVE-DUPLICATES"
	tREMOVE_ITEM                   = "REMOVE-ITEM"
	tREPEAT                        = "REPEAT"
	tREPLACE_ITEM                  = "REPLACE-ITEM"
	tREPORT                        = "REPORT"
	tRESET_PERSPECTIVE             = "RESET-PERSPECTIVE"
	tRESET_TICKS                   = "RESET-TICKS"
	tRESET_TIMER                   = "RESET-TIMER"
	tRESIZE_WORLD                  = "RESIZE-WORLD"
	tREVERSE                       = "REVERSE"
	tRGB                           = "RGB"
	tRIDE                          = "RIDE"
	tRIDE_ME                       = "RIDE-ME"
	tRIGHT                         = "RIGHT"
	tROUND                         = "ROUND"
	tSCALE_COLOR                   = "SCALE-COLOR"
	tSELF                          = "SELF"
	tSEMICOLON                     = "SEMICOLON"
	tSENTENCE                      = "SENTENCE"
	tSET                           = "SET"
	tSET_DEFAULT_SHAPE             = "SET-DEFAULT-SHAPE"
	tSET_HISTOGRAM_NUM_BARS        = "SET-HISTOGRAM-NUM-BARS"
	tSET_PATCH_SIZE                = "SET-PATCH-SIZE"
	tSET_PLOT_BACKGROUND_COLOR     = "SET-PLOT-BACKGROUND-COLOR"
	tSET_PLOT_PEN_COLOR            = "SET-PLOT-PEN-COLOR"
	tSET_PLOT_PEN_INTERVAL         = "SET-PLOT-PEN-INTERVAL"
	tSET_PLOT_PEN_MODE             = "SET-PLOT-PEN-MODE"
	tSETUP_PLOTS                   = "SETUP-PLOTS"
	tSET_PLOT_X_RANGE              = "SET-PLOT-X-RANGE"
	tSET_PLOT_Y_RANGE              = "SET-PLOT-Y-RANGE"
	tSET_XY                        = "SET-XY"
	tSHADE_OF                      = "SHADE-OF"
	tSHAPE                         = "SHAPE"
	tSHAPES                        = "SHAPES"
	tSHOW                          = "SHOW"
	tSHOW_TURTLE                   = "SHOW-TURTLE"
	tSHOW_LINK                     = "SHOW-LINK"
	tSHUFFLE                       = "SHUFFLE"
	tSIN                           = "SIN"
	tSIZE                          = "SIZE"
	tSORT                          = "SORT"
	tSORT_BY                       = "SORT-BY"
	tSORT_ON                       = "SORT-ON"
	tSPROUT                        = "SPROUT"
	tSQRT                          = "SQRT"
	tSTAMP                         = "STAMP"
	tSTAMP_ERASE                   = "STAMP-ERASE"
	tSTANDARD_DEVIATION            = "STANDARD-DEVIATION"
	tSTARTUP                       = "STARTUP"
	tSTOP                          = "STOP"
	tSUBJECT                       = "SUBJECT"
	tSUBLIST                       = "SUBLIST"
	tSUBSTRING                     = "SUBSTRING"
	tSUBTRACT_HEADINGS             = "SUBTRACT-HEADINGS"
	tSUM                           = "SUM"
	tTAN                           = "TAN"
	tTHICKNESS                     = "THICKNESS"
	tTICK                          = "TICK"
	tTICK_ADVANCE                  = "TICK-ADVANCE"
	tTICKS                         = "TICKS"
	tTIE                           = "TIE"
	tTIE_MODE                      = "TIE-MODE"
	tTIMER                         = "TIMER"
	tTO                            = "TO"
	tTO_REPORT                     = "TO-REPORT"
	tTOWARDS                       = "TOWARDS"
	tTOWARDSXY                     = "TOWARDSXY"
	tTURTLE                        = "TURTLE"
	tTURTLE_SET                    = "TURTLE-SET"
	tTURTLES                       = "TURTLES"
	tTURTLES_AT                    = "TURTLES-AT"
	tTURTLES_HERE                  = "TURTLES-HERE"
	tTURTLES_ON                    = "TURTLES-ON"
	tTURTLES_OWN                   = "TURTLES-OWN"
	tTYPE                          = "TYPE"
	tUNDIRECTED_LINK_BREED         = "UNDIRECTED-LINK-BREED"
	tUNTIE                         = "UNTIE"
	tUP_TO_N_OF                    = "UP-TO-N-OF"
	tUPDATE_PLOTS                  = "UPDATE-PLOTS"
	tUPHILL                        = "UPHILL"
	tUPHILL4                       = "UPHILL4"
	tUSER_DIRECTORY                = "USER-DIRECTORY"
	tUSER_FILE                     = "USER-FILE"
	tUSER_NEW_FILE                 = "USER-NEW-FILE"
	tUSER_INPUT                    = "USER-INPUT"
	tUSER_MESSAGE                  = "USER-MESSAGE"
	tUSER_ONE_OF                   = "USER-ONE-OF"
	tUSER_YES_OR_NO                = "USER-YES-OR-NO"
	tVARIANCE                      = "VARIANCE"
	tWAIT                          = "WAIT"
	tWHILE                         = "WHILE"
	tWHO                           = "WHO"
	tWHO_ARE_NOT                   = "WHO-ARE-NOT"
	tWITH                          = "WITH"
	tLINK_WITH                     = "LINK-WITH"
	tWITH_MAX                      = "WITH-MAX"
	tWITH_MIN                      = "WITH-MIN"
	tWITH_LOCAL_RANDOMNESS         = "WITH-LOCAL-RANDOMNESS"
	tWITHOUT_INTERRUPTION          = "WITHOUT-INTERRUPTION"
	tWORD                          = "WORD"
	tWORLD_WIDTH                   = "WORLD-WIDTH"
	tWORLD_HEIGHT                  = "WORLD-HEIGHT"
	tWRAP_COLOR                    = "WRAP-COLOR"
	tWRITE                         = "WRITE"
	tXCOR                          = "XCOR"
	tXOR                           = "XOR"
	tYCOR                          = "YCOR"
	tARROW                         = "ARROW"

	tNUMBER  = "NUMBER"
	tSTRING  = "STRING"
	tBOOLEAN = "BOOLEAN"
	tVAR     = "VAR"
	tNEWLINE = "NEWLINE"

	tCREATE_ORDERED_BREEDS = "CREATE-ORDERED-BREEDS"
	tCREATE_BREED_TO       = "CREATE-BREED-TO"
	tCREATE_BREEDS_TO      = "CREATE-BREEDS-TO"
	tCREATE_BREED_FROM     = "CREATE-BREED-FROM"
	tCREATE_BREEDS_FROM    = "CREATE-BREEDS-FROM"
	tCREATE_BREED_WITH     = "CREATE-BREED-WITH"
	tCREATE_BREEDS_WITH    = "CREATE-BREEDS-WITH"
	tCREATE_BREEDS         = "CREATE-BREEDS"
	tIN_BREED_NEIGHBOR     = "IN-BREED-NEIGHBOR"
	tIN_BREED_NEIGHBORS    = "IN-BREED-NEIGHBORS"
	tIN_BREED_FROM         = "IN-BREED-FROM"
	tMY_BREEDS             = "MY-BREEDS"
	tMY_IN_BREEDS          = "MY-IN-BREEDS"
	tMY_OUT_BREEDS         = "MY-OUT-BREEDS"
	tBREED_NEIGHBOR        = "BREED-NEIGHBOR"
	tOUT_BREED_NEIGHBOR    = "OUT-BREED-NEIGHBOR"
	tOUT_BREED_NEIGHBORS   = "OUT-BREED-NEIGHBORS"
	tOUT_BREED_TO          = "OUT-BREED-TO"
	tBREEDS_AT             = "BREEDS-AT"
	tBREEDS_HERE           = "BREEDS-HERE"
	tBREEDS_ON             = "BREEDS-ON"
	tBREEDS_OWN            = "BREEDS-OWN"
	tBREED_WITH            = "BREED-WITH"

	//tokens added by the token augmenter
	tOPEN_BRACKET_OF                  = "OPEN-BRACKET-OF"
	tCLOSE_BRACKET_OF                 = "CLOSE-BRACKET-OF"
	tOPEN_BRACKET_FUNCTION_PARAMETER  = "OPEN-BRACKET-FUNCTION-PARAMETER"
	tCLOSE_BRACKET_FUNCTION_PARAMETER = "CLOSE-BRACKET-FUNCTION-PARAMETER"
	tFUNCTION_NAME                    = "FUNCTION-NAME"
	tFUNCTION_CALL                    = "FUNCTION-CALL"

	tROOT = "ROOT"
)