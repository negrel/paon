package property

var (
	colorBlack                = ColorFromHex(0x000000)
	colorMaroon               = ColorFromHex(0x800000)
	colorGreen                = ColorFromHex(0x008000)
	colorOlive                = ColorFromHex(0x808000)
	colorNavy                 = ColorFromHex(0x000080)
	colorPurple               = ColorFromHex(0x800080)
	colorTeal                 = ColorFromHex(0x008080)
	colorSilver               = ColorFromHex(0xC0C0C0)
	colorGray                 = ColorFromHex(0x808080)
	colorRed                  = ColorFromHex(0xFF0000)
	colorLime                 = ColorFromHex(0x00FF00)
	colorYellow               = ColorFromHex(0xFFFF00)
	colorBlue                 = ColorFromHex(0x0000FF)
	colorFuchsia              = ColorFromHex(0xFF00FF)
	colorAqua                 = ColorFromHex(0x00FFFF)
	colorWhite                = ColorFromHex(0xFFFFFF)
	color16                   = ColorFromHex(0x000000)
	color17                   = ColorFromHex(0x00005F)
	color18                   = ColorFromHex(0x000087)
	color19                   = ColorFromHex(0x0000AF)
	color20                   = ColorFromHex(0x0000D7)
	color21                   = ColorFromHex(0x0000FF)
	color22                   = ColorFromHex(0x005F00)
	color23                   = ColorFromHex(0x005F5F)
	color24                   = ColorFromHex(0x005F87)
	color25                   = ColorFromHex(0x005FAF)
	color26                   = ColorFromHex(0x005FD7)
	color27                   = ColorFromHex(0x005FFF)
	color28                   = ColorFromHex(0x008700)
	color29                   = ColorFromHex(0x00875F)
	color30                   = ColorFromHex(0x008787)
	color31                   = ColorFromHex(0x0087Af)
	color32                   = ColorFromHex(0x0087D7)
	color33                   = ColorFromHex(0x0087FF)
	color34                   = ColorFromHex(0x00AF00)
	color35                   = ColorFromHex(0x00AF5F)
	color36                   = ColorFromHex(0x00AF87)
	color37                   = ColorFromHex(0x00AFAF)
	color38                   = ColorFromHex(0x00AFD7)
	color39                   = ColorFromHex(0x00AFFF)
	color40                   = ColorFromHex(0x00D700)
	color41                   = ColorFromHex(0x00D75F)
	color42                   = ColorFromHex(0x00D787)
	color43                   = ColorFromHex(0x00D7AF)
	color44                   = ColorFromHex(0x00D7D7)
	color45                   = ColorFromHex(0x00D7FF)
	color46                   = ColorFromHex(0x00FF00)
	color47                   = ColorFromHex(0x00FF5F)
	color48                   = ColorFromHex(0x00FF87)
	color49                   = ColorFromHex(0x00FFAF)
	color50                   = ColorFromHex(0x00FFd7)
	color51                   = ColorFromHex(0x00FFFF)
	color52                   = ColorFromHex(0x5F0000)
	color53                   = ColorFromHex(0x5F005F)
	color54                   = ColorFromHex(0x5F0087)
	color55                   = ColorFromHex(0x5F00AF)
	color56                   = ColorFromHex(0x5F00D7)
	color57                   = ColorFromHex(0x5F00FF)
	color58                   = ColorFromHex(0x5F5F00)
	color59                   = ColorFromHex(0x5F5F5F)
	color60                   = ColorFromHex(0x5F5F87)
	color61                   = ColorFromHex(0x5F5FAF)
	color62                   = ColorFromHex(0x5F5FD7)
	color63                   = ColorFromHex(0x5F5FFF)
	color64                   = ColorFromHex(0x5F8700)
	color65                   = ColorFromHex(0x5F875F)
	color66                   = ColorFromHex(0x5F8787)
	color67                   = ColorFromHex(0x5F87AF)
	color68                   = ColorFromHex(0x5F87D7)
	color69                   = ColorFromHex(0x5F87FF)
	color70                   = ColorFromHex(0x5FAF00)
	color71                   = ColorFromHex(0x5FAF5F)
	color72                   = ColorFromHex(0x5FAF87)
	color73                   = ColorFromHex(0x5FAFAF)
	color74                   = ColorFromHex(0x5FAFD7)
	color75                   = ColorFromHex(0x5FAFFF)
	color76                   = ColorFromHex(0x5FD700)
	color77                   = ColorFromHex(0x5FD75F)
	color78                   = ColorFromHex(0x5FD787)
	color79                   = ColorFromHex(0x5FD7AF)
	color80                   = ColorFromHex(0x5FD7D7)
	color81                   = ColorFromHex(0x5FD7FF)
	color82                   = ColorFromHex(0x5FFF00)
	color83                   = ColorFromHex(0x5FFF5F)
	color84                   = ColorFromHex(0x5FFF87)
	color85                   = ColorFromHex(0x5FFFAF)
	color86                   = ColorFromHex(0x5FFFD7)
	color87                   = ColorFromHex(0x5FFFFF)
	color88                   = ColorFromHex(0x870000)
	color89                   = ColorFromHex(0x87005F)
	color90                   = ColorFromHex(0x870087)
	color91                   = ColorFromHex(0x8700AF)
	color92                   = ColorFromHex(0x8700D7)
	color93                   = ColorFromHex(0x8700FF)
	color94                   = ColorFromHex(0x875F00)
	color95                   = ColorFromHex(0x875F5F)
	color96                   = ColorFromHex(0x875F87)
	color97                   = ColorFromHex(0x875FAF)
	color98                   = ColorFromHex(0x875FD7)
	color99                   = ColorFromHex(0x875FFF)
	color100                  = ColorFromHex(0x878700)
	color101                  = ColorFromHex(0x87875F)
	color102                  = ColorFromHex(0x878787)
	color103                  = ColorFromHex(0x8787AF)
	color104                  = ColorFromHex(0x8787D7)
	color105                  = ColorFromHex(0x8787FF)
	color106                  = ColorFromHex(0x87AF00)
	color107                  = ColorFromHex(0x87AF5F)
	color108                  = ColorFromHex(0x87AF87)
	color109                  = ColorFromHex(0x87AFAF)
	color110                  = ColorFromHex(0x87AFD7)
	color111                  = ColorFromHex(0x87AFFF)
	color112                  = ColorFromHex(0x87D700)
	color113                  = ColorFromHex(0x87D75F)
	color114                  = ColorFromHex(0x87D787)
	color115                  = ColorFromHex(0x87D7AF)
	color116                  = ColorFromHex(0x87D7D7)
	color117                  = ColorFromHex(0x87D7FF)
	color118                  = ColorFromHex(0x87FF00)
	color119                  = ColorFromHex(0x87FF5F)
	color120                  = ColorFromHex(0x87FF87)
	color121                  = ColorFromHex(0x87FFAF)
	color122                  = ColorFromHex(0x87FFD7)
	color123                  = ColorFromHex(0x87FFFF)
	color124                  = ColorFromHex(0xAF0000)
	color125                  = ColorFromHex(0xAF005F)
	color126                  = ColorFromHex(0xAF0087)
	color127                  = ColorFromHex(0xAF00AF)
	color128                  = ColorFromHex(0xAF00D7)
	color129                  = ColorFromHex(0xAF00FF)
	color130                  = ColorFromHex(0xAF5F00)
	color131                  = ColorFromHex(0xAF5F5F)
	color132                  = ColorFromHex(0xAF5F87)
	color133                  = ColorFromHex(0xAF5FAF)
	color134                  = ColorFromHex(0xAF5FD7)
	color135                  = ColorFromHex(0xAF5FFF)
	color136                  = ColorFromHex(0xAF8700)
	color137                  = ColorFromHex(0xAF875F)
	color138                  = ColorFromHex(0xAF8787)
	color139                  = ColorFromHex(0xAF87AF)
	color140                  = ColorFromHex(0xAF87D7)
	color141                  = ColorFromHex(0xAF87FF)
	color142                  = ColorFromHex(0xAFAF00)
	color143                  = ColorFromHex(0xAFAF5F)
	color144                  = ColorFromHex(0xAFAF87)
	color145                  = ColorFromHex(0xAFAFAF)
	color146                  = ColorFromHex(0xAFAFD7)
	color147                  = ColorFromHex(0xAFAFFF)
	color148                  = ColorFromHex(0xAFD700)
	color149                  = ColorFromHex(0xAFD75F)
	color150                  = ColorFromHex(0xAFD787)
	color151                  = ColorFromHex(0xAFD7AF)
	color152                  = ColorFromHex(0xAFD7D7)
	color153                  = ColorFromHex(0xAFD7FF)
	color154                  = ColorFromHex(0xAFFF00)
	color155                  = ColorFromHex(0xAFFF5F)
	color156                  = ColorFromHex(0xAFFF87)
	color157                  = ColorFromHex(0xAFFFAF)
	color158                  = ColorFromHex(0xAFFFD7)
	color159                  = ColorFromHex(0xAFFFFF)
	color160                  = ColorFromHex(0xD70000)
	color161                  = ColorFromHex(0xD7005F)
	color162                  = ColorFromHex(0xD70087)
	color163                  = ColorFromHex(0xD700AF)
	color164                  = ColorFromHex(0xD700D7)
	color165                  = ColorFromHex(0xD700FF)
	color166                  = ColorFromHex(0xD75F00)
	color167                  = ColorFromHex(0xD75F5F)
	color168                  = ColorFromHex(0xD75F87)
	color169                  = ColorFromHex(0xD75FAF)
	color170                  = ColorFromHex(0xD75FD7)
	color171                  = ColorFromHex(0xD75FFF)
	color172                  = ColorFromHex(0xD78700)
	color173                  = ColorFromHex(0xD7875F)
	color174                  = ColorFromHex(0xD78787)
	color175                  = ColorFromHex(0xD787AF)
	color176                  = ColorFromHex(0xD787D7)
	color177                  = ColorFromHex(0xD787FF)
	color178                  = ColorFromHex(0xD7AF00)
	color179                  = ColorFromHex(0xD7AF5F)
	color180                  = ColorFromHex(0xD7AF87)
	color181                  = ColorFromHex(0xD7AFAF)
	color182                  = ColorFromHex(0xD7AFD7)
	color183                  = ColorFromHex(0xD7AFFF)
	color184                  = ColorFromHex(0xD7D700)
	color185                  = ColorFromHex(0xD7D75F)
	color186                  = ColorFromHex(0xD7D787)
	color187                  = ColorFromHex(0xD7D7AF)
	color188                  = ColorFromHex(0xD7D7D7)
	color189                  = ColorFromHex(0xD7D7FF)
	color190                  = ColorFromHex(0xD7FF00)
	color191                  = ColorFromHex(0xD7FF5F)
	color192                  = ColorFromHex(0xD7FF87)
	color193                  = ColorFromHex(0xD7FFAF)
	color194                  = ColorFromHex(0xD7FFD7)
	color195                  = ColorFromHex(0xD7FFFF)
	color196                  = ColorFromHex(0xFF0000)
	color197                  = ColorFromHex(0xFF005F)
	color198                  = ColorFromHex(0xFF0087)
	color199                  = ColorFromHex(0xFF00AF)
	color200                  = ColorFromHex(0xFF00D7)
	color201                  = ColorFromHex(0xFF00FF)
	color202                  = ColorFromHex(0xFF5F00)
	color203                  = ColorFromHex(0xFF5F5F)
	color204                  = ColorFromHex(0xFF5F87)
	color205                  = ColorFromHex(0xFF5FAF)
	color206                  = ColorFromHex(0xFF5FD7)
	color207                  = ColorFromHex(0xFF5FFF)
	color208                  = ColorFromHex(0xFF8700)
	color209                  = ColorFromHex(0xFF875F)
	color210                  = ColorFromHex(0xFF8787)
	color211                  = ColorFromHex(0xFF87AF)
	color212                  = ColorFromHex(0xFF87D7)
	color213                  = ColorFromHex(0xFF87FF)
	color214                  = ColorFromHex(0xFFAF00)
	color215                  = ColorFromHex(0xFFAF5F)
	color216                  = ColorFromHex(0xFFAF87)
	color217                  = ColorFromHex(0xFFAFAF)
	color218                  = ColorFromHex(0xFFAFD7)
	color219                  = ColorFromHex(0xFFAFFF)
	color220                  = ColorFromHex(0xFFD700)
	color221                  = ColorFromHex(0xFFD75F)
	color222                  = ColorFromHex(0xFFD787)
	color223                  = ColorFromHex(0xFFD7AF)
	color224                  = ColorFromHex(0xFFD7D7)
	color225                  = ColorFromHex(0xFFD7FF)
	color226                  = ColorFromHex(0xFFFF00)
	color227                  = ColorFromHex(0xFFFF5F)
	color228                  = ColorFromHex(0xFFFF87)
	color229                  = ColorFromHex(0xFFFFAF)
	color230                  = ColorFromHex(0xFFFFD7)
	color231                  = ColorFromHex(0xFFFFFF)
	color232                  = ColorFromHex(0x080808)
	color233                  = ColorFromHex(0x121212)
	color234                  = ColorFromHex(0x1C1C1C)
	color235                  = ColorFromHex(0x262626)
	color236                  = ColorFromHex(0x303030)
	color237                  = ColorFromHex(0x3A3A3A)
	color238                  = ColorFromHex(0x444444)
	color239                  = ColorFromHex(0x4E4E4E)
	color240                  = ColorFromHex(0x585858)
	color241                  = ColorFromHex(0x626262)
	color242                  = ColorFromHex(0x6C6C6C)
	color243                  = ColorFromHex(0x767676)
	color244                  = ColorFromHex(0x808080)
	color245                  = ColorFromHex(0x8A8A8A)
	color246                  = ColorFromHex(0x949494)
	color247                  = ColorFromHex(0x9E9E9E)
	color248                  = ColorFromHex(0xA8A8A8)
	color249                  = ColorFromHex(0xB2B2B2)
	color250                  = ColorFromHex(0xBCBCBC)
	color251                  = ColorFromHex(0xC6C6C6)
	color252                  = ColorFromHex(0xD0D0D0)
	color253                  = ColorFromHex(0xDADADA)
	color254                  = ColorFromHex(0xE4E4E4)
	color255                  = ColorFromHex(0xEEEEEE)
	colorAliceBlue            = ColorFromHex(0xF0F8FF)
	colorAntiqueWhite         = ColorFromHex(0xFAEBD7)
	colorAquaMarine           = ColorFromHex(0x7FFFD4)
	colorAzure                = ColorFromHex(0xF0FFFF)
	colorBeige                = ColorFromHex(0xF5F5DC)
	colorBisque               = ColorFromHex(0xFFE4C4)
	colorBlanchedAlmond       = ColorFromHex(0xFFEBCD)
	colorBlueViolet           = ColorFromHex(0x8A2BE2)
	colorBrown                = ColorFromHex(0xA52A2A)
	colorBurlyWood            = ColorFromHex(0xDEB887)
	colorCadetBlue            = ColorFromHex(0x5F9EA0)
	colorChartreuse           = ColorFromHex(0x7FFF00)
	colorChocolate            = ColorFromHex(0xD2691E)
	colorCoral                = ColorFromHex(0xFF7F50)
	colorCornflowerBlue       = ColorFromHex(0x6495ED)
	colorCornsilk             = ColorFromHex(0xFFF8DC)
	colorCrimson              = ColorFromHex(0xDC143C)
	colorDarkBlue             = ColorFromHex(0x00008B)
	colorDarkCyan             = ColorFromHex(0x008B8B)
	colorDarkGoldenrod        = ColorFromHex(0xB8860B)
	colorDarkGray             = ColorFromHex(0xA9A9A9)
	colorDarkGreen            = ColorFromHex(0x006400)
	colorDarkKhaki            = ColorFromHex(0xBDB76B)
	colorDarkMagenta          = ColorFromHex(0x8B008B)
	colorDarkOliveGreen       = ColorFromHex(0x556B2F)
	colorDarkOrange           = ColorFromHex(0xFF8C00)
	colorDarkOrchid           = ColorFromHex(0x9932CC)
	colorDarkRed              = ColorFromHex(0x8B0000)
	colorDarkSalmon           = ColorFromHex(0xE9967A)
	colorDarkSeaGreen         = ColorFromHex(0x8FBC8F)
	colorDarkSlateBlue        = ColorFromHex(0x483D8B)
	colorDarkSlateGray        = ColorFromHex(0x2F4F4F)
	colorDarkTurquoise        = ColorFromHex(0x00CED1)
	colorDarkViolet           = ColorFromHex(0x9400D3)
	colorDeepPink             = ColorFromHex(0xFF1493)
	colorDeepSkyBlue          = ColorFromHex(0x00BFFF)
	colorDimGray              = ColorFromHex(0x696969)
	colorDodgerBlue           = ColorFromHex(0x1E90FF)
	colorFireBrick            = ColorFromHex(0xB22222)
	colorFloralWhite          = ColorFromHex(0xFFFAF0)
	colorForestGreen          = ColorFromHex(0x228B22)
	colorGainsboro            = ColorFromHex(0xDCDCDC)
	colorGhostWhite           = ColorFromHex(0xF8F8FF)
	colorGold                 = ColorFromHex(0xFFD700)
	colorGoldenrod            = ColorFromHex(0xDAA520)
	colorGreenYellow          = ColorFromHex(0xADFF2F)
	colorHoneydew             = ColorFromHex(0xF0FFF0)
	colorHotPink              = ColorFromHex(0xFF69B4)
	colorIndianRed            = ColorFromHex(0xCD5C5C)
	colorIndigo               = ColorFromHex(0x4B0082)
	colorIvory                = ColorFromHex(0xFFFFF0)
	colorKhaki                = ColorFromHex(0xF0E68C)
	colorLavender             = ColorFromHex(0xE6E6FA)
	colorLavenderBlush        = ColorFromHex(0xFFF0F5)
	colorLawnGreen            = ColorFromHex(0x7CFC00)
	colorLemonChiffon         = ColorFromHex(0xFFFACD)
	colorLightBlue            = ColorFromHex(0xADD8E6)
	colorLightCoral           = ColorFromHex(0xF08080)
	colorLightCyan            = ColorFromHex(0xE0FFFF)
	colorLightGoldenrodYellow = ColorFromHex(0xFAFAD2)
	colorLightGray            = ColorFromHex(0xD3D3D3)
	colorLightGreen           = ColorFromHex(0x90EE90)
	colorLightPink            = ColorFromHex(0xFFB6C1)
	colorLightSalmon          = ColorFromHex(0xFFA07A)
	colorLightSeaGreen        = ColorFromHex(0x20B2AA)
	colorLightSkyBlue         = ColorFromHex(0x87CEFA)
	colorLightSlateGray       = ColorFromHex(0x778899)
	colorLightSteelBlue       = ColorFromHex(0xB0C4DE)
	colorLightYellow          = ColorFromHex(0xFFFFE0)
	colorLimeGreen            = ColorFromHex(0x32CD32)
	colorLinen                = ColorFromHex(0xFAF0E6)
	colorMediumAquamarine     = ColorFromHex(0x66CDAA)
	colorMediumBlue           = ColorFromHex(0x0000CD)
	colorMediumOrchid         = ColorFromHex(0xBA55D3)
	colorMediumPurple         = ColorFromHex(0x9370DB)
	colorMediumSeaGreen       = ColorFromHex(0x3CB371)
	colorMediumSlateBlue      = ColorFromHex(0x7B68EE)
	colorMediumSpringGreen    = ColorFromHex(0x00FA9A)
	colorMediumTurquoise      = ColorFromHex(0x48D1CC)
	colorMediumVioletRed      = ColorFromHex(0xC71585)
	colorMidnightBlue         = ColorFromHex(0x191970)
	colorMintCream            = ColorFromHex(0xF5FFFA)
	colorMistyRose            = ColorFromHex(0xFFE4E1)
	colorMoccasin             = ColorFromHex(0xFFE4B5)
	colorNavajoWhite          = ColorFromHex(0xFFDEAD)
	colorOldLace              = ColorFromHex(0xFDF5E6)
	colorOliveDrab            = ColorFromHex(0x6B8E23)
	colorOrange               = ColorFromHex(0xFFA500)
	colorOrangeRed            = ColorFromHex(0xFF4500)
	colorOrchid               = ColorFromHex(0xDA70D6)
	colorPaleGoldenrod        = ColorFromHex(0xEEE8AA)
	colorPaleGreen            = ColorFromHex(0x98FB98)
	colorPaleTurquoise        = ColorFromHex(0xAFEEEE)
	colorPaleVioletRed        = ColorFromHex(0xDB7093)
	colorPapayaWhip           = ColorFromHex(0xFFEFD5)
	colorPeachPuff            = ColorFromHex(0xFFDAB9)
	colorPeru                 = ColorFromHex(0xCD853F)
	colorPink                 = ColorFromHex(0xFFC0CB)
	colorPlum                 = ColorFromHex(0xDDA0DD)
	colorPowderBlue           = ColorFromHex(0xB0E0E6)
	colorRebeccaPurple        = ColorFromHex(0x663399)
	colorRosyBrown            = ColorFromHex(0xBC8F8F)
	colorRoyalBlue            = ColorFromHex(0x4169E1)
	colorSaddleBrown          = ColorFromHex(0x8B4513)
	colorSalmon               = ColorFromHex(0xFA8072)
	colorSandyBrown           = ColorFromHex(0xF4A460)
	colorSeaGreen             = ColorFromHex(0x2E8B57)
	colorSeashell             = ColorFromHex(0xFFF5EE)
	colorSienna               = ColorFromHex(0xA0522D)
	colorSkyblue              = ColorFromHex(0x87CEEB)
	colorSlateBlue            = ColorFromHex(0x6A5ACD)
	colorSlateGray            = ColorFromHex(0x708090)
	colorSnow                 = ColorFromHex(0xFFFAFA)
	colorSpringGreen          = ColorFromHex(0x00FF7F)
	colorSteelBlue            = ColorFromHex(0x4682B4)
	colorTan                  = ColorFromHex(0xD2B48C)
	colorThistle              = ColorFromHex(0xD8BFD8)
	colorTomato               = ColorFromHex(0xFF6347)
	colorTurquoise            = ColorFromHex(0x40E0D0)
	colorViolet               = ColorFromHex(0xEE82EE)
	colorWheat                = ColorFromHex(0xF5DEB3)
	colorWhiteSmoke           = ColorFromHex(0xF5F5F5)
	colorYellowGreen          = ColorFromHex(0x9ACD32)
)

