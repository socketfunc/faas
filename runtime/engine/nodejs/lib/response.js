class Response {
  constructor (conn) {
    this._conn = conn
  }

  static get CMD () {
    return {
      SEND: 0,
      REPLY: 1,
      BROADCAST: 2
    }
  }

  _send (cmd, topic, event, id, payload) {
    if (!this._conn.writable) {
      return
    }
    const version = 1
    this._conn.write({version, cmd, packet: {topic, event, id, payload}})
  }

  /**
   *
   * @param topic
   * @param event
   * @param id
   * @param payload
   */
  send ({topic, event, id, payload}) {
    this._send(Response.CMD.SEND, topic, event, id, payload)
  }

  close () {
    if (!this._conn.writable) {
      return
    }
    this._conn.end()
  }
}

module.exports = Response
