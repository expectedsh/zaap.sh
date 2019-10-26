import styled from "utils/styled"

export const Root = styled.div`
  margin: 42px 0;
`

export const Logo = styled.img`
  display: block;
  margin: 0 auto 32px;
  height: 52px;
`

export const Container = styled.div`
  margin: 0 auto;
  padding: 48px 32px 32px;
  border-radius: 4px;
  max-width: 450px;
  background: ${props => props.theme.color.white};

  @media screen and (max-width: ${props => props.theme.breakpoint.sm}) {
    max-width: 100%;
    border-radius: 0;
  }
`

export const Title = styled.div`
  font-weight: ${props => props.theme.text.fontWeightThin};
  // font-size: ${props => props.theme.text.fontSizeLarge};
  font-size: 2rem;
  text-align: center;
  margin-bottom: 32px;
`

export const OAuthText = styled.div`
  text-align: center;
  margin: 24px 0;
  text-transform: uppercase;
  font-size: 0.875rem;
  font-weight: 500;
  color: #555555;
`

export const OAuthButtonGroup = styled.div`
  display: flex;
  flex-direction: row;

  @media screen and (max-width: ${props => props.theme.breakpoint.sm}) {
    flex-direction: column;
  }
`

export const OAuthButton = styled.a`
  position: relative;
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  padding: 12px 16px 12px 40px;
  border: 2px solid ${props => props.theme.color.grey};
  border-radius: 3px;
  font-size: 18px;
  font-weight: ${props => props.theme.text.fontWeightRegular};
  color: ${props => props.theme.color.textNormal};
  text-decoration: none;
  cursor: pointer;

  & + & {
    margin-left: 16px;
  }

  &:hover {
    background: ${props => props.theme.color.grey};
  }

  img {
    position: absolute;
    left: 16px;
  }

  @media screen and (max-width: ${props => props.theme.breakpoint.sm}) {
    & + & {
      margin-left: 0;
      margin-top: 24px;
    }
  }
`

export const AlternativeLink = styled.div`
  margin: 24px 0;
  text-align: center;
  color: ${props => props.theme.color.white};

  a {
    color: ${props => props.theme.color.white};
  }
`
