import React from "react"
import { addDecorator } from "@storybook/react"
import ThemeProvider from "~/style/themeProvider"

addDecorator(storyFn => (
  <ThemeProvider>
    {storyFn()}
  </ThemeProvider>
))
