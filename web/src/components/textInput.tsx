import styled from "utils/styled"

export default styled.input`
  width: 100%;
  border: 2px solid ${props => props.theme.color.grey};
  border-radius: 4px;
  margin-top: 24px;
  padding: 14px;
  font-size: 18px;
  background: ${props => props.theme.color.greyLight};

  &:focus {
    outline: none;
  }
`