// ColorBlack returns a preallocated color.
func ColorBlack() *Color {
	return &colorBlack
}

// ColorMaroon returns a preallocated color.
func ColorMaroon() *Color {
	return &colorMaroon
}

// ColorGreen returns a preallocated color.
func ColorGreen() *Color {
	return &colorGreen
}

// ColorOlive returns a preallocated color.
func ColorOlive() *Color {
	return &colorOlive
}

// ColorNavy returns a preallocated color.
func ColorNavy() *Color {
	return &colorNavy
}

// ColorPurple returns a preallocated color.
func ColorPurple() *Color {
	return &colorPurple
}

// ColorTeal returns a preallocated color.
func ColorTeal() *Color {
	return &colorTeal
}

// ColorSilver returns a preallocated color.
func ColorSilver() *Color {
	return &colorSilver
}

// ColorGray returns a preallocated color.
func ColorGray() *Color {
	return &colorGray
}

// ColorRed returns a preallocated color.
func ColorRed() *Color {
	return &colorRed
}

// ColorLime returns a preallocated color.
func ColorLime() *Color {
	return &colorLime
}

// ColorYellow returns a preallocated color.
func ColorYellow() *Color {
	return &colorYellow
}

// ColorBlue returns a preallocated color.
func ColorBlue() *Color {
	return &colorBlue
}

