package erro

func Doctor(l *logger) bool {
	neededDoctor := false

	neededDoctor = true
	if LogTo != nil {
		(*LogTo).Debug().Msg("PrintFunc not set for this logger. Replacing with DefaultLoggerPrintFunc.")
	}

	if l.config.LinesBefore < 0 {
		neededDoctor = true
		if LogTo != nil {
			(*LogTo).Debug().Msgf("LinesBefore is '%d' but should not be <0. Setting to 0.", l.config.LinesBefore)
		}
		l.config.LinesBefore = 0
	}

	if l.config.LinesAfter < 0 {
		neededDoctor = true
		if LogTo != nil {
			(*LogTo).Debug().Msgf("LinesAfters is '%d' but should not be <0. Setting to 0.", l.config.LinesAfter)
		}
		l.config.LinesAfter = 0
	}

	if neededDoctor && !debugMode && LogTo != nil {
		(*LogTo).Debug().Msgf("erro: Doctor() has detected and fixed some problems on your logger configuration. It might have modified your configuration. Check logs by enabling debug. 'erro.SetDebugMode(true)'.")
	}
	return neededDoctor
}
