import { ApolloServer } from 'apollo-server-koa'
import { readdirSync, readFileSync } from 'fs'
import { join } from 'path'
import Koa from 'koa'
import bodyParser from 'koa-bodyparser'
import cors from 'koa-cors'
import router from './router'
import resolvers from './resolvers'

const app = new Koa()

app.use(bodyParser())
app.use(cors())
app.use(router.routes())

const SCHEMA_DIRECTORY = join(__dirname, '..', 'schema')
const server = new ApolloServer({
  typeDefs: readdirSync(SCHEMA_DIRECTORY)
    .map(filename => readFileSync(join(SCHEMA_DIRECTORY, filename), 'utf8')),
  resolvers,
})

server.applyMiddleware({
  app,
  path: '/graphql',
  disableHealthCheck: true,
})

export default app