// ColorFuchsia returns a preallocated color.
func ColorFuchsia() *Color {
	return &colorFuchsia
}

// ColorAqua returns a preallocated color.
func ColorAqua() *Color {
	return &colorAqua
}

// ColorWhite returns a preallocated color.
func ColorWhite() *Color {
	return &colorWhite
}

// Color16 returns a preallocated color.
func Color16() *Color {
	return &color16
}

// Color17 returns a preallocated color.
func Color17() *Color {
	return &color17
}

// Color18 returns a preallocated color.
func Color18() *Color {
	return &color18
}

// Color19 returns a preallocated color.
func Color19() *Color {
	return &color19
}

// Color20 returns a preallocated color.
func Color20() *Color {
	return &color20
}

// Color21 returns a preallocated color.
func Color21() *Color {
	return &color21
}

// Color22 returns a preallocated color.
func Color22() *Color {
	return &color22
}

// Color23 returns a preallocated color.
func Color23() *Color {
	return &color23
}

// Color24 returns a preallocated color.
func Color24() *Color {
	return &color24
}

// Color25 returns a preallocated color.
func Color25() *Color {
	return &color25
}

// Color26 returns a preallocated color.
func Color26() *Color {
	return &color26
}

// Color27 returns a preallocated color.
func Color27() *Color {
	return &color27
}

