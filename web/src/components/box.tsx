import React from "react"
import { css } from "@emotion/core"
import { AxiosError } from "axios"
import styled from "utils/styled"

type Props =
  { variant: "success" | "warning", message: string } |
  { variant: "error", error: AxiosError | Error }

const Wrapper = styled.div<Props>`
  border-radius: 4px;
  padding: 16px;

  ${props => props.variant === "error" && css`
    background: ${props.theme.color.red};
    color: ${props.theme.color.white};
  `}
`

const Box = (props: Props) => {
  const getMessage = () => {
    if (props.variant === "error") {
      if ((props.error as AxiosError).isAxiosError) {
        return (props.error as any).response.data.error.message
      }
      return props.error.message
    }
    return props.message
  }
  return (
    <Wrapper {...props}>
      {getMessage()}
    </Wrapper>
  )
}

export default Box
