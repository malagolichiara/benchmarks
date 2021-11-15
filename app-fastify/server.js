const cluster = require('cluster');
const numCPUs = require('os').cpus().length;
const fastify = require('fastify')({ logger: false })
fastify.register(require('fastify-formbody'))
const port = 8000;

var values = {}

if (cluster.isMaster) {
  console.log(`Master ${process.pid} is running`);
  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }
  cluster.on('exit', worker => {
    console.log(`Worker ${worker.process.pid} died`);
  });
} else {
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

  fastify.listen(port, () => {
    console.log(`Fastify listening on port ${port}, PID: ${process.pid}`);
  });

  // Run the server!
  // const start = async () => {
  //   try {
  //     await fastify.listen(8000)
  //   } catch (err) {
  //     fastify.log.error(err)
  //     process.exit(1)
  //   }
  // }

  // start()
}