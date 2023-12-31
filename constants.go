package captcha

const (
	characters               string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	mathMinDefault           uint8  = 1
	mathMaxDefault           uint16 = 9
	widthDefault             uint16 = 150
	heightDefault            uint16 = 50
	noiseDefault             uint8  = 1
	sizeDefault              uint8  = 4
	fontSizeDefault          uint8  = 12
	ratioFontSize            uint8  = 4
	noiseGreyColorMinDefault uint8  = 1
	noiseGreyColorMaxDefault uint8  = 9
	noiseGreyColorMinInverse uint8  = 7
	noiseGreyColorMaxInverse uint8  = 15
	paddingHorizontalDefault uint8  = 5
	fillColorMinDefault      uint8  = 0
	fillColorMaxDefault      uint8  = 4
	fillColorMinInverse      uint8  = 10
	fillColorMaxInverse      uint8  = 14
	regexColor               string = "^#(?:[0-9a-fA-F]{3}){1,2}$"
)
