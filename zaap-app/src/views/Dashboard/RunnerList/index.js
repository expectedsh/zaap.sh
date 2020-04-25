import React, { useEffect } from "react"
import { Link, useHistory } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import moment from "moment"
import { fetchRunners } from "~/store/runners/actions"
import Alert from "~/components/Alert"
import RunnerStateBadge from "~/components/RunnerStatusBadge"
import Header from "~/components/Header"
import Table from "~/components/Table"
import style from "./RunnerList.module.scss"
import RunnerTypeLogo from "~/components/RunnerTypeLogo"

const cx = classnames.bind(style)

const tableConfig = [
  {
    renderHeader: () => "Name",
    renderCell: runner => (
      <>
        <RunnerTypeLogo type={runner.type} className={cx('runner-logo')}/>
        {runner.name}
      </>
    ),
    cellClassName: cx("cell-name"),
  },
  {
    renderHeader: () => "Status",
    renderCell: runner => <RunnerStateBadge status={runner.status}/>,
    cellClassName: cx("cell-state"),
  },
  {
    renderHeader: () => "Endpoint",
    renderCell: runner => runner.url,
    cellClassName: cx("cell-url"),
  },
  {
    renderHeader: () => "Last update",
    renderCell: runner => moment(runner.updatedAt).fromNow(),
    cellClassName: cx("cell-created"),
  },
]

function RunnerList() {
  const dispatch = useDispatch()
  const history = useHistory()
  const { pending, runners, error } = useSelector(state => state.runners)

  useEffect(() => {
    dispatch(fetchRunners())
  }, [])

  function renderBody() {
    if (pending) {
      return <div>Loading...</div>
    }
    if (error) {
      return <Alert className="alert alert-error" error={error}/>
    }
    return runners ? (
      <Table
        config={tableConfig}
        dataSource={runners}
        noData={
          <div className={cx('no-runner')}>
            <div className={cx('title')}>
              You don't have runner 😭
            </div>
            <div className={cx('description')}>
              Register a runner and it will show up here.
            </div>
          </div>
        }
      />
    ) : null
  }

  return (
    <>
      <Header preTitle="Overview" title="Runners">
        <Link className="btn btn-secondary" to="/runners/new">
          Register runner
        </Link>
      </Header>
      <div className="container">
        {renderBody()}
      </div>
    </>
  )
}

export default RunnerList
