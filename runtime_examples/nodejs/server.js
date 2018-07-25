const server = require('socketfunc')

async function handler (req, res) {
  res.send({
    topic: 'topic',
    event: 'event',
    id: 100,
    payload: Buffer.from('test message 01')
  })
  await timeout(1000)
  res.send({
    topic: 'topic',
    event: 'event',
    id: 101,
    payload: Buffer.from('test message 02')
  })
  await timeout(1000)
  res.send({
    topic: 'topic',
    event: 'event',
    id: 102,
    payload: Buffer.from('test message 03')
  })
}

server.start(handler)

function timeout (ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}
