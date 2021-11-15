// Require the framework and instantiate it
const fastify = require('fastify')({ logger: false })
fastify.register(require('fastify-formbody'))

var values = {}

// Declare a route
fastify.get('/', async (request, reply) => {
  return { hello: 'world' }
})

fastify.get('/items/:idx', async (request, reply) => {
  return {
    value: values[request.params.idx]
  }
})

fastify.post('/items/:idx', async (request, reply) => {
  values[request.params.idx] = request.body.value
  return {
    value: values[request.params.idx]
  }
})

// Run the server!
const start = async () => {
  try {
    await fastify.listen(8000)
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}

start()