// Color28 returns a preallocated color.
func Color28() *Color {
	return &color28
}

// Color29 returns a preallocated color.
func Color29() *Color {
	return &color29
}

// Color30 returns a preallocated color.
func Color30() *Color {
	return &color30
}

// Color31 returns a preallocated color.
func Color31() *Color {
	return &color31
}

// Color32 returns a preallocated color.
func Color32() *Color {
	return &color32
}

// Color33 returns a preallocated color.
func Color33() *Color {
	return &color33
}

// Color34 returns a preallocated color.
func Color34() *Color {
	return &color34
}

// Color35 returns a preallocated color.
func Color35() *Color {
	return &color35
}

// Color36 returns a preallocated color.
func Color36() *Color {
	return &color36
}

// Color37 returns a preallocated color.
func Color37() *Color {
	return &color37
}

// Color38 returns a preallocated color.
func Color38() *Color {
	return &color38
}

// Color39 returns a preallocated color.
func Color39() *Color {
	return &color39
}

// Color40 returns a preallocated color.
func Color40() *Color {
	return &color40
}

// Color41 returns a preallocated color.
func Color41() *Color {
	return &color41
}

// Color42 returns a preallocated color.
func Color42() *Color {
	return &color42
}

// Color43 returns a preallocated color.
func Color43() *Color {
	return &color43
}

// Color44 returns a preallocated color.
func Color44() *Color {
	return &color44
}

// Color45 returns a preallocated color.
func Color45() *Color {
	return &color45
}

// Color46 returns a preallocated color.
func Color46() *Color {
	return &color46
}

// Color47 returns a preallocated color.
func Color47() *Color {
	return &color47
}

// Color48 returns a preallocated color.
func Color48() *Color {
	return &color48
}

// Color49 returns a preallocated color.
func Color49() *Color {
	return &color49
}

// Color50 returns a preallocated color.
func Color50() *Color {
	return &color50
}

// Color51 returns a preallocated color.
func Color51() *Color {
	return &color51
}

// Color52 returns a preallocated color.
func Color52() *Color {
	return &color52
}

// Color53 returns a preallocated color.
func Color53() *Color {
	return &color53
}

// Color54 returns a preallocated color.
func Color54() *Color {
	return &color54
}

// Color55 returns a preallocated color.
func Color55() *Color {
	return &color55
}

// Color56 returns a preallocated color.
func Color56() *Color {
	return &color56
}

// Color57 returns a preallocated color.
func Color57() *Color {
	return &color57
}

// Color58 returns a preallocated color.
func Color58() *Color {
	return &color58
}

// Color59 returns a preallocated color.
func Color59() *Color {
	return &color59
}

// Color60 returns a preallocated color.
func Color60() *Color {
	return &color60
}

// Color61 returns a preallocated color.
func Color61() *Color {
	return &color61
}

// Color62 returns a preallocated color.
func Color62() *Color {
	return &color62
}

// Color63 returns a preallocated color.
func Color63() *Color {
	return &color63
}

// Color64 returns a preallocated color.
func Color64() *Color {
	return &color64
}

// Color65 returns a preallocated color.
func Color65() *Color {
	return &color65
}

// Color66 returns a preallocated color.
func Color66() *Color {
	return &color66
}

// Color67 returns a preallocated color.
func Color67() *Color {
	return &color67
}

// Color68 returns a preallocated color.
func Color68() *Color {
	return &color68
}

// Color69 returns a preallocated color.
func Color69() *Color {
	return &color69
}

// Color70 returns a preallocated color.
func Color70() *Color {
	return &color70
}

// Color71 returns a preallocated color.
func Color71() *Color {
	return &color71
}

// Color72 returns a preallocated color.
func Color72() *Color {
	return &color72
}

// Color73 returns a preallocated color.
func Color73() *Color {
	return &color73
}

// Color74 returns a preallocated color.
func Color74() *Color {
	return &color74
}

// Color75 returns a preallocated color.
func Color75() *Color {
	return &color75
}

// Color76 returns a preallocated color.
func Color76() *Color {
	return &color76
}

// Color77 returns a preallocated color.
func Color77() *Color {
	return &color77
}

