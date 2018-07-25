const grpc = require('grpc')

const proto = grpc.load('../../proto/runtime.proto')

const client = new proto.faas.Runtime('localhost:9000', grpc.credentials.createInsecure())

const conn = client.stream({
  version: 1,
  id: 'test',
  packet: {
    topic: 'topic',
    event: 'event',
    id: 1,
    payload: Buffer.from('test message')
  }
})

conn.on('data', (data) => {
  console.log('---')
  console.log(data.version)
  console.log(data.cmd)
  console.log(data.packet.topic)
  console.log(data.packet.event)
  console.log(data.packet.id)
  console.log(data.packet.payload.toString('utf8'))
})
