package palenight

import (
	"image/color"

	"barista.run/colors"
)

func getColors() palenightColors {
	return palenightColors{
		red:          fromHex("#ff5370"),
		lightRed:     fromHex("#ff869a"),
		darkRed:      fromHex("#BE5046"),
		green:        fromHex("#C3E88D"),
		yellow:       fromHex("#ffcb6b"),
		darkYellow:   fromHex("#F78C6C"),
		blue:         fromHex("#82b1ff"),
		purple:       fromHex("#c792ea"),
		bluePurple:   fromHex("#939ede"),
		cyan:         fromHex("#89DDFF"),
		white:        fromHex("#bfc7d5"),
		black:        fromHex("#292D3E"),
		commentGrey:  fromHex("#697098"),
		gutterFgGrey: fromHex("#4B5263"),
		cursorGrey:   fromHex("#2C323C"),
		visualGrey:   fromHex("#3E4452"),
		menuGrey:     fromHex("#3E4452"),
		specialGrey:  fromHex("#3B4048"),
		vertsplit:    fromHex("#181A1F"),
		whiteMask1:   fromHex("#333747"),
		whiteMask3:   fromHex("#474b59"),
		whiteMask11:  fromHex("#989aa2"),

		matBlackRegular:   fromHex("#292d3e"),
		matBlackLight:     fromHex("#434758"),
		matRedRegular:     fromHex("#f07178"),
		matRedLight:       fromHex("#ff8b92"),
		matGreenRegular:   fromHex("#c3e88d"),
		matGreenLight:     fromHex("#ddffa7"),
		matYellowRegular:  fromHex("#ffcb6b"),
		matYellowLight:    fromHex("#ffe585"),
		matBlueRegular:    fromHex("#82aaff"),
		matBlueLight:      fromHex("#9cc4ff"),
		matMagentaRegular: fromHex("#c792ea"),
		matMagentaLight:   fromHex("#e1acff"),
		matCyanRegular:    fromHex("#89ddff"),
		matCyanLight:      fromHex("#a3f7ff"),
		matWhiteRegular:   fromHex("#d0d0d0"),
		matWhiteLight:     fromHex("#ffffff"),
	}
}

type palenightColor color.Color

func fromHex(hex string) palenightColor {
	return palenightColor(colors.Hex(hex))
}

type palenightColors struct {
	red          palenightColor
	lightRed     palenightColor
	darkRed      palenightColor
	green        palenightColor
	yellow       palenightColor
	darkYellow   palenightColor
	blue         palenightColor
	purple       palenightColor
	bluePurple   palenightColor
	cyan         palenightColor
	white        palenightColor
	black        palenightColor
	commentGrey  palenightColor
	gutterFgGrey palenightColor
	cursorGrey   palenightColor
	visualGrey   palenightColor
	menuGrey     palenightColor
	specialGrey  palenightColor
	vertsplit    palenightColor
	whiteMask1   palenightColor
	whiteMask3   palenightColor
	whiteMask11  palenightColor

	matBlackRegular   palenightColor
	matBlackLight     palenightColor
	matRedRegular     palenightColor
	matRedLight       palenightColor
	matGreenRegular   palenightColor
	matGreenLight     palenightColor
	matYellowRegular  palenightColor
	matYellowLight    palenightColor
	matBlueRegular    palenightColor
	matBlueLight      palenightColor
	matMagentaRegular palenightColor
	matMagentaLight   palenightColor
	matCyanRegular    palenightColor
	matCyanLight      palenightColor
	matWhiteRegular   palenightColor
	matWhiteLight     palenightColor
}
