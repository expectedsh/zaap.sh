import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames/bind';
import style from './Table.module.scss';

const cx = classnames.bind(style);

function Table({
  config, dataSource, onRowClick, noData = null,
}) {
  function onClick(row) {
    return (event) => {
      if (!onRowClick) {
        return;
      }
      if (['a', 'button'].includes(event.nativeEvent.path[0].tagName.toLowerCase())) {
        return;
      }
      onRowClick(row, event);
    };
  }

  return (
    <div className={cx('root')}>
      {dataSource.length ? (
        <>
          <div className={cx('header')}>
            {config.map((item, index) => (
              <div key={index} className={cx('cell', item.cellClassName)}>
                {item.renderHeader()}
              </div>
            ))}
          </div>
          <div>
            {dataSource.map((row, rIndex) => (
              <div
                key={rIndex}
                className={cx('row')}
                onClick={onClick(row)}
                style={onRowClick ? { cursor: 'pointer' } : {}}
              >
                {config.map((item, index) => (
                  <div key={index} className={cx('cell', item.cellClassName)}>
                    {item.renderCell(row)}
                  </div>
                ))}
              </div>
            ))}
          </div>
        </>
      ) : noData}
    </div>
  );
}

Table.propTypes = {
  config: PropTypes.arrayOf(
    PropTypes.shape({
      renderHeader: PropTypes.func.isRequired,
      renderCell: PropTypes.func.isRequired,
      cellClassName: PropTypes.string,
    }),
  ).isRequired,
  dataSource: PropTypes.array.isRequired,
  onRowClick: PropTypes.func,
  noData: PropTypes.node,
};

export default Table;
