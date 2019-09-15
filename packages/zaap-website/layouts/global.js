import React from 'react'
import { ThemeProvider } from 'emotion-theming'

const theme = {
  fontFamily: 'Montserrat',
  fontWeightMedium: 400,
  fontWeightSemiBold: 500,
  fontWeightBold: 600,
  colorPrimary: '#0069FF',
  colorRed: '#D32F2F',
  colorRedDark: '#9A0007',
  colorGreen: '#15CD72',
  colorGreenDark: '#0CB863',
  colorGrey: '#E3EBF6',
  colorGreyDark: '#95AAC9',
  colorGreyLight: '#F9FBFD',
  colorTextPrimary: '#000000',
  colorTextSecondary: '#6C757D',
}

const GlobalLayout = ({ children }) => (
  <ThemeProvider theme={theme}>
    {children}
  </ThemeProvider>
)

export default GlobalLayout
