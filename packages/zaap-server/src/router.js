import Router from 'koa-router'
import graphqlHTTP from 'koa-graphql'
import { makeExecutableSchema, SchemaDirectiveVisitor } from 'graphql-tools'
import { promises as fs } from 'fs'
import { join } from 'path'
import knex from 'knex'

const router = new Router()

class ExecuteQueryDirective extends SchemaDirectiveVisitor {
  visitFieldDefinition(field) {
    const { query } = this.args
    field.resolve = async (_, args) => {
      const data = await knex({
        client: 'postgres',
        connection: {
          host: process.env.DB_HOST,
          user: process.env.DB_USER,
          password: process.env.DB_PASSWORD,
          database: process.env.DB_NAME,
        },
      }).raw(query, args)
      return field.astNode.type.type && field.astNode.type.type.kind === 'ListType'
        ? data.rows : data.rows[0]
    }
  }
}

class FieldDirective extends SchemaDirectiveVisitor {
  visitFieldDefinition(field) {
    const { from } = this.args
    field.resolve = root => root[from]
  }
}

router.all('/hello-api', async (ctx, next) => {
  const schema = makeExecutableSchema({
    typeDefs: await fs.readFile(join(__dirname, '..', '..', '..', 'schema.graphql'), 'utf8'),
    schemaDirectives: {
      executeQuery: ExecuteQueryDirective,
      field: FieldDirective,
    },
  })

  return graphqlHTTP({
    schema,
    graphiql: true,
  })(ctx, next)
})

export default router
