import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames/bind';
import ApplicationStatusName from '~/oldcomponents/ApplicationStatusName';
import style from './ApplicationStatusBadge.module.scss';

const cx = classnames.bind(style);

function getClassName(status) {
  switch (status) {
    case 'deploying':
    case 'running':
      return 'green';
    case 'crashed':
    case 'failed':
      return 'red';
    default:
      return 'grey';
  }
}

function ApplicationStatusBadge({ status }) {
  return (
    <div className={cx('root', getClassName(status))}>
      <ApplicationStatusName status={status} />
    </div>
  );
}

ApplicationStatusBadge.propTypes = {
  status: PropTypes.string,
};

export default ApplicationStatusBadge;