// Color78 returns a preallocated color.
func Color78() *Color {
	return &color78
}

// Color79 returns a preallocated color.
func Color79() *Color {
	return &color79
}

// Color80 returns a preallocated color.
func Color80() *Color {
	return &color80
}

// Color81 returns a preallocated color.
func Color81() *Color {
	return &color81
}

// Color82 returns a preallocated color.
func Color82() *Color {
	return &color82
}

// Color83 returns a preallocated color.
func Color83() *Color {
	return &color83
}

// Color84 returns a preallocated color.
func Color84() *Color {
	return &color84
}

// Color85 returns a preallocated color.
func Color85() *Color {
	return &color85
}

// Color86 returns a preallocated color.
func Color86() *Color {
	return &color86
}

// Color87 returns a preallocated color.
func Color87() *Color {
	return &color87
}

// Color88 returns a preallocated color.
func Color88() *Color {
	return &color88
}

// Color89 returns a preallocated color.
func Color89() *Color {
	return &color89
}

// Color90 returns a preallocated color.
func Color90() *Color {
	return &color90
}

// Color91 returns a preallocated color.
func Color91() *Color {
	return &color91
}

// Color92 returns a preallocated color.
func Color92() *Color {
	return &color92
}

// Color93 returns a preallocated color.
func Color93() *Color {
	return &color93
}

// Color94 returns a preallocated color.
func Color94() *Color {
	return &color94
}

// Color95 returns a preallocated color.
func Color95() *Color {
	return &color95
}

// Color96 returns a preallocated color.
func Color96() *Color {
	return &color96
}

// Color97 returns a preallocated color.
func Color97() *Color {
	return &color97
}

// Color98 returns a preallocated color.
func Color98() *Color {
	return &color98
}

// Color99 returns a preallocated color.
func Color99() *Color {
	return &color99
}

// Color100 returns a preallocated color.
func Color100() *Color {
	return &color100
}

// Color101 returns a preallocated color.
func Color101() *Color {
	return &color101
}

// Color102 returns a preallocated color.
func Color102() *Color {
	return &color102
}

// Color103 returns a preallocated color.
func Color103() *Color {
	return &color103
}

// Color104 returns a preallocated color.
func Color104() *Color {
	return &color104
}

// Color105 returns a preallocated color.
func Color105() *Color {
	return &color105
}

// Color106 returns a preallocated color.
func Color106() *Color {
	return &color106
}

// Color107 returns a preallocated color.
func Color107() *Color {
	return &color107
}

// Color108 returns a preallocated color.
func Color108() *Color {
	return &color108
}

// Color109 returns a preallocated color.
func Color109() *Color {
	return &color109
}

// Color110 returns a preallocated color.
func Color110() *Color {
	return &color110
}

// Color111 returns a preallocated color.
func Color111() *Color {
	return &color111
}

// Color112 returns a preallocated color.
func Color112() *Color {
	return &color112
}

// Color113 returns a preallocated color.
func Color113() *Color {
	return &color113
}

// Color114 returns a preallocated color.
func Color114() *Color {
	return &color114
}

// Color115 returns a preallocated color.
func Color115() *Color {
	return &color115
}

// Color116 returns a preallocated color.
func Color116() *Color {
	return &color116
}

// Color117 returns a preallocated color.
func Color117() *Color {
	return &color117
}

// Color118 returns a preallocated color.
func Color118() *Color {
	return &color118
}

// Color119 returns a preallocated color.
func Color119() *Color {
	return &color119
}

// Color120 returns a preallocated color.
func Color120() *Color {
	return &color120
}

// Color121 returns a preallocated color.
func Color121() *Color {
	return &color121
}

// Color122 returns a preallocated color.
func Color122() *Color {
	return &color122
}

// Color123 returns a preallocated color.
func Color123() *Color {
	return &color123
}

// Color124 returns a preallocated color.
func Color124() *Color {
	return &color124
}

// Color125 returns a preallocated color.
func Color125() *Color {
	return &color125
}

// Color126 returns a preallocated color.
func Color126() *Color {
	return &color126
}

// Color127 returns a preallocated color.
func Color127() *Color {
	return &color127
}

// Color128 returns a preallocated color.
func Color128() *Color {
	return &color128
}

// Color129 returns a preallocated color.
func Color129() *Color {
	return &color129
}

// Color130 returns a preallocated color.
func Color130() *Color {
	return &color130
}

// Color131 returns a preallocated color.
func Color131() *Color {
	return &color131
}

// Color132 returns a preallocated color.
func Color132() *Color {
	return &color132
}

// Color133 returns a preallocated color.
func Color133() *Color {
	return &color133
}

// Color134 returns a preallocated color.
func Color134() *Color {
	return &color134
}

// Color135 returns a preallocated color.
func Color135() *Color {
	return &color135
}

// Color136 returns a preallocated color.
func Color136() *Color {
	return &color136
}

// Color137 returns a preallocated color.
func Color137() *Color {
	return &color137
}

// Color138 returns a preallocated color.
func Color138() *Color {
	return &color138
}

// Color139 returns a preallocated color.
func Color139() *Color {
	return &color139
}

// Color140 returns a preallocated color.
func Color140() *Color {
	return &color140
}

// Color141 returns a preallocated color.
func Color141() *Color {
	return &color141
}

// Color142 returns a preallocated color.
func Color142() *Color {
	return &color142
}

// Color143 returns a preallocated color.
func Color143() *Color {
	return &color143
}

// Color144 returns a preallocated color.
func Color144() *Color {
	return &color144
}

// Color145 returns a preallocated color.
func Color145() *Color {
	return &color145
}

// Color146 returns a preallocated color.
func Color146() *Color {
	return &color146
}

// Color147 returns a preallocated color.
func Color147() *Color {
	return &color147
}

// Color148 returns a preallocated color.
func Color148() *Color {
	return &color148
}

// Color149 returns a preallocated color.
func Color149() *Color {
	return &color149
}

// Color150 returns a preallocated color.
func Color150() *Color {
	return &color150
}

// Color151 returns a preallocated color.
func Color151() *Color {
	return &color151
}

// Color152 returns a preallocated color.
func Color152() *Color {
	return &color152
}

// Color153 returns a preallocated color.
func Color153() *Color {
	return &color153
}

// Color154 returns a preallocated color.
func Color154() *Color {
	return &color154
}

// Color155 returns a preallocated color.
func Color155() *Color {
	return &color155
}

// Color156 returns a preallocated color.
func Color156() *Color {
	return &color156
}

// Color157 returns a preallocated color.
func Color157() *Color {
	return &color157
}

// Color158 returns a preallocated color.
func Color158() *Color {
	return &color158
}

// Color159 returns a preallocated color.
func Color159() *Color {
	return &color159
}

// Color160 returns a preallocated color.
func Color160() *Color {
	return &color160
}

// Color161 returns a preallocated color.
func Color161() *Color {
	return &color161
}

// Color162 returns a preallocated color.
func Color162() *Color {
	return &color162
}

// Color163 returns a preallocated color.
func Color163() *Color {
	return &color163
}

// Color164 returns a preallocated color.
func Color164() *Color {
	return &color164
}

// Color165 returns a preallocated color.
func Color165() *Color {
	return &color165
}

// Color166 returns a preallocated color.
func Color166() *Color {
	return &color166
}

// Color167 returns a preallocated color.
func Color167() *Color {
	return &color167
}

