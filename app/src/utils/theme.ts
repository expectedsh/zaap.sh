export interface Theme {
  breakpoint: {
    xs: string
    sm: string
    md: string
    lg: string
    xl: string
  }
  color: {
    white: string

    greyLight: string
    grey: string
    greyDark: string

    greenLight: string
    green: string
    greenDark: string

    redLight: string
    red: string
    redDark: string

    blueLight: string
    blue: string
    blueDark: string

    textNormal: string
    textMuted: string
  },
  text: {
    fontFamily: string

    lineHeightHeading: string
    lineHeightBody: string

    fontSizeExtraSmall: string
    fontSizeSmall: string
    fontSizeMedium: string
    fontSizeLarge: string
    fontSizeExtraLarge: string

    fontWeightThin: string
    fontWeightRegular: string
    fontWeightSemiBold: string
    fontWeightBold: string
  }
}

export const theme: Theme = {
  breakpoint: {
    xs: "0",
    sm: "576px",
    md: "768px",
    lg: "992px",
    xl: "1200px",
  },
  color: {
    white: "#FFF",

    greyLight: "#F9FBFD",
    grey: "#E3EBF6",
    greyDark: "#95AAC9",

    redLight: "#FF7B72",
    red: "#F24646",
    redDark: "#B8001E",

    greenLight: "#56DD7F",
    green: "#04AA51",
    greenDark: "#007925",

    blueLight: "#0065FF",
    blue: "#0052CC",
    blueDark: "#05367F",

    textNormal: "#212529",
    textMuted: "#6C757D",
  },
  text: {
    fontFamily: "'Lato', Arial, Helvetica, sans-serif",

    lineHeightHeading: "1.15",
    lineHeightBody: "1.4",

    fontSizeExtraSmall: "11px",
    fontSizeSmall: "12px",
    fontSizeMedium: "14px",
    fontSizeLarge: "16px",
    fontSizeExtraLarge: "18px",

    fontWeightThin: "300",
    fontWeightRegular: "400",
    fontWeightSemiBold: "500",
    fontWeightBold: "600",
  }
}
