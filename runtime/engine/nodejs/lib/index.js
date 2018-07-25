const path = require('path')

const grpc = require('grpc')

const Request = require('./request')
const Response = require('./response')

const PORT = process.env._RUNTIME_PORT

const proto = grpc.load(path.join(__dirname, '/../proto/internal.proto'))

const server = new grpc.Server()

exports.start = function (handler) {
  server.addService(proto.runtime.Internal.service, {
    pipe: async (conn) => {
      console.log(conn.request)
      const req = new Request(conn.request)
      const res = new Response(conn)

      try {
        await handler(req, res)
      } catch (err) {
        console.log(err)
      } finally {
        res.close()
      }
    },
    healthz: (call, callback) => {
      callback(null, {
        status: 1
      })
    }
  })
  server.bind(`localhost:${PORT}`, grpc.ServerCredentials.createInsecure())
  server.start()
}

process.on('SIGTERM', () => {
  server.tryShutdown(() => {
    console.log('shutdown')
  })
})
