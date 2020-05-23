import './Button.scss';
import React from 'react';
import PropTypes from 'prop-types';

function Button({
  loading, children, disabled, ...props
}) {
  return (
    <button {...props} disabled={loading || disabled}>
      {children}
    </button>
  );
}

Button.propTypes = {
  loading: PropTypes.bool,
  className: PropTypes.string,
  type: PropTypes.string,
  onClick: PropTypes.func,
};

export default Button;
