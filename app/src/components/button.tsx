import { css } from "@emotion/core"
import styled from "utils/styled"

interface Props {
  variant?: "success"
}

export default styled.button<Props>`
  margin-top: 24px;
  display: block;
  width: 100%;
  border: none;
  border-radius: 4px;
  font-size: 1.125rem;
  font-weight: 500;
  padding: 14px 0;
  cursor: pointer;

  ${props => props.variant === "success" && css`
    color: ${props.theme.color.white};
    background: ${props.theme.color.greenLight};
    &:hover {
      background: ${props.theme.color.green};
    }
  `}
`