// Color168 returns a preallocated color.
func Color168() *Color {
	return &color168
}

// Color169 returns a preallocated color.
func Color169() *Color {
	return &color169
}

// Color170 returns a preallocated color.
func Color170() *Color {
	return &color170
}

// Color171 returns a preallocated color.
func Color171() *Color {
	return &color171
}

// Color172 returns a preallocated color.
func Color172() *Color {
	return &color172
}

// Color173 returns a preallocated color.
func Color173() *Color {
	return &color173
}

// Color174 returns a preallocated color.
func Color174() *Color {
	return &color174
}

// Color175 returns a preallocated color.
func Color175() *Color {
	return &color175
}

// Color176 returns a preallocated color.
func Color176() *Color {
	return &color176
}

// Color177 returns a preallocated color.
func Color177() *Color {
	return &color177
}

// Color178 returns a preallocated color.
func Color178() *Color {
	return &color178
}

// Color179 returns a preallocated color.
func Color179() *Color {
	return &color179
}

// Color180 returns a preallocated color.
func Color180() *Color {
	return &color180
}

// Color181 returns a preallocated color.
func Color181() *Color {
	return &color181
}

// Color182 returns a preallocated color.
func Color182() *Color {
	return &color182
}

// Color183 returns a preallocated color.
func Color183() *Color {
	return &color183
}

// Color184 returns a preallocated color.
func Color184() *Color {
	return &color184
}

// Color185 returns a preallocated color.
func Color185() *Color {
	return &color185
}

// Color186 returns a preallocated color.
func Color186() *Color {
	return &color186
}

// Color187 returns a preallocated color.
func Color187() *Color {
	return &color187
}

// Color188 returns a preallocated color.
func Color188() *Color {
	return &color188
}

// Color189 returns a preallocated color.
func Color189() *Color {
	return &color189
}

// Color190 returns a preallocated color.
func Color190() *Color {
	return &color190
}

// Color191 returns a preallocated color.
func Color191() *Color {
	return &color191
}

// Color192 returns a preallocated color.
func Color192() *Color {
	return &color192
}

// Color193 returns a preallocated color.
func Color193() *Color {
	return &color193
}

// Color194 returns a preallocated color.
func Color194() *Color {
	return &color194
}

// Color195 returns a preallocated color.
func Color195() *Color {
	return &color195
}

// Color196 returns a preallocated color.
func Color196() *Color {
	return &color196
}

// Color197 returns a preallocated color.
func Color197() *Color {
	return &color197
}

// Color198 returns a preallocated color.
func Color198() *Color {
	return &color198
}

// Color199 returns a preallocated color.
func Color199() *Color {
	return &color199
}

// Color200 returns a preallocated color.
func Color200() *Color {
	return &color200
}

// Color201 returns a preallocated color.
func Color201() *Color {
	return &color201
}

// Color202 returns a preallocated color.
func Color202() *Color {
	return &color202
}

// Color203 returns a preallocated color.
func Color203() *Color {
	return &color203
}

// Color204 returns a preallocated color.
func Color204() *Color {
	return &color204
}

// Color205 returns a preallocated color.
func Color205() *Color {
	return &color205
}

// Color206 returns a preallocated color.
func Color206() *Color {
	return &color206
}

// Color207 returns a preallocated color.
func Color207() *Color {
	return &color207
}

// Color208 returns a preallocated color.
func Color208() *Color {
	return &color208
}

// Color209 returns a preallocated color.
func Color209() *Color {
	return &color209
}

// Color210 returns a preallocated color.
func Color210() *Color {
	return &color210
}

// Color211 returns a preallocated color.
func Color211() *Color {
	return &color211
}

// Color212 returns a preallocated color.
func Color212() *Color {
	return &color212
}

// Color213 returns a preallocated color.
func Color213() *Color {
	return &color213
}

// Color214 returns a preallocated color.
func Color214() *Color {
	return &color214
}

// Color215 returns a preallocated color.
func Color215() *Color {
	return &color215
}

// Color216 returns a preallocated color.
func Color216() *Color {
	return &color216
}

// Color217 returns a preallocated color.
func Color217() *Color {
	return &color217
}

// Color218 returns a preallocated color.
func Color218() *Color {
	return &color218
}

// Color219 returns a preallocated color.
func Color219() *Color {
	return &color219
}

// Color220 returns a preallocated color.
func Color220() *Color {
	return &color220
}

// Color221 returns a preallocated color.
func Color221() *Color {
	return &color221
}

// Color222 returns a preallocated color.
func Color222() *Color {
	return &color222
}

// Color223 returns a preallocated color.
func Color223() *Color {
	return &color223
}

// Color224 returns a preallocated color.
func Color224() *Color {
	return &color224
}

// Color225 returns a preallocated color.
func Color225() *Color {
	return &color225
}

// Color226 returns a preallocated color.
func Color226() *Color {
	return &color226
}

// Color227 returns a preallocated color.
func Color227() *Color {
	return &color227
}

// Color228 returns a preallocated color.
func Color228() *Color {
	return &color228
}

// Color229 returns a preallocated color.
func Color229() *Color {
	return &color229
}

// Color230 returns a preallocated color.
func Color230() *Color {
	return &color230
}

// Color231 returns a preallocated color.
func Color231() *Color {
	return &color231
}

// Color232 returns a preallocated color.
func Color232() *Color {
	return &color232
}

// Color233 returns a preallocated color.
func Color233() *Color {
	return &color233
}

// Color234 returns a preallocated color.
func Color234() *Color {
	return &color234
}

// Color235 returns a preallocated color.
func Color235() *Color {
	return &color235
}

// Color236 returns a preallocated color.
func Color236() *Color {
	return &color236
}

// Color237 returns a preallocated color.
func Color237() *Color {
	return &color237
}

// Color238 returns a preallocated color.
func Color238() *Color {
	return &color238
}

// Color239 returns a preallocated color.
func Color239() *Color {
	return &color239
}

// Color240 returns a preallocated color.
func Color240() *Color {
	return &color240
}

// Color241 returns a preallocated color.
func Color241() *Color {
	return &color241
}

// Color242 returns a preallocated color.
func Color242() *Color {
	return &color242
}

// Color243 returns a preallocated color.
func Color243() *Color {
	return &color243
}

// Color244 returns a preallocated color.
func Color244() *Color {
	return &color244
}

// Color245 returns a preallocated color.
func Color245() *Color {
	return &color245
}

// Color246 returns a preallocated color.
func Color246() *Color {
	return &color246
}

// Color247 returns a preallocated color.
func Color247() *Color {
	return &color247
}

// Color248 returns a preallocated color.
func Color248() *Color {
	return &color248
}

// Color249 returns a preallocated color.
func Color249() *Color {
	return &color249
}

// Color250 returns a preallocated color.
func Color250() *Color {
	return &color250
}

// Color251 returns a preallocated color.
func Color251() *Color {
	return &color251
}

// Color252 returns a preallocated color.
func Color252() *Color {
	return &color252
}

// Color253 returns a preallocated color.
func Color253() *Color {
	return &color253
}

// Color254 returns a preallocated color.
func Color254() *Color {
	return &color254
}

// Color255 returns a preallocated color.
func Color255() *Color {
	return &color255
}

// ColorAliceBlue returns a preallocated color.
func ColorAliceBlue() *Color {
	return &colorAliceBlue
}

// ColorAntiqueWhite returns a preallocated color.
func ColorAntiqueWhite() *Color {
	return &colorAntiqueWhite
}

