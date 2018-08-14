const grpc = require('grpc')
const proto = grpc.load(`${__dirname}/../../proto/runtime.proto`)

const server = new grpc.Server()

server.addService(proto.runtime.Runtime.service, {
  stream: async (conn) => {

  },
  healthCheck: () => {

  }
})
