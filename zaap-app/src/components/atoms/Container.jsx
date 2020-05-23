import styled from '@emotion/styled'
import { css } from '@emotion/core'

const Container = styled.div((props) => css`
  width: 100%;
  padding: 0 16px;
  margin: 0 auto;

  @media (min-width: ${props.theme.breakpoints.small}) {
    max-width: 540px;
  }

  @media (min-width: ${props.theme.breakpoints.medium}) {
    max-width: 720px;
  }

  @media (min-width: ${props.theme.breakpoints.large}) {
    max-width: 960px;
  }

  @media (min-width: ${props.theme.breakpoints.extraLarge}) {
    max-width: 1140px;
  }
`)

export default Container
