package converter

type parseFunc func(tokens []token, index *int) (ast, error)

var parseMap = map[string]*parseExpression{
	tROOT:                             &rootExp,
	tE:                                nil,
	tPI:                               nil,
	tTRUE:                             nil,
	tFALSE:                            nil,
	tBLACK:                            &singleValueExpr,
	tGRAY:                             &singleValueExpr,
	tWHITE:                            &singleValueExpr,
	tRED:                              &singleValueExpr,
	tORANGE:                           &singleValueExpr,
	tBROWN:                            &singleValueExpr,
	tYELLOW:                           &singleValueExpr,
	tGREEN:                            &singleValueExpr,
	tLIME:                             &singleValueExpr,
	tTURQUOISE:                        &singleValueExpr,
	tCYAN:                             &singleValueExpr,
	tSKY:                              &singleValueExpr,
	tBLUE:                             &singleValueExpr,
	tVIOLET:                           &singleValueExpr,
	tMAGENTA:                          &singleValueExpr,
	tPINK:                             &singleValueExpr,
	tABS:                              nil,
	tACOS:                             nil,
	tALL:                              &allExpr,
	tAND:                              &singleValueExpr,
	tANY:                              nil,
	tAPPROXIMATE_HSB:                  nil,
	tAPPROXIMATE_RGB:                  nil,
	tPLUS:                             &singleValueExpr,
	tMULTIPLY:                         &singleValueExpr,
	tMINUS:                            &singleValueExpr,
	tDIVIDE:                           &singleValueExpr,
	tPOWER:                            &singleValueExpr,
	tLESS:                             &singleValueExpr,
	tGREATER:                          &singleValueExpr,
	tLESS_EQUAL:                       &singleValueExpr,
	tGREATER_EQUAL:                    &singleValueExpr,
	tEQUAL:                            &singleValueExpr,
	tNOT_EQUAL:                        &singleValueExpr,
	tOPEN_BRACKET:                     nil,
	tCLOSE_BRACKET:                    nil,
	tOPEN_BRACKET_OF:                  &openBracketOfExpr,
	tCLOSE_BRACKET_OF:                 &singleValueExpr,
	tOPEN_BRACKET_FUNCTION_PARAMETER:  &openBracketFunctionParameterExpr,
	tCLOSE_BRACKET_FUNCTION_PARAMETER: &singleValueExpr,
	tOPEN_PARENTHESIS:                 &leftParenthesisExpr,
	tCLOSE_PARENTHESIS:                nil,
	tASIN:                             nil,
	tASK:                              &askExpr,
	tAT_POINTS:                        nil,
	tATAN:                             nil,
	tBACK:                             nil,
	tBASE_COLORS:                      nil,
	tBOTH_ENDS:                        nil,
	tBREED:                            &breedExp,
	tBUT_FIRST:                        nil,
	tBUT_LAST:                         nil,
	tCAN_MOVE:                         nil,
	tCEILING:                          nil,
	tCLEAR_ALL:                        &singleValueExpr,
	tCLEAR_ALL_PLOTS:                  nil,
	tCLEAR_DRAWING:                    nil,
	tCLEAR_GLOBALS:                    nil,
	tCLEAR_LINKS:                      nil,
	tCLEAR_OUTPUT:                     nil,
	tCLEAR_PATCHES:                    nil,
	tCLEAR_PLOT:                       nil,
	tCLEAR_PLOTS:                      nil,
	tCLEAR_TICKS:                      nil,
	tCOLOR:                            &singleValueExpr,
	tCOS:                              nil,
	tCOUNT:                            nil,
	tCREATE_ORDERED_BREEDS:            nil,
	tCREATE_BREED_TO:                  nil,
	tCREATE_BREEDS_TO:                 nil,
	tCREATE_BREED_FROM:                nil,
	tCREATE_BREEDS_FROM:               nil,
	tCREATE_BREED_WITH:                nil,
	tCREATE_BREEDS_WITH:               nil,
	tCREATE_BREEDS:                    &createBreedsExpr,
	tCREATE_ORDERED_TURTLES:           nil,
	tCREATE_LINK_TO:                   nil,
	tCREATE_LINKS_TO:                  nil,
	tCREATE_LINK_FROM:                 nil,
	tCREATE_LINKS_FROM:                nil,
	tCREATE_LINK_WITH:                 nil,
	tCREATE_LINKS_WITH:                nil,
	tCREATE_TURTLES:                   nil,
	tDATE_AND_TIME:                    nil,
	tDIE:                              &singleValueExpr,
	tDIFFUSE:                          nil,
	tDIFFUSE4:                         nil,
	tDIRECTED_LINK_BREED:              nil,
	tDISPLAY:                          nil,
	tDISTANCE:                         nil,
	tDISTANCEXY:                       nil,
	tDOWNHILL:                         nil,
	tDOWNHILL4:                        nil,
	tDX:                               nil,
	tDY:                               nil,
	tEMPTY:                            nil,
	tEND:                              nil,
	tEND1:                             nil,
	tEND2:                             nil,
	tERROR:                            nil,
	tERROR_MESSAGE:                    nil,
	tEVERY:                            nil,
	tEXP:                              nil,
	tEXTENSIONS:                       nil,
	tEXTRACT_HSB:                      nil,
	tEXTRACT_RGB:                      nil,
	tFACE:                             &faceExpr,
	tFACEXY:                           &faceXYCallExpr,
	tFILTER:                           nil,
	tFIRST:                            nil,
	tFLOOR:                            nil,
	tFOLLOW:                           nil,
	tFOLLOW_ME:                        nil,
	tFOREACH:                          nil,
	tFORWARD:                          &forwardCallExpr,
	tFPUT:                             nil,
	tFUNCTION_NAME:                    &functionNameExpr,
	tFUNCTION_CALL:                    &functionCallExpr,
	tGLOBALS:                          &globalsExpr,
	tHATCH:                            nil,
	tHEADING:                          &singleValueExpr,
	tHIDDEN:                           nil,
	tHIDE_LINK:                        nil,
	tHIDE_TURTLE:                      nil,
	tHISTOGRAM:                        nil,
	tHOME:                             nil,
	tHSB:                              nil,
	tIF:                               &ifExpr,
	tIFELSE:                           &ifElseExpr,
	tIFELSE_VALUE:                     nil,
	tIMPORT_DRAWING:                   nil,
	tIMPORT_PCOLORS:                   nil,
	tIMPORT_PCOLORS_RGB:               nil,
	tIMPORT_WORLD:                     nil,
	tIN_BREED_NEIGHBOR:                nil,
	tIN_BREED_NEIGHBORS:               nil,
	tIN_BREED_FROM:                    nil,
	tIN_CONE:                          nil,
	tIN_LINK_NEIGHBOR:                 nil,
	tIN_LINK_NEIGHBORS:                nil,
	tIN_LINK_FROM:                     nil,
	tIN_RADIUS:                        nil,
	tINSERT_ITEM:                      nil,
	tINSPECT:                          nil,
	tINT:                              nil,
	tIS_AGENT:                         nil,
	tIS_AGENTSET:                      nil,
	tIS_ANONYMOUS_COMMAND:             nil,
	tIS_ANONYMOUS_REPORTER:            nil,
	tIS_BOOLEAN:                       nil,
	tIS_DIRECTED_LINK:                 nil,
	tIS_LINK:                          nil,
	tIS_LINKSET:                       nil,
	tIS_LIST:                          nil,
	tIS_NUMBER:                        nil,
	tIS_PATCH:                         nil,
	tIS_PATCHSET:                      nil,
	tIS_STRING:                        nil,
	tIS_TURTLE:                        nil,
	tIS_TURTLESET:                     nil,
	tIS_UNDIRECTED_LINK:               nil,
	tITEM:                             nil,
	tJUMP:                             nil,
	tLABEL:                            nil,
	tLABEL_COLOR:                      nil,
	tLAST:                             nil,
	tLAYOUT_CIRCLE:                    nil,
	tLAYOUT_RADIAL:                    nil,
	tLAYOUT_SPRING:                    nil,
	tLAYOUT_TUTTET:                    nil,
	tLEFT:                             &leftExpr,
	tLENGTH:                           nil,
	tLET:                              nil,
	tLINK:                             nil,
	tLINK_HEADING:                     nil,
	tLINK_LENGTH:                      nil,
	tLINK_SET:                         nil,
	tLINK_SHAPES:                      nil,
	tLINKS:                            nil,
	tLINKS_OWN:                        nil,
	tLIST:                             nil,
	tLN:                               nil,
	tLOG:                              nil,
	tLOOP:                             nil,
	tLPUT:                             nil,
	tMAP:                              nil,
	tMAX:                              &maxExpr,
	tMAX_N_OF:                         nil,
	tMAX_ONE_OF:                       nil,
	tMAX_PXCOR:                        &singleValueExpr,
	tMAX_PYCOR:                        nil,
	tMEAN:                             nil,
	tMEDIAN:                           nil,
	tMEMBER:                           nil,
	tMIN:                              nil,
	tMIN_N_OF:                         nil,
	tMIN_ONE_OF:                       nil,
	tMIN_PXCOR:                        &singleValueExpr,
	tMIN_PYCOR:                        nil,
	tMOD:                              nil,
	tMODES:                            nil,
	tMOUSE_DOWN:                       nil,
	tMOUSE_INSIDE:                     nil,
	tMOUSE_XCOR:                       nil,
	tMOUSE_YCOR:                       nil,
	tMOVE_TO:                          nil,
	tMY_BREEDS:                        nil,
	tMY_LINKS:                         nil,
	tMY_IN_BREEDS:                     nil,
	tMY_IN_LINKS:                      nil,
	tMY_OUT_BREEDS:                    nil,
	tMY_OUT_LINKS:                     nil,
	tBREED_NEIGHBOR:                   nil,
	tOUT_BREED_NEIGHBOR:               nil,
	tOUT_BREED_NEIGHBORS:              nil,
	tOUT_BREED_TO:                     nil,
	tBREEDS_AT:                        nil,
	tBREEDS_HERE:                      nil,
	tBREEDS_ON:                        nil,
	tBREEDS_OWN:                       nil,
	tBREED_WITH:                       nil,
	tMYSELF:                           nil,
	tN_OF:                             nil,
	tN_VALUES:                         nil,
	tNEIGHBORS:                        nil,
	tNEIGHBORS4:                       nil,
	tLINK_NEIGHBOR:                    nil,
	tNEW_SEED:                         nil,
	tNO_DISPLAY:                       nil,
	tNOBODY:                           &singleValueExpr,
	tNO_LINKS:                         nil,
	tNO_PATCHES:                       nil,
	tNOT:                              nil,
	tNO_TURTLES:                       nil,
	tNUMBER:                           &singleValueExpr,
	tOF:                               &singleValueExpr,
	tONE_OF:                           nil,
	tOR:                               nil,
	tOTHER:                            nil,
	tOTHER_END:                        nil,
	tOUT_LINK_NEIGHBOR:                nil,
	tOUT_LINK_NEIGHBORS:               nil,
	tOUT_LINK_TO:                      nil,
	tOUTPUT_PRINT:                     nil,
	tOUTPUT_TYPE:                      nil,
	tOUTPUT_SHOW:                      nil,
	tPATCH:                            &patchExpr,
	tOUTPUT_WRITE:                     nil,
	tPATCH_AT:                         &patchAtExpr,
	tPATCH_AHEAD:                      nil,
	tPATCH_HERE:                       nil,
	tPATCH_AT_HEADING_AND_DISTANCE:    nil,
	tPATCH_RIGHT_AND_AHEAD:            nil,
	tPATCH_LEFT_AND_AHEAD:             nil,
	tPATCH_SIZE:                       nil,
	tPATCH_SET:                        nil,
	tPATCHES_OWN:                      nil,
	tPATCHES:                          nil,
	tPLABEL:                           nil,
	tPCOLOR:                           nil,
	tPOSITION:                         nil,
	tPLABEL_COLOR:                     nil,
	tPRINT:                            nil,
	tPRECISION:                        nil,
	tPYCOR:                            nil,
	tPXCOR:                            nil,
	tRANDOM_FLOAT:                     &randomFloatExpr,
	tRANDOM:                           &randomExpr,
	tRANDOM_GAMMA:                     nil,
	tRANDOM_EXPONENTIAL:               nil,
	tRANDOM_POISSON:                   nil,
	tRANDOM_NORMAL:                    nil,
	tRANDOM_PYCOR:                     nil,
	tRANDOM_PXCOR:                     nil,
	tRANDOM_XCOR:                      nil,
	tRANDOM_SEED:                      nil,
	tRANGE:                            nil,
	tRANDOM_YCOR:                      nil,
	tREDUCE:                           nil,
	tREAD_FROM_STRING:                 nil,
	tREMOVE:                           nil,
	tREMAINDER:                        nil,
	tREMOVE_ITEM:                      nil,
	tREMOVE_DUPLICATES:                nil,
	tREPLACE_ITEM:                     nil,
	tREPEAT:                           nil,
	tRESET_PERSPECTIVE:                nil,
	tREPORT:                           &reportExpr,
	tRESET_TICKS:                      &singleValueExpr,
	tRESIZE_WORLD:                     nil,
	tRESET_TIMER:                      nil,
	tRGB:                              nil,
	tREVERSE:                          nil,
	tRIDE_ME:                          nil,
	tRIDE:                             nil,
	tRIGHT:                            &rightExpr,
	tSCALE_COLOR:                      nil,
	tROUND:                            nil,
	tSEMICOLON:                        nil,
	tSELF:                             nil,
	tSENTENCE:                         nil,
	tSET_DEFAULT_SHAPE:                &setDefaultShapeExpr,
	tSET:                              &setExpr,
	tSET_PATCH_SIZE:                   nil,
	tSET_HISTOGRAM_NUM_BARS:           nil,
	tSET_PLOT_PEN_COLOR:               nil,
	tSET_PLOT_BACKGROUND_COLOR:        nil,
	tSET_PLOT_PEN_MODE:                nil,
	tSET_PLOT_PEN_INTERVAL:            nil,
	tSET_PLOT_X_RANGE:                 nil,
	tSETUP_PLOTS:                      nil,
	tSET_XY:                           &setXYExpr,
	tSET_PLOT_Y_RANGE:                 nil,
	tSHAPE:                            &singleValueExpr,
	tSHADE_OF:                         nil,
	tSHOW:                             nil,
	tSHAPES:                           nil,
	tSHOW_TURTLE:                      nil,
	tSHUFFLE:                          nil,
	tSHOW_LINK:                        nil,
	tSIZE:                             &singleValueExpr,
	tSIN:                              nil,
	tSORT_BY:                          nil,
	tSORT:                             nil,
	tSPROUT:                           &sproutExpr,
	tSORT_ON:                          nil,
	tSTAMP:                            &singleValueExpr,
	tSTRING:                           &singleValueExpr,
	tSQRT:                             nil,
	tSTANDARD_DEVIATION:               nil,
	tSTAMP_ERASE:                      nil,
	tSTOP:                             &singleValueExpr,
	tSTARTUP:                          nil,
	tSUBLIST:                          nil,
	tSUBJECT:                          nil,
	tSUBTRACT_HEADINGS:                nil,
	tSUBSTRING:                        nil,
	tTAN:                              nil,
	tSUM:                              nil,
	tTICK:                             &singleValueExpr,
	tTHICKNESS:                        nil,
	tTICKS:                            nil,
	tTICK_ADVANCE:                     nil,
	tTIE_MODE:                         nil,
	tTIE:                              nil,
	tTO:                               &toExpr,
	tTIMER:                            nil,
	tTOWARDS:                          nil,
	tTO_REPORT:                        &toExpr,
	tTURTLE:                           &turtleExpr,
	tTOWARDSXY:                        nil,
	tTURTLES:                          &singleValueExpr,
	tTURTLE_SET:                       nil,
	tTURTLES_HERE:                     nil,
	tTURTLES_AT:                       nil,
	tTURTLES_OWN:                      nil,
	tTURTLES_ON:                       nil,
	tUNDIRECTED_LINK_BREED:            nil,
	tTYPE:                             nil,
	tUP_TO_N_OF:                       nil,
	tUNTIE:                            nil,
	tUPHILL:                           nil,
	tUPDATE_PLOTS:                     nil,
	tUSER_DIRECTORY:                   nil,
	tUPHILL4:                          nil,
	tUSER_NEW_FILE:                    nil,
	tUSER_FILE:                        nil,
	tUSER_MESSAGE:                     nil,
	tUSER_INPUT:                       nil,
	tUSER_YES_OR_NO:                   nil,
	tUSER_ONE_OF:                      nil,
	tWAIT:                             nil,
	tVAR:                              &singleValueExpr,
	tVARIANCE:                         nil,
	tWHO:                              &singleValueExpr,
	tWHILE:                            nil,
	tWITH:                             nil,
	tWHO_ARE_NOT:                      nil,
	tWITH_MAX:                         nil,
	tLINK_WITH:                        nil,
	tWITH_LOCAL_RANDOMNESS:            nil,
	tWITH_MIN:                         nil,
	tWORD:                             nil,
	tWITHOUT_INTERRUPTION:             nil,
	tWORLD_HEIGHT:                     nil,
	tWORLD_WIDTH:                      nil,
	tWRITE:                            nil,
	tWRAP_COLOR:                       nil,
	tXOR:                              nil,
	tXCOR:                             &singleValueExpr,
	tARROW:                            nil,
	tYCOR:                             nil,
}
