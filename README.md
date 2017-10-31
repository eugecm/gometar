# GOMETAR
[![Build Status](https://travis-ci.org/eugecm/gometar.svg?branch=master)](https://travis-ci.org/eugecm/gometar)
[![Coverage](https://codecov.io/gh/eugecm/gometar/branch/master/graph/badge.svg)]()

GOMETAR is a METAR parser written in Go. A METAR is a routine report of
meteorological conditions at an aerodrome.

It supports International and US-like METARs.

## Example
```go
    import (
        "fmt"

        "github.com/eugecm/gometar/metar/parser"
    )

    p := parser.New()
    report, _ := p.Parse("EGNM 191750Z 13011KT 1400 R14/P1500 RA BR SCT001 BKN002 13/13 Q0997")
    fmt.Println(report.Clouds)
```

## TODO
* Add support for remarks
* Add support for TAF
* Clean up API

# Links

* [Australian Government: Bureau of Meteorology - METAR/SPECI](http://www.bom.gov.au/aviation/data/education/metar-speci.pdf)
* [Embry-Riddle Aeronautical University - U.S METAR/SPECI CODE FORMAT WITH REMARKS](http://wx.erau.edu/reference/text/metar_code_format.pdf)
* [US Department of Transportation - Aviation Weather Formats: METAR/TAF](https://www.uscg.mil/auxiliary/missions/auxair/metar_taf.pdf)
* [Wunderground - Metar Tutorial](https://www.wunderground.com/metarFAQ.asp)
* [VATSIM.net - Interpreting METARs and TAFs](https://www.vatsim.net/pilot-resource-centre/general-lessons/interpreting-metars-and-tafs)