// ColorAquaMarine returns a preallocated color.
func ColorAquaMarine() *Color {
	return &colorAquaMarine
}

// ColorAzure returns a preallocated color.
func ColorAzure() *Color {
	return &colorAzure
}

// ColorBeige returns a preallocated color.
func ColorBeige() *Color {
	return &colorBeige
}

// ColorBisque returns a preallocated color.
func ColorBisque() *Color {
	return &colorBisque
}

// ColorBlanchedAlmond returns a preallocated color.
func ColorBlanchedAlmond() *Color {
	return &colorBlanchedAlmond
}

// ColorBlueViolet returns a preallocated color.
func ColorBlueViolet() *Color {
	return &colorBlueViolet
}

// ColorBrown returns a preallocated color.
func ColorBrown() *Color {
	return &colorBrown
}

// ColorBurlyWood returns a preallocated color.
func ColorBurlyWood() *Color {
	return &colorBurlyWood
}

// ColorCadetBlue returns a preallocated color.
func ColorCadetBlue() *Color {
	return &colorCadetBlue
}

// ColorChartreuse returns a preallocated color.
func ColorChartreuse() *Color {
	return &colorChartreuse
}

// ColorChocolate returns a preallocated color.
func ColorChocolate() *Color {
	return &colorChocolate
}

// ColorCoral returns a preallocated color.
func ColorCoral() *Color {
	return &colorCoral
}

// ColorCornflowerBlue returns a preallocated color.
func ColorCornflowerBlue() *Color {
	return &colorCornflowerBlue
}

// ColorCornsilk returns a preallocated color.
func ColorCornsilk() *Color {
	return &colorCornsilk
}

// ColorCrimson returns a preallocated color.
func ColorCrimson() *Color {
	return &colorCrimson
}

// ColorDarkBlue returns a preallocated color.
func ColorDarkBlue() *Color {
	return &colorDarkBlue
}

// ColorDarkCyan returns a preallocated color.
func ColorDarkCyan() *Color {
	return &colorDarkCyan
}

// ColorDarkGoldenrod returns a preallocated color.
func ColorDarkGoldenrod() *Color {
	return &colorDarkGoldenrod
}

// ColorDarkGray returns a preallocated color.
func ColorDarkGray() *Color {
	return &colorDarkGray
}

// ColorDarkGreen returns a preallocated color.
func ColorDarkGreen() *Color {
	return &colorDarkGreen
}

// ColorDarkKhaki returns a preallocated color.
func ColorDarkKhaki() *Color {
	return &colorDarkKhaki
}

// ColorDarkMagenta returns a preallocated color.
func ColorDarkMagenta() *Color {
	return &colorDarkMagenta
}

// ColorDarkOliveGreen returns a preallocated color.
func ColorDarkOliveGreen() *Color {
	return &colorDarkOliveGreen
}

// ColorDarkOrange returns a preallocated color.
func ColorDarkOrange() *Color {
	return &colorDarkOrange
}

// ColorDarkOrchid returns a preallocated color.
func ColorDarkOrchid() *Color {
	return &colorDarkOrchid
}

// ColorDarkRed returns a preallocated color.
func ColorDarkRed() *Color {
	return &colorDarkRed
}

// ColorDarkSalmon returns a preallocated color.
func ColorDarkSalmon() *Color {
	return &colorDarkSalmon
}

// ColorDarkSeaGreen returns a preallocated color.
func ColorDarkSeaGreen() *Color {
	return &colorDarkSeaGreen
}

// ColorDarkSlateBlue returns a preallocated color.
func ColorDarkSlateBlue() *Color {
	return &colorDarkSlateBlue
}

// ColorDarkSlateGray returns a preallocated color.
func ColorDarkSlateGray() *Color {
	return &colorDarkSlateGray
}

// ColorDarkTurquoise returns a preallocated color.
func ColorDarkTurquoise() *Color {
	return &colorDarkTurquoise
}

// ColorDarkViolet returns a preallocated color.
func ColorDarkViolet() *Color {
	return &colorDarkViolet
}

// ColorDeepPink returns a preallocated color.
func ColorDeepPink() *Color {
	return &colorDeepPink
}

// ColorDeepSkyBlue returns a preallocated color.
func ColorDeepSkyBlue() *Color {
	return &colorDeepSkyBlue
}

// ColorDimGray returns a preallocated color.
func ColorDimGray() *Color {
	return &colorDimGray
}

// ColorDodgerBlue returns a preallocated color.
func ColorDodgerBlue() *Color {
	return &colorDodgerBlue
}

// ColorFireBrick returns a preallocated color.
func ColorFireBrick() *Color {
	return &colorFireBrick
}

// ColorFloralWhite returns a preallocated color.
func ColorFloralWhite() *Color {
	return &colorFloralWhite
}

// ColorForestGreen returns a preallocated color.
func ColorForestGreen() *Color {
	return &colorForestGreen
}

// ColorGainsboro returns a preallocated color.
func ColorGainsboro() *Color {
	return &colorGainsboro
}

// ColorGhostWhite returns a preallocated color.
func ColorGhostWhite() *Color {
	return &colorGhostWhite
}

// ColorGold returns a preallocated color.
func ColorGold() *Color {
	return &colorGold
}

// ColorGoldenrod returns a preallocated color.
func ColorGoldenrod() *Color {
	return &colorGoldenrod
}

// ColorGreenYellow returns a preallocated color.
func ColorGreenYellow() *Color {
	return &colorGreenYellow
}

// ColorHoneydew returns a preallocated color.
func ColorHoneydew() *Color {
	return &colorHoneydew
}

// ColorHotPink returns a preallocated color.
func ColorHotPink() *Color {
	return &colorHotPink
}

// ColorIndianRed returns a preallocated color.
func ColorIndianRed() *Color {
	return &colorIndianRed
}

// ColorIndigo returns a preallocated color.
func ColorIndigo() *Color {
	return &colorIndigo
}

// ColorIvory returns a preallocated color.
func ColorIvory() *Color {
	return &colorIvory
}

// ColorKhaki returns a preallocated color.
func ColorKhaki() *Color {
	return &colorKhaki
}

// ColorLavender returns a preallocated color.
func ColorLavender() *Color {
	return &colorLavender
}

// ColorLavenderBlush returns a preallocated color.
func ColorLavenderBlush() *Color {
	return &colorLavenderBlush
}

// ColorLawnGreen returns a preallocated color.
func ColorLawnGreen() *Color {
	return &colorLawnGreen
}

// ColorLemonChiffon returns a preallocated color.
func ColorLemonChiffon() *Color {
	return &colorLemonChiffon
}

// ColorLightBlue returns a preallocated color.
func ColorLightBlue() *Color {
	return &colorLightBlue
}

// ColorLightCoral returns a preallocated color.
func ColorLightCoral() *Color {
	return &colorLightCoral
}

// ColorLightCyan returns a preallocated color.
func ColorLightCyan() *Color {
	return &colorLightCyan
}

// ColorLightGoldenrodYellow returns a preallocated color.
func ColorLightGoldenrodYellow() *Color {
	return &colorLightGoldenrodYellow
}

// ColorLightGray returns a preallocated color.
func ColorLightGray() *Color {
	return &colorLightGray
}

// ColorLightGreen returns a preallocated color.
func ColorLightGreen() *Color {
	return &colorLightGreen
}

