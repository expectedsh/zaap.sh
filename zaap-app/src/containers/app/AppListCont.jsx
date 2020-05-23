import React from 'react'
import AppList from '~/components/pages/app/AppList'

function AppListCont() {
  return <AppList loading={false} error={undefined} applications={[]} />
}

export default AppListCont
