import React from "react"
import PropTypes from "prop-types"
import classnames from "classnames/bind"
import style from "./Table.module.scss"

const cx = classnames.bind(style)

function Table({ config, dataSource, onRowClick }) {
  return (
    <div className={cx("root")}>
      <div className={cx("header")}>
        {config.map((item, index) => (
          <div key={index} className={cx("cell", item.cellClassName)}>
            {item.renderHeader()}
          </div>
        ))}
      </div>
      <div>
        {dataSource.map((row, rIndex) => (
          <div key={rIndex} className={cx("row")} onClick={e => onRowClick?.(row, e)}>
            {config.map((item, index) => (
              <div key={index} className={cx("cell", item.cellClassName)}>
                {item.renderCell(row)}
              </div>
            ))}
          </div>
        ))}
      </div>
    </div>
  )
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
}

export default Table