// ColorLightPink returns a preallocated color.
func ColorLightPink() *Color {
	return &colorLightPink
}

// ColorLightSalmon returns a preallocated color.
func ColorLightSalmon() *Color {
	return &colorLightSalmon
}

// ColorLightSeaGreen returns a preallocated color.
func ColorLightSeaGreen() *Color {
	return &colorLightSeaGreen
}

// ColorLightSkyBlue returns a preallocated color.
func ColorLightSkyBlue() *Color {
	return &colorLightSkyBlue
}

// ColorLightSlateGray returns a preallocated color.
func ColorLightSlateGray() *Color {
	return &colorLightSlateGray
}

// ColorLightSteelBlue returns a preallocated color.
func ColorLightSteelBlue() *Color {
	return &colorLightSteelBlue
}

// ColorLightYellow returns a preallocated color.
func ColorLightYellow() *Color {
	return &colorLightYellow
}

// ColorLimeGreen returns a preallocated color.
func ColorLimeGreen() *Color {
	return &colorLimeGreen
}

// ColorLinen returns a preallocated color.
func ColorLinen() *Color {
	return &colorLinen
}

// ColorMediumAquamarine returns a preallocated color.
func ColorMediumAquamarine() *Color {
	return &colorMediumAquamarine
}

// ColorMediumBlue returns a preallocated color.
func ColorMediumBlue() *Color {
	return &colorMediumBlue
}

// ColorMediumOrchid returns a preallocated color.
func ColorMediumOrchid() *Color {
	return &colorMediumOrchid
}

// ColorMediumPurple returns a preallocated color.
func ColorMediumPurple() *Color {
	return &colorMediumPurple
}

// ColorMediumSeaGreen returns a preallocated color.
func ColorMediumSeaGreen() *Color {
	return &colorMediumSeaGreen
}

// ColorMediumSlateBlue returns a preallocated color.
func ColorMediumSlateBlue() *Color {
	return &colorMediumSlateBlue
}

// ColorMediumSpringGreen returns a preallocated color.
func ColorMediumSpringGreen() *Color {
	return &colorMediumSpringGreen
}

// ColorMediumTurquoise returns a preallocated color.
func ColorMediumTurquoise() *Color {
	return &colorMediumTurquoise
}

// ColorMediumVioletRed returns a preallocated color.
func ColorMediumVioletRed() *Color {
	return &colorMediumVioletRed
}

// ColorMidnightBlue returns a preallocated color.
func ColorMidnightBlue() *Color {
	return &colorMidnightBlue
}

// ColorMintCream returns a preallocated color.
func ColorMintCream() *Color {
	return &colorMintCream
}

// ColorMistyRose returns a preallocated color.
func ColorMistyRose() *Color {
	return &colorMistyRose
}

// ColorMoccasin returns a preallocated color.
func ColorMoccasin() *Color {
	return &colorMoccasin
}

// ColorNavajoWhite returns a preallocated color.
func ColorNavajoWhite() *Color {
	return &colorNavajoWhite
}

// ColorOldLace returns a preallocated color.
func ColorOldLace() *Color {
	return &colorOldLace
}

// ColorOliveDrab returns a preallocated color.
func ColorOliveDrab() *Color {
	return &colorOliveDrab
}

// ColorOrange returns a preallocated color.
func ColorOrange() *Color {
	return &colorOrange
}

// ColorOrangeRed returns a preallocated color.
func ColorOrangeRed() *Color {
	return &colorOrangeRed
}

// ColorOrchid returns a preallocated color.
func ColorOrchid() *Color {
	return &colorOrchid
}

// ColorPaleGoldenrod returns a preallocated color.
func ColorPaleGoldenrod() *Color {
	return &colorPaleGoldenrod
}

// ColorPaleGreen returns a preallocated color.
func ColorPaleGreen() *Color {
	return &colorPaleGreen
}

// ColorPaleTurquoise returns a preallocated color.
func ColorPaleTurquoise() *Color {
	return &colorPaleTurquoise
}

// ColorPaleVioletRed returns a preallocated color.
func ColorPaleVioletRed() *Color {
	return &colorPaleVioletRed
}

// ColorPapayaWhip returns a preallocated color.
func ColorPapayaWhip() *Color {
	return &colorPapayaWhip
}

// ColorPeachPuff returns a preallocated color.
func ColorPeachPuff() *Color {
	return &colorPeachPuff
}

// ColorPeru returns a preallocated color.
func ColorPeru() *Color {
	return &colorPeru
}

// ColorPink returns a preallocated color.
func ColorPink() *Color {
	return &colorPink
}

// ColorPlum returns a preallocated color.
func ColorPlum() *Color {
	return &colorPlum
}

// ColorPowderBlue returns a preallocated color.
func ColorPowderBlue() *Color {
	return &colorPowderBlue
}

// ColorRebeccaPurple returns a preallocated color.
func ColorRebeccaPurple() *Color {
	return &colorRebeccaPurple
}

// ColorRosyBrown returns a preallocated color.
func ColorRosyBrown() *Color {
	return &colorRosyBrown
}

// ColorRoyalBlue returns a preallocated color.
func ColorRoyalBlue() *Color {
	return &colorRoyalBlue
}

// ColorSaddleBrown returns a preallocated color.
func ColorSaddleBrown() *Color {
	return &colorSaddleBrown
}

// ColorSalmon returns a preallocated color.
func ColorSalmon() *Color {
	return &colorSalmon
}

// ColorSandyBrown returns a preallocated color.
func ColorSandyBrown() *Color {
	return &colorSandyBrown
}

// ColorSeaGreen returns a preallocated color.
func ColorSeaGreen() *Color {
	return &colorSeaGreen
}

// ColorSeashell returns a preallocated color.
func ColorSeashell() *Color {
	return &colorSeashell
}

// ColorSienna returns a preallocated color.
func ColorSienna() *Color {
	return &colorSienna
}

// ColorSkyblue returns a preallocated color.
func ColorSkyblue() *Color {
	return &colorSkyblue
}

// ColorSlateBlue returns a preallocated color.
func ColorSlateBlue() *Color {
	return &colorSlateBlue
}

// ColorSlateGray returns a preallocated color.
func ColorSlateGray() *Color {
	return &colorSlateGray
}

// ColorSnow returns a preallocated color.
func ColorSnow() *Color {
	return &colorSnow
}

// ColorSpringGreen returns a preallocated color.
func ColorSpringGreen() *Color {
	return &colorSpringGreen
}

// ColorSteelBlue returns a preallocated color.
func ColorSteelBlue() *Color {
	return &colorSteelBlue
}

// ColorTan returns a preallocated color.
func ColorTan() *Color {
	return &colorTan
}

// ColorThistle returns a preallocated color.
func ColorThistle() *Color {
	return &colorThistle
}

// ColorTomato returns a preallocated color.
func ColorTomato() *Color {
	return &colorTomato
}

// ColorTurquoise returns a preallocated color.
func ColorTurquoise() *Color {
	return &colorTurquoise
}

// ColorViolet returns a preallocated color.
func ColorViolet() *Color {
	return &colorViolet
}

// ColorWheat returns a preallocated color.
func ColorWheat() *Color {
	return &colorWheat
}

// ColorWhiteSmoke returns a preallocated color.
func ColorWhiteSmoke() *Color {
	return &colorWhiteSmoke
}

// ColorYellowGreen returns a preallocated color.
func ColorYellowGreen() *Color {
	return &colorYellowGreen
}
