import React, { useMemo } from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames/bind';
import style from './FormSection.module.scss';

const cx = classnames.bind(style);

function FormSection({
  name, description, className, children,
}) {
  return (
    <div className={cx('root', className)}>
      <div className={cx('left-pane')}>
        <h3 className={cx('title')}>{name}</h3>
        {description && (
          <p className={cx('description')}>
            {description}
          </p>
        )}
      </div>
      <div className={cx('right-pane')}>
        {children}
      </div>
    </div>
  );
}

FormSection.propTypes = {
  name: PropTypes.string,
  description: PropTypes.string,
  className: PropTypes.string,
  children: PropTypes.node,
};

export default FormSection;
